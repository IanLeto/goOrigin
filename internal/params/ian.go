package params

type CreateIanRequestInfo struct {
	Body struct {
		Weight int `json:"weight"`
	} `json:"body"`
	BETre struct {
		Runner  int `json:"runner"`
		Core    int `json:"core"`
		Support int `json:"support"`
	} `json:"BETre"`
	Worker struct {
		Vol1 string `json:"vol1"`
		Vol2 string `json:"vol2"`
		Vol3 string `json:"vol3"`
		Vol4 string `json:"vol4"`
	} `json:"worker"`
}
