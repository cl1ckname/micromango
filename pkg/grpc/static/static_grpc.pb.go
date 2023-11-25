// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: static.proto

package static

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Static_GetImage_FullMethodName    = "/micromango.Static/GetImage"
	Static_UploadCover_FullMethodName = "/micromango.Static/UploadCover"
)

// StaticClient is the client API for Static service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticClient interface {
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	UploadCover(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error)
}

type staticClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticClient(cc grpc.ClientConnInterface) StaticClient {
	return &staticClient{cc}
}

func (c *staticClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, Static_GetImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadCover(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error) {
	out := new(UploadImageResponse)
	err := c.cc.Invoke(ctx, Static_UploadCover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticServer is the server API for Static service.
// All implementations must embed UnimplementedStaticServer
// for forward compatibility
type StaticServer interface {
	GetImage(context.Context, *GetImageRequest) (*ImageResponse, error)
	UploadCover(context.Context, *UploadImageRequest) (*UploadImageResponse, error)
	mustEmbedUnimplementedStaticServer()
}

// UnimplementedStaticServer must be embedded to have forward compatible implementations.
type UnimplementedStaticServer struct {
}

func (UnimplementedStaticServer) GetImage(context.Context, *GetImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedStaticServer) UploadCover(context.Context, *UploadImageRequest) (*UploadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCover not implemented")
}
func (UnimplementedStaticServer) mustEmbedUnimplementedStaticServer() {}

// UnsafeStaticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticServer will
// result in compilation errors.
type UnsafeStaticServer interface {
	mustEmbedUnimplementedStaticServer()
}

func RegisterStaticServer(s grpc.ServiceRegistrar, srv StaticServer) {
	s.RegisterService(&Static_ServiceDesc, srv)
}

func _Static_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_GetImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadCover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadCover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadCover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadCover(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Static_ServiceDesc is the grpc.ServiceDesc for Static service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Static_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "micromango.Static",
	HandlerType: (*StaticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImage",
			Handler:    _Static_GetImage_Handler,
		},
		{
			MethodName: "UploadCover",
			Handler:    _Static_UploadCover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "static.proto",
}
