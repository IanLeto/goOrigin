package entity

import "time"

type ProjectAggDocEntity struct {
	Took     int                 `json:"took"`
	TimedOut bool                `json:"timed_out"`
	Shards   ShardsInfo          `json:"_shards"`
	Hits     HitsInfo            `json:"hits"`
	Aggs     ProjectAggregations `json:"aggregations"`
}

type ShardsInfo struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type HitsInfo struct {
	Total    TotalHits `json:"total"`
	MaxScore float64   `json:"max_score"`
	Hits     []Hit     `json:"hits"`
}

type TotalHits struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type Hit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source struct {
		// 根据你的文档结构添加相应的字段
	} `json:"_source"`
}

type ProjectAggregations struct {
	ProjectStats ProjectStatsAggregation `json:"project_stats"`
	// 添加其他聚合结果的结构体
}

type ProjectStatsAggregation struct {
	DocCountErrorUpperBound int            `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int            `json:"sum_other_doc_count"`
	Buckets                 []ProjectStats `json:"buckets"`
}

type ProjectStats struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
	// 添加其他聚合度量的字段
}

type Project struct {
	ID        uint          `json:"id"`
	Project   string        `json:"project"`
	ProjectCN string        `json:"project_cn"`
	AZ        string        `json:"az"`
	TracePOD  string        `json:"trace_pod"`
	CreateAT  time.Duration `json:"create_at"`
	UpdateAT  time.Duration `json:"update_at"`
}

type ProjectBizCode struct {
	ID          uint   `json:"id"`
	BizKey      string `json:"biz_key"`
	BizValue    string `json:"biz_value"`
	BizType     string `json:"biz_type"`
	Project     string `json:"project"`
	Cluster     string `json:"cluster"`
	Service     string `json:"service"`
	SpanAlias   string `json:"span_alias"`
	CreatedUser string `json:"created_user"`
	UpdateUser  string `json:"update_user"`
}
