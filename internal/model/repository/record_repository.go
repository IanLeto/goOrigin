package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToRecordDAO(record *entity.RecordEntity) *dao.TRecord {
	return &dao.TRecord{
		Name:    record.Name,
		Weight:  record.Weight,
		Vol1:    record.Vol1,
		Vol2:    record.Vol2,
		Vol3:    record.Vol3,
		Vol4:    record.Vol4,
		Content: record.Content,
		Region:  record.Region,
		IsFuck:  record.IsFuck,
		Cost:    record.Cost,
		Dev:     record.Dev,
		Coding:  record.Coding,
		Social:  record.Social,
	}
}

func ToRecordEntity(tRecord *dao.TRecord) *entity.RecordEntity {
	return &entity.RecordEntity{
		Name:       tRecord.Name,
		Weight:     tRecord.Weight,
		Content:    tRecord.Content,
		Cost:       tRecord.Cost,
		Vol1:       tRecord.Vol1,
		Vol2:       tRecord.Vol2,
		Vol3:       tRecord.Vol3,
		Vol4:       tRecord.Vol4,
		Dev:        tRecord.Dev,
		Coding:     tRecord.Coding,
		Social:     tRecord.Social,
		IsFuck:     tRecord.IsFuck,
		CreateTime: tRecord.CreateTime,
		ModifyTime: tRecord.ModifyTime,
	}
}
