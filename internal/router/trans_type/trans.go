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
)

func CreateTransInfo(c *gin.Context) {
	var (
		req = &V1.CreateTransInfoReq{}
		res = &V1.CreateTransInfoResponse{}
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
		req   = &V1.GetTransInfoListReq{}      // 请求结构体
		res   = &V1.GetTransInfoListResponse{} // 响应结构体
		err   error
		total int64
	)
	// 调用逻辑层查询函数
	var list []*entity.TransInfoEntity
	var aggs []*entity.UrlPathAggEntity
	var reqInfo = &V1.SearchUrlPathWithReturnCodesInfo{}

	// 设置默认 region
	if req.Region == "" {
		c.Set("region", "win")
	} else {
		c.Set("region", req.Region)
	}

	// 绑定 JSON 请求体
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("failed to bind JSON: %v", err)
		goto ERR
	}

	list, total, err = logic.QueryTransTypeList(c, req.Region, req.Project, req.TransType, req.Page, req.PageSize)
	if err != nil {
		logrus.Errorf("query logic failed: %v", err)
		goto ERR
	}
	reqInfo.Project = req.Project
	reqInfo.TransTypes = []string{req.TransType}
	reqInfo.StartTime = 0 // 可以根据实际需要设置
	reqInfo.EndTime = 0   // 可以根据实际需要设置
	//reqInfo.Keyword = req.Keyword
	//reqInfo.OrderBy = req.OrderBy
	aggs, err = logic.SearchUrlPathWithReturnCode(c, req.Region, reqInfo)
	list = append(list, entity.ConvertUrlPathAggListToTransInfoList(aggs)...)

	res.Items = list
	res.Total = total
	res.Page = req.Page
	res.PageSize = req.PageSize
	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("query trans info failed: %s", err)))
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
			ProjectID:  rc.Project,
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
	startTime, _ := conv.Int(c.Query("start_time"))
	endTime, _ := conv.Int(c.Query("end_time"))

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
