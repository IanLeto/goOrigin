package dao

type TProject struct {
	Project  string `json:"project" gorm:"type:varchar(100);not null" json:"project,omitempty"`
	AZ       string `json:"az" gorm:"type:varchar(100);not null" json:"az,omitempty"`
	TracePod string `json:"trace_pod" gorm:"type:varchar(100);not null" json:"trace_pod,omitempty"`
	CreateAt int    `json:"create_at" gorm:"type:int;not null" json:"create_at,omitempty"`
	Update   int    `json:"update" gorm:"type:int;not null" json:"update,omitempty"`
}
