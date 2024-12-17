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
		Dimension: &entity.Dimension{
			Cluster: tRecord.Dimension.Cluster,
			//Src:           tRecord.Dimension.Src,
			//Psrc:          tRecord.Dimension.Psrc,
			TransType:     tRecord.Dimension.TransType,
			TransTypeCode: tRecord.Dimension.TransTypeCode,
			TransTypeDesc: tRecord.Dimension.TransTypeDesc,
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
			Cluster: tRecord.Dimension.Cluster,
			//Src:           tRecord.Dimension.Src,
			//Psrc:          tRecord.Dimension.Psrc,
			TransType:     tRecord.Dimension.TransType,
			TransTypeCode: tRecord.Dimension.TransTypeCode,
			TransTypeDesc: tRecord.Dimension.TransTypeDesc,
			TransChannel:  tRecord.Dimension.TransChannel,
			RetCode:       tRecord.Dimension.RetCode,
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
