package trans_type

import (
	"fmt"
	"github.com/cstockton/go-conv"
	_ "github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
	"goOrigin/internal/model/entity"
	"time"
)

func CreateTransInfo(c *gin.Context) {
	var (
		req = &V1.CreateTransInfoReq{}
		res = &V1.CreateTransInfoRes{}
		err error
	)
	if req.Region == "" {
		c.Set("region", "win")
	} else {
		c.Set("region", req.Region)
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	err = logic.CreateType(c, req.Region, req.Items)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func GetTransInfoList(c *gin.Context) {
	var (
		req   = &V1.SearchTransInfoReq{} // 请求结构体
		res   = &V1.SearchTransInfoRes{} // 响应结构体
		err   error
		total int64
	)

	// 绑定 JSON 请求体
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("failed to bind JSON: %v", err)
		V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("invalid request: %s", err)))
		return
	}

	// 设置默认值
	if req.Region == "" {
		req.Region = "win"
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// 解析时间参数
	var startTime, endTime *time.Time
	if req.StartTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err != nil {
			V1.BuildErrResponse(c, V1.BuildErrInfo(0, "invalid start_time format, should be: 2006-01-02 15:04:05"))
			return
		}
		startTime = &t
	}
	if req.EndTime != "" {
		t, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err != nil {
			V1.BuildErrResponse(c, V1.BuildErrInfo(0, "invalid end_time format, should be: 2006-01-02 15:04:05"))
			return
		}
		endTime = &t
	}

	// 设置 region 到 context
	c.Set("region", req.Region)

	// 调用带时间参数的查询函数
	list, total, err := logic.SearchTransInfo(c, req.Region, req.Project, req.TransType, startTime, endTime, req.Page, req.PageSize)
	if err != nil {
		logrus.Errorf("query logic failed: %v", err)
		V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("query trans info failed: %s", err)))
		return
	}

	// 处理聚合查询
	if req.TransType != "" || req.Project != "" {
		// 转换时间为时间戳（如果 SearchUrlPathWithReturnCode 需要时间戳）
		var startTimestamp, endTimestamp int64
		if startTime != nil {
			startTimestamp = startTime.Unix()
		}
		if endTime != nil {
			endTimestamp = endTime.Unix()
		}

		reqInfo := &V1.SearchUrlPathWithReturnCodesInfo{
			Project:    req.Project,
			TransTypes: []string{},
			StartTime:  startTimestamp,
			EndTime:    endTimestamp,
		}

		if req.TransType != "" {
			reqInfo.TransTypes = []string{req.TransType}
		}

		aggs, err := logic.SearchUrlPathWithReturnCode(c, req.Region, reqInfo)
		if err != nil {
			logrus.Warnf("search url path with return code failed: %v", err)
		} else if len(aggs) > 0 {
			// 合并聚合结果（带去重）
			mergeAggregationResults(list, aggs, &total)
		}
	}

	// 构建响应
	res.Items = list
	res.Total = total
	res.Page = req.Page
	res.PageSize = req.PageSize

	V1.BuildResponse(c, V1.BuildInfo(res))
}

// 辅助函数：合并聚合结果
func mergeAggregationResults(list []*entity.TransInfoEntity, aggs []*entity.UrlPathAggEntity, total *int64) {
	// 创建已存在项的映射
	existingMap := make(map[string]bool)
	for _, item := range list {
		key := fmt.Sprintf("%s_%s", item.Project, item.TransType)
		existingMap[key] = true
	}

	// 转换并添加不重复的聚合结果
	aggList := entity.ConvertUrlPathAggListToTransInfoList(aggs)
	for _, aggItem := range aggList {
		key := fmt.Sprintf("%s_%s", aggItem.Project, aggItem.TransType)
		if !existingMap[key] {
			list = append(list, aggItem)
			*total++
		}
	}
}

func DeleteTransInfo(c *gin.Context) {
	var (
		req = &V1.DeleteTransInfoReq{}
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("bind json failed: %v", err)
		goto ERR
	}

	if req.Region == "" {
		req.Region = "win"
	}

	err = logic.DeleteTransInfo(c, req.Region, req.Project, req.TransType)
	if err != nil {
		logrus.Errorf("delete trans info failed: %v", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(nil))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("delete failed: %s", err)))
}

func UpdateTransInfo(c *gin.Context) {
	var (
		req = &V1.UpdateTransInfoReq{}
		err error
	)
	infoEntity := convertToEntity(req.Item)

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("bind json failed: %v", err)
		goto ERR
	}

	if req.Region == "" {
		req.Region = "win"
	}

	// ✨ 转换请求结构体为 infoEntity

	err = logic.UpdateTransInfo(c, req.Region, infoEntity)
	if err != nil {
		logrus.Errorf("update failed: %v", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(nil))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("update failed: %s", err)))
}

// 转换函数：将 API 层结构体转换为 entity 层结构体
func convertToEntity(item *V1.UpdateTransInfo) *entity.TransInfoEntity {
	if item == nil {
		return nil
	}

	var codes []*entity.ReturnCodeEntity
	for _, rc := range item.ReturnCodes {
		codes = append(codes, &entity.ReturnCodeEntity{
			ReturnCode: rc.ReturnCode,
			TransType:  rc.TransType,
			Project:    rc.Project,
			Status:     rc.Status,
		})
	}

	return &entity.TransInfoEntity{
		Project:     item.Project,
		TransType:   item.TransType,
		ReturnCodes: codes,
	}
}

func SearchTransTypeReturnCodes(c *gin.Context) {
	var (
		req    = &V1.SearchUrlPathWithReturnCodesReq{}
		res    = &V1.SearchUrlPathWithReturnCodesInfoResponse{}
		result []*entity.UrlPathAggEntity
		err    error
	)

	// 获取查询参数
	region, _ := conv.String(c.Query("region"))
	if region == "" {
		region = "win"
	}
	page, _ := conv.Int(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := conv.Int(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = 107
	}

	// 构建请求参数
	startTime, _ := conv.Int64(c.Query("start_time"))
	endTime, _ := conv.Int64(c.Query("end_time"))

	req.Region = region
	req.Page = page
	req.PageSize = pageSize
	req.SearchUrlPathWithReturnCodesInfo = &V1.SearchUrlPathWithReturnCodesInfo{
		Project:    c.Query("project"),
		Az:         c.Query("az"),
		TransTypes: c.QueryArray("trans_types"),
		StartTime:  startTime,
		EndTime:    endTime,
		Keyword:    c.Query("keyword"),
		OrderBy:    c.Query("order_by"),
	}

	// 执行查询
	result, err = logic.SearchUrlPathWithReturnCode(c, region, req.SearchUrlPathWithReturnCodesInfo)
	if err != nil {
		goto ERR
	}

	// 构建响应
	for _, record := range result {
		res.Items = append(res.Items, record)
	}
	res.Total = len(result)
	res.Page = page
	res.PageSize = pageSize

	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("search trans type return codes failed by %s", err)))
}
