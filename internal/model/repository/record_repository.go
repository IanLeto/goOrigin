package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"time"
)

func ToRecordDAO(record *entity.RecordEntity) *dao.TRecord {
	return &dao.TRecord{

		Name:    record.Title,
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

// ToRecordEntity 方法：转换 TRecord 为 RecordEntity，并填充时间格式
func ToRecordEntity(tRecord *dao.TRecord) *entity.RecordEntity {
	// 创建 `RecordEntity` 实例
	record := &entity.RecordEntity{
		ID:         tRecord.ID,
		Title:      tRecord.Name,
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

	// 如果 CreateTime 存在，则填充时间格式
	if tRecord.CreateTime > 0 {
		// 转换为 `time.Time`
		createTime := time.Unix(tRecord.CreateTime, 0)

		// 生成不同格式的时间
		record.TimeInfo = entity.RecordTimeInfo{
			CreateTimeISO8601:   createTime.UTC().Format(time.RFC3339),                                            // UTC 时间（ISO 8601）
			CreateTimeISO8601CN: createTime.In(time.FixedZone("CST", 8*3600)).Format("2006-01-02T15:04:05+08:00"), // 东八区
			CreateTimeRFC822:    createTime.Format(time.RFC1123Z),                                                 // RFC 822 格式
			CreateTimeUnix:      tRecord.CreateTime,                                                               // Unix 时间戳（秒）
			CreateTimeUnixMs:    tRecord.CreateTime * 1000,                                                        // Unix 时间戳（毫秒）
			CreateTimeDBFormat:  createTime.Format("2006-01-02 15:04:05"),                                         // 数据库格式
			CreateTimeCompact:   createTime.Format("20060102150405"),                                              // 紧凑格式
		}
	}

	return record
}
