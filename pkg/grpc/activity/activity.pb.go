// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: activity.proto

package activity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	share "micromango/pkg/grpc/share"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0}
}

func (x *LikeRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *LikeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DislikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DislikeRequest) Reset() {
	*x = DislikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DislikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DislikeRequest) ProtoMessage() {}

func (x *DislikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DislikeRequest.ProtoReflect.Descriptor instead.
func (*DislikeRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{1}
}

func (x *DislikeRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *DislikeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type HasLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MangaId string `protobuf:"bytes,2,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
}

func (x *HasLikeRequest) Reset() {
	*x = HasLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasLikeRequest) ProtoMessage() {}

func (x *HasLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasLikeRequest.ProtoReflect.Descriptor instead.
func (*HasLikeRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{2}
}

func (x *HasLikeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HasLikeRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

type HasLikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Has bool `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
}

func (x *HasLikeResponse) Reset() {
	*x = HasLikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasLikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasLikeResponse) ProtoMessage() {}

func (x *HasLikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasLikeResponse.ProtoReflect.Descriptor instead.
func (*HasLikeResponse) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{3}
}

func (x *HasLikeResponse) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

type RateMangaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MangaId string `protobuf:"bytes,2,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	Rate    uint32 `protobuf:"varint,3,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *RateMangaRequest) Reset() {
	*x = RateMangaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateMangaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateMangaRequest) ProtoMessage() {}

func (x *RateMangaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateMangaRequest.ProtoReflect.Descriptor instead.
func (*RateMangaRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{4}
}

func (x *RateMangaRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RateMangaRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *RateMangaRequest) GetRate() uint32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type UserRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MangaId string `protobuf:"bytes,2,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
}

func (x *UserRateRequest) Reset() {
	*x = UserRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRateRequest) ProtoMessage() {}

func (x *UserRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRateRequest.ProtoReflect.Descriptor instead.
func (*UserRateRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{5}
}

func (x *UserRateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRateRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

type UserRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate *float32 `protobuf:"fixed32,1,opt,name=rate,proto3,oneof" json:"rate,omitempty"`
}

func (x *UserRateResponse) Reset() {
	*x = UserRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRateResponse) ProtoMessage() {}

func (x *UserRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRateResponse.ProtoReflect.Descriptor instead.
func (*UserRateResponse) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{6}
}

func (x *UserRateResponse) GetRate() float32 {
	if x != nil && x.Rate != nil {
		return *x.Rate
	}
	return 0
}

var File_activity_proto protoreflect.FileDescriptor

var file_activity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x1a, 0x0b, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0b, 0x4c, 0x69, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67,
	0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x44, 0x69,
	0x73, 0x6c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42,
	0x0a, 0x0e, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67,
	0x61, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61,
	0x49, 0x64, 0x22, 0x23, 0x0a, 0x0f, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x68, 0x61, 0x73, 0x22, 0x58, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x22, 0x43, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x32, 0xa0, 0x02, 0x0a,
	0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x6b,
	0x65, 0x12, 0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4c,
	0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x1a, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x6c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x42, 0x0a, 0x07, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1a, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x67, 0x61, 0x12, 0x1c, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1e, 0x5a, 0x1c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_proto_rawDescOnce sync.Once
	file_activity_proto_rawDescData = file_activity_proto_rawDesc
)

func file_activity_proto_rawDescGZIP() []byte {
	file_activity_proto_rawDescOnce.Do(func() {
		file_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_proto_rawDescData)
	})
	return file_activity_proto_rawDescData
}

var file_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_activity_proto_goTypes = []interface{}{
	(*LikeRequest)(nil),      // 0: micromango.LikeRequest
	(*DislikeRequest)(nil),   // 1: micromango.DislikeRequest
	(*HasLikeRequest)(nil),   // 2: micromango.HasLikeRequest
	(*HasLikeResponse)(nil),  // 3: micromango.HasLikeResponse
	(*RateMangaRequest)(nil), // 4: micromango.RateMangaRequest
	(*UserRateRequest)(nil),  // 5: micromango.UserRateRequest
	(*UserRateResponse)(nil), // 6: micromango.UserRateResponse
	(*share.Empty)(nil),      // 7: Empty
}
var file_activity_proto_depIdxs = []int32{
	0, // 0: micromango.Activity.Like:input_type -> micromango.LikeRequest
	1, // 1: micromango.Activity.Dislike:input_type -> micromango.DislikeRequest
	2, // 2: micromango.Activity.HasLike:input_type -> micromango.HasLikeRequest
	4, // 3: micromango.Activity.RateManga:input_type -> micromango.RateMangaRequest
	5, // 4: micromango.Activity.UserRate:input_type -> micromango.UserRateRequest
	7, // 5: micromango.Activity.Like:output_type -> Empty
	7, // 6: micromango.Activity.Dislike:output_type -> Empty
	3, // 7: micromango.Activity.HasLike:output_type -> micromango.HasLikeResponse
	7, // 8: micromango.Activity.RateManga:output_type -> Empty
	6, // 9: micromango.Activity.UserRate:output_type -> micromango.UserRateResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_activity_proto_init() }
func file_activity_proto_init() {
	if File_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DislikeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasLikeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasLikeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateMangaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_activity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_activity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_activity_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_proto_goTypes,
		DependencyIndexes: file_activity_proto_depIdxs,
		MessageInfos:      file_activity_proto_msgTypes,
	}.Build()
	File_activity_proto = out.File
	file_activity_proto_rawDesc = nil
	file_activity_proto_goTypes = nil
	file_activity_proto_depIdxs = nil
}
