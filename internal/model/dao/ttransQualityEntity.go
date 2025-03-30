package dao

import "time"

const TableNameEcampProjectInfoTb = "ecamp_project_info_tb"

// EcampProjectInfoTb mapped from table <ecamp_project_info_tb>
type EcampProjectInfoTb struct {
	ID        int32     `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
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
	ID           uint   `gorm:"primaryKey"`
	Code         string `gorm:"size:50"`  // T001
	CodeCN       string `gorm:"size:100"` // 开户7
	ProjectID    uint
	ServiceCodes []EcampServiceCodeTb `gorm:"foreignKey:TransTypeID"`
}

type EcampServiceCodeTb struct {
	TransTypeID   uint   `gorm:"primaryKey"`
	ServiceCode   string `gorm:"size:50"`  // SC001
	ServiceCodeCN string `gorm:"size:100"` // 开户成功
	TraceID       string `gorm:"size:255"`
	Cluster       string `gorm:"size:100"`
	PodName       string `gorm:"size:100"`
}

const TableNameEcampTransTypeTb = "ecamp_trans_type_tb"

func (*EcampTransTypeTb) TableName() string {
	return TableNameEcampTransTypeTb
}

const TableNameEcampServiceCodeTb = "ecamp_service_code_tb"

func (*EcampServiceCodeTb) TableName() string {
	return TableNameEcampServiceCodeTb
}
