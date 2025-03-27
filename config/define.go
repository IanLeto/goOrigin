package config

import (
	"fmt"
	"github.com/spf13/viper"
	"goOrigin/pkg/utils"
	"log"
	"os"
	"strings"
)

var ConfV2 *V2Config

func NewV2Config() *V2Config {
	var (
		path string
	)
	// 当然环境变量的优先级最高
	if os.Getenv("configPath") != "" {
		path = os.Getenv("configPath")
	}
	paths := strings.Split(path, ",")
	for _, p := range paths {
		viper.SetConfigFile(p)
		// 创建一个临时 Viper 实例来读取当前配置文件
		tempV := viper.New()
		tempV.SetConfigFile(p)
		if err := tempV.ReadInConfig(); err != nil {
			log.Fatalf("读取配置文件 %s 时出错: %v", p, err)
		}

		// 将当前配置文件的内容合并到主配置中
		configMap := tempV.AllSettings()
		if err := viper.MergeConfigMap(configMap); err != nil {
			log.Fatalf("合并配置文件 %s 时出错: %v", p, err)
		}
	}

	return &V2Config{
		Base:      NewBaseConfig(),
		Env:       NewComponentConfig(),
		Component: viper.GetStringSlice("component"),
		Cron:      viper.GetStringSlice("cron"),
	}
}

func InitConf() {
	ConfV2 = NewV2Config()
	fmt.Println(utils.ToJson(ConfV2))
}
