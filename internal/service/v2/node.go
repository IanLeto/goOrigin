package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goOrigin/API/V1"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

func SearchNode(c *gin.Context, region, father, keyword string) {

}

func CreateNode(c *gin.Context, req *V1.CreateNodeRequest) (id uint, err error) {
	var (
		logger = logger2.NewLogger()
		node   *model.NodeEntity
	)

	node = &model.NodeEntity{
		Name:     req.Name,
		Content:  req.Content,
		Depend:   req.Depend,
		FatherID: req.FatherId,
		Done:     req.Done,
		Region:   req.Region,
		Status:   "New",
		Note:     req.Note,
		Tags:     req.Tags,
		Children: req.Children,
	}

	id, err = model.CreateNodeAdapter(c, node, req.Region, false)
	if err != nil {
		logger.Error("创建node 失败")
		return id, err
	}
	return
}

func CreateNodes(c *gin.Context, topoInfo, region string) (interface{}, error) {
	var (
		err   error
		node  *model.NodeEntity
		nodes []*model.NodeEntity
	)
	err = json.Unmarshal([]byte(topoInfo), node)
	if err != nil {
		return nil, err
	}
	db := clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	nodes = node.ToNodes()
	record, _, err := mysql.Create(db, nodes)
	if err != nil {
		return nil, err
	}
	return record, err

}
