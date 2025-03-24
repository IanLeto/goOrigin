package dao

// TTransInfo 数据库表结构（GORM ORM 结构）
type TTransInfo struct {
	*Meta     `json:"*_meta,omitempty"`
	TraceID   string `json:"trace_id" gorm:"column:trace_id;type:varchar(255);not null"`
	Cluster   string `json:"cluster" gorm:"column:cluster;type:varchar(100);not null"`
	Channel   string `json:"channel" gorm:"column:channel;type:varchar(100);not null"`
	PodName   string `json:"pod_name" gorm:"column:pod_name;type:varchar(100);not null"`
	SvcName   string `json:"svc_name" gorm:"column:svc_name;type:varchar(100);not null"`
	TransType string `json:"trans_type" gorm:"column:trans_type;type:varchar(50);not null"`
}

// TableName 指定数据库表名
func (TTransInfo) TableName() string {
	return "trans_info"
}
