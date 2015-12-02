// Code generated by protoc-gen-go.
// source: papertrail/papertrail.proto
// DO NOT EDIT!

/*
Package papertrail is a generated protocol buffer package.

It is generated from these files:
	papertrail/papertrail.proto

It has these top-level messages:
	SearchRequest
	SearchResponse
	LogEvent
*/
package papertrail

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

type SearchRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SearchResponse struct {
	LogEvents []*LogEvent `protobuf:"bytes,1,rep,name=log_events" json:"log_events,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetLogEvents() []*LogEvent {
	if m != nil {
		return m.LogEvents
	}
	return nil
}

type LogEvent struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Source     string `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Program    string `protobuf:"bytes,3,opt,name=program" json:"program,omitempty"`
	LogMessage string `protobuf:"bytes,4,opt,name=log_message" json:"log_message,omitempty"`
}

func (m *LogEvent) Reset()                    { *m = LogEvent{} }
func (m *LogEvent) String() string            { return proto.CompactTextString(m) }
func (*LogEvent) ProtoMessage()               {}
func (*LogEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*SearchRequest)(nil), "papertrail.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "papertrail.SearchResponse")
	proto.RegisterType((*LogEvent)(nil), "papertrail.LogEvent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for PapertrailService service

type PapertrailServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type papertrailServiceClient struct {
	cc *grpc.ClientConn
}

func NewPapertrailServiceClient(cc *grpc.ClientConn) PapertrailServiceClient {
	return &papertrailServiceClient{cc}
}

func (c *papertrailServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/papertrail.PapertrailService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PapertrailService service

type PapertrailServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterPapertrailServiceServer(s *grpc.Server, srv PapertrailServiceServer) {
	s.RegisterService(&_PapertrailService_serviceDesc, srv)
}

func _PapertrailService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PapertrailServiceServer).Search(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PapertrailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "papertrail.PapertrailService",
	HandlerType: (*PapertrailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _PapertrailService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0x4f, 0x4b, 0x87, 0x40,
	0x10, 0x45, 0x2d, 0xab, 0x11, 0x8d, 0xb6, 0x0e, 0x9b, 0x41, 0x84, 0x27, 0x4f, 0x06, 0x76, 0xeb,
	0xd2, 0xa9, 0x9b, 0x87, 0xc8, 0xee, 0xb1, 0xd9, 0xb0, 0x09, 0xea, 0x6e, 0xb3, 0xab, 0xd0, 0xb7,
	0xcf, 0x3f, 0x98, 0x06, 0xbf, 0xdb, 0x9b, 0xf7, 0x86, 0x79, 0xef, 0x0d, 0xdc, 0x68, 0xa1, 0x91,
	0x2c, 0x89, 0xba, 0xb9, 0xdf, 0x60, 0xa6, 0x49, 0x59, 0xc5, 0x60, 0x63, 0x92, 0x5b, 0x08, 0x4b,
	0x14, 0x54, 0x7d, 0xbd, 0xe2, 0x77, 0x8f, 0xc6, 0xb2, 0x10, 0x8e, 0x47, 0x40, 0x3f, 0xdc, 0xb9,
	0x73, 0xd2, 0xb3, 0xe4, 0x11, 0xa2, 0x55, 0x37, 0x5a, 0x75, 0x06, 0x59, 0x0a, 0xd0, 0x28, 0xf9,
	0x8e, 0x03, 0x76, 0xd6, 0x8c, 0x5b, 0x5e, 0x1a, 0xe4, 0x57, 0xd9, 0xce, 0xa4, 0x50, 0xf2, 0x79,
	0x12, 0x93, 0x02, 0x4e, 0x57, 0xcc, 0x00, 0xdc, 0xfa, 0x73, 0xb9, 0xc9, 0x22, 0xf0, 0x8d, 0xea,
	0xa9, 0x42, 0xee, 0xce, 0xf3, 0x39, 0x9c, 0x8c, 0xc1, 0x24, 0x89, 0x96, 0x7b, 0x33, 0x71, 0x09,
	0xc1, 0x64, 0xd1, 0xa2, 0x31, 0x42, 0x22, 0x3f, 0x9a, 0xc8, 0xfc, 0x0d, 0x2e, 0x5e, 0xfe, 0x4c,
	0x4a, 0xa4, 0xa1, 0xae, 0x90, 0x3d, 0x81, 0xbf, 0xc4, 0x63, 0xd7, 0xfb, 0x08, 0xff, 0x2a, 0xc5,
	0xf1, 0x21, 0x69, 0x69, 0xf3, 0xe1, 0xcf, 0x2f, 0x79, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x92,
	0xda, 0xcf, 0xf0, 0x31, 0x01, 0x00, 0x00,
}
