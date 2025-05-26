package entity

import (
	"time"
)

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

type SpanEntity struct {
	Stats []SpanDatEntity `json:"stats"`
}

type SpanDatEntity struct {
	TransType    string `json:"trans_type"`
	TransTypeCN  string `json:"trans_type_cn"`
	SuccessCount int64  `json:"success_count"`
	FailedCount  int64  `json:"failed_count"`
	UnknownCount int64  `json:"unknown_count"`
	Total        int64  `json:"total"`
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
	Project     string              `json:"project"`
	TransType   string              `json:"trans_type"`
	TransTypeCn string              `json:"trans_type_cn"`
	ReturnCode  []*ReturnCodeEntity `json:"return_code"`
	Interval    int                 `json:"interval"`
	Dimension1  string              `json:"dimension_1"`
	Dimension2  string              `json:"dimension_2"`
}
type ReturnCodeEntity struct {
	ReturnCode   string `json:"return_code"`
	ReturnCodeCn string `json:"return_code_cn"`
	ProjectID    string `json:"project_id"`
	TransType    string `json:"trans_type"`
	Status       string `json:"status"`
}

type TradeReturnCodeEntity struct {
	UrlPath       string
	SuccessCount  int
	FailedCount   int
	UnKnownCount  string
	Total         string
	TransTypeCn   string
	ResponseCount string
}

type TransTypeEntity struct {
	TransType   string   `json:"trans_type"`
	TransTypeCn string   `json:"trans_type_cn"`
	ReturnCodes []string `json:"return_codes"`
}

type TransTypeResponseEntity struct {
	Items []*TransTypeEntity `json:"items"`
}
