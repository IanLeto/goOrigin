package service

import (
	"context"
	"errors"
	"goOrigin/agent/pbs/service"
	"goOrigin/pkg/utils"
	"os/exec"
	"time"
)

func StartSyncTask(ctx context.Context, request *service.StartSyncTaskRequest) (*service.StartSyncTaskResponse, error) {
	//var task = model.ShellTask{}
	//res ,err := task.Exec()
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func MakeShell(ctx context.Context, request *service.MakeShellRequest) (res []byte, err error) {
	type taskResult struct {
		Err    error
		Result []byte
	}
	var (
		ch = make(chan *taskResult)
		c, cancel = context.WithCancel(ctx)
	)
	defer cancel()
	go func(c context.Context) {
		var r = &taskResult{}
		select {
		case <-ctx.Done():
			return
		default:
			r.Result, r.Err = exec.Command("/bin/bash", utils.GetFilePath("template/test.sh")).CombinedOutput()
			ch <- r
		}

	}(c)

	select {
	case <-ctx.Done():
		return nil, errors.New("canceled")
	case <-time.After(time.Duration(request.Timeout) * time.Second):
		return nil, errors.New("timeout")
	case result, ok := <-ch:
		if ok {
			return result.Result, result.Err
		}
		return nil, errors.New("unknown error")
	}

}
