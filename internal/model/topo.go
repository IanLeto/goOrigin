package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
	"goOrigin/pkg/utils"
)

type NodeEntity struct {
	ID       string        `json:"_id"`
	Name     string        `json:"name"`
	Content  string        `json:"content"`
	Depend   string        `json:"depend"`
	Father   string        `json:"father"`
	FatherID string        `json:"father_id"`
	Done     bool          `json:"done"`
	Status   string        `json:"status"`
	Note     string        `json:"note"`
	Tags     []string      `json:"tags"`
	Children []string      `json:"children"`
	Nodes    []*NodeEntity `json:"nodes"`
}

type Topo struct {
	*NodeEntity
	Children []*NodeEntity `json:"children"`
}

func GetTopo(ctx context.Context, root *NodeEntity) *NodeEntity {
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
			ephemeralNode NodeEntity
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

func (node *NodeEntity) CreateNode(c *gin.Context) (id string, err error) {
	var (
		conn   *clients.EsV2Conn
		res    *elastic.IndexResponse
		father *NodeEntity
		logger = logger2.NewLogger()
	)
	conn = clients.NewEsV2Conn(nil)
	_, err = conn.Client.Info()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return "", err
	}
	var (
		query = map[string]interface{}{}
		boolq = map[string]interface{}{}
	)
	var (
		getFather = func() {
			conn.Query("topo", map[string]interface{}{
				"bool": map[string]interface{}{
					"term": map[string]interface{}{
						"_id": node.FatherID,
					},
				},
			})
		}
	)

	//if node.FatherID != "" {

	//
	//}
	// 说明是root节点
	if node.FatherID == "" {
		goto Create
	}

	//_, err = conn.Query("", query)
	//if err != nil {
	//	goto ERR
	//}
Create:
	_, err = conn.Create(nil)
	if err != nil {
		goto ERR
	}
	node.Tags = utils.Set(append(node.Tags, father.Tags...))
	node.Father = father.Name
	return res.Id, nil
ERR:
	{
		return "", err
	}
}

func (node *NodeEntity) UpdateNode(c *gin.Context) (id string, err error) {
	var (
		conn   *clients.EsConn
		res    *elastic.IndexResponse
		father *NodeEntity
		logger = logger2.NewLogger()
	)
	conn, err = clients.NewEsConn(nil)
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return "", err
	}
	_, err = conn.Update(nil)
	if err != nil {
		goto ERR
	}
	node.Tags = utils.Set(append(node.Tags, father.Tags...))
	node.Father = father.Name
	return res.Id, nil
ERR:
	{
		return "", err
	}
}
