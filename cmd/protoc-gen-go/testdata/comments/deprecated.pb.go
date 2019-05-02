// Code generated by protoc-gen-go. DO NOT EDIT.
// comments/deprecated.proto is a deprecated file.

package comments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

// Deprecated: Use DeprecatedEnum.Type.Values instead.
var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

// Deprecated: Use DeprecatedEnum.Type.Values instead.
var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) Enum() *DeprecatedEnum {
	p := new(DeprecatedEnum)
	*p = x
	return p
}

func (x DeprecatedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeprecatedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_comments_deprecated_proto_enumTypes[0].Descriptor()
}

func (x DeprecatedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeprecatedEnum.Type instead.
func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return file_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
type DeprecatedMessage struct {
	DeprecatedField      string                  `protobuf:"bytes,1,opt,name=deprecated_field,json=deprecatedField,proto3" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *DeprecatedMessage) Reset() {
	*x = DeprecatedMessage{}
}

func (x *DeprecatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeprecatedMessage) ProtoMessage() {}

func (x *DeprecatedMessage) ProtoReflect() protoreflect.Message {
	return file_comments_deprecated_proto_msgTypes[0].MessageOf(x)
}

func (m *DeprecatedMessage) XXX_Methods() *protoiface.Methods {
	return file_comments_deprecated_proto_msgTypes[0].Methods()
}

// Deprecated: Use DeprecatedMessage.ProtoReflect.Type instead.
func (*DeprecatedMessage) Descriptor() ([]byte, []int) {
	return file_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *DeprecatedMessage) GetDeprecatedField() string {
	if x != nil {
		return x.DeprecatedField
	}
	return ""
}

var File_comments_deprecated_proto protoreflect.FileDescriptor

var file_comments_deprecated_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x2a, 0x28, 0x0a, 0x0e,
	0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12,
	0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x02,
	0x08, 0x01, 0x1a, 0x02, 0x18, 0x01, 0x42, 0x43, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0xb8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_comments_deprecated_proto_rawDescOnce sync.Once
	file_comments_deprecated_proto_rawDescData = file_comments_deprecated_proto_rawDesc
)

func file_comments_deprecated_proto_rawDescGZIP() []byte {
	file_comments_deprecated_proto_rawDescOnce.Do(func() {
		file_comments_deprecated_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_deprecated_proto_rawDescData)
	})
	return file_comments_deprecated_proto_rawDescData
}

var file_comments_deprecated_proto_enumTypes = make([]protoreflect.EnumType, 1)
var file_comments_deprecated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_comments_deprecated_proto_goTypes = []interface{}{
	(DeprecatedEnum)(0),       // 0: goproto.protoc.comments.DeprecatedEnum
	(*DeprecatedMessage)(nil), // 1: goproto.protoc.comments.DeprecatedMessage
}
var file_comments_deprecated_proto_depIdxs = []int32{}

func init() { file_comments_deprecated_proto_init() }
func file_comments_deprecated_proto_init() {
	if File_comments_deprecated_proto != nil {
		return
	}
	File_comments_deprecated_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_comments_deprecated_proto_rawDesc,
		GoTypes:            file_comments_deprecated_proto_goTypes,
		DependencyIndexes:  file_comments_deprecated_proto_depIdxs,
		EnumOutputTypes:    file_comments_deprecated_proto_enumTypes,
		MessageOutputTypes: file_comments_deprecated_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_comments_deprecated_proto_rawDesc = nil
	file_comments_deprecated_proto_goTypes = nil
	file_comments_deprecated_proto_depIdxs = nil
}
