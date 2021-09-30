package main

import (
	"goOrigin/cmd/event"
	"goOrigin/config"
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

var connFactory = make([]storage.Conn, 0)

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
	if mode == "" {
		mode = config.Conf.RunMode
	}
	return nil
}

// step 4 初始化组件
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

// step 7 初始化factory
var initData = func() error {
	for _, d := range config.Conf.Data {
		switch d {
		case "mongo":
			connFactory = append(connFactory, storage.Mongo)
		}
	}
	// factory 行为执行
	for _, conn := range connFactory {
		if err := conn.InitData(mode); err != nil {
			return err
		}
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
	preCheck = append(preCheck, initEvent, envCheck,
		initConfig, initComponents, initMode, initData)
}
