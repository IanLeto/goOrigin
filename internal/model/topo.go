package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	elastic2 "goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

type NodeEntity struct {
	ID       uint          `json:"id"`
	Name     string        `json:"name"`
	Content  string        `json:"content"`
	Depend   string        `json:"depend"`
	Father   string        `json:"father"`
	FatherID uint          `json:"father_id"`
	Done     bool          `json:"done"`
	Status   string        `json:"status"`
	Tags     []string      `json:"tags"`
	Note     string        `json:"note"`
	Region   string        `json:"region"`
	Children []string      `json:"children"`
	Nodes    []*NodeEntity `json:"nodes"`
}

func NewNodeEntityFromTnode(node *TNode) *NodeEntity {
	var (
		tags []string
	)
	value, err := json.Marshal(node.Tags)
	if err != nil {
		logrus.Errorf("处理tag 失败 %s", err)
		return nil
	}
	err = json.Unmarshal(value, &tags)
	if err != nil {
		logrus.Errorf("处理tag 失败 %s", err)
		return nil
	}
	return &NodeEntity{
		ID:       node.ID,
		Name:     node.Name,
		Content:  node.Content,
		Depend:   node.Depend,
		Father:   node.Father,
		FatherID: node.FatherID,
		Done:     node.Done,
		Status:   node.Status,
		Tags:     tags,
		Note:     node.Note,
		Region:   node.Region,
		//Children: node.Children,
		//Nodes:    node.Nodes,
	}
}

func (n *NodeEntity) ToTNode() *TNode {
	var node *TNode
	value, err := json.Marshal(n.Tags)
	if err != nil {
		return nil
	}
	node = &TNode{
		Name:     n.Name,
		Content:  n.Content,
		Depend:   n.Depend,
		Father:   n.Father,
		FatherID: n.FatherID,
		Done:     n.Done,
		Status:   n.Status,
		Region:   n.Region,
		Note:     n.Note,
		Tags:     string(value),
	}
	return node
}

type Topo struct {
	*NodeEntity
	Children []*NodeEntity `json:"children"`
}

func GetTopo(ctx context.Context, root *NodeEntity) *NodeEntity {
	return nil
}

func CreateNodeAdapter(c *gin.Context, node *NodeEntity, region string, sync bool) (id uint, err error) {
	var (
		db *gorm.DB
	)
	tNode := node.ToTNode()
	db = clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	res, _, err := mysql.Create(db, &tNode)
	if sync {
		return node.CreateNode(c)
	}
	fmt.Println(res)
	return tNode.ID, err

}

func GetNodeAdapter(c *gin.Context, name, father, region string) ([]*NodeEntity, error) {
	var (
		db    *gorm.DB
		res   []*NodeEntity
		dbRes []*TNode
	)
	tNode := TNode{
		Name:   name,
		Father: father,
	}
	db = clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	data, _, err := mysql.GetValues(db, &tNode)
	if err != nil {
		logrus.Errorf("get node failed by %s", err)
		return nil, err
	}
	values, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("get node failed by %s", err)
		return nil, err
	}
	err = json.Unmarshal(values, &dbRes)
	if err != nil {
		goto ERR
	}
	for _, v := range dbRes {
		bytes, err := json.Marshal(v.Tags)
		if err != nil {
			goto ERR
		}
		ephemeralTags := []string{}
		err = json.Unmarshal(bytes, &ephemeralTags)
		if err != nil {
			goto ERR
		}
		res = append(res, &NodeEntity{
			ID:       v.ID,
			Name:     v.Name,
			Content:  v.Content,
			Depend:   v.Depend,
			Father:   v.Father,
			FatherID: v.FatherID,
			Done:     v.Done,
			Status:   v.Status,
			Tags:     ephemeralTags,
			Note:     v.Note,
			Region:   v.Region,
		})
	}
	return res, err
ERR:
	return nil, err
}

func (node *NodeEntity) CreateNode(c *gin.Context) (id uint, err error) {
	var (
		conn   *elastic2.EsV2Conn
		father *NodeEntity
		logger = logger2.NewLogger()
	)
	conn = elastic2.EsConns[node.Region]
	_, err = conn.Client.Info()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return 0, err
	}
	var (
		query = map[string]interface{}{}
	)
	var (
		doc                   *elastic2.EsDoc
		insertResultInfo      *elastic2.InsertResultInfo
		insertResultInfoValue []byte

		insertInfo      map[string]interface{}
		insertInfoValue []byte
		value           []byte
		source          []byte
	)
	switch {
	case node.Father != "":
		query = map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"term": map[string]interface{}{
						"ToTNode": node.Father,
					},
				},
			},
		}
		goto Query
	case node.FatherID != 0:
		query = map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"term": map[string]interface{}{
						"_id": node.ID,
					},
				},
			},
		}
		goto Query
	default:
		goto Create
	}

Query:
	logrus.Debugf("query: %s", func() string {
		b, _ := json.Marshal(query)
		return string(b)
	}())
	value, err = conn.Query(config.NodeMapping, query)
	if err != nil {
		goto ERR
	}
	err = json.Unmarshal(value, &doc)
	if err != nil {
		logrus.Debugf("query: %s", func() string {
			b, _ := json.Marshal(value)
			return string(b)
		}())
	}
	if doc.Hits.Total.Value == 0 {
		err = errors.New("father node not found")
		goto Create
	}
	source, err = json.Marshal(doc.Hits.Hits[0].Source)
	if err != nil {
		goto ERR
	}
	err = json.Unmarshal(source, &father)

	if err != nil {
		goto ERR
	}
	node.Father = father.Name
	node.FatherID = father.ID
	if source == nil {
		err = errors.New("father node not found")
		goto ERR
	}
	goto Create

Create:
	insertInfoValue, err = json.Marshal(node)
	if err != nil {
		goto ERR
	}
	err = json.Unmarshal(insertInfoValue, &insertInfo)
	if err != nil {
		goto ERR
	}
	insertResultInfoValue, err = conn.Create("node", insertInfoValue)
	if err != nil {
		goto ERR
	}
	err = json.Unmarshal(insertResultInfoValue, &insertResultInfo)
	id, _ = conv.Uint(insertResultInfo.Id)
	return id, err
ERR:
	{
		return 0, err
	}
}
