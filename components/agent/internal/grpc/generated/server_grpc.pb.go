// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: server.proto

package generated

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

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Server_ConnectClient, error)
	Join(ctx context.Context, opts ...grpc.CallOption) (Server_JoinClient, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Server_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Server_ServiceDesc.Streams[0], "/server.Server/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Server_ConnectClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type serverConnectClient struct {
	grpc.ClientStream
}

func (x *serverConnectClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serverClient) Join(ctx context.Context, opts ...grpc.CallOption) (Server_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &Server_ServiceDesc.Streams[1], "/server.Server/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverJoinClient{stream}
	return x, nil
}

type Server_JoinClient interface {
	Send(*Message) error
	Recv() (*Order, error)
	grpc.ClientStream
}

type serverJoinClient struct {
	grpc.ClientStream
}

func (x *serverJoinClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverJoinClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	Connect(*ConnectRequest, Server_ConnectServer) error
	Join(Server_JoinServer) error
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) Connect(*ConnectRequest, Server_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedServerServer) Join(Server_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerServer).Connect(m, &serverConnectServer{stream})
}

type Server_ConnectServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type serverConnectServer struct {
	grpc.ServerStream
}

func (x *serverConnectServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _Server_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerServer).Join(&serverJoinServer{stream})
}

type Server_JoinServer interface {
	Send(*Order) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type serverJoinServer struct {
	grpc.ServerStream
}

func (x *serverJoinServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverJoinServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Server_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Join",
			Handler:       _Server_Join_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}
