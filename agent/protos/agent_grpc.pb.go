// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.0
// source: agent.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	PingTask(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	RunScript(ctx context.Context, in *RunScriptRequest, opts ...grpc.CallOption) (*RunScriptResponse, error)
	RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (*RunJobResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) PingTask(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/service.v1.Agent/PingTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RunScript(ctx context.Context, in *RunScriptRequest, opts ...grpc.CallOption) (*RunScriptResponse, error) {
	out := new(RunScriptResponse)
	err := c.cc.Invoke(ctx, "/service.v1.Agent/RunScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RunJob(ctx context.Context, in *RunJobRequest, opts ...grpc.CallOption) (*RunJobResponse, error) {
	out := new(RunJobResponse)
	err := c.cc.Invoke(ctx, "/service.v1.Agent/RunJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	PingTask(context.Context, *Ping) (*Pong, error)
	RunScript(context.Context, *RunScriptRequest) (*RunScriptResponse, error)
	RunJob(context.Context, *RunJobRequest) (*RunJobResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) PingTask(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingTask not implemented")
}
func (UnimplementedAgentServer) RunScript(context.Context, *RunScriptRequest) (*RunScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunScript not implemented")
}
func (UnimplementedAgentServer) RunJob(context.Context, *RunJobRequest) (*RunJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunJob not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_PingTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).PingTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.Agent/PingTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PingTask(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RunScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RunScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.Agent/RunScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RunScript(ctx, req.(*RunScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RunJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RunJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.Agent/RunJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RunJob(ctx, req.(*RunJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingTask",
			Handler:    _Agent_PingTask_Handler,
		},
		{
			MethodName: "RunScript",
			Handler:    _Agent_RunScript_Handler,
		},
		{
			MethodName: "RunJob",
			Handler:    _Agent_RunJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
