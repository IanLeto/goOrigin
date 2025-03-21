package k8sHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
)

func DeleteDeploy(c *gin.Context) {
	var (
		name = c.GetString("name")
		ns   = c.GetString("namespace")
		res  string
		err  error
	)
	res, err = logic.DeleteDeployment(c, name, ns)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateDeploy(c *gin.Context) {
	var (
		req  = V1.UpdateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = logic.UpdateDeployment(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func ListDeploy(c *gin.Context) {
	var (
		ns  = c.GetString("namespace")
		res interface{}
		err error
	)
	res, err = logic.ListDeployments(c, ns)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func CreateDeployDynamic(c *gin.Context) {
	var (
		req  = V1.CreateDeploymentDynamicRequest{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = logic.CreateDeploymentDynamic(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteDeployDynamic(c *gin.Context) {
	var (
		name = c.GetString("name")
		ns   = c.GetString("namespace")
		res  string
		err  error
	)
	err = logic.DeleteDeploymentDynamic(c, name, ns)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateDeployDynamic(c *gin.Context) {
	var (
		req  = V1.UpdateDeploymentDynamicRequest{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = logic.UpdateDeploymentDynamicRequest(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
