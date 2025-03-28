package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

// ToTransInfoDAO 方法：转换 TransInfoEntity 为 EcampTransInfoTb（基础结构 → GORM 结构）
func ToTransInfoDAO(trans *entity.TransInfoEntity) *dao.EcampTransInfoTb {
	return &dao.EcampTransInfoTb{
		TraceID:       trans.TraceID,
		Cluster:       trans.Cluster,
		PodName:       trans.PodName,
		SvcName:       trans.SvcName,
		TransType:     trans.TransType,
		ServiceCode:   trans.ServiceCode,
		TransTypeCN:   trans.TransTypeCN,
		ServiceCodeCN: trans.ServiceCodeCN,
	}
}

// ToTransInfoEntity 方法：转换 EcampTransInfoTb 为 TransInfoEntity（GORM 结构 → 基础结构）
func ToTransInfoEntity(tTrans *dao.EcampTransInfoTb) *entity.TransInfoEntity {
	return &entity.TransInfoEntity{
		TraceID:       tTrans.TraceID,
		Cluster:       tTrans.Cluster,
		PodName:       tTrans.PodName,
		SvcName:       tTrans.SvcName,
		TransType:     tTrans.TransType,
		ServiceCode:   tTrans.ServiceCode,
		TransTypeCN:   tTrans.TransTypeCN,
		ServiceCodeCN: tTrans.ServiceCodeCN,
	}
}
