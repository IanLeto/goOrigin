package v2

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

func CreateNode(c *gin.Context, region string, entity *model.NodeEntity) (id uint, err error) {
	var (
		logger = logger2.NewLogger()
	)

	id, err = model.CreateNodeAdapter(c, entity, region, false)
	tNode, err := entity.ToMySQLTable()
	record, _, err := mysql.Create(mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client, tNode)
	logger.Debug(fmt.Sprintf("create node %s", record))
	if err != nil {
		logger.Error("创建node 失败")
		return id, err
	}
	return
}

func CreateNodes(c *gin.Context, nodes []*model.NodeEntity) (interface{}, error) {
	var (
		err error
	)
	if err != nil {
		return nil, err
	}
	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[""]).Client
	record, _, err := mysql.Create(db, nodes)
	if err != nil {
		return nil, err
	}
	return record, err
}

func UpdateNode(c *gin.Context, id uint, region string, nodeUpdate *model.NodeEntity) (interface{}, error) {
	var (
		logger = logger2.NewLogger()
		err    error
	)
	nodeEntity, err := GetNodeDetail(c, region, id)
	nodeEntity.MergeWith(nodeUpdate)
	tNode, err := nodeEntity.ToMySQLTable()
	record, _, err := mysql.Create(mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client, tNode)
	if err != nil {
		logger.Error("创建node 失败")
		return "", err
	}
	return record, nil
}

func NewExistEsQuery(param string, query elastic.Query) elastic.Query {
	if param == "" {
		return nil
	}
	return query
}

func SearchNodes(ctx *gin.Context, name string, region string, content string) {

}

func GetNodeDetail(c *gin.Context, region string, id uint) (*model.NodeEntity, error) {
	var (
		db     *gorm.DB
		tNode  = &mysql.TNode{}
		result *model.NodeEntity
	)
	tNode.ID = id
	db = mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	_, err := mysql.GetValueByID(db, tNode)
	if err != nil {
		goto ERR
	}
	result = model.NewNodeEntityFromTnode(tNode)
	return result, err
ERR:
	{
		return nil, err
	}
}

func GetNodes(c *gin.Context, name, father, region string) (node []*model.NodeEntity, err error) {
	return model.GetNodeAdapter(c, name, father, region)
	//var (
	//	logger  = logger2.NewLogger()
	//	queries = map[string]interface{}{}
	//	conn    *clients.EsV2Conn
	//)
	//conn = clients.NewEsV2Conn(config.Conf)
	//
	//if name != "" {
	//	queries["term"] = map[string]interface{}{
	//		"name": name,
	//	}
	//
	//}
	//res, err := conn.Search("node", queries)
	//
	//if err != nil {
	//	logger.Error(fmt.Sprintf("查询topo失败%s", err.Error()))
	//	goto ERR
	//}
	//
	//for _, hit := range res.Hits.Hits {
	//	var ephemeralNode *model.NodeEntity
	//	data, err := json.Marshal(hit.Source)
	//	if err != nil {
	//		goto ERR
	//	}
	//	err = json.Unmarshal(data, &ephemeralNode)
	//	ephemeralNode.ID = hit.Id
	//	if err != nil {
	//		goto ERR
	//	}
	//	node = append(node, ephemeralNode)
	//}
	//ERR:
	//	{
	//		return nil, err
	//	}
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

func SearchNodeDetail(c *gin.Context, id, name, father string) (interface{}, error) {
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
