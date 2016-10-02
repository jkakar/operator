// Code generated by protoc-gen-go.
// source: pb/bread.proto
// DO NOT EDIT!

/*
Package breadpb is a generated protocol buffer package.

It is generated from these files:
	pb/bread.proto
	pb/deploy.proto
	pb/ping.proto

It has these top-level messages:
	OperatorRequest
	OperatorMessage
	HTTPRequest
	ServerStartupNotice
	ServiceStartupError
	ListTargetsRequest
	ListBuildsRequest
	TriggerRequest
	OtpRequest
	PingRequest
	SlowLorisRequest
	WhoamiRequest
*/
package breadpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import operator "github.com/sr/operator"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OperatorRequest struct {
	Event   string            `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	Request *operator.Request `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	Message *OperatorMessage  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Error   string            `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *OperatorRequest) Reset()                    { *m = OperatorRequest{} }
func (m *OperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*OperatorRequest) ProtoMessage()               {}
func (*OperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OperatorRequest) GetRequest() *operator.Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *OperatorRequest) GetMessage() *OperatorMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type OperatorMessage struct {
	Source *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Text   string           `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Html   string           `protobuf:"bytes,3,opt,name=html" json:"html,omitempty"`
}

func (m *OperatorMessage) Reset()                    { *m = OperatorMessage{} }
func (m *OperatorMessage) String() string            { return proto.CompactTextString(m) }
func (*OperatorMessage) ProtoMessage()               {}
func (*OperatorMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OperatorMessage) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

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
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

type ServerStartupNotice struct {
	Address    string   `protobuf:"bytes,1,opt,name=Address,json=address" json:"Address,omitempty"`
	Protocol   string   `protobuf:"bytes,2,opt,name=Protocol,json=protocol" json:"Protocol,omitempty"`
	Services   []string `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
	HalAddress string   `protobuf:"bytes,4,opt,name=HalAddress,json=halAddress" json:"HalAddress,omitempty"`
}

func (m *ServerStartupNotice) Reset()                    { *m = ServerStartupNotice{} }
func (m *ServerStartupNotice) String() string            { return proto.CompactTextString(m) }
func (*ServerStartupNotice) ProtoMessage()               {}
func (*ServerStartupNotice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ServiceStartupError struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ServiceStartupError) Reset()                    { *m = ServiceStartupError{} }
func (m *ServiceStartupError) String() string            { return proto.CompactTextString(m) }
func (*ServiceStartupError) ProtoMessage()               {}
func (*ServiceStartupError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*OperatorRequest)(nil), "bread.OperatorRequest")
	proto.RegisterType((*OperatorMessage)(nil), "bread.OperatorMessage")
	proto.RegisterType((*HTTPRequest)(nil), "bread.HTTPRequest")
	proto.RegisterType((*ServerStartupNotice)(nil), "bread.ServerStartupNotice")
	proto.RegisterType((*ServiceStartupError)(nil), "bread.ServiceStartupError")
}

func init() { proto.RegisterFile("pb/bread.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x26, 0x8e, 0xc7, 0xa2, 0x94, 0x05, 0x55, 0x26, 0x12, 0x25, 0xf2, 0xc9, 0x12,
	0x92, 0x83, 0x52, 0x21, 0x55, 0xdc, 0xf8, 0x88, 0xd4, 0x0b, 0x50, 0x36, 0x3d, 0x71, 0x41, 0x6b,
	0x7b, 0x48, 0x22, 0x9c, 0xac, 0xbb, 0xbb, 0x8e, 0xc8, 0x2f, 0xe0, 0x1f, 0x70, 0xe2, 0xc7, 0x56,
	0xfb, 0x65, 0xe7, 0x36, 0x6f, 0xe6, 0xcd, 0xdb, 0x37, 0x33, 0x0b, 0xe7, 0x4d, 0x31, 0x2f, 0x04,
	0xb2, 0x2a, 0x6f, 0x04, 0x57, 0x9c, 0x8c, 0x0c, 0x98, 0x5e, 0xad, 0x39, 0x5f, 0xd7, 0x38, 0x37,
	0xc9, 0xa2, 0xfd, 0x35, 0xaf, 0x5a, 0xc1, 0xd4, 0x96, 0xef, 0x2d, 0x6d, 0x7a, 0xce, 0x1b, 0x14,
	0x4c, 0x71, 0x61, 0x71, 0xfa, 0x3f, 0x80, 0xa7, 0xdf, 0x5c, 0x8a, 0xe2, 0x43, 0x8b, 0x52, 0x91,
	0x17, 0x30, 0xc2, 0x03, 0xee, 0x55, 0x12, 0xcc, 0x82, 0x2c, 0xa2, 0x16, 0x90, 0x37, 0x10, 0x0a,
	0x4b, 0x48, 0x06, 0xb3, 0x20, 0x8b, 0x17, 0xcf, 0xf2, 0x4e, 0xcb, 0x75, 0x52, 0xcf, 0x20, 0x6f,
	0x21, 0xdc, 0xa1, 0x94, 0x6c, 0x8d, 0xc9, 0xd0, 0x90, 0x2f, 0x73, 0x6b, 0xd6, 0xbf, 0xf5, 0xc5,
	0x56, 0xa9, 0xa7, 0x99, 0x47, 0x85, 0xe0, 0x22, 0x39, 0x73, 0x8f, 0x6a, 0x90, 0x96, 0xbd, 0x3b,
	0xd7, 0x41, 0x32, 0x18, 0x4b, 0xde, 0x8a, 0x12, 0x8d, 0xbd, 0x78, 0x71, 0xd1, 0xdb, 0x58, 0x99,
	0x3c, 0x75, 0x75, 0x42, 0xe0, 0x4c, 0xe1, 0x1f, 0x6b, 0x37, 0xa2, 0x26, 0xd6, 0xb9, 0x8d, 0xda,
	0xd5, 0xc6, 0x55, 0x44, 0x4d, 0x9c, 0xfe, 0x1b, 0x40, 0x7c, 0x7b, 0x7f, 0x7f, 0xe7, 0xe7, 0xbf,
	0x84, 0xf1, 0x0e, 0xd5, 0x86, 0x57, 0x6e, 0x01, 0x0e, 0xe9, 0xde, 0x86, 0xa9, 0x8d, 0xd7, 0xd3,
	0x31, 0xb9, 0x86, 0xd1, 0x43, 0x8b, 0xe2, 0x98, 0x0c, 0x67, 0xc3, 0x2c, 0x5e, 0xbc, 0x72, 0x63,
	0x9e, 0xc8, 0xe5, 0xdf, 0x75, 0x7d, 0xb9, 0x57, 0xe2, 0x48, 0x2d, 0x97, 0xbc, 0x86, 0x58, 0x2a,
	0xa6, 0x5a, 0xf9, 0xb3, 0xe4, 0x15, 0x26, 0xe1, 0x2c, 0xc8, 0x9e, 0x50, 0xb0, 0xa9, 0x4f, 0xbc,
	0x42, 0xf2, 0x0e, 0x26, 0xfe, 0x6e, 0xc9, 0xc4, 0x4c, 0xf9, 0x32, 0xb7, 0x87, 0xcd, 0xfd, 0x61,
	0xf3, 0xcf, 0x8e, 0x40, 0x3b, 0x6a, 0xbf, 0xc3, 0xe8, 0x64, 0x87, 0xd3, 0x1b, 0x80, 0xde, 0x02,
	0xb9, 0x80, 0xe1, 0x6f, 0x3c, 0xba, 0xc9, 0x74, 0xa8, 0xbb, 0x0e, 0xac, 0x6e, 0xd1, 0xcd, 0x65,
	0xc1, 0xfb, 0xc1, 0x4d, 0x90, 0xfe, 0x0d, 0xe0, 0xf9, 0x0a, 0xc5, 0x01, 0xc5, 0x4a, 0x31, 0xa1,
	0xda, 0xe6, 0x2b, 0x57, 0xdb, 0x12, 0x49, 0x02, 0xe1, 0x87, 0xaa, 0x12, 0x28, 0xa5, 0xd3, 0x09,
	0x99, 0x85, 0x64, 0x0a, 0x93, 0x3b, 0x6d, 0xb0, 0xe4, 0xb5, 0x93, 0x9b, 0x34, 0x0e, 0xeb, 0x9a,
	0x44, 0x71, 0xd8, 0x96, 0x28, 0xcd, 0xb6, 0x22, 0xda, 0x61, 0x72, 0x05, 0x70, 0xcb, 0x6a, 0x2f,
	0x6a, 0xbf, 0x00, 0x6c, 0xba, 0x4c, 0xba, 0xb4, 0x46, 0xb6, 0x25, 0x3a, 0x27, 0x4b, 0x3d, 0x9a,
	0x36, 0xe2, 0x24, 0xbc, 0x11, 0x07, 0xfb, 0x55, 0x0c, 0x4e, 0x56, 0xf1, 0x31, 0xfa, 0x11, 0x9a,
	0xfb, 0x34, 0x45, 0x31, 0x36, 0xbe, 0xae, 0x1f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x5a, 0x16,
	0x83, 0x48, 0x03, 0x00, 0x00,
}
