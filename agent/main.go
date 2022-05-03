package main

import (
	"goOrigin/agent/handlers"
	pbs "goOrigin/agent/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9991")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pbs.RegisterAgentServer(server, &handlers.TaskHandler{})
	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
