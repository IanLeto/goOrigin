package cmdHandlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pbs "goOrigin/agent/pb"
	"goOrigin/backend"
	"goOrigin/internal/params"
	"io"
)

func ssh(c *gin.Context) {

}

type JobGroup struct {
}

type Job struct {
	JobID   string `json:"job_id"`
	Content string `json:"content"`
	Client  pbs.AgentClient
}

func Ping(c *gin.Context) {
	var (
		err error
		ctx = context.Background()
		res *pbs.Pong
	)
	cli, err := backend.NewAgentClient()
	if err != nil {
		goto ERR
	}
	res, err = cli.PingTask(ctx, &pbs.Ping{})
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func StartTask(c *gin.Context) {
	var (
		err error
		ctx = context.Background()
		res *pbs.StartSyncTaskResponse
	)
	cli, err := backend.NewAgentClient()
	if err != nil {
		goto ERR
	}
	res, err = cli.StartSyncTask(ctx, &pbs.StartSyncTaskRequest{
		TaskID:       "",
		ShellContent: "",
		Sync:         "",
		Params:       nil,
	})
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func MakeShell(c *gin.Context) {
	var (
		err error
		ctx = context.Background()
		res *pbs.MakeShellResponse
		req = &pbs.MakeShellRequest{}
	)
	cli, err := backend.NewAgentClient()
	if err != nil {
		goto ERR
	}
	if err = c.ShouldBindJSON(req); err != nil {
		goto ERR
	}

	res, err = cli.MakeShell(ctx, &pbs.MakeShellRequest{
		TaskID:       "1",
		ShellContent: "",
		Timeout:      req.Timeout,
	})
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func MakeStream(c *gin.Context) {
	var (
		err error
		ctx = context.Background()
		res *pbs.MakeShellResponse
		req = &pbs.MakeShellRequest{}
	)
	cli, err := backend.NewAgentClient()
	searchS, _ := cli.GetFileStream(ctx, &pbs.Ping{})
	for {
		r, err := searchS.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(r)
		if err != nil {
			goto ERR
		}

	}
	if err != nil {
		goto ERR
	}
	if err = c.ShouldBindJSON(req); err != nil {
		goto ERR
	}

	res, err = cli.MakeShell(ctx, &pbs.MakeShellRequest{
		TaskID:       "1",
		ShellContent: "",
		Timeout:      req.Timeout,
	})

	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
