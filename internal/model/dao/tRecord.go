package dao

type TRecord struct {
	*Meta     `json:"*_meta,omitempty"`
	Title     string  `json:"name" gorm:"column:name;type:varchar(100);not null"`
	MorWeight float32 `json:"mor_weight" gorm:"column:mor_weight;type:float"`
	NigWeight float32 `json:"nig_weight" gorm:"column:nig_weight;type:float"`
	Weight    float32 `json:"weight" gorm:"column:weight;type:float"`
	IsFuck    bool    `json:"is_fuck" gorm:"column:is_fuck;type:boolean;not null"`
	Vol1      string  `json:"vol1" gorm:"column:vol1;type:text"`
	Vol2      string  `json:"vol2" gorm:"column:vol2;type:text"`
	Vol3      string  `json:"vol3" gorm:"column:vol3;type:text"`
	Vol4      string  `json:"vol4" gorm:"column:vol4;type:text"`
	Content   string  `json:"content" gorm:"column:content;type:text"`
	Cost      int     `json:"cost" gorm:"column:cost;type:int"`
	Social    string  `json:"social" gorm:"column:content;type:text"`
}
