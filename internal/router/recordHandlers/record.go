package recordHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
	"goOrigin/internal/model/entity"
)

func CreateRecord(c *gin.Context) {
	var (
		req    = &V1.CreateIanRecordRequest{}
		res    = &V1.CreateIanRecordResponse{}
		err    error
		record = &entity.Record{}
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	record.Name = req.Name
	record.Weight = req.Weight
	record.Vol1 = req.Vol1
	record.Vol2 = req.Vol2
	record.Vol3 = req.Vol3
	record.Vol4 = req.Vol4
	record.Content = req.Content
	record.Region = req.Region
	record.Retire = req.Retire
	record.Cost = req.Cost
	res.Id, err = logic.CreateRecord(c, record)
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
		//region = c.Query("region")
		//name   = c.Query("name")
		res = &V1.SelectIanRecordResponse{}
		err error
	)
	//startTime, _ := conv.Int64(c.Query("start_time"))
	//endTime, _ := conv.Int64(c.Query("modify_time"))
	//res, err = logic.QueryIanRecordsV2(c, region, name, startTime, endTime, 0)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateRecord(c *gin.Context) {
	var (
		req = &V1.UpdateIanRecordRequest{}
		res = &V1.UpdateIanRecordResponse{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	//res, err = logic.UpdateIanRecordsV2(c, req)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
