package service

import (
	"context"
	"goOrigin/agent/model"
	"goOrigin/agent/pbs/service"
)

func StartSyncTask(ctx context.Context, request *service.StartSyncTaskRequest) (*service.StartSyncTaskResponse, error) {
	var task = model.ShellTask{}
	res ,err := task.Exec()
	if err != nil {
		return nil, err
	}
	return res, nil
}
