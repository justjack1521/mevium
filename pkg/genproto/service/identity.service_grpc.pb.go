// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: service/identity.service.proto

package services

import (
	context "context"
	protoidentity "github.com/justjack1521/mevium/pkg/genproto/protoidentity"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MeviusIdentityService_GetSinglePlayerIdentity_FullMethodName = "/service.MeviusIdentityService/GetSinglePlayerIdentity"
	MeviusIdentityService_GetSinglePlayerLoadout_FullMethodName  = "/service.MeviusIdentityService/GetSinglePlayerLoadout"
	MeviusIdentityService_GetMultiPlayerLoadout_FullMethodName   = "/service.MeviusIdentityService/GetMultiPlayerLoadout"
)

// MeviusIdentityServiceClient is the client API for MeviusIdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeviusIdentityServiceClient interface {
	GetSinglePlayerIdentity(ctx context.Context, in *protoidentity.GetSinglePlayerIdentityRequest, opts ...grpc.CallOption) (*protoidentity.GetSinglePlayerIdentityResponse, error)
	GetSinglePlayerLoadout(ctx context.Context, in *protoidentity.GetSinglePlayerLoadoutRequest, opts ...grpc.CallOption) (*protoidentity.GetSinglePlayerLoadoutResponse, error)
	GetMultiPlayerLoadout(ctx context.Context, in *protoidentity.GetMultiPlayerLoadoutRequest, opts ...grpc.CallOption) (*protoidentity.GetMultiPlayerLoadoutResponse, error)
}

type meviusIdentityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeviusIdentityServiceClient(cc grpc.ClientConnInterface) MeviusIdentityServiceClient {
	return &meviusIdentityServiceClient{cc}
}

func (c *meviusIdentityServiceClient) GetSinglePlayerIdentity(ctx context.Context, in *protoidentity.GetSinglePlayerIdentityRequest, opts ...grpc.CallOption) (*protoidentity.GetSinglePlayerIdentityResponse, error) {
	out := new(protoidentity.GetSinglePlayerIdentityResponse)
	err := c.cc.Invoke(ctx, MeviusIdentityService_GetSinglePlayerIdentity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusIdentityServiceClient) GetSinglePlayerLoadout(ctx context.Context, in *protoidentity.GetSinglePlayerLoadoutRequest, opts ...grpc.CallOption) (*protoidentity.GetSinglePlayerLoadoutResponse, error) {
	out := new(protoidentity.GetSinglePlayerLoadoutResponse)
	err := c.cc.Invoke(ctx, MeviusIdentityService_GetSinglePlayerLoadout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusIdentityServiceClient) GetMultiPlayerLoadout(ctx context.Context, in *protoidentity.GetMultiPlayerLoadoutRequest, opts ...grpc.CallOption) (*protoidentity.GetMultiPlayerLoadoutResponse, error) {
	out := new(protoidentity.GetMultiPlayerLoadoutResponse)
	err := c.cc.Invoke(ctx, MeviusIdentityService_GetMultiPlayerLoadout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeviusIdentityServiceServer is the server API for MeviusIdentityService service.
// All implementations should embed UnimplementedMeviusIdentityServiceServer
// for forward compatibility
type MeviusIdentityServiceServer interface {
	GetSinglePlayerIdentity(context.Context, *protoidentity.GetSinglePlayerIdentityRequest) (*protoidentity.GetSinglePlayerIdentityResponse, error)
	GetSinglePlayerLoadout(context.Context, *protoidentity.GetSinglePlayerLoadoutRequest) (*protoidentity.GetSinglePlayerLoadoutResponse, error)
	GetMultiPlayerLoadout(context.Context, *protoidentity.GetMultiPlayerLoadoutRequest) (*protoidentity.GetMultiPlayerLoadoutResponse, error)
}

// UnimplementedMeviusIdentityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMeviusIdentityServiceServer struct {
}

func (UnimplementedMeviusIdentityServiceServer) GetSinglePlayerIdentity(context.Context, *protoidentity.GetSinglePlayerIdentityRequest) (*protoidentity.GetSinglePlayerIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSinglePlayerIdentity not implemented")
}
func (UnimplementedMeviusIdentityServiceServer) GetSinglePlayerLoadout(context.Context, *protoidentity.GetSinglePlayerLoadoutRequest) (*protoidentity.GetSinglePlayerLoadoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSinglePlayerLoadout not implemented")
}
func (UnimplementedMeviusIdentityServiceServer) GetMultiPlayerLoadout(context.Context, *protoidentity.GetMultiPlayerLoadoutRequest) (*protoidentity.GetMultiPlayerLoadoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiPlayerLoadout not implemented")
}

// UnsafeMeviusIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeviusIdentityServiceServer will
// result in compilation errors.
type UnsafeMeviusIdentityServiceServer interface {
	mustEmbedUnimplementedMeviusIdentityServiceServer()
}

func RegisterMeviusIdentityServiceServer(s grpc.ServiceRegistrar, srv MeviusIdentityServiceServer) {
	s.RegisterService(&MeviusIdentityService_ServiceDesc, srv)
}

func _MeviusIdentityService_GetSinglePlayerIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoidentity.GetSinglePlayerIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusIdentityServiceServer).GetSinglePlayerIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusIdentityService_GetSinglePlayerIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusIdentityServiceServer).GetSinglePlayerIdentity(ctx, req.(*protoidentity.GetSinglePlayerIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusIdentityService_GetSinglePlayerLoadout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoidentity.GetSinglePlayerLoadoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusIdentityServiceServer).GetSinglePlayerLoadout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusIdentityService_GetSinglePlayerLoadout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusIdentityServiceServer).GetSinglePlayerLoadout(ctx, req.(*protoidentity.GetSinglePlayerLoadoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusIdentityService_GetMultiPlayerLoadout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoidentity.GetMultiPlayerLoadoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusIdentityServiceServer).GetMultiPlayerLoadout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusIdentityService_GetMultiPlayerLoadout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusIdentityServiceServer).GetMultiPlayerLoadout(ctx, req.(*protoidentity.GetMultiPlayerLoadoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeviusIdentityService_ServiceDesc is the grpc.ServiceDesc for MeviusIdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeviusIdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MeviusIdentityService",
	HandlerType: (*MeviusIdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSinglePlayerIdentity",
			Handler:    _MeviusIdentityService_GetSinglePlayerIdentity_Handler,
		},
		{
			MethodName: "GetSinglePlayerLoadout",
			Handler:    _MeviusIdentityService_GetSinglePlayerLoadout_Handler,
		},
		{
			MethodName: "GetMultiPlayerLoadout",
			Handler:    _MeviusIdentityService_GetMultiPlayerLoadout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/identity.service.proto",
}
