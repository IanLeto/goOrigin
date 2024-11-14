package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goOrigin/API/V1"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/model/entity"

	"goOrigin/pkg/utils"
	"strings"
)

// OdaSuccessAndFailedRateMetric 返回成功和失败的比率
func OdaSuccessAndFailedRateMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SuccessRateEntity, error) {
	var (
		projectDocEntity = &entity.ProjectAggDocEntity{}
		err              error
		//conn             = elastic.GlobalEsConns[region]
	)
	var (
		query = map[string]interface{}{}
		//alias = ""
		aggs = map[string]interface{}{}
	)
	var (
		totalTranslations      = map[string]interface{}{}
		successfulTranslations = map[string]interface{}{}
	)
	var (
		ret_code string = "aaaa"
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
	// 聚合算成功率
	if ret_code != "" {
		totalTranslations = map[string]interface{}{
			"filter": map[string]interface{}{
				"value_count": map[string]interface{}{
					"field": "ret_code.keyword",
				},
			},
		}
		successfulTranslations = map[string]interface{}{
			"filter": map[string]interface{}{
				"terms": map[string]interface{}{
					"ret_code.keyword": []string{"0"},
				},
			},
		}
	}
	aggs = map[string]interface{}{
		"total_translations":      totalTranslations,
		"successful_translations": successfulTranslations,
	}
	query = map[string]interface{}{
		"aggs": aggs,
		"size": 0,
		"bool": map[string]interface{}{
			"query": query,
			"sort":  sort,
		},
	}
	// todo
	value := []byte{}
	//value, err := conn.Search(alias, query)
	log.Debug(fmt.Sprintf("query => %s", func() string {
		//s, _ := json.MarshalIndent(query, "", "  ")
		s, _ := json.Marshal(query)
		return string(s)
	}()))
	if err != nil {
		log.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, projectDocEntity)
	if err != nil {
		log.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	return nil, err
ERR:
	return nil, err

}

// OdaSuccessCountAndFailedCountMetric 返回成功和失败的数量
func OdaSuccessCountAndFailedCountMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SuccessRateEntity, error) {
	var (
		projectDocEntity = &entity.ProjectAggDocEntity{}
		err              error
		conn             = elastic.GlobalEsConns[region]
	)
	var (
		query = map[string]interface{}{}
		alias = ""
		aggs  = map[string]interface{}{}
	)
	var ()
	var (
		ret_code string
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
	// 聚合算成功率
	if ret_code != "" {
		filterCallback(&filter, "ret_code", []string{ret_code})
	}
	if region != "" {
		filterCallback(&filter, "region", strings.Split(region, ","))
	}

	query = map[string]interface{}{
		"aggs": aggs,
		"size": 0,
		"bool": map[string]interface{}{
			"must": []map[string]interface{}{
				{
					"exists": map[string]interface{}{
						"field": "trans.ret_code.keyword",
					},
				},
			},
			"filter": filter,
		},
		"sort": sort,
	}
	value, err := conn.Search(alias, query)
	if err != nil {
		log.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, projectDocEntity)
	if err != nil {
		log.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	return nil, err
ERR:
	return nil, err

}

// OdaRespRateMetric  返回响应率
func OdaRespRateMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SuccessRateEntity, error) {
	var (
		successRateEntity = &entity.SuccessRateEntity{}
		err               error
		conn              = elastic.GlobalEsConns[region]
	)
	var (
		query = map[string]interface{}{}
		alias = ""
		aggs  = map[string]interface{}{}
	)
	var (
		total_docs = map[string]interface{}{}
		valid_docs = map[string]interface{}{}
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
	// 聚合计算总文档数和有效文档数
	total_docs = map[string]interface{}{
		"value_count": map[string]interface{}{
			"field": "_id",
		},
	}
	valid_docs = map[string]interface{}{
		"filter": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "ret_code.keyword",
						},
					},
					{
						"bool": map[string]interface{}{
							"must_not": []map[string]interface{}{
								{
									"term": map[string]interface{}{
										"ret_code.keyword": "",
									},
								},
								{
									"term": map[string]interface{}{
										"ret_code.keyword": nil,
									},
								},
							},
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"valid_count": map[string]interface{}{
				"value_count": map[string]interface{}{
					"field": "_id",
				},
			},
		},
	}
	aggs = map[string]interface{}{
		"total_docs": total_docs,
		"valid_docs": valid_docs,
	}
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "ret_code.keyword",
						},
					},
				},
			},
		},
		"aggs": aggs,
		"size": 0,
		"sort": sort,
	}
	value, err := conn.Search(alias, query)
	if err != nil {
		log.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, successRateEntity)
	if err != nil {
		log.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	return successRateEntity, nil
ERR:
	return nil, err
}

// OdaSuccessCountMetric 成功/失败/错误数
func OdaSuccessCountMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SuccessRateEntity, error) {
	var (
		successRateEntity = &entity.SuccessRateEntity{}
		err               error
		conn              = elastic.GlobalEsConns[region]
	)
	var (
		query = map[string]interface{}{}
		alias = ""
		aggs  = map[string]interface{}{}
	)
	var (
		filters_count = map[string]interface{}{}
	)
	var (
		req_count  = map[string]interface{}{}
		resp_count = map[string]interface{}{}
		err_count  = map[string]interface{}{}
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

	filters_count = map[string]interface{}{
		"filters": map[string]interface{}{
			"filters": map[string]interface{}{
				"req_count":  req_count,
				"resp_count": resp_count,
				"err_count":  err_count,
			},
		},
	}
	aggs = map[string]interface{}{
		"filters_count": filters_count,
	}
	query = map[string]interface{}{
		"query": map[string]interface{}{},
		"aggs":  aggs,
		"size":  0,
		"sort":  sort,
	}
	value, err := conn.Search(alias, query)
	if err != nil {
		log.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, successRateEntity)
	if err != nil {
		log.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	return successRateEntity, nil
ERR:
	return nil, err
}
