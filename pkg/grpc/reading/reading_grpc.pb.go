// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api/protobuf/reading.proto

package reading

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
	Reading_GetMangaContent_FullMethodName = "/reading.Reading/GetMangaContent"
	Reading_AddMangaContent_FullMethodName = "/reading.Reading/AddMangaContent"
	Reading_GetChapter_FullMethodName      = "/reading.Reading/GetChapter"
	Reading_AddChapter_FullMethodName      = "/reading.Reading/AddChapter"
	Reading_GetPage_FullMethodName         = "/reading.Reading/GetPage"
	Reading_AddPage_FullMethodName         = "/reading.Reading/AddPage"
)

// ReadingClient is the client API for Reading service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReadingClient interface {
	GetMangaContent(ctx context.Context, in *MangaContentRequest, opts ...grpc.CallOption) (*MangaContentResponse, error)
	AddMangaContent(ctx context.Context, in *AddMangaContentRequest, opts ...grpc.CallOption) (*MangaContentResponse, error)
	GetChapter(ctx context.Context, in *ChapterRequest, opts ...grpc.CallOption) (*ChapterResponse, error)
	AddChapter(ctx context.Context, in *AddChapterRequest, opts ...grpc.CallOption) (*ChapterResponse, error)
	GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageResponse, error)
	AddPage(ctx context.Context, in *AddPageRequest, opts ...grpc.CallOption) (*PageResponse, error)
}

type readingClient struct {
	cc grpc.ClientConnInterface
}

func NewReadingClient(cc grpc.ClientConnInterface) ReadingClient {
	return &readingClient{cc}
}

func (c *readingClient) GetMangaContent(ctx context.Context, in *MangaContentRequest, opts ...grpc.CallOption) (*MangaContentResponse, error) {
	out := new(MangaContentResponse)
	err := c.cc.Invoke(ctx, Reading_GetMangaContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readingClient) AddMangaContent(ctx context.Context, in *AddMangaContentRequest, opts ...grpc.CallOption) (*MangaContentResponse, error) {
	out := new(MangaContentResponse)
	err := c.cc.Invoke(ctx, Reading_AddMangaContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readingClient) GetChapter(ctx context.Context, in *ChapterRequest, opts ...grpc.CallOption) (*ChapterResponse, error) {
	out := new(ChapterResponse)
	err := c.cc.Invoke(ctx, Reading_GetChapter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readingClient) AddChapter(ctx context.Context, in *AddChapterRequest, opts ...grpc.CallOption) (*ChapterResponse, error) {
	out := new(ChapterResponse)
	err := c.cc.Invoke(ctx, Reading_AddChapter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readingClient) GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageResponse, error) {
	out := new(PageResponse)
	err := c.cc.Invoke(ctx, Reading_GetPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readingClient) AddPage(ctx context.Context, in *AddPageRequest, opts ...grpc.CallOption) (*PageResponse, error) {
	out := new(PageResponse)
	err := c.cc.Invoke(ctx, Reading_AddPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadingServer is the server API for Reading service.
// All implementations must embed UnimplementedReadingServer
// for forward compatibility
type ReadingServer interface {
	GetMangaContent(context.Context, *MangaContentRequest) (*MangaContentResponse, error)
	AddMangaContent(context.Context, *AddMangaContentRequest) (*MangaContentResponse, error)
	GetChapter(context.Context, *ChapterRequest) (*ChapterResponse, error)
	AddChapter(context.Context, *AddChapterRequest) (*ChapterResponse, error)
	GetPage(context.Context, *PageRequest) (*PageResponse, error)
	AddPage(context.Context, *AddPageRequest) (*PageResponse, error)
	mustEmbedUnimplementedReadingServer()
}

// UnimplementedReadingServer must be embedded to have forward compatible implementations.
type UnimplementedReadingServer struct {
}

func (UnimplementedReadingServer) GetMangaContent(context.Context, *MangaContentRequest) (*MangaContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMangaContent not implemented")
}
func (UnimplementedReadingServer) AddMangaContent(context.Context, *AddMangaContentRequest) (*MangaContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMangaContent not implemented")
}
func (UnimplementedReadingServer) GetChapter(context.Context, *ChapterRequest) (*ChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChapter not implemented")
}
func (UnimplementedReadingServer) AddChapter(context.Context, *AddChapterRequest) (*ChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChapter not implemented")
}
func (UnimplementedReadingServer) GetPage(context.Context, *PageRequest) (*PageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedReadingServer) AddPage(context.Context, *AddPageRequest) (*PageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPage not implemented")
}
func (UnimplementedReadingServer) mustEmbedUnimplementedReadingServer() {}

// UnsafeReadingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReadingServer will
// result in compilation errors.
type UnsafeReadingServer interface {
	mustEmbedUnimplementedReadingServer()
}

func RegisterReadingServer(s grpc.ServiceRegistrar, srv ReadingServer) {
	s.RegisterService(&Reading_ServiceDesc, srv)
}

func _Reading_GetMangaContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MangaContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadingServer).GetMangaContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reading_GetMangaContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadingServer).GetMangaContent(ctx, req.(*MangaContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reading_AddMangaContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMangaContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadingServer).AddMangaContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reading_AddMangaContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadingServer).AddMangaContent(ctx, req.(*AddMangaContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reading_GetChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadingServer).GetChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reading_GetChapter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadingServer).GetChapter(ctx, req.(*ChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reading_AddChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadingServer).AddChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reading_AddChapter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadingServer).AddChapter(ctx, req.(*AddChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reading_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadingServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reading_GetPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadingServer).GetPage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reading_AddPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadingServer).AddPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reading_AddPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadingServer).AddPage(ctx, req.(*AddPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reading_ServiceDesc is the grpc.ServiceDesc for Reading service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reading_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reading.Reading",
	HandlerType: (*ReadingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMangaContent",
			Handler:    _Reading_GetMangaContent_Handler,
		},
		{
			MethodName: "AddMangaContent",
			Handler:    _Reading_AddMangaContent_Handler,
		},
		{
			MethodName: "GetChapter",
			Handler:    _Reading_GetChapter_Handler,
		},
		{
			MethodName: "AddChapter",
			Handler:    _Reading_AddChapter_Handler,
		},
		{
			MethodName: "GetPage",
			Handler:    _Reading_GetPage_Handler,
		},
		{
			MethodName: "AddPage",
			Handler:    _Reading_AddPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf/reading.proto",
}
