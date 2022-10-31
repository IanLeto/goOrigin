package service

import (
	"github.com/gin-gonic/gin"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	logger2 "goOrigin/pkg/logger"
)

func CreateNode(c *gin.Context, req *params.CreateNodeRequest) (id string, err error) {
	var (
		logger = logger2.NewLogger()
		node   *model.Node
	)
	node = &model.Node{
		Name:    req.Name,
		Content: req.Content,
		Father:  req.Father,
	}
	id, err = node.CreateNode(c)
	if err != nil {
		logger.Error("创建node 失败")
		return "", err
	}
	return
}
