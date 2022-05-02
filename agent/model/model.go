package model

import (
	"context"
	"errors"
	"goOrigin/pkg/utils"
	"os/exec"
	"time"
)

type TaskPool struct {
	Tasks []Task
}

func (t *TaskPool) Reg(j Task)  {
	t.Tasks = append(t.Tasks, j)
}

type Task interface {
	Exec() error
}

type ShellTask struct {
	Content string
	Path string
	Timeout string

}

type SyncTask struct {
	Url     string
	Content string
	Timeout int
	Ctx     context.Context
}

type AsyncTask struct {
	Url     string
	Content string
	Timeout int
	Ctx     context.Context
}

type TaskResult struct {
	Err    error
	Result []byte
}

func (t *SyncTask) ExecSingle(ctx context.Context) (res []byte, err error) {
	var (
		ch = make(chan *TaskResult)
	)
	go func() {
		var r = &TaskResult{}
		r.Result, r.Err = exec.Command("/bin/bash", utils.GetFilePath("template/test.sh")).CombinedOutput()
		ch <- r
	}()
	select {
	case <-ctx.Done():
		return nil, errors.New("canceled")
	case <-time.After(time.Duration(t.Timeout) * time.Second):
		return nil, errors.New("timeout")
	case result, ok := <-ch:
		if ok {
			return result.Result, result.Err
		}
		return nil, errors.New("unknown error")
	}

}


