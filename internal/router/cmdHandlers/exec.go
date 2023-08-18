package cmdHandlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goOrigin/API/V1"
	pbs "goOrigin/agent/protos"
	"goOrigin/backend"
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
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
