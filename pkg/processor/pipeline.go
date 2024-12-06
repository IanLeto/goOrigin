package processor

import (
	"context"
	"time"
)

// Node 接口定义了节点的处理方法
type Node interface {
	Process(ctx context.Context, input []byte) ([]byte, error)
	ProcessWithChannel(ctx context.Context, input <-chan []byte, output chan<- []byte) error
}

// SimpleNode 是一个简单的节点实现
type SimpleNode struct{}

func (n *SimpleNode) Process(ctx context.Context, input []byte) ([]byte, error) {
	// 模拟数据处理,这里简单地将输入数据原样返回
	time.Sleep(1 * time.Second)
	return input, nil
}

func (n *SimpleNode) ProcessWithChannel(ctx context.Context, input <-chan []byte, output chan<- []byte) error {
	for data := range input {
		time.Sleep(1 * time.Millisecond)
		output <- data
	}
	return nil
}

// Pipeline 是使用通道实现的流水线
type Pipeline struct {
	Nodes []Node
}

func (p *Pipeline) Add(nodes ...Node) {
	p.Nodes = append(p.Nodes, nodes...)
}

func (p *Pipeline) Start(ctx context.Context, input <-chan []byte) <-chan []byte {
	var output chan []byte
	for i, node := range p.Nodes {
		if i == 0 {
			output = make(chan []byte)
		} else {
			input = output
			output = make(chan []byte)
		}
		go func(n Node, in <-chan []byte, out chan<- []byte) {
			err := n.ProcessWithChannel(ctx, in, out)
			if err != nil {
				logger.Sugar().Errorln(err)
			}
			close(out)
		}(node, input, output)
	}
	return output
}

// PipelineWithoutChannel 是不使用通道实现的流水线
type PipelineWithoutChannel struct {
	Nodes []Node
}

func (p *PipelineWithoutChannel) Add(nodes ...Node) {
	p.Nodes = append(p.Nodes, nodes...)
}

func (p *PipelineWithoutChannel) Start(ctx context.Context, input []byte) ([]byte, error) {
	var err error
	for _, node := range p.Nodes {
		input, err = node.Process(ctx, input)
		if err != nil {
			return nil, err
		}
	}
	return input, nil
}
