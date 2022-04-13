package task

import (
	"context"
	"goOrigin/agent/pbs/service"
)

type Task struct {
}

func (t *Task) PingTask(ctx context.Context, ping *service.Ping) (*service.Pong, error) {
	return &service.Pong{Version: "v0.0.1"}, nil
}

func (t *Task) StartTask(ctx context.Context, request *service.StartTaskRequest) (*service.StartTaskResponse, error) {
	panic("implement me")
}

func (t *Task) StopTask(ctx context.Context, request *service.StopTaskRequest) (*service.StopTaskResponse, error) {
	panic("implement me")
}

func (t *Task) GetTaskDetailOne(ctx context.Context, request *service.GetTaskDetailRequest) (*service.GetTaskDetailResponse, error) {
	panic("implement me")
}
