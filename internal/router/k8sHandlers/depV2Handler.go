package k8sHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/service"
)

func CreateV2Deploy(c *gin.Context) {
	var (
		req  = V1.CreateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.CreateDeploymentV2(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
