package V1

type CreateIanRecordRequest struct {
	*CreateIanRecordRequestInfo
	Region string `json:"region"`
}

type CreateIanRecordRequestInfo struct {
	Title     string  `json:"title"`
	MorWeight float32 `json:"mor_weight"`
	NigWeight float32 `json:"nig_weight"`
	IsFuck    bool    `json:"is_fuck"`
	Vol1      string  `json:"vol1"`
	Vol2      string  `json:"vol2"`
	Vol3      string  `json:"vol3"`
	Vol4      string  `json:"vol4"`
	Content   string  `json:"content"`
	Cost      int     `json:"cost"`
	Coding    string  `json:"coding"`
	Social    string  `json:"social"`
}

type CreateIanRecordResponse struct {
	Id uint `json:"id"`
}

type QueryIanRecordsResponse struct {
	Items []interface{} `json:"items"`
}

type UpdateIanRecordRequest struct {
	ID uint `json:"id"`
	*CreateIanRecordRequestInfo
	Region string `json:"region"`
}

type UpdateIanRecordResponse struct {
	ID uint `json:"id"`
}

type QueryPodHistoryRes struct {
	Data []PodHistoryInfo
}

type PodHistoryInfo struct {
	HostIP  string
	PodName string
	PodIP   string
	Time    int64
	CPU     string
	Memory  string
}
