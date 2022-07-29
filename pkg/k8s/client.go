package k8s

import (
	"context"
	"flag"
	"goOrigin/config"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"log"
)

var K8S *KubeConn

type KubeConn struct {
	Client        *kubernetes.Clientset
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
	K8S = NewK8sConn(context.TODO(), nil)
	return nil
}

func NewK8sConn(ctx context.Context, conf *config.Config) *KubeConn {
	if conf == nil {
		conf = config.Conf
	}
	k8sconfig := flag.String("k8sconfig", "/Users/ian/.kube/config", "kubernetes config file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
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

	//fmt.Println(client.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{}))
	return &KubeConn{Client: client, DynamicClient: dyClient}
}
