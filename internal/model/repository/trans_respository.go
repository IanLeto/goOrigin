package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

// ToTransInfoDAO 方法：转换 TransInfoEntity 为 TTransInfo（基础结构 → GORM 结构）
func ToTransInfoDAO(trans *entity.TransInfoEntity) *dao.TTransInfo {
	return &dao.TTransInfo{
		TraceID:   trans.TraceID,
		Cluster:   trans.Cluster,
		Channel:   trans.Channel,
		PodName:   trans.PodName,
		SvcName:   trans.SvcName,
		TransType: trans.TransType,
	}
}

// ToTransInfoEntity 方法：转换 TTransInfo 为 TransInfoEntity（GORM 结构 → 基础结构）
func ToTransInfoEntity(tTrans *dao.TTransInfo) *entity.TransInfoEntity {
	return &entity.TransInfoEntity{
		TraceID:   tTrans.TraceID,
		Cluster:   tTrans.Cluster,
		Channel:   tTrans.Channel,
		PodName:   tTrans.PodName,
		SvcName:   tTrans.SvcName,
		TransType: tTrans.TransType,
	}
}
