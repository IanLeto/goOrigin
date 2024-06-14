package config

import (
	"encoding/json"
	"github.com/spf13/viper"
)

type Component interface {
	ToJSON() string
}

type V2Config struct {
	Base      BaseConfig                 `yaml:"base" json:"base"`
	Logger    LoggerConfig               `yaml:"logger" json:"logger"`
	Trace     TraceConfig                `yaml:"trace" json:"trace"`
	Component []string                   `yaml:"component" json:"component"`
	Env       map[string]ComponentConfig `yaml:"env" json:"env"`
}

type BaseConfig struct {
	Name string `yaml:"name" json:"name"`
	Port int    `yaml:"port" json:"port"`
	Mode string `yaml:"mode" json:"mode"`
}

func NewBaseConfig() BaseConfig {
	return BaseConfig{
		Name: viper.Get("base.name").(string),
		Port: viper.Get("base.port").(int),
		Mode: viper.Get("base.mode").(string),
	}
}

type LoggerConfig struct {
	Level   string `yaml:"level" json:"level"`
	Format  string `yaml:"format" json:"format"`
	Output  string `yaml:"output" json:"output"`
	File    string `yaml:"file" json:"file"`
	MaxSize int    `yaml:"maxsize" json:"maxsize"`
	Backup  int    `yaml:"backup" json:"backup"`
	MaxAge  int    `yaml:"max_age" json:"max_age"`
}

type TraceConfig struct {
	Az       string `yaml:"az" json:"az"`
	App      string `yaml:"app" json:"app"`
	Biz      string `yaml:"biz" json:"biz"`
	System   string `yaml:"system" json:"system"`
	Project  string `yaml:"project" json:"project"`
	Author   string `yaml:"author" json:"author"`
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	Interval string `yaml:"interval" json:"interval"`
	SvcName  string `yaml:"svcname" json:"svcname"`
}

func NewTraceConfig() TraceConfig {
	return TraceConfig{
		Az:       viper.GetString("trace.az"),
		App:      viper.GetString("trace.app"),
		Biz:      viper.GetString("trace.biz"),
		System:   viper.GetString("trace.system"),
		Project:  viper.GetString("trace.project"),
		Author:   viper.GetString("trace.author"),
		Endpoint: viper.GetString("trace.endpoint"),
		Interval: viper.GetString("trace.interval"),
		SvcName:  viper.GetString("trace.svcname"),
	}
}

type ConnConfig struct {
	Env map[string]interface{} `yaml:"env" json:"env"`
}
type Conn interface {
}

func NewComponentConfig() map[string]ComponentConfig {
	var (
		res = make(map[string]ComponentConfig)
	)
	for env, componentInfo := range viper.GetStringMap("env") {
		component := componentInfo.(map[string]interface{})
		res[env] = ComponentConfig{
			MysqlSQLConfig: MySQLConfig{
				DBName:   component["mysql"].(map[string]interface{})["dbname"].(string),
				User:     component["mysql"].(map[string]interface{})["user"].(string),
				Password: component["mysql"].(map[string]interface{})["password"].(string),
				Address:  component["mysql"].(map[string]interface{})["address"].(string),
			},
		}
	}
	return res
}

type EnvConfig struct {
	Env map[string]ComponentConfig `yaml:"env" json:"env"`
}

type ComponentConfig struct {
	MysqlSQLConfig MySQLConfig `yaml:"mysql" json:"mysql"`
}

func NewConnConfig() map[string]interface{} {
	//var (
	//	conns = make(map[string]interface{})
	//)

	return viper.Get("env").(map[string]interface{})
}

type EnvWindowConfig struct {
	MySQL EnvWindowMySQLConfig `yaml:"mysql" json:"mysql"`
	ES    EnvWindowESConfig    `yaml:"es" json:"es"`
}

type EnvWindowMySQLConfig struct {
	Local EnvWindowMySQLLocalConfig `yaml:"local" json:"local"`
}

func NewEnvWindowMySQLConfig() EnvWindowMySQLConfig {
	return EnvWindowMySQLConfig{
		Local: EnvWindowMySQLLocalConfig{
			Address: viper.Get("env.window.mysql.local.address").(string),
		},
	}
}

type MySQLConfig struct {
	DBName   string `yaml:"dbname" json:"dbname"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Address  string `yaml:"address" json:"address"`
}

type EnvWindowMySQLLocalConfig struct {
	Address string `yaml:"address" json:"address"`
}

type EnvWindowESConfig struct {
	Conn EnvWindowESConnConfig `yaml:"conn" json:"conn"`
}

type EnvWindowESConnConfig struct {
	Address string `yaml:"address" json:"address"`
}

type EnvMacConfig struct {
	MySQL EnvMacMySQLConfig `yaml:"mysql" json:"mysql"`
	ES    EnvMacESConfig    `yaml:"es" json:"es"`
}

type EnvMacMySQLConfig struct {
	Local EnvMacMySQLLocalConfig `yaml:"local" json:"local"`
}

type EnvMacMySQLLocalConfig struct {
	Address string `yaml:"address" json:"address"`
}

type EnvMacESConfig struct {
	Conn EnvMacESConnConfig `yaml:"conn" json:"conn"`
}

type EnvMacESConnConfig struct {
	Address string `yaml:"address" json:"address"`
}

type EnvProdConfig struct {
	MySQL EnvProdMySQLConfig `yaml:"mysql" json:"mysql"`
	ES    EnvProdESConfig    `yaml:"es" json:"es"`
}

type EnvProdMySQLConfig struct {
	Local EnvProdMySQLLocalConfig `yaml:"local" json:"local"`
}

type EnvProdMySQLLocalConfig struct {
	Address string `yaml:"address" json:"address"`
}

type EnvProdESConfig struct {
	Conn EnvProdESConnConfig `yaml:"conn" json:"conn"`
}

type EnvProdESConnConfig struct {
	Address string `yaml:"address" json:"address"`
}

func (c *Config) ToJSON() string {
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	return string(jsonData)
}

func (c *BaseConfig) ToJSON() string {
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	return string(jsonData)
}

func (c *LoggerConfig) ToJSON() string {
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	return string(jsonData)
}
