package entity

import "time"

type SuccessRateEntity struct {
}

type ODAMetricEntity struct {
	Interval time.Duration `json:"interval"`
}

type Dimension struct {
	Cluster       string `json:"cluster"`
	Src           string `json:"src"`
	Psrc          string `json:"psrc"`
	TransType     string `json:"trans_type"`
	TransTypeCode string `json:"trans_type_code"`
	TransTypeDesc string `json:"trans_type_desc"`
	TransChannel  string `json:"trans_channel"`
	RetCode       string `json:"ret_code"`
}

type Indicator struct {
	SuccessCount  int `json:"success_count"`
	SuccessRate   int `json:"success_rate"`
	FailedCount   int `json:"failed_count"`
	FailedRate    int `json:"failed_rate"`
	ResponseCount int `json:"response_count"`
	ResponseRate  int `json:"response_rate"`
}
