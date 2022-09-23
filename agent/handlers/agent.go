package handlers

import (
	"context"
	pb "goOrigin/agent/protos"
	"goOrigin/agent/service"
	"log"
	"os"
	"os/exec"
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

func (a AgentHandler) RunScript(ctx context.Context, req *pb.RunScriptRequest) (*pb.RunScriptResponse, error) {
	var (
		res = &pb.RunScriptResponse{}
	)

	cmd := exec.Command("bash", "-c", req.Content)
	cmd.Stderr = os.Stderr
	d, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	log.Println(string(d))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	res.Result = string(out)
	return res, err
}

func (a AgentHandler) RunJob(ctx context.Context, req *pb.RunJobRequest) (*pb.RunJobResponse, error) {
	var (
		res = &pb.RunJobResponse{}
		err error
	)
	for _, s := range req.GetContents() {
		cmd := exec.Command("bash", "-c", s)
		cmd.Stderr = os.Stderr
		d, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		log.Println(string(d))
		out, err := cmd.CombinedOutput()
		if err != nil {
			return nil, err
		}
		res.Result = append(res.Result, string(out))
	}
	return res, err

}
