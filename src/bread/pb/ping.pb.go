// Code generated by protoc-gen-go.
// source: pb/ping.proto
// DO NOT EDIT!

package breadpb

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

type OtpRequest struct {
	Request *operator.Request `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *OtpRequest) Reset()                    { *m = OtpRequest{} }
func (m *OtpRequest) String() string            { return proto.CompactTextString(m) }
func (*OtpRequest) ProtoMessage()               {}
func (*OtpRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *OtpRequest) GetRequest() *operator.Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type SalesforceAuthRequest struct {
	Request *operator.Request `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *SalesforceAuthRequest) Reset()                    { *m = SalesforceAuthRequest{} }
func (m *SalesforceAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*SalesforceAuthRequest) ProtoMessage()               {}
func (*SalesforceAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SalesforceAuthRequest) GetRequest() *operator.Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type PingRequest struct {
	Request *operator.Request `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Arg1    string            `protobuf:"bytes,2,opt,name=arg1" json:"arg1,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *PingRequest) GetRequest() *operator.Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type SlowLorisRequest struct {
	Request *operator.Request `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Wait    string            `protobuf:"bytes,2,opt,name=wait" json:"wait,omitempty"`
}

func (m *SlowLorisRequest) Reset()                    { *m = SlowLorisRequest{} }
func (m *SlowLorisRequest) String() string            { return proto.CompactTextString(m) }
func (*SlowLorisRequest) ProtoMessage()               {}
func (*SlowLorisRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *SlowLorisRequest) GetRequest() *operator.Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type WhoamiRequest struct {
	Request *operator.Request `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *WhoamiRequest) Reset()                    { *m = WhoamiRequest{} }
func (m *WhoamiRequest) String() string            { return proto.CompactTextString(m) }
func (*WhoamiRequest) ProtoMessage()               {}
func (*WhoamiRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *WhoamiRequest) GetRequest() *operator.Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*OtpRequest)(nil), "bread.OtpRequest")
	proto.RegisterType((*SalesforceAuthRequest)(nil), "bread.SalesforceAuthRequest")
	proto.RegisterType((*PingRequest)(nil), "bread.PingRequest")
	proto.RegisterType((*SlowLorisRequest)(nil), "bread.SlowLorisRequest")
	proto.RegisterType((*WhoamiRequest)(nil), "bread.WhoamiRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Ping service

type PingClient interface {
	// Test OTP verification
	Otp(ctx context.Context, in *OtpRequest, opts ...grpc.CallOption) (*operator.Response, error)
	// Test authentication via Salesforce Authenticator push notification
	SalesforceAuth(ctx context.Context, in *SalesforceAuthRequest, opts ...grpc.CallOption) (*operator.Response, error)
	// Reply with PONG if everything is working
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*operator.Response, error)
	// Trigger a slow request, for testing timeout handling
	SlowLoris(ctx context.Context, in *SlowLorisRequest, opts ...grpc.CallOption) (*operator.Response, error)
	// Reply with the email of the current authenticated user
	Whoami(ctx context.Context, in *WhoamiRequest, opts ...grpc.CallOption) (*operator.Response, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) Otp(ctx context.Context, in *OtpRequest, opts ...grpc.CallOption) (*operator.Response, error) {
	out := new(operator.Response)
	err := grpc.Invoke(ctx, "/bread.Ping/Otp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingClient) SalesforceAuth(ctx context.Context, in *SalesforceAuthRequest, opts ...grpc.CallOption) (*operator.Response, error) {
	out := new(operator.Response)
	err := grpc.Invoke(ctx, "/bread.Ping/SalesforceAuth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*operator.Response, error) {
	out := new(operator.Response)
	err := grpc.Invoke(ctx, "/bread.Ping/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingClient) SlowLoris(ctx context.Context, in *SlowLorisRequest, opts ...grpc.CallOption) (*operator.Response, error) {
	out := new(operator.Response)
	err := grpc.Invoke(ctx, "/bread.Ping/SlowLoris", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingClient) Whoami(ctx context.Context, in *WhoamiRequest, opts ...grpc.CallOption) (*operator.Response, error) {
	out := new(operator.Response)
	err := grpc.Invoke(ctx, "/bread.Ping/Whoami", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ping service

type PingServer interface {
	// Test OTP verification
	Otp(context.Context, *OtpRequest) (*operator.Response, error)
	// Test authentication via Salesforce Authenticator push notification
	SalesforceAuth(context.Context, *SalesforceAuthRequest) (*operator.Response, error)
	// Reply with PONG if everything is working
	Ping(context.Context, *PingRequest) (*operator.Response, error)
	// Trigger a slow request, for testing timeout handling
	SlowLoris(context.Context, *SlowLorisRequest) (*operator.Response, error)
	// Reply with the email of the current authenticated user
	Whoami(context.Context, *WhoamiRequest) (*operator.Response, error)
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_Otp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Otp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Ping/Otp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Otp(ctx, req.(*OtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ping_SalesforceAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalesforceAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).SalesforceAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Ping/SalesforceAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).SalesforceAuth(ctx, req.(*SalesforceAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ping_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Ping/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ping_SlowLoris_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlowLorisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).SlowLoris(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Ping/SlowLoris",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).SlowLoris(ctx, req.(*SlowLorisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ping_Whoami_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoamiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Whoami(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Ping/Whoami",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Whoami(ctx, req.(*WhoamiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bread.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Otp",
			Handler:    _Ping_Otp_Handler,
		},
		{
			MethodName: "SalesforceAuth",
			Handler:    _Ping_SalesforceAuth_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Ping_Ping_Handler,
		},
		{
			MethodName: "SlowLoris",
			Handler:    _Ping_SlowLoris_Handler,
		},
		{
			MethodName: "Whoami",
			Handler:    _Ping_Whoami_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("pb/ping.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0xdf, 0xbe, 0x2d, 0xb9, 0xa5, 0x62, 0x2f, 0x8a, 0xa5, 0xb8, 0x28, 0x01, 0xa1,
	0xa0, 0x4c, 0x30, 0x6e, 0x14, 0xdc, 0x58, 0x5c, 0x8a, 0x95, 0x64, 0x21, 0xb8, 0x9b, 0xd4, 0x31,
	0x19, 0x68, 0x73, 0xc7, 0x99, 0x09, 0xfd, 0x27, 0xfe, 0x5e, 0x69, 0x3e, 0xfa, 0x21, 0x59, 0x65,
	0x77, 0x16, 0xe7, 0x39, 0x97, 0x3c, 0x19, 0x18, 0xaa, 0xd8, 0x57, 0x32, 0x4b, 0x98, 0xd2, 0x64,
	0x09, 0xff, 0xc7, 0x5a, 0xf0, 0xcf, 0xc9, 0x55, 0x22, 0x6d, 0x9a, 0xc7, 0x6c, 0x49, 0x6b, 0xdf,
	0x68, 0x9f, 0x94, 0xd0, 0xdc, 0xd2, 0x3e, 0x94, 0x6d, 0xef, 0x01, 0x60, 0x61, 0x55, 0x28, 0xbe,
	0x73, 0x61, 0x2c, 0x5e, 0x43, 0x5f, 0x97, 0x71, 0xec, 0x4c, 0x9d, 0xd9, 0x20, 0x18, 0xb1, 0x5d,
	0xbf, 0xea, 0x84, 0x75, 0xc3, 0x7b, 0x86, 0xf3, 0x88, 0xaf, 0x84, 0xf9, 0x22, 0xbd, 0x14, 0x4f,
	0xb9, 0x4d, 0x5b, 0xad, 0xbc, 0xc2, 0xe0, 0x4d, 0x66, 0x49, 0x1b, 0x16, 0x11, 0xba, 0x5c, 0x27,
	0xb7, 0xe3, 0xce, 0xd4, 0x99, 0xb9, 0x61, 0x91, 0xbd, 0x08, 0x4e, 0xa3, 0x15, 0x6d, 0x5e, 0x48,
	0x4b, 0xd3, 0x76, 0x74, 0xc3, 0xa5, 0xad, 0x47, 0xb7, 0xd9, 0x7b, 0x84, 0xe1, 0x7b, 0x4a, 0x7c,
	0x2d, 0xdb, 0x2c, 0x06, 0x3f, 0x1d, 0xe8, 0x6e, 0xbf, 0x11, 0x6f, 0xe0, 0xdf, 0xc2, 0x2a, 0x1c,
	0xb1, 0xe2, 0x17, 0xb1, 0xbd, 0xf8, 0x09, 0x1e, 0xe2, 0x46, 0x51, 0x66, 0x04, 0xce, 0xe1, 0xe4,
	0xd8, 0x2f, 0x5e, 0x56, 0x60, 0xa3, 0xf6, 0xc6, 0x0d, 0x56, 0x5d, 0xc6, 0x8a, 0x3c, 0x50, 0xdd,
	0xd8, 0xbf, 0x07, 0x77, 0x67, 0x0f, 0x2f, 0xea, 0x73, 0x7f, 0x7c, 0x36, 0x92, 0x01, 0xf4, 0x4a,
	0x45, 0x78, 0x56, 0x61, 0x47, 0xc6, 0x9a, 0x98, 0xb9, 0xfb, 0xd1, 0x2f, 0xaa, 0x2a, 0x8e, 0x7b,
	0xc5, 0x73, 0xbc, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x69, 0x02, 0x00, 0xaf, 0xcd, 0x02, 0x00,
	0x00,
}
