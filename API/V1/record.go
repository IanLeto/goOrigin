package V1

type CreateIanRecordRequest struct {
	*CreateIanRecordRequestInfo
	Region string `json:"region"`
}

type CreateIanRecordRequestInfo struct {
	Name      string  `json:"name"`
	Weight    float32 `json:"weight"`
	IsFuck    bool    `json:"is_fuck"`
	Vol1      string  `json:"vol1"`
	Vol2      string  `json:"vol2"`
	Vol3      string  `json:"vol3"`
	Vol4      string  `json:"vol4"`
	Content   string  `json:"content"`
	Cost      int     `json:"cost"`
	Coding    string  `json:"coding"`
	Social    string  `json:"social"`
	NigWeight float32 `json:"nigWeight"`
}

type CreateIanRecordResponse struct {
	Id uint `json:"id"`
}

type QueryIanRecordsResponse struct {
	Items []interface{} `json:"items"`
}

type UpdateIanRecordRequest struct {
	Info CreateIanRecordRequestInfo `json:"info"`
}

type UpdateIanRecordResponse struct {
	ID uint `json:"id"`
}
