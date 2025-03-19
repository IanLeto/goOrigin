package logic

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
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

func NewExistEsQuery(param string, query elastic.Query) elastic.Query {
	if param == "" {
		return nil
	}
	return query
}

func GetNodes(c *gin.Context, name string, status string, parentId int, startTime, endTime int64, pageSize, page int) ([]*entity.NodeEntity, error) {
	var (
		nodeEntities = make([]*dao.TNode, 0)
		err          error
		res          = make([]*entity.NodeEntity, 0)
		region       = c.GetString("region")
	)
	db := mysql.GlobalMySQLConns[region]
	sql := db.Client.Debug().Table("t_nodes")
	if name != "" {
		sql = sql.Where("name = ?", name)
	}
	if startTime != 0 {
		sql = sql.Where("create_time > ?", startTime)
	}
	if endTime != 0 {
		sql = sql.Where("modify_time < ?", endTime)
	}
	// 添加查询条件
	if name != "" {
		sql = sql.Where("name = ?", name)
	}
	sql = sql.Where("parent_id = ?", parentId)
	if pageSize == 0 {
		pageSize = 50
	}

	// 分页查询
	tRecords := sql.Order("create_time DESC").Limit(pageSize).Find(&nodeEntities)
	if tRecords.Error != nil {
		logrus.Errorf("query records failed: %s", tRecords.Error)
		goto ERR
	}

	// 转换数据
	for _, nodeEntity := range nodeEntities {
		res = append(res, repository.ToNodeEntity(nodeEntity))
	}

	return res, nil

ERR:
	return nil, err
}

func SearchNodes(c *gin.Context, id, name, father, content string, done bool) (node []*entity.NodeEntity, err error) {
	var (
		logger  = logger2.NewLogger()
		queries []elastic.Query
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		daoRes  *elastic.SearchResult
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(entity.EsNode) }()
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
	daoRes, err = client.Search().Index(entity.EsNode).Query(bq).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("查询topo失败%s", err.Error()))
		goto ERR
	}
	for _, hit := range daoRes.Hits.Hits {
		var ephemeralNode *entity.NodeEntity
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
		nodes []*entity.NodeEntity
		root  *entity.NodeEntity
	)
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client
	// 先找到该节点信息
	record, _, err := mysql.GetValue(db, &entity.NodeEntity{ID: id}, "")
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
		node *entity.NodeEntity
	)
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client

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
	daoRes, err := client.DeleteByQuery().Index(entity.EsNode).Query(bq).Do(c)
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
		node    *entity.NodeEntity
		err     error
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(entity.EsNode) }()
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
		root *entity.NodeEntity
	)
	queryMysqlCallback := func(node *entity.NodeEntity) (*entity.NodeEntity, error) {
		var (
			err error
		)
		return nil, err
	}
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client
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
		node   *entity.NodeEntity
		conn   *elastic2.EsV2Conn
		doc    *dao.EsDoc
	)
	var (
		queries map[string]interface{}
	)
	if region == "" {
		conn = elastic2.GlobalEsConns[DefaultRegion]
	} else {
		conn = elastic2.GlobalEsConns[region]
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
		node    *entity.NodeEntity
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(entity.EsNode) }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化es 失败 %s", err))
		return nil, err
	}
	queries = append(queries, NewExistEsQuery(name, elastic.NewTermsQuery("name", name)))
	bq.Filter(queries...)
	daoRes, err = client.Search().Index(entity.EsNode).Query(bq).Do(c)
	if len(daoRes.Hits.Hits) == 0 {
		err = errors.New("不存在该数据")
		goto ERR
	}
	err = json.Unmarshal(daoRes.Hits.Hits[0].Source, &node)
	if err != nil {
		logger.Error(fmt.Sprintf("json 错误 %s", err.Error()))
		goto ERR
	}
	node = entity.GetTopo(c, node)
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
