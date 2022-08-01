package k8sHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
)

func CreateDeploy(c *gin.Context) {
	var (
		req  = params.CreateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.CreateDeployment(c, &req)
	params.BuildResponse(c, params.BuildInfo(name))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
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

	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func UpdateDeploy(c *gin.Context) {
	var (
		req  = params.UpdateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.UpdateDeployment(c, &req)
	params.BuildResponse(c, params.BuildInfo(name))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
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

	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func CreateConfigMap(c *gin.Context) {
	var (
		req  = params.CreateDeploymentReq{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.CreateDeployment(c, &req)
	params.BuildResponse(c, params.BuildInfo(name))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func SelectConfigMap(c *gin.Context) {
	var (
		req  = params.GetConfigMapRequestInfo{}
		name interface{}
		err  error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	name, err = service.GetConfigMapDetail(c, &req)
	params.BuildResponse(c, params.BuildInfo(name))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
