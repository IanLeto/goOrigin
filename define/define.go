package define

import (
	"github.com/spf13/viper"
)

var ConfigPath = "config"

var Viper = viper.New()

var InitHandler []func() error



// 全部放到一个define 里面 大概率会出现循环引用
// define 之后仅放置常量
// var StdLogger *logging.Logger
