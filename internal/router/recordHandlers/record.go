package recordHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/params"
)

func CreateRecord(c *gin.Context) {
	var (
		req = params.CreateRecordReqInfo{}
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
