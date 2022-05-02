package cmdHandlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "goOrigin/agent/pbs/service"
	"goOrigin/backend"
	"goOrigin/internal/params"
)

func ssh(c *gin.Context) {

}

type JobGroup struct {
}

type Job struct {
	JobID   string `json:"job_id"`
	Content string `json:"content"`
	Client  pb.AgentClient
}

func Ping(c *gin.Context) {
	var (
		err error
		ctx = context.Background()
		res *pb.Pong
	)
	cli, err := backend.NewAgentClient()
	if err != nil {
		goto ERR
	}
	res, err = cli.PingTask(ctx, &pb.Ping{})
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
		res *pb.StartSyncTaskResponse
	)
	cli, err := backend.NewAgentClient()
	if err != nil {
		goto ERR
	}
	res, err = cli.StartSyncTask(ctx, &pb.StartSyncTaskRequest{
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
		res *pb.MakeShellResponse
		req = &pb.MakeShellRequest{}
	)
	cli, err := backend.NewAgentClient()
	if err != nil {
		goto ERR
	}
	if err = c.ShouldBindJSON(req); err != nil {
		goto ERR
	}

	res, err = cli.MakeShell(ctx, &pb.MakeShellRequest{
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
