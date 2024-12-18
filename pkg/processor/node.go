package processor

import (
	"context"
	"encoding/json"
	"goOrigin/internal/model/entity"
)

type FilterTraceNode struct {
}

func (f *FilterTraceNode) Process(ctx context.Context, input []byte) ([]byte, error) {
	var err error
	var spanInfo entity.KafkaLogEntity
	err = json.Unmarshal(input, &spanInfo)
	if spanInfo.Trans.TransTypeCode == "" {
		return nil, nil
	}
	return input, err
}

func (f *FilterTraceNode) ProcessWithChannel(ctx context.Context, input <-chan []byte, output chan<- []byte) error {
	//TODO implement me
	panic("implement me")
}
