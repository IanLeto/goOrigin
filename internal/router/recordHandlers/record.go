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
		//entity = &entity.Record{}
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	res.Id, err = logic.CreateRecord(c, req.Region, req.CreateIanRecordRequestInfo)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
func QueryRecord(c *gin.Context) {
	var (
		res    = &V1.QueryIanRecordsResponse{}
		result []*entity.Record
		err    error
	)
	startTime, _ := conv.Int64(c.Query("start_time"))
	endTime, _ := conv.Int64(c.Query("modify_time"))
	region, _ := conv.String(c.Query("region"))
	name, _ := conv.String(c.Query("name"))
	result, err = logic.QueryRecords(c, region, name, startTime, endTime)
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
		record = &entity.Record{}
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res.ID, err = logic.UpdateRecord(c, record)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
