package topoHandlers

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
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
	entity.Depend = req.Depend
	entity.ParentID = req.ParentId
	entity.Done = req.Done
	entity.Region = req.Region
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
			Depend:   info.Depend,
			ParentID: info.ParentId,
			Done:     info.Done,
			Tags:     info.Tags,
			Note:     info.Note,
			Region:   info.Region,
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
		Depend:   req.Depend,
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

	res, err = logic.GetTopo(c, name, id, region)
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

	res, err = logic.GetTopoList(c, region)
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
	res, err := logic.DeleteNodes(c, ids)
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

func EDeleteNode(c *gin.Context) {
	var (
		err error
		res interface{}
	)
	id, _ := conv.Uint(c.Query("id"))
	region := c.Query("region")

	single := c.Query("single")
	isSingle, _ := conv.Bool(single)
	if isSingle {
		res, err = logic.DeleteSingleNode(c, id, region)
	} else {
		res, err = logic.DeleteNode(c, id, region)
	}
	if err != nil {
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
