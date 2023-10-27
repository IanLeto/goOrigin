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

type QueryResponse struct {
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Body       struct {
		Weight float32 `json:"weight"`
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

type CreateIanRecordRequestInfo struct {
	Name       string            `json:"name"`
	Weight     float32           `json:"weight"`
	BF         string            `json:"bf"`
	LUN        string            `json:"lun"`
	DIN        string            `json:"din"`
	EXTRA      string            `json:"extra"`
	Core       int               `json:"core"`
	Runner     int               `json:"runner"`
	Support    int               `json:"support"`
	Squat      int               `json:"squat"`
	EasyBurpee int               `json:"easy_burpee"`
	Chair      int               `json:"chair"`
	Stretch    int               `json:"stretch"`
	Vol1       string            `json:"vol1"`
	Vol2       string            `json:"vol2"`
	Vol3       string            `json:"vol3"`
	Vol4       string            `json:"vol4"`
	Content    string            `json:"content"`
	Extra      map[string]string `json:"Extra"`
}

type CreateIanResponseRequestInfo struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Weight     float32 `json:"weight"`
	BF         string  `json:"bf"`
	LUN        string  `json:"lun"`
	DIN        string  `json:"din"`
	EXTRA      string  `json:"extra"`
	Core       int     `json:"core"`
	Runner     int     `json:"runner"`
	Support    int     `json:"support"`
	Squat      int     `json:"squat"`
	EasyBurpee int     `json:"easy_burpee"`
	Chair      int     `json:"chair"`
	Stretch    int     `json:"stretch"`
	Vol1       string  `json:"vol1"`
	Vol2       string  `json:"vol2"`
	Vol3       string  `json:"vol3"`
	Vol4       string  `json:"vol4"`
	Content    string  `json:"content"`
	CreateTime string  `json:"create_time"`
	UpdateTime string  `json:"update_time"`
}
