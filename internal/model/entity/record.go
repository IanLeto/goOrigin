package entity

type RecordEntity struct {
	Title      string  `json:"title" bson:"title"`
	Weight     float32 `json:"weight" bson:"weight"`
	IsFuck     bool    `json:"is_fuck"`
	Vol1       string  `json:"vol1" bson:"vol1"`
	Vol2       string  `json:"vol2" bson:"vol2"`
	Vol3       string  `json:"vol3" bson:"vol3"`
	Vol4       string  `json:"vol4" bson:"vol4"`
	Cost       int     `json:"cost" bson:"cost"`
	Content    string  `json:"content" bson:"content"`
	Region     string  `json:"region" bson:"region"`
	Dev        string  `json:"dev"`
	Coding     string  `json:"coding"`
	Social     string  `json:"social"`
	CreateTime int64   `json:"create_time"`
	ModifyTime int64   `json:"update_time"`
}
