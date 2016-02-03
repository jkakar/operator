// Code generated by protoc-gen-go.
// source: services/buildkite/buildkite.proto
// DO NOT EDIT!

/*
Package buildkite is a generated protocol buffer package.

It is generated from these files:
	services/buildkite/buildkite.proto

It has these top-level messages:
	StatusRequest
	StatusResponse
	ListBuildsRequest
	ListBuildsResponse
*/
package buildkite

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import operator "github.com/sr/operator/proto"

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

type StatusRequest struct {
	// Optional project slug.
	Slug string `protobuf:"bytes,1,opt,name=slug" json:"slug,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type ListBuildsRequest struct {
	ProjectSlug string `protobuf:"bytes,1,opt,name=project_slug" json:"project_slug,omitempty"`
}

func (m *ListBuildsRequest) Reset()                    { *m = ListBuildsRequest{} }
func (m *ListBuildsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBuildsRequest) ProtoMessage()               {}
func (*ListBuildsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListBuildsResponse struct {
	Output *operator.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *ListBuildsResponse) Reset()                    { *m = ListBuildsResponse{} }
func (m *ListBuildsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBuildsResponse) ProtoMessage()               {}
func (*ListBuildsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListBuildsResponse) GetOutput() *operator.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "buildkite.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "buildkite.StatusResponse")
	proto.RegisterType((*ListBuildsRequest)(nil), "buildkite.ListBuildsRequest")
	proto.RegisterType((*ListBuildsResponse)(nil), "buildkite.ListBuildsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for BuildkiteService service

type BuildkiteServiceClient interface {
	// List the status of all (i.e. the status of the last build) of one or
	// all projects.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// List the last builds of one or all projects, optionally limited to a
	// project.
	ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error)
}

type buildkiteServiceClient struct {
	cc *grpc.ClientConn
}

func NewBuildkiteServiceClient(cc *grpc.ClientConn) BuildkiteServiceClient {
	return &buildkiteServiceClient{cc}
}

func (c *buildkiteServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/buildkite.BuildkiteService/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildkiteServiceClient) ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error) {
	out := new(ListBuildsResponse)
	err := grpc.Invoke(ctx, "/buildkite.BuildkiteService/ListBuilds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BuildkiteService service

type BuildkiteServiceServer interface {
	// List the status of all (i.e. the status of the last build) of one or
	// all projects.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// List the last builds of one or all projects, optionally limited to a
	// project.
	ListBuilds(context.Context, *ListBuildsRequest) (*ListBuildsResponse, error)
}

func RegisterBuildkiteServiceServer(s *grpc.Server, srv BuildkiteServiceServer) {
	s.RegisterService(&_BuildkiteService_serviceDesc, srv)
}

func _BuildkiteService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BuildkiteServiceServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BuildkiteService_ListBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BuildkiteServiceServer).ListBuilds(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BuildkiteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildkite.BuildkiteService",
	HandlerType: (*BuildkiteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _BuildkiteService_Status_Handler,
		},
		{
			MethodName: "ListBuilds",
			Handler:    _BuildkiteService_ListBuilds_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0xc9, 0xce, 0x2c, 0x49, 0x45, 0xb0,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xe1, 0x02, 0x52, 0x7c, 0xf9, 0x05, 0xa9, 0x45,
	0x89, 0x25, 0xf9, 0x45, 0x10, 0x29, 0x25, 0x59, 0x2e, 0xde, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0xe2,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x1e, 0x2e, 0x96, 0xe2, 0x9c, 0xd2, 0x74, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x4e, 0x25, 0x23, 0x2e, 0x3e, 0x98, 0x74, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x02, 0x17, 0x5b, 0x7e, 0x69, 0x49, 0x41, 0x69, 0x09, 0x58, 0x05, 0xb7, 0x91, 0x80,
	0x1e, 0xdc, 0x44, 0x7f, 0xb0, 0xb8, 0x92, 0x26, 0x97, 0xa0, 0x4f, 0x66, 0x71, 0x89, 0x13, 0xc8,
	0x4e, 0xb8, 0xb1, 0x22, 0x5c, 0x3c, 0x40, 0x0b, 0xb3, 0x52, 0x93, 0x4b, 0xe2, 0x91, 0x8c, 0x37,
	0xe3, 0x12, 0x42, 0x56, 0x4a, 0xac, 0x15, 0x46, 0xab, 0x19, 0xb9, 0x04, 0x9c, 0x60, 0x7e, 0x0a,
	0x86, 0x04, 0x80, 0x90, 0x2d, 0x17, 0x1b, 0xc4, 0xad, 0x42, 0x12, 0x7a, 0x88, 0x10, 0x40, 0xf1,
	0x9d, 0x94, 0x24, 0x16, 0x19, 0xa8, 0xad, 0x9e, 0x5c, 0x5c, 0x08, 0xb7, 0x08, 0xc9, 0x20, 0x29,
	0xc4, 0xf0, 0x8d, 0x94, 0x2c, 0x0e, 0x59, 0x88, 0x51, 0x52, 0xbc, 0x93, 0x9a, 0x24, 0x11, 0x61,
	0x9e, 0xc4, 0x06, 0x0e, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd1, 0xaa, 0xa3,
	0xab, 0x01, 0x00, 0x00,
}