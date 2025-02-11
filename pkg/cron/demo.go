package cron

import (
	"context"
	"fmt"
	"time"
)

// Demo 是一个获取 Pod 信息的任务
type Demo struct {
	name string
}

// Exec 实现 Job 接口中的 Run 方法
func (p *Demo) Exec(ctx context.Context) error {
	for {
		fmt.Println("Attempting to fetch Pod information...")
		// 模拟获取 Pod 信息
		podInfo, err := getPodInfo()
		if err != nil {
			fmt.Printf("Failed to fetch Pod information: %s\n", err)
		} else {
			fmt.Printf("Successfully fetched Pod information: %s\n", podInfo)
		}

		// 每隔 10 秒获取一次 Pod 信息
		time.Sleep(10 * time.Second)
	}
	return nil
}

// GetName Title 实现 Job 接口中的 Title 方法
func (p *Demo) GetName() string {
	return p.name
}

// 模拟获取 Pod 信息的函数
// 这里可以替换成真实的 Kubernetes API 调用
func getPodInfo() (string, error) {
	// 模拟一个成功和失败的情况
	if time.Now().Unix()%2 == 0 {
		return "Pod-1234", nil
	}
	// 模拟获取失败的情况
	return "", fmt.Errorf("unable to fetch Pod information")
}

func DemoCronFactory() error {
	GTM.AddJob(&Demo{name: "ian"})
	return nil
}
