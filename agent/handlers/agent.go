package handlers

import (
	"context"
	pb "goOrigin/agent/protos"
	"goOrigin/agent/service"
	"log"
)

type AgentHandler struct {
	pb.UnimplementedAgentServer // 这个东西是防止重新编译proto 报错
}

func (a AgentHandler) PingTask(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	var (
		res = &pb.Pong{}
		err error
	)
	if err != nil {
		log.Printf("初始化失败 %s", err)
	}
	service.Pong()
	res.Version = "v0.0.1"
	return res, err

}
