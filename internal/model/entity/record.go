package entity

type Record struct {
	Name     string   `json:"name" bson:"name"`
	Weight   float32  `json:"weight" bson:"weight"`
	IsFuck   bool     `json:"is_fuck"`
	Vol1     string   `json:"vol1" bson:"vol1"`
	Vol2     string   `json:"vol2" bson:"vol2"`
	Vol3     string   `json:"vol3" bson:"vol3"`
	Vol4     string   `json:"vol4" bson:"vol4"`
	Cost     int      `json:"cost" bson:"cost"`
	Content  string   `json:"content" bson:"content"`
	Region   string   `json:"region" bson:"region"`
	Retire   int      `json:"retire" bson:"retire"`
	Dev      string   `json:"dev"`
	Coding   string   `json:"coding"`
	Reading  string   `json:"reading"`
	Social   string   `json:"social"`
	Property Property `json:"property"`
}

type Property struct {
	Credit        int `json:"credit"`
	ProvidentFund int `json:"providentFund"`
	Liabilities   int `json:"liabilities"`
}
