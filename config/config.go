package config

import "github.com/spf13/viper"

type ComponentConfig interface {
	NewComponent() ComponentConfig
}

// base Backend
type BackendConfig struct {
	*MySqlBackendConfig
}

func NewBackendConfig() *BackendConfig {
	return &BackendConfig{
		NewMySqlBackendConfig(),
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

// base backend
type BaseBackend struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (b *BaseBackend) GetType() string {
	panic("implement me")
}

func (b *BaseBackend) GetUrl() string {
	panic("implement me")
}

func (b *BaseBackend) GetAddress() string {
	panic("implement me")
}

func (b *BaseBackend) GetPort() string {
	panic("implement me")
}

func (b *BaseBackend) Close() error {
	panic("implement me")
}

type MySqlBackend struct {
	*BaseBackend
}

// 配置中心httpclient 配置参数

type CCConf struct {
	Address   string
	HeartBeat int
}

func NewCCClientConf() *CCConf {
	return &CCConf{
		Address:   viper.GetString("address"),
		HeartBeat: viper.GetInt("heart_beat"),
	}

}
