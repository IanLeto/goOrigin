package topoHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
	"goOrigin/internal/model/entity"
)

func CreateNode(c *gin.Context) {
	var (
		req    = V1.CreateNodeRequest{}
		res    interface{}
		err    error
		entity = &entity.NodeEntity{}
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	entity.Name = req.Name
	entity.Content = req.Content
	entity.Done = req.Done
	entity.Status = "New"
	entity.Note = req.Note
	entity.Tags = req.Tags
	res, err = logic.CreateNode(c, req.CreateNodeInfo)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func GetNodeByID(c *gin.Context) {
	var (
		req = V1.GetNodeRequest{}
		res interface{}
		err error
	)

	if err = c.ShouldBindQuery(&req); err != nil {
		logrus.Errorf("bind query err: %s", err)
		goto ERR
	}

	res, err = logic.GetNodeByID(c, req.ID)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("get node failed: %s", err)))
}

func UpdateNode(c *gin.Context) {
	var (
		req = V1.UpdateNodeRequest{}
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("bind json err: %s", err)
		goto ERR
	}

	err = logic.UpdateNode(c, req.UpdateNodeInfo)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo("ok"))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("update node failed: %s", err)))
}

func DeleteNode(c *gin.Context) {
	var (
		req = V1.DeleteNodeRequest{}
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("bind json err: %s", err)
		goto ERR
	}

	err = logic.DeleteNodeByID(c, req.ID)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo("deleted"))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("delete node failed: %s", err)))
}

func GetNodeDetail(c *gin.Context) {
	var (
		region = c.Query("region")
		res    interface{}
		err    error
	)
	id, _ := conv.Uint(c.Query("id"))
	res, err = logic.GetNodeDetail(c, region, id)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func ListNodes(c *gin.Context) {
	var (
		req = V1.ListNodeRequest{}
		err error
	)

	if err = c.ShouldBindQuery(&req); err != nil {
		logrus.Errorf("bind query err: %s", err)
	}

	// 调用 logic 层
	data, total, err := logic.ListNodes(c, req.ParentID, req.Status, req.Keyword, req.Page, req.PageSize)
	if err != nil {
		goto ERR
	}

	// 返回分页结构
	V1.BuildResponse(c, V1.BuildInfo(gin.H{
		"list":  data,
		"total": total,
	}))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("list nodes failed: %s", err)))
}
