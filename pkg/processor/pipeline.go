package processor

import (
	"context"
)

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
