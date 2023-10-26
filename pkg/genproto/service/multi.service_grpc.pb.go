// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: service/multi.service.proto

package services

import (
	context "context"
	protomulti "github.com/justjack1521/mevium/pkg/genproto/protomulti"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MeviusMultiService_CreateSession_FullMethodName = "/service.MeviusMultiService/CreateSession"
	MeviusMultiService_EndSession_FullMethodName    = "/service.MeviusMultiService/EndSession"
	MeviusMultiService_SearchLobby_FullMethodName   = "/service.MeviusMultiService/SearchLobby"
	MeviusMultiService_WatchLobby_FullMethodName    = "/service.MeviusMultiService/WatchLobby"
	MeviusMultiService_DiscardLobby_FullMethodName  = "/service.MeviusMultiService/DiscardLobby"
	MeviusMultiService_CreateLobby_FullMethodName   = "/service.MeviusMultiService/CreateLobby"
	MeviusMultiService_JoinLobby_FullMethodName     = "/service.MeviusMultiService/JoinLobby"
	MeviusMultiService_ReadyLobby_FullMethodName    = "/service.MeviusMultiService/ReadyLobby"
	MeviusMultiService_UnreadyLobby_FullMethodName  = "/service.MeviusMultiService/UnreadyLobby"
	MeviusMultiService_CancelLobby_FullMethodName   = "/service.MeviusMultiService/CancelLobby"
	MeviusMultiService_GetPlayer_FullMethodName     = "/service.MeviusMultiService/GetPlayer"
)

// MeviusMultiServiceClient is the client API for MeviusMultiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeviusMultiServiceClient interface {
	CreateSession(ctx context.Context, in *protomulti.CreateSessionRequest, opts ...grpc.CallOption) (*protomulti.CreateSessionResponse, error)
	EndSession(ctx context.Context, in *protomulti.EndSessionRequest, opts ...grpc.CallOption) (*protomulti.EndSessionResponse, error)
	SearchLobby(ctx context.Context, in *protomulti.SearchLobbyRequest, opts ...grpc.CallOption) (*protomulti.SearchLobbyResponse, error)
	WatchLobby(ctx context.Context, in *protomulti.WatchLobbyRequest, opts ...grpc.CallOption) (*protomulti.WatchLobbyResponse, error)
	DiscardLobby(ctx context.Context, in *protomulti.DiscardLobbyRequest, opts ...grpc.CallOption) (*protomulti.DiscardLobbyResponse, error)
	CreateLobby(ctx context.Context, in *protomulti.CreateLobbyRequest, opts ...grpc.CallOption) (*protomulti.CreateLobbyResponse, error)
	JoinLobby(ctx context.Context, in *protomulti.JoinLobbyRequest, opts ...grpc.CallOption) (*protomulti.JoinLobbyResponse, error)
	ReadyLobby(ctx context.Context, in *protomulti.ReadyLobbyRequest, opts ...grpc.CallOption) (*protomulti.ReadyLobbyResponse, error)
	UnreadyLobby(ctx context.Context, in *protomulti.UnreadyLobbyRequest, opts ...grpc.CallOption) (*protomulti.UnreadyLobbyResponse, error)
	CancelLobby(ctx context.Context, in *protomulti.CancelLobbyRequest, opts ...grpc.CallOption) (*protomulti.CancelLobbyResponse, error)
	GetPlayer(ctx context.Context, in *protomulti.GetLobbyPlayerRequest, opts ...grpc.CallOption) (*protomulti.GetLobbyPlayerResponse, error)
}

type meviusMultiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeviusMultiServiceClient(cc grpc.ClientConnInterface) MeviusMultiServiceClient {
	return &meviusMultiServiceClient{cc}
}

func (c *meviusMultiServiceClient) CreateSession(ctx context.Context, in *protomulti.CreateSessionRequest, opts ...grpc.CallOption) (*protomulti.CreateSessionResponse, error) {
	out := new(protomulti.CreateSessionResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_CreateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) EndSession(ctx context.Context, in *protomulti.EndSessionRequest, opts ...grpc.CallOption) (*protomulti.EndSessionResponse, error) {
	out := new(protomulti.EndSessionResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_EndSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) SearchLobby(ctx context.Context, in *protomulti.SearchLobbyRequest, opts ...grpc.CallOption) (*protomulti.SearchLobbyResponse, error) {
	out := new(protomulti.SearchLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_SearchLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) WatchLobby(ctx context.Context, in *protomulti.WatchLobbyRequest, opts ...grpc.CallOption) (*protomulti.WatchLobbyResponse, error) {
	out := new(protomulti.WatchLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_WatchLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) DiscardLobby(ctx context.Context, in *protomulti.DiscardLobbyRequest, opts ...grpc.CallOption) (*protomulti.DiscardLobbyResponse, error) {
	out := new(protomulti.DiscardLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_DiscardLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) CreateLobby(ctx context.Context, in *protomulti.CreateLobbyRequest, opts ...grpc.CallOption) (*protomulti.CreateLobbyResponse, error) {
	out := new(protomulti.CreateLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_CreateLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) JoinLobby(ctx context.Context, in *protomulti.JoinLobbyRequest, opts ...grpc.CallOption) (*protomulti.JoinLobbyResponse, error) {
	out := new(protomulti.JoinLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_JoinLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) ReadyLobby(ctx context.Context, in *protomulti.ReadyLobbyRequest, opts ...grpc.CallOption) (*protomulti.ReadyLobbyResponse, error) {
	out := new(protomulti.ReadyLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_ReadyLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) UnreadyLobby(ctx context.Context, in *protomulti.UnreadyLobbyRequest, opts ...grpc.CallOption) (*protomulti.UnreadyLobbyResponse, error) {
	out := new(protomulti.UnreadyLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_UnreadyLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) CancelLobby(ctx context.Context, in *protomulti.CancelLobbyRequest, opts ...grpc.CallOption) (*protomulti.CancelLobbyResponse, error) {
	out := new(protomulti.CancelLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_CancelLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) GetPlayer(ctx context.Context, in *protomulti.GetLobbyPlayerRequest, opts ...grpc.CallOption) (*protomulti.GetLobbyPlayerResponse, error) {
	out := new(protomulti.GetLobbyPlayerResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_GetPlayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeviusMultiServiceServer is the server API for MeviusMultiService service.
// All implementations should embed UnimplementedMeviusMultiServiceServer
// for forward compatibility
type MeviusMultiServiceServer interface {
	CreateSession(context.Context, *protomulti.CreateSessionRequest) (*protomulti.CreateSessionResponse, error)
	EndSession(context.Context, *protomulti.EndSessionRequest) (*protomulti.EndSessionResponse, error)
	SearchLobby(context.Context, *protomulti.SearchLobbyRequest) (*protomulti.SearchLobbyResponse, error)
	WatchLobby(context.Context, *protomulti.WatchLobbyRequest) (*protomulti.WatchLobbyResponse, error)
	DiscardLobby(context.Context, *protomulti.DiscardLobbyRequest) (*protomulti.DiscardLobbyResponse, error)
	CreateLobby(context.Context, *protomulti.CreateLobbyRequest) (*protomulti.CreateLobbyResponse, error)
	JoinLobby(context.Context, *protomulti.JoinLobbyRequest) (*protomulti.JoinLobbyResponse, error)
	ReadyLobby(context.Context, *protomulti.ReadyLobbyRequest) (*protomulti.ReadyLobbyResponse, error)
	UnreadyLobby(context.Context, *protomulti.UnreadyLobbyRequest) (*protomulti.UnreadyLobbyResponse, error)
	CancelLobby(context.Context, *protomulti.CancelLobbyRequest) (*protomulti.CancelLobbyResponse, error)
	GetPlayer(context.Context, *protomulti.GetLobbyPlayerRequest) (*protomulti.GetLobbyPlayerResponse, error)
}

// UnimplementedMeviusMultiServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMeviusMultiServiceServer struct {
}

func (UnimplementedMeviusMultiServiceServer) CreateSession(context.Context, *protomulti.CreateSessionRequest) (*protomulti.CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedMeviusMultiServiceServer) EndSession(context.Context, *protomulti.EndSessionRequest) (*protomulti.EndSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndSession not implemented")
}
func (UnimplementedMeviusMultiServiceServer) SearchLobby(context.Context, *protomulti.SearchLobbyRequest) (*protomulti.SearchLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) WatchLobby(context.Context, *protomulti.WatchLobbyRequest) (*protomulti.WatchLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) DiscardLobby(context.Context, *protomulti.DiscardLobbyRequest) (*protomulti.DiscardLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) CreateLobby(context.Context, *protomulti.CreateLobbyRequest) (*protomulti.CreateLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) JoinLobby(context.Context, *protomulti.JoinLobbyRequest) (*protomulti.JoinLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) ReadyLobby(context.Context, *protomulti.ReadyLobbyRequest) (*protomulti.ReadyLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadyLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) UnreadyLobby(context.Context, *protomulti.UnreadyLobbyRequest) (*protomulti.UnreadyLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnreadyLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) CancelLobby(context.Context, *protomulti.CancelLobbyRequest) (*protomulti.CancelLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) GetPlayer(context.Context, *protomulti.GetLobbyPlayerRequest) (*protomulti.GetLobbyPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}

// UnsafeMeviusMultiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeviusMultiServiceServer will
// result in compilation errors.
type UnsafeMeviusMultiServiceServer interface {
	mustEmbedUnimplementedMeviusMultiServiceServer()
}

func RegisterMeviusMultiServiceServer(s grpc.ServiceRegistrar, srv MeviusMultiServiceServer) {
	s.RegisterService(&MeviusMultiService_ServiceDesc, srv)
}

func _MeviusMultiService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).CreateSession(ctx, req.(*protomulti.CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_EndSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.EndSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).EndSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_EndSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).EndSession(ctx, req.(*protomulti.EndSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_SearchLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.SearchLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).SearchLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_SearchLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).SearchLobby(ctx, req.(*protomulti.SearchLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_WatchLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.WatchLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).WatchLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_WatchLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).WatchLobby(ctx, req.(*protomulti.WatchLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_DiscardLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.DiscardLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).DiscardLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_DiscardLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).DiscardLobby(ctx, req.(*protomulti.DiscardLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_CreateLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.CreateLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).CreateLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_CreateLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).CreateLobby(ctx, req.(*protomulti.CreateLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_JoinLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.JoinLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).JoinLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_JoinLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).JoinLobby(ctx, req.(*protomulti.JoinLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_ReadyLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.ReadyLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).ReadyLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_ReadyLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).ReadyLobby(ctx, req.(*protomulti.ReadyLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_UnreadyLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.UnreadyLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).UnreadyLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_UnreadyLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).UnreadyLobby(ctx, req.(*protomulti.UnreadyLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_CancelLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.CancelLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).CancelLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_CancelLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).CancelLobby(ctx, req.(*protomulti.CancelLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.GetLobbyPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_GetPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).GetPlayer(ctx, req.(*protomulti.GetLobbyPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeviusMultiService_ServiceDesc is the grpc.ServiceDesc for MeviusMultiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeviusMultiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MeviusMultiService",
	HandlerType: (*MeviusMultiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _MeviusMultiService_CreateSession_Handler,
		},
		{
			MethodName: "EndSession",
			Handler:    _MeviusMultiService_EndSession_Handler,
		},
		{
			MethodName: "SearchLobby",
			Handler:    _MeviusMultiService_SearchLobby_Handler,
		},
		{
			MethodName: "WatchLobby",
			Handler:    _MeviusMultiService_WatchLobby_Handler,
		},
		{
			MethodName: "DiscardLobby",
			Handler:    _MeviusMultiService_DiscardLobby_Handler,
		},
		{
			MethodName: "CreateLobby",
			Handler:    _MeviusMultiService_CreateLobby_Handler,
		},
		{
			MethodName: "JoinLobby",
			Handler:    _MeviusMultiService_JoinLobby_Handler,
		},
		{
			MethodName: "ReadyLobby",
			Handler:    _MeviusMultiService_ReadyLobby_Handler,
		},
		{
			MethodName: "UnreadyLobby",
			Handler:    _MeviusMultiService_UnreadyLobby_Handler,
		},
		{
			MethodName: "CancelLobby",
			Handler:    _MeviusMultiService_CancelLobby_Handler,
		},
		{
			MethodName: "GetPlayer",
			Handler:    _MeviusMultiService_GetPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/multi.service.proto",
}
