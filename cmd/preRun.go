package cmd

import (
	"context"
	"fmt"
	"goOrigin/cmd/event"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/pkg"
	"goOrigin/pkg/cron"
	"goOrigin/pkg/k8s"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
	"os"
)

var preCheck []func() error
var mode string

// 初始化组件
var compInit = map[string]func() error{
	"mongo": storage.InitMongo,
	"zk":    storage.InitZk,
	"k8s":   k8s.InitK8s,
	"redis": storage.InitRedis,
	"mysql": storage.InitMySQL,
	"es":    elastic.InitEs,
}

var cronTask = map[string]func() error{
	//"ian": cron.RegisterNoteIan, // 定期创建日报
	"logger": cron.RegLoggerCron,
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
	config.InitConf()
	if mode == "" {
		mode = config.Conf.RunMode
	}
	return nil
}

var initLogger = func() error {
	return nil
}

// step 4 初始化组件
var initComponents = func() error {
	// 如果pass ,初始化不执行组件检查
	//if pass := os.Getenv("PASS"); pass != "true" {
	//	return nil
	//}
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
		event.Bus.Publish("mode", "debug")
	}
	return nil
}

// step 7 初始化factory 初始数据
var initData = func() error {
	return nil
}

// step 8 初始化定时任务
var initCronTask = func() error {
	var taskRootCtx = context.Background()
	for _, t := range config.Conf.Cron {
		if err := cronTask[t](); err != nil {
			return err
		}
	}
	for _, t := range cron.QueueCron {
		go func(task pkg.Job) {
			_ = task.Exec(taskRootCtx, nil)
		}(t)
	}
	return nil
}

func PreRun(configPath string) string {
	if configPath != "" {
		utils.NoError(os.Setenv("configPath", configPath))
		fmt.Println("配置文件路径为:", configPath)
	}
	for _, f := range preCheck {
		utils.NoError(f())
	}
	event.Bus.Publish("run_mode", "debug")
	return mode
}

func init() {
	preCheck = append(preCheck, initEvent, envCheck,
		initConfig, initLogger, initComponents, initMode, initData, initCronTask)

}
