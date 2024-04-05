package dao

type TRecord struct {
	*Meta   `json:"*_meta,omitempty"`
	Name    string  `json:"name" gorm:"type:varchar(100);not null" json:"name,omitempty"`
	Weight  float32 `json:"weight" gorm:"type:float;not null" json:"weight,omitempty"`
	IsFuck  bool    `json:"is_fuck" gorm:"type:boolean;not null" json:"is_fuck,omitempty"`
	Vol1    string  `json:"vol1" gorm:"type:varchar(100)" json:"vol_1,omitempty"`
	Vol2    string  `json:"vol2" gorm:"type:varchar(100)" json:"vol_2,omitempty"`
	Vol3    string  `json:"vol3" gorm:"type:varchar(100)" json:"vol_3,omitempty"`
	Vol4    string  `json:"vol4" gorm:"type:varchar(100)" json:"vol_4,omitempty"`
	Content string  `json:"content" gorm:"type:text" json:"content,omitempty"`
	Region  string  `json:"region" gorm:"type:varchar(100)" json:"region,omitempty"`
	Retire  int     `gorm:"type:int" json:"retire,omitempty"`
	Cost    int     `gorm:"type:int" json:"cost,omitempty"`
	Dev     string  `gorm:"type:varchar(100)" json:"dev,omitempty"`
	Coding  string  `gorm:"type:varchar(100)" json:"coding,omitempty"`
	Social  string  `gorm:"type:varchar(100)" json:"social,omitempty"`
}
