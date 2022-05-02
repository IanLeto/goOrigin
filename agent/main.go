package main

import (
	"goOrigin/agent/handlers"
	pb "goOrigin/agent/pbs/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9991")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentServer(server, &handlers.Task{})
	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
