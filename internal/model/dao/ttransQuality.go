package dao

import "time"

const TableNameEcampProjectInfoTb = "ecamp_project_info_tb"

// EcampProjectInfoTb mapped from table <ecamp_project_info_tb>
type EcampProjectInfoTb struct {
	*Meta
	Project   string    `gorm:"column:project;type:varchar(253);not null;uniqueIndex:az_pod_proj_uniqe,priority:3" json:"project"`
	ProjectCn string    `gorm:"column:project_cn;type:varchar(253);not null" json:"project_cn"`
	Az        string    `gorm:"column:az;type:varchar(253);not null;uniqueIndex:az_pod_proj_uniqe,priority:1" json:"az"`
	TracePod  string    `gorm:"column:trace_pod;type:varchar(253);not null;uniqueIndex:az_pod_proj_uniqe,priority:2" json:"trace_pod"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);not null" json:"updated_at"`
}

// TableName EcampProjectInfoTb's table name
func (*EcampProjectInfoTb) TableName() string {
	return TableNameEcampProjectInfoTb
}

type EcampTransTypeTb struct {
	*Meta
	TransType   string    `gorm:"size:50;not null;uniqueIndex:trans_proj_unique,priority:1" json:"trans_type"`
	TransTypeCN string    `gorm:"size:100" json:"trans_type_cn"`                                             // 开户7
	Project     string    `gorm:"size:255;not null;uniqueIndex:trans_proj_unique,priority:2" json:"project"` // 组合唯一索引
	IsAlert     bool      `gorm:"default:false" json:"is_alert"`
	Threshold   int       `gorm:"default:0" json:"threshold"` // 新增阈值字段
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime(3);not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime(3);not null" json:"updated_at"`
}

type EcampReturnCodeTb struct {
	*Meta
	TransType  string    `gorm:"size:50;not null;index:idx_trans_proj,priority:1" json:"trans_type"`
	ReturnCode string    `gorm:"size:50;not null;index:idx_trans_proj,priority:3" json:"return_code"`
	Project    string    `gorm:"size:255;not null;index:idx_trans_proj,priority:2" json:"project"`
	Status     string    `gorm:"size:50;default:'active'" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(3);not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(3);not null" json:"updated_at"`
}

// EcampTransTypeCNTb  表结构体（简化版）
type EcampTransTypeCNTb struct {
	*Meta
	SysName         string    `gorm:"size:100;not null" json:"sys_name"`         // 系统名称代码（相当于project）
	ServiceName     string    `gorm:"size:100;not null" json:"service_name"`     // 服务名称代码
	InterfaceEnname string    `gorm:"size:255;not null" json:"interface_enname"` // 接口英文名称
	InterfaceName   string    `gorm:"size:255;not null" json:"interface_name"`   // 接口中文名称
	Url             string    `gorm:"size:255;not null" json:"url"`              // URL（相当于trans_type）
	CreatedAt       time.Time `gorm:"type:datetime(3);not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:datetime(3);not null" json:"updated_at"`
}

// TableNameEcampTransTypeCNTb TableName 指定表名
const TableNameEcampTransTypeCNTb = "ecamp_trans_infoCN_tb"

func (*EcampTransTypeCNTb) TableName() string {
	return TableNameEcampTransTypeCNTb
}

const TableNameEcampTransTypeTb = "ecamp_trans_info_tb"

func (*EcampTransTypeTb) TableName() string {
	return TableNameEcampTransTypeTb
}

const TableNameEcampReturnCodeTb = "ecamp_return_code_tb"

func (*EcampReturnCodeTb) TableName() string {
	return TableNameEcampReturnCodeTb
}

type ESAggResponse struct {
	Aggregations struct {
		ReqURLAgg struct {
			Buckets []struct {
				Key           string `json:"key"`       // 聚合键：request.url
				DocCount      int64  `json:"doc_count"` // 当前 URL 的总文档数
				ReturnCodeAgg struct {
					Buckets []struct {
						Key      string `json:"key"`       // 聚合键：return_code
						DocCount int64  `json:"doc_count"` // 每个 return_code 的文档数
					} `json:"buckets"`
				} `json:"returnCodeAgg"` // 子聚合：return_code 聚合
			} `json:"buckets"`
		} `json:"reqUrlAgg"` // 顶层聚合：按 request.url 进行聚合
	} `json:"aggregations"`
}

type EcampAggUrlDoc struct {
	ReqURL        string                   `json:"req_url"`         // 请求路径
	TotalCount    int64                    `json:"total_count"`     // 总调用次数7
	ReturnCodeAgg []EcampReturnCodeAggItem `json:"return_code_agg"` // 每个 return_code 的调用次数
}

type EcampReturnCodeAggItem struct {
	ReturnCode string `json:"return_code"` // 返回码
	Count      int64  `json:"count"`       // 调用次数
}

type AggUrlPathDoc struct {
	Took         int                    `json:"took"`
	TimedOut     bool                   `json:"timed_out"`
	Shards       AggUrlPathShards       `json:"_shards"`
	Hits         AggUrlPathHits         `json:"hits"`
	Aggregations AggUrlPathAggregations `json:"aggregations"`
}

type AggUrlPathShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type AggUrlPathHits struct {
	Total    AggUrlPathTotal `json:"total"`
	MaxScore interface{}     `json:"max_score"`
	Hits     []interface{}   `json:"hits"`
}

type AggUrlPathTotal struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type AggUrlPathAggregations struct {
	ByTransType AggUrlPathByTransType `json:"by_transType"`
}

type AggUrlPathByTransType struct {
	DocCountErrorUpperBound int              `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int              `json:"sum_other_doc_count"`
	Buckets                 []AggUrlPathInfo `json:"buckets"`
}

type AggUrlPathInfo struct {
	Key          string                 `json:"key"`
	DocCount     int                    `json:"doc_count"`
	ByReturnCode AggUrlPathByReturnCode `json:"by_return_code"`
}

type AggUrlPathByReturnCode struct {
	DocCountErrorUpperBound int                        `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int                        `json:"sum_other_doc_count"`
	Buckets                 []AggUrlPathReturnCodeInfo `json:"buckets"`
}

type AggUrlPathReturnCodeInfo struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

type AggProjectUrlPathDoc struct {
	Aggregation struct {
		AggsByTime struct {
			Buckets []AggsByTimeBucket `json:"buckets"`
		} `json:"aggs_by_time"`
	} `json:"aggregation"`
}

type AggsByTimeBucket struct {
	Key          int64           `json:"key"`          // 时间戳（毫秒）
	DocCount     int             `json:"doc_count"`    // 当前时间段的文档数
	GroupByField GroupByFieldAgg `json:"group_by_url"` // 嵌套聚合结果（比如按 URL 聚合）
}

// GroupByFieldAgg 嵌套维度聚合结果结构（如按 URL、项目、返回码等聚合）
type GroupByFieldAgg struct {
	Buckets []GroupByFieldBucket `json:"buckets"`
}

type GroupByFieldBucket struct {
	Key      string `json:"key"`       // URL、项目、返回码等
	DocCount int    `json:"doc_count"` // 每个值的文档数
}
