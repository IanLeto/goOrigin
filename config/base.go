package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goOrigin/logging"
	"log"
	"os"
)

type Config struct {
	Name    string `yaml:"name"`
	Port    string `yaml:"port"`
	RunMode string `yaml:"run_mode"`
	Backend *BackendConfig
	Client  *HttpClientConfig
}

func NewConfig(path string) *Config {
	if initConfig(path) != nil {
		panic("init config failed")
	}
	return &Config{
		Name:    viper.GetString("name"),
		Port:    viper.GetString("addr"),
		RunMode: viper.GetString("run_mode"),
		Backend: NewBackendConfig(),
		Client:  NewHttpClientConfig(),
	}
}

func init() {
	//define.InitHandler = append(define.InitHandler, initConfig)
}
func initConfig(path string) error {
	var logger = logging.GetStdLogger()
	if path == "" {
		// 这里的配置文件一定要放到项目根目录上
		// viper 读取文件的特性导致被不同包调用时，该路径会根据调用方变化
		viper.AddConfigPath("../")
		viper.SetConfigName("config") // 配置文件名称(无扩展名)
		viper.SetConfigType("yaml")
	} else {
		viper.SetConfigFile(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.Debugf("error in init config %s", err)
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
