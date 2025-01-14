package dao

// TODAMetric 表映射结构体
type TODAMetric struct {
	*Meta `json:"*_meta,omitempty"`

	Interval         int64  `gorm:"type:bigint;not null" json:"interval"` // 存储为毫秒
	Cluster          string `gorm:"type:varchar(255)" json:"cluster"`
	TransType        string `gorm:"type:varchar(255)" json:"trans_type"`
	TransTypeCode    string `gorm:"type:varchar(255)" json:"trans_type_code"` // 锚定字段
	TransChannel     string `gorm:"type:varchar(255)" json:"trans_channel"`
	RetCode          string `gorm:"type:varchar(255)" json:"ret_code"`
	SvcName          string `gorm:"type:varchar(255)" json:"svc_name"`
	SuccessCount     int    `gorm:"type:int" json:"success_count"`
	SuccessRate      int    `gorm:"type:int" json:"success_rate"`
	FailedCount      int    `gorm:"type:int" json:"failed_count"`
	FailedRate       int    `gorm:"type:int" json:"failed_rate"`
	ResponseCount    int    `gorm:"type:int" json:"response_count"`
	ResponseRate     int    `gorm:"type:int" json:"response_rate"`
	CustomDimensions string `gorm:"type:text" json:"custom_dimensions"` // 存储为 JSON 格式字符串
}
