// Code generated by protoc-gen-go.
// source: gcloud/gcloud.proto
// DO NOT EDIT!

/*
Package gcloud is a generated protocol buffer package.

It is generated from these files:
	gcloud/gcloud.proto

It has these top-level messages:
	CreateDevInstanceRequest
	CreateDevInstanceResponse
	StopRequest
	StopResponse
	Instance
	ListInstancesRequest
	ListInstancesResponse
*/
package gcloud

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

type CreateDevInstanceRequest struct {
	Source *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
}

func (m *CreateDevInstanceRequest) Reset()                    { *m = CreateDevInstanceRequest{} }
func (m *CreateDevInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDevInstanceRequest) ProtoMessage()               {}
func (*CreateDevInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateDevInstanceRequest) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type CreateDevInstanceResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *CreateDevInstanceResponse) Reset()                    { *m = CreateDevInstanceResponse{} }
func (m *CreateDevInstanceResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDevInstanceResponse) ProtoMessage()               {}
func (*CreateDevInstanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateDevInstanceResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type StopRequest struct {
	Source   *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Instance string           `protobuf:"bytes,2,opt,name=instance" json:"instance,omitempty"`
	Zone     string           `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StopRequest) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type StopResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *StopResponse) Reset()                    { *m = StopResponse{} }
func (m *StopResponse) String() string            { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()               {}
func (*StopResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StopResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type Instance struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Zone   string `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListInstancesRequest struct {
	Source    *operator.Source `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ProjectId string           `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *ListInstancesRequest) Reset()                    { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()               {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListInstancesRequest) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type ListInstancesResponse struct {
	Output  *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
	Source  *operator.Source `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Objects []*Instance      `protobuf:"bytes,3,rep,name=objects" json:"objects,omitempty"`
}

func (m *ListInstancesResponse) Reset()                    { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()               {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListInstancesResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *ListInstancesResponse) GetSource() *operator.Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ListInstancesResponse) GetObjects() []*Instance {
	if m != nil {
		return m.Objects
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDevInstanceRequest)(nil), "gcloud.CreateDevInstanceRequest")
	proto.RegisterType((*CreateDevInstanceResponse)(nil), "gcloud.CreateDevInstanceResponse")
	proto.RegisterType((*StopRequest)(nil), "gcloud.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "gcloud.StopResponse")
	proto.RegisterType((*Instance)(nil), "gcloud.Instance")
	proto.RegisterType((*ListInstancesRequest)(nil), "gcloud.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "gcloud.ListInstancesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for GcloudService service

type GcloudServiceClient interface {
	// Provision a development instance using the configured image.
	CreateDevInstance(ctx context.Context, in *CreateDevInstanceRequest, opts ...grpc.CallOption) (*CreateDevInstanceResponse, error)
	// List all instances under the configured project.
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Stop a running instance.
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type gcloudServiceClient struct {
	cc *grpc.ClientConn
}

func NewGcloudServiceClient(cc *grpc.ClientConn) GcloudServiceClient {
	return &gcloudServiceClient{cc}
}

func (c *gcloudServiceClient) CreateDevInstance(ctx context.Context, in *CreateDevInstanceRequest, opts ...grpc.CallOption) (*CreateDevInstanceResponse, error) {
	out := new(CreateDevInstanceResponse)
	err := grpc.Invoke(ctx, "/gcloud.GcloudService/CreateDevInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcloudServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := grpc.Invoke(ctx, "/gcloud.GcloudService/ListInstances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gcloudServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := grpc.Invoke(ctx, "/gcloud.GcloudService/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GcloudService service

type GcloudServiceServer interface {
	// Provision a development instance using the configured image.
	CreateDevInstance(context.Context, *CreateDevInstanceRequest) (*CreateDevInstanceResponse, error)
	// List all instances under the configured project.
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Stop a running instance.
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

func RegisterGcloudServiceServer(s *grpc.Server, srv GcloudServiceServer) {
	s.RegisterService(&_GcloudService_serviceDesc, srv)
}

func _GcloudService_CreateDevInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDevInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcloudServiceServer).CreateDevInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gcloud.GcloudService/CreateDevInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcloudServiceServer).CreateDevInstance(ctx, req.(*CreateDevInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcloudService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcloudServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gcloud.GcloudService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcloudServiceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GcloudService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GcloudServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gcloud.GcloudService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GcloudServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GcloudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcloud.GcloudService",
	HandlerType: (*GcloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevInstance",
			Handler:    _GcloudService_CreateDevInstance_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _GcloudService_ListInstances_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _GcloudService_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x4d, 0x81, 0x20, 0x0c, 0x42, 0x74, 0x41, 0x53, 0x1a, 0x49, 0xb0, 0x27, 0xe2, 0x01, 0x23,
	0x5e, 0xbc, 0x8b, 0x31, 0x24, 0x24, 0x26, 0xe5, 0x62, 0xbc, 0x90, 0x52, 0x36, 0xa4, 0x46, 0xbb,
	0xb5, 0x3b, 0xe5, 0xe0, 0xd1, 0x9f, 0xe0, 0xc5, 0x7f, 0xea, 0xd9, 0x76, 0x3f, 0x9a, 0x22, 0xa0,
	0xe9, 0xa9, 0xbb, 0xf3, 0xf1, 0xde, 0x9b, 0xe9, 0x5b, 0x68, 0xaf, 0xbc, 0x17, 0x16, 0x2f, 0x2f,
	0xe5, 0x67, 0x18, 0x46, 0x0c, 0x19, 0xa9, 0xca, 0x9b, 0xd5, 0x62, 0x21, 0x8d, 0x5c, 0x64, 0x91,
	0x8c, 0xdb, 0x63, 0x30, 0x6f, 0x23, 0xea, 0x22, 0x1d, 0xd3, 0xf5, 0x24, 0xe0, 0xe8, 0x06, 0x1e,
	0x75, 0xe8, 0x5b, 0x4c, 0x39, 0x92, 0x01, 0x54, 0x39, 0x8b, 0x23, 0x8f, 0x9a, 0x46, 0xdf, 0x18,
	0x34, 0x46, 0x47, 0xc3, 0xac, 0x79, 0x26, 0xe2, 0x8e, 0xca, 0xdb, 0x77, 0xd0, 0xdd, 0x81, 0xc2,
	0x43, 0x16, 0x70, 0x9a, 0xc2, 0xb0, 0x18, 0xc3, 0x18, 0xb7, 0x61, 0x1e, 0x44, 0xdc, 0x51, 0x79,
	0x7b, 0x05, 0x8d, 0x19, 0xb2, 0xb0, 0x30, 0x3f, 0xb1, 0xa0, 0xe6, 0x2b, 0x5a, 0xb3, 0x94, 0xd4,
	0xd6, 0x9d, 0xec, 0x4e, 0x08, 0x54, 0xde, 0x59, 0x40, 0xcd, 0xb2, 0x88, 0x8b, 0xb3, 0x7d, 0x03,
	0x87, 0x92, 0xa8, 0xb0, 0xc4, 0x27, 0xa8, 0xe9, 0x01, 0x49, 0x0b, 0x4a, 0xfe, 0x52, 0x74, 0xd4,
	0x9d, 0xe4, 0x94, 0x32, 0x05, 0xee, 0xab, 0x56, 0x20, 0xce, 0xe4, 0x34, 0x99, 0x01, 0x5d, 0x8c,
	0xb9, 0xe2, 0x57, 0xb7, 0x4c, 0x55, 0x25, 0xa7, 0x6a, 0x0e, 0x9d, 0xa9, 0xcf, 0x51, 0xe3, 0xf3,
	0xe2, 0x7b, 0xe8, 0x01, 0x24, 0xbf, 0xf5, 0x99, 0x7a, 0x38, 0x4f, 0x94, 0x49, 0x1d, 0x75, 0x15,
	0x99, 0x2c, 0xed, 0x2f, 0x03, 0x4e, 0x7e, 0x31, 0x14, 0x5d, 0x40, 0x4e, 0x4c, 0xe9, 0x1f, 0x31,
	0x17, 0x70, 0xc0, 0x16, 0x29, 0x73, 0x3a, 0x7b, 0x59, 0x94, 0x2a, 0x4b, 0x66, 0x16, 0xd1, 0x05,
	0xa3, 0x6f, 0x03, 0x9a, 0xf7, 0x22, 0x39, 0xa3, 0xd1, 0xda, 0x4f, 0xba, 0x1f, 0xe1, 0x78, 0xcb,
	0x52, 0xa4, 0xaf, 0x11, 0xf6, 0x79, 0xd6, 0x3a, 0xff, 0xa3, 0x42, 0xcd, 0x3a, 0x85, 0xe6, 0xc6,
	0x12, 0xc8, 0x99, 0xee, 0xd9, 0xb5, 0x7d, 0xab, 0xb7, 0x27, 0xab, 0xd0, 0xae, 0xa0, 0x92, 0x5a,
	0x89, 0xb4, 0x75, 0x59, 0xce, 0xc1, 0x56, 0x67, 0x33, 0x28, 0x5b, 0x2c, 0xf8, 0xfc, 0xe8, 0xaa,
	0xf7, 0xb8, 0xa8, 0x8a, 0x67, 0x78, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x74, 0x97, 0xee, 0x12,
	0xb5, 0x03, 0x00, 0x00,
}
