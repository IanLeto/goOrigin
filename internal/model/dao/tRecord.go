package dao

type TRecord struct {
	*Meta
	Name    string  `json:"name" gorm:"type:varchar(100);not null"`
	Weight  float32 `json:"weight" gorm:"type:float;not null"`
	IsFuck  bool    `json:"is_fuck" gorm:"type:boolean;not null"`
	Vol1    string  `json:"vol1" gorm:"type:varchar(100)"`
	Vol2    string  `json:"vol2" gorm:"type:varchar(100)"`
	Vol3    string  `json:"vol3" gorm:"type:varchar(100)"`
	Vol4    string  `json:"vol4" gorm:"type:varchar(100)"`
	Content string  `json:"content" gorm:"type:text"`
	Region  string  `json:"region" gorm:"type:varchar(100)"`
	Retire  int     `gorm:"type:int"`
	Cost    int     `gorm:"type:int"`
}
