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
		id  int
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	if id ,err = service.CreateRecord(req);err != nil{
		goto ERR

	}
	params.BuildResponse(c, params.BuildInfo(res))
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
