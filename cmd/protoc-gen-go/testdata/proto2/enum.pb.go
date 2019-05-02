// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/enum.proto

package proto2

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

// EnumType1 comment.
type EnumType1 int32

const (
	// EnumType1_ONE comment.
	EnumType1_ONE EnumType1 = 1
	// EnumType1_TWO comment.
	EnumType1_TWO EnumType1 = 2
)

// Deprecated: Use EnumType1.Type.Values instead.
var EnumType1_name = map[int32]string{
	1: "ONE",
	2: "TWO",
}

// Deprecated: Use EnumType1.Type.Values instead.
var EnumType1_value = map[string]int32{
	"ONE": 1,
	"TWO": 2,
}

func (x EnumType1) Enum() *EnumType1 {
	p := new(EnumType1)
	*p = x
	return p
}

func (x EnumType1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumType1) Descriptor() protoreflect.EnumDescriptor {
	return file_proto2_enum_proto_enumTypes[0].Descriptor()
}

func (x EnumType1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumType1) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumType1(num)
	return nil
}

// Deprecated: Use EnumType1.Type instead.
func (EnumType1) EnumDescriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0}
}

type EnumType2 int32

const (
	EnumType2_duplicate1 EnumType2 = 1
	EnumType2_duplicate2 EnumType2 = 1
)

// Deprecated: Use EnumType2.Type.Values instead.
var EnumType2_name = map[int32]string{
	1: "duplicate1",
	// Duplicate value: 1: "duplicate2",
}

// Deprecated: Use EnumType2.Type.Values instead.
var EnumType2_value = map[string]int32{
	"duplicate1": 1,
	"duplicate2": 1,
}

func (x EnumType2) Enum() *EnumType2 {
	p := new(EnumType2)
	*p = x
	return p
}

func (x EnumType2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumType2) Descriptor() protoreflect.EnumDescriptor {
	return file_proto2_enum_proto_enumTypes[1].Descriptor()
}

func (x EnumType2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumType2) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumType2(num)
	return nil
}

// Deprecated: Use EnumType2.Type instead.
func (EnumType2) EnumDescriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{1}
}

// NestedEnumType1A comment.
type EnumContainerMessage1_NestedEnumType1A int32

const (
	// NestedEnumType1A_VALUE comment.
	EnumContainerMessage1_NESTED_1A_VALUE EnumContainerMessage1_NestedEnumType1A = 0
)

// Deprecated: Use EnumContainerMessage1_NestedEnumType1A.Type.Values instead.
var EnumContainerMessage1_NestedEnumType1A_name = map[int32]string{
	0: "NESTED_1A_VALUE",
}

// Deprecated: Use EnumContainerMessage1_NestedEnumType1A.Type.Values instead.
var EnumContainerMessage1_NestedEnumType1A_value = map[string]int32{
	"NESTED_1A_VALUE": 0,
}

func (x EnumContainerMessage1_NestedEnumType1A) Enum() *EnumContainerMessage1_NestedEnumType1A {
	p := new(EnumContainerMessage1_NestedEnumType1A)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1A) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_NestedEnumType1A) Descriptor() protoreflect.EnumDescriptor {
	return file_proto2_enum_proto_enumTypes[2].Descriptor()
}

func (x EnumContainerMessage1_NestedEnumType1A) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_NestedEnumType1A) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1A(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_NestedEnumType1A.Type instead.
func (EnumContainerMessage1_NestedEnumType1A) EnumDescriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0, 0}
}

type EnumContainerMessage1_NestedEnumType1B int32

const (
	EnumContainerMessage1_NESTED_1B_VALUE EnumContainerMessage1_NestedEnumType1B = 0
)

// Deprecated: Use EnumContainerMessage1_NestedEnumType1B.Type.Values instead.
var EnumContainerMessage1_NestedEnumType1B_name = map[int32]string{
	0: "NESTED_1B_VALUE",
}

// Deprecated: Use EnumContainerMessage1_NestedEnumType1B.Type.Values instead.
var EnumContainerMessage1_NestedEnumType1B_value = map[string]int32{
	"NESTED_1B_VALUE": 0,
}

func (x EnumContainerMessage1_NestedEnumType1B) Enum() *EnumContainerMessage1_NestedEnumType1B {
	p := new(EnumContainerMessage1_NestedEnumType1B)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1B) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_NestedEnumType1B) Descriptor() protoreflect.EnumDescriptor {
	return file_proto2_enum_proto_enumTypes[3].Descriptor()
}

func (x EnumContainerMessage1_NestedEnumType1B) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_NestedEnumType1B) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1B(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_NestedEnumType1B.Type instead.
func (EnumContainerMessage1_NestedEnumType1B) EnumDescriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0, 1}
}

// NestedEnumType2A comment.
type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A int32

const (
	// NestedEnumType2A_VALUE comment.
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2A_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A = 0
)

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A.Type.Values instead.
var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name = map[int32]string{
	0: "NESTED_2A_VALUE",
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A.Type.Values instead.
var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value = map[string]int32{
	"NESTED_2A_VALUE": 0,
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Descriptor() protoreflect.EnumDescriptor {
	return file_proto2_enum_proto_enumTypes[4].Descriptor()
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A.Type instead.
func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) EnumDescriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0, 0, 0}
}

type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B int32

const (
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2B_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B = 0
)

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B.Type.Values instead.
var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name = map[int32]string{
	0: "NESTED_2B_VALUE",
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B.Type.Values instead.
var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value = map[string]int32{
	"NESTED_2B_VALUE": 0,
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Descriptor() protoreflect.EnumDescriptor {
	return file_proto2_enum_proto_enumTypes[5].Descriptor()
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B.Type instead.
func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) EnumDescriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0, 0, 1}
}

type EnumContainerMessage1 struct {
	DefaultDuplicate1    *EnumType2              `protobuf:"varint,1,opt,name=default_duplicate1,json=defaultDuplicate1,enum=goproto.protoc.proto2.EnumType2,def=1" json:"default_duplicate1,omitempty"`
	DefaultDuplicate2    *EnumType2              `protobuf:"varint,2,opt,name=default_duplicate2,json=defaultDuplicate2,enum=goproto.protoc.proto2.EnumType2,def=1" json:"default_duplicate2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *EnumContainerMessage1) Reset() {
	*x = EnumContainerMessage1{}
}

func (x *EnumContainerMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumContainerMessage1) ProtoMessage() {}

func (x *EnumContainerMessage1) ProtoReflect() protoreflect.Message {
	return file_proto2_enum_proto_msgTypes[0].MessageOf(x)
}

func (m *EnumContainerMessage1) XXX_Methods() *protoiface.Methods {
	return file_proto2_enum_proto_msgTypes[0].Methods()
}

// Deprecated: Use EnumContainerMessage1.ProtoReflect.Type instead.
func (*EnumContainerMessage1) Descriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0}
}

const Default_EnumContainerMessage1_DefaultDuplicate1 EnumType2 = EnumType2_duplicate1
const Default_EnumContainerMessage1_DefaultDuplicate2 EnumType2 = EnumType2_duplicate2

func (x *EnumContainerMessage1) GetDefaultDuplicate1() EnumType2 {
	if x != nil && x.DefaultDuplicate1 != nil {
		return *x.DefaultDuplicate1
	}
	return Default_EnumContainerMessage1_DefaultDuplicate1
}

func (x *EnumContainerMessage1) GetDefaultDuplicate2() EnumType2 {
	if x != nil && x.DefaultDuplicate2 != nil {
		return *x.DefaultDuplicate2
	}
	return Default_EnumContainerMessage1_DefaultDuplicate2
}

type EnumContainerMessage1_EnumContainerMessage2 struct {
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *EnumContainerMessage1_EnumContainerMessage2) Reset() {
	*x = EnumContainerMessage1_EnumContainerMessage2{}
}

func (x *EnumContainerMessage1_EnumContainerMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumContainerMessage1_EnumContainerMessage2) ProtoMessage() {}

func (x *EnumContainerMessage1_EnumContainerMessage2) ProtoReflect() protoreflect.Message {
	return file_proto2_enum_proto_msgTypes[1].MessageOf(x)
}

func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Methods() *protoiface.Methods {
	return file_proto2_enum_proto_msgTypes[1].Methods()
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2.ProtoReflect.Type instead.
func (*EnumContainerMessage1_EnumContainerMessage2) Descriptor() ([]byte, []int) {
	return file_proto2_enum_proto_rawDescGZIP(), []int{0, 0}
}

var File_proto2_enum_proto protoreflect.FileDescriptor

var file_proto2_enum_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0x8e, 0x03, 0x0a, 0x15, 0x45,
	0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x12, 0x5b, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x32, 0x3a, 0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x31, 0x52, 0x11,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x31, 0x12, 0x5b, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x75, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x32, 0x3a,
	0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x32, 0x52, 0x11, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x32, 0x1a, 0x69,
	0x0a, 0x15, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x22, 0x27, 0x0a, 0x10, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x32, 0x41, 0x12, 0x13, 0x0a, 0x0f, 0x4e,
	0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x32, 0x41, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00,
	0x22, 0x27, 0x0a, 0x10, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x32, 0x42, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x32,
	0x42, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x22, 0x27, 0x0a, 0x10, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x31, 0x41, 0x12, 0x13, 0x0a,
	0x0f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x31, 0x41, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x00, 0x22, 0x27, 0x0a, 0x10, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x31, 0x42, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44,
	0x5f, 0x31, 0x42, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x2a, 0x1d, 0x0a, 0x09, 0x45,
	0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x31, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x2a, 0x51, 0x0a, 0x09, 0x45, 0x6e,
	0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x32, 0x12, 0x0e, 0x0a, 0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x31, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x32, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x04, 0x08, 0x02, 0x10,
	0x02, 0x22, 0x04, 0x08, 0x03, 0x10, 0x03, 0x2a, 0x09, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x44, 0x31, 0x2a, 0x09, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x32, 0x42, 0x3e, 0x5a,
	0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var (
	file_proto2_enum_proto_rawDescOnce sync.Once
	file_proto2_enum_proto_rawDescData = file_proto2_enum_proto_rawDesc
)

func file_proto2_enum_proto_rawDescGZIP() []byte {
	file_proto2_enum_proto_rawDescOnce.Do(func() {
		file_proto2_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto2_enum_proto_rawDescData)
	})
	return file_proto2_enum_proto_rawDescData
}

var file_proto2_enum_proto_enumTypes = make([]protoreflect.EnumType, 6)
var file_proto2_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto2_enum_proto_goTypes = []interface{}{
	(EnumType1)(0), // 0: goproto.protoc.proto2.EnumType1
	(EnumType2)(0), // 1: goproto.protoc.proto2.EnumType2
	(EnumContainerMessage1_NestedEnumType1A)(0),                       // 2: goproto.protoc.proto2.EnumContainerMessage1.NestedEnumType1A
	(EnumContainerMessage1_NestedEnumType1B)(0),                       // 3: goproto.protoc.proto2.EnumContainerMessage1.NestedEnumType1B
	(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)(0), // 4: goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2.NestedEnumType2A
	(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)(0), // 5: goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2.NestedEnumType2B
	(*EnumContainerMessage1)(nil),                                     // 6: goproto.protoc.proto2.EnumContainerMessage1
	(*EnumContainerMessage1_EnumContainerMessage2)(nil),               // 7: goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2
}
var file_proto2_enum_proto_depIdxs = []int32{
	1, // goproto.protoc.proto2.EnumContainerMessage1.default_duplicate1:type_name -> goproto.protoc.proto2.EnumType2
	1, // goproto.protoc.proto2.EnumContainerMessage1.default_duplicate2:type_name -> goproto.protoc.proto2.EnumType2
}

func init() { file_proto2_enum_proto_init() }
func file_proto2_enum_proto_init() {
	if File_proto2_enum_proto != nil {
		return
	}
	File_proto2_enum_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_proto2_enum_proto_rawDesc,
		GoTypes:            file_proto2_enum_proto_goTypes,
		DependencyIndexes:  file_proto2_enum_proto_depIdxs,
		EnumOutputTypes:    file_proto2_enum_proto_enumTypes,
		MessageOutputTypes: file_proto2_enum_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_proto2_enum_proto_rawDesc = nil
	file_proto2_enum_proto_goTypes = nil
	file_proto2_enum_proto_depIdxs = nil
}
