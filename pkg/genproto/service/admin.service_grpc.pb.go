// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: service/admin.service.proto

package services

import (
	context "context"
	protoadmin "github.com/justjack1521/mevium/pkg/genproto/protoadmin"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MeviusAdminService_GrantItem_FullMethodName         = "/service.MeviusAdminService/GrantItem"
	MeviusAdminService_CreateBaseJobCard_FullMethodName = "/service.MeviusAdminService/CreateBaseJobCard"
	MeviusAdminService_CreateSkillPanel_FullMethodName  = "/service.MeviusAdminService/CreateSkillPanel"
)

// MeviusAdminServiceClient is the client API for MeviusAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeviusAdminServiceClient interface {
	GrantItem(ctx context.Context, in *protoadmin.GrantItemRequest, opts ...grpc.CallOption) (*protoadmin.GrantItemResponse, error)
	CreateBaseJobCard(ctx context.Context, in *protoadmin.CreateBaseJobCardRequest, opts ...grpc.CallOption) (*protoadmin.CreateBaseJobCardResponse, error)
	CreateSkillPanel(ctx context.Context, in *protoadmin.CreateSkillPanelRequest, opts ...grpc.CallOption) (*protoadmin.CreateSkillPanelResponse, error)
}

type meviusAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeviusAdminServiceClient(cc grpc.ClientConnInterface) MeviusAdminServiceClient {
	return &meviusAdminServiceClient{cc}
}

func (c *meviusAdminServiceClient) GrantItem(ctx context.Context, in *protoadmin.GrantItemRequest, opts ...grpc.CallOption) (*protoadmin.GrantItemResponse, error) {
	out := new(protoadmin.GrantItemResponse)
	err := c.cc.Invoke(ctx, MeviusAdminService_GrantItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusAdminServiceClient) CreateBaseJobCard(ctx context.Context, in *protoadmin.CreateBaseJobCardRequest, opts ...grpc.CallOption) (*protoadmin.CreateBaseJobCardResponse, error) {
	out := new(protoadmin.CreateBaseJobCardResponse)
	err := c.cc.Invoke(ctx, MeviusAdminService_CreateBaseJobCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusAdminServiceClient) CreateSkillPanel(ctx context.Context, in *protoadmin.CreateSkillPanelRequest, opts ...grpc.CallOption) (*protoadmin.CreateSkillPanelResponse, error) {
	out := new(protoadmin.CreateSkillPanelResponse)
	err := c.cc.Invoke(ctx, MeviusAdminService_CreateSkillPanel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeviusAdminServiceServer is the server API for MeviusAdminService service.
// All implementations should embed UnimplementedMeviusAdminServiceServer
// for forward compatibility
type MeviusAdminServiceServer interface {
	GrantItem(context.Context, *protoadmin.GrantItemRequest) (*protoadmin.GrantItemResponse, error)
	CreateBaseJobCard(context.Context, *protoadmin.CreateBaseJobCardRequest) (*protoadmin.CreateBaseJobCardResponse, error)
	CreateSkillPanel(context.Context, *protoadmin.CreateSkillPanelRequest) (*protoadmin.CreateSkillPanelResponse, error)
}

// UnimplementedMeviusAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMeviusAdminServiceServer struct {
}

func (UnimplementedMeviusAdminServiceServer) GrantItem(context.Context, *protoadmin.GrantItemRequest) (*protoadmin.GrantItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantItem not implemented")
}
func (UnimplementedMeviusAdminServiceServer) CreateBaseJobCard(context.Context, *protoadmin.CreateBaseJobCardRequest) (*protoadmin.CreateBaseJobCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBaseJobCard not implemented")
}
func (UnimplementedMeviusAdminServiceServer) CreateSkillPanel(context.Context, *protoadmin.CreateSkillPanelRequest) (*protoadmin.CreateSkillPanelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSkillPanel not implemented")
}

// UnsafeMeviusAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeviusAdminServiceServer will
// result in compilation errors.
type UnsafeMeviusAdminServiceServer interface {
	mustEmbedUnimplementedMeviusAdminServiceServer()
}

func RegisterMeviusAdminServiceServer(s grpc.ServiceRegistrar, srv MeviusAdminServiceServer) {
	s.RegisterService(&MeviusAdminService_ServiceDesc, srv)
}

func _MeviusAdminService_GrantItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoadmin.GrantItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusAdminServiceServer).GrantItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusAdminService_GrantItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusAdminServiceServer).GrantItem(ctx, req.(*protoadmin.GrantItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusAdminService_CreateBaseJobCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoadmin.CreateBaseJobCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusAdminServiceServer).CreateBaseJobCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusAdminService_CreateBaseJobCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusAdminServiceServer).CreateBaseJobCard(ctx, req.(*protoadmin.CreateBaseJobCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusAdminService_CreateSkillPanel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoadmin.CreateSkillPanelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusAdminServiceServer).CreateSkillPanel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusAdminService_CreateSkillPanel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusAdminServiceServer).CreateSkillPanel(ctx, req.(*protoadmin.CreateSkillPanelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeviusAdminService_ServiceDesc is the grpc.ServiceDesc for MeviusAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeviusAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MeviusAdminService",
	HandlerType: (*MeviusAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrantItem",
			Handler:    _MeviusAdminService_GrantItem_Handler,
		},
		{
			MethodName: "CreateBaseJobCard",
			Handler:    _MeviusAdminService_CreateBaseJobCard_Handler,
		},
		{
			MethodName: "CreateSkillPanel",
			Handler:    _MeviusAdminService_CreateSkillPanel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/admin.service.proto",
}
