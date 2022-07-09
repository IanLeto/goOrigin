package recordHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
)

func CreateIanRecord(c *gin.Context) {
	var (
		req = params.CreateIanRequestInfo{}
		//res   = params.CreatRecordResInfo{}
		objID interface{}
		//objID string
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	objID, err = service.CreateIanRecord(c, req)
	params.BuildResponse(c, params.BuildInfo(objID))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteIanRecord(c *gin.Context) {
	var (
		id  = c.GetString("id")
		res int64
		err error
	)
	res, err = service.DeleteIanRecord(c, id)
	if err != nil {
		goto ERR
	}

	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateIanRecord(c *gin.Context) {
	var (
		req = params.CreateIanRequestInfo{}
		res = params.CreatRecordResInfo{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
