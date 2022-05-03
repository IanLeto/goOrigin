package handlers

import (
	"context"
	"goOrigin/agent/service"
	"goOrigin/pkg/utils"
	pbs "goOrigin/agent/pb"
	"sync"
)

var ShPath = utils.GetFilePath("template/test.sh")
var TaskPool sync.Map

type TaskHandler struct {
	Content string

}

func (t *TaskHandler) MakeShell(ctx context.Context, request *pbs.MakeShellRequest) ( *pbs.MakeShellResponse,  error) {
	var (

		res = &pbs.MakeShellResponse{}
		err error
		content []byte

	)

	res.TaskID = request.TaskID
	content ,err = service.MakeShell(ctx,request)
	if err != nil {
		goto ERR
	}
	res.Content = string(content)
	ERR:
		return res, err
}

func (t *TaskHandler) PingTask(ctx context.Context, ping *pbs.Ping) (*pbs.Pong, error) {
	return &pbs.Pong{Version: "v0.0.1"}, nil
}

func (t *TaskHandler) StartTask(ctx context.Context, request *pbs.StartTaskRequest) (*pbs.StartTaskResponse, error) {
	var (
		err      error
		response = &pbs.StartTaskResponse{}
	)
	return response, err
}

func (t *TaskHandler) StartSyncTask(ctx context.Context, request *pbs.StartSyncTaskRequest) (*pbs.StartSyncTaskResponse, error) {
	panic(1)
	//return service.StartSyncTask(ctx, request)

}

func (t *TaskHandler) StopTask(ctx context.Context, request *pbs.StopTaskRequest) (*pbs.StopTaskResponse, error) {
	panic("implement me")
}

func (t *TaskHandler) GetTaskDetailOne(ctx context.Context, request *pbs.GetTaskDetailRequest) (*pbs.GetTaskDetailResponse, error) {
	panic("implement me")
}
