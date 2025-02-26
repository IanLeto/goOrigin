package dao

type TAchievementRecord struct {
	*Meta       `json:"*_meta,omitempty"`
	Name        string `json:"name" gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Description string `json:"description" gorm:"type:text" json:"description,omitempty"`
	Points      uint   `json:"points" gorm:"type:int unsigned;not null;default:0" json:"points,omitempty"`
	Type        string `json:"type" gorm:"type:enum('normal','daily');not null;default:'normal'" json:"type,omitempty"`
	AchievedAt  int64  `json:"achieved_at" gorm:"type:bigint" json:"achieved_at,omitempty"`
}
