package logic

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func GetClusterIPPool(ctx context.Context, proxyURL string) ([]string, error) {
	// 通过集群代理器URL初始化配置
	config := &rest.Config{
		Host: proxyURL,
		// 如果需要认证，可以添加以下字段
		// BearerToken: "your-token",
		// 或者使用证书认证
		// TLSClientConfig: rest.TLSClientConfig{
		//     CAFile:   "/path/to/ca.crt",
		//     CertFile: "/path/to/client.crt",
		//     KeyFile:  "/path/to/client.key",
		// },
	}

	// 创建动态客户端
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %v", err)
	}

	// 定义 IPOOL 资源的 GVR (Group, Version, Resource)
	// 根据实际的 CRD 定义调整这些值
	ipoolGVR := schema.GroupVersionResource{
		Group:    "networking.example.com", // 替换为实际的 API Group
		Version:  "v1",                     // 替换为实际的版本
		Resource: "ipools",                 // 资源名称（通常是复数形式）
	}

	// 列出所有命名空间中的 IPOOL 资源
	ipoolList, err := dynamicClient.Resource(ipoolGVR).
		Namespace(""). // 空字符串表示所有命名空间
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list IPOOL resources: %v", err)
	}

	// 提取 IPOOL 名称列表
	var ipoolNames []string
	for _, item := range ipoolList.Items {
		name := item.GetName()
		namespace := item.GetNamespace()

		// 如果是命名空间级别的资源，可以包含命名空间信息
		if namespace != "" {
			ipoolNames = append(ipoolNames, fmt.Sprintf("%s/%s", namespace, name))
		} else {
			// 如果是集群级别的资源
			ipoolNames = append(ipoolNames, name)
		}
	}

	return ipoolNames, nil
}

//func QueryHistroy(ctx context.Context, req interface{}) {
//	re, err := svc.Metric().Query(ctx, req)
//}

type QueryResult struct {
	ID string
}
