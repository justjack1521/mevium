// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: social.service.proto

package services

import (
	context "context"
	protop "github.com/justjack1521/mevium/pkg/genproto/protop"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MeviusSocialService_FollowPlayer_FullMethodName         = "/service.MeviusSocialService/FollowPlayer"
	MeviusSocialService_UnfollowPlayer_FullMethodName       = "/service.MeviusSocialService/UnfollowPlayer"
	MeviusSocialService_PlayerSearch_FullMethodName         = "/service.MeviusSocialService/PlayerSearch"
	MeviusSocialService_UpdatePlayerPresence_FullMethodName = "/service.MeviusSocialService/UpdatePlayerPresence"
	MeviusSocialService_UpdatePlayerPosition_FullMethodName = "/service.MeviusSocialService/UpdatePlayerPosition"
	MeviusSocialService_UpdateCompanion_FullMethodName      = "/service.MeviusSocialService/UpdateCompanion"
)

// MeviusSocialServiceClient is the client API for MeviusSocialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeviusSocialServiceClient interface {
	FollowPlayer(ctx context.Context, in *protop.FollowPlayerRequest, opts ...grpc.CallOption) (*protop.FollowPlayerResponse, error)
	UnfollowPlayer(ctx context.Context, in *protop.UnfollowPlayerRequest, opts ...grpc.CallOption) (*protop.UnfollowPlayerResponse, error)
	PlayerSearch(ctx context.Context, in *protop.PlayerSearchRequest, opts ...grpc.CallOption) (*protop.PlayerSearchResponse, error)
	UpdatePlayerPresence(ctx context.Context, in *protop.UpdatePlayerPresenceRequest, opts ...grpc.CallOption) (*protop.UpdatePlayerPresenceResponse, error)
	UpdatePlayerPosition(ctx context.Context, in *protop.UpdatePlayerPositionRequest, opts ...grpc.CallOption) (*protop.UpdatePlayerPositionResponse, error)
	UpdateCompanion(ctx context.Context, in *protop.UpdatePlayerCompanionRequest, opts ...grpc.CallOption) (*protop.UpdatePlayerCompanionResponse, error)
}

type meviusSocialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeviusSocialServiceClient(cc grpc.ClientConnInterface) MeviusSocialServiceClient {
	return &meviusSocialServiceClient{cc}
}

func (c *meviusSocialServiceClient) FollowPlayer(ctx context.Context, in *protop.FollowPlayerRequest, opts ...grpc.CallOption) (*protop.FollowPlayerResponse, error) {
	out := new(protop.FollowPlayerResponse)
	err := c.cc.Invoke(ctx, MeviusSocialService_FollowPlayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusSocialServiceClient) UnfollowPlayer(ctx context.Context, in *protop.UnfollowPlayerRequest, opts ...grpc.CallOption) (*protop.UnfollowPlayerResponse, error) {
	out := new(protop.UnfollowPlayerResponse)
	err := c.cc.Invoke(ctx, MeviusSocialService_UnfollowPlayer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusSocialServiceClient) PlayerSearch(ctx context.Context, in *protop.PlayerSearchRequest, opts ...grpc.CallOption) (*protop.PlayerSearchResponse, error) {
	out := new(protop.PlayerSearchResponse)
	err := c.cc.Invoke(ctx, MeviusSocialService_PlayerSearch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusSocialServiceClient) UpdatePlayerPresence(ctx context.Context, in *protop.UpdatePlayerPresenceRequest, opts ...grpc.CallOption) (*protop.UpdatePlayerPresenceResponse, error) {
	out := new(protop.UpdatePlayerPresenceResponse)
	err := c.cc.Invoke(ctx, MeviusSocialService_UpdatePlayerPresence_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusSocialServiceClient) UpdatePlayerPosition(ctx context.Context, in *protop.UpdatePlayerPositionRequest, opts ...grpc.CallOption) (*protop.UpdatePlayerPositionResponse, error) {
	out := new(protop.UpdatePlayerPositionResponse)
	err := c.cc.Invoke(ctx, MeviusSocialService_UpdatePlayerPosition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meviusSocialServiceClient) UpdateCompanion(ctx context.Context, in *protop.UpdatePlayerCompanionRequest, opts ...grpc.CallOption) (*protop.UpdatePlayerCompanionResponse, error) {
	out := new(protop.UpdatePlayerCompanionResponse)
	err := c.cc.Invoke(ctx, MeviusSocialService_UpdateCompanion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeviusSocialServiceServer is the server API for MeviusSocialService service.
// All implementations should embed UnimplementedMeviusSocialServiceServer
// for forward compatibility
type MeviusSocialServiceServer interface {
	FollowPlayer(context.Context, *protop.FollowPlayerRequest) (*protop.FollowPlayerResponse, error)
	UnfollowPlayer(context.Context, *protop.UnfollowPlayerRequest) (*protop.UnfollowPlayerResponse, error)
	PlayerSearch(context.Context, *protop.PlayerSearchRequest) (*protop.PlayerSearchResponse, error)
	UpdatePlayerPresence(context.Context, *protop.UpdatePlayerPresenceRequest) (*protop.UpdatePlayerPresenceResponse, error)
	UpdatePlayerPosition(context.Context, *protop.UpdatePlayerPositionRequest) (*protop.UpdatePlayerPositionResponse, error)
	UpdateCompanion(context.Context, *protop.UpdatePlayerCompanionRequest) (*protop.UpdatePlayerCompanionResponse, error)
}

// UnimplementedMeviusSocialServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMeviusSocialServiceServer struct {
}

func (UnimplementedMeviusSocialServiceServer) FollowPlayer(context.Context, *protop.FollowPlayerRequest) (*protop.FollowPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowPlayer not implemented")
}
func (UnimplementedMeviusSocialServiceServer) UnfollowPlayer(context.Context, *protop.UnfollowPlayerRequest) (*protop.UnfollowPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowPlayer not implemented")
}
func (UnimplementedMeviusSocialServiceServer) PlayerSearch(context.Context, *protop.PlayerSearchRequest) (*protop.PlayerSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerSearch not implemented")
}
func (UnimplementedMeviusSocialServiceServer) UpdatePlayerPresence(context.Context, *protop.UpdatePlayerPresenceRequest) (*protop.UpdatePlayerPresenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayerPresence not implemented")
}
func (UnimplementedMeviusSocialServiceServer) UpdatePlayerPosition(context.Context, *protop.UpdatePlayerPositionRequest) (*protop.UpdatePlayerPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayerPosition not implemented")
}
func (UnimplementedMeviusSocialServiceServer) UpdateCompanion(context.Context, *protop.UpdatePlayerCompanionRequest) (*protop.UpdatePlayerCompanionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanion not implemented")
}

// UnsafeMeviusSocialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeviusSocialServiceServer will
// result in compilation errors.
type UnsafeMeviusSocialServiceServer interface {
	mustEmbedUnimplementedMeviusSocialServiceServer()
}

func RegisterMeviusSocialServiceServer(s grpc.ServiceRegistrar, srv MeviusSocialServiceServer) {
	s.RegisterService(&MeviusSocialService_ServiceDesc, srv)
}

func _MeviusSocialService_FollowPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protop.FollowPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusSocialServiceServer).FollowPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusSocialService_FollowPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusSocialServiceServer).FollowPlayer(ctx, req.(*protop.FollowPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusSocialService_UnfollowPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protop.UnfollowPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusSocialServiceServer).UnfollowPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusSocialService_UnfollowPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusSocialServiceServer).UnfollowPlayer(ctx, req.(*protop.UnfollowPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusSocialService_PlayerSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protop.PlayerSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusSocialServiceServer).PlayerSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusSocialService_PlayerSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusSocialServiceServer).PlayerSearch(ctx, req.(*protop.PlayerSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusSocialService_UpdatePlayerPresence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protop.UpdatePlayerPresenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusSocialServiceServer).UpdatePlayerPresence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusSocialService_UpdatePlayerPresence_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusSocialServiceServer).UpdatePlayerPresence(ctx, req.(*protop.UpdatePlayerPresenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusSocialService_UpdatePlayerPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protop.UpdatePlayerPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusSocialServiceServer).UpdatePlayerPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusSocialService_UpdatePlayerPosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusSocialServiceServer).UpdatePlayerPosition(ctx, req.(*protop.UpdatePlayerPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeviusSocialService_UpdateCompanion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protop.UpdatePlayerCompanionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusSocialServiceServer).UpdateCompanion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusSocialService_UpdateCompanion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusSocialServiceServer).UpdateCompanion(ctx, req.(*protop.UpdatePlayerCompanionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeviusSocialService_ServiceDesc is the grpc.ServiceDesc for MeviusSocialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeviusSocialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MeviusSocialService",
	HandlerType: (*MeviusSocialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowPlayer",
			Handler:    _MeviusSocialService_FollowPlayer_Handler,
		},
		{
			MethodName: "UnfollowPlayer",
			Handler:    _MeviusSocialService_UnfollowPlayer_Handler,
		},
		{
			MethodName: "PlayerSearch",
			Handler:    _MeviusSocialService_PlayerSearch_Handler,
		},
		{
			MethodName: "UpdatePlayerPresence",
			Handler:    _MeviusSocialService_UpdatePlayerPresence_Handler,
		},
		{
			MethodName: "UpdatePlayerPosition",
			Handler:    _MeviusSocialService_UpdatePlayerPosition_Handler,
		},
		{
			MethodName: "UpdateCompanion",
			Handler:    _MeviusSocialService_UpdateCompanion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "social.service.proto",
}
