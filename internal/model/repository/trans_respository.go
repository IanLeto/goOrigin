package repository

import (
	"goOrigin/API/V1"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"strconv"
)

func ConvertToTransInfoEntity(transType *dao.EcampTransTypeTb) *entity.TransInfoEntity {
	var returnCodeEntities []*entity.ReturnCodeEntity
	for _, rc := range transType.ReturnCodes {
		returnCodeEntities = append(returnCodeEntities, &entity.ReturnCodeEntity{
			ReturnCode: rc.ReturnCode,
			ProjectID:  rc.Project,
			TransType:  rc.TransType,
			Status:     rc.Status,
		})
	}

	return &entity.TransInfoEntity{
		Project:     transType.Project,
		TransType:   transType.TransType,
		ReturnCodes: returnCodeEntities,
		Dimension1:  transType.Dimension1,
		Dimension2:  transType.Dimension2,
	}
}

func ToEcampTransTypeTb(entity *entity.TransInfoEntity) *dao.EcampTransTypeTb {
	if entity == nil {
		return nil
	}

	model := &dao.EcampTransTypeTb{
		TransType:   entity.TransType,
		Project:     entity.Project,
		TransTypeCN: entity.TransTypeCn, // 可补充
		IsAlert:     false,
		Dimension1:  entity.Dimension1,
		Dimension2:  entity.Dimension2,
	}

	for _, rc := range entity.ReturnCodes {
		returnCode := dao.EcampReturnCodeTb{
			TransType:  rc.TransType,
			ReturnCode: rc.ReturnCode,
			Project:    rc.ProjectID,
			Status:     rc.Status,
		}
		model.ReturnCodes = append(model.ReturnCodes, returnCode)
	}

	return model
}

// ToTransTypeCountEntity 按照文档和表结构，计算成功率情况
func ToTransTypeCountEntity(doc *dao.EcampAggUrlDoc, transInfos []entity.TransInfoEntity) *entity.TradeReturnCodeEntity {
	if doc == nil {
		return nil
	}

	// 1. 查找匹配的交易类型
	var matchedTrans *entity.TransInfoEntity
	for i := range transInfos {
		if transInfos[i].TransType == doc.ReqURL {
			matchedTrans = &transInfos[i]
			break
		}
	}

	// 2. 构建 returnCode -> status 映射（success / failed / unknown）
	returnCodeStatusMap := make(map[string]string)
	if matchedTrans != nil {
		for _, rc := range matchedTrans.ReturnCodes {
			if rc != nil {
				returnCodeStatusMap[rc.ReturnCode] = rc.Status
			}
		}
	}

	// 3. 分类计算
	var (
		successCount int
		failedCount  int
		unknownCount int
	)

	for _, rcAgg := range doc.ReturnCodeAgg {
		status, exists := returnCodeStatusMap[rcAgg.ReturnCode]
		count := int(rcAgg.Count)

		switch {
		case exists && status == "success":
			successCount += count
		case exists && status == "failed":
			failedCount += count
		default:
			unknownCount += count
		}
	}

	total := successCount + failedCount + unknownCount

	// 4. 构建返回结构
	return &entity.TradeReturnCodeEntity{
		UrlPath:       doc.ReqURL,
		SuccessCount:  successCount,
		FailedCount:   failedCount,
		UnKnownCount:  strconv.Itoa(unknownCount),
		Total:         strconv.Itoa(total),
		ResponseCount: strconv.Itoa(int(doc.TotalCount)),
	}
}

// ConvertToTransTypeResponse 将ES查询结果和配置信息转换为响应格式
func ConvertToTransTypeResponse(
	kafkaLogs []*entity.KafkaLogEntity,
	transInfoMap map[string]*entity.TransInfoEntity, // key: trans_type
	page, pageSize int,
) *V1.TransTypeResponse2 {
	// 用于去重和聚合相同URL的返回码
	urlReturnCodeMap := make(map[string]map[string]string)
	urlCnNameMap := make(map[string]string)

	// 遍历ES数据
	for _, log := range kafkaLogs {
		// 获取URL路径和返回码
		urlPath := log.ReqURL
		returnCode := log.ReturnCode

		if urlPath == "" || returnCode == "" {
			continue
		}

		// 从Trans扩展信息中获取trans_type
		transType := log.Trans.TransType // 假设SpanTransTypeInfoEntity有TransType字段

		// 查找对应的配置信息
		transInfo, exists := transInfoMap[transType]
		if !exists {
			continue
		}

		// 初始化URL的返回码map
		if _, ok := urlReturnCodeMap[urlPath]; !ok {
			urlReturnCodeMap[urlPath] = make(map[string]string)
			urlCnNameMap[urlPath] = transInfo.TransTypeCn
		}

		// 查找返回码对应的中文描述
		for _, rc := range transInfo.ReturnCodes {
			if rc.ReturnCode == returnCode {
				break
			}
		}
	}

	// 转换为最终响应格式
	items := make([]*V1.TransTypeItem2, 0, len(urlReturnCodeMap))
	for urlPath, returnCodeMap := range urlReturnCodeMap {
		item := &V1.TransTypeItem2{
			UrlPath:    urlPath,
			UrlPathCN:  urlCnNameMap[urlPath],
			ReturnCode: returnCodeMap,
		}
		items = append(items, item)
	}

	// 计算分页
	total := len(items)
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	// 返回分页后的数据
	return &V1.TransTypeResponse2{
		Items:    items[start:end],
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

func ToUrlPathAggEntity(record *dao.AggUrlPathDoc) *entity.UrlPathAggEntity {
	panic(1)
}
