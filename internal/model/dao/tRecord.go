package dao

type TRecord struct {
	*Meta   `json:"*_meta,omitempty"`
	Name    string  `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Weight  float32 `json:"weight" gorm:"column:weight;type:float;not null"`
	IsFuck  bool    `json:"is_fuck" gorm:"column:is_fuck;type:boolean;not null"`
	Vol1    string  `json:"vol1" gorm:"column:vol1;type:text"`
	Vol2    string  `json:"vol2" gorm:"column:vol2;type:text"`
	Vol3    string  `json:"vol3" gorm:"column:vol3;type:text"`
	Vol4    string  `json:"vol4" gorm:"column:vol4;type:text"`
	Content string  `json:"content" gorm:"column:content;type:text"`
	Region  string  `json:"region" gorm:"column:region;type:varchar(100)"`
	Retire  int     `json:"retire" gorm:"column:retire;type:int"`
	Cost    int     `json:"cost" gorm:"column:cost;type:int"`
	Dev     string  `json:"dev" gorm:"column:dev;type:varchar(100)"`
	Coding  string  `json:"coding" gorm:"column:coding;type:varchar(100)"`
	Social  string  `json:"social" gorm:"column:social;type:varchar(100)"`
}

type DocRecord struct {
}
