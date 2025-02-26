// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proxy.proto

package api

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

// ShareBuildProxyClient is the client API for ShareBuildProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShareBuildProxyClient interface {
	InitializeBuildEnv(ctx context.Context, in *InitializeBuildEnvRequest, opts ...grpc.CallOption) (*InitializeBuildEnvResponse, error)
	ForwardAndExecute(ctx context.Context, in *ForwardAndExecuteRequest, opts ...grpc.CallOption) (*ForwardAndExecuteResponse, error)
	ClearBuildEnv(ctx context.Context, in *ClearBuildEnvRequest, opts ...grpc.CallOption) (*ClearBuildEnvResponse, error)
}

type shareBuildProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewShareBuildProxyClient(cc grpc.ClientConnInterface) ShareBuildProxyClient {
	return &shareBuildProxyClient{cc}
}

func (c *shareBuildProxyClient) InitializeBuildEnv(ctx context.Context, in *InitializeBuildEnvRequest, opts ...grpc.CallOption) (*InitializeBuildEnvResponse, error) {
	out := new(InitializeBuildEnvResponse)
	err := c.cc.Invoke(ctx, "/api.ShareBuildProxy/InitializeBuildEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareBuildProxyClient) ForwardAndExecute(ctx context.Context, in *ForwardAndExecuteRequest, opts ...grpc.CallOption) (*ForwardAndExecuteResponse, error) {
	out := new(ForwardAndExecuteResponse)
	err := c.cc.Invoke(ctx, "/api.ShareBuildProxy/ForwardAndExecute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareBuildProxyClient) ClearBuildEnv(ctx context.Context, in *ClearBuildEnvRequest, opts ...grpc.CallOption) (*ClearBuildEnvResponse, error) {
	out := new(ClearBuildEnvResponse)
	err := c.cc.Invoke(ctx, "/api.ShareBuildProxy/ClearBuildEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareBuildProxyServer is the server API for ShareBuildProxy service.
// All implementations must embed UnimplementedShareBuildProxyServer
// for forward compatibility
type ShareBuildProxyServer interface {
	InitializeBuildEnv(context.Context, *InitializeBuildEnvRequest) (*InitializeBuildEnvResponse, error)
	ForwardAndExecute(context.Context, *ForwardAndExecuteRequest) (*ForwardAndExecuteResponse, error)
	ClearBuildEnv(context.Context, *ClearBuildEnvRequest) (*ClearBuildEnvResponse, error)
	mustEmbedUnimplementedShareBuildProxyServer()
}

// UnimplementedShareBuildProxyServer must be embedded to have forward compatible implementations.
type UnimplementedShareBuildProxyServer struct {
}

func (UnimplementedShareBuildProxyServer) InitializeBuildEnv(context.Context, *InitializeBuildEnvRequest) (*InitializeBuildEnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeBuildEnv not implemented")
}
func (UnimplementedShareBuildProxyServer) ForwardAndExecute(context.Context, *ForwardAndExecuteRequest) (*ForwardAndExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardAndExecute not implemented")
}
func (UnimplementedShareBuildProxyServer) ClearBuildEnv(context.Context, *ClearBuildEnvRequest) (*ClearBuildEnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearBuildEnv not implemented")
}
func (UnimplementedShareBuildProxyServer) mustEmbedUnimplementedShareBuildProxyServer() {}

// UnsafeShareBuildProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShareBuildProxyServer will
// result in compilation errors.
type UnsafeShareBuildProxyServer interface {
	mustEmbedUnimplementedShareBuildProxyServer()
}

func RegisterShareBuildProxyServer(s grpc.ServiceRegistrar, srv ShareBuildProxyServer) {
	s.RegisterService(&ShareBuildProxy_ServiceDesc, srv)
}

func _ShareBuildProxy_InitializeBuildEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeBuildEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareBuildProxyServer).InitializeBuildEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShareBuildProxy/InitializeBuildEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareBuildProxyServer).InitializeBuildEnv(ctx, req.(*InitializeBuildEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareBuildProxy_ForwardAndExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardAndExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareBuildProxyServer).ForwardAndExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShareBuildProxy/ForwardAndExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareBuildProxyServer).ForwardAndExecute(ctx, req.(*ForwardAndExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareBuildProxy_ClearBuildEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearBuildEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareBuildProxyServer).ClearBuildEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShareBuildProxy/ClearBuildEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareBuildProxyServer).ClearBuildEnv(ctx, req.(*ClearBuildEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShareBuildProxy_ServiceDesc is the grpc.ServiceDesc for ShareBuildProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShareBuildProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ShareBuildProxy",
	HandlerType: (*ShareBuildProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeBuildEnv",
			Handler:    _ShareBuildProxy_InitializeBuildEnv_Handler,
		},
		{
			MethodName: "ForwardAndExecute",
			Handler:    _ShareBuildProxy_ForwardAndExecute_Handler,
		},
		{
			MethodName: "ClearBuildEnv",
			Handler:    _ShareBuildProxy_ClearBuildEnv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}
