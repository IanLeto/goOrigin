package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/API/V1"
	pbs "goOrigin/agent/protos"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
	"goOrigin/rpcClient"
)

type ScriptApis interface {
	CreateScript(c *gin.Context, req V1.CreateScriptRequest) (result interface{}, err error)
}

func CreateScript(c *gin.Context, req V1.CreateScriptRequest) (result interface{}, err error) {
	var (
		logger = logger2.NewLogger()
		script = &entity.BaseScript{
			Name:       req.Name,
			Comment:    req.Comment,
			Type:       req.Type,
			Content:    req.Content,
			File:       req.File,
			Uploader:   req.Uploader,
			CreateTime: req.CreateTime,
			UpdateTime: req.UpdateTime,
			System:     req.System,
			IsFile:     req.IsFile,
			Timeout:    req.Timeout,
			Tags:       req.Tags,
			UsedTime:   req.UsedTime,
		}
		client *elastic.Client
		res    *elastic.IndexResponse
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex("script") }()
	if err != nil {
		logger.Error(fmt.Sprintf("errors: %s", err))
		goto ERR
	}

	res, err = client.Index().Index("script").BodyJson(script).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("请求es失败 : %s", err))
		goto ERR
	}
	return res, err
ERR:
	{
		return nil, err
	}

}

func QueryScript(c *gin.Context, req V1.QueryScriptRequest) (res *V1.QueryScriptListResponse, err error) {
	var (
		bq      = elastic.NewBoolQuery()
		eq      = elastic.NewExistsQuery("Uploader") // 排除无效脚本
		logger  = logger2.NewLogger()
		queries []elastic.Query
		client  *elastic.Client
		daoRes  *elastic.SearchResult
		infos   []*V1.QueryScriptListResponseInfo
	)

	client, err = clients.NewESClient()
	if err != nil {
		logger.Error(fmt.Sprintf("errors : %s", err))
		goto ERR
	}
	if req.Key != "" {
		queries = append(queries, elastic.NewMatchQuery("Content", req.Key))
	}
	if req.Name != "" {
		queries = append(queries, elastic.NewTermQuery("GetName", req.Name))
	}
	if req.Type != "" {
		queries = append(queries, elastic.NewTermQuery("Type", req.Type))
	}
	bq.Must(queries...)

	daoRes, err = client.Search().Index("script").Query(eq).Query(bq).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("请求es失败 : %s", err))
		goto ERR
	}

	for _, hit := range daoRes.Hits.Hits {
		var ephemeralSc entity.BaseScript
		err = json.Unmarshal(hit.Source, &ephemeralSc)

		infos = append(infos, &V1.QueryScriptListResponseInfo{
			ID:         hit.Id,
			Name:       ephemeralSc.Name,
			Comment:    ephemeralSc.Comment,
			Type:       ephemeralSc.Type,
			Content:    ephemeralSc.Content,
			File:       ephemeralSc.File,
			Uploader:   ephemeralSc.Uploader,
			CreateTime: ephemeralSc.CreateTime,
			UpdateTime: ephemeralSc.UpdateTime,
			System:     ephemeralSc.System,
			IsFile:     ephemeralSc.IsFile,
			Timeout:    ephemeralSc.Timeout,
			Tags:       ephemeralSc.Tags,
		})

	}

	res = &V1.QueryScriptListResponse{Infos: infos}
	return res, err
ERR:
	{
		return nil, err
	}
}

func RunScript(c *gin.Context, id string) (*pbs.RunScriptResponse, error) {
	var (
		bq     = elastic.NewBoolQuery()
		logger = logger2.NewLogger()
		client *elastic.Client
		script *entity.BaseScript
	)
	agent, err := rpcClient.NewAgentClient()
	client, err = clients.NewESClient()
	bq.Filter(elastic.NewTermQuery("_id", id))
	result, err := client.Get().Index("script").Id(id).Do(c)

	if err != nil {
		logger.Error(fmt.Sprintf("查询es 失败 %s", err.Error()))
		return nil, err
	}
	err = json.Unmarshal(result.Source, &script)
	if err != nil {
		logger.Error(fmt.Sprintf("format es info 失败 %s", err.Error()))
		return nil, err
	}
	script.UsedTime += 1
	_, err = client.Index().Index("script").Id(script.ID).BodyJson(script).Do(c)
	return agent.RunScript(c, &pbs.RunScriptRequest{
		Name:    "",
		Content: script.Content,
	})

}
