package entity

// RecordTimeInfo 结构体，存储所有时间格式
type RecordTimeInfo struct {
	CreateTimeISO8601   string `json:"create_time_iso8601"`    // ISO 8601 标准格式（UTC）
	CreateTimeISO8601CN string `json:"create_time_iso8601_cn"` // ISO 8601 北京时间（东八区）
	CreateTimeRFC822    string `json:"create_time_rfc822"`     // RFC 822 格式
	CreateTimeUnix      int64  `json:"create_time_unix"`       // Unix 时间戳（秒）
	CreateTimeUnixMs    int64  `json:"create_time_unix_ms"`    // Unix 时间戳（毫秒）
	CreateTimeDBFormat  string `json:"create_time_db"`         // 数据库常用格式
	CreateTimeCompact   string `json:"create_time_compact"`    // 紧凑格式（无分隔符）
}

// BalanceInfo 表示余额和消费相关的信息
type BalanceInfo struct {
	CurrentBalance float64 `json:"current_balance"` // 当前余额
	AvgDailyCost   float64 `json:"avg_daily_cost"`  // 平均每日支出（可通过历史数据计算得到）
	PredictDays    int     `json:"predict_days"`    // 可支撑天数（系统根据 balance / avg_cost 推算）
	UpdateTime     int64   `json:"update_time"`     // 更新时间戳
}

// RecordEntity 结构体，包含原始数据 + 时间信息 + 余额信息
type RecordEntity struct {
	ID        uint    `json:"id" bson:"_id"`
	Title     string  `json:"title" bson:"title"`
	MorWeight float32 `json:"mor_weight" bson:"weight"`
	NigWeight float32 `json:"nig_weight" bson:"weight"`

	IsFuck     bool           `json:"is_fuck"`
	Vol1       string         `json:"vol1" bson:"vol1"`
	Vol2       string         `json:"vol2" bson:"vol2"`
	Vol3       string         `json:"vol3" bson:"vol3"`
	Vol4       string         `json:"vol4" bson:"vol4"`
	Cost       int            `json:"cost" bson:"cost"` // 单日或一次性支出
	Content    string         `json:"content" bson:"content"`
	Coding     string         `json:"coding"`
	Social     string         `json:"social"`
	CreateTime int64          `json:"create_time"`
	ModifyTime int64          `json:"modify_time"`
	TimeInfo   RecordTimeInfo `json:"time_info"`

	Balance *BalanceInfo `json:"balance,omitempty"` // 新增余额信息，可选字段
}
