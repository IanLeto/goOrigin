package utils

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

var SelfPodInfo *PodInfo

type PodInfo struct {
	PodName         string            // Pod名称
	PodNamespace    string            // Pod所在的命名空间
	PodIP           string            // Pod的IP地址
	PodUID          string            // Pod的唯一标识符
	NodeName        string            // Pod所在的节点名称
	ContainerName   string            // 容器名称
	ContainerImage  string            // 容器镜像
	ContainerID     string            // 容器ID
	Labels          map[string]string // Pod的标签
	Annotations     map[string]string // Pod的注解
	EnvironmentVars map[string]string // 指定的环境变量
}

func GetPodInfo(envVars []string) (*PodInfo, error) {
	// 创建 Kubernetes 客户端配置
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	// 创建 Kubernetes 客户端实例
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	// 获取当前 Pod 的名称
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		return nil, fmt.Errorf("POD_NAME environment variable not set")
	}

	// 获取当前 Pod 的命名空间
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		return nil, fmt.Errorf("POD_NAMESPACE environment variable not set")
	}

	// 获取当前 Pod 的信息
	pod, err := clientset.CoreV1().Pods(podNamespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 创建PodInfo实例
	podInfo := &PodInfo{
		PodName:         podName,
		PodNamespace:    podNamespace,
		PodIP:           pod.Status.PodIP,
		PodUID:          string(pod.UID),
		NodeName:        pod.Spec.NodeName,
		ContainerName:   pod.Spec.Containers[0].Name,
		ContainerImage:  pod.Spec.Containers[0].Image,
		ContainerID:     pod.Status.ContainerStatuses[0].ContainerID,
		Labels:          pod.Labels,
		Annotations:     pod.Annotations,
		EnvironmentVars: make(map[string]string),
	}

	// 获取指定的环境变量
	for _, envVar := range envVars {
		podInfo.EnvironmentVars[envVar] = os.Getenv(envVar)
	}

	return podInfo, nil
}
