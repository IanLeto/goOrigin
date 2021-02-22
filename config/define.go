package config

var GlobalConfig *Config

func init() {
	GlobalConfig = NewConfig()
}
