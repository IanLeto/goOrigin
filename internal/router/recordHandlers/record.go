package recordHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
	"goOrigin/internal/model/entity"
	"time"
)

func CreateRecord(c *gin.Context) {
	var (
		req = &V1.CreateIanRecordRequest{}
		res = &V1.CreateIanRecordResponse{}
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

	res.Id, err = logic.CreateRecord(c, req.CreateIanRecordRequestInfo)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func CreateFileRecord(c *gin.Context) {
	var (
		req = &V1.CreateIanRecordRequest{}
		res = &V1.CreateIanRecordResponse{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	res.Id, err = logic.CreateFileRecord(c, req.CreateIanRecordRequestInfo)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
func QueryRecords(c *gin.Context) {
	var (
		res    = &V1.QueryIanRecordsResponse{}
		result []*entity.RecordEntity
		err    error
	)

	startTime, _ := conv.Int64(c.Query("start_time"))
	endTime, _ := conv.Int64(c.Query("modify_time"))
	region, _ := conv.String(c.Query("region"))
	if region == "" {
		region = "win"
	}
	name, _ := conv.String(c.Query("name"))
	page, _ := conv.Int(c.Query("page"))
	pageSize, _ := conv.Int(c.Query("page_size"))
	result, err = logic.QueryRecords(c, region, name, startTime, endTime, pageSize, page)
	if err != nil {
		goto ERR
	}
	for _, record := range result {
		res.Items = append(res.Items, record)
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateRecord(c *gin.Context) {
	var (
		req    = &V1.UpdateIanRecordRequest{}
		res    = &V1.UpdateIanRecordResponse{}
		err    error
		record = &entity.RecordEntity{}
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	// 将请求体转换为记录实体
	record.ID = req.ID

	// 检查是否有CreateIanRecordRequestInfo数据
	if req.CreateIanRecordRequestInfo != nil {
		record.Title = req.Title
		record.MorWeight = req.MorWeight
		record.NigWeight = req.NigWeight
		record.IsFuck = req.IsFuck
		record.Vol1 = req.Vol1
		record.Vol2 = req.Vol2
		record.Vol3 = req.Vol3
		record.Vol4 = req.Vol4
		record.Content = req.Content
		record.Cost = req.Cost
		record.Coding = req.Coding
		record.Social = req.Social
	}

	// 更新修改时间
	record.ModifyTime = time.Now().Unix()

	res.ID, err = logic.UpdateRecord(c, req.Region, record)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

// DeleteRecord 处理 HTTP 请求
func DeleteRecord(c *gin.Context) {
	var (
		//res = map[string]string
		err error
	)

	// 获取 ID 和 region
	recordIDStr := c.Query("id")
	region := c.Query("region")
	if region == "" {
		region = "win"

	}
	// 转换 ID
	recordID, err := conv.Uint(recordIDStr)
	if err != nil || recordID <= 0 {
		logrus.Errorf("invalid record ID: %s", recordIDStr)
		goto ERR
	}

	// 调用删除逻辑
	err = logic.DeleteRecord(c, region, uint(recordID))
	if err != nil {
		logrus.Errorf("delete record failed: %s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(recordID))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("delete record failed: %s", err)))
}

func SearchSuccessRate(c *gin.Context) {
	var (
		res    = &V1.SuccessRateResponse{}
		result *entity.SpanEntity
		err    error
	)

	// 从请求中提取参数
	region, _ := conv.String(c.Query("region"))
	if region == "" {
		region = "default"
	}
	project, _ := conv.String(c.Query("project"))
	transTypes := c.QueryArray("trans_types")
	startTime, _ := conv.Int64(c.Query("start_time"))
	endTime, _ := conv.Int64(c.Query("end_time"))

	// 构造参数结构体
	reqInfo := &V1.SuccessRateReqInfo{
		Project:    project,
		TransTypes: transTypes,
		StartTime:  startTime,
		EndTime:    endTime,
		Region:     region,
	}

	// 调用核心逻辑
	result, err = logic.SearchTransTypeSuccessStatsMetric(c, region, reqInfo)
	if err != nil {
		goto ERR
	}

	// 数据格式转换为响应结构体 SuccessRateItem
	for _, stat := range result.Stats {
		res.Items = append(res.Items, &V1.SuccessRateItem{
			TransType:     stat.TransType,
			TransTypeCn:   stat.TransTypeCN,
			SuccessCount:  int(stat.SuccessCount),
			FailedCount:   int(stat.FailedCount),
			UnknownCount:  int(stat.UnknownCount),
			Total:         int(stat.Total),
			ResponseCount: int(stat.Total), // 可自定义逻辑
		})
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("query failed: %s", err)))
}
func QueryTransTypeReturnCodes(c *gin.Context) {
	var (
		req  = &V1.TransTypeQueryInfo{}
		resp *entity.TransTypeResponseEntity
		err  error
	)

	// 解析请求参数
	if err := c.ShouldBindQuery(req); err != nil {
		V1.BuildErrResponse(c, V1.BuildErrInfo(1001, "参数解析失败"))
		return
	}
	if req.Region == "" {
		req.Region = "default"
	}

	// 调用逻辑层
	resp, err = logic.QueryTransTypeWithReturnCodesInfo(c, req)
	if err != nil {
		V1.BuildErrResponse(c, V1.BuildErrInfo(1002, fmt.Sprintf("查询失败: %s", err)))
		return
	}

	V1.BuildResponse(c, V1.BuildInfo(resp))
}
