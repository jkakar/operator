// Code generated by protoc-gen-go.
// source: services/controller/controller.proto
// DO NOT EDIT!

/*
Package controller is a generated protocol buffer package.

It is generated from these files:
	services/controller/controller.proto

It has these top-level messages:
	CreateClusterRequest
	DeployRequest
	CreateClusterResponse
	DeployResponse
*/
package controller

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import operator "operator"

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

type CreateClusterRequest struct {
}

func (m *CreateClusterRequest) Reset()                    { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()               {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DeployRequest struct {
	BuildId          string `protobuf:"bytes,1,opt,name=build_id" json:"build_id,omitempty"`
	HubotBuildId     string `protobuf:"bytes,2,opt,name=hubot_build_id" json:"hubot_build_id,omitempty"`
	OperatordBuildId string `protobuf:"bytes,3,opt,name=operatord_build_id" json:"operatord_build_id,omitempty"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateClusterResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *CreateClusterResponse) Reset()                    { *m = CreateClusterResponse{} }
func (m *CreateClusterResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterResponse) ProtoMessage()               {}
func (*CreateClusterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateClusterResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type DeployResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *DeployResponse) Reset()                    { *m = DeployResponse{} }
func (m *DeployResponse) String() string            { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()               {}
func (*DeployResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeployResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateClusterRequest)(nil), "controller.CreateClusterRequest")
	proto.RegisterType((*DeployRequest)(nil), "controller.DeployRequest")
	proto.RegisterType((*CreateClusterResponse)(nil), "controller.CreateClusterResponse")
	proto.RegisterType((*DeployResponse)(nil), "controller.DeployResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Controller service

type ControllerClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*CreateClusterResponse, error)
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployResponse, error)
}

type controllerClient struct {
	cc *grpc.ClientConn
}

func NewControllerClient(cc *grpc.ClientConn) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*CreateClusterResponse, error) {
	out := new(CreateClusterResponse)
	err := grpc.Invoke(ctx, "/controller.Controller/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployResponse, error) {
	out := new(DeployResponse)
	err := grpc.Invoke(ctx, "/controller.Controller/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Controller service

type ControllerServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*CreateClusterResponse, error)
	Deploy(context.Context, *DeployRequest) (*DeployResponse, error)
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ControllerServer).CreateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Controller_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ControllerServer).Deploy(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "controller.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _Controller_CreateCluster_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Controller_Deploy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x0a, 0x41, 0x47, 0x1a, 0xca, 0xa2, 0xd5, 0xee, 0x29, 0x06, 0x0f, 0x9e, 0x52, 0x88,
	0x27, 0x4f, 0x1e, 0xe2, 0x5d, 0x10, 0x3d, 0x97, 0xa6, 0x19, 0x30, 0xb0, 0x74, 0xd6, 0xd9, 0x59,
	0xc1, 0xab, 0x9f, 0xe0, 0xaf, 0xf8, 0x83, 0x62, 0xd2, 0xa4, 0x5b, 0x09, 0xe2, 0x6d, 0x78, 0xef,
	0xcd, 0xbc, 0xc7, 0x1b, 0xb8, 0x72, 0xc8, 0x6f, 0xcd, 0x1a, 0xdd, 0x62, 0x4d, 0x1b, 0x61, 0x32,
	0x06, 0x39, 0x18, 0x73, 0xcb, 0x24, 0xa4, 0x60, 0x87, 0xe8, 0x73, 0xb2, 0xc8, 0x2b, 0x21, 0x5e,
	0xf4, 0x43, 0x27, 0xca, 0x66, 0x70, 0x5a, 0x32, 0xae, 0x04, 0x4b, 0xe3, 0x9d, 0x20, 0x3f, 0xe2,
	0xab, 0x47, 0x27, 0xd9, 0x33, 0x4c, 0xee, 0xd1, 0x1a, 0x7a, 0xdf, 0x02, 0x6a, 0x0a, 0x47, 0x95,
	0x6f, 0x4c, 0xbd, 0x6c, 0xea, 0x8b, 0x28, 0x8d, 0xae, 0x8f, 0xd5, 0x0c, 0x92, 0x17, 0x5f, 0x91,
	0x2c, 0x07, 0xfc, 0xa0, 0xc5, 0x35, 0xa8, 0xde, 0xa4, 0xde, 0x71, 0x87, 0x3f, 0x5c, 0x76, 0x0b,
	0x67, 0xbf, 0xec, 0x9c, 0xa5, 0x8d, 0x43, 0x95, 0x42, 0x4c, 0x5e, 0xac, 0x97, 0xf6, 0xf8, 0x49,
	0x31, 0xcd, 0x87, 0xa0, 0x0f, 0x2d, 0x9e, 0x15, 0x90, 0xf4, 0x89, 0xfe, 0xbb, 0x53, 0x7c, 0x45,
	0x00, 0xe5, 0xd0, 0x82, 0x7a, 0x82, 0xc9, 0x9e, 0xbb, 0x4a, 0xf3, 0xa0, 0xb5, 0xb1, 0x1e, 0xf4,
	0xe5, 0x1f, 0x8a, 0x6d, 0x8c, 0x3b, 0x88, 0xbb, 0x60, 0x6a, 0x1e, 0x8a, 0xf7, 0xea, 0xd3, 0x7a,
	0x8c, 0xea, 0x0e, 0xe8, 0xe4, 0xf3, 0x63, 0x1e, 0x3c, 0xab, 0x8a, 0xdb, 0xd7, 0xdc, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x17, 0xc8, 0x21, 0xe7, 0xe7, 0x01, 0x00, 0x00,
}