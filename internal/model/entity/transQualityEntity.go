package entity

import "time"

type ODAMetricEntity struct {
	Interval time.Duration `json:"interval"`
	*PredefinedDimensions
	*Indicator
	CustomDimensions []string `json:"custom_dimensions"`
}

type PredefinedDimensions struct {
	TraceID       string `json:"trace_id"`
	Cluster       string `json:"cluster"`
	TransTypeCode string `json:"trans_type_code"` // 锚定字段
	TransChannel  string `json:"trans_channel"`
	RetCode       string `json:"ret_code"`
	SvcName       string `json:"svc_name"`
}

type Indicator struct {
	SuccessCount  int `json:"success_count"`
	SuccessRate   int `json:"success_rate"`
	FailedCount   int `json:"failed_count"`
	FailedRate    int `json:"failed_rate"`
	ResponseCount int `json:"response_count"`
	ResponseRate  int `json:"response_rate"`
}

type SuccessRateEntity struct {
}

func ConvertLogToMetric(log *KafkaLogEntity) ODAMetricEntity {
	// 组装目标结构
	metric := ODAMetricEntity{
		PredefinedDimensions: &PredefinedDimensions{
			Cluster:       log.InstanceZone,
			TransTypeCode: log.LogType,
			TransChannel:  log.RemoteApp,
			RetCode:       log.ResultCode,
			SvcName:       log.Service,
		},
	}
	return metric
}

// TransInfoEntity 网关交易
type TransInfoEntity struct {
	Cluster    string            `json:"cluster"`
	Project    string            `json:"project"`
	TransType  string            `json:"trans_type"`
	ReturnCode *ReturnCodeEntity `json:"return_code"`
	Interval   int               `json:"interval"`
}
type ReturnCodeEntity struct {
	ReturnCode   string `json:"return_code"`
	ReturnCodeCn string `json:"return_code_cn"`
	ProjectID    string `json:"project_id"`
	TransType    string `json:"trans_type"`
	Status       string `json:"status"`
}
