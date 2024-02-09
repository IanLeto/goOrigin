package v2

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"goOrigin/API/V1"
	"goOrigin/config"
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
	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client
	nodes = node.ToNodes()
	record, _, err := mysql.Create(db, nodes)
	if err != nil {
		return nil, err
	}
	return record, err
}

func UpdateNode(c *gin.Context, req *V1.UpdateNodeRequest) (id string, err error) {
	var (
		logger = logger2.NewLogger()
		node   *model.NodeEntity
	)
	if req.Name != "" {
		node.Name = req.Name
	}
	if req.Content != "" {
		node.Content = req.Content
	}
	if req.Depend != "" {
		node.Depend = req.Depend
	}
	if req.FatherId != 0 {
		node.FatherID = req.FatherId
	}
	if req.Done != nil {
		node.Done = *req.Done
	}
	if req.Note != "" {
		node.Note = req.Note
	}
	if req.Status != "" {
		node.Status = req.Status
	}
	if len(req.Tags) != 0 {
		node.Tags = req.Tags
	}
	if len(req.Children) != 0 {
		node.Children = req.Children
	}

	if err != nil {
		logger.Error("创建node 失败")
		return "", err
	}
	return
}

func NewExistEsQuery(param string, query elastic.Query) elastic.Query {
	if param == "" {
		return nil
	}
	return query
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
