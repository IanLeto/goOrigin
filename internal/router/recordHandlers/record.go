package recordHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/service"
)

func CreateRecord(c *gin.Context) {
	var (
		req = &V1.CreateIanRecordRequest{}
		res = &V1.CreateIanRecordResponse{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = service.CreateIanRecordV2(c, req)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
func SelectRecord(c *gin.Context) {
	var (
		region = c.GetString("region")
		name   = c.GetString("name")
		res    = &V1.SelectIanRecordResponse{}
		err    error
	)
	startTime, _ := conv.Int64(c.Query("start_time"))
	endTime, _ := conv.Int64(c.Query("modify_time"))
	res, err = service.SelectIanRecordsV2(c, region, name, startTime, endTime, 0)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
