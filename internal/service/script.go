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

func QueryScript(c *gin.Context, req params.QueryScriptRequest) (res *params.QueryScriptListResponse, err error) {
	var (
		logger = logger2.NewLogger()
		client *elastic.Client
		result *elastic.SearchResult
		script *model.BaseScript
	)
	client, err = clients.NewESClient()
	if err != nil {
		logger.Error(fmt.Sprintf("errors : %s", err))
		goto ERR
	}
	// todo 增加查询相关
	result, err = client.Search().Index("script").Query(elastic.NewMatchAllQuery()).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("请求es失败 : %s", err))
		goto ERR
	}
	for _, item := range result.Each(reflect.TypeOf(script)) {
		v, ok := item.(model.BaseScript)
		if !ok {
			goto ERR
		}
		res.Infos = append(res.Infos, &params.QueryScriptListResponseInfo{
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
	return res, err
ERR:
	{
		return nil, err
	}
}
