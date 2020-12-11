package define

import "github.com/spf13/viper"

var ConfigPath = "config"

var Viper = viper.New()

var InitHandler []func() error
