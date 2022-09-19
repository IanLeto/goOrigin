package params

type QueryWeightRequest struct {
	Start  int64  `json:"start"`
	End    int64  `json:"end"`
	Metric string `json:"metric"`
	Labels struct {
	} `json:"labels"`
}
