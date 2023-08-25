package k8s

import (
	"context"
	"goOrigin/config"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

var K8SConn *KubeConn

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
	K8SConn = NewK8sConn(context.TODO(), nil)
	return nil
}

func NewK8sConn(ctx context.Context, conf *config.Config) *KubeConn {
	if conf == nil {
		conf = config.Conf
	}
	// 这里的flag 可以重置命令行
	//k8sconfig := flag.String("k8sconfig1", "/Users/ian/.kube/config", "kubernetes config file path")
	//flag.Parse()
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// 使用 kubeconfig 文件创建 config 对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	dyClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return &KubeConn{ClientSet: client, DynamicClient: dyClient}
}
