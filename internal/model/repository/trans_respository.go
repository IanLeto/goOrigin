package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToServiceCodeDAOs(
	entity *entity.TransInfoEntity,
	projectID uint,
	transTypeMap map[string]uint, // transTypeCode => TransTypeID
) []*dao.EcampServiceCodeTb {
	var results []*dao.EcampServiceCodeTb

	for transCode, _ := range entity.TransType {
		transTypeID, ok := transTypeMap[transCode]
		if !ok {
			continue // 跳过未找到交易类型 ID 的项
		}
		for svcCode, svcCN := range entity.ServiceCode {
			results = append(results, &dao.EcampServiceCodeTb{
				ServiceCode:   svcCode,
				ServiceCodeCN: svcCN,
				TransTypeID:   transTypeID,
				Cluster:       entity.Cluster,
			})
		}
	}
	return results
}

func ToTransInfoEntityFromProject(
	project *dao.EcampProjectInfoTb,
	transTypes []dao.EcampTransTypeTb,
) *entity.TransInfoEntity {
	entity := &entity.TransInfoEntity{
		TraceID:     "",
		Cluster:     "",
		PodName:     "",
		Project:     project.Project,
		TransType:   make(map[string]string),
		ServiceCode: make(map[string]string),
	}

	for _, t := range transTypes {
		entity.TransType[t.TransType] = t.TransTypeCN
		for _, svc := range t.ServiceCodes {
			entity.ServiceCode[svc.ServiceCode] = svc.ServiceCodeCN
			// 只取第一条记录的 ProjectID/Cluster/PodName（假定相同）

		}
	}
	return entity
}
