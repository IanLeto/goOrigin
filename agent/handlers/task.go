package handlers

import (
	"context"
	"goOrigin/agent/pbs/service"
	"goOrigin/pkg/utils"
	"sync"
)

var ShPath = utils.GetFilePath("template/test.sh")
var TaskPool sync.Map

type TaskHandler struct {
	Content string
}

func (t *TaskHandler) PingTask(ctx context.Context, ping *service.Ping) (*service.Pong, error) {
	return &service.Pong{Version: "v0.0.1"}, nil
}

func (t *TaskHandler) StartTask(ctx context.Context, request *service.StartTaskRequest) (*service.StartTaskResponse, error) {
	var (
		err      error
		response = &service.StartTaskResponse{}
	)
	return response, err
}

func (t *TaskHandler) StartSyncTask(ctx context.Context, request *service.StartSyncTaskRequest) (*service.StartSyncTaskResponse, error) {
	return service.StartSyncTask(ctx, request)

}

func (t *TaskHandler) StopTask(ctx context.Context, request *service.StopTaskRequest) (*service.StopTaskResponse, error) {
	panic("implement me")
}

func (t *TaskHandler) GetTaskDetailOne(ctx context.Context, request *service.GetTaskDetailRequest) (*service.GetTaskDetailResponse, error) {
	panic("implement me")
}
