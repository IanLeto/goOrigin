package conf

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
}

func NewConfig() *Config {
	InitConfig()
	return &Config{
		Name:    viper.GetString("name"),
		Port:    viper.GetString("addr"),
		Backend: NewBackendConfig(),
	}
}

type BackendConfig struct {
	*MySqlBackendConfig
}

func NewBackendConfig() *BackendConfig {
	return &BackendConfig{
		NewMySqlBackendConfig(),
	}
}

type MySqlBackendConfig struct {
	Address  string
	Port     string
	Password string
	User     string
}

func NewMySqlBackendConfig() *MySqlBackendConfig {
	return &MySqlBackendConfig{
		Address:  viper.GetString("backend.MySql.address"),
		Port:     viper.GetString("backend.MySql.port"),
		User:     viper.GetString("backend.MySql.user"),
		Password: viper.GetString("backend.MySql.password"),
	}
}

type BaseBackend struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (b *BaseBackend) GetType() string {
	panic("implement me")
}

func (b *BaseBackend) GetUrl() string {
	panic("implement me")
}

func (b *BaseBackend) GetAddress() string {
	panic("implement me")
}

func (b *BaseBackend) GetPort() string {
	panic("implement me")
}

func (b *BaseBackend) Close() error {
	panic("implement me")
}

type MySqlBackend struct {
	*BaseBackend
}

func init() {
	define.InitHandler = append(define.InitHandler, InitConfig)
	for _, v := range define.InitHandler {
		if err := v(); err != nil {
			panic("init config file %v")
		}
	}
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
