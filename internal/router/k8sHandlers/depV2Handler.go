package k8sHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
)

func CreateV2Deploy(c *gin.Context) {
	var (
		req  = params.CreateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.CreateDeploymentV2(c, &req)
	params.BuildResponse(c, params.BuildInfo(name))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
