// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: src/grpc/platform.proto

package rpc

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

// UserInfoServiceClient is the client API for UserInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoServiceClient interface {
	GetList(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoRepeatResponse, error)
	GetByUsername(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type userInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoServiceClient(cc grpc.ClientConnInterface) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) GetList(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoRepeatResponse, error) {
	out := new(UserInfoRepeatResponse)
	err := c.cc.Invoke(ctx, "/rpc.UserInfoService/getList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoServiceClient) GetByUsername(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/rpc.UserInfoService/getByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServiceServer is the server API for UserInfoService service.
// All implementations must embed UnimplementedUserInfoServiceServer
// for forward compatibility
type UserInfoServiceServer interface {
	GetList(context.Context, *UserInfoRequest) (*UserInfoRepeatResponse, error)
	GetByUsername(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	mustEmbedUnimplementedUserInfoServiceServer()
}

// UnimplementedUserInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServiceServer struct {
}

func (UnimplementedUserInfoServiceServer) GetList(context.Context, *UserInfoRequest) (*UserInfoRepeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedUserInfoServiceServer) GetByUsername(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (UnimplementedUserInfoServiceServer) mustEmbedUnimplementedUserInfoServiceServer() {}

// UnsafeUserInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServiceServer will
// result in compilation errors.
type UnsafeUserInfoServiceServer interface {
	mustEmbedUnimplementedUserInfoServiceServer()
}

func RegisterUserInfoServiceServer(s grpc.ServiceRegistrar, srv UserInfoServiceServer) {
	s.RegisterService(&UserInfoService_ServiceDesc, srv)
}

func _UserInfoService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UserInfoService/getList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetList(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoService_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UserInfoService/getByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetByUsername(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfoService_ServiceDesc is the grpc.ServiceDesc for UserInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getList",
			Handler:    _UserInfoService_GetList_Handler,
		},
		{
			MethodName: "getByUsername",
			Handler:    _UserInfoService_GetByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/platform.proto",
}