package processor_test

import (
	"context"
	"fmt"
	"goOrigin/pkg/processor"
	"testing"
)

// 基准测试函数
func BenchmarkPipelineWithChannel(b *testing.B) {
	ctx := context.Background()
	p := &processor.Pipeline{}
	for i := 0; i < 10; i++ {
		p.Add(&processor.SimpleNode{})
	}
	var count = 0
	input := make(chan []byte, 1)
	input <- []byte("benchmark input")
	close(input)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output := p.Start(ctx, input)
		for range output {
			count++
		}
	}
	fmt.Println(count)
}

func BenchmarkPipelineWithoutChannel(b *testing.B) {
	ctx := context.Background()
	p := &processor.PipelineWithoutChannel{}
	for i := 0; i < 10; i++ {
		p.Add(&processor.SimpleNode{})
	}
	var count = 0
	input := []byte("benchmark input")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = p.Start(ctx, input)
		count++

	}
	fmt.Println(count)
}
