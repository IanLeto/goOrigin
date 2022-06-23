package cmdHandlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pbs "goOrigin/agent/protos"
	"goOrigin/backend"
	"goOrigin/internal/params"
)

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
