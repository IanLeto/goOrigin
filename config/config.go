package config

import "github.com/spf13/viper"

type ComponentConfig interface {
	NewComponent() ComponentConfig
}

// base Backend
type BackendConfig struct {
	*MySqlBackendConfig
	*MongoBackendConfig
}

func NewBackendConfig() *BackendConfig {
	return &BackendConfig{
		NewMySqlBackendConfig(),
		NewMongoBackendConfig(),
	}
}

// base Http client conf

type HttpClientConfig struct {
	CC *CCConf
}

func NewHttpClientConfig() *HttpClientConfig {
	return &HttpClientConfig{
		CC: NewCCClientConf(),
	}
}

// mysql backend config
type MySqlBackendConfig struct {
	Address  string
	Port     string
	Password string
	User     string
}

func NewMySqlBackendConfig() *MySqlBackendConfig {
	return &MySqlBackendConfig{
		Address:  viper.GetString("backend.MySql.address"),
		Port:     viper.GetString("backend.MySql.port"),
		User:     viper.GetString("backend.MySql.user"),
		Password: viper.GetString("backend.MySql.password"),
	}
}

// mongoDB
type MongoBackendConfig struct {
	Address  string
	Port     string
	Password string
	User     string
	DB       string
}

func NewMongoBackendConfig() *MongoBackendConfig {
	return &MongoBackendConfig{
		Address:  viper.GetString("backend.mongo.address"),
		Port:     viper.GetString("backend.mongo.port"),
		User:     viper.GetString("backend.mongo.user"),
		Password: viper.GetString("backend.mongo.password"),
		DB:       viper.GetString("backend.mongo.DB"),
	}
}

// 配置中心httpclient 配置参数

type CCConf struct {
	Address   string
	HeartBeat int
}

func NewCCClientConf() *CCConf {
	return &CCConf{
		Address:   viper.GetString("client.CC.address"),
		HeartBeat: viper.GetInt("client.CC.heart_beat"),
	}

}
