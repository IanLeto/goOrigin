package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/utils"
	"strings"
)

func AggProject(ctx *gin.Context, region string, project string) (*entity.ProjectAggDocEntity, error) {
	var (
		//tMessage         = &dao.ODAMetric{}
		projectDocEntity = &entity.ProjectAggDocEntity{}
		err              error
		conn             = elastic.GlobalEsConns[region]
	)
	var (
		query = map[string]interface{}{}
		alias = ""
		agg   = map[string]interface{}{}
	)
	var (
	//start int
	//end   int
	)
	var (
		sort = []map[string]interface{}{
			{
				"time": map[string]interface{}{
					"order": "desc",
				},
			},
		}
		filter         []map[string]interface{}
		filterCallback = func(filter *[]map[string]interface{}, key string, value []string) {
			if len(value) > 0 {
				*filter = append(*filter, map[string]interface{}{
					"terms": map[string]interface{}{
						fmt.Sprintf("%s.keyword", key): value,
					},
				})
			}
		}
	)
	if region != "" {
		filterCallback(&filter, "region", strings.Split(region, ","))
	}
	query = map[string]interface{}{
		"bool": map[string]interface{}{
			"query":        query,
			"aggregations": agg,
			"sort":         sort,
		},
	}
	value, err := conn.Search(alias, query)
	if err != nil {
		logger.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, projectDocEntity)
	if err != nil {
		logger.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	return projectDocEntity, err
ERR:
	return nil, err

}
