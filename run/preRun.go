package run

import (
	"goOrigin/config"
	"goOrigin/utils"
)

var defaultConfigPath = "/Users/ian/go/src/goOrigin/config.yaml"

var preCheck []func() error

var Conf *config.Config

// step 1 本地环境变量检查
var envCheck = func() error {
	return nil
}

// step 2 本地配置文件检查
var initConfig = func() error {
	Conf = config.NewConfig(defaultConfigPath)
	return nil
}

// step 3 初始化组件
var initMySQL = func() {

}

func PreRun() {
	for _, f := range preCheck {
		utils.CheckNoError(f())
	}
}

func init() {
	preCheck = append(preCheck, envCheck, initConfig)
}
