// Code generated by protoc-gen-go.
// source: operator/operator.proto
// DO NOT EDIT!

/*
Package operator is a generated protocol buffer package.

It is generated from these files:
	operator/operator.proto

It has these top-level messages:
	Output
	ServerStartupNotice
*/
package operator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Output struct {
	PlainText string `protobuf:"bytes,1,opt,name=PlainText" json:"PlainText,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerStartupNotice struct {
	Address  string `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	Protocol string `protobuf:"bytes,2,opt,name=Protocol" json:"Protocol,omitempty"`
}

func (m *ServerStartupNotice) Reset()                    { *m = ServerStartupNotice{} }
func (m *ServerStartupNotice) String() string            { return proto.CompactTextString(m) }
func (*ServerStartupNotice) ProtoMessage()               {}
func (*ServerStartupNotice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Output)(nil), "operator.Output")
	proto.RegisterType((*ServerStartupNotice)(nil), "operator.ServerStartupNotice")
}

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x2f, 0x48, 0x2d,
	0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x25, 0x69, 0x2e, 0x36, 0xff, 0xd2, 0x92, 0x82, 0xd2, 0x12, 0x21, 0x41, 0x2e, 0xce, 0x80,
	0x9c, 0xc4, 0xcc, 0xbc, 0x90, 0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x25, 0x0b,
	0x2e, 0xe1, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0xa2, 0xe0, 0x92, 0xc4, 0xa2, 0x92, 0xd2, 0x02, 0xbf,
	0xfc, 0x92, 0xcc, 0xe4, 0x54, 0x21, 0x7e, 0x2e, 0x76, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62,
	0x88, 0x3a, 0x21, 0x01, 0x2e, 0x8e, 0x00, 0x90, 0xb9, 0xc9, 0xf9, 0x39, 0x12, 0x4c, 0x20, 0x91,
	0x24, 0x36, 0xb0, 0x3d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x65, 0x5e, 0x00, 0x82,
	0x00, 0x00, 0x00,
}
