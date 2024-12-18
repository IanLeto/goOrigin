package processor_test

import (
	"bufio"
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/processor"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// RedisSuite :
type PipelineTest struct {
	suite.Suite
	FilterNode processor.FilterTraceNode
	ctx        context.Context
	Pipe       processor.Pipeline
}

func (s *PipelineTest) SetupTest() {

}

func (s *PipelineTest) TestPing() {
	// 初始化 Pipeline
	s.Pipe = processor.Pipeline{}
	s.Pipe.Nodes = append(s.Pipe.Nodes, &processor.FilterTraceNode{})
	s.ctx = context.TODO()

	go func() {
		// 定义文件路径
		filePath := "test.json"

		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			// 打印调试信息
			fmt.Println("Failed to open file:", filePath)
			fmt.Println("Error:", err)

			// 获取文件的绝对路径
			absolutePath, _ := filepath.Abs(filePath)
			fmt.Println("Absolute path:", absolutePath)

			// 获取当前工作目录
			currentDir, _ := os.Getwd()
			fmt.Println("Current working directory:", currentDir)

			// 提示可能的解决方法
			fmt.Println("Make sure the file exists at the specified location, or check the working directory.")
			return
		}
		defer file.Close()

		// 创建一个读取器
		reader := bufio.NewScanner(file)
		var ch = make(chan []byte)
		defer func() { close(ch) }()

		for {
			select {
			default:
				go func() {
					for v := range s.Pipe.Start(s.ctx, ch) {
						fmt.Println(v)
					}
				}()
				for reader.Scan() {
					ch <- reader.Bytes()
					// 处理日志行

				}
			}
		}
	}()

	// 让 Goroutine 有足够时间运行
	time.Sleep(100 * time.Second)
}

// TestHttpClient :
func TestPipelineTestConfiguration(t *testing.T) {
	suite.Run(t, new(PipelineTest))
}
