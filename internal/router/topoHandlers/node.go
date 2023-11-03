package topoHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/service"
	v2 "goOrigin/internal/service/v2"
)

func CreateNode(c *gin.Context) {
	var (
		req = V1.CreateNodeRequest{}
		res interface{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = v2.CreateNode(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func CreateNodes(c *gin.Context) {
	var (
		req = V1.CreateNodesRequest{}
		res interface{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = v2.CreateNodes(c, req.Info, req.Region)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

//func GetNodeList(c *gin.Context) {
//	var (
//		req = V1.GetNodeListRequest{}
//		res = V1.GetNodeListResponse{}
//		err error
//	)
//	var (
//		name    = c.Search("name")
//		content = c.Search("content")
//		father  = c.Search("father")
//		region  = c.Search("region")
//	)
//
//}

func UpdateNode(c *gin.Context) {
	var (
		req = V1.UpdateNodeRequest{}
		res interface{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = v2.UpdateNode(c, &req)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func GetNodes(c *gin.Context) {
	var (
		name   = c.Query("name")
		father = c.Query("father")
		region = c.Query("region")
		err    error
	)
	res, err := v2.GetNodes(c, name, father, region)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func GetNodeDetail(c *gin.Context) {
	var (
		region = c.Query("region")
		res    interface{}
		err    error
	)
	id, _ := conv.Uint(c.Query("id"))
	res, err = v2.GetNodeDetail(c, region, id)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func GetTopo(c *gin.Context) {
	var (
		name   = c.Query("name")
		region = c.Query("region")
		res    interface{}
		err    error
	)
	id, err := conv.Uint(c.Query("id"))
	if err != nil {
		goto ERR
	}

	res, err = service.GetTopo(c, name, id, region)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
func GetTopoList(c *gin.Context) {
	var (
		res    interface{}
		err    error
		region = c.Query("region")
	)

	res, err = service.GetTopoList(c, region)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteNodes(c *gin.Context) {
	var (
		ids = c.QueryArray("ids")
		err error
	)
	res, err := service.DeleteNodes(c, ids)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
func DeleteNode(c *gin.Context) {
	var (
		err error
		res interface{}
	)
	id, _ := conv.Uint(c.Query("id"))
	region := c.Query("region")

	single := c.Query("single")
	isSingle, _ := conv.Bool(single)
	if isSingle {
		res, err = service.DeleteSingleNode(c, id, region)
	} else {
		res, err = service.DeleteNode(c, id, region)
	}
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
