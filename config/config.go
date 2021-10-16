package config

import "github.com/spf13/viper"

type ComponentConfig interface {
	NewComponent() ComponentConfig
}

// base Backend
type BackendConfig struct {
	*MySqlBackendConfig
	*MongoBackendConfig
	*ZKConfig
}

func NewBackendConfig() *BackendConfig {
	return &BackendConfig{
		NewMySqlBackendConfig(),
		NewMongoBackendConfig(),
		NewZkConfig(),
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
	Name     string
}

func NewMySqlBackendConfig() *MySqlBackendConfig {
	return &MySqlBackendConfig{
		Address:  viper.GetString("backend.MySql.address"),
		Port:     viper.GetString("backend.MySql.port"),
		User:     viper.GetString("backend.MySql.user"),
		Password: viper.GetString("backend.MySql.password"),
		Name:     viper.GetString("backend.MySql.name"),
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

// zookeeper 配置
type ZKConfig struct {
	Address []string
	Master  string
	Auth    string
}

func NewZkConfig() *ZKConfig {
	return &ZKConfig{
		Address: viper.GetStringSlice("backend.zk.Address"),
		Master:  viper.GetString("backend.zk.Master"),
		Auth:    viper.GetString("backend.zk.Auth"),
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

// logging 配置
type LoggingConfig struct {
	FileName string
	Level    string
	Path     string
	Rotation RotationConfig
}

type RotationConfig struct {
	Time  int
	Count int
}

func NewLoggingConfig() *LoggingConfig {
	return &LoggingConfig{
		FileName: viper.GetString("logging.fileName"),
		Level:    viper.GetString("logging.level"),
		Path:     viper.GetString("logging.path"),
		Rotation: RotationConfig{
			Time:  viper.GetInt("logging.rotation.time"),
			Count: viper.GetInt("logging.rotation.Count"),
		},
	}
}

// ssh
type SSHConfig struct {
	Address string
	User    string
	Type    string
	KeyPath string // ssh_id 路径
	Port    int
	Auth    string
}

func NewSSHConfig() *SSHConfig {
	return &SSHConfig{
		Address: viper.GetString("ssh.address"),
		User:    viper.GetString("ssh.user"),
		Type:    viper.GetString("ssh.type"),
		KeyPath: viper.GetString("ssh.key_path"),
		Auth:    viper.GetString("ssh.auth"),
		Port:    viper.GetInt("ssh.port"),
	}
}
