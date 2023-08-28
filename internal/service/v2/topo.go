package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/API/V1"
	"goOrigin/internal/model"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
)

func CreateNode(c *gin.Context, req *V1.CreateNodeRequest) (id string, err error) {
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

	id, err = node.CreateNode(c)
	if err != nil {
		logger.Error("创建node 失败")
		return id, err
	}
	return
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
	if req.FatherId != "" {
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

	id, err = node.UpdateNode(c)
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

func GetNodes(c *gin.Context, id, name, father, content string, done bool) (node []*model.NodeEntity, err error) {
	panic(1)
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
	//res, err := conn.Query("node", queries)
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
	return node, nil
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
func GetTopoList(c *gin.Context) (res []*V1.GetTopoResponse, err error) {
	var (
		logger = logger2.NewLogger()
		client *elastic.Client
		daoRes *elastic.SearchResult
		node   *model.NodeEntity
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex(model.EsNode) }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化es 失败 %s", err))
		return nil, err
	}
	daoRes, err = client.Search().Index(model.EsNode).Query(elastic.NewTermsQuery("father", "")).Do(c)
	if len(daoRes.Hits.Hits) == 0 {
		err = errors.New("不存在该数据")
		goto ERR
	}
	for _, hit := range daoRes.Hits.Hits {
		err = json.Unmarshal(hit.Source, &node)
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

	node = model.GetTopo(c, node)
	return

ERR:
	{
		return nil, err
	}
}

func GetTopo(c *gin.Context, name string) (res *V1.GetTopoResponse, err error) {
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
