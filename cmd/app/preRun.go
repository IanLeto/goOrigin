package main

import (
	"goOrigin/cmd/event"
	"goOrigin/config"
	"goOrigin/pkg/logging"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
	"os"
)

var defaultConfigPath = "/Users/ian/go/src/goOrigin/config.yaml"

var preCheck []func() error
var mode string

var compInit = map[string]func() error{
	"mongo": storage.InitMongo,
	"zk":    storage.InitZk,
}

// step 1 本地环境变量检查
var envCheck = func() error {
	mode = os.Getenv("mode")
	return nil
}

// step 2 初始化 event
var initEvent = func() error {
	return event.InitEvent()
}

// step 3 本地配置文件检查
var initConfig = func() error {
	config.InitConf(defaultConfigPath)
	return nil
}

// step 4 初始化配置
var initLog = func() error {
	return logging.InitLogger()
}

// step 5 初始化组件
var initComponents = func() error {
	for _, component := range config.Conf.Components {
		if fn, ok := compInit[component]; ok {
			utils.NoError(fn())
		}
	}
	return nil
}

// step 6 启动模式

var initMode = func() error {
	switch mode {
	default:
		event.Bus.Pub("goDebug")
	}
	return nil
}

func PreRun() string {
	for _, f := range preCheck {
		utils.NoError(f())
	}
	return mode
}

func init() {
	preCheck = append(preCheck, initEvent, envCheck, initConfig, initLog, initComponents, initMode)
}
