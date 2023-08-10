package params

type GetLogsReq struct {
	Container string `json:"container"`
	PodID     string `json:"pod_id"`
	Ns        string `json:"ns"`
	Cluster   string `json:"cluster"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	Size      int    `json:"size"`
	Location  string `json:"location"`
	Step      int    `json:"step"`

	LimitByte int `json:"limit_byte"`
	LimitLine int `json:"limit_line"`
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
