// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: activity.proto

package activity

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	share "micromango/pkg/grpc/share"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Activity_Like_FullMethodName        = "/micromango.Activity/Like"
	Activity_Dislike_FullMethodName     = "/micromango.Activity/Dislike"
	Activity_LikesNumber_FullMethodName = "/micromango.Activity/LikesNumber"
	Activity_HasLike_FullMethodName     = "/micromango.Activity/HasLike"
)

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*share.Empty, error)
	Dislike(ctx context.Context, in *DislikeRequest, opts ...grpc.CallOption) (*share.Empty, error)
	LikesNumber(ctx context.Context, in *LikesNumberRequest, opts ...grpc.CallOption) (*LikesNumberResponse, error)
	HasLike(ctx context.Context, in *HasLikeRequest, opts ...grpc.CallOption) (*HasLikeResponse, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*share.Empty, error) {
	out := new(share.Empty)
	err := c.cc.Invoke(ctx, Activity_Like_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) Dislike(ctx context.Context, in *DislikeRequest, opts ...grpc.CallOption) (*share.Empty, error) {
	out := new(share.Empty)
	err := c.cc.Invoke(ctx, Activity_Dislike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) LikesNumber(ctx context.Context, in *LikesNumberRequest, opts ...grpc.CallOption) (*LikesNumberResponse, error) {
	out := new(LikesNumberResponse)
	err := c.cc.Invoke(ctx, Activity_LikesNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) HasLike(ctx context.Context, in *HasLikeRequest, opts ...grpc.CallOption) (*HasLikeResponse, error) {
	out := new(HasLikeResponse)
	err := c.cc.Invoke(ctx, Activity_HasLike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	Like(context.Context, *LikeRequest) (*share.Empty, error)
	Dislike(context.Context, *DislikeRequest) (*share.Empty, error)
	LikesNumber(context.Context, *LikesNumberRequest) (*LikesNumberResponse, error)
	HasLike(context.Context, *HasLikeRequest) (*HasLikeResponse, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) Like(context.Context, *LikeRequest) (*share.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedActivityServer) Dislike(context.Context, *DislikeRequest) (*share.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dislike not implemented")
}
func (UnimplementedActivityServer) LikesNumber(context.Context, *LikesNumberRequest) (*LikesNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikesNumber not implemented")
}
func (UnimplementedActivityServer) HasLike(context.Context, *HasLikeRequest) (*HasLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasLike not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_Dislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DislikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).Dislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_Dislike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).Dislike(ctx, req.(*DislikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_LikesNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikesNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).LikesNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_LikesNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).LikesNumber(ctx, req.(*LikesNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_HasLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).HasLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_HasLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).HasLike(ctx, req.(*HasLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "micromango.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Like",
			Handler:    _Activity_Like_Handler,
		},
		{
			MethodName: "Dislike",
			Handler:    _Activity_Dislike_Handler,
		},
		{
			MethodName: "LikesNumber",
			Handler:    _Activity_LikesNumber_Handler,
		},
		{
			MethodName: "HasLike",
			Handler:    _Activity_HasLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity.proto",
}
