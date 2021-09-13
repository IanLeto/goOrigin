package main

import (
	"goOrigin/cmd/event"
	"goOrigin/config"
	"goOrigin/pkg/utils"
	"os"
)

var defaultConfigPath = "/Users/ian/go/src/goOrigin/config.yaml"

var preCheck []func() error
var mode string

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

// step 4 初始化组件
var initComponents = func() error {
	for _, component := range config.Conf.Components {
		event.Bus.Pub(component)
	}
	return nil
}

// step 5 启动模式

var initMode = func() error {
	switch mode {
	default:
		event.Bus.Pub("goDebug")
	}
	return nil
}

//var initMongoSession = func() error {
//	event.Bus.Pub("")
//	storage.InitMongo()
//	return nil
//}

func PreRun() string {
	for _, f := range preCheck {
		utils.CheckNoError(f())
	}
	return mode
}

func init() {
	preCheck = append(preCheck, initEvent, envCheck, initConfig, initComponents, initMode)
}
