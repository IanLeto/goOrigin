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

// RecordEntity 结构体，包含原始数据 + 时间信息
type RecordEntity struct {
	Title      string         `json:"title" bson:"title"`
	Weight     float32        `json:"weight" bson:"weight"`
	IsFuck     bool           `json:"is_fuck"`
	Vol1       string         `json:"vol1" bson:"vol1"`
	Vol2       string         `json:"vol2" bson:"vol2"`
	Vol3       string         `json:"vol3" bson:"vol3"`
	Vol4       string         `json:"vol4" bson:"vol4"`
	Cost       int            `json:"cost" bson:"cost"`
	Content    string         `json:"content" bson:"content"`
	Region     string         `json:"region" bson:"region"`
	Dev        string         `json:"dev"`
	Coding     string         `json:"coding"`
	Social     string         `json:"social"`
	CreateTime int64          `json:"create_time"`
	ModifyTime int64          `json:"modify_time"` // 修正 UpdateTime 命名
	TimeInfo   RecordTimeInfo `json:"time_info"`
}
