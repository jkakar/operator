// Code generated by protoc-gen-go.
// source: testing/ping.proto
// DO NOT EDIT!

/*
Package operatortesting is a generated protocol buffer package.

It is generated from these files:
	testing/ping.proto

It has these top-level messages:
	PingerConfig
	PingRequest
*/
package operatortesting

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import operator "github.com/sr/operator"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PingerConfig struct {
}

func (m *PingerConfig) Reset()                    { *m = PingerConfig{} }
func (m *PingerConfig) String() string            { return proto.CompactTextString(m) }
func (*PingerConfig) ProtoMessage()               {}
func (*PingerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingRequest struct {
	Source *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func init() {
	proto.RegisterType((*PingerConfig)(nil), "testing.PingerConfig")
	proto.RegisterType((*PingRequest)(nil), "testing.PingRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Pinger service

type PingerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*operator.Response, error)
}

type pingerClient struct {
	cc *grpc.ClientConn
}

func NewPingerClient(cc *grpc.ClientConn) PingerClient {
	return &pingerClient{cc}
}

func (c *pingerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*operator.Response, error) {
	out := new(operator.Response)
	err := grpc.Invoke(ctx, "/testing.Pinger/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pinger service

type PingerServer interface {
	Ping(context.Context, *PingRequest) (*operator.Response, error)
}

func RegisterPingerServer(s *grpc.Server, srv PingerServer) {
	s.RegisterService(&_Pinger_serviceDesc, srv)
}

func _Pinger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testing.Pinger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pinger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testing.Pinger",
	HandlerType: (*PingerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Pinger_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x49, 0x2d, 0x2e,
	0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x00, 0x12, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50,
	0x31, 0x29, 0xbe, 0xfc, 0x82, 0xd4, 0xa2, 0xc4, 0x92, 0xfc, 0x22, 0x88, 0x84, 0x12, 0x1f, 0x17,
	0x4f, 0x00, 0x50, 0x3c, 0xb5, 0xc8, 0x39, 0x3f, 0x2f, 0x2d, 0x33, 0x5d, 0xc9, 0x9c, 0x8b, 0x1b,
	0xc4, 0x0f, 0x4a, 0x2d, 0x2c, 0x05, 0xea, 0x10, 0xd2, 0xe0, 0x62, 0x2b, 0xce, 0x2f, 0x2d, 0x4a,
	0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd0, 0x83, 0xeb, 0x0f, 0x06, 0x8b, 0x07,
	0x41, 0xe5, 0x8d, 0x5c, 0xb8, 0xd8, 0x20, 0x06, 0x09, 0x19, 0x70, 0xb1, 0x80, 0x58, 0x42, 0x22,
	0x7a, 0x50, 0x4b, 0xf5, 0x90, 0x4c, 0x94, 0x12, 0x42, 0x98, 0x10, 0x94, 0x5a, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0x2a, 0xc5, 0x31, 0xa9, 0x49, 0x92, 0x05, 0xe4, 0x5a, 0x27, 0xc1, 0x28, 0x7e, 0x98,
	0x34, 0x54, 0x73, 0x12, 0x1b, 0xd8, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x83,
	0xd5, 0x5f, 0xd7, 0x00, 0x00, 0x00,
}
