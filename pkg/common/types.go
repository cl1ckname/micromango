package common

type File struct {
	File     []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}
