package dao

// EcampTransInfoTb 数据库表结构（GORM ORM 结构）
type EcampTransInfoTb struct {
	*Meta         `json:"*_meta,omitempty"`
	TraceID       string `json:"trace_id" gorm:"column:trace_id;type:varchar(255);not null"`
	Cluster       string `json:"cluster" gorm:"column:cluster;type:varchar(100);not null"`
	PodName       string `json:"pod_name" gorm:"column:pod_name;type:varchar(100);not null"`
	SvcName       string `json:"svc_name" gorm:"column:svc_name;type:varchar(100);not null"`
	TransType     string `json:"trans_type" gorm:"column:trans_type;type:varchar(50);not null"`
	ServiceCode   string `json:"service_code" gorm:"column:service_code;type:varchar(50);not null"`
	TransTypeCN   string `json:"trans_type_cn" gorm:"column:trans_type_cn;type:varchar(100);"`
	ServiceCodeCN string `json:"service_code_cn" gorm:"column:service_code_cn;type:varchar(100);"`
}

// TableName 指定数据库表名
func (EcampTransInfoTb) TableName() string {
	return "ecamp_trans_info_tb"
}
