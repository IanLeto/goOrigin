package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
	"gorm.io/gorm"
)

func CreateNode(c *gin.Context, region string, entity *entity.NodeEntity) (id uint, err error) {
	var (
		logger = logger2.NewLogger()
	)
	tNode := repository.ToTnode(entity)
	record, _, err := mysql.Create(mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client, tNode)
	logger.Debug(fmt.Sprintf("create node %s", record))
	if err != nil {
		logger.Error("创建node 失败")
		return id, err
	}
	return
}

func CreateNodes(c *gin.Context, nodes []*entity.NodeEntity, region string) (interface{}, error) {
	var (
		err error
	)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, err
	}
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client
	record, _, err := mysql.Create(db, nodes)
	if err != nil {
		return nil, err
	}
	return record, err
}

func UpdateNode(c *gin.Context, id uint, region string, nodeUpdate *entity.NodeEntity) (interface{}, error) {
	var (
		logger = logger2.NewLogger()
		err    error
	)
	nodeEntity, err := GetNodeDetail(c, region, id)
	record, _, err := mysql.Create(mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client, repository.ToTnode(nodeEntity))
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

func GetNodeDetail(c *gin.Context, region string, id uint) (*entity.NodeEntity, error) {
	var (
		db     *gorm.DB
		tNode  = &dao.TNode{}
		result *entity.NodeEntity
	)
	tNode.ID = id
	db = mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client
	_, err := mysql.GetValueByID(db, tNode)
	if err != nil {
		goto ERR
	}
	result = repository.ToNodeEntity(tNode)
	return result, err
ERR:
	{
		return nil, err
	}
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
func DeleteNode(c *gin.Context, id uint, region string) (interface{}, error) {
	var (
		err   error
		tNode = &dao.TNode{}
	)
	tNode.ID = id
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig).Client
	err = mysql.DeleteValue(db, tNode)
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

func SearchNodeDetail(c *gin.Context, id, name, father string) (interface{}, error) {
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
