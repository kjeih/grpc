// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: grpc.proto

package grpc

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

const (
	GRPC_SubscribeMessage_FullMethodName = "/GRPC/SubscribeMessage"
	GRPC_SendMessage_FullMethodName      = "/GRPC/SendMessage"
)

// GRPCClient is the client API for GRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCClient interface {
	// Sends a greeting
	SubscribeMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (GRPC_SubscribeMessageClient, error)
	SendMessage(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error)
}

type gRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCClient(cc grpc.ClientConnInterface) GRPCClient {
	return &gRPCClient{cc}
}

func (c *gRPCClient) SubscribeMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (GRPC_SubscribeMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &GRPC_ServiceDesc.Streams[0], GRPC_SubscribeMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCSubscribeMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GRPC_SubscribeMessageClient interface {
	Recv() (*MessageReply, error)
	grpc.ClientStream
}

type gRPCSubscribeMessageClient struct {
	grpc.ClientStream
}

func (x *gRPCSubscribeMessageClient) Recv() (*MessageReply, error) {
	m := new(MessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCClient) SendMessage(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := c.cc.Invoke(ctx, GRPC_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCServer is the server API for GRPC service.
// All implementations must embed UnimplementedGRPCServer
// for forward compatibility
type GRPCServer interface {
	// Sends a greeting
	SubscribeMessage(*MessageRequest, GRPC_SubscribeMessageServer) error
	SendMessage(context.Context, *SendRequest) (*SendReply, error)
	mustEmbedUnimplementedGRPCServer()
}

// UnimplementedGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCServer struct {
}

func (UnimplementedGRPCServer) SubscribeMessage(*MessageRequest, GRPC_SubscribeMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeMessage not implemented")
}
func (UnimplementedGRPCServer) SendMessage(context.Context, *SendRequest) (*SendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedGRPCServer) mustEmbedUnimplementedGRPCServer() {}

// UnsafeGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCServer will
// result in compilation errors.
type UnsafeGRPCServer interface {
	mustEmbedUnimplementedGRPCServer()
}

func RegisterGRPCServer(s grpc.ServiceRegistrar, srv GRPCServer) {
	s.RegisterService(&GRPC_ServiceDesc, srv)
}

func _GRPC_SubscribeMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GRPCServer).SubscribeMessage(m, &gRPCSubscribeMessageServer{stream})
}

type GRPC_SubscribeMessageServer interface {
	Send(*MessageReply) error
	grpc.ServerStream
}

type gRPCSubscribeMessageServer struct {
	grpc.ServerStream
}

func (x *gRPCSubscribeMessageServer) Send(m *MessageReply) error {
	return x.ServerStream.SendMsg(m)
}

func _GRPC_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPC_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).SendMessage(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPC_ServiceDesc is the grpc.ServiceDesc for GRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GRPC",
	HandlerType: (*GRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _GRPC_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeMessage",
			Handler:       _GRPC_SubscribeMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
