// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package localnetv1

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

// EndpointsClient is the client API for Endpoints service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointsClient interface {
	// Returns all the endpoints for this node.
	Watch(ctx context.Context, opts ...grpc.CallOption) (Endpoints_WatchClient, error)
}

type endpointsClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointsClient(cc grpc.ClientConnInterface) EndpointsClient {
	return &endpointsClient{cc}
}

func (c *endpointsClient) Watch(ctx context.Context, opts ...grpc.CallOption) (Endpoints_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Endpoints_ServiceDesc.Streams[0], "/localnetv1.Endpoints/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointsWatchClient{stream}
	return x, nil
}

type Endpoints_WatchClient interface {
	Send(*WatchReq) error
	Recv() (*OpItem, error)
	grpc.ClientStream
}

type endpointsWatchClient struct {
	grpc.ClientStream
}

func (x *endpointsWatchClient) Send(m *WatchReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointsWatchClient) Recv() (*OpItem, error) {
	m := new(OpItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EndpointsServer is the server API for Endpoints service.
// All implementations must embed UnimplementedEndpointsServer
// for forward compatibility
type EndpointsServer interface {
	// Returns all the endpoints for this node.
	Watch(Endpoints_WatchServer) error
	mustEmbedUnimplementedEndpointsServer()
}

// UnimplementedEndpointsServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointsServer struct {
}

func (UnimplementedEndpointsServer) Watch(Endpoints_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedEndpointsServer) mustEmbedUnimplementedEndpointsServer() {}

// UnsafeEndpointsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointsServer will
// result in compilation errors.
type UnsafeEndpointsServer interface {
	mustEmbedUnimplementedEndpointsServer()
}

func RegisterEndpointsServer(s grpc.ServiceRegistrar, srv EndpointsServer) {
	s.RegisterService(&Endpoints_ServiceDesc, srv)
}

func _Endpoints_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointsServer).Watch(&endpointsWatchServer{stream})
}

type Endpoints_WatchServer interface {
	Send(*OpItem) error
	Recv() (*WatchReq, error)
	grpc.ServerStream
}

type endpointsWatchServer struct {
	grpc.ServerStream
}

func (x *endpointsWatchServer) Send(m *OpItem) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointsWatchServer) Recv() (*WatchReq, error) {
	m := new(WatchReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Endpoints_ServiceDesc is the grpc.ServiceDesc for Endpoints service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Endpoints_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "localnetv1.Endpoints",
	HandlerType: (*EndpointsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Endpoints_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/localnetv1/services.proto",
}

// GlobalClient is the client API for Global service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalClient interface {
	Watch(ctx context.Context, opts ...grpc.CallOption) (Global_WatchClient, error)
}

type globalClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalClient(cc grpc.ClientConnInterface) GlobalClient {
	return &globalClient{cc}
}

func (c *globalClient) Watch(ctx context.Context, opts ...grpc.CallOption) (Global_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Global_ServiceDesc.Streams[0], "/localnetv1.Global/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &globalWatchClient{stream}
	return x, nil
}

type Global_WatchClient interface {
	Send(*GlobalWatchReq) error
	Recv() (*OpItem, error)
	grpc.ClientStream
}

type globalWatchClient struct {
	grpc.ClientStream
}

func (x *globalWatchClient) Send(m *GlobalWatchReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *globalWatchClient) Recv() (*OpItem, error) {
	m := new(OpItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GlobalServer is the server API for Global service.
// All implementations must embed UnimplementedGlobalServer
// for forward compatibility
type GlobalServer interface {
	Watch(Global_WatchServer) error
	mustEmbedUnimplementedGlobalServer()
}

// UnimplementedGlobalServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalServer struct {
}

func (UnimplementedGlobalServer) Watch(Global_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedGlobalServer) mustEmbedUnimplementedGlobalServer() {}

// UnsafeGlobalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalServer will
// result in compilation errors.
type UnsafeGlobalServer interface {
	mustEmbedUnimplementedGlobalServer()
}

func RegisterGlobalServer(s grpc.ServiceRegistrar, srv GlobalServer) {
	s.RegisterService(&Global_ServiceDesc, srv)
}

func _Global_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GlobalServer).Watch(&globalWatchServer{stream})
}

type Global_WatchServer interface {
	Send(*OpItem) error
	Recv() (*GlobalWatchReq, error)
	grpc.ServerStream
}

type globalWatchServer struct {
	grpc.ServerStream
}

func (x *globalWatchServer) Send(m *OpItem) error {
	return x.ServerStream.SendMsg(m)
}

func (x *globalWatchServer) Recv() (*GlobalWatchReq, error) {
	m := new(GlobalWatchReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Global_ServiceDesc is the grpc.ServiceDesc for Global service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Global_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "localnetv1.Global",
	HandlerType: (*GlobalServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Global_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/localnetv1/services.proto",
	
	
}
