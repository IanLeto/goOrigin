package config

var GlobalConfig *Config

func init() {
	GlobalConfig = NewConfig("/Users/ian/go/src/goOrigin/config.yaml")
}
