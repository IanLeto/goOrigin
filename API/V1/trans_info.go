package V1

type CreateTransInfoReq struct {
	Region string            `json:"region"` // 所属区域，批量共用
	Items  []CreateTransInfo `json:"items"`  // 多个交易类型
}

// CreateTransInfo 请求结构体也需要更新
type CreateTransInfo struct {
	Project     string            `json:"project" binding:"required"`
	TransType   string            `json:"trans_type" binding:"required"`
	TransTypeCN string            `json:"trans_type_cn"`
	ReturnCodes map[string]string `json:"return_codes"`
	IsAlert     bool              `json:"is_alert"`
	Threshold   int               `json:"threshold"`
}

type CreateTransInfoRes struct {
	Id uint `json:"id"`
}

type BatchCreateTransInfoRes struct {
	Success []CreateTransInfoRes `json:"success"` // 成功记录
	Failed  []FailedItem         `json:"failed"`  // 失败记录
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
type SearchTransInfoReq struct {
	Region    string `json:"region"`
	Keyword   string `json:"keyword"`
	Project   string `json:"project"`
	TransType string `json:"trans_type"`
	StartTime string `json:"start_time"` // 格式: 2006-01-02 15:04:05
	EndTime   string `json:"end_time"`   // 格式: 2006-01-02 15:04:05
	Page      int    `json:"page"`       // 当前页
	PageSize  int    `json:"page_size"`  // 每页数量
}

type SearchTransInfoRes struct {
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
	ReturnCode string `json:"return_code"`
	TransType  string `json:"trans_type"`
	Project    string `json:"project"`
	Status     string `json:"status"`
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

type TransTypeQueryInfo struct {
	Project    string   `json:"project" form:"project"`
	Az         string   `json:"az"`
	TransTypes []string `json:"trans_types" form:"trans_types"`
	StartTime  int      `json:"start_time"`
	EndTime    int      `json:"end_time"`
	Keyword    string   `json:"keyword"`
}

// SearchUrlPathWithReturnCodesReq 检索交易设置页列表2
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
	StartTime  int64    `json:"start_time"`
	EndTime    int64    `json:"end_time"`
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

	Svcname string `json:"svcname"`

	SourceIP string `json:"sourceIP"`
	TargetIP string `json:"targetIP"`

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

// 交易质量dashboard
type TransDashboardHisgromReq struct {
	Region    string `json:"region"`
	Az        string `json:"az"`
	Cluster   string `json:"cluster"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}

type TransDashboardHisgromRes struct {
}

// 交易质量指标req
type TransDashboardHisgromMetricReq struct {
	Region    string `json:"region"`
	Az        string `json:"az"`
	Cluster   string `json:"cluster"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}

type TransDashboardHisgromMetricRes struct {
}
