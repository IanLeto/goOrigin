package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/repository"

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
	logger.Debug(fmt.Sprintf("query => %s", func() string {
		//s, _ := json.MarshalIndent(query, "", "  ")
		s, _ := json.Marshal(query)
		return string(s)
	}()))
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
		logger.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, successRateEntity)
	if err != nil {
		logger.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
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
		logger.Error(fmt.Sprintf("create record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	err = utils.JsonToStruct(value, successRateEntity)
	if err != nil {
		logger.Error(fmt.Sprintf("conv record failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}
	return successRateEntity, nil
ERR:
	return nil, err
}

func CreateType(ctx context.Context, region string, reqs []V1.CreateTransInfo) error {
	// 获取数据库连接
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig)

	// 开启事务
	tx := db.Client.Begin()
	if tx.Error != nil {
		logrus.Errorf("failed to begin transaction: %v", tx.Error)
		return tx.Error
	}

	// 捕获异常，确保出错时回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			logrus.Errorf("panic occurred during CreateType: %v", r)
		}
	}()

	for _, req := range reqs {
		// 1. 查询项目信息，确保存在
		var projectInfo dao.EcampProjectInfoTb
		if err := tx.Table(dao.TableNameEcampProjectInfoTb).
			Where("project = ?", req.Project).
			First(&projectInfo).Error; err != nil {
			logrus.Errorf("project [%s] not found: %v", req.Project, err)
			tx.Rollback()
			return err
		}

		// 2. 检查是否已存在相同 trans_type + project
		var existing dao.EcampTransTypeTb
		if err := tx.Table(dao.TableNameEcampTransTypeTb).
			Where("trans_type = ? AND project = ?", req.TransType, req.Project).
			First(&existing).Error; err == nil {
			logrus.Infof("trans_type [%s] for project [%s] already exists, skipping", req.TransType, req.Project)
			continue
		}

		// 3. 插入交易类型记录
		transType := &dao.EcampTransTypeTb{
			TransType:   req.TransType,
			TransTypeCN: "", // 可扩展字段
			Project:     req.Project,
			IsAlert:     false,
			Dimension1:  req.Dimension1,
			Dimension2:  req.Dimension2,
		}

		if err := tx.Table(dao.TableNameEcampTransTypeTb).Create(transType).Error; err != nil {
			logrus.Errorf("failed to insert trans_type [%s]: %v", req.TransType, err)
			tx.Rollback()
			return err
		}

		// 4. 插入 return_code 记录
		var returnCodes []dao.EcampReturnCodeTb
		for code, cn := range req.ServiceCode {
			returnCodes = append(returnCodes, dao.EcampReturnCodeTb{
				TransType:    req.TransType,
				ReturnCode:   code,
				ReturnCodeCN: cn,
				Project:      req.Project,
				Status:       "active",
			})
		}

		if len(returnCodes) > 0 {
			if err := tx.Table(dao.TableNameEcampReturnCodeTb).
				Create(&returnCodes).Error; err != nil {
				logrus.Errorf("failed to insert return codes for trans_type [%s]: %v", req.TransType, err)
				tx.Rollback()
				return err
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		logrus.Errorf("commit transaction failed: %v", err)
		return err
	}

	return nil
}

func QueryTransTypeList(ctx context.Context, region string, project, transType string) ([]*entity.TransInfoEntity, error) {
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig)

	var transTypeTbList []dao.EcampTransTypeTb

	// 查询并预加载 return codes
	query := db.Client.Debug().Debug().
		Preload("ReturnCodes").
		Table(dao.TableNameEcampTransTypeTb).
		Where("project = ?", project)

	if transType != "" {
		query = query.Where("trans_type = ?", transType)
	}

	if err := query.Find(&transTypeTbList).Error; err != nil {
		logrus.Errorf("query trans types failed: %v", err)
		return nil, err
	}

	// 使用转换函数
	var result []*entity.TransInfoEntity
	for _, t := range transTypeTbList {
		result = append(result, repository.ConvertToTransInfoEntity(&t))
	}

	return result, nil
}

func DeleteTransInfo(ctx context.Context, region, project, transType string) error {
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig)

	tx := db.Client.Begin()

	// 删除返回码
	if err := tx.
		Where("project = ? AND trans_type = ?", project, transType).
		Delete(&dao.EcampReturnCodeTb{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除交易类型
	if err := tx.
		Where("project = ? AND trans_type = ?", project, transType).
		Delete(&dao.EcampTransTypeTb{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateTransInfo(ctx context.Context, region string, item *entity.TransInfoEntity) error {
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig)

	tx := db.Client.Begin()

	// 1. 更新主表
	if err := tx.Model(&dao.EcampTransTypeTb{}).
		Where("project = ? AND trans_type = ?", item.Project, item.TransType).
		Updates(map[string]interface{}{
			"trans_type_cn": item.TransType, // 示例字段，如有中文名可替换
			"dimension1":    "",             // 可补充
			"dimension2":    "",
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2. 删除旧返回码
	if err := tx.
		Where("project = ? AND trans_type = ?", item.Project, item.TransType).
		Delete(&dao.EcampReturnCodeTb{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3. 插入新返回码
	var newCodes []dao.EcampReturnCodeTb
	for _, rc := range item.ReturnCode {
		newCodes = append(newCodes, dao.EcampReturnCodeTb{
			TransType:    rc.TransType,
			ReturnCode:   rc.ReturnCode,
			ReturnCodeCN: rc.ReturnCodeCn,
			Project:      rc.ProjectID,
			Status:       rc.Status,
		})
	}
	if len(newCodes) > 0 {
		if err := tx.Create(&newCodes).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
