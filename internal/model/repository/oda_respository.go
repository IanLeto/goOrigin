package repository

import (
	"encoding/json"
	"goOrigin/API/V1"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"time"
)

// SvcTransAlertRecordInfo 转 ODAMetricEntity
func ToODAMetricEntityFromInfo(info *V1.SvcTransAlertRecordInfo) *entity.ODAMetricEntity {
	return &entity.ODAMetricEntity{
		Interval: time.Duration(info.Interval), // 直接赋值
		PredefinedDimensions: &entity.PredefinedDimensions{
			Cluster:       info.Cluster,
			TransType:     info.TransType,
			TransTypeCode: info.TransTypeCode,
			TransChannel:  info.TransChannel,
			RetCode:       info.RetCode,
			SvcName:       info.SvcName,
		},
		Indicator: &entity.Indicator{
			SuccessCount:  info.SuccessCount,
			SuccessRate:   info.SuccessRate,
			FailedCount:   info.FailedCount,
			FailedRate:    info.FailedRate,
			ResponseCount: info.ResponseCount,
			ResponseRate:  info.ResponseRate,
		},
		CustomDimensions: info.CustomDimensions, // 直接赋值为切片
	}
}

// ODAMetricEntity 转 TODAMetric
func ToODAMetricDAO(metric *entity.ODAMetricEntity) *dao.TODAMetric {
	// 将 CustomDimensions 转换为 JSON 字符串
	customDimensions, _ := json.Marshal(metric.CustomDimensions)

	return &dao.TODAMetric{
		Interval:         int64(metric.Interval.Milliseconds()), // 将 time.Duration 转换为毫秒
		Cluster:          metric.PredefinedDimensions.Cluster,
		TransType:        metric.PredefinedDimensions.TransType,
		TransTypeCode:    metric.PredefinedDimensions.TransTypeCode,
		TransChannel:     metric.PredefinedDimensions.TransChannel,
		RetCode:          metric.PredefinedDimensions.RetCode,
		SvcName:          metric.PredefinedDimensions.SvcName,
		SuccessCount:     metric.Indicator.SuccessCount,
		SuccessRate:      metric.Indicator.SuccessRate,
		FailedCount:      metric.Indicator.FailedCount,
		FailedRate:       metric.Indicator.FailedRate,
		ResponseCount:    metric.Indicator.ResponseCount,
		ResponseRate:     metric.Indicator.ResponseRate,
		CustomDimensions: string(customDimensions), // 转换为字符串存储
	}
}

// TODAMetric 转 ODAMetricEntity
func ToODAMetricEntity(tMetric *dao.TODAMetric) *entity.ODAMetricEntity {
	// 将 CustomDimensions 从 JSON 字符串转换为切片
	var customDimensions []string
	_ = json.Unmarshal([]byte(tMetric.CustomDimensions), &customDimensions)

	return &entity.ODAMetricEntity{
		Interval: time.Duration(tMetric.Interval) * time.Millisecond, // 将毫秒转换回 time.Duration
		PredefinedDimensions: &entity.PredefinedDimensions{
			Cluster:       tMetric.Cluster,
			TransType:     tMetric.TransType,
			TransTypeCode: tMetric.TransTypeCode,
			TransChannel:  tMetric.TransChannel,
			RetCode:       tMetric.RetCode,
			SvcName:       tMetric.SvcName,
		},
		Indicator: &entity.Indicator{
			SuccessCount:  tMetric.SuccessCount,
			SuccessRate:   tMetric.SuccessRate,
			FailedCount:   tMetric.FailedCount,
			FailedRate:    tMetric.FailedRate,
			ResponseCount: tMetric.ResponseCount,
			ResponseRate:  tMetric.ResponseRate,
		},
		CustomDimensions: customDimensions, // 转换为切片返回
	}
}
