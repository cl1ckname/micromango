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

var File_share_proto protoreflect.FileDescriptor

var file_share_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x49, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x41, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x4f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50, 0x10, 0x04, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_share_proto_goTypes = []interface{}{
	(ListName)(0), // 0: ListName
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_share_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_share_proto_goTypes,
		DependencyIndexes: file_share_proto_depIdxs,
		EnumInfos:         file_share_proto_enumTypes,
	}.Build()
	File_share_proto = out.File
	file_share_proto_rawDesc = nil
	file_share_proto_goTypes = nil
	file_share_proto_depIdxs = nil
}