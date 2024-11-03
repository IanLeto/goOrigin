package config

import "goOrigin/pkg/utils"

type CronjobConfig struct {
	TransferConfig *TransferConfig `json:"transfer"`
}

type DemoJobConfig struct {
	// 任务名称
	Name string `json:"name"`
	// 任务描述
	Desc string `json:"desc"`
	// 任务类型
	Type string `json:"type"`
	// 任务执行时间
	ExecTime string `json:"exec_time"`
	// 任务执行间隔
	Interval string `json:"interval"`
	// 任务执行超时时间
	Timeout string `json:"timeout"`
	// 任务执行次数
	ExecCount int `json:"exec_count"`
	// 任务执行失败次数
	ExecFailCount int `json:"exec_fail_count"`
	// 任务执行失败重试次数
	ExecFailRetryCount int `json:"exec_fail_retry_count"`
	// 任务执行失败重试间隔
	ExecFailRetryInterval string `json:"exec_fail_retry_interval"`
	// 任务执行失败重试超时时间
	ExecFailRetryTimeout string `json:"exec_fail_retry_timeout"`
	// 任务执行失败重试次数
}

type TransferConfig struct {
	Interval int `yaml:"interval"`
	Timeout  int `yaml:"timeout"`
	Retry    int `yaml:"retry"`
	MySQL    struct {
		DBName      string `yaml:"dbname"`
		Address     string `yaml:"address"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		IsMigration bool   `yaml:"is_migration"`
		Table       string `yaml:"table"`
	} `yaml:"mysql"`
	ES struct {
		Address string `yaml:"address"`
		Alias   string `yaml:"alias"`
	} `yaml:"es"`
}

func initCronJobConfig(input interface{}) *CronjobConfig {
	var (
		err           error
		cronjobConfig CronjobConfig
	)
	err = utils.JsonToStruct(input, &cronjobConfig)
	utils.NoError(err)
	return &cronjobConfig
}
