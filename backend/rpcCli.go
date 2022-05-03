package backend

import (
	pbs "goOrigin/agent/pb"
	"google.golang.org/grpc"
)

func NewAgentClient() (pbs.AgentClient,error) {

	conn, err := grpc.Dial("localhost:9991", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pbs.NewAgentClient(conn), nil
}
