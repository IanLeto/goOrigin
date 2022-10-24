package backend

import (
	pbs "goOrigin/agent/protos"
	"google.golang.org/grpc"
)

func NewAgentClient() (pbs.AgentClient, error) {
	//cred, err := credentials.NewClientTLSFromFile("/Users/ian/go/src/goOrigin/agent/server.crt", "goOrigin")
	//if err != nil {
	//	return nil, err
	//}
	//conn, err := grpc.Dial("localhost:9991", grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial("localhost:9991", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pbs.NewAgentClient(conn), nil
}
