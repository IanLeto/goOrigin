package main

import (
	"goOrigin/agent/handlers"
	pb "goOrigin/agent/protos"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9991")
	if err != nil {
		panic(err)
	}
	//creds, err := credentials.NewServerTLSFromFile("/Users/ian/go/src/goOrigin/agent/server.crt",
	//	"/Users/ian/go/src/goOrigin/agent/server.key")
	//if err != nil {
	//	panic(err)
	//}
	server := grpc.NewServer()
	pb.RegisterAgentServer(server, &handlers.AgentHandler{})
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
