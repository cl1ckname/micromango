// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: catalog.proto

package catalog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reading "micromango/pkg/grpc/reading"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{0}
}

type MangasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mangas []*share.MangaPreviewResponse `protobuf:"bytes,1,rep,name=mangas,proto3" json:"mangas,omitempty"`
}

func (x *MangasResponse) Reset() {
	*x = MangasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangasResponse) ProtoMessage() {}

func (x *MangasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangasResponse.ProtoReflect.Descriptor instead.
func (*MangasResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *MangasResponse) GetMangas() []*share.MangaPreviewResponse {
	if x != nil {
		return x.Mangas
	}
	return nil
}

type MangaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string  `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	UserId  *string `protobuf:"bytes,2,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
}

func (x *MangaRequest) Reset() {
	*x = MangaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangaRequest) ProtoMessage() {}

func (x *MangaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangaRequest.ProtoReflect.Descriptor instead.
func (*MangaRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *MangaRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *MangaRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type MangaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId     string                        `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	Title       string                        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Cover       string                        `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Description string                        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Content     *reading.MangaContentResponse `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt   string                        `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	List        *share.ListName               `protobuf:"varint,7,opt,name=list,proto3,enum=ListName,oneof" json:"list,omitempty"`
	Genres      []uint32                      `protobuf:"varint,8,rep,packed,name=genres,proto3" json:"genres,omitempty"`
	ListStats   map[uint32]uint64             `protobuf:"bytes,9,rep,name=listStats,proto3" json:"listStats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Likes       uint64                        `protobuf:"varint,10,opt,name=likes,proto3" json:"likes,omitempty"`
	Liked       bool                          `protobuf:"varint,11,opt,name=liked,proto3" json:"liked,omitempty"`
	Rate        *share.AvgMangaRateResponse   `protobuf:"bytes,12,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *MangaResponse) Reset() {
	*x = MangaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangaResponse) ProtoMessage() {}

func (x *MangaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangaResponse.ProtoReflect.Descriptor instead.
func (*MangaResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *MangaResponse) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *MangaResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MangaResponse) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *MangaResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MangaResponse) GetContent() *reading.MangaContentResponse {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *MangaResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MangaResponse) GetList() share.ListName {
	if x != nil && x.List != nil {
		return *x.List
	}
	return share.ListName(0)
}

func (x *MangaResponse) GetGenres() []uint32 {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *MangaResponse) GetListStats() map[uint32]uint64 {
	if x != nil {
		return x.ListStats
	}
	return nil
}

func (x *MangaResponse) GetLikes() uint64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *MangaResponse) GetLiked() bool {
	if x != nil {
		return x.Liked
	}
	return false
}

func (x *MangaResponse) GetRate() *share.AvgMangaRateResponse {
	if x != nil {
		return x.Rate
	}
	return nil
}

type GetMangasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Starts        *string  `protobuf:"bytes,1,opt,name=starts,proto3,oneof" json:"starts,omitempty"`
	GenresInclude []uint32 `protobuf:"varint,3,rep,packed,name=genresInclude,proto3" json:"genresInclude,omitempty"`
	GenresExclude []uint32 `protobuf:"varint,4,rep,packed,name=genresExclude,proto3" json:"genresExclude,omitempty"`
}

func (x *GetMangasRequest) Reset() {
	*x = GetMangasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMangasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMangasRequest) ProtoMessage() {}

func (x *GetMangasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMangasRequest.ProtoReflect.Descriptor instead.
func (*GetMangasRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *GetMangasRequest) GetStarts() string {
	if x != nil && x.Starts != nil {
		return *x.Starts
	}
	return ""
}

func (x *GetMangasRequest) GetGenresInclude() []uint32 {
	if x != nil {
		return x.GenresInclude
	}
	return nil
}

func (x *GetMangasRequest) GetGenresExclude() []uint32 {
	if x != nil {
		return x.GenresExclude
	}
	return nil
}

type AddMangaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Cover       []byte   `protobuf:"bytes,2,opt,name=cover,proto3,oneof" json:"cover,omitempty"`
	Description *string  `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Genres      []uint32 `protobuf:"varint,4,rep,packed,name=genres,proto3" json:"genres,omitempty"`
}

func (x *AddMangaRequest) Reset() {
	*x = AddMangaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMangaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMangaRequest) ProtoMessage() {}

func (x *AddMangaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMangaRequest.ProtoReflect.Descriptor instead.
func (*AddMangaRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *AddMangaRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddMangaRequest) GetCover() []byte {
	if x != nil {
		return x.Cover
	}
	return nil
}

func (x *AddMangaRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *AddMangaRequest) GetGenres() []uint32 {
	if x != nil {
		return x.Genres
	}
	return nil
}

type UpdateMangaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId     string   `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	Title       *string  `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Cover       []byte   `protobuf:"bytes,3,opt,name=cover,proto3,oneof" json:"cover,omitempty"`
	Description *string  `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Genres      []uint32 `protobuf:"varint,5,rep,packed,name=genres,proto3" json:"genres,omitempty"`
}

func (x *UpdateMangaRequest) Reset() {
	*x = UpdateMangaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMangaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMangaRequest) ProtoMessage() {}

func (x *UpdateMangaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMangaRequest.ProtoReflect.Descriptor instead.
func (*UpdateMangaRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMangaRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *UpdateMangaRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateMangaRequest) GetCover() []byte {
	if x != nil {
		return x.Cover
	}
	return nil
}

func (x *UpdateMangaRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateMangaRequest) GetGenres() []uint32 {
	if x != nil {
		return x.Genres
	}
	return nil
}

type DeleteMangaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
}

func (x *DeleteMangaRequest) Reset() {
	*x = DeleteMangaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMangaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMangaRequest) ProtoMessage() {}

func (x *DeleteMangaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMangaRequest.ProtoReflect.Descriptor instead.
func (*DeleteMangaRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMangaRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaList []string `protobuf:"bytes,1,rep,name=mangaList,proto3" json:"mangaList,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{8}
}

func (x *GetListRequest) GetMangaList() []string {
	if x != nil {
		return x.MangaList
	}
	return nil
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviewList []*share.MangaPreviewResponse `protobuf:"bytes,1,rep,name=previewList,proto3" json:"previewList,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{9}
}

func (x *GetListResponse) GetPreviewList() []*share.MangaPreviewResponse {
	if x != nil {
		return x.PreviewList
	}
	return nil
}

var File_catalog_proto protoreflect.FileDescriptor

var file_catalog_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x3f, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x6d, 0x61, 0x6e, 0x67, 0x61,
	0x73, 0x22, 0x50, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xf3, 0x03, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67,
	0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x76, 0x67, 0x4d, 0x61, 0x6e, 0x67, 0x61,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0d, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x45, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xc7, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x6e, 0x67, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xa5, 0x03, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x12, 0x18,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x73,
	0x12, 0x1c, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67,
	0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f,
	0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x12, 0x1e, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x12, 0x1e, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d,
	0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x67, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d,
	0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d,
	0x5a, 0x1b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_proto_rawDescOnce sync.Once
	file_catalog_proto_rawDescData = file_catalog_proto_rawDesc
)

func file_catalog_proto_rawDescGZIP() []byte {
	file_catalog_proto_rawDescOnce.Do(func() {
		file_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_proto_rawDescData)
	})
	return file_catalog_proto_rawDescData
}

var file_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_catalog_proto_goTypes = []interface{}{
	(*Empty)(nil),                        // 0: micromango.Empty
	(*MangasResponse)(nil),               // 1: micromango.MangasResponse
	(*MangaRequest)(nil),                 // 2: micromango.MangaRequest
	(*MangaResponse)(nil),                // 3: micromango.MangaResponse
	(*GetMangasRequest)(nil),             // 4: micromango.GetMangasRequest
	(*AddMangaRequest)(nil),              // 5: micromango.AddMangaRequest
	(*UpdateMangaRequest)(nil),           // 6: micromango.UpdateMangaRequest
	(*DeleteMangaRequest)(nil),           // 7: micromango.DeleteMangaRequest
	(*GetListRequest)(nil),               // 8: micromango.GetListRequest
	(*GetListResponse)(nil),              // 9: micromango.GetListResponse
	nil,                                  // 10: micromango.MangaResponse.ListStatsEntry
	(*share.MangaPreviewResponse)(nil),   // 11: MangaPreviewResponse
	(*reading.MangaContentResponse)(nil), // 12: micromango.MangaContentResponse
	(share.ListName)(0),                  // 13: ListName
	(*share.AvgMangaRateResponse)(nil),   // 14: AvgMangaRateResponse
}
var file_catalog_proto_depIdxs = []int32{
	11, // 0: micromango.MangasResponse.mangas:type_name -> MangaPreviewResponse
	12, // 1: micromango.MangaResponse.content:type_name -> micromango.MangaContentResponse
	13, // 2: micromango.MangaResponse.list:type_name -> ListName
	10, // 3: micromango.MangaResponse.listStats:type_name -> micromango.MangaResponse.ListStatsEntry
	14, // 4: micromango.MangaResponse.rate:type_name -> AvgMangaRateResponse
	11, // 5: micromango.GetListResponse.previewList:type_name -> MangaPreviewResponse
	2,  // 6: micromango.Catalog.GetManga:input_type -> micromango.MangaRequest
	4,  // 7: micromango.Catalog.GetMangas:input_type -> micromango.GetMangasRequest
	5,  // 8: micromango.Catalog.AddManga:input_type -> micromango.AddMangaRequest
	6,  // 9: micromango.Catalog.UpdateManga:input_type -> micromango.UpdateMangaRequest
	7,  // 10: micromango.Catalog.DeleteManga:input_type -> micromango.DeleteMangaRequest
	8,  // 11: micromango.Catalog.GetList:input_type -> micromango.GetListRequest
	3,  // 12: micromango.Catalog.GetManga:output_type -> micromango.MangaResponse
	1,  // 13: micromango.Catalog.GetMangas:output_type -> micromango.MangasResponse
	3,  // 14: micromango.Catalog.AddManga:output_type -> micromango.MangaResponse
	3,  // 15: micromango.Catalog.UpdateManga:output_type -> micromango.MangaResponse
	0,  // 16: micromango.Catalog.DeleteManga:output_type -> micromango.Empty
	9,  // 17: micromango.Catalog.GetList:output_type -> micromango.GetListResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_catalog_proto_init() }
func file_catalog_proto_init() {
	if File_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangasResponse); i {
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
		file_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangaRequest); i {
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
		file_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangaResponse); i {
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
		file_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMangasRequest); i {
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
		file_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMangaRequest); i {
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
		file_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMangaRequest); i {
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
		file_catalog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMangaRequest); i {
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
		file_catalog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_catalog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
	file_catalog_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_catalog_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_catalog_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_catalog_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_catalog_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_proto_goTypes,
		DependencyIndexes: file_catalog_proto_depIdxs,
		MessageInfos:      file_catalog_proto_msgTypes,
	}.Build()
	File_catalog_proto = out.File
	file_catalog_proto_rawDesc = nil
	file_catalog_proto_goTypes = nil
	file_catalog_proto_depIdxs = nil
}
