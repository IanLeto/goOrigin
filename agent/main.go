package main

import (
	pb "goOrigin/agent/pbs/service"
	"goOrigin/agent/task"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9991")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterAgentServer(server, &task.Task{})
	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
