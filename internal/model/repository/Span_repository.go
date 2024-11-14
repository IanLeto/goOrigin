package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToOdaMetricEntity(tRecord *dao.ODAMetricMessage) *entity.ODAMetricEntity {
	return &entity.ODAMetricEntity{
		Interval: tRecord.Interval,
	}
}

func ToOdaMetricMessage(tRecord *entity.ODAMetricEntity) *dao.ODAMetricMessage {
	return &dao.ODAMetricMessage{
		Interval: tRecord.Interval,
	}
}
