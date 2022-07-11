package params

type CreateIanRequestInfo struct {
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

type QueryRequest struct {
	Name string `json:"name"`
}

type QueryResponse struct {
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

type AppendRequestInfo struct {
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