package entity

import "time"

type ODAMetricEntity struct {
	Interval time.Duration `json:"interval"`
	*PredefinedDimensions
	*Indicator
	CustomDimensions []string `json:"custom_dimensions"`
}

type PredefinedDimensions struct {
	Cluster       string `json:"cluster"`
	TransType     string `json:"trans_type"`
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
