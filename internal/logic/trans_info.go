package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/repository"
	"gorm.io/gorm"

	"goOrigin/internal/model/entity"
	"goOrigin/pkg/utils"
	"strings"
)

func SearchTransInfo(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.TradeReturnCodeEntity, error) {
	panic(1)
}

func SearchTransTypeSuccessStatsMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SpanEntity, error) {
	var (
		conn             = elastic.GlobalEsConns[region]
		ret              = &entity.SpanEntity{}
		err              error
		esIndex          = "span"
		projectAggResult = map[string]interface{}{}
	)

	// -----------------------------
	// 1. 获取数据库连接并加载返回码配置
	// -----------------------------
	db := mysql.GlobalMySQLConns[region]
	var returnCodes []dao.EcampReturnCodeTb

	codeQuery := db.Client.Debug().Table("ecamp_return_code_tb")

	// 可选查询条件（按项目或交易类型过滤）
	//if info.Project != "" {
	//	codeQuery = codeQuery.Where("project = ?", info.Project)
	//}
	//if len(info.TransTypes) > 0 {
	//	codeQuery = codeQuery.Where("trans_type IN ?", info.TransTypes)
	//}

	err = codeQuery.Find(&returnCodes).Error
	if err != nil {
		logger.Error(fmt.Sprintf("查询返回码配置失败: %v", err))
		return nil, err
	}

	// 构建映射 map[transType][return_code] = status
	returnCodeMap := make(map[string]map[string]string)
	for _, rc := range returnCodes {
		if _, ok := returnCodeMap[rc.TransType]; !ok {
			returnCodeMap[rc.TransType] = make(map[string]string)
		}
		returnCodeMap[rc.TransType][rc.ReturnCode] = strings.ToLower(rc.Status)
	}

	// -----------------------------
	// 2. 构建 Elasticsearch 查询
	// -----------------------------
	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"by_transType": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "transType.keyword",
					"size":  1000,
				},
				"aggs": map[string]interface{}{
					"by_returnCode": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "return_code.keyword",
							"size":  1000,
						},
					},
				},
			},
		},
	}

	// -----------------------------
	// 3. 执行 Elasticsearch 查询
	// -----------------------------
	value, err := conn.Search(esIndex, query)
	if err != nil {
		logger.Error(fmt.Sprintf("ES 查询失败: %s\nQuery: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		return nil, err
	}

	err = json.Unmarshal(value, &projectAggResult)
	if err != nil {
		logger.Error(fmt.Sprintf("解析 ES 返回失败: %s", err))
		return nil, err
	}

	// -----------------------------
	// 4. 加载交易类型中文名称映射
	// -----------------------------
	var transTypes []dao.EcampTransTypeTb
	typeQuery := db.Client.Debug().Table("ecamp_trans_type_tb")
	//if info.Project != "" {
	//	typeQuery = typeQuery.Where("project = ?", info.Project)
	//}
	err = typeQuery.Find(&transTypes).Error
	if err != nil {
		logger.Error(fmt.Sprintf("加载交易类型失败: %s", err))
		return nil, err
	}

	transTypeCNMap := make(map[string]string)
	for _, t := range transTypes {
		transTypeCNMap[t.TransType] = t.TransTypeCN
	}

	// -----------------------------
	// 5. 组装统计结果
	// -----------------------------
	aggregations := projectAggResult["aggregations"].(map[string]interface{})
	transBuckets := aggregations["by_transType"].(map[string]interface{})["buckets"].([]interface{})

	for _, bucket := range transBuckets {
		b := bucket.(map[string]interface{})
		transType := b["key"].(string)
		returnBuckets := b["by_returnCode"].(map[string]interface{})["buckets"].([]interface{})

		var (
			successCount int64
			failedCount  int64
			unknownCount int64
			total        int64
		)

		for _, rb := range returnBuckets {
			r := rb.(map[string]interface{})
			retCode := r["key"].(string)
			count := int64(r["doc_count"].(float64))

			status := returnCodeMap[transType][retCode]
			switch status {
			case "success":
				successCount += count
			case "failed":
				failedCount += count
			default:
				unknownCount += count
			}
			total += count
		}

		ret.Stats = append(ret.Stats, entity.SpanDatEntity{
			TransType:    transType,
			TransTypeCN:  transTypeCNMap[transType],
			SuccessCount: successCount,
			FailedCount:  failedCount,
			UnknownCount: unknownCount,
			Total:        total,
		})
	}

	return ret, nil
}

// OdaSuccessAndFailedRateMetric 返回成功和失败的比率
func OdaSuccessAndFailedRateMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SpanEntity, error) {
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
func OdaSuccessCountAndFailedCountMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SpanEntity, error) {
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
func OdaRespRateMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SpanEntity, error) {
	var (
		successRateEntity = &entity.SpanEntity{}
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
func OdaSuccessCountMetric(ctx *gin.Context, region string, info *V1.SuccessRateReqInfo) (*entity.SpanEntity, error) {
	var (
		successRateEntity = &entity.SpanEntity{}
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

		// 2. 查询是否存在该 trans_type + project
		var existing dao.EcampTransTypeTb
		err := tx.Table(dao.TableNameEcampTransTypeTb).
			Where("trans_type = ? AND project = ?", req.TransType, req.Project).
			First(&existing).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("query existing trans_type [%s] failed: %v", req.TransType, err)
			tx.Rollback()
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在：插入交易类型
			newTrans := &dao.EcampTransTypeTb{
				TransType:   req.TransType,
				TransTypeCN: "", // 可扩展字段
				Project:     req.Project,
				IsAlert:     false,
				Dimension1:  req.Dimension1,
				Dimension2:  req.Dimension2,
			}

			if err := tx.Table(dao.TableNameEcampTransTypeTb).Create(newTrans).Error; err != nil {
				logrus.Errorf("failed to insert trans_type [%s]: %v", req.TransType, err)
				tx.Rollback()
				return err
			}
		} else {
			// 存在：更新字段
			updateFields := map[string]interface{}{
				"dimension1": req.Dimension1,
				"dimension2": req.Dimension2,
				"is_alert":   false,
			}

			if err := tx.Table(dao.TableNameEcampTransTypeTb).
				Where("trans_type = ? AND project = ?", req.TransType, req.Project).
				Updates(updateFields).Error; err != nil {
				logrus.Errorf("failed to update trans_type [%s]: %v", req.TransType, err)
				tx.Rollback()
				return err
			}
		}

		// 3. 删除旧的 return_code
		if err := tx.Table(dao.TableNameEcampReturnCodeTb).
			Where("trans_type = ? AND project = ?", req.TransType, req.Project).
			Delete(&dao.EcampReturnCodeTb{}).Error; err != nil {
			logrus.Errorf("failed to delete return codes for trans_type [%s]: %v", req.TransType, err)
			tx.Rollback()
			return err
		}

		// 4. 插入新的 return_codes
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

func QueryTransTypeList(ctx context.Context, region, project, transType string, page, pageSize int) ([]*entity.TransInfoEntity, int64, error) {
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig)

	var (
		transTypeTbList []dao.EcampTransTypeTb
		total           int64
	)

	query := db.Client.
		Debug().
		Model(&dao.EcampTransTypeTb{}).
		Preload("ReturnCodes").
		Where("project = ?", project)

	if transType != "" {
		query = query.Where("trans_type = ?", transType)
	}

	// 1. 查询总数
	if err := query.Count(&total).Error; err != nil {
		logrus.Errorf("count query failed: %v", err)
		return nil, 0, err
	}

	// 2. 分页查询数据
	offset := (page - 1) * pageSize
	if err := query.
		Limit(pageSize).
		Offset(offset).
		Find(&transTypeTbList).Error; err != nil {
		logrus.Errorf("query data failed: %v", err)
		return nil, 0, err
	}

	// 3. 转换为 entity
	var result []*entity.TransInfoEntity
	for _, t := range transTypeTbList {
		result = append(result, repository.ConvertToTransInfoEntity(&t))
	}

	return result, total, nil
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
	for _, rc := range item.ReturnCodes {
		newCodes = append(newCodes, dao.EcampReturnCodeTb{
			TransType:  rc.TransType,
			ReturnCode: rc.ReturnCode,
			Project:    rc.ProjectID,
			Status:     rc.Status,
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

func QueryTransTypeWithReturnCodesInfo(ctx *gin.Context, region string, info *V1.TransTypeQueryInfo) (*entity.TransTypeResponseEntity, error) {
	var (
		conn         = elastic.GlobalEsConns[region]
		result       = &entity.TransTypeResponseEntity{}
		transTypeMap = make(map[string]*entity.TransTypeEntity)
		db           = mysql.GlobalMySQLConns[region]
	)

	// -----------------------------
	// Step 1: 查询 ES 聚合数据
	// -----------------------------
	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"by_transType": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "transType.keyword",
					"size":  1000,
				},
				"aggs": map[string]interface{}{
					"by_return_code": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "return_code.keyword",
							"size":  1000,
						},
					},
				},
			},
		},
	}

	raw, err := conn.Search("span", query)
	if err != nil {
		return nil, fmt.Errorf("ES 查询失败: %v", err)
	}

	var aggResult map[string]interface{}
	err = json.Unmarshal(raw, &aggResult)
	if err != nil {
		return nil, fmt.Errorf("ES 结果解析失败: %v", err)
	}

	buckets := aggResult["aggregations"].(map[string]interface{})["by_transType"].(map[string]interface{})["buckets"].([]interface{})

	for _, b := range buckets {
		bucket := b.(map[string]interface{})
		transType := bucket["key"].(string)
		returnCodeBuckets := bucket["by_return_code"].(map[string]interface{})["buckets"].([]interface{})

		entityItem := &entity.TransTypeEntity{
			TransType:   transType,
			ReturnCodes: []string{},
		}

		for _, rc := range returnCodeBuckets {
			code := rc.(map[string]interface{})["key"].(string)
			entityItem.ReturnCodes = append(entityItem.ReturnCodes, code)
		}

		transTypeMap[transType] = entityItem
	}

	// -----------------------------
	// Step 2: 查询数据库中的中文名
	// -----------------------------
	var transTypeDBs []dao.EcampTransTypeTb
	sql := db.Client.Debug().Model(&dao.EcampTransTypeTb{})
	if info.Project != "" {
		sql = sql.Where("project = ?", info.Project)
	}
	if len(info.TransTypes) > 0 {
		sql = sql.Where("trans_type IN ?", info.TransTypes)
	}
	if err := sql.Find(&transTypeDBs).Error; err != nil {
		return nil, fmt.Errorf("数据库查询失败: %v", err)
	}

	// 设置中文名
	for _, t := range transTypeDBs {
		if val, ok := transTypeMap[t.TransType]; ok {
			val.TransTypeCn = t.TransTypeCN
		}
	}

	// -----------------------------
	// Step 3: 转换结果为列表
	// -----------------------------
	for _, v := range transTypeMap {
		result.Items = append(result.Items, v)
	}

	return result, nil
}
func SearchUrlPathWithReturnCode2(ctx *gin.Context, region string, info *V1.SearchUrlPathWithReturnCodesInfo) ([]*entity.UrlPathAggEntity, error) {
	var (
		aggUrlPathDoc = &dao.AggUrlPathDoc{}
		result        []*entity.UrlPathAggEntity
		err           error
		conn          = elastic.GlobalEsConns[region]
	)
	var (
		query = map[string]interface{}{}
		alias = "your-index-alias" // 需要替换为实际的索引别名
		aggs  = map[string]interface{}{}
	)
	var (
		byTransType  = map[string]interface{}{}
		byReturnCode = map[string]interface{}{}
	)
	var (
		mustConditions []map[string]interface{}
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

	// 构建过滤条件
	if info.Project != "" {
		mustConditions = append(mustConditions, map[string]interface{}{
			"term": map[string]interface{}{
				"project.keyword": info.Project,
			},
		})
	}

	if info.Az != "" {
		mustConditions = append(mustConditions, map[string]interface{}{
			"term": map[string]interface{}{
				"az.keyword": info.Az,
			},
		})
	}

	if len(info.TransTypes) > 0 {
		filterCallback(&mustConditions, "trans_type", info.TransTypes)
	}

	if info.StartTime > 0 && info.EndTime > 0 {
		mustConditions = append(mustConditions, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": info.StartTime,
					"lte": info.EndTime,
				},
			},
		})
	}

	if info.Keyword != "" {
		mustConditions = append(mustConditions, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  info.Keyword,
				"fields": []string{"trans_type", "url_path"},
			},
		})
	}

	// 构建聚合
	byReturnCode = map[string]interface{}{
		"terms": map[string]interface{}{
			"field": "return_code.keyword",
			"size":  100,
		},
	}

	byTransType = map[string]interface{}{
		"terms": map[string]interface{}{
			"field": "trans_type.keyword",
			"size":  1000,
		},
		"aggs": map[string]interface{}{
			"by_return_code": byReturnCode,
		},
	}

	aggs = map[string]interface{}{
		"by_transType": byTransType,
	}

	// 构建查询
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustConditions,
			},
		},
		"aggs": aggs,
		"size": 0,
	}

	// 添加排序
	if info.OrderBy != "" {
		query["sort"] = []map[string]interface{}{
			{
				info.OrderBy: map[string]interface{}{
					"order": "desc",
				},
			},
		}
	}

	// 执行查询
	value, err := conn.Search(alias, query)
	if err != nil {
		logger.Error(fmt.Sprintf("search url path failed %s: %s", err, func() string {
			s, _ := json.Marshal(query)
			return string(s)
		}()))
		goto ERR
	}

	// 转换结果
	err = utils.JsonToStruct(value, aggUrlPathDoc)
	if err != nil {
		logger.Error(fmt.Sprintf("conv url path result failed %s: %s", err, func() string {
			s, _ := json.Marshal(value)
			return string(s)
		}()))
		goto ERR
	}

	// 处理聚合结果
	for _, bucket := range aggUrlPathDoc.Aggregations.ByTransType.Buckets {
		urlPathEntity := &entity.UrlPathAggEntity{
			UrlPath:         bucket.Key,
			UrlPathCN:       bucket.Key, // 这里可能需要转换逻辑
			ReturnCode:      make(map[string]string),
			ReturnCodeCount: make(map[string]int),
		}

		for _, codeBucket := range bucket.ByReturnCode.Buckets {
			urlPathEntity.ReturnCode[codeBucket.Key] = codeBucket.Key
			urlPathEntity.ReturnCodeCount[codeBucket.Key] = codeBucket.DocCount
		}

		result = append(result, urlPathEntity)
	}

	return result, nil

ERR:
	return nil, err
}

func SearchUrlPathWithReturnCode(ctx *gin.Context, region string, info *V1.SearchUrlPathWithReturnCodesInfo) ([]*entity.UrlPathAggEntity, error) {
	// 直接返回固定的mock数据，不做任何过滤或判断
	mockData := []*entity.UrlPathAggEntity{
		{
			UrlPath:   "/api/v1/user/login",
			UrlPathCN: "用户登录接口",
			Project:   "user-service",
			ReturnCode: map[string]string{
				"AA200": "成功",
				"AA201": "创建成功",
				"DD400": "请求参数错误",
				"DD401": "未授权",
				"ZZ403": "禁止访问",
				"SS404": "资源不存在",
				"VV500": "服务器内部错误",
				"VV502": "网关错误",
				"VV503": "服务不可用",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 15678,
				"AA201": 234,
				"DD400": 1256,
				"DD401": 890,
				"ZZ403": 456,
				"SS404": 234,
				"VV500": 89,
				"VV502": 23,
				"VV503": 12,
			},
		},
		{
			UrlPath:   "/api/v2/order/create",
			UrlPathCN: "订单创建接口",
			Project:   "order-service",
			ReturnCode: map[string]string{
				"AA200": "成功",
				"DD400": "请求参数错误",
				"DD409": "资源冲突",
				"DD429": "请求过于频繁",
				"VV500": "服务器内部错误",
				"VV504": "网关超时",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 98765,
				"DD400": 3456,
				"DD409": 567,
				"DD429": 1234,
				"VV500": 234,
				"VV504": 45,
			},
		},
		{
			UrlPath:   "/api/v3/payment/process",
			UrlPathCN: "支付处理接口",
			Project:   "payment-service",
			ReturnCode: map[string]string{
				"AA200": "支付成功",
				"AA202": "处理中",
				"DD400": "参数错误",
				"DD402": "支付失败",
				"VV500": "系统异常",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 45678,
				"AA202": 12345,
				"DD400": 789,
				"DD402": 456,
				"VV500": 123,
			},
		},
		{
			UrlPath:   "/admin/v1/dashboard",
			UrlPathCN: "管理后台仪表盘",
			Project:   "admin-service",
			ReturnCode: map[string]string{
				"AA200": "成功",
				"DD401": "未授权",
				"ZZ403": "权限不足",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 5678,
				"DD401": 234,
				"ZZ403": 89,
			},
		},
		{
			UrlPath:   "/health/check",
			UrlPathCN: "健康检查接口",
			Project:   "monitor-service",
			ReturnCode: map[string]string{
				"AA200": "健康",
				"VV503": "服务不可用",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 999999,
				"VV503": 12,
			},
		},
		{
			UrlPath:   "/api/v4/product/search",
			UrlPathCN: "商品搜索接口",
			Project:   "product-service",
			ReturnCode: map[string]string{
				"AA200": "成功",
				"AA204": "无内容",
				"AA206": "部分内容",
				"DD400": "搜索条件错误",
				"DD408": "请求超时",
				"DD413": "请求体过大",
				"DD422": "无法处理的实体",
				"VV500": "搜索引擎错误",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 87654,
				"AA204": 5432,
				"AA206": 1234,
				"DD400": 876,
				"DD408": 234,
				"DD413": 56,
				"DD422": 123,
				"VV500": 45,
			},
		},
		{
			UrlPath:   "/websocket/connect",
			UrlPathCN: "WebSocket连接",
			Project:   "websocket-service",
			ReturnCode: map[string]string{
				"AA101": "切换协议",
				"DD400": "错误的请求",
				"DD426": "需要升级",
				"VV500": "内部错误",
			},
			ReturnCodeCount: map[string]int{
				"AA101": 23456,
				"DD400": 567,
				"DD426": 123,
				"VV500": 34,
			},
		},
		{
			UrlPath:   "/api/v5/file/upload",
			UrlPathCN: "文件上传接口",
			Project:   "file-service",
			ReturnCode: map[string]string{
				"AA200": "上传成功",
				"AA201": "创建成功",
				"DD400": "文件格式错误",
				"DD413": "文件过大",
				"DD415": "不支持的媒体类型",
				"VV500": "存储服务错误",
				"VV507": "存储空间不足",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 34567,
				"AA201": 12345,
				"DD400": 2345,
				"DD413": 678,
				"DD415": 345,
				"VV500": 123,
				"VV507": 23,
			},
		},
		{
			UrlPath:   "/batch/job/execute",
			UrlPathCN: "批处理任务执行",
			Project:   "batch-service",
			ReturnCode: map[string]string{
				"AA200": "执行成功",
				"AA202": "已接受，处理中",
				"DD423": "资源锁定",
				"DD424": "失败的依赖",
				"VV500": "执行失败",
				"VV504": "执行超时",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 8765,
				"AA202": 4321,
				"DD423": 234,
				"DD424": 123,
				"VV500": 56,
				"VV504": 12,
			},
		},
		{
			UrlPath:   "/metrics/prometheus",
			UrlPathCN: "监控指标接口",
			Project:   "monitor-service",
			ReturnCode: map[string]string{
				"AA200": "成功",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 9999999,
			},
		},
		{
			UrlPath:   "/api/legacy/v0/deprecated",
			UrlPathCN: "已废弃的旧接口",
			Project:   "legacy-service",
			ReturnCode: map[string]string{
				"AA301": "永久重定向",
				"AA308": "永久重定向(保持方法)",
				"DD410": "已删除",
			},
			ReturnCodeCount: map[string]int{
				"AA301": 456,
				"AA308": 234,
				"DD410": 123,
			},
		},
		{
			UrlPath:   "/internal/debug/trace",
			UrlPathCN: "内部调试追踪",
			Project:   "debug-service",
			ReturnCode: map[string]string{
				"AA200": "成功",
				"DD401": "未授权",
				"ZZ403": "禁止访问",
				"DD405": "方法不允许",
				"VV501": "未实现",
				"VV511": "需要网络认证",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 123,
				"DD401": 456,
				"ZZ403": 789,
				"DD405": 234,
				"VV501": 56,
				"VV511": 12,
			},
		},
		{
			UrlPath:   "/graphql/query",
			UrlPathCN: "GraphQL查询接口",
			Project:   "graphql-service",
			ReturnCode: map[string]string{
				"AA200": "查询成功",
				"DD400": "查询语法错误",
				"DD401": "未授权",
				"DD429": "查询过于频繁",
				"VV500": "解析器错误",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 56789,
				"DD400": 1234,
				"DD401": 567,
				"DD429": 234,
				"VV500": 89,
			},
		},
		{
			UrlPath:   "/api/v6/stream/video",
			UrlPathCN: "视频流媒体接口",
			Project:   "stream-service",
			ReturnCode: map[string]string{
				"AA200": "流传输成功",
				"AA206": "部分内容",
				"DD400": "无效的范围请求",
				"DD416": "请求范围不满足",
				"VV500": "流媒体服务错误",
				"VV504": "流超时",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 123456,
				"AA206": 98765,
				"DD400": 2345,
				"DD416": 678,
				"VV500": 345,
				"VV504": 123,
			},
		},
		{
			UrlPath:   "/oauth2/token",
			UrlPathCN: "OAuth2令牌接口",
			Project:   "auth-service",
			ReturnCode: map[string]string{
				"AA200": "令牌颁发成功",
				"DD400": "无效的授权请求",
				"DD401": "客户端认证失败",
				"DD403": "禁止的授权范围",
				"VV500": "认证服务错误",
			},
			ReturnCodeCount: map[string]int{
				"AA200": 678901,
				"DD400": 12345,
				"DD401": 6789,
				"DD403": 1234,
				"VV500": 567,
			},
		},
	}

	return mockData, nil
}
