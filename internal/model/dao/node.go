package dao

type TNode struct {
	*Meta
	Name     string `gorm:"type:varchar(255)"`
	Content  string `gorm:"type:varchar(255)"`
	Depend   string `gorm:"type:varchar(255)"`
	ParentID uint   `gorm:"type:int(11) unsigned"` // 0 means root
	Done     bool   `gorm:"type:tinyint(1)"`
	Status   string `gorm:"type:varchar(255)"`
	Region   string `gorm:"type:varchar(255)"`
	Note     string `gorm:"type:text"`
	Tags     string `gorm:"type:text"`
}
