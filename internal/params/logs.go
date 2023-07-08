package params

type GetLogsReq struct {
	Container string `json:"container"`
	PodID     string `json:"pod_id"`
	Ns        string `json:"ns"`
	Cluster   string `json:"cluster"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	Size      int    `json:"size"`
	Location  string `json:"Location"`
	Step      int    `json:"step"`
}

type GetLogsRes struct {
	Contents []interface{}
	FromDate string
	ToDate   string
}
type Entry struct {
	Timestamp int64  `json:"timestamp"`
	Date      string `json:"date"`
	Content   int    `json:"content"`
}
