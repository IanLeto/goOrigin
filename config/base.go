package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goOrigin/define"
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

func (c *Config) ReceiveConfigPath(path string) {


}
func NewConfig() *Config {
	if InitConfig() != nil {
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
	define.InitHandler = append(define.InitHandler, InitConfig)
}
func InitConfig() error {

	viper.SetConfigFile("../config/config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
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
