package config

import "encoding/json"

type Component interface {
	ToJSON() string
}

type V2Config struct {
	Base   BaseConfig   `yaml:"base" json:"base"`
	Logger LoggerConfig `yaml:"logger" json:"logger"`
	Env    EnvConfig    `yaml:"env" json:"env"`
}

type BaseConfig struct {
	Name string `yaml:"name" json:"name"`
	Port int    `yaml:"port" json:"port"`
	Mode string `yaml:"mode" json:"mode"`
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

type EnvConfig struct {
	Window EnvWindowConfig `yaml:"window" json:"window"`
	Mac    EnvMacConfig    `yaml:"mac" json:"mac"`
	Prod   EnvProdConfig   `yaml:"prod" json:"prod"`
}

type EnvWindowConfig struct {
	MySQL EnvWindowMySQLConfig `yaml:"mysql" json:"mysql"`
	ES    EnvWindowESConfig    `yaml:"es" json:"es"`
}

type EnvWindowMySQLConfig struct {
	Local EnvWindowMySQLLocalConfig `yaml:"local" json:"local"`
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

func (c *EnvConfig) ToJSON() string {
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	return string(jsonData)
}
