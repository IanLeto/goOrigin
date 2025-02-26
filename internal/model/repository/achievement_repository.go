package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToAchievementRecordDAO(record *entity.AchievementRecordEntity) *dao.TAchievementRecord {
	return &dao.TAchievementRecord{
		Name:        record.Name,
		Description: record.Description,
		Points:      record.Points,
		Type:        record.Type,
		AchievedAt:  record.AchievedAt,
	}
}

func ToAchievementRecordEntity(tRecord *dao.TAchievementRecord) *entity.AchievementRecordEntity {
	return &entity.AchievementRecordEntity{
		Name:        tRecord.Name,
		Description: tRecord.Description,
		Points:      tRecord.Points,
		Type:        tRecord.Type,
		AchievedAt:  tRecord.AchievedAt,
		CreateTime:  tRecord.CreateTime,
		ModifyTime:  tRecord.ModifyTime,
	}
}
