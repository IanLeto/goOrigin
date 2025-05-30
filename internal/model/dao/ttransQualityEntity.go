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
	TransType   string              `gorm:"size:50;not null;uniqueIndex"`              // 加 uniqueIndex
	TransTypeCN string              `gorm:"size:100"`                                  // 开户7
	Project     string              `gorm:"size:255;not null"`                         // 改为字符串，与 TransInfoEntity 一致
	ReturnCodes []EcampReturnCodeTb `gorm:"foreignKey:TransType;references:TransType"` // 通过字符串关联
	IsAlert     bool
	Dimension1  string `gorm:"size:100"` // 交易类型
	Dimension2  string `gorm:"size:100"` // 交易渠道
}

type EcampReturnCodeTb struct {
	*Meta
	TransType    string `gorm:"size:50;index"`    // 外键字段
	ReturnCode   string `gorm:"size:50;not null"` // SC001
	ReturnCodeCN string `gorm:"size:100"`         // 开户成功
	Project      string `gorm:"size:255"`         // 与 ReturnCodeEntity.ProjectID 匹配
	Status       string `gorm:"size:50"`
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
}

type AggUrlPathInfo struct {
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
