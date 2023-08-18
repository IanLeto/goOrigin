package elastic

type Doc struct {
}

// Aggregations is the struct for elasticsearch aggregation
type Aggregations struct {
	*Doc
	Aggregations map[string]Buckets `json:"aggregations"`
}

type Buckets struct {
	DocCountErrorUpperBound int64 `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int64 `json:"sum_other_doc_count"`
	Buckets                 []Bucket
}
type Bucket struct {
	Key      string `json:"key"`
	DocCount int64  `json:"doc_count"`
}
