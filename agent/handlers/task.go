package handlers

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	pbs "goOrigin/agent/pb"
	"goOrigin/agent/service"
	"goOrigin/pkg/utils"
	"os"
	"os/exec"
	"sync"
)

var ShPath = utils.GetFilePath("template/test.sh")
var TaskPool sync.Map

type TaskHandler struct {
	Content string
}

func (t *TaskHandler) GetFileStream(ping *pbs.Ping, server pbs.Agent_GetFileStreamServer) error {
	var (
		fileObj, fileErr = os.OpenFile("test", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		outWriter        = bufio.NewWriterSize(fileObj, 1)
		err              error
	)

	if fileErr != nil {
		return err
	}
	defer func() {
		_ = outWriter.Flush()
		_ = fileObj.Close()
	}()
	cmd := exec.Command("/bin/bash", utils.GetFilePath("template/test.sh"))
	stdout, _ := cmd.StdoutPipe()

	err = cmd.Start()
	if err != nil {
		logrus.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			err := server.Send(&pbs.FileStream{
				Id:    "",
				Items: []string{scanner.Text()},
			})
			if err != nil {
				logrus.Errorf("err occur in send stream %s", err)
				return
			}
		}
	}()
	err = cmd.Wait()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return nil
}

func (t *TaskHandler) MakeShell(ctx context.Context, request *pbs.MakeShellRequest) (*pbs.MakeShellResponse, error) {
	var (
		res     = &pbs.MakeShellResponse{}
		err     error
		content []byte
	)

	res.TaskID = request.TaskID
	content, err = service.MakeShell(ctx, request)
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
