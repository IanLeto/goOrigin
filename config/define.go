package config

import (
	"fmt"
	"goOrigin/pkg/logger"
	"goOrigin/pkg/utils"
)

var Conf *Config
var ConfV2 *V2Config

var logger2, err = logger.InitZap()

func InitConf() {
	Conf = NewConfig()
	ConfV2 = NewV2Config()
	//logger2.Sugar().Infof("config init success", utils.ToJson(ConfV2))
	fmt.Println(utils.ToJson(ConfV2))
}
