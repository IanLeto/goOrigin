package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/API/V1"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
	"gorm.io/gorm"
)

func CreateNode(ctx *gin.Context, info *V1.CreateNodeInfo) (uint, error) {
	var (
		nodeEntity = &entity.NodeEntity{}
		region     = ctx.GetString("region")
	)

	// 填充字段
	nodeEntity.Name = info.Name
	nodeEntity.Content = info.Content
	//nodeEntity.DependIDs = info.Depend
	nodeEntity.ParentID = info.ParentID
	nodeEntity.Region = region
	nodeEntity.Done = info.Done
	nodeEntity.Tags = info.Tags
	nodeEntity.Note = info.Note
	nodeEntity.Status = info.Status

	// 转换为 DAO
	tNode := repository.ToTNode(nodeEntity)

	// 获取 DB 实例
	db := mysql.GlobalMySQLConns[region]
	res, _, err := mysql.Create(db.Client, tNode)
	if err != nil {
		logger.Error(fmt.Sprintf("create node failed %s: %s", err, res))
		return 0, err
	}

	return tNode.ID, nil
}

func GetNodeByID(ctx *gin.Context, id uint) (*entity.NodeEntity, error) {
	var (
		err error
	)
	region := ctx.GetString("region")
	db := mysql.GlobalMySQLConns[region]

	var tNode dao.TNode
	record := db.Client.Where("id = ?", id).First(&tNode)
	err = record.Error
	if err != nil {
		logger.Error(fmt.Sprintf("get node by id failed: %s", err))
		return nil, err
	}

	return repository.ToNodeEntity(&tNode), nil
}

func UpdateNode(ctx *gin.Context, info *V1.UpdateNodeInfo) error {
	var (
		err    error
		region = ctx.GetString("region")
		db     = mysql.GlobalMySQLConns[region]
		tNode  dao.TNode
	)

	// 查找原始数据

	record := db.Client.Where("id = ?", info.ID).First(&tNode)
	err = record.Error
	if err != nil {
		logger.Error(fmt.Sprintf("update node failed, not found: %s", err))
		return err
	}

	// 更新字段
	tNode.Name = info.Name
	tNode.Content = info.Content
	tNode.ParentID = info.ParentID
	tNode.Done = info.Done
	tNode.Note = info.Note
	tNode.Status = info.Status
	tNode.Region = region

	// Tags 处理（如需要）
	tagStr, _ := json.Marshal(info.Tags)
	tNode.Tags = string(tagStr)

	err = db.Client.Save(&tNode).Error
	if err != nil {
		logger.Error(fmt.Sprintf("update node failed: %s", err))
	}
	return err
}

func DeleteNodeByID(ctx *gin.Context, id uint) error {
	region := ctx.GetString("region")
	db := mysql.GlobalMySQLConns[region]

	err := db.Client.Where("id = ?", id).Delete(&dao.TNode{}).Error
	if err != nil {
		logger.Error(fmt.Sprintf("delete node failed: %s", err))
	}
	return err
}

func ListNodes(ctx *gin.Context, parentID uint, status string, keyword string, page, pageSize int) ([]*entity.NodeEntity, int64, error) {
	region := ctx.GetString("region")
	db := mysql.GlobalMySQLConns[region]

	var (
		tNodes []*dao.TNode
		total  int64
	)

	query := db.Client.Model(&dao.TNode{}).Where("region = ?", region)

	if parentID != 0 {
		query = query.Where("parent_id = ?", parentID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if keyword != "" {
		query = query.Where("name LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&tNodes).Error
	if err != nil {
		logger.Error(fmt.Sprintf("list nodes failed: %s", err))
		return nil, 0, err
	}

	entities := make([]*entity.NodeEntity, 0, len(tNodes))
	for _, t := range tNodes {
		entities = append(entities, repository.ToNodeEntity(t))
	}

	return entities, total, nil
}

func GetFullNodeTree(ctx *gin.Context) ([]*entity.NodeEntity, error) {
	region := ctx.GetString("region")
	db := mysql.GlobalMySQLConns[region]

	var tNodes []*dao.TNode
	err := db.Client.Where("region = ?", region).Find(&tNodes).Error
	if err != nil {
		logger.Error(fmt.Sprintf("get tree failed: %s", err))
		return nil, err
	}

	nodeMap := make(map[uint]*entity.NodeEntity)
	var allNodes []*entity.NodeEntity

	for _, t := range tNodes {
		node := repository.ToNodeEntity(t)
		node.Children = []*entity.NodeEntity{}
		allNodes = append(allNodes, node)
		nodeMap[node.ID] = node
	}

	var roots []*entity.NodeEntity
	for _, node := range allNodes {
		if node.ParentID == 0 {
			roots = append(roots, node)
		} else {
			if parent, ok := nodeMap[node.ParentID]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}

	return roots, nil
}

func NewExistEsQuery(param string, query elastic.Query) elastic.Query {
	if param == "" {
		return nil
	}
	return query
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
