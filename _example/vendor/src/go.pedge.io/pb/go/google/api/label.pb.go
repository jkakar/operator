// Code generated by protoc-gen-go.
// source: google/api/label.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Value types that can be used as label values.
type LabelDescriptor_ValueType int32

const (
	// A variable-length string. This is the default.
	LabelDescriptor_STRING LabelDescriptor_ValueType = 0
	// Boolean; true or false.
	LabelDescriptor_BOOL LabelDescriptor_ValueType = 1
	// A 64-bit signed integer.
	LabelDescriptor_INT64 LabelDescriptor_ValueType = 2
)

var LabelDescriptor_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
}
var LabelDescriptor_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
}

func (x LabelDescriptor_ValueType) String() string {
	return proto.EnumName(LabelDescriptor_ValueType_name, int32(x))
}
func (LabelDescriptor_ValueType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A description of a label.
type LabelDescriptor struct {
	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The type of data that can be assigned to the label.
	ValueType LabelDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,enum=google.api.LabelDescriptor_ValueType" json:"value_type,omitempty"`
	// A human-readable description for the label.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *LabelDescriptor) Reset()                    { *m = LabelDescriptor{} }
func (m *LabelDescriptor) String() string            { return proto.CompactTextString(m) }
func (*LabelDescriptor) ProtoMessage()               {}
func (*LabelDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*LabelDescriptor)(nil), "google.api.LabelDescriptor")
	proto.RegisterEnum("google.api.LabelDescriptor_ValueType", LabelDescriptor_ValueType_name, LabelDescriptor_ValueType_value)
}

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x49, 0x4c, 0x4a, 0xcd, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x88, 0xeb, 0x01, 0xc5, 0x95, 0xe6, 0x30, 0x72, 0xf1, 0xfb, 0x80,
	0xe4, 0x5c, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x84, 0xb8, 0xb9, 0x98, 0xb3,
	0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x85, 0x2c, 0xb9, 0xb8, 0xca, 0x12, 0x73, 0x4a,
	0x53, 0xe3, 0x4b, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x80, 0x62, 0x7c, 0x46, 0xaa, 0x7a, 0x08, 0x13,
	0xf4, 0xd0, 0x74, 0xeb, 0x85, 0x81, 0x54, 0x87, 0x00, 0x15, 0x0b, 0x09, 0x73, 0x71, 0xa7, 0x40,
	0xc5, 0x33, 0xf3, 0xf3, 0x24, 0x98, 0x41, 0xe6, 0x29, 0xe9, 0x70, 0x71, 0x22, 0x54, 0x70, 0x71,
	0xb1, 0x05, 0x87, 0x04, 0x79, 0xfa, 0xb9, 0x0b, 0x30, 0x08, 0x71, 0x70, 0xb1, 0x38, 0xf9, 0xfb,
	0xfb, 0x08, 0x30, 0x0a, 0x71, 0x72, 0xb1, 0x7a, 0xfa, 0x85, 0x98, 0x99, 0x08, 0x30, 0x39, 0xc9,
	0x71, 0xf1, 0x25, 0xe7, 0xe7, 0x22, 0x59, 0xe7, 0xc4, 0x05, 0xb6, 0x2f, 0x00, 0xe4, 0x91, 0x00,
	0xc6, 0x24, 0x36, 0xb0, 0x8f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x45, 0xbb, 0x66,
	0xeb, 0x00, 0x00, 0x00,
}
