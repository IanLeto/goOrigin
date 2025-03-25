package recordHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
	"goOrigin/internal/model/entity"
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
