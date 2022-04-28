package task

import (
	"context"
	"goOrigin/agent/model"
	"goOrigin/agent/pbs/service"
	"goOrigin/pkg/utils"
	"sync"
)

var ShPath = utils.GetFilePath("template/test.sh")
var TaskPool sync.Map

type Task struct {
	Content string
}

func (t *Task) PingTask(ctx context.Context, ping *service.Ping) (*service.Pong, error) {
	return &service.Pong{Version: "v0.0.1"}, nil
}

func (t *Task) StartTask(ctx context.Context, request *service.StartTaskRequest) (*service.StartTaskResponse, error) {
	var (
		err      error
		response = &service.StartTaskResponse{}
	)
	TaskPool.Store(request.TaskID, )
	return response, err
}

func (t *Task) StartSyncTask(ctx context.Context, request *service.StartSyncTaskRequest) (*service.StartSyncTaskResponse, error) {
	var (
		err      error
		response = &service.StartSyncTaskResponse{}
		sync     = &model.SyncTask{
			Url:     "",
			Content: "",
			Timeout: 20,
			Ctx:     ctx,
		}
	)
	switch request.Sync {
	case "sync":

	default:

	}
	res, err := sync.ExecSingle(ctx)
	response.Result = string(res)
	return response, err

}

func (t *Task) StopTask(ctx context.Context, request *service.StopTaskRequest) (*service.StopTaskResponse, error) {
	panic("implement me")
}

func (t *Task) GetTaskDetailOne(ctx context.Context, request *service.GetTaskDetailRequest) (*service.GetTaskDetailResponse, error) {
	panic("implement me")
}
