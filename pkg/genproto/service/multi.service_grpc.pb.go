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
	MeviusMultiService_UnwatchLobby_FullMethodName  = "/service.MeviusMultiService/UnwatchLobby"
	MeviusMultiService_CreateLobby_FullMethodName   = "/service.MeviusMultiService/CreateLobby"
	MeviusMultiService_CancelLobby_FullMethodName   = "/service.MeviusMultiService/CancelLobby"
	MeviusMultiService_StartLobby_FullMethodName    = "/service.MeviusMultiService/StartLobby"
	MeviusMultiService_JoinLobby_FullMethodName     = "/service.MeviusMultiService/JoinLobby"
	MeviusMultiService_LeaveLobby_FullMethodName    = "/service.MeviusMultiService/LeaveLobby"
	MeviusMultiService_ReadyLobby_FullMethodName    = "/service.MeviusMultiService/ReadyLobby"
	MeviusMultiService_UnreadyLobby_FullMethodName  = "/service.MeviusMultiService/UnreadyLobby"
	MeviusMultiService_SendStamp_FullMethodName     = "/service.MeviusMultiService/SendStamp"
	MeviusMultiService_GetGame_FullMethodName       = "/service.MeviusMultiService/GetGame"
	MeviusMultiService_EnqueueAction_FullMethodName = "/service.MeviusMultiService/EnqueueAction"
	MeviusMultiService_DequeueAction_FullMethodName = "/service.MeviusMultiService/DequeueAction"
	MeviusMultiService_LockAction_FullMethodName    = "/service.MeviusMultiService/LockAction"
)

// MeviusMultiServiceClient is the client API for MeviusMultiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeviusMultiServiceClient interface {
	CreateSession(ctx context.Context, in *protomulti.CreateSessionRequest, opts ...grpc.CallOption) (*protomulti.CreateSessionResponse, error)
	EndSession(ctx context.Context, in *protomulti.EndSessionRequest, opts ...grpc.CallOption) (*protomulti.EndSessionResponse, error)
	SearchLobby(ctx context.Context, in *protomulti.SearchLobbyRequest, opts ...grpc.CallOption) (*protomulti.SearchLobbyResponse, error)
	WatchLobby(ctx context.Context, in *protomulti.WatchLobbyRequest, opts ...grpc.CallOption) (*protomulti.WatchLobbyResponse, error)
	UnwatchLobby(ctx context.Context, in *protomulti.UnwatchLobbyRequest, opts ...grpc.CallOption) (*protomulti.UnwatchLobbyResponse, error)
	CreateLobby(ctx context.Context, in *protomulti.CreateLobbyRequest, opts ...grpc.CallOption) (*protomulti.CreateLobbyResponse, error)
	CancelLobby(ctx context.Context, in *protomulti.CancelLobbyRequest, opts ...grpc.CallOption) (*protomulti.CancelLobbyResponse, error)
	StartLobby(ctx context.Context, in *protomulti.StartLobbyRequest, opts ...grpc.CallOption) (*protomulti.StartLobbyResponse, error)
	JoinLobby(ctx context.Context, in *protomulti.JoinLobbyRequest, opts ...grpc.CallOption) (*protomulti.JoinLobbyResponse, error)
	LeaveLobby(ctx context.Context, in *protomulti.LeaveLobbyRequest, opts ...grpc.CallOption) (*protomulti.LeaveLobbyResponse, error)
	ReadyLobby(ctx context.Context, in *protomulti.ReadyLobbyRequest, opts ...grpc.CallOption) (*protomulti.ReadyLobbyResponse, error)
	UnreadyLobby(ctx context.Context, in *protomulti.UnreadyLobbyRequest, opts ...grpc.CallOption) (*protomulti.UnreadyLobbyResponse, error)
	SendStamp(ctx context.Context, in *protomulti.SendStampRequest, opts ...grpc.CallOption) (*protomulti.SendStampResponse, error)
	GetGame(ctx context.Context, in *protomulti.GetGameRequest, opts ...grpc.CallOption) (*protomulti.GetGameResponse, error)
	EnqueueAction(ctx context.Context, in *protomulti.GameEnqueueActionRequest, opts ...grpc.CallOption) (*protomulti.GameEnqueueActionResponse, error)
	DequeueAction(ctx context.Context, in *protomulti.GameDequeueActionRequest, opts ...grpc.CallOption) (*protomulti.GameDequeueActionResponse, error)
	LockAction(ctx context.Context, in *protomulti.GameLockActionRequest, opts ...grpc.CallOption) (*protomulti.GameLockActionResponse, error)
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

func (c *meviusMultiServiceClient) UnwatchLobby(ctx context.Context, in *protomulti.UnwatchLobbyRequest, opts ...grpc.CallOption) (*protomulti.UnwatchLobbyResponse, error) {
	out := new(protomulti.UnwatchLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_UnwatchLobby_FullMethodName, in, out, opts...)
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

func (c *meviusMultiServiceClient) CancelLobby(ctx context.Context, in *protomulti.CancelLobbyRequest, opts ...grpc.CallOption) (*protomulti.CancelLobbyResponse, error) {
	out := new(protomulti.CancelLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_CancelLobby_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) StartLobby(ctx context.Context, in *protomulti.StartLobbyRequest, opts ...grpc.CallOption) (*protomulti.StartLobbyResponse, error) {
	out := new(protomulti.StartLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_StartLobby_FullMethodName, in, out, opts...)
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

func (c *meviusMultiServiceClient) LeaveLobby(ctx context.Context, in *protomulti.LeaveLobbyRequest, opts ...grpc.CallOption) (*protomulti.LeaveLobbyResponse, error) {
	out := new(protomulti.LeaveLobbyResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_LeaveLobby_FullMethodName, in, out, opts...)
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

func (c *meviusMultiServiceClient) SendStamp(ctx context.Context, in *protomulti.SendStampRequest, opts ...grpc.CallOption) (*protomulti.SendStampResponse, error) {
	out := new(protomulti.SendStampResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_SendStamp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) GetGame(ctx context.Context, in *protomulti.GetGameRequest, opts ...grpc.CallOption) (*protomulti.GetGameResponse, error) {
	out := new(protomulti.GetGameResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_GetGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) EnqueueAction(ctx context.Context, in *protomulti.GameEnqueueActionRequest, opts ...grpc.CallOption) (*protomulti.GameEnqueueActionResponse, error) {
	out := new(protomulti.GameEnqueueActionResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_EnqueueAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) DequeueAction(ctx context.Context, in *protomulti.GameDequeueActionRequest, opts ...grpc.CallOption) (*protomulti.GameDequeueActionResponse, error) {
	out := new(protomulti.GameDequeueActionResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_DequeueAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusMultiServiceClient) LockAction(ctx context.Context, in *protomulti.GameLockActionRequest, opts ...grpc.CallOption) (*protomulti.GameLockActionResponse, error) {
	out := new(protomulti.GameLockActionResponse)
	err := c.cc.Invoke(ctx, MeviusMultiService_LockAction_FullMethodName, in, out, opts...)
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
	UnwatchLobby(context.Context, *protomulti.UnwatchLobbyRequest) (*protomulti.UnwatchLobbyResponse, error)
	CreateLobby(context.Context, *protomulti.CreateLobbyRequest) (*protomulti.CreateLobbyResponse, error)
	CancelLobby(context.Context, *protomulti.CancelLobbyRequest) (*protomulti.CancelLobbyResponse, error)
	StartLobby(context.Context, *protomulti.StartLobbyRequest) (*protomulti.StartLobbyResponse, error)
	JoinLobby(context.Context, *protomulti.JoinLobbyRequest) (*protomulti.JoinLobbyResponse, error)
	LeaveLobby(context.Context, *protomulti.LeaveLobbyRequest) (*protomulti.LeaveLobbyResponse, error)
	ReadyLobby(context.Context, *protomulti.ReadyLobbyRequest) (*protomulti.ReadyLobbyResponse, error)
	UnreadyLobby(context.Context, *protomulti.UnreadyLobbyRequest) (*protomulti.UnreadyLobbyResponse, error)
	SendStamp(context.Context, *protomulti.SendStampRequest) (*protomulti.SendStampResponse, error)
	GetGame(context.Context, *protomulti.GetGameRequest) (*protomulti.GetGameResponse, error)
	EnqueueAction(context.Context, *protomulti.GameEnqueueActionRequest) (*protomulti.GameEnqueueActionResponse, error)
	DequeueAction(context.Context, *protomulti.GameDequeueActionRequest) (*protomulti.GameDequeueActionResponse, error)
	LockAction(context.Context, *protomulti.GameLockActionRequest) (*protomulti.GameLockActionResponse, error)
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
func (UnimplementedMeviusMultiServiceServer) UnwatchLobby(context.Context, *protomulti.UnwatchLobbyRequest) (*protomulti.UnwatchLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnwatchLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) CreateLobby(context.Context, *protomulti.CreateLobbyRequest) (*protomulti.CreateLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) CancelLobby(context.Context, *protomulti.CancelLobbyRequest) (*protomulti.CancelLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) StartLobby(context.Context, *protomulti.StartLobbyRequest) (*protomulti.StartLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) JoinLobby(context.Context, *protomulti.JoinLobbyRequest) (*protomulti.JoinLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) LeaveLobby(context.Context, *protomulti.LeaveLobbyRequest) (*protomulti.LeaveLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) ReadyLobby(context.Context, *protomulti.ReadyLobbyRequest) (*protomulti.ReadyLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadyLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) UnreadyLobby(context.Context, *protomulti.UnreadyLobbyRequest) (*protomulti.UnreadyLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnreadyLobby not implemented")
}
func (UnimplementedMeviusMultiServiceServer) SendStamp(context.Context, *protomulti.SendStampRequest) (*protomulti.SendStampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendStamp not implemented")
}
func (UnimplementedMeviusMultiServiceServer) GetGame(context.Context, *protomulti.GetGameRequest) (*protomulti.GetGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedMeviusMultiServiceServer) EnqueueAction(context.Context, *protomulti.GameEnqueueActionRequest) (*protomulti.GameEnqueueActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueAction not implemented")
}
func (UnimplementedMeviusMultiServiceServer) DequeueAction(context.Context, *protomulti.GameDequeueActionRequest) (*protomulti.GameDequeueActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DequeueAction not implemented")
}
func (UnimplementedMeviusMultiServiceServer) LockAction(context.Context, *protomulti.GameLockActionRequest) (*protomulti.GameLockActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockAction not implemented")
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

func _MeviusMultiService_UnwatchLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.UnwatchLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).UnwatchLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_UnwatchLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).UnwatchLobby(ctx, req.(*protomulti.UnwatchLobbyRequest))
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

func _MeviusMultiService_StartLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.StartLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).StartLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_StartLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).StartLobby(ctx, req.(*protomulti.StartLobbyRequest))
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

func _MeviusMultiService_LeaveLobby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.LeaveLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).LeaveLobby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_LeaveLobby_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).LeaveLobby(ctx, req.(*protomulti.LeaveLobbyRequest))
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

func _MeviusMultiService_SendStamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.SendStampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).SendStamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_SendStamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).SendStamp(ctx, req.(*protomulti.SendStampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.GetGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_GetGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).GetGame(ctx, req.(*protomulti.GetGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_EnqueueAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.GameEnqueueActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).EnqueueAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_EnqueueAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).EnqueueAction(ctx, req.(*protomulti.GameEnqueueActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_DequeueAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.GameDequeueActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).DequeueAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_DequeueAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).DequeueAction(ctx, req.(*protomulti.GameDequeueActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusMultiService_LockAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protomulti.GameLockActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusMultiServiceServer).LockAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusMultiService_LockAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusMultiServiceServer).LockAction(ctx, req.(*protomulti.GameLockActionRequest))
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
			MethodName: "UnwatchLobby",
			Handler:    _MeviusMultiService_UnwatchLobby_Handler,
		},
		{
			MethodName: "CreateLobby",
			Handler:    _MeviusMultiService_CreateLobby_Handler,
		},
		{
			MethodName: "CancelLobby",
			Handler:    _MeviusMultiService_CancelLobby_Handler,
		},
		{
			MethodName: "StartLobby",
			Handler:    _MeviusMultiService_StartLobby_Handler,
		},
		{
			MethodName: "JoinLobby",
			Handler:    _MeviusMultiService_JoinLobby_Handler,
		},
		{
			MethodName: "LeaveLobby",
			Handler:    _MeviusMultiService_LeaveLobby_Handler,
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
			MethodName: "SendStamp",
			Handler:    _MeviusMultiService_SendStamp_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _MeviusMultiService_GetGame_Handler,
		},
		{
			MethodName: "EnqueueAction",
			Handler:    _MeviusMultiService_EnqueueAction_Handler,
		},
		{
			MethodName: "DequeueAction",
			Handler:    _MeviusMultiService_DequeueAction_Handler,
		},
		{
			MethodName: "LockAction",
			Handler:    _MeviusMultiService_LockAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/multi.service.proto",
}
