package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goOrigin/pkg/utils"
	"log"
	"os"
	"path/filepath"
)

const k8sConfigMap = "/root/config/config.yaml"
const tencentLocal = "/root/config/config.yaml"
const debug = "/Users/ian/go/src/goOrigin/config.yaml"
const NodeMapping = "node"
const configPath = "/app/config.yaml"

type Config struct {
	Name       string `yaml:"name"`
	Port       string `yaml:"port"`
	RunMode    string `yaml:"run_mode"`
	Data       []string
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
	viper.SetConfigFile(filepath.Join(path, "config.yaml"))
	utils.NoError(viper.ReadInConfig())
	return &Config{
		Name:       viper.GetString("name"),
		Port:       viper.GetString("addr"),
		RunMode:    viper.GetString("run_mode"),
		Components: viper.GetStringSlice("components"),
		Data:       viper.GetStringSlice("data"),
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
