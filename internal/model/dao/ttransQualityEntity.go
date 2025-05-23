package dao

import "time"

const TableNameEcampProjectInfoTb = "ecamp_project_info_tb"

// EcampProjectInfoTb mapped from table <ecamp_project_info_tb>
type EcampProjectInfoTb struct {
	*Meta
	Project   string    `gorm:"column:project;type:varchar(253);not null;uniqueIndex:az_pod_proj_uniqe,priority:3" json:"project"`
	ProjectCn string    `gorm:"column:project_cn;type:varchar(253);not null" json:"project_cn"`
	Az        string    `gorm:"column:az;type:varchar(253);not null;uniqueIndex:az_pod_proj_uniqe,priority:1" json:"az"`
	TracePod  string    `gorm:"column:trace_pod;type:varchar(253);not null;uniqueIndex:az_pod_proj_uniqe,priority:2" json:"trace_pod"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);not null" json:"updated_at"`
}

// TableName EcampProjectInfoTb's table name
func (*EcampProjectInfoTb) TableName() string {
	return TableNameEcampProjectInfoTb
}

type EcampTransTypeTb struct {
	*Meta
	TransType   string              `gorm:"size:50;not null;uniqueIndex"`              // 加 uniqueIndex
	TransTypeCN string              `gorm:"size:100"`                                  // 开户7
	Project     string              `gorm:"size:255;not null"`                         // 改为字符串，与 TransInfoEntity 一致
	ReturnCodes []EcampReturnCodeTb `gorm:"foreignKey:TransType;references:TransType"` // 通过字符串关联
	IsAlert     bool
	Dimension1  string `gorm:"size:100"` // 交易类型
	Dimension2  string `gorm:"size:100"` // 交易渠道
}

type EcampReturnCodeTb struct {
	*Meta
	TransType    string `gorm:"size:50;index"`    // 外键字段
	ReturnCode   string `gorm:"size:50;not null"` // SC001
	ReturnCodeCN string `gorm:"size:100"`         // 开户成功
	Project      string `gorm:"size:255"`         // 与 ReturnCodeEntity.ProjectID 匹配
	Status       string `gorm:"size:50"`
}

const TableNameEcampTransTypeTb = "ecamp_trans_info_tb"

func (*EcampTransTypeTb) TableName() string {
	return TableNameEcampTransTypeTb
}

const TableNameEcampReturnCodeTb = "ecamp_return_code_tb"

func (*EcampReturnCodeTb) TableName() string {
	return TableNameEcampReturnCodeTb
}
