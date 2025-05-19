package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

// 将 GORM 模型转换为业务结构体
func ToTransInfoEntity(model *dao.EcampTransTypeTb) *entity.TransInfoEntity {
	if model == nil {
		return nil
	}

	// 默认只取第一个 ReturnCode
	var returnCode *entity.ReturnCodeEntity
	if len(model.ReturnCodes) > 0 {
		rc := model.ReturnCodes[0] // 可根据需要调整逻辑
		returnCode = &entity.ReturnCodeEntity{
			ReturnCode:   rc.ReturnCode,
			ReturnCodeCn: rc.ReturnCodeCN,
			ProjectID:    rc.Project,
			TransType:    rc.TransType,
			Status:       rc.Status,
		}
	}

	return &entity.TransInfoEntity{
		Cluster:    "", // cluster 信息未提供，可根据上下文设置
		Project:    model.Project,
		TransType:  model.TransType,
		ReturnCode: returnCode,
		Interval:   0, // interval 信息未存储于数据库，可后期扩展
	}
}

// 将业务结构体转换为 GORM 模型
func ToEcampTransTypeTb(entity *entity.TransInfoEntity) *dao.EcampTransTypeTb {
	if entity == nil {
		return nil
	}

	model := &dao.EcampTransTypeTb{
		TransType:   entity.TransType,
		Project:     entity.Project,
		TransTypeCN: "", // 如有需要可以补充
		IsAlert:     false,
		Dimension1:  "",
		Dimension2:  "",
	}

	if entity.ReturnCode != nil {
		model.ReturnCodes = []dao.EcampReturnCodeTb{
			{
				TransType:    entity.ReturnCode.TransType,
				ReturnCode:   entity.ReturnCode.ReturnCode,
				ReturnCodeCN: entity.ReturnCode.ReturnCodeCn,
				Project:      entity.ReturnCode.ProjectID,
				Status:       entity.ReturnCode.Status,
			},
		}
	}

	return model
}
