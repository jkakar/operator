// Code generated by protoc-gen-go.
// source: httplog.proto
// DO NOT EDIT!

/*
Package bread is a generated protocol buffer package.

It is generated from these files:
	httplog.proto

It has these top-level messages:
	HTTPRequest
*/
package bread

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HTTPRequest struct {
	Method     string                    `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path       string                    `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Query      map[string]string         `protobuf:"bytes,3,rep,name=query" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	StatusCode uint32                    `protobuf:"varint,7,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Duration   *google_protobuf.Duration `protobuf:"bytes,8,opt,name=duration" json:"duration,omitempty"`
	Error      string                    `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string            { return proto.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HTTPRequest) GetQuery() map[string]string {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *HTTPRequest) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*HTTPRequest)(nil), "bread.HTTPRequest")
}

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x4f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x89, 0xa9, 0xed, 0x84, 0x82, 0x2c, 0x22, 0x6b, 0x41, 0x2d, 0x9e, 0x7a, 0xda, 0x42,
	0x8b, 0x50, 0xbc, 0xaa, 0xe0, 0x51, 0x43, 0xef, 0xb2, 0x31, 0x63, 0x22, 0xc6, 0x4e, 0xdc, 0xcc,
	0x0a, 0xf9, 0x11, 0xbf, 0xd7, 0x64, 0x37, 0xd5, 0xde, 0xe6, 0xbd, 0x79, 0x33, 0xef, 0x3d, 0x98,
	0x96, 0xcc, 0x75, 0x45, 0x85, 0xaa, 0x0d, 0x31, 0x89, 0x38, 0x33, 0xa8, 0xf3, 0xd9, 0x65, 0x41,
	0x54, 0x54, 0xb8, 0x74, 0x64, 0x66, 0xdf, 0x96, 0xb9, 0x35, 0x9a, 0xdf, 0x69, 0xe7, 0x65, 0xd7,
	0x3f, 0x21, 0x24, 0x8f, 0xdb, 0xed, 0x53, 0x8a, 0x5f, 0x16, 0x1b, 0x16, 0x67, 0x30, 0xfa, 0x44,
	0x2e, 0x29, 0x97, 0xc1, 0x3c, 0x58, 0x4c, 0xd2, 0x01, 0x09, 0x01, 0x47, 0xb5, 0xe6, 0x52, 0x86,
	0x8e, 0x75, 0xb3, 0x58, 0x43, 0xdc, 0x1d, 0x99, 0x56, 0x46, 0xf3, 0x68, 0x91, 0xac, 0x2e, 0x94,
	0xb3, 0x54, 0x07, 0xef, 0xd4, 0x73, 0xbf, 0x7f, 0xd8, 0xb1, 0x69, 0x53, 0xaf, 0x15, 0x57, 0x90,
	0x34, 0xac, 0xd9, 0x36, 0x2f, 0xaf, 0x94, 0xa3, 0x3c, 0xee, 0xfe, 0x4d, 0x53, 0xf0, 0xd4, 0x5d,
	0xc7, 0x88, 0x1b, 0x18, 0xef, 0x33, 0xca, 0x71, 0xb7, 0x4d, 0x56, 0xe7, 0xca, 0x97, 0x50, 0xfb,
	0x12, 0xea, 0x7e, 0x10, 0xa4, 0x7f, 0x52, 0x71, 0x0a, 0x31, 0x1a, 0x43, 0x46, 0x4e, 0x5c, 0x42,
	0x0f, 0x66, 0x1b, 0x80, 0xff, 0x08, 0xe2, 0x04, 0xa2, 0x0f, 0x6c, 0x87, 0x66, 0xfd, 0xd8, 0x5f,
	0x7d, 0xeb, 0xca, 0xe2, 0xd0, 0xcb, 0x83, 0xdb, 0x70, 0x13, 0x64, 0x23, 0x67, 0xb6, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x2d, 0x14, 0x1b, 0x57, 0x01, 0x00, 0x00,
}
