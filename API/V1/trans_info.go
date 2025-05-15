package V1

type SuccessRateReqInfo struct {
}

type BatchCreateTransInfoReq struct {
	Region string            `json:"region"` // 所属区域，批量共用
	Items  []CreateTransInfo `json:"items"`  // 多个交易类型
}

type CreateTransInfo struct {
	Cluster     string            `json:"cluster"`
	Project     string            `json:"project"`
	TransType   string            `json:"trans_type"`
	ServiceCode map[string]string `json:"service_code"` // return_code => return_code_cn
	Interval    int               `json:"interval"`
	Dimension1  string            `json:"dimension1"` // 交易类型
	Dimension2  string            `json:"dimension2"` // 交易渠道
}

type CreateTransInfoReq struct {
	*CreateTransInfo
	Region string `json:"region"`
}

type CreateTransInfoResponse struct {
	Id uint `json:"id"`
}

type BatchCreateTransInfoResponse struct {
	Success []CreateTransInfoResponse `json:"success"` // 成功记录
	Failed  []FailedItem              `json:"failed"`  // 失败记录
}
type FailedItem struct {
	TransType string `json:"trans_type"`
	Error     string `json:"error"`
}
