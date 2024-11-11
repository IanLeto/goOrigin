package dao

import "time"

type TProject struct {
	*Meta     `json:"*_meta,omitempty"`
	Project   string        `json:"project" gorm:"type:varchar(100);not null" json:"project,omitempty"`
	ProjectCN string        `json:"project_cn" gorm:"type:varchar(100);not null" json:"project_cn,omitempty"`
	AZ        string        `json:"az" gorm:"type:varchar(100);not null" json:"az,omitempty"`
	TracePOD  string        `json:"trace_pod" gorm:"type:varchar(100);not null" json:"trace_pod,omitempty"`
	CreateAT  time.Duration `json:"create_at" gorm:"type:bigint;not null" json:"create_at,omitempty"`
	UpdateAT  time.Duration `json:"update_at" gorm:"type:bigint;not null" json:"update_at,omitempty"`
}

type TProjectBizCode struct {
	*Meta       `json:"*_meta,omitempty"`
	BizKey      string `json:"biz_key" gorm:"type:varchar(100);not null" json:"biz_key,omitempty"`
	BizValue    string `json:"biz_value" gorm:"type:varchar(100);not null" json:"biz_value,omitempty"`
	BizType     string `json:"biz_type" gorm:"type:varchar(100);not null" json:"biz_type,omitempty"`
	Project     string `json:"project" gorm:"type:varchar(100);not null" json:"project,omitempty"`
	Cluster     string `json:"cluster" gorm:"type:varchar(100);not null" json:"cluster,omitempty"`
	Service     string `json:"service" gorm:"type:varchar(100);not null" json:"service,omitempty"`
	SpanAlias   string `json:"span_alias" gorm:"type:varchar(100);not null" json:"span_alias,omitempty"`
	CreatedUser string `json:"created_user" gorm:"type:varchar(100);not null" json:"created_user,omitempty"`
	UpdateUser  string `json:"update_user" gorm:"type:varchar(100);not null" json:"update_user,omitempty"`
}
