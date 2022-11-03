package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	logger2 "goOrigin/pkg/logger"
)

func CreateNode(c *gin.Context, req *params.CreateNodeRequest) (id string, err error) {
	var (
		logger = logger2.NewLogger()
		node   *model.Node
	)
	node = &model.Node{
		Name:    req.Name,
		Content: req.Content,
		Father:  req.Father,
	}
	id, err = node.CreateNode(c)
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

func GetNodes(c *gin.Context, id, name string) (node []*model.Node, err error) {
	var (
		logger  = logger2.NewLogger()
		queries []elastic.Query
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		daoRes  *elastic.SearchResult
		//eq     = elastic.NewExistsQuery()
	)
	queries = append(queries, NewExistEsQuery(name, elastic.NewTermQuery("name", name)))
	queries = append(queries, NewExistEsQuery(id, elastic.NewTermQuery("_id", id)))
	bq.Must(queries...)
	daoRes, err = client.Search().Index("topo").Query(bq).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("查询topo失败%s", err.Error()))
		goto ERR
	}
	for _, hit := range daoRes.Hits.Hits {
		var ephemeralNode *model.Node
		err = json.Unmarshal(hit.Source, &ephemeralNode)
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

func DeleteNodes(c *gin.Context, ids []string) (interface{}, error) {
	var (
		logger  = logger2.NewLogger()
		filters []elastic.Query
		bq      = elastic.NewBoolQuery()
		client  *elastic.Client
		//node    *model.Node
		//daoRes  *elastic.DeleteResponse
		//err error
	)
	filters = append(filters, NewExistEsQuery("_id", elastic.NewTermsQuery("_id", ids)))
	bq.Filter(filters...)
	daoRes, err := client.DeleteByQuery().Index("topo").Query(bq).Do(c)
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
		node    *model.Node
		err     error
	)
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
