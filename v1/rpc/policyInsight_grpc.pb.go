// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: src/grpc/policyInsight.proto

package rpc

import (
	context "context"
	common "github.com/sky-cloud-tec/proto/v1/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ObjectAnalyzeServiceClient is the client API for ObjectAnalyzeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectAnalyzeServiceClient interface {
	Analyze(ctx context.Context, in *ByDeviceRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type objectAnalyzeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectAnalyzeServiceClient(cc grpc.ClientConnInterface) ObjectAnalyzeServiceClient {
	return &objectAnalyzeServiceClient{cc}
}

func (c *objectAnalyzeServiceClient) Analyze(ctx context.Context, in *ByDeviceRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := c.cc.Invoke(ctx, "/rpc.ObjectAnalyzeService/Analyze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectAnalyzeServiceServer is the server API for ObjectAnalyzeService service.
// All implementations must embed UnimplementedObjectAnalyzeServiceServer
// for forward compatibility
type ObjectAnalyzeServiceServer interface {
	Analyze(context.Context, *ByDeviceRequest) (*AnalyzeResponse, error)
	mustEmbedUnimplementedObjectAnalyzeServiceServer()
}

// UnimplementedObjectAnalyzeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObjectAnalyzeServiceServer struct {
}

func (UnimplementedObjectAnalyzeServiceServer) Analyze(context.Context, *ByDeviceRequest) (*AnalyzeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedObjectAnalyzeServiceServer) mustEmbedUnimplementedObjectAnalyzeServiceServer() {}

// UnsafeObjectAnalyzeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectAnalyzeServiceServer will
// result in compilation errors.
type UnsafeObjectAnalyzeServiceServer interface {
	mustEmbedUnimplementedObjectAnalyzeServiceServer()
}

func RegisterObjectAnalyzeServiceServer(s grpc.ServiceRegistrar, srv ObjectAnalyzeServiceServer) {
	s.RegisterService(&ObjectAnalyzeService_ServiceDesc, srv)
}

func _ObjectAnalyzeService_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectAnalyzeServiceServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ObjectAnalyzeService/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectAnalyzeServiceServer).Analyze(ctx, req.(*ByDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectAnalyzeService_ServiceDesc is the grpc.ServiceDesc for ObjectAnalyzeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectAnalyzeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ObjectAnalyzeService",
	HandlerType: (*ObjectAnalyzeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _ObjectAnalyzeService_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}

// PolicyOptimizationClient is the client API for PolicyOptimization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyOptimizationClient interface {
	PermissivePolicyClustering(ctx context.Context, in *BySessionsAndParameters, opts ...grpc.CallOption) (*PermissivePolicyClusteringResponse, error)
}

type policyOptimizationClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyOptimizationClient(cc grpc.ClientConnInterface) PolicyOptimizationClient {
	return &policyOptimizationClient{cc}
}

func (c *policyOptimizationClient) PermissivePolicyClustering(ctx context.Context, in *BySessionsAndParameters, opts ...grpc.CallOption) (*PermissivePolicyClusteringResponse, error) {
	out := new(PermissivePolicyClusteringResponse)
	err := c.cc.Invoke(ctx, "/rpc.PolicyOptimization/PermissivePolicyClustering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyOptimizationServer is the server API for PolicyOptimization service.
// All implementations must embed UnimplementedPolicyOptimizationServer
// for forward compatibility
type PolicyOptimizationServer interface {
	PermissivePolicyClustering(context.Context, *BySessionsAndParameters) (*PermissivePolicyClusteringResponse, error)
	mustEmbedUnimplementedPolicyOptimizationServer()
}

// UnimplementedPolicyOptimizationServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyOptimizationServer struct {
}

func (UnimplementedPolicyOptimizationServer) PermissivePolicyClustering(context.Context, *BySessionsAndParameters) (*PermissivePolicyClusteringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissivePolicyClustering not implemented")
}
func (UnimplementedPolicyOptimizationServer) mustEmbedUnimplementedPolicyOptimizationServer() {}

// UnsafePolicyOptimizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyOptimizationServer will
// result in compilation errors.
type UnsafePolicyOptimizationServer interface {
	mustEmbedUnimplementedPolicyOptimizationServer()
}

func RegisterPolicyOptimizationServer(s grpc.ServiceRegistrar, srv PolicyOptimizationServer) {
	s.RegisterService(&PolicyOptimization_ServiceDesc, srv)
}

func _PolicyOptimization_PermissivePolicyClustering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BySessionsAndParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyOptimizationServer).PermissivePolicyClustering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.PolicyOptimization/PermissivePolicyClustering",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyOptimizationServer).PermissivePolicyClustering(ctx, req.(*BySessionsAndParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyOptimization_ServiceDesc is the grpc.ServiceDesc for PolicyOptimization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyOptimization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.PolicyOptimization",
	HandlerType: (*PolicyOptimizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PermissivePolicyClustering",
			Handler:    _PolicyOptimization_PermissivePolicyClustering_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}

// PolicyAnalyzeServiceClient is the client API for PolicyAnalyzeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyAnalyzeServiceClient interface {
	Analyze(ctx context.Context, in *ByDeviceRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type policyAnalyzeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyAnalyzeServiceClient(cc grpc.ClientConnInterface) PolicyAnalyzeServiceClient {
	return &policyAnalyzeServiceClient{cc}
}

func (c *policyAnalyzeServiceClient) Analyze(ctx context.Context, in *ByDeviceRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := c.cc.Invoke(ctx, "/rpc.PolicyAnalyzeService/Analyze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyAnalyzeServiceServer is the server API for PolicyAnalyzeService service.
// All implementations must embed UnimplementedPolicyAnalyzeServiceServer
// for forward compatibility
type PolicyAnalyzeServiceServer interface {
	Analyze(context.Context, *ByDeviceRequest) (*AnalyzeResponse, error)
	mustEmbedUnimplementedPolicyAnalyzeServiceServer()
}

// UnimplementedPolicyAnalyzeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyAnalyzeServiceServer struct {
}

func (UnimplementedPolicyAnalyzeServiceServer) Analyze(context.Context, *ByDeviceRequest) (*AnalyzeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedPolicyAnalyzeServiceServer) mustEmbedUnimplementedPolicyAnalyzeServiceServer() {}

// UnsafePolicyAnalyzeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyAnalyzeServiceServer will
// result in compilation errors.
type UnsafePolicyAnalyzeServiceServer interface {
	mustEmbedUnimplementedPolicyAnalyzeServiceServer()
}

func RegisterPolicyAnalyzeServiceServer(s grpc.ServiceRegistrar, srv PolicyAnalyzeServiceServer) {
	s.RegisterService(&PolicyAnalyzeService_ServiceDesc, srv)
}

func _PolicyAnalyzeService_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyAnalyzeServiceServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.PolicyAnalyzeService/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyAnalyzeServiceServer).Analyze(ctx, req.(*ByDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyAnalyzeService_ServiceDesc is the grpc.ServiceDesc for PolicyAnalyzeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyAnalyzeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.PolicyAnalyzeService",
	HandlerType: (*PolicyAnalyzeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _PolicyAnalyzeService_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}

// PolicyConvergenceServiceClient is the client API for PolicyConvergenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyConvergenceServiceClient interface {
	Convergence(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*common.SimpleJsonResponse, error)
}

type policyConvergenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyConvergenceServiceClient(cc grpc.ClientConnInterface) PolicyConvergenceServiceClient {
	return &policyConvergenceServiceClient{cc}
}

func (c *policyConvergenceServiceClient) Convergence(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*common.SimpleJsonResponse, error) {
	out := new(common.SimpleJsonResponse)
	err := c.cc.Invoke(ctx, "/rpc.PolicyConvergenceService/Convergence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyConvergenceServiceServer is the server API for PolicyConvergenceService service.
// All implementations must embed UnimplementedPolicyConvergenceServiceServer
// for forward compatibility
type PolicyConvergenceServiceServer interface {
	Convergence(context.Context, *ByIdRequest) (*common.SimpleJsonResponse, error)
	mustEmbedUnimplementedPolicyConvergenceServiceServer()
}

// UnimplementedPolicyConvergenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyConvergenceServiceServer struct {
}

func (UnimplementedPolicyConvergenceServiceServer) Convergence(context.Context, *ByIdRequest) (*common.SimpleJsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convergence not implemented")
}
func (UnimplementedPolicyConvergenceServiceServer) mustEmbedUnimplementedPolicyConvergenceServiceServer() {
}

// UnsafePolicyConvergenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyConvergenceServiceServer will
// result in compilation errors.
type UnsafePolicyConvergenceServiceServer interface {
	mustEmbedUnimplementedPolicyConvergenceServiceServer()
}

func RegisterPolicyConvergenceServiceServer(s grpc.ServiceRegistrar, srv PolicyConvergenceServiceServer) {
	s.RegisterService(&PolicyConvergenceService_ServiceDesc, srv)
}

func _PolicyConvergenceService_Convergence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyConvergenceServiceServer).Convergence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.PolicyConvergenceService/Convergence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyConvergenceServiceServer).Convergence(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyConvergenceService_ServiceDesc is the grpc.ServiceDesc for PolicyConvergenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyConvergenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.PolicyConvergenceService",
	HandlerType: (*PolicyConvergenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convergence",
			Handler:    _PolicyConvergenceService_Convergence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}

// ComplianceAnalyzeServiceClient is the client API for ComplianceAnalyzeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplianceAnalyzeServiceClient interface {
	Analyze(ctx context.Context, in *ByDeviceRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type complianceAnalyzeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceAnalyzeServiceClient(cc grpc.ClientConnInterface) ComplianceAnalyzeServiceClient {
	return &complianceAnalyzeServiceClient{cc}
}

func (c *complianceAnalyzeServiceClient) Analyze(ctx context.Context, in *ByDeviceRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := c.cc.Invoke(ctx, "/rpc.ComplianceAnalyzeService/Analyze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceAnalyzeServiceServer is the server API for ComplianceAnalyzeService service.
// All implementations must embed UnimplementedComplianceAnalyzeServiceServer
// for forward compatibility
type ComplianceAnalyzeServiceServer interface {
	Analyze(context.Context, *ByDeviceRequest) (*AnalyzeResponse, error)
	mustEmbedUnimplementedComplianceAnalyzeServiceServer()
}

// UnimplementedComplianceAnalyzeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComplianceAnalyzeServiceServer struct {
}

func (UnimplementedComplianceAnalyzeServiceServer) Analyze(context.Context, *ByDeviceRequest) (*AnalyzeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedComplianceAnalyzeServiceServer) mustEmbedUnimplementedComplianceAnalyzeServiceServer() {
}

// UnsafeComplianceAnalyzeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplianceAnalyzeServiceServer will
// result in compilation errors.
type UnsafeComplianceAnalyzeServiceServer interface {
	mustEmbedUnimplementedComplianceAnalyzeServiceServer()
}

func RegisterComplianceAnalyzeServiceServer(s grpc.ServiceRegistrar, srv ComplianceAnalyzeServiceServer) {
	s.RegisterService(&ComplianceAnalyzeService_ServiceDesc, srv)
}

func _ComplianceAnalyzeService_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceAnalyzeServiceServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ComplianceAnalyzeService/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceAnalyzeServiceServer).Analyze(ctx, req.(*ByDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComplianceAnalyzeService_ServiceDesc is the grpc.ServiceDesc for ComplianceAnalyzeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplianceAnalyzeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ComplianceAnalyzeService",
	HandlerType: (*ComplianceAnalyzeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _ComplianceAnalyzeService_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}

// ContrailServiceRpcClient is the client API for ContrailServiceRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContrailServiceRpcClient interface {
	CreatePipeline(ctx context.Context, in *ByContrailTicketRequest, opts ...grpc.CallOption) (*common.EmptyResponse, error)
}

type contrailServiceRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewContrailServiceRpcClient(cc grpc.ClientConnInterface) ContrailServiceRpcClient {
	return &contrailServiceRpcClient{cc}
}

func (c *contrailServiceRpcClient) CreatePipeline(ctx context.Context, in *ByContrailTicketRequest, opts ...grpc.CallOption) (*common.EmptyResponse, error) {
	out := new(common.EmptyResponse)
	err := c.cc.Invoke(ctx, "/rpc.ContrailServiceRpc/CreatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContrailServiceRpcServer is the server API for ContrailServiceRpc service.
// All implementations must embed UnimplementedContrailServiceRpcServer
// for forward compatibility
type ContrailServiceRpcServer interface {
	CreatePipeline(context.Context, *ByContrailTicketRequest) (*common.EmptyResponse, error)
	mustEmbedUnimplementedContrailServiceRpcServer()
}

// UnimplementedContrailServiceRpcServer must be embedded to have forward compatible implementations.
type UnimplementedContrailServiceRpcServer struct {
}

func (UnimplementedContrailServiceRpcServer) CreatePipeline(context.Context, *ByContrailTicketRequest) (*common.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedContrailServiceRpcServer) mustEmbedUnimplementedContrailServiceRpcServer() {}

// UnsafeContrailServiceRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContrailServiceRpcServer will
// result in compilation errors.
type UnsafeContrailServiceRpcServer interface {
	mustEmbedUnimplementedContrailServiceRpcServer()
}

func RegisterContrailServiceRpcServer(s grpc.ServiceRegistrar, srv ContrailServiceRpcServer) {
	s.RegisterService(&ContrailServiceRpc_ServiceDesc, srv)
}

func _ContrailServiceRpc_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByContrailTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContrailServiceRpcServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ContrailServiceRpc/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContrailServiceRpcServer).CreatePipeline(ctx, req.(*ByContrailTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContrailServiceRpc_ServiceDesc is the grpc.ServiceDesc for ContrailServiceRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContrailServiceRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ContrailServiceRpc",
	HandlerType: (*ContrailServiceRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _ContrailServiceRpc_CreatePipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	Execute(ctx context.Context, in *ReportExecuteRequest, opts ...grpc.CallOption) (*ReportExecuteResponse, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) Execute(ctx context.Context, in *ReportExecuteRequest, opts ...grpc.CallOption) (*ReportExecuteResponse, error) {
	out := new(ReportExecuteResponse)
	err := c.cc.Invoke(ctx, "/rpc.ReportService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations must embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	Execute(context.Context, *ReportExecuteRequest) (*ReportExecuteResponse, error)
	mustEmbedUnimplementedReportServiceServer()
}

// UnimplementedReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) Execute(context.Context, *ReportExecuteRequest) (*ReportExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedReportServiceServer) mustEmbedUnimplementedReportServiceServer() {}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ReportService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).Execute(ctx, req.(*ReportExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _ReportService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/grpc/policyInsight.proto",
}
