// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: catalog.proto

package catalog

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
	Catalog_GetManga_FullMethodName    = "/micromango.Catalog/GetManga"
	Catalog_GetMangas_FullMethodName   = "/micromango.Catalog/GetMangas"
	Catalog_AddManga_FullMethodName    = "/micromango.Catalog/AddManga"
	Catalog_UpdateManga_FullMethodName = "/micromango.Catalog/UpdateManga"
)

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	GetManga(ctx context.Context, in *MangaRequest, opts ...grpc.CallOption) (*MangaResponse, error)
	GetMangas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MangasResponse, error)
	AddManga(ctx context.Context, in *AddMangaRequest, opts ...grpc.CallOption) (*MangaResponse, error)
	UpdateManga(ctx context.Context, in *UpdateMangaRequest, opts ...grpc.CallOption) (*MangaResponse, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) GetManga(ctx context.Context, in *MangaRequest, opts ...grpc.CallOption) (*MangaResponse, error) {
	out := new(MangaResponse)
	err := c.cc.Invoke(ctx, Catalog_GetManga_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetMangas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MangasResponse, error) {
	out := new(MangasResponse)
	err := c.cc.Invoke(ctx, Catalog_GetMangas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) AddManga(ctx context.Context, in *AddMangaRequest, opts ...grpc.CallOption) (*MangaResponse, error) {
	out := new(MangaResponse)
	err := c.cc.Invoke(ctx, Catalog_AddManga_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) UpdateManga(ctx context.Context, in *UpdateMangaRequest, opts ...grpc.CallOption) (*MangaResponse, error) {
	out := new(MangaResponse)
	err := c.cc.Invoke(ctx, Catalog_UpdateManga_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	GetManga(context.Context, *MangaRequest) (*MangaResponse, error)
	GetMangas(context.Context, *Empty) (*MangasResponse, error)
	AddManga(context.Context, *AddMangaRequest) (*MangaResponse, error)
	UpdateManga(context.Context, *UpdateMangaRequest) (*MangaResponse, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) GetManga(context.Context, *MangaRequest) (*MangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManga not implemented")
}
func (UnimplementedCatalogServer) GetMangas(context.Context, *Empty) (*MangasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMangas not implemented")
}
func (UnimplementedCatalogServer) AddManga(context.Context, *AddMangaRequest) (*MangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddManga not implemented")
}
func (UnimplementedCatalogServer) UpdateManga(context.Context, *UpdateMangaRequest) (*MangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateManga not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_GetManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_GetManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetManga(ctx, req.(*MangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetMangas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetMangas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_GetMangas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetMangas(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_AddManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).AddManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_AddManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).AddManga(ctx, req.(*AddMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_UpdateManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).UpdateManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Catalog_UpdateManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).UpdateManga(ctx, req.(*UpdateMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "micromango.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManga",
			Handler:    _Catalog_GetManga_Handler,
		},
		{
			MethodName: "GetMangas",
			Handler:    _Catalog_GetMangas_Handler,
		},
		{
			MethodName: "AddManga",
			Handler:    _Catalog_AddManga_Handler,
		},
		{
			MethodName: "UpdateManga",
			Handler:    _Catalog_UpdateManga_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}
