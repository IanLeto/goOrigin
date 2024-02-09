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
