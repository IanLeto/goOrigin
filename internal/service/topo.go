package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/config"
	elastic2 "goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

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
		Status:   "New",
		Note:     req.Note,
		Tags:     req.Tags,
		Children: req.Children,
		Region:   req.Region,
	}

	id, err = node.CreateNode(c)
	if err != nil {
		logger.Error("创建node 失败")
		return 0, err
	}
	return
}

func NewExistEsQuery(param string, query elastic.Query) elastic.Query {
	if param == "" {
		return nil
	}
	return query
}

func GetNodes(c *gin.Context, id, name, father, content string, done bool) (node []*model.NodeEntity, err error) {
	var (
		logger  = logger2.NewLogger()
		queries []elastic.Query
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		daoRes  *elastic.SearchResult
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(model.EsNode) }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return nil, err
	}
	if name != "" {
		queries = append(queries, elastic.NewTermsQuery("name", name))

	}
	if father != "" {
		queries = append(queries, elastic.NewTermsQuery("father", father))

	}
	if content != "" {
		queries = append(queries, elastic.NewMatchQuery("content", content))

	}
	bq.Must(queries...)
	daoRes, err = client.Search().Index(model.EsNode).Query(bq).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("查询topo失败%s", err.Error()))
		goto ERR
	}
	for _, hit := range daoRes.Hits.Hits {
		var ephemeralNode *model.NodeEntity
		err = json.Unmarshal(hit.Source, &ephemeralNode)
		//ephemeralNode.ID, _ = conv.Uint(hit.Id)
		if err != nil {
			goto ERR
		}
		node = append(node, ephemeralNode)
	}
	return node, nil
ERR:
	{
		return nil, err
	}
}

func DeleteNode(c *gin.Context, id uint, region string) (interface{}, error) {
	var (
		err   error
		nodes []*model.NodeEntity
		root  *model.NodeEntity
	)
	db := clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	// 先找到该节点信息
	record, _, err := mysql.GetValue(db, &model.NodeEntity{ID: id}, "")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("del data failed by %s", err))
	}
	values, err := json.Marshal(record)
	if err != nil {
		goto ERR
	}
	err = json.Unmarshal(values, &root)
	if err != nil {
		goto ERR
	}
	// 正式环境，删除之前要归档到es中
	if region == "" {
		// todo es 归档
	}
	nodes = root.ToNodes()
	// 1. 查询node 对应的所有topo数据
	// todo 依赖tree结构，先搞定脚手架
	err = mysql.DeleteValues(db, nodes)
	return nil, err
ERR:
	return nil, errors.New(fmt.Sprintf("del data failed by %s", err))

}

func DeleteSingleNode(c *gin.Context, id uint, region string) (interface{}, error) {
	var (
		err  error
		node *model.NodeEntity
	)
	db := clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client

	err = mysql.DeleteValues(db, node)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("del data failed by %s", err))
	}
	return nil, err
}

func DeleteNodes(c *gin.Context, ids []string) (interface{}, error) {
	var (
		logger  = logger2.NewLogger()
		filters []elastic.Query
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		//node    *model.NodeEntity
		//daoRes  *elastic.DeleteResponse
		//err error
	)
	filters = append(filters, NewExistEsQuery("_id", elastic.NewTermsQuery("_id", ids)))
	bq.Filter(filters...)
	daoRes, err := client.DeleteByQuery().Index(model.EsNode).Query(bq).Do(c)
	if err != nil {
		logger.Error("delete 失败")
		return nil, err
	}
	return daoRes, err
}

func GetNodeDetail(c *gin.Context, id, name, father string) (interface{}, error) {
	var (
		logger  = logger2.NewLogger()
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		daoRes  *elastic.SearchResult
		queries []elastic.Query
		node    *model.NodeEntity
		err     error
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(model.EsNode) }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return nil, err
	}
	queries = append(queries, NewExistEsQuery(id, elastic.NewTermsQuery("_id", id)))
	queries = append(queries, NewExistEsQuery(name, elastic.NewTermsQuery("name", name)))
	queries = append(queries, NewExistEsQuery(father, elastic.NewTermsQuery("father", father)))
	daoRes, err = client.Search().Index("topo").Query(bq).Do(c)
	hit := daoRes.Hits.Hits[0]

	if err != nil {
		logger.Error(fmt.Sprintf("删除错误"))
		goto ERR
	}
	err = json.Unmarshal(hit.Source, &node)
	if err != nil {
		goto ERR
	}
	return node, nil

ERR:
	{
		return nil, err
	}
}

const DefaultRegion = "conn"

func GetTopo(c *gin.Context, name string, id uint, region string) (interface{}, error) {
	var (
		err  error
		root *model.NodeEntity
	)
	queryMysqlCallback := func(node *model.NodeEntity) (*model.NodeEntity, error) {
		var (
			err error
		)
		return nil, err
	}
	db := clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	record, _, err := mysql.GetValue(db, root, "node")
	root.Epl(queryMysqlCallback)
	if err != nil {
		goto ERR
	}

ERR:
	logrus.Errorf("获取topo 失败 %s", err)
	return record, err
}

func GetTopoList(c *gin.Context, region string) (res []*V1.GetTopoResponse, err error) {
	var (
		logger = logger2.NewLogger()
		node   *model.NodeEntity
		conn   *elastic2.EsV2Conn
		doc    *elastic2.EsDoc
	)
	var (
		queries map[string]interface{}
	)
	if region == "" {
		conn = elastic2.EsConns[DefaultRegion]
	} else {
		conn = elastic2.EsConns[region]
	}
	queries = map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"region.keyword": region,
			},
		},
	}
	value, err := conn.Search(config.NodeMapping, queries)
	err = json.Unmarshal(value, &doc)
	if err != nil {
		logrus.Debugf("query: %s", func() string {
			b, _ := json.Marshal(value)
			return string(b)
		}())
	}
	for _, hit := range doc.Hits.Hits {
		data, err := json.Marshal(hit.Source)
		if err != nil {
			goto ERR
		}
		err = json.Unmarshal(data, &node)
		if err != nil {
			logger.Error(fmt.Sprintf("json 错误 %s", err.Error()))
			goto ERR
		}
		res = append(res, &V1.GetTopoResponse{
			Name:    node.Name,
			Content: node.Content,
			Depend:  node.Depend,
			Done:    node.Done,
			Tags:    node.Tags,
			Note:    node.Note,
			Nodes:   node.Nodes,
		})
	}

	return

ERR:
	{
		return nil, err
	}
}

func GetTopo2(c *gin.Context, name string) (res *V1.GetTopoResponse, err error) {
	var (
		logger  = logger2.NewLogger()
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		daoRes  *elastic.SearchResult
		queries []elastic.Query
		node    *model.NodeEntity
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(model.EsNode) }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化es 失败 %s", err))
		return nil, err
	}
	queries = append(queries, NewExistEsQuery(name, elastic.NewTermsQuery("name", name)))
	bq.Filter(queries...)
	daoRes, err = client.Search().Index(model.EsNode).Query(bq).Do(c)
	if len(daoRes.Hits.Hits) == 0 {
		err = errors.New("不存在该数据")
		goto ERR
	}
	err = json.Unmarshal(daoRes.Hits.Hits[0].Source, &node)
	if err != nil {
		logger.Error(fmt.Sprintf("json 错误 %s", err.Error()))
		goto ERR
	}
	node = model.GetTopo(c, node)
	res = &V1.GetTopoResponse{
		Name:    node.Name,
		Content: node.Content,
		Depend:  node.Depend,
		Done:    node.Done,
		Tags:    node.Tags,
		Note:    node.Note,
		Nodes:   node.Nodes,
	}
	return

ERR:
	{
		return nil, err
	}
}
