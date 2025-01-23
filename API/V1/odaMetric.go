package V1

import "time"

type SuccessRateReqInfo struct {
}

type SvcTransAlertRecordInfo struct {
	Interval         time.Duration `json:"interval,omitempty"`
	Cluster          string        `json:"cluster,omitempty"`
	TransType        string        `json:"trans_type,omitempty"`
	TransTypeCode    string        `json:"trans_type_code,omitempty"`
	TransChannel     string        `json:"trans_channel,omitempty"`
	RetCode          string        `json:"ret_code,omitempty"`
	SvcName          string        `json:"svc_name,omitempty"`
	SuccessCount     int           `json:"success_count,omitempty"`
	SuccessRate      int           `json:"success_rate,omitempty"`
	FailedCount      int           `json:"failed_count,omitempty"`
	FailedRate       int           `json:"failed_rate,omitempty"`
	ResponseCount    int           `json:"response_count,omitempty"`
	ResponseRate     int           `json:"response_rate,omitempty"`
	CustomDimensions []string      `json:"custom_dimensions,omitempty"`
}
