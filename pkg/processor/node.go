package processor

import "context"

type Filter struct {
}

func (f *Filter) Process(ctx context.Context, input []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (f *Filter) ProcessWithChannel(ctx context.Context, input <-chan []byte, output chan<- []byte) error {
	//TODO implement me
	panic("implement me")
}
