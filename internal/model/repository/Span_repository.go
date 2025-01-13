package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToOdaMetricEntity(tRecord *dao.ODAMetricMessage) *entity.ODAMetricEntity {
	if tRecord == nil {
		return nil
	}

	return &entity.ODAMetricEntity{
		Interval: tRecord.Interval,
		PredefinedDimensions: &entity.PredefinedDimensions{
			Cluster: tRecord.Dimension.Cluster,

			TransType:     tRecord.Dimension.TransType,
			TransTypeCode: tRecord.Dimension.TransTypeCode,
			TransChannel:  tRecord.Dimension.TransChannel,
			RetCode:       tRecord.Dimension.RetCode,
		},
		Indicator: &entity.Indicator{
			SuccessCount:  tRecord.Indicator.SuccessCount,
			SuccessRate:   tRecord.Indicator.SuccessRate,
			FailedCount:   tRecord.Indicator.FailedCount,
			FailedRate:    tRecord.Indicator.FailedRate,
			ResponseCount: tRecord.Indicator.ResponseCount,
			ResponseRate:  tRecord.Indicator.ResponseRate,
		},
	}
}

func ToOdaMetricMessage(tRecord *entity.ODAMetricEntity) *dao.ODAMetricMessage {
	if tRecord == nil {
		return nil
	}

	return &dao.ODAMetricMessage{
		Interval: tRecord.Interval,
		Dimension: &dao.Dimension{
			Cluster: tRecord.PredefinedDimensions.Cluster,
			//Src:           tRecord.PredefinedDimensions.Src,
			//Psrc:          tRecord.PredefinedDimensions.Psrc,
			TransType:     tRecord.PredefinedDimensions.TransType,
			TransTypeCode: tRecord.PredefinedDimensions.TransTypeCode,
			TransChannel:  tRecord.PredefinedDimensions.TransChannel,
			RetCode:       tRecord.PredefinedDimensions.RetCode,
		},
		Indicator: &dao.Indicator{
			SuccessCount:  tRecord.Indicator.SuccessCount,
			SuccessRate:   tRecord.Indicator.SuccessRate,
			FailedCount:   tRecord.Indicator.FailedCount,
			FailedRate:    tRecord.Indicator.FailedRate,
			ResponseCount: tRecord.Indicator.ResponseCount,
			ResponseRate:  tRecord.Indicator.ResponseRate,
		},
	}
}
