package model

import "goOrigin/internal/dao/mysql"

type TNode struct {
	*mysql.Meta
	Name     string `gorm:"type:varchar(255)"`
	Content  string `gorm:"type:varchar(255)"`
	Depend   string `gorm:"type:varchar(255)"`
	Father   string `gorm:"type:varchar(255)"`
	FatherID uint   `gorm:"type:int(11) unsigned"` // 0 means root
	Done     bool
	Status   string `gorm:"type:varchar(255)"`
	Region   string `gorm:"type:varchar(255)"`
	Note     string `gorm:"type:text"`
	Tags     string `gorm:"type:text"`
}
