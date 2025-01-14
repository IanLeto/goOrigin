package repository

import (
	"encoding/json"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"time"
)

// 从 ODAMetricEntity 转换为 TODAMetric
func ToODAMetricDAO(metric *entity.ODAMetricEntity) *dao.TODAMetric {
	// 将 CustomDimensions 转换为 JSON 字符串
	customDimensions, _ := json.Marshal(metric.CustomDimensions)

	return &dao.TODAMetric{
		Interval:         int64(metric.Interval.Milliseconds()), // 将 time.Duration 转换为毫秒
		Cluster:          metric.Cluster,
		TransType:        metric.TransType,
		TransTypeCode:    metric.TransTypeCode,
		TransChannel:     metric.TransChannel,
		RetCode:          metric.RetCode,
		SvcName:          metric.SvcName,
		SuccessCount:     metric.SuccessCount,
		SuccessRate:      metric.SuccessRate,
		FailedCount:      metric.FailedCount,
		FailedRate:       metric.FailedRate,
		ResponseCount:    metric.ResponseCount,
		ResponseRate:     metric.ResponseRate,
		CustomDimensions: string(customDimensions), // 转换为字符串存储
	}
}

// 从 TODAMetric 转换为 ODAMetricEntity
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
