// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: rank.service.proto

package services

import (
	context "context"
	protor "github.com/justjack1521/mevium/pkg/genproto/protor"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MeviusRankService_SubmitScore_FullMethodName = "/service.MeviusRankService/SubmitScore"
)

// MeviusRankServiceClient is the client API for MeviusRankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeviusRankServiceClient interface {
	SubmitScore(ctx context.Context, in *protor.SubmitScoreRequest, opts ...grpc.CallOption) (*protor.SubmitScoreResponse, error)
}

type meviusRankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeviusRankServiceClient(cc grpc.ClientConnInterface) MeviusRankServiceClient {
	return &meviusRankServiceClient{cc}
}

func (c *meviusRankServiceClient) SubmitScore(ctx context.Context, in *protor.SubmitScoreRequest, opts ...grpc.CallOption) (*protor.SubmitScoreResponse, error) {
	out := new(protor.SubmitScoreResponse)
	err := c.cc.Invoke(ctx, MeviusRankService_SubmitScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeviusRankServiceServer is the server API for MeviusRankService service.
// All implementations should embed UnimplementedMeviusRankServiceServer
// for forward compatibility
type MeviusRankServiceServer interface {
	SubmitScore(context.Context, *protor.SubmitScoreRequest) (*protor.SubmitScoreResponse, error)
}

// UnimplementedMeviusRankServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMeviusRankServiceServer struct {
}

func (UnimplementedMeviusRankServiceServer) SubmitScore(context.Context, *protor.SubmitScoreRequest) (*protor.SubmitScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScore not implemented")
}

// UnsafeMeviusRankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeviusRankServiceServer will
// result in compilation errors.
type UnsafeMeviusRankServiceServer interface {
	mustEmbedUnimplementedMeviusRankServiceServer()
}

func RegisterMeviusRankServiceServer(s grpc.ServiceRegistrar, srv MeviusRankServiceServer) {
	s.RegisterService(&MeviusRankService_ServiceDesc, srv)
}

func _MeviusRankService_SubmitScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protor.SubmitScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeviusRankServiceServer).SubmitScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeviusRankService_SubmitScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeviusRankServiceServer).SubmitScore(ctx, req.(*protor.SubmitScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeviusRankService_ServiceDesc is the grpc.ServiceDesc for MeviusRankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeviusRankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MeviusRankService",
	HandlerType: (*MeviusRankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitScore",
			Handler:    _MeviusRankService_SubmitScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rank.service.proto",
}
