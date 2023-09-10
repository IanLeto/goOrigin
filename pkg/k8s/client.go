package k8s

import (
	"context"
	"goOrigin/config"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

var Conn *KubeConn

type KubeConn struct {
	ClientSet     kubernetes.Interface
	DynamicClient dynamic.Interface
}

func (k *KubeConn) Close() error {
	//TODO implement me
	panic("implement me")
}

func (k *KubeConn) InitData(mode string) error {
	//TODO implement me
	panic("implement me")
}

func InitK8s() error {
	Conn = NewK8sConn(context.TODO(), nil)
	return nil
}

func NewK8sConn(ctx context.Context, conf *config.Config) *KubeConn {
	if conf == nil {
		conf = config.Conf
	}
	subconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// 使用 subconfig 文件创建 configFromFlags 对象
	configFromFlags, err := clientcmd.BuildConfigFromFlags("", subconfig)
	if err != nil {
		configFromFlags, err = rest.InClusterConfig()
	}
	client, err := kubernetes.NewForConfig(configFromFlags)
	if err != nil {
		log.Fatal(err)
	}
	dyClient, err := dynamic.NewForConfig(configFromFlags)
	if err != nil {
		log.Fatal(err)
	}
	return &KubeConn{ClientSet: client, DynamicClient: dyClient}
}
