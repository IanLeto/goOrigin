package model

type Doc struct {
}

// Aggregations is the struct for elasticsearch aggregation
type Aggregations struct {
	*Doc
	Aggregations map[string]Buckets `json:"aggregations"`
}

type Buckets struct {
	Buckets []Bucket
}
type Bucket struct {
	Key      string `json:"key"`
	DocCount int64  `json:"doc_count"`
}
