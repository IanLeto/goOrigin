package backend

import (
	pb "goOrigin/agent/pbs/service"
	"google.golang.org/grpc"
)

func NewAgentClient() (pb.AgentClient,error) {
	conn, err := grpc.Dial("localhost:9991", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewAgentClient(conn), nil
}
