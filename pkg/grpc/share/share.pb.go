// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: share.proto

package share

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

type ListName int32

const (
	ListName_UNKNOWN   ListName = 0
	ListName_READING   ListName = 1
	ListName_TOREAD    ListName = 2
	ListName_COMPLETED ListName = 3
	ListName_DROP      ListName = 4
)

// Enum value maps for ListName.
var (
	ListName_name = map[int32]string{
		0: "UNKNOWN",
		1: "READING",
		2: "TOREAD",
		3: "COMPLETED",
		4: "DROP",
	}
	ListName_value = map[string]int32{
		"UNKNOWN":   0,
		"READING":   1,
		"TOREAD":    2,
		"COMPLETED": 3,
		"DROP":      4,
	}
)

func (x ListName) Enum() *ListName {
	p := new(ListName)
	*p = x
	return p
}

func (x ListName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListName) Descriptor() protoreflect.EnumDescriptor {
	return file_share_proto_enumTypes[0].Descriptor()
}

func (ListName) Type() protoreflect.EnumType {
	return &file_share_proto_enumTypes[0]
}

func (x ListName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListName.Descriptor instead.
func (ListName) EnumDescriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

type MangaPreviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MangaId string `protobuf:"bytes,1,opt,name=mangaId,proto3" json:"mangaId,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Cover   string `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
}

func (x *MangaPreviewResponse) Reset() {
	*x = MangaPreviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MangaPreviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MangaPreviewResponse) ProtoMessage() {}

func (x *MangaPreviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MangaPreviewResponse.ProtoReflect.Descriptor instead.
func (*MangaPreviewResponse) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

func (x *MangaPreviewResponse) GetMangaId() string {
	if x != nil {
		return x.MangaId
	}
	return ""
}

func (x *MangaPreviewResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MangaPreviewResponse) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[1]
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
	return file_share_proto_rawDescGZIP(), []int{1}
}

var File_share_proto protoreflect.FileDescriptor

var file_share_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a,
	0x14, 0x4d, 0x61, 0x6e, 0x67, 0x61, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x67, 0x61, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x2a, 0x49, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x45, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x4f,
	0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50, 0x10, 0x04, 0x42,
	0x1b, 0x5a, 0x19, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_share_proto_rawDescOnce sync.Once
	file_share_proto_rawDescData = file_share_proto_rawDesc
)

func file_share_proto_rawDescGZIP() []byte {
	file_share_proto_rawDescOnce.Do(func() {
		file_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_share_proto_rawDescData)
	})
	return file_share_proto_rawDescData
}

var file_share_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_share_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_share_proto_goTypes = []interface{}{
	(ListName)(0),                // 0: ListName
	(*MangaPreviewResponse)(nil), // 1: MangaPreviewResponse
	(*Empty)(nil),                // 2: Empty
}
var file_share_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_share_proto_init() }
func file_share_proto_init() {
	if File_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MangaPreviewResponse); i {
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
		file_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_share_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_share_proto_goTypes,
		DependencyIndexes: file_share_proto_depIdxs,
		EnumInfos:         file_share_proto_enumTypes,
		MessageInfos:      file_share_proto_msgTypes,
	}.Build()
	File_share_proto = out.File
	file_share_proto_rawDesc = nil
	file_share_proto_goTypes = nil
	file_share_proto_depIdxs = nil
}
