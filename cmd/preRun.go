package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"goOrigin/cmd/event"
	"goOrigin/config"
	"goOrigin/internal/db"
	"goOrigin/pkg"
	"goOrigin/pkg/cron"
	"goOrigin/pkg/k8s"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
	"os"
)

//var defaultConfigPath = "/Users/ian/go/src/goOrigin/config.yaml"
var defaultConfigPath = ""
var preCheck []func() error
var mode string
var connFactory = make([]storage.Conn, 0)

var migrate = map[string]interface{}{
	"t_jobs": &db.TJob{},
}

// 初始化组件
var compInit = map[string]func() error{
	"mongo": storage.InitMongo,
	"zk":    storage.InitZk,
	"k8s":   k8s.InitK8s,
	"redis": storage.InitRedis,
	"mysql": storage.InitMySQL,
}

var cronTask = map[string]func() error{
	"ian": cron.RegisterNoteIan, // 定期创建日报
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
	//for _, d := range config.Conf.Data {
	//	switch d {
	//	case "mongo":
	//		connFactory = append(connFactory, storage.GloablMongo)
	//	}
	//}
	//// factory 行为执行
	//for _, conn := range connFactory {
	//	if err := conn.InitData(mode); err != nil {
	//		return err
	//	}
	//}
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
var dbMigrate = func() error {
	for name, i := range migrate {
		if err := storage.GlobalMySQL.AutoMigrate(i); err != nil {
			logrus.Errorf("初始化 table %s failed: %s", name, err.Error)
		}
	}
	return nil
}

func PreRun(configPath string) string {
	if configPath != "" {
		defaultConfigPath = configPath
	}
	for _, f := range preCheck {
		utils.NoError(f())
	}
	event.Bus.Publish("run_mode", "debug")
	return mode
}

func init() {
	preCheck = append(preCheck, initEvent, envCheck,
		initConfig, initLogger, initComponents, initMode, initData, initCronTask, dbMigrate)

}
