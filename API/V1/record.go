package V1

type CreateIanRequestInfo struct {
	Name string `json:"name"`
	Body struct {
		Weight float32 `json:"weight"`
		BF     string  `json:"bf"`
		LUN    string  `json:"lun"`
		DIN    string  `json:"din"`
		EXTRA  string  `json:"extra"`
	} `json:"body"`
	BETre struct {
		Core       int `json:"core"`
		Runner     int `json:"runner"`
		Support    int `json:"support"`
		Squat      int `json:"squat"`
		EasyBurpee int `json:"easy_burpee"`
		Chair      int `json:"chair"`
		Stretch    int `json:"stretch"`
	} `json:"BETre"`
	Worker struct {
		Vol1 string `json:"vol1"`
		Vol2 string `json:"vol2"`
		Vol3 string `json:"vol3"`
		Vol4 string `json:"vol4"`
	} `json:"worker"`
	Content string            `json:"content"`
	Extra   map[string]string `json:"Extra"`
}

type QueryRequest struct {
	Name string `json:"name"`
}

type AppendRequestInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Body struct {
		Weight int `json:"weight"`
	} `json:"body"`
	BETre struct {
		Core       int `json:"core"`
		Runner     int `json:"runner"`
		Support    int `json:"support"`
		Squat      int `json:"squat"`
		EasyBurpee int `json:"easy_burpee"`
		Chair      int `json:"chair"`
		Stretch    int `json:"stretch"`
	} `json:"BETre"`
	Worker struct {
		Vol1 string `json:"vol1"`
		Vol2 string `json:"vol2"`
		Vol3 string `json:"vol3"`
		Vol4 string `json:"vol4"`
	} `json:"worker"`
}
type BatchCreateIanRecordRequest struct {
	Items []CreateIanRecordRequest `json:"items"`
}

type BatchCreateIanRecordResponse struct {
	Items []interface{} `json:"items"`
}

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
