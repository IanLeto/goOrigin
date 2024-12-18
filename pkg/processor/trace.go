package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"goOrigin/internal/model/entity"
)

var (
	UnDoneSpanInfoEntity map[string][]*entity.KafkaLogEntity
)

type Agg struct {
}

func (a *Agg) Process(ctx context.Context, input []byte) ([]byte, error) {
	var (
		spanInfo *entity.KafkaLogEntity
	)
	fmt.Println(input)
	_ = json.Unmarshal(input, spanInfo)
	switch spanInfo.SpanID {
	// 结束span
	case "0":
		return input, nil
	default:
		UnDoneSpanInfoEntity[spanInfo.TraceId] = append(UnDoneSpanInfoEntity[spanInfo.TraceId], spanInfo)
	}
	return nil, nil
}

func (a *Agg) ProcessWithChannel(ctx context.Context, input <-chan []byte, output chan<- []byte) error {
	//TODO implement me
	panic("implement me")
}
