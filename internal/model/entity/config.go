package entity

const (
	TD = "td"
	SP = "sp"
	MS = "ms"
)

type ConfigEntity struct {
	KafkaConfig map[string]KafkaConfigEntity `json:"kafkaConfig,omitempty"`
}

type KafkaConfigEntity struct {
	Topic    string `json:"topic,omitempty"`
	Brokers  string `json:"brokers,omitempty"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	AuthType string `json:"authType,omitempty"`
	GroupID  string `json:"groupID,omitempty"`
	Domain   string `json:"domain,omitempty"`
}
