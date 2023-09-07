package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Checker interface {
	Check()
}

type ComponentConfig interface {
	NewComponent() ComponentConfig
}

type BackendConfig struct {
	*MySqlBackendConfig
	*MongoBackendConfig
	*ZKConfig
	*RedisConfig
	*K8sConfig
	*PromConfig
	*EsConfig
	*JaegerConfig
	*MysqlConfig
}

func NewBackendConfig() *BackendConfig {
	return &BackendConfig{
		NewMySqlBackendConfig(),
		NewMongoBackendConfig(),
		NewZkConfig(),
		NewRedisConfig(),
		NewK8sConfig(),
		NewPromConfig(),
		NewEsConfig(),
		NewJaegerConfig(),
		NewMysqlConfig(),
	}
}

// base Http client conf

type HttpClientConfig struct {
	CC *CCConf
}

func NewHttpClientConfig() *HttpClientConfig {
	return &HttpClientConfig{
		CC: NewCCClientConf(),
	}
}

// MySqlBackendConfig mysql backend config
type MySqlBackendConfig struct {
	Address  string
	Port     string
	Password string
	User     string
	Name     string
}

func NewMySqlBackendConfig() *MySqlBackendConfig {
	return &MySqlBackendConfig{
		Address:  viper.GetString("backend.MySql.address"),
		Port:     viper.GetString("backend.MySql.port"),
		User:     viper.GetString("backend.MySql.user"),
		Password: viper.GetString("backend.MySql.password"),
		Name:     viper.GetString("backend.MySql.name"),
	}
}

// MongoBackendConfig mongoDB
type MongoBackendConfig struct {
	Address  string
	Port     string
	Password string
	User     string
	DB       string
}

func NewMongoBackendConfig() *MongoBackendConfig {
	return &MongoBackendConfig{
		Address:  viper.GetString("backend.mongo.address"),
		Port:     viper.GetString("backend.mongo.port"),
		User:     viper.GetString("backend.mongo.user"),
		Password: viper.GetString("backend.mongo.password"),
		DB:       viper.GetString("backend.mongo.DB"),
	}
}

// zookeeper 配置
type ZKConfig struct {
	Address []string
	Master  string
	Auth    string
}

func NewZkConfig() *ZKConfig {
	return &ZKConfig{
		Address: viper.GetStringSlice("backend.zk.Address"),
		Master:  viper.GetString("backend.zk.Master"),
		Auth:    viper.GetString("backend.zk.Auth"),
	}
}

type RedisConfig struct {
	DB         int
	Addr       string
	IsSentinel bool
	Auth       string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		DB:         viper.GetInt("backend.redis.DB"),
		Addr:       viper.GetString("backend.redis.Addr"),
		IsSentinel: viper.GetBool("backend.redis.IsSentinel"),
		Auth:       viper.GetString("backend.redis.Auth"),
	}
}

type PromConfig struct {
	Address string
	Push    string
	Group   string
}

func NewPromConfig() *PromConfig {
	return &PromConfig{
		Push:    viper.GetString("backend.prom.push"),
		Address: viper.GetString("backend.prom.address"),
		Group:   viper.GetString("backend.prom.group"),
	}
}

// es

type EsConfig struct {
	ElasticSearchRegion map[string]EsInfo `json:"elasticsearch"`
}
type EsInfo struct {
	Address string `json:"address"`
	Region  string `json:"region"`
}

func NewEsConfig() *EsConfig {
	esRegion := viper.GetStringMap("backend.es")
	esRegions := make(map[string]EsInfo)
	for s, info := range esRegion {
		regionInfo := info.(map[string]interface{})
		esRegions[s] = EsInfo{
			Address: regionInfo["address"].(string),
			Region:  s,
		}
	}
	return &EsConfig{
		ElasticSearchRegion: esRegions,
	}
}

type JaegerConfig struct {
	Address string
}

func NewJaegerConfig() *JaegerConfig {
	return &JaegerConfig{Address: viper.GetString("backend.jaeger.address")}
}

// 配置中心httpclient 配置参数

type CCConf struct {
	Address   string
	HeartBeat int
}

func NewCCClientConf() *CCConf {
	return &CCConf{
		Address:   viper.GetString("client.CC.address"),
		HeartBeat: viper.GetInt("client.CC.heart_beat"),
	}

}

// LoggingConfig logging 配置
type LoggingConfig struct {
	FileName string
	Level    string
	Path     string
	Rotation RotationConfig
}

type RotationConfig struct {
	Time  int
	Count int
}

func NewLoggingConfig() *LoggingConfig {
	return &LoggingConfig{
		FileName: viper.GetString("logging.fileName"),
		Level:    viper.GetString("logging.level"),
		Path:     viper.GetString("logging.path"),
		Rotation: RotationConfig{
			Time:  viper.GetInt("logging.rotation.time"),
			Count: viper.GetInt("logging.rotation.Count"),
		},
	}
}

// SSHConfig ssh
type SSHConfig struct {
	Address string
	User    string
	Type    string
	KeyPath string // ssh_id 路径
	Port    int
	Auth    string
}

func NewSSHConfig() *SSHConfig {
	return &SSHConfig{
		Address: viper.GetString("ssh.address"),
		User:    viper.GetString("ssh.user"),
		Type:    viper.GetString("ssh.type"),
		KeyPath: viper.GetString("ssh.key_path"),
		Auth:    viper.GetString("ssh.auth"),
		Port:    viper.GetInt("ssh.port"),
	}
}

// K8sConfig k8s 配置
type K8sConfig struct {
	Clusters map[string]ClusterInfo
}

type ClusterInfo struct {
	ClusterName      string `json:"cluster_name"`
	APIServerAddress string `json:"apiserver_address"`
	ConfigAddress    string `json:"config_address"`
	IsInCluster      bool   `json:"is_in_cluster"`
	Token            string `json:"token"`
}

func NewK8sConfig() *K8sConfig {
	cluster := viper.GetStringMap("backend.cluster")
	clusters := make(map[string]ClusterInfo)
	for s, info := range cluster {
		clusterInfo := info.(map[string]interface{})
		clusters[s] = ClusterInfo{
			ClusterName:      clusterInfo["cluster"].(string),
			APIServerAddress: clusterInfo["apiserver_address"].(string),
			ConfigAddress:    clusterInfo["config_address"].(string),
			IsInCluster:      clusterInfo["is_in_cluster"].(bool),
			Token:            clusterInfo["token"].(string),
		}
	}

	return &K8sConfig{
		Clusters: clusters,
	}
}

type MysqlConfig struct {
	Clusters map[string]*MysqlInfo `json:"clusters"`
}

func (m MysqlConfig) Check() (string, error) {

	for cluster, info := range m.Clusters {
		if !info.IsMigration {
			continue
		}
		if info.Address == "" {
			return cluster, errors.New("mysql address is empty")
		}
		if info.Port == "" {
			return cluster, errors.New("mysql port is empty")
		}
		if info.User == "" {
			return cluster, errors.New("mysql user is empty")
		}
		if info.Password == "" {
			return cluster, errors.New("mysql password is empty")
		}
		if info.Name == "" {
			return cluster, errors.New("mysql name is empty")
		}
	}
	return "", nil
}

type MysqlInfo struct {
	Address     string `json:"address"`
	Port        string `json:"port"`
	Password    string `json:"password"`
	User        string `json:"user"`
	Name        string `json:"name"`
	IsMigration bool   `json:"isMigration"`
}

func NewMysqlConfig() *MysqlConfig {
	cluster := viper.GetStringMap("backend.MySql")
	clusters := make(map[string]*MysqlInfo)
	for s, info := range cluster {
		clusterInfo := info.(map[string]interface{})
		clusters[s] = &MysqlInfo{
			Address:     clusterInfo["address"].(string),
			User:        clusterInfo["user"].(string),
			Password:    clusterInfo["password"].(string),
			Name:        clusterInfo["name"].(string),
			IsMigration: clusterInfo["ismigration"].(bool),
		}
	}

	return &MysqlConfig{
		Clusters: clusters,
	}
}

func (k K8sConfig) Check() (string, error) {
	for cluster, info := range k.Clusters {
		if info.IsInCluster {
			continue
		}
		if info.ConfigAddress != "" {
			config := flag.String("kubeconfig", info.ConfigAddress, "absolute path to the kubeconfig file")
			flag.Parse()
			restConfig, err := clientcmd.BuildConfigFromFlags("", *config)
			if err != nil {
				return cluster, err
			}
			client, err := kubernetes.NewForConfig(restConfig)
			version := client.RESTClient().APIVersion().Version
			if err != nil {
				return cluster, err
			}
			fmt.Println(version)
			continue
		}
		if info.Token == "" {
			return cluster, errors.New("token is empty")
		}

		if info.APIServerAddress == "" {
			return cluster, errors.New("api server address is empty")
		}

	}
	return "", nil
}
