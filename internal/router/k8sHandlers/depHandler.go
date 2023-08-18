package k8sHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/service"
)

func CreateDeploy(c *gin.Context) {
	var (
		req  = V1.CreateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.CreateDeployment(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteDeploy(c *gin.Context) {
	var (
		name = c.GetString("name")
		ns   = c.GetString("namespace")
		res  string
		err  error
	)
	res, err = service.DeleteDeployment(c, name, ns)
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
	name, err = service.UpdateDeployment(c, &req)
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
	res, err = service.ListDeployments(c, ns)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func CreateConfigMap(c *gin.Context) {
	var (
		req  = V1.CreateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.CreateDeployment(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func SelectConfigMap(c *gin.Context) {
	var (
		req  = V1.GetConfigMapRequestInfo{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.GetConfigMapDetail(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
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
	name, err = service.CreateDeploymentDynamic(c, &req)
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
	err = service.DeleteDeploymentDynamic(c, name, ns)
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
	name, err = service.UpdateDeploymentDynamicRequest(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(name))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
