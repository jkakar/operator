// Code generated by protoc-gen-go.
// source: pb/protologpb.proto
// DO NOT EDIT!

/*
Package protologpb is a generated protocol buffer package.

It is generated from these files:
	pb/protologpb.proto

It has these top-level messages:
	Entry
*/
package protologpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Level is a logging level.
type Level int32

const (
	Level_LEVEL_DEBUG Level = 0
	Level_LEVEL_INFO  Level = 1
	Level_LEVEL_WARN  Level = 2
	Level_LEVEL_ERROR Level = 3
)

var Level_name = map[int32]string{
	0: "LEVEL_DEBUG",
	1: "LEVEL_INFO",
	2: "LEVEL_WARN",
	3: "LEVEL_ERROR",
}
var Level_value = map[string]int32{
	"LEVEL_DEBUG": 0,
	"LEVEL_INFO":  1,
	"LEVEL_WARN":  2,
	"LEVEL_ERROR": 3,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}
func (Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Entry is the object serialized for logging.
type Entry struct {
	// id may not be set depending on logger options
	// it is up to the user to determine if id is required
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Level     Level                      `protobuf:"varint,2,opt,name=level,enum=protolog.Level" json:"level,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// both context and fields may be set
	Context []*Entry_Message  `protobuf:"bytes,4,rep,name=context" json:"context,omitempty"`
	Fields  map[string]string `protobuf:"bytes,5,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// one of event, message, writer_output will be set
	Event        *Entry_Message `protobuf:"bytes,6,opt,name=event" json:"event,omitempty"`
	Message      string         `protobuf:"bytes,7,opt,name=message" json:"message,omitempty"`
	WriterOutput []byte         `protobuf:"bytes,8,opt,name=writer_output,json=writerOutput,proto3" json:"writer_output,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetContext() []*Entry_Message {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Entry) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Entry) GetEvent() *Entry_Message {
	if m != nil {
		return m.Event
	}
	return nil
}

// Message represents a serialized protobuf message.
// The name is the name registered with protolog.
type Entry_Message struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Entry_Message) Reset()                    { *m = Entry_Message{} }
func (m *Entry_Message) String() string            { return proto.CompactTextString(m) }
func (*Entry_Message) ProtoMessage()               {}
func (*Entry_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto.RegisterType((*Entry)(nil), "protolog.Entry")
	proto.RegisterType((*Entry_Message)(nil), "protolog.Entry.Message")
	proto.RegisterEnum("protolog.Level", Level_name, Level_value)
}

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x5d, 0x4b, 0xe3, 0x40,
	0x14, 0xdd, 0x24, 0x4d, 0xd3, 0xde, 0x66, 0xdb, 0x30, 0xbb, 0xb0, 0x43, 0xf6, 0x61, 0xcb, 0x8a,
	0x20, 0x82, 0x29, 0xb6, 0x2f, 0xd5, 0x37, 0x8b, 0xa9, 0x14, 0x6a, 0x0b, 0x83, 0x1f, 0xe0, 0x4b,
	0x49, 0xec, 0x34, 0x04, 0xf3, 0x45, 0x32, 0xa9, 0xf6, 0xcf, 0xf9, 0xdb, 0x4c, 0x67, 0x12, 0xa3,
	0x82, 0x6f, 0xe7, 0x9e, 0x7b, 0xce, 0xbd, 0x87, 0x3b, 0x03, 0xbf, 0x12, 0x77, 0x90, 0xa4, 0x31,
	0x8b, 0x83, 0xd8, 0x4b, 0x5c, 0x8b, 0x43, 0xd4, 0xaa, 0x18, 0xf3, 0x9f, 0x17, 0xc7, 0x5e, 0x40,
	0x85, 0xc4, 0xcd, 0x37, 0x03, 0xe6, 0x87, 0x34, 0x63, 0x4e, 0x98, 0x08, 0xe9, 0xff, 0x57, 0x05,
	0x54, 0x3b, 0x62, 0xe9, 0x0e, 0x75, 0x41, 0xf6, 0xd7, 0x58, 0xea, 0x4b, 0x47, 0x6d, 0x52, 0x20,
	0x74, 0x08, 0x6a, 0x40, 0xb7, 0x34, 0xc0, 0x72, 0x41, 0x75, 0x87, 0x3d, 0xab, 0x1a, 0x6a, 0xcd,
	0xf7, 0x34, 0x11, 0x5d, 0x34, 0x86, 0xf6, 0xfb, 0x4c, 0xac, 0x14, 0xd2, 0xce, 0xd0, 0xb4, 0xc4,
	0x56, 0xab, 0xda, 0x6a, 0xdd, 0x54, 0x0a, 0x52, 0x8b, 0xd1, 0x29, 0x68, 0x8f, 0x71, 0xc4, 0xe8,
	0x0b, 0xc3, 0x8d, 0xbe, 0x52, 0xf8, 0xfe, 0xd4, 0x2b, 0x78, 0x24, 0xeb, 0x9a, 0x66, 0x99, 0xe3,
	0x51, 0x52, 0xe9, 0xd0, 0x08, 0x9a, 0x1b, 0x9f, 0x06, 0xeb, 0x0c, 0xab, 0xdc, 0xf1, 0xf7, 0xab,
	0x63, 0xca, 0xbb, 0x1c, 0x93, 0x52, 0x8a, 0x4e, 0x40, 0x2d, 0x92, 0x46, 0x0c, 0x37, 0x79, 0xba,
	0x6f, 0xb7, 0x08, 0x15, 0xc2, 0xa0, 0x85, 0x82, 0xc1, 0x1a, 0x3f, 0x46, 0x55, 0xa2, 0x03, 0xf8,
	0xf9, 0x9c, 0xfa, 0x8c, 0xa6, 0xab, 0x38, 0x67, 0x49, 0xce, 0x70, 0xab, 0xe8, 0xeb, 0x44, 0x17,
	0xe4, 0x92, 0x73, 0xe6, 0x08, 0xb4, 0x72, 0x20, 0x42, 0xd0, 0x88, 0x9c, 0x90, 0x96, 0x37, 0xe5,
	0x18, 0xfd, 0x06, 0x75, 0xeb, 0x04, 0x39, 0xe5, 0x57, 0xd5, 0x89, 0x28, 0xcc, 0x33, 0xe8, 0x7c,
	0x48, 0x8e, 0x0c, 0x50, 0x9e, 0xe8, 0xae, 0xf4, 0xed, 0xe1, 0x67, 0x5b, 0xbb, 0xb4, 0x9d, 0xcb,
	0x63, 0xe9, 0x78, 0x06, 0x2a, 0x7f, 0x0f, 0xd4, 0x83, 0xce, 0xdc, 0xbe, 0xb3, 0xe7, 0xab, 0x4b,
	0x7b, 0x72, 0x7b, 0x65, 0xfc, 0x28, 0x1e, 0x14, 0x04, 0x31, 0x5b, 0x4c, 0x97, 0x86, 0x54, 0xd7,
	0xf7, 0x17, 0x64, 0x61, 0xc8, 0xb5, 0xc1, 0x26, 0x64, 0x49, 0x0c, 0x65, 0xa2, 0x3f, 0x40, 0xfd,
	0x95, 0xdc, 0x26, 0xc7, 0xa3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x7f, 0x51, 0x26, 0x62,
	0x02, 0x00, 0x00,
}
