package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Component interface {
	ToJSON() string
	InitSelf() Component
}

type V2Config struct {
	Base      BaseConfig                 `yaml:"base" json:"base"`
	Logger    LoggerConfig               `yaml:"logger" json:"logger"`
	Trace     TraceConfig                `yaml:"trace" json:"trace"`
	Component []string                   `yaml:"component" json:"component"`
	Env       map[string]ComponentConfig `yaml:"env" json:"env"`
	Cron      []string                   `yaml:"cron" json:"cron"`
	//Cron      map[string]interface{}     `yaml:"cron" json:"cron"`
}

type BaseConfig struct {
	Name   string `yaml:"name" json:"name"`
	Port   int    `yaml:"port" json:"port"`
	Mode   string `yaml:"mode" json:"mode"`
	Region string `yaml:"region" json:"region"`
}

func NewBaseConfig() BaseConfig {
	return BaseConfig{
		Name:   viper.Get("base.name").(string),
		Port:   viper.Get("base.port").(int),
		Mode:   viper.Get("base.mode").(string),
		Region: viper.Get("base.region").(string),
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

func NewComponentConfig() map[string]ComponentConfig {
	var (
		res = make(map[string]ComponentConfig)
	)
	envMap := viper.GetStringMap("env")
	for env, componentInfo := range envMap {
		component, ok := componentInfo.(map[string]interface{})
		if !ok {
			fmt.Printf("Invalid componentInfo format for environment %s", env)
			continue
		}
		var eph = ComponentConfig{}
		mysqlInfo, ok := component["mysql"].(map[string]interface{})
		if ok {
			mysqlConfInfo := initMysqlConfig(mysqlInfo)
			eph.MysqlSQLConfig = *mysqlConfInfo
		}
		esInfo, ok := component["es"].(map[string]interface{})
		if ok {
			esConfInfo := initEsConfig(esInfo)
			eph.EsConfig = *esConfInfo
		}
		cronInfo, ok := component["cronjob"].(map[string]interface{})
		if ok {
			cronjobConfig := initCronJobConfig(cronInfo)
			eph.CronJobConfig = *cronjobConfig
		}

		prome, ok := component["prom"].(map[string]interface{})
		if ok {
			promeConfig := initPromeConfig(prome)
			eph.PromeConfig = *promeConfig
		}
		res[env] = eph
	}
	return res
}

type EnvConfig struct {
	Env map[string]ComponentConfig `yaml:"env" json:"env"`
}

type ComponentConfig struct {
	MysqlSQLConfig MySQLConfig   `yaml:"mysql" json:"mysql"`
	EsConfig       ESConfigV2    `yaml:"es" json:"es"`
	CronJobConfig  CronjobConfig `yaml:"cronJob" json:"cronJob"`
	PromeConfig    PromeConfig   `yaml:"prome" json:"prome"`
}

type MySQLConfig struct {
	DBName      string `yaml:"dbname" json:"dbname"`
	User        string `yaml:"user" json:"user"`
	Password    string `yaml:"password" json:"password"`
	Address     string `yaml:"address" json:"address"`
	IsMigration bool   `yaml:"is_migration" json:"is_migration"`
}

func initMysqlConfig(mysql map[string]interface{}) *MySQLConfig {
	dbName, ok := mysql["dbname"].(string)
	if !ok {
		return nil
	}
	user, ok := mysql["user"].(string)
	if !ok {
		return nil
	}
	password, ok := mysql["password"].(string)
	if !ok {
		return nil
	}
	address, ok := mysql["address"].(string)
	if !ok {
		return nil
	}
	isMigration, ok := mysql["is_migration"].(bool)
	if !ok {
		return nil
	}
	return &MySQLConfig{
		DBName:      dbName,
		User:        user,
		Password:    password,
		Address:     address,
		IsMigration: isMigration,
	}
}

type ESConfigV2 struct {
	Address string `yaml:"address" json:"address"`
}

func initEsConfig(es map[string]interface{}) *ESConfigV2 {
	address, ok := es["address"].(string)
	if !ok {
		return nil
	}
	return &ESConfigV2{Address: address}
}

type PromeConfig struct {
	Address string `yaml:"address" json:"address"`
}

func initPromeConfig(prome map[string]interface{}) *PromeConfig {
	address, ok := prome["address"].(string)
	if !ok {
		return nil
	}

	return &PromeConfig{Address: address}
}
