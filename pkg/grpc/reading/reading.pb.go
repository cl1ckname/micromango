// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: reading.proto

package reading

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddMangaContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
}

func (x *AddMangaContentRequest) Reset() {
	*x = AddMangaContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMangaContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMangaContentRequest) ProtoMessage() {}

func (x *AddMangaContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMangaContentRequest.ProtoReflect.Descriptor instead.
func (*AddMangaContentRequest) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{0}
}

func (x *AddMangaContentRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

type MangaContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: body:manga
	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
}

func (x *MangaContentRequest) Reset() {
	*x = MangaContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangaContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangaContentRequest) ProtoMessage() {}

func (x *MangaContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangaContentRequest.ProtoReflect.Descriptor instead.
func (*MangaContentRequest) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{1}
}

func (x *MangaContentRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

type MangaContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId  string                              `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	Chapters []*MangaContentResponse_ChapterHead `protobuf:"bytes,2,rep,name=chapters,proto3" json:"chapters,omitempty"`
}

func (x *MangaContentResponse) Reset() {
	*x = MangaContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangaContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangaContentResponse) ProtoMessage() {}

func (x *MangaContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangaContentResponse.ProtoReflect.Descriptor instead.
func (*MangaContentResponse) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{2}
}

func (x *MangaContentResponse) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *MangaContentResponse) GetChapters() []*MangaContentResponse_ChapterHead {
	if x != nil {
		return x.Chapters
	}
	return nil
}

type AddChapterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *AddChapterRequest) Reset() {
	*x = AddChapterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChapterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChapterRequest) ProtoMessage() {}

func (x *AddChapterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChapterRequest.ProtoReflect.Descriptor instead.
func (*AddChapterRequest) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{3}
}

func (x *AddChapterRequest) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *AddChapterRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ChapterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterId string `protobuf:"bytes,1,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
}

func (x *ChapterRequest) Reset() {
	*x = ChapterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterRequest) ProtoMessage() {}

func (x *ChapterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterRequest.ProtoReflect.Descriptor instead.
func (*ChapterRequest) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{4}
}

func (x *ChapterRequest) GetChapterId() string {
	if x != nil {
		return x.ChapterId
	}
	return ""
}

type ChapterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterId     string                      `protobuf:"bytes,1,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
	MangaId       string                      `protobuf:"bytes,2,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	ChapterNumber uint32                      `protobuf:"varint,3,opt,name=chapterNumber,proto3" json:"chapterNumber,omitempty"`
	Title         string                      `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Pages         []*ChapterResponse_PageHead `protobuf:"bytes,5,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *ChapterResponse) Reset() {
	*x = ChapterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterResponse) ProtoMessage() {}

func (x *ChapterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterResponse.ProtoReflect.Descriptor instead.
func (*ChapterResponse) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{5}
}

func (x *ChapterResponse) GetChapterId() string {
	if x != nil {
		return x.ChapterId
	}
	return ""
}

func (x *ChapterResponse) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *ChapterResponse) GetChapterNumber() uint32 {
	if x != nil {
		return x.ChapterNumber
	}
	return 0
}

func (x *ChapterResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChapterResponse) GetPages() []*ChapterResponse_PageHead {
	if x != nil {
		return x.Pages
	}
	return nil
}

type AddPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterId  string `protobuf:"bytes,1,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
	PageNumber uint32 `protobuf:"varint,2,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	Image      string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *AddPageRequest) Reset() {
	*x = AddPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPageRequest) ProtoMessage() {}

func (x *AddPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPageRequest.ProtoReflect.Descriptor instead.
func (*AddPageRequest) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{6}
}

func (x *AddPageRequest) GetChapterId() string {
	if x != nil {
		return x.ChapterId
	}
	return ""
}

func (x *AddPageRequest) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *AddPageRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId string `protobuf:"bytes,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{7}
}

func (x *PageRequest) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

type PageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId     string `protobuf:"bytes,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
	ChapterId  string `protobuf:"bytes,2,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
	PageNumber uint32 `protobuf:"varint,3,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	Image      string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *PageResponse) Reset() {
	*x = PageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResponse) ProtoMessage() {}

func (x *PageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResponse.ProtoReflect.Descriptor instead.
func (*PageResponse) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{8}
}

func (x *PageResponse) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

func (x *PageResponse) GetChapterId() string {
	if x != nil {
		return x.ChapterId
	}
	return ""
}

func (x *PageResponse) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PageResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type MangaContentResponse_ChapterHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterId     string `protobuf:"bytes,1,opt,name=chapterId,proto3" json:"chapterId,omitempty"`
	ChapterNumber uint32 `protobuf:"varint,2,opt,name=chapterNumber,proto3" json:"chapterNumber,omitempty"`
	Title         string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *MangaContentResponse_ChapterHead) Reset() {
	*x = MangaContentResponse_ChapterHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangaContentResponse_ChapterHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangaContentResponse_ChapterHead) ProtoMessage() {}

func (x *MangaContentResponse_ChapterHead) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangaContentResponse_ChapterHead.ProtoReflect.Descriptor instead.
func (*MangaContentResponse_ChapterHead) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MangaContentResponse_ChapterHead) GetChapterId() string {
	if x != nil {
		return x.ChapterId
	}
	return ""
}

func (x *MangaContentResponse_ChapterHead) GetChapterNumber() uint32 {
	if x != nil {
		return x.ChapterNumber
	}
	return 0
}

func (x *MangaContentResponse_ChapterHead) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ChapterResponse_PageHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId     string `protobuf:"bytes,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
	PageNumber uint32 `protobuf:"varint,3,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
}

func (x *ChapterResponse_PageHead) Reset() {
	*x = ChapterResponse_PageHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterResponse_PageHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterResponse_PageHead) ProtoMessage() {}

func (x *ChapterResponse_PageHead) ProtoReflect() protoreflect.Message {
	mi := &file_reading_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterResponse_PageHead.ProtoReflect.Descriptor instead.
func (*ChapterResponse_PageHead) Descriptor() ([]byte, []int) {
	return file_reading_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ChapterResponse_PageHead) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

func (x *ChapterResponse_PageHead) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

var File_reading_proto protoreflect.FileDescriptor

var file_reading_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x22, 0x32, 0x0a, 0x16, 0x41,
	0x64, 0x64, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x22,
	0x2f, 0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64,
	0x22, 0xe3, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e,
	0x67, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67,
	0x61, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x48,
	0x65, 0x61, 0x64, 0x52, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x67, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61,
	0x6e, 0x67, 0x61, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x43,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x85, 0x02, 0x0a, 0x0f,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x1a,
	0x42, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x22, 0x7a, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0xc8, 0x03, 0x0a,
	0x07, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x41,
	0x64, 0x64, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x43,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_reading_proto_rawDescOnce sync.Once
	file_reading_proto_rawDescData = file_reading_proto_rawDesc
)

func file_reading_proto_rawDescGZIP() []byte {
	file_reading_proto_rawDescOnce.Do(func() {
		file_reading_proto_rawDescData = protoimpl.X.CompressGZIP(file_reading_proto_rawDescData)
	})
	return file_reading_proto_rawDescData
}

var file_reading_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_reading_proto_goTypes = []interface{}{
	(*AddMangaContentRequest)(nil),           // 0: micromango.AddMangaContentRequest
	(*MangaContentRequest)(nil),              // 1: micromango.MangaContentRequest
	(*MangaContentResponse)(nil),             // 2: micromango.MangaContentResponse
	(*AddChapterRequest)(nil),                // 3: micromango.AddChapterRequest
	(*ChapterRequest)(nil),                   // 4: micromango.ChapterRequest
	(*ChapterResponse)(nil),                  // 5: micromango.ChapterResponse
	(*AddPageRequest)(nil),                   // 6: micromango.AddPageRequest
	(*PageRequest)(nil),                      // 7: micromango.PageRequest
	(*PageResponse)(nil),                     // 8: micromango.PageResponse
	(*MangaContentResponse_ChapterHead)(nil), // 9: micromango.MangaContentResponse.ChapterHead
	(*ChapterResponse_PageHead)(nil),         // 10: micromango.ChapterResponse.PageHead
}
var file_reading_proto_depIdxs = []int32{
	9,  // 0: micromango.MangaContentResponse.chapters:type_name -> micromango.MangaContentResponse.ChapterHead
	10, // 1: micromango.ChapterResponse.pages:type_name -> micromango.ChapterResponse.PageHead
	1,  // 2: micromango.Reading.GetMangaContent:input_type -> micromango.MangaContentRequest
	0,  // 3: micromango.Reading.AddMangaContent:input_type -> micromango.AddMangaContentRequest
	4,  // 4: micromango.Reading.GetChapter:input_type -> micromango.ChapterRequest
	3,  // 5: micromango.Reading.AddChapter:input_type -> micromango.AddChapterRequest
	7,  // 6: micromango.Reading.GetPage:input_type -> micromango.PageRequest
	6,  // 7: micromango.Reading.AddPage:input_type -> micromango.AddPageRequest
	2,  // 8: micromango.Reading.GetMangaContent:output_type -> micromango.MangaContentResponse
	2,  // 9: micromango.Reading.AddMangaContent:output_type -> micromango.MangaContentResponse
	5,  // 10: micromango.Reading.GetChapter:output_type -> micromango.ChapterResponse
	5,  // 11: micromango.Reading.AddChapter:output_type -> micromango.ChapterResponse
	8,  // 12: micromango.Reading.GetPage:output_type -> micromango.PageResponse
	8,  // 13: micromango.Reading.AddPage:output_type -> micromango.PageResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_reading_proto_init() }
func file_reading_proto_init() {
	if File_reading_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMangaContentRequest); i {
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
		file_reading_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangaContentRequest); i {
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
		file_reading_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangaContentResponse); i {
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
		file_reading_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChapterRequest); i {
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
		file_reading_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterRequest); i {
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
		file_reading_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterResponse); i {
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
		file_reading_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPageRequest); i {
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
		file_reading_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_reading_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResponse); i {
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
		file_reading_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangaContentResponse_ChapterHead); i {
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
		file_reading_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterResponse_PageHead); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reading_proto_goTypes,
		DependencyIndexes: file_reading_proto_depIdxs,
		MessageInfos:      file_reading_proto_msgTypes,
	}.Build()
	File_reading_proto = out.File
	file_reading_proto_rawDesc = nil
	file_reading_proto_goTypes = nil
	file_reading_proto_depIdxs = nil
}
