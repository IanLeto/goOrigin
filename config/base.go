package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

const k8sConfigMap = "/root/config/config.yaml"
const tencentLocal = "/root/config/config.yaml"
const debug = "/Users/ian/go/src/goOrigin/config.yaml"

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

func NewConfig(path string) *Config {
	if initConfig(path) != nil {
		panic("init config failed")
	}
	return &Config{
		Name:       viper.GetString("name"),
		Port:       viper.GetString("addr"),
		RunMode:    viper.GetString("run_mode"),
		Components: viper.GetStringSlice("components"),
		Data:       viper.GetStringSlice("data"),
		Logging:    NewLoggingConfig(),
		Backend:    NewBackendConfig(),
		Client:     NewHttpClientConfig(),
		Cron:       viper.GetStringSlice("cron"),
		SSH:        NewSSHConfig(),
	}
}

func init() {
	//define.InitHandler = append(define.InitHandler, initConfig)
}

func initConfig(path string) error {
	var configPath = ""
	fmt.Println(os.Getenv("mode"))
	switch os.Getenv("mode") {
	case "k8s":
		configPath = k8sConfigMap
	case "debug":
		configPath = debug
	case "tencent":
		configPath = tencentLocal
	default:
		configPath = debug

	}

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error in init config %s", err)
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("reading config err %v", err)
		}
		viper.AddConfigPath(dir)
		if err = viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config file changed")
	})
}
