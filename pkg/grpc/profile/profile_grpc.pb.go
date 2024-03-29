// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: profile.proto

package profile

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
	Profile_Create_FullMethodName         = "/Profile/Create"
	Profile_Update_FullMethodName         = "/Profile/Update"
	Profile_Get_FullMethodName            = "/Profile/GetRate"
	Profile_GetList_FullMethodName        = "/Profile/GetRateList"
	Profile_AddToList_FullMethodName      = "/Profile/AddToList"
	Profile_RemoveFromList_FullMethodName = "/Profile/RemoveFromList"
	Profile_IsInList_FullMethodName       = "/Profile/IsInList"
	Profile_ListStats_FullMethodName      = "/Profile/ListStats"
)

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	AddToList(ctx context.Context, in *AddToListRequest, opts ...grpc.CallOption) (*share.Empty, error)
	RemoveFromList(ctx context.Context, in *RemoveFromListRequest, opts ...grpc.CallOption) (*share.Empty, error)
	IsInList(ctx context.Context, in *IsInListRequest, opts ...grpc.CallOption) (*IsInListResponse, error)
	ListStats(ctx context.Context, in *ListStatsRequests, opts ...grpc.CallOption) (*ListStatsResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Profile_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Profile_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Profile_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Profile_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) AddToList(ctx context.Context, in *AddToListRequest, opts ...grpc.CallOption) (*share.Empty, error) {
	out := new(share.Empty)
	err := c.cc.Invoke(ctx, Profile_AddToList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) RemoveFromList(ctx context.Context, in *RemoveFromListRequest, opts ...grpc.CallOption) (*share.Empty, error) {
	out := new(share.Empty)
	err := c.cc.Invoke(ctx, Profile_RemoveFromList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) IsInList(ctx context.Context, in *IsInListRequest, opts ...grpc.CallOption) (*IsInListResponse, error) {
	out := new(IsInListResponse)
	err := c.cc.Invoke(ctx, Profile_IsInList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ListStats(ctx context.Context, in *ListStatsRequests, opts ...grpc.CallOption) (*ListStatsResponse, error) {
	out := new(ListStatsResponse)
	err := c.cc.Invoke(ctx, Profile_ListStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
	Update(context.Context, *UpdateRequest) (*Response, error)
	Get(context.Context, *GetRequest) (*Response, error)
	GetList(context.Context, *GetListRequest) (*ListResponse, error)
	AddToList(context.Context, *AddToListRequest) (*share.Empty, error)
	RemoveFromList(context.Context, *RemoveFromListRequest) (*share.Empty, error)
	IsInList(context.Context, *IsInListRequest) (*IsInListResponse, error)
	ListStats(context.Context, *ListStatsRequests) (*ListStatsResponse, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) Create(context.Context, *CreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProfileServer) Update(context.Context, *UpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProfileServer) Get(context.Context, *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (UnimplementedProfileServer) GetList(context.Context, *GetListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRateList not implemented")
}
func (UnimplementedProfileServer) AddToList(context.Context, *AddToListRequest) (*share.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToList not implemented")
}
func (UnimplementedProfileServer) RemoveFromList(context.Context, *RemoveFromListRequest) (*share.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromList not implemented")
}
func (UnimplementedProfileServer) IsInList(context.Context, *IsInListRequest) (*IsInListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInList not implemented")
}
func (UnimplementedProfileServer) ListStats(context.Context, *ListStatsRequests) (*ListStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStats not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_AddToList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddToList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_AddToList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddToList(ctx, req.(*AddToListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_RemoveFromList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).RemoveFromList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_RemoveFromList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).RemoveFromList(ctx, req.(*RemoveFromListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_IsInList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).IsInList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_IsInList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).IsInList(ctx, req.(*IsInListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ListStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatsRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ListStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_ListStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ListStats(ctx, req.(*ListStatsRequests))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Profile_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Profile_Update_Handler,
		},
		{
			MethodName: "GetRate",
			Handler:    _Profile_Get_Handler,
		},
		{
			MethodName: "GetRateList",
			Handler:    _Profile_GetList_Handler,
		},
		{
			MethodName: "AddToList",
			Handler:    _Profile_AddToList_Handler,
		},
		{
			MethodName: "RemoveFromList",
			Handler:    _Profile_RemoveFromList_Handler,
		},
		{
			MethodName: "IsInList",
			Handler:    _Profile_IsInList_Handler,
		},
		{
			MethodName: "ListStats",
			Handler:    _Profile_ListStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
