package recordHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/model/entity"
)

// param $参数名 参数类型 参数数据类型 是否必须 描述 其他

// CreateIanRecord @Summary
// @Description
// @Tags RecordEntity
// @Accept json
// @param record body params.CreateIanRequestInfo true "1"
// @Router /v1/record [POST]
func CreateIanRecord(c *gin.Context) {
	var (
		req   = V1.CreateIanRequestInfo{}
		objID interface{}
		err   error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	//objID, err = logic.CreateIanRecord(c, req)
	V1.BuildResponse(c, V1.BuildInfo(objID))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteIanRecord(c *gin.Context) {
	var (
		//id  = c.GetString("id")
		res int64
		err error
	)
	//res, err = logic.DeleteIanRecord(c, id)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateIanRecord(c *gin.Context) {
	var (
		req = V1.CreateIanRequestInfo{}
		res interface{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func AppendIanRecord(c *gin.Context) {
	var (
		req = V1.AppendRequestInfo{}
		res *entity.RecordEntity
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	//res, err = logic.AppendIanRecord(c, &req)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
