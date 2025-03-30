package V1

type SuccessRateReqInfo struct {
}

type CreateTransInfoReq struct {
	*CreateTransInfo
	Region string `json:"region"`
}

type CreateTransInfo struct {
	TraceID     string            `json:"trace_id"`
	Cluster     string            `json:"cluster"`
	PodName     string            `json:"pod_name"`
	Project     string            `json:"project"`
	TransType   map[string]string `json:"trans_type"`
	ServiceCode map[string]string `json:"service_code"`
	Interval    int               `json:"interval"`
}

type CreateTransInfoResponse struct {
	Id uint `json:"id"`
}
