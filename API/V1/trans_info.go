package V1

type CreateTransInfoReq struct {
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

type DeleteTransInfoReq struct {
	Region    string `json:"region"`
	Project   string `json:"project"`
	TransType string `json:"trans_type"`
}
type GetTransInfoListReq struct {
	Region    string `json:"region"`
	Project   string `json:"project"`
	TransType string `json:"trans_type"`
	Page      int    `json:"page"`      // 当前页
	PageSize  int    `json:"page_size"` // 每页数量
}

type GetTransInfoListResponse struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type UpdateTransInfoReq struct {
	Region string           `json:"region"`
	Item   *UpdateTransInfo `json:"item"`
}

type UpdateTransInfo struct {
	Project     string              `json:"project"`
	TransType   string              `json:"trans_type"`
	Interval    int                 `json:"interval"`
	ReturnCodes []*UpdateReturnCode `json:"return_codes"`
}

type UpdateReturnCode struct {
	ReturnCode   string `json:"return_code"`
	ReturnCodeCn string `json:"return_code_cn"`
	TransType    string `json:"trans_type"`
	Project      string `json:"project"`
	Status       string `json:"status"`
}

type CreateTradeReturnCodeRequest struct {
	UrlPath       string `json:"url_path"`
	SuccessCount  int    `json:"success_count"`
	FailedCount   int    `json:"failed_count"`
	UnKnownCount  string `json:"unknown_count"`
	Total         string `json:"total"`
	TransTypeCn   string `json:"trans_type_cn"`
	ResponseCount string `json:"response_count"`
}

// 搜索es 列表使用
type SearchTradeReturnCodeRequest struct {
	UrlPath       string `json:"url_path"`
	SuccessCount  int    `json:"success_count"`
	FailedCount   int    `json:"failed_count"`
	UnKnownCount  int    `json:"unknown_count"`
	Total         int    `json:"total"`
	TransTypeCn   string `json:"trans_type_cn"`
	ResponseCount int    `json:"response_count"`
}

type SuccessRateReqInfo struct {
	Project    string   `form:"project" json:"project"`         // 项目标识
	TransTypes []string `form:"trans_types" json:"trans_types"` // 交易类型列表
	StartTime  int64    `form:"start_time" json:"start_time"`   // 开始时间（可扩展）
	EndTime    int64    `form:"end_time" json:"end_time"`       // 结束时间
	Region     string   `form:"region" json:"region"`           // 查询区域
}

type SuccessRateItem struct {
	TransType     string `json:"trans_type"`
	TransTypeCn   string `json:"trans_type_cn"`
	SuccessCount  int    `json:"success_count"`
	FailedCount   int    `json:"failed_count"`
	UnknownCount  int    `json:"unknown_count"`
	Total         int    `json:"total"`
	ResponseCount int    `json:"response_count"`
}

type SuccessRateResponse struct {
	Items []*SuccessRateItem `json:"items"`
}

// 检索交易设置页列表
type TransTypeQueryReq struct {
	Region   string `json:"region" form:"region"`
	Page     int    `json:"page" form:"page"`           // 当前页码，从 1 开始
	PageSize int    `json:"page_size" form:"page_size"` // 每页大小，默认 10
	*TransTypeQueryInfo
}

type TransTypeQueryInfo struct {
	Project    string   `json:"project" form:"project"`
	Az         string   `json:"az"`
	TransTypes []string `json:"trans_types" form:"trans_types"`
	StartTime  int      `json:"start_time"`
	EndTime    int      `json:"end_time"`
	Keyword    string   `json:"keyword"`
	OrderBy    string   `json:"order_by"`
}

// 检索交易设置页列表2
type SearchUrlPathWithReturnCodesReq struct {
	Region   string `json:"region" form:"region"`
	Page     int    `json:"page" form:"page"`           // 当前页码，从 1 开始
	PageSize int    `json:"page_size" form:"page_size"` // 每页大小，默认 10
	*SearchUrlPathWithReturnCodesInfo
}

type SearchUrlPathWithReturnCodesInfo struct {
	Project    string   `json:"project" form:"project"`
	Az         string   `json:"az"`
	TransTypes []string `json:"trans_types" form:"trans_types"`
	StartTime  int      `json:"start_time"`
	EndTime    int      `json:"end_time"`
	Keyword    string   `json:"keyword"`
	OrderBy    string   `json:"order_by"`
}

// 交易聚合查询接口
type AggUrlPathWithReturnCodesReq struct {
	Region   string `json:"region" form:"region"`
	Page     int    `json:"page" form:"page"`           // 当前页码，从 1 开始
	PageSize int    `json:"page_size" form:"page_size"` // 每页大小，默认 10
	*SearchUrlPathWithReturnCodesInfo
}

type AggUrlPathWithReturnCodesInfo struct {
	Project string `json:"project" form:"project"`
	Az      string `json:"az"`

	Svcname    string `json:"svcname"`
	Dimension1 string `json:"dimension1"`
	Dimension2 string `json:"dimension2"`
	SourceIP   string `json:"sourceIP"`
	TargetIP   string `json:"targetIP"`

	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
	OrderBy   string `json:"order_by"`
}

// 交易页面返回值
type SearchUrlPathWithReturnCodesInfoResponse struct {
	Items    []interface{} `json:"items"`
	Total    int           `json:"total"`     // 总条数
	Page     int           `json:"page"`      // 当前页
	PageSize int           `json:"page_size"` // 单页条数

}

// 交易页面返回值
type TransTypeResponse struct {
	Items    []*TransTypeItem `json:"items"`
	Total    int              `json:"total"`     // 总条数
	Page     int              `json:"page"`      // 当前页
	PageSize int              `json:"page_size"` // 单页条数

}

type TransTypeItem struct {
	TransType   string   `json:"trans_type"`
	TransTypeCn string   `json:"trans_type_cn"`
	ReturnCode  []string `json:"return_code"`
}
type TransTypeResponse2 struct {
	Items    []*TransTypeItem2 `json:"items"`
	Total    int               `json:"total"`     // 总条数
	Page     int               `json:"page"`      // 当前页
	PageSize int               `json:"page_size"` // 单页条数

}

type TransTypeItem2 struct {
	UrlPath         string            `json:"url_path"`
	UrlPathCN       string            `json:"url_path_cn"`
	ReturnCode      map[string]string `json:"return_code"`
	ReturnCodeCount map[string]int    `json:"return_code_count"`
}
