package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
	"reflect"
)

func CreateScript(c *gin.Context, req params.CreateScriptRequest) (result interface{}, err error) {
	var (
		logger = logger2.NewLogger()
		script = &model.BaseScript{
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
	if err != nil {
		logger.Error(fmt.Sprintf("errors: %s", err))
		goto ERR
	}

	//switch req.Type {
	//case "py":
	//	script = &model.PythonScript{BaseScript: baseScript}
	//case "sh":
	//	script = &model.SellScript{BaseScript: baseScript}
	//}
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

func DelScript(c *gin.Context) {

}

func QueryScript(c *gin.Context, req params.QueryScriptRequest) (res *params.QueryScriptListResponse, err error) {
	var (
		bp      = elastic.NewBoolQuery()
		eq      = elastic.NewExistsQuery("Uploader") // 排除无效脚本
		logger  = logger2.NewLogger()
		queries []elastic.Query
		client  *elastic.Client
		result  *elastic.SearchResult
		script  *model.BaseScript
		infos   []*params.QueryScriptListResponseInfo
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
		queries = append(queries, elastic.NewTermQuery("Name", req.Name))
	}
	if req.Type != "" {
		queries = append(queries, elastic.NewTermQuery("Type", req.Type))
	}
	bp.Must(queries...)

	if req.Tags != "" {

	}

	result, err = client.Search().Index("script").Query(eq).Query(bp).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("请求es失败 : %s", err))
		goto ERR
	}
	for _, item := range result.Each(reflect.TypeOf(script)) {
		v, ok := item.(*model.BaseScript)
		if !ok {
			goto ERR
		}
		infos = append(infos, &params.QueryScriptListResponseInfo{
			ID:         v.ID,
			Name:       v.Name,
			Comment:    v.Comment,
			Type:       v.Type,
			Content:    v.Content,
			File:       v.File,
			Uploader:   v.Uploader,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			System:     v.System,
			IsFile:     v.IsFile,
			Timeout:    v.Timeout,
			Tags:       v.Tags,
		})
	}
	res = &params.QueryScriptListResponse{Infos: infos}
	return res, err
ERR:
	{
		return nil, err
	}
}
