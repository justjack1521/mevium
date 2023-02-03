// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: web.service.proto

package services

import (
	context "context"
	protoc "github.com/justjack1521/mevium/pkg/genproto/protoc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoreWebServiceClient is the client API for CoreWebService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreWebServiceClient interface {
	CardSale(ctx context.Context, in *protoc.CardSaleRequest, opts ...grpc.CallOption) (*protoc.CardSaleResponse, error)
	CardFusion(ctx context.Context, in *protoc.CardFusionRequest, opts ...grpc.CallOption) (*protoc.CardFusionResponse, error)
	CardBoostFusion(ctx context.Context, in *protoc.CardBoostFusionRequest, opts ...grpc.CallOption) (*protoc.CardBoostFusionResponse, error)
	CreateProfile(ctx context.Context, in *protoc.CreateProfileRequest, opts ...grpc.CallOption) (*protoc.CreateProfileResponse, error)
	ClaimMailboxItem(ctx context.Context, in *protoc.ClaimMailBoxItemRequest, opts ...grpc.CallOption) (*protoc.ClaimMailBoxItemResponse, error)
	ConfirmDailyMission(ctx context.Context, in *protoc.ConfirmDailyMissionRequest, opts ...grpc.CallOption) (*protoc.ConfirmDailyMissionResponse, error)
	FavouriteCard(ctx context.Context, in *protoc.CardFavouriteRequest, opts ...grpc.CallOption) (*protoc.CardFavouriteResponse, error)
	FollowPlayer(ctx context.Context, in *protoc.FollowPlayerRequest, opts ...grpc.CallOption) (*protoc.FollowPlayerResponse, error)
	RestoreStamina(ctx context.Context, in *protoc.StaminaRestoreRequest, opts ...grpc.CallOption) (*protoc.StaminaRestoreResponse, error)
	Teleport(ctx context.Context, in *protoc.TeleportRequest, opts ...grpc.CallOption) (*protoc.TeleportResponse, error)
	UnfollowPlayer(ctx context.Context, in *protoc.UnfollowPlayerRequest, opts ...grpc.CallOption) (*protoc.UnfollowPlayerResponse, error)
}

type coreWebServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreWebServiceClient(cc grpc.ClientConnInterface) CoreWebServiceClient {
	return &coreWebServiceClient{cc}
}

func (c *coreWebServiceClient) CardSale(ctx context.Context, in *protoc.CardSaleRequest, opts ...grpc.CallOption) (*protoc.CardSaleResponse, error) {
	out := new(protoc.CardSaleResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/CardSale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) CardFusion(ctx context.Context, in *protoc.CardFusionRequest, opts ...grpc.CallOption) (*protoc.CardFusionResponse, error) {
	out := new(protoc.CardFusionResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/CardFusion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) CardBoostFusion(ctx context.Context, in *protoc.CardBoostFusionRequest, opts ...grpc.CallOption) (*protoc.CardBoostFusionResponse, error) {
	out := new(protoc.CardBoostFusionResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/CardBoostFusion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) CreateProfile(ctx context.Context, in *protoc.CreateProfileRequest, opts ...grpc.CallOption) (*protoc.CreateProfileResponse, error) {
	out := new(protoc.CreateProfileResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) ClaimMailboxItem(ctx context.Context, in *protoc.ClaimMailBoxItemRequest, opts ...grpc.CallOption) (*protoc.ClaimMailBoxItemResponse, error) {
	out := new(protoc.ClaimMailBoxItemResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/ClaimMailboxItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) ConfirmDailyMission(ctx context.Context, in *protoc.ConfirmDailyMissionRequest, opts ...grpc.CallOption) (*protoc.ConfirmDailyMissionResponse, error) {
	out := new(protoc.ConfirmDailyMissionResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/ConfirmDailyMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) FavouriteCard(ctx context.Context, in *protoc.CardFavouriteRequest, opts ...grpc.CallOption) (*protoc.CardFavouriteResponse, error) {
	out := new(protoc.CardFavouriteResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/FavouriteCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) FollowPlayer(ctx context.Context, in *protoc.FollowPlayerRequest, opts ...grpc.CallOption) (*protoc.FollowPlayerResponse, error) {
	out := new(protoc.FollowPlayerResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/FollowPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) RestoreStamina(ctx context.Context, in *protoc.StaminaRestoreRequest, opts ...grpc.CallOption) (*protoc.StaminaRestoreResponse, error) {
	out := new(protoc.StaminaRestoreResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/RestoreStamina", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) Teleport(ctx context.Context, in *protoc.TeleportRequest, opts ...grpc.CallOption) (*protoc.TeleportResponse, error) {
	out := new(protoc.TeleportResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/Teleport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreWebServiceClient) UnfollowPlayer(ctx context.Context, in *protoc.UnfollowPlayerRequest, opts ...grpc.CallOption) (*protoc.UnfollowPlayerResponse, error) {
	out := new(protoc.UnfollowPlayerResponse)
	err := c.cc.Invoke(ctx, "/core.CoreWebService/UnfollowPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreWebServiceServer is the server API for CoreWebService service.
// All implementations should embed UnimplementedCoreWebServiceServer
// for forward compatibility
type CoreWebServiceServer interface {
	CardSale(context.Context, *protoc.CardSaleRequest) (*protoc.CardSaleResponse, error)
	CardFusion(context.Context, *protoc.CardFusionRequest) (*protoc.CardFusionResponse, error)
	CardBoostFusion(context.Context, *protoc.CardBoostFusionRequest) (*protoc.CardBoostFusionResponse, error)
	CreateProfile(context.Context, *protoc.CreateProfileRequest) (*protoc.CreateProfileResponse, error)
	ClaimMailboxItem(context.Context, *protoc.ClaimMailBoxItemRequest) (*protoc.ClaimMailBoxItemResponse, error)
	ConfirmDailyMission(context.Context, *protoc.ConfirmDailyMissionRequest) (*protoc.ConfirmDailyMissionResponse, error)
	FavouriteCard(context.Context, *protoc.CardFavouriteRequest) (*protoc.CardFavouriteResponse, error)
	FollowPlayer(context.Context, *protoc.FollowPlayerRequest) (*protoc.FollowPlayerResponse, error)
	RestoreStamina(context.Context, *protoc.StaminaRestoreRequest) (*protoc.StaminaRestoreResponse, error)
	Teleport(context.Context, *protoc.TeleportRequest) (*protoc.TeleportResponse, error)
	UnfollowPlayer(context.Context, *protoc.UnfollowPlayerRequest) (*protoc.UnfollowPlayerResponse, error)
}

// UnimplementedCoreWebServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCoreWebServiceServer struct {
}

func (UnimplementedCoreWebServiceServer) CardSale(context.Context, *protoc.CardSaleRequest) (*protoc.CardSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardSale not implemented")
}
func (UnimplementedCoreWebServiceServer) CardFusion(context.Context, *protoc.CardFusionRequest) (*protoc.CardFusionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardFusion not implemented")
}
func (UnimplementedCoreWebServiceServer) CardBoostFusion(context.Context, *protoc.CardBoostFusionRequest) (*protoc.CardBoostFusionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardBoostFusion not implemented")
}
func (UnimplementedCoreWebServiceServer) CreateProfile(context.Context, *protoc.CreateProfileRequest) (*protoc.CreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedCoreWebServiceServer) ClaimMailboxItem(context.Context, *protoc.ClaimMailBoxItemRequest) (*protoc.ClaimMailBoxItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimMailboxItem not implemented")
}
func (UnimplementedCoreWebServiceServer) ConfirmDailyMission(context.Context, *protoc.ConfirmDailyMissionRequest) (*protoc.ConfirmDailyMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDailyMission not implemented")
}
func (UnimplementedCoreWebServiceServer) FavouriteCard(context.Context, *protoc.CardFavouriteRequest) (*protoc.CardFavouriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavouriteCard not implemented")
}
func (UnimplementedCoreWebServiceServer) FollowPlayer(context.Context, *protoc.FollowPlayerRequest) (*protoc.FollowPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowPlayer not implemented")
}
func (UnimplementedCoreWebServiceServer) RestoreStamina(context.Context, *protoc.StaminaRestoreRequest) (*protoc.StaminaRestoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreStamina not implemented")
}
func (UnimplementedCoreWebServiceServer) Teleport(context.Context, *protoc.TeleportRequest) (*protoc.TeleportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teleport not implemented")
}
func (UnimplementedCoreWebServiceServer) UnfollowPlayer(context.Context, *protoc.UnfollowPlayerRequest) (*protoc.UnfollowPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowPlayer not implemented")
}

// UnsafeCoreWebServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreWebServiceServer will
// result in compilation errors.
type UnsafeCoreWebServiceServer interface {
	mustEmbedUnimplementedCoreWebServiceServer()
}

func RegisterCoreWebServiceServer(s grpc.ServiceRegistrar, srv CoreWebServiceServer) {
	s.RegisterService(&CoreWebService_ServiceDesc, srv)
}

func _CoreWebService_CardSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.CardSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).CardSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/CardSale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).CardSale(ctx, req.(*protoc.CardSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_CardFusion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.CardFusionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).CardFusion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/CardFusion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).CardFusion(ctx, req.(*protoc.CardFusionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_CardBoostFusion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.CardBoostFusionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).CardBoostFusion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/CardBoostFusion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).CardBoostFusion(ctx, req.(*protoc.CardBoostFusionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).CreateProfile(ctx, req.(*protoc.CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_ClaimMailboxItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.ClaimMailBoxItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).ClaimMailboxItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/ClaimMailboxItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).ClaimMailboxItem(ctx, req.(*protoc.ClaimMailBoxItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_ConfirmDailyMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.ConfirmDailyMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).ConfirmDailyMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/ConfirmDailyMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).ConfirmDailyMission(ctx, req.(*protoc.ConfirmDailyMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_FavouriteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.CardFavouriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).FavouriteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/FavouriteCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).FavouriteCard(ctx, req.(*protoc.CardFavouriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_FollowPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.FollowPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).FollowPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/FollowPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).FollowPlayer(ctx, req.(*protoc.FollowPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_RestoreStamina_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.StaminaRestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).RestoreStamina(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/RestoreStamina",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).RestoreStamina(ctx, req.(*protoc.StaminaRestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_Teleport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.TeleportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).Teleport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/Teleport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).Teleport(ctx, req.(*protoc.TeleportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreWebService_UnfollowPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoc.UnfollowPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreWebServiceServer).UnfollowPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.CoreWebService/UnfollowPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreWebServiceServer).UnfollowPlayer(ctx, req.(*protoc.UnfollowPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreWebService_ServiceDesc is the grpc.ServiceDesc for CoreWebService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreWebService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.CoreWebService",
	HandlerType: (*CoreWebServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CardSale",
			Handler:    _CoreWebService_CardSale_Handler,
		},
		{
			MethodName: "CardFusion",
			Handler:    _CoreWebService_CardFusion_Handler,
		},
		{
			MethodName: "CardBoostFusion",
			Handler:    _CoreWebService_CardBoostFusion_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _CoreWebService_CreateProfile_Handler,
		},
		{
			MethodName: "ClaimMailboxItem",
			Handler:    _CoreWebService_ClaimMailboxItem_Handler,
		},
		{
			MethodName: "ConfirmDailyMission",
			Handler:    _CoreWebService_ConfirmDailyMission_Handler,
		},
		{
			MethodName: "FavouriteCard",
			Handler:    _CoreWebService_FavouriteCard_Handler,
		},
		{
			MethodName: "FollowPlayer",
			Handler:    _CoreWebService_FollowPlayer_Handler,
		},
		{
			MethodName: "RestoreStamina",
			Handler:    _CoreWebService_RestoreStamina_Handler,
		},
		{
			MethodName: "Teleport",
			Handler:    _CoreWebService_Teleport_Handler,
		},
		{
			MethodName: "UnfollowPlayer",
			Handler:    _CoreWebService_UnfollowPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web.service.proto",
}
