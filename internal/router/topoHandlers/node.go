package topoHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	v2 "goOrigin/internal/logic/v2"
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
	entity.ParentID = req.ParentId
	entity.Done = req.Done
	entity.Status = "New"
	entity.Note = req.Note
	entity.Tags = req.Tags
	res, err = v2.CreateNode(c, req.Region, entity)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func CreateNodes(c *gin.Context) {
	var (
		req      = V1.CreateNodesRequest{}
		res      interface{}
		err      error
		entities []*entity.NodeEntity
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	for _, info := range req.Info {
		entities = append(entities, &entity.NodeEntity{
			Name:     info.Name,
			Content:  info.Content,
			ParentID: info.ParentId,
			Done:     info.Done,
			Tags:     info.Tags,
			Note:     info.Note,
		})
	}
	res, err = v2.CreateNodes(c, entities, "")
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
	nodeEntity := &entity.NodeEntity{
		Name:     req.Name,
		Content:  req.Content,
		ParentID: req.ParentID,
		Status:   req.Status,
		Note:     req.Note,
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	res, err = v2.UpdateNode(c, req.ID, req.Region, nodeEntity)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func SearchNode(c *gin.Context) {
	var (
		name    = c.Query("name")
		region  = c.Query("region")
		content = c.Query("content")
		res     interface{}
		err     error
	)
	v2.SearchNodes(c, name, region, content)
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

func DeleteNode(c *gin.Context) {
	var (
		region = c.Query("region")
		err    error
	)
	id, err := conv.Uint(c.Query("id"))

	res, err := v2.DeleteNode(c, id, region)
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
