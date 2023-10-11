// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: service/access.service.proto

package services

import (
	context "context"
	protoaccess "github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AccessService_AuthToken_FullMethodName          = "/service.AccessService/AuthToken"
	AccessService_ChangePassword_FullMethodName     = "/service.AccessService/ChangePassword"
	AccessService_LoginUser_FullMethodName          = "/service.AccessService/LoginUser"
	AccessService_RefreshToken_FullMethodName       = "/service.AccessService/RefreshToken"
	AccessService_RegisterUser_FullMethodName       = "/service.AccessService/RegisterUser"
	AccessService_UserHasRole_FullMethodName        = "/service.AccessService/UserHasRole"
	AccessService_CustomerSearch_FullMethodName     = "/service.AccessService/CustomerSearch"
	AccessService_RememberTokenQuery_FullMethodName = "/service.AccessService/RememberTokenQuery"
)

// AccessServiceClient is the client API for AccessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessServiceClient interface {
	AuthToken(ctx context.Context, in *protoaccess.AuthTokenRequest, opts ...grpc.CallOption) (*protoaccess.AuthTokenResponse, error)
	ChangePassword(ctx context.Context, in *protoaccess.ChangePasswordRequest, opts ...grpc.CallOption) (*protoaccess.ChangePasswordResponse, error)
	LoginUser(ctx context.Context, in *protoaccess.LoginUserRequest, opts ...grpc.CallOption) (*protoaccess.LoginUserResponse, error)
	RefreshToken(ctx context.Context, in *protoaccess.RefreshTokenRequest, opts ...grpc.CallOption) (*protoaccess.RefreshTokenResponse, error)
	RegisterUser(ctx context.Context, in *protoaccess.RegisterUserRequest, opts ...grpc.CallOption) (*protoaccess.RegisterUserResponse, error)
	UserHasRole(ctx context.Context, in *protoaccess.UserHasRoleRequest, opts ...grpc.CallOption) (*protoaccess.UserHasRoleResponse, error)
	CustomerSearch(ctx context.Context, in *protoaccess.CustomerSearchRequest, opts ...grpc.CallOption) (*protoaccess.CustomerSearchResponse, error)
	RememberTokenQuery(ctx context.Context, in *protoaccess.RememberTokenQueryRequest, opts ...grpc.CallOption) (*protoaccess.RememberTokenQueryResponse, error)
}

type accessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessServiceClient(cc grpc.ClientConnInterface) AccessServiceClient {
	return &accessServiceClient{cc}
}

func (c *accessServiceClient) AuthToken(ctx context.Context, in *protoaccess.AuthTokenRequest, opts ...grpc.CallOption) (*protoaccess.AuthTokenResponse, error) {
	out := new(protoaccess.AuthTokenResponse)
	err := c.cc.Invoke(ctx, AccessService_AuthToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) ChangePassword(ctx context.Context, in *protoaccess.ChangePasswordRequest, opts ...grpc.CallOption) (*protoaccess.ChangePasswordResponse, error) {
	out := new(protoaccess.ChangePasswordResponse)
	err := c.cc.Invoke(ctx, AccessService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) LoginUser(ctx context.Context, in *protoaccess.LoginUserRequest, opts ...grpc.CallOption) (*protoaccess.LoginUserResponse, error) {
	out := new(protoaccess.LoginUserResponse)
	err := c.cc.Invoke(ctx, AccessService_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) RefreshToken(ctx context.Context, in *protoaccess.RefreshTokenRequest, opts ...grpc.CallOption) (*protoaccess.RefreshTokenResponse, error) {
	out := new(protoaccess.RefreshTokenResponse)
	err := c.cc.Invoke(ctx, AccessService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) RegisterUser(ctx context.Context, in *protoaccess.RegisterUserRequest, opts ...grpc.CallOption) (*protoaccess.RegisterUserResponse, error) {
	out := new(protoaccess.RegisterUserResponse)
	err := c.cc.Invoke(ctx, AccessService_RegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) UserHasRole(ctx context.Context, in *protoaccess.UserHasRoleRequest, opts ...grpc.CallOption) (*protoaccess.UserHasRoleResponse, error) {
	out := new(protoaccess.UserHasRoleResponse)
	err := c.cc.Invoke(ctx, AccessService_UserHasRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) CustomerSearch(ctx context.Context, in *protoaccess.CustomerSearchRequest, opts ...grpc.CallOption) (*protoaccess.CustomerSearchResponse, error) {
	out := new(protoaccess.CustomerSearchResponse)
	err := c.cc.Invoke(ctx, AccessService_CustomerSearch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServiceClient) RememberTokenQuery(ctx context.Context, in *protoaccess.RememberTokenQueryRequest, opts ...grpc.CallOption) (*protoaccess.RememberTokenQueryResponse, error) {
	out := new(protoaccess.RememberTokenQueryResponse)
	err := c.cc.Invoke(ctx, AccessService_RememberTokenQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessServiceServer is the server API for AccessService service.
// All implementations should embed UnimplementedAccessServiceServer
// for forward compatibility
type AccessServiceServer interface {
	AuthToken(context.Context, *protoaccess.AuthTokenRequest) (*protoaccess.AuthTokenResponse, error)
	ChangePassword(context.Context, *protoaccess.ChangePasswordRequest) (*protoaccess.ChangePasswordResponse, error)
	LoginUser(context.Context, *protoaccess.LoginUserRequest) (*protoaccess.LoginUserResponse, error)
	RefreshToken(context.Context, *protoaccess.RefreshTokenRequest) (*protoaccess.RefreshTokenResponse, error)
	RegisterUser(context.Context, *protoaccess.RegisterUserRequest) (*protoaccess.RegisterUserResponse, error)
	UserHasRole(context.Context, *protoaccess.UserHasRoleRequest) (*protoaccess.UserHasRoleResponse, error)
	CustomerSearch(context.Context, *protoaccess.CustomerSearchRequest) (*protoaccess.CustomerSearchResponse, error)
	RememberTokenQuery(context.Context, *protoaccess.RememberTokenQueryRequest) (*protoaccess.RememberTokenQueryResponse, error)
}

// UnimplementedAccessServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccessServiceServer struct {
}

func (UnimplementedAccessServiceServer) AuthToken(context.Context, *protoaccess.AuthTokenRequest) (*protoaccess.AuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthToken not implemented")
}
func (UnimplementedAccessServiceServer) ChangePassword(context.Context, *protoaccess.ChangePasswordRequest) (*protoaccess.ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAccessServiceServer) LoginUser(context.Context, *protoaccess.LoginUserRequest) (*protoaccess.LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAccessServiceServer) RefreshToken(context.Context, *protoaccess.RefreshTokenRequest) (*protoaccess.RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAccessServiceServer) RegisterUser(context.Context, *protoaccess.RegisterUserRequest) (*protoaccess.RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAccessServiceServer) UserHasRole(context.Context, *protoaccess.UserHasRoleRequest) (*protoaccess.UserHasRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHasRole not implemented")
}
func (UnimplementedAccessServiceServer) CustomerSearch(context.Context, *protoaccess.CustomerSearchRequest) (*protoaccess.CustomerSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerSearch not implemented")
}
func (UnimplementedAccessServiceServer) RememberTokenQuery(context.Context, *protoaccess.RememberTokenQueryRequest) (*protoaccess.RememberTokenQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RememberTokenQuery not implemented")
}

// UnsafeAccessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessServiceServer will
// result in compilation errors.
type UnsafeAccessServiceServer interface {
	mustEmbedUnimplementedAccessServiceServer()
}

func RegisterAccessServiceServer(s grpc.ServiceRegistrar, srv AccessServiceServer) {
	s.RegisterService(&AccessService_ServiceDesc, srv)
}

func _AccessService_AuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.AuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).AuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_AuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).AuthToken(ctx, req.(*protoaccess.AuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).ChangePassword(ctx, req.(*protoaccess.ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).LoginUser(ctx, req.(*protoaccess.LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).RefreshToken(ctx, req.(*protoaccess.RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).RegisterUser(ctx, req.(*protoaccess.RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_UserHasRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.UserHasRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).UserHasRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_UserHasRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).UserHasRole(ctx, req.(*protoaccess.UserHasRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_CustomerSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.CustomerSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).CustomerSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_CustomerSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).CustomerSearch(ctx, req.(*protoaccess.CustomerSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessService_RememberTokenQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protoaccess.RememberTokenQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).RememberTokenQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessService_RememberTokenQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).RememberTokenQuery(ctx, req.(*protoaccess.RememberTokenQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessService_ServiceDesc is the grpc.ServiceDesc for AccessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AccessService",
	HandlerType: (*AccessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthToken",
			Handler:    _AccessService_AuthToken_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AccessService_ChangePassword_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AccessService_LoginUser_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AccessService_RefreshToken_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _AccessService_RegisterUser_Handler,
		},
		{
			MethodName: "UserHasRole",
			Handler:    _AccessService_UserHasRole_Handler,
		},
		{
			MethodName: "CustomerSearch",
			Handler:    _AccessService_CustomerSearch_Handler,
		},
		{
			MethodName: "RememberTokenQuery",
			Handler:    _AccessService_RememberTokenQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/access.service.proto",
}