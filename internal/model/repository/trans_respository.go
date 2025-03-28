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
				Code:        svcCode,
				NameCN:      svcCN,
				TransTypeID: transTypeID,
				TraceID:     entity.TraceID,
				Cluster:     entity.Cluster,
				PodName:     entity.PodName,
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
		entity.TransType[t.Code] = t.NameCN
		for _, svc := range t.ServiceCodes {
			entity.ServiceCode[svc.Code] = svc.NameCN

			// 只取第一条记录的 TraceID/Cluster/PodName（假定相同）
			if entity.TraceID == "" {
				entity.TraceID = svc.TraceID
				entity.Cluster = svc.Cluster
				entity.PodName = svc.PodName
			}
		}
	}
	return entity
}
