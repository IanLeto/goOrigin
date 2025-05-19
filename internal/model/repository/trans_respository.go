package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ConvertToTransInfoEntity(transType *dao.EcampTransTypeTb) *entity.TransInfoEntity {
	var returnCodeEntities []*entity.ReturnCodeEntity
	for _, rc := range transType.ReturnCodes {
		returnCodeEntities = append(returnCodeEntities, &entity.ReturnCodeEntity{
			ReturnCode:   rc.ReturnCode,
			ReturnCodeCn: rc.ReturnCodeCN,
			ProjectID:    rc.Project,
			TransType:    rc.TransType,
			Status:       rc.Status,
		})
	}

	return &entity.TransInfoEntity{
		Project:    transType.Project,
		TransType:  transType.TransType,
		ReturnCode: returnCodeEntities,
		Dimension1: transType.Dimension1,
		Dimension2: transType.Dimension2,
		Interval:   0,
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

	for _, rc := range entity.ReturnCode {
		returnCode := dao.EcampReturnCodeTb{
			TransType:    rc.TransType,
			ReturnCode:   rc.ReturnCode,
			ReturnCodeCN: rc.ReturnCodeCn,
			Project:      rc.ProjectID,
			Status:       rc.Status,
		}
		model.ReturnCodes = append(model.ReturnCodes, returnCode)
	}

	return model
}
