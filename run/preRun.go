package run

import (
	"goOrigin/config"
	"goOrigin/storage"
	"goOrigin/utils"
)

var defaultConfigPath = "/Users/ian/go/src/goOrigin/config.yaml"

var preCheck []func() error

// step 1 本地环境变量检查
var envCheck = func() error {
	return nil
}

// step 2 本地配置文件检查
var initConfig = func() error {
	config.InitConf(defaultConfigPath)
	return nil
}

// step 3 初始化组件
var initMySQL = func() {

}

var initMongoSession = func() error {
	storage.InitMongo()
	return nil
}

func PreRun() {
	for _, f := range preCheck {
		utils.CheckNoError(f())
	}
}

func init() {
	preCheck = append(preCheck, envCheck, initConfig, initMongoSession)
}
