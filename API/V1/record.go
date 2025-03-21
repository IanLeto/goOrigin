package V1

type CreateIanRecordRequest struct {
	*CreateIanRecordRequestInfo
	Region string `json:"region"`
}

type CreateIanRecordRequestInfo struct {
	Name    string  `json:"name"`
	Weight  float32 `json:"weight"`
	IsFuck  bool    `json:"is_fuck"`
	Vol1    string  `json:"vol1"`
	Vol2    string  `json:"vol2"`
	Vol3    string  `json:"vol3"`
	Vol4    string  `json:"vol4"`
	Content string  `json:"content"`
	Cost    int     `json:"cost"`
	Dev     string  `json:"dev"`
	Coding  string  `json:"coding"`
	Social  string  `json:"social"`
}

type CreateIanRecordResponse struct {
	Id uint `json:"id"`
}

type QueryIanRecordsResponse struct {
	Items []interface{} `json:"items"`
}

type UpdateIanRecordRequest struct {
	Info RecordInfo `json:"info"`
}

type UpdateIanRecordResponse struct {
	//Item RecordInfo `json:"item"`
	ID uint `json:"id"`
}
type RecordInfo struct {
	Id         uint              `json:"id"`
	CreateTime int64             `json:"create_time"`
	ModifyTime int64             `json:"modify_time"`
	Name       string            `json:"name"`
	Weight     float32           `json:"weight"`
	EXTRA      string            `json:"extra"`
	Core       int               `json:"core"`
	Runner     int               `json:"runner"`
	Support    int               `json:"support"`
	Squat      int               `json:"squat"`
	Chair      int               `json:"chair"`
	Stretch    int               `json:"stretch"`
	Vol1       string            `json:"vol1"`
	Vol2       string            `json:"vol2"`
	Vol3       string            `json:"vol3"`
	Vol4       string            `json:"vol4"`
	Content    string            `json:"content"`
	Extra      map[string]string `json:"Extra"`
	Region     string            `json:"region"`
}
