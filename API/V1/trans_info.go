package V1

type SuccessRateReqInfo struct {
}

type CreateTransInfoReq struct {
	*CreateTransInfo
	Region string `json:"region"`
}

type CreateTransInfo struct {
	Cluster     string            `json:"cluster"`
	Project     string            `json:"project"`
	TransType   map[string]string `json:"trans_type"`
	ServiceCode map[string]string `json:"service_code"`
	Interval    int               `json:"interval"`
	Dimension1  string            `json:"Dimension1"` // 交易类型
	Dimension2  string            `json:"Dimension2"` // 交易渠道
}

type CreateTransInfoResponse struct {
	Id uint `json:"id"`
}
