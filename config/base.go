package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goOrigin/pkg/utils"
	"log"
	"os"
)

const k8sConfigMap = "/root/config/config.yaml"
const NodeMapping = "node"

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
	//viper.SetConfigFile(filepath.Join(path, "config.yaml"))
	viper.SetConfigFile(path)
	utils.NoError(viper.ReadInConfig())
	viper.OnConfigChange(func(in fsnotify.Event) {
		Conf = NewConfig()
	})
	return &Config{
		Name:       viper.GetString("name"),
		Port:       viper.GetString("addr"),
		RunMode:    viper.GetString("run_mode"),
		Components: viper.GetStringSlice("components"),
		SSH:        NewSSHConfig(),
		Logging:    NewLoggingConfig(),
		Backend:    NewBackendConfig(),
		Client:     NewHttpClientConfig(),
		Cron:       viper.GetStringSlice("cron"),
	}
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config file changed")
	})
}
