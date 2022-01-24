package model

type ConfInfo struct {
	ServiceID       string `gorm:"unique_index:tag"` // 之前通过 服务模板版本 和 配置组合来获取唯一配置
	Version         string `gorm:"unique_index:tag"` // 现在 通过 服务ID 和 Version 来获取唯一配置 即 服务版本 和 配置版本由多对多关系 转为一对一
	TemplateContent string  // 模板信息
	BizContent      string  // 该服务的业务数据
	// 公共配置 和 公共模块配置单独处理
}

type Req struct {
	ServiceName    string
	ServiceVersion string
	ConfVersion    string
}
