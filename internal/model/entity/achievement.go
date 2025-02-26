package entity

type AchievementRecordEntity struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Points      uint   `json:"points" bson:"points"`
	Type        string `json:"type" bson:"type"`
	AchievedAt  int64  `json:"achieved_at" bson:"achieved_at"`
	CreateTime  int64  `json:"create_time" bson:"create_time"`
	ModifyTime  int64  `json:"update_time" bson:"update_time"`
}
