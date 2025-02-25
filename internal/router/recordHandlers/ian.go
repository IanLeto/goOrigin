package recordHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/model/entity"
)

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
