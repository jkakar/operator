// Code generated by protoc-gen-go.
// source: hal/hal9000.proto
// DO NOT EDIT!

/*
Package breadhal is a generated protocol buffer package.

It is generated from these files:
	hal/hal9000.proto

It has these top-level messages:
	Message
	Response
*/
package breadhal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Text      string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	Room      string `protobuf:"bytes,3,opt,name=room" json:"room,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Match bool `protobuf:"varint,1,opt,name=match" json:"match,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Message)(nil), "hal.Message")
	proto.RegisterType((*Response)(nil), "hal.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Robot service

type RobotClient interface {
	IsMatch(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
	Dispatch(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type robotClient struct {
	cc *grpc.ClientConn
}

func NewRobotClient(cc *grpc.ClientConn) RobotClient {
	return &robotClient{cc}
}

func (c *robotClient) IsMatch(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/hal.Robot/IsMatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotClient) Dispatch(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/hal.Robot/Dispatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Robot service

type RobotServer interface {
	IsMatch(context.Context, *Message) (*Response, error)
	Dispatch(context.Context, *Message) (*Response, error)
}

func RegisterRobotServer(s *grpc.Server, srv RobotServer) {
	s.RegisterService(&_Robot_serviceDesc, srv)
}

func _Robot_IsMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServer).IsMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hal.Robot/IsMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServer).IsMatch(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robot_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hal.Robot/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServer).Dispatch(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Robot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hal.Robot",
	HandlerType: (*RobotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsMatch",
			Handler:    _Robot_IsMatch_Handler,
		},
		{
			MethodName: "Dispatch",
			Handler:    _Robot_Dispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("hal/hal9000.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcc, 0xd1,
	0xcf, 0x48, 0xcc, 0xb1, 0x34, 0x30, 0x30, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce,
	0x48, 0xcc, 0x51, 0x0a, 0xe0, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2,
	0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x64,
	0xb9, 0xb8, 0x4a, 0x8b, 0x53, 0x8b, 0xe2, 0x53, 0x73, 0x13, 0x33, 0x73, 0x24, 0x98, 0xc0, 0x32,
	0x9c, 0x20, 0x11, 0x57, 0x90, 0x00, 0x48, 0x4b, 0x51, 0x7e, 0x7e, 0xae, 0x04, 0x33, 0x44, 0x0b,
	0x88, 0xad, 0xa4, 0xc0, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc2,
	0xc5, 0x9a, 0x9b, 0x58, 0x92, 0x9c, 0x01, 0x36, 0x93, 0x23, 0x08, 0xc2, 0x31, 0x8a, 0xe0, 0x62,
	0x0d, 0xca, 0x4f, 0xca, 0x2f, 0x11, 0x52, 0xe3, 0x62, 0xf7, 0x2c, 0xf6, 0x05, 0x89, 0x09, 0xf1,
	0xe8, 0x65, 0x24, 0xe6, 0xe8, 0x41, 0x9d, 0x22, 0xc5, 0x0b, 0xe6, 0xc1, 0x8d, 0x51, 0xe7, 0xe2,
	0x70, 0xc9, 0x2c, 0x2e, 0x20, 0xa8, 0xd0, 0x89, 0x2b, 0x8a, 0x23, 0xa9, 0x28, 0x35, 0x31, 0x25,
	0x23, 0x31, 0x27, 0x89, 0x0d, 0xec, 0x4b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x46,
	0xff, 0xe3, 0xfa, 0x00, 0x00, 0x00,
}
