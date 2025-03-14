package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goOrigin/pkg/utils"
	"log"
	"os"
	"strings"
)

const k8sConfigMap = "/root/config/config.yaml"
const NodeMapping = "node"

var BaseInfo = map[string]string{}

type Config struct {
	Name       string `yaml:"name"`
	Port       string `yaml:"port"`
	RunMode    string `yaml:"run_mode"`
	SSH        *SSHConfig
	Backend    *BackendConfig
	Client     *HttpClientConfig
	Logging    *LoggingConfig
	Cron       []string
	Components []string `yaml:"components"`
}

func NewConfig() *Config {
	var (
		path string
	)
	// 我们现在三个环境, 本地, k8s, 腾讯云 之后会有工作站
	switch os.Getenv("mode") {
	case "k8s":
		path = k8sConfigMap
	case "local":
		dir, err := os.Getwd()
		utils.NoError(err)
		path = dir
	case "remote":
		path = ""
	default:
		dir, err := os.Getwd()
		utils.NoError(err)
		path = dir
	}
	// 当然环境变量的优先级最高
	if os.Getenv("configPath") != "" {
		path = os.Getenv("configPath")
	}
	paths := strings.Split(path, ",")
	for _, p := range paths {
		viper.SetConfigFile(p)
		utils.NoError(viper.MergeInConfig())
	}
	utils.NoError(viper.ReadInConfig())
	return &Config{
		Name:       viper.GetString("name"),
		Port:       viper.GetString("port"),
		RunMode:    viper.GetString("run_mode"),
		Components: viper.GetStringSlice("components"),
		SSH:        NewSSHConfig(),
		Logging:    NewLoggingConfig(),
		Backend:    NewBackendConfig(),
		Client:     NewHttpClientConfig(),
		Cron:       viper.GetStringSlice("cron"),
	}

}

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

func NewV2ConfigFromPath(p string) *V2Config {
	viper.SetConfigFile(p)
	// 创建一个临时 Viper 实例来读取当前配置文件
	tempV := viper.New()
	tempV.SetConfigFile(p)
	if err := tempV.ReadInConfig(); err != nil {
		logger2.Sugar().Errorf("读取配置文件 %s 时出错: %v", p, err)
	}

	// 将当前配置文件的内容合并到主配置中
	configMap := tempV.AllSettings()
	if err := viper.MergeConfigMap(configMap); err != nil {
		logger2.Sugar().Errorf("读取配置文件 %s 时出错: %v", p, err)
	}
	return &V2Config{
		Base:      NewBaseConfig(),
		Env:       NewComponentConfig(),
		Component: viper.GetStringSlice("component"),
		Cron:      viper.GetStringSlice("cron"),
	}
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config file changed")
	})
}
