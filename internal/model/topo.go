package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

type Node struct {
	ID       string   `json:"_id"`
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	Depend   string   `json:"depend"`
	Father   string   `json:"father"`
	FatherID string   `json:"father_id"`
	Done     bool     `json:"done"`
	Status   string   `json:"status"`
	Note     string   `json:"note"`
	Tags     []string `json:"tags"`
	Children []string `json:"children"`
	Nodes    []*Node  `json:"nodes"`
}

type Topo struct {
	*Node
	Children []*Node `json:"children"`
}

func GetTopo(ctx context.Context, root *Node) *Node {
	var (
		client *elastic.Client
		daoRes *elastic.SearchResult
		bq     = elastic.NewBoolQuery()
		err    error
	)
	client, err = clients.NewESClient()
	defer func() { _ = client.CloseIndex("") }()
	if err != nil {

	}

	bq.Filter(elastic.NewTermQuery("father", root.Name))
	daoRes, err = client.Search().Index(EsNode).Query(bq).Do(ctx)
	for _, hit := range daoRes.Hits.Hits {
		var (
			ephemeralNode Node
		)
		err = json.Unmarshal(hit.Source, &ephemeralNode)
		root.Nodes = append(root.Nodes, &ephemeralNode)
		GetTopo(ctx, &ephemeralNode)
	}
	if err != nil {
		return nil
	}
	return root
}

func (node *Node) CreateNode(c *gin.Context) (id string, err error) {
	var (
		client *elastic.Client
		res    *elastic.IndexResponse
		logger = logger2.NewLogger()
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(EsNode) }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return "", err
	}
	res, err = client.Index().Index(EsNode).BodyJson(node).Do(c)
	if err != nil {
		goto ERR
	}
	return res.Id, nil
ERR:
	{
		return "", err
	}
}
