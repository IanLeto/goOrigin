package repository

import (
	"fmt"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"strconv"
)

func ConvertToTransInfoEntity(transType *dao.EcampTransTypeTb, returnCodes []dao.EcampReturnCodeTb) *entity.TransInfoEntity {
	var returnCodeEntities []*entity.ReturnCodeEntity

	for _, rc := range returnCodes {
		returnCodeEntities = append(returnCodeEntities, &entity.ReturnCodeEntity{
			ReturnCode: rc.ReturnCode,
			Project:    rc.Project,
			TransType:  rc.TransType,
			Status:     rc.Status,
			Count:      0,
		})
	}

	return &entity.TransInfoEntity{
		Project:     transType.Project,
		TransType:   transType.TransType,
		TransTypeCn: []string{}, // 初始化为空数组，后续会在查询中填充
		ReturnCodes: returnCodeEntities,
		IsAlert:     transType.IsAlert,
		Threshold:   transType.Threshold,
	}
}

// 批量转换函数（处理多个TransType）
func ConvertToTransInfoEntities(transTypes []dao.EcampTransTypeTb, returnCodesMap map[string][]dao.EcampReturnCodeTb) []*entity.TransInfoEntity {
	var result []*entity.TransInfoEntity

	for _, transType := range transTypes {
		// 使用 transType + project 作为key查找对应的return codes
		key := fmt.Sprintf("%s_%s", transType.TransType, transType.Project)
		returnCodes := returnCodesMap[key]

		result = append(result, ConvertToTransInfoEntity(&transType, returnCodes))
	}

	return result
}

//// 辅助函数：查询TransType及其关联的ReturnCodes
//func GetTransTypeWithReturnCodes(db *gorm.DB, transType string, project string) (*entity.TransInfoEntity, error) {
//	var trans dao.EcampTransTypeTb
//	var returnCodes []dao.EcampReturnCodeTb
//
//	// 查询 TransType
//	if err := db.Table(dao.TableNameEcampTransTypeTb).
//		Where("trans_type = ? AND project = ?", transType, project).
//		First(&trans).Error; err != nil {
//		return nil, err
//	}
//
//	// 查询关联的 ReturnCodes
//	if err := db.Table(dao.TableNameEcampReturnCodeTb).
//		Where("trans_type = ? AND project = ?", transType, project).
//		Find(&returnCodes).Error; err != nil {
//		return nil, err
//	}
//
//	return ConvertToTransInfoEntity(&trans, returnCodes), nil
//}

// 批量查询函数
//func GetAllTransTypesWithReturnCodes(db *gorm.DB, project string) ([]*entity.TransInfoEntity, error) {
//	var transTypes []dao.EcampTransTypeTb
//	var returnCodes []dao.EcampReturnCodeTb
//
//	// 查询所有 TransTypes
//	query := db.Table(dao.TableNameEcampTransTypeTb)
//	if project != "" {
//		query = query.Where("project = ?", project)
//	}
//	if err := query.Find(&transTypes).Error; err != nil {
//		return nil, err
//	}
//
//	// 查询所有 ReturnCodes
//	rcQuery := db.Table(dao.TableNameEcampReturnCodeTb)
//	if project != "" {
//		rcQuery = rcQuery.Where("project = ?", project)
//	}
//	if err := rcQuery.Find(&returnCodes).Error; err != nil {
//		return nil, err
//	}
//
//	// 构建 returnCodes map
//	returnCodesMap := make(map[string][]dao.EcampReturnCodeTb)
//	for _, rc := range returnCodes {
//		key := fmt.Sprintf("%s_%s", rc.TransType, rc.Project)
//		returnCodesMap[key] = append(returnCodesMap[key], rc)
//	}
//
//	return ConvertToTransInfoEntities(transTypes, returnCodesMap), nil
//}

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
