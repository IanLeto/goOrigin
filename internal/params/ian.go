package params

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
	Extra map[string]string `json:"Extra"`
}

type QueryRequest struct {
	Name string `json:"name"`
}

type QueryResponse struct {
	Name string `json:"name"`
	Body struct {
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
