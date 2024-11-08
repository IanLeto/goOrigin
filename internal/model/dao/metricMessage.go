package dao

import "time"

type ODAMetric struct {
	Interval time.Duration `json:"interval"`
}

type Dimension struct {
	Cluster string `json:"cluster"`
	Src     string `json:"src"`
	Psrc    string `json:"psrc"`
}

type Indicator struct {
	SuccessCount  int `json:"success_count"`
	SuccessRate   int `json:"success_rate"`
	FailedCount   int `json:"failed_count"`
	FailedRate    int `json:"failed_rate"`
	ResponseCount int `json:"response_count"`
	ResponseRate  int `json:"response_rate"`
}
