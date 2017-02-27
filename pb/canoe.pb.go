// Code generated by protoc-gen-go.
// source: canoe.proto
// DO NOT EDIT!

package breadpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateDeployRequest struct {
	UserEmail   string `protobuf:"bytes,1,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	Project     string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	TargetName  string `protobuf:"bytes,3,opt,name=target_name,json=targetName" json:"target_name,omitempty"`
	ArtifactUrl string `protobuf:"bytes,4,opt,name=artifact_url,json=artifactUrl" json:"artifact_url,omitempty"`
	Lock        bool   `protobuf:"varint,5,opt,name=lock" json:"lock,omitempty"`
}

func (m *CreateDeployRequest) Reset()                    { *m = CreateDeployRequest{} }
func (m *CreateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDeployRequest) ProtoMessage()               {}
func (*CreateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateDeployRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *CreateDeployRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *CreateDeployRequest) GetTargetName() string {
	if m != nil {
		return m.TargetName
	}
	return ""
}

func (m *CreateDeployRequest) GetArtifactUrl() string {
	if m != nil {
		return m.ArtifactUrl
	}
	return ""
}

func (m *CreateDeployRequest) GetLock() bool {
	if m != nil {
		return m.Lock
	}
	return false
}

type CreateDeployResponse struct {
	Error    bool   `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	DeployId int64  `protobuf:"varint,3,opt,name=deploy_id,json=deployId" json:"deploy_id,omitempty"`
}

func (m *CreateDeployResponse) Reset()                    { *m = CreateDeployResponse{} }
func (m *CreateDeployResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDeployResponse) ProtoMessage()               {}
func (*CreateDeployResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateDeployResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *CreateDeployResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateDeployResponse) GetDeployId() int64 {
	if m != nil {
		return m.DeployId
	}
	return 0
}

type CreateTerraformDeployRequest struct {
	UserEmail        string `protobuf:"bytes,1,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	Branch           string `protobuf:"bytes,2,opt,name=branch" json:"branch,omitempty"`
	Commit           string `protobuf:"bytes,3,opt,name=commit" json:"commit,omitempty"`
	Project          string `protobuf:"bytes,4,opt,name=project" json:"project,omitempty"`
	TerraformVersion string `protobuf:"bytes,5,opt,name=terraform_version,json=terraformVersion" json:"terraform_version,omitempty"`
}

func (m *CreateTerraformDeployRequest) Reset()                    { *m = CreateTerraformDeployRequest{} }
func (m *CreateTerraformDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTerraformDeployRequest) ProtoMessage()               {}
func (*CreateTerraformDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CreateTerraformDeployRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *CreateTerraformDeployRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *CreateTerraformDeployRequest) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *CreateTerraformDeployRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *CreateTerraformDeployRequest) GetTerraformVersion() string {
	if m != nil {
		return m.TerraformVersion
	}
	return ""
}

type CompleteTerraformDeployRequest struct {
	UserEmail  string `protobuf:"bytes,1,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	DeployId   int64  `protobuf:"varint,2,opt,name=deploy_id,json=deployId" json:"deploy_id,omitempty"`
	Successful bool   `protobuf:"varint,3,opt,name=successful" json:"successful,omitempty"`
	RequestId  string `protobuf:"bytes,4,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Project    string `protobuf:"bytes,5,opt,name=project" json:"project,omitempty"`
}

func (m *CompleteTerraformDeployRequest) Reset()                    { *m = CompleteTerraformDeployRequest{} }
func (m *CompleteTerraformDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CompleteTerraformDeployRequest) ProtoMessage()               {}
func (*CompleteTerraformDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CompleteTerraformDeployRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *CompleteTerraformDeployRequest) GetDeployId() int64 {
	if m != nil {
		return m.DeployId
	}
	return 0
}

func (m *CompleteTerraformDeployRequest) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *CompleteTerraformDeployRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *CompleteTerraformDeployRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type UnlockTerraformProjectRequest struct {
	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	Project   string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
}

func (m *UnlockTerraformProjectRequest) Reset()                    { *m = UnlockTerraformProjectRequest{} }
func (m *UnlockTerraformProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*UnlockTerraformProjectRequest) ProtoMessage()               {}
func (*UnlockTerraformProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UnlockTerraformProjectRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *UnlockTerraformProjectRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type TerraformDeployResponse struct {
	Error     bool   `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	DeployId  int64  `protobuf:"varint,3,opt,name=deploy_id,json=deployId" json:"deploy_id,omitempty"`
	RequestId string `protobuf:"bytes,4,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Project   string `protobuf:"bytes,5,opt,name=project" json:"project,omitempty"`
}

func (m *TerraformDeployResponse) Reset()                    { *m = TerraformDeployResponse{} }
func (m *TerraformDeployResponse) String() string            { return proto.CompactTextString(m) }
func (*TerraformDeployResponse) ProtoMessage()               {}
func (*TerraformDeployResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *TerraformDeployResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *TerraformDeployResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TerraformDeployResponse) GetDeployId() int64 {
	if m != nil {
		return m.DeployId
	}
	return 0
}

func (m *TerraformDeployResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *TerraformDeployResponse) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type PhoneAuthenticationRequest struct {
	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	Action    string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *PhoneAuthenticationRequest) Reset()                    { *m = PhoneAuthenticationRequest{} }
func (m *PhoneAuthenticationRequest) String() string            { return proto.CompactTextString(m) }
func (*PhoneAuthenticationRequest) ProtoMessage()               {}
func (*PhoneAuthenticationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *PhoneAuthenticationRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *PhoneAuthenticationRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type PhoneAuthenticationResponse struct {
	Error     bool   `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	UserEmail string `protobuf:"bytes,3,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
}

func (m *PhoneAuthenticationResponse) Reset()                    { *m = PhoneAuthenticationResponse{} }
func (m *PhoneAuthenticationResponse) String() string            { return proto.CompactTextString(m) }
func (*PhoneAuthenticationResponse) ProtoMessage()               {}
func (*PhoneAuthenticationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *PhoneAuthenticationResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *PhoneAuthenticationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PhoneAuthenticationResponse) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateDeployRequest)(nil), "bread.CreateDeployRequest")
	proto.RegisterType((*CreateDeployResponse)(nil), "bread.CreateDeployResponse")
	proto.RegisterType((*CreateTerraformDeployRequest)(nil), "bread.CreateTerraformDeployRequest")
	proto.RegisterType((*CompleteTerraformDeployRequest)(nil), "bread.CompleteTerraformDeployRequest")
	proto.RegisterType((*UnlockTerraformProjectRequest)(nil), "bread.UnlockTerraformProjectRequest")
	proto.RegisterType((*TerraformDeployResponse)(nil), "bread.TerraformDeployResponse")
	proto.RegisterType((*PhoneAuthenticationRequest)(nil), "bread.PhoneAuthenticationRequest")
	proto.RegisterType((*PhoneAuthenticationResponse)(nil), "bread.PhoneAuthenticationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Canoe service

type CanoeClient interface {
	CreateDeploy(ctx context.Context, in *CreateDeployRequest, opts ...grpc.CallOption) (*CreateDeployResponse, error)
	CreateTerraformDeploy(ctx context.Context, in *CreateTerraformDeployRequest, opts ...grpc.CallOption) (*TerraformDeployResponse, error)
	CompleteTerraformDeploy(ctx context.Context, in *CompleteTerraformDeployRequest, opts ...grpc.CallOption) (*TerraformDeployResponse, error)
	UnlockTerraformProject(ctx context.Context, in *UnlockTerraformProjectRequest, opts ...grpc.CallOption) (*TerraformDeployResponse, error)
	PhoneAuthentication(ctx context.Context, in *PhoneAuthenticationRequest, opts ...grpc.CallOption) (*PhoneAuthenticationResponse, error)
}

type canoeClient struct {
	cc *grpc.ClientConn
}

func NewCanoeClient(cc *grpc.ClientConn) CanoeClient {
	return &canoeClient{cc}
}

func (c *canoeClient) CreateDeploy(ctx context.Context, in *CreateDeployRequest, opts ...grpc.CallOption) (*CreateDeployResponse, error) {
	out := new(CreateDeployResponse)
	err := grpc.Invoke(ctx, "/bread.Canoe/CreateDeploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *canoeClient) CreateTerraformDeploy(ctx context.Context, in *CreateTerraformDeployRequest, opts ...grpc.CallOption) (*TerraformDeployResponse, error) {
	out := new(TerraformDeployResponse)
	err := grpc.Invoke(ctx, "/bread.Canoe/CreateTerraformDeploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *canoeClient) CompleteTerraformDeploy(ctx context.Context, in *CompleteTerraformDeployRequest, opts ...grpc.CallOption) (*TerraformDeployResponse, error) {
	out := new(TerraformDeployResponse)
	err := grpc.Invoke(ctx, "/bread.Canoe/CompleteTerraformDeploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *canoeClient) UnlockTerraformProject(ctx context.Context, in *UnlockTerraformProjectRequest, opts ...grpc.CallOption) (*TerraformDeployResponse, error) {
	out := new(TerraformDeployResponse)
	err := grpc.Invoke(ctx, "/bread.Canoe/UnlockTerraformProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *canoeClient) PhoneAuthentication(ctx context.Context, in *PhoneAuthenticationRequest, opts ...grpc.CallOption) (*PhoneAuthenticationResponse, error) {
	out := new(PhoneAuthenticationResponse)
	err := grpc.Invoke(ctx, "/bread.Canoe/PhoneAuthentication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Canoe service

type CanoeServer interface {
	CreateDeploy(context.Context, *CreateDeployRequest) (*CreateDeployResponse, error)
	CreateTerraformDeploy(context.Context, *CreateTerraformDeployRequest) (*TerraformDeployResponse, error)
	CompleteTerraformDeploy(context.Context, *CompleteTerraformDeployRequest) (*TerraformDeployResponse, error)
	UnlockTerraformProject(context.Context, *UnlockTerraformProjectRequest) (*TerraformDeployResponse, error)
	PhoneAuthentication(context.Context, *PhoneAuthenticationRequest) (*PhoneAuthenticationResponse, error)
}

func RegisterCanoeServer(s *grpc.Server, srv CanoeServer) {
	s.RegisterService(&_Canoe_serviceDesc, srv)
}

func _Canoe_CreateDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanoeServer).CreateDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Canoe/CreateDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanoeServer).CreateDeploy(ctx, req.(*CreateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Canoe_CreateTerraformDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTerraformDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanoeServer).CreateTerraformDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Canoe/CreateTerraformDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanoeServer).CreateTerraformDeploy(ctx, req.(*CreateTerraformDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Canoe_CompleteTerraformDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTerraformDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanoeServer).CompleteTerraformDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Canoe/CompleteTerraformDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanoeServer).CompleteTerraformDeploy(ctx, req.(*CompleteTerraformDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Canoe_UnlockTerraformProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockTerraformProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanoeServer).UnlockTerraformProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Canoe/UnlockTerraformProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanoeServer).UnlockTerraformProject(ctx, req.(*UnlockTerraformProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Canoe_PhoneAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanoeServer).PhoneAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bread.Canoe/PhoneAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanoeServer).PhoneAuthentication(ctx, req.(*PhoneAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Canoe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bread.Canoe",
	HandlerType: (*CanoeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeploy",
			Handler:    _Canoe_CreateDeploy_Handler,
		},
		{
			MethodName: "CreateTerraformDeploy",
			Handler:    _Canoe_CreateTerraformDeploy_Handler,
		},
		{
			MethodName: "CompleteTerraformDeploy",
			Handler:    _Canoe_CompleteTerraformDeploy_Handler,
		},
		{
			MethodName: "UnlockTerraformProject",
			Handler:    _Canoe_UnlockTerraformProject_Handler,
		},
		{
			MethodName: "PhoneAuthentication",
			Handler:    _Canoe_PhoneAuthentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "canoe.proto",
}

func init() { proto.RegisterFile("canoe.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcf, 0x4e, 0x14, 0x4f,
	0x10, 0xc7, 0xd3, 0xc0, 0xc2, 0x6e, 0xc1, 0xe1, 0xf7, 0x6b, 0x10, 0x36, 0x0b, 0xac, 0xd0, 0xf8,
	0x07, 0x51, 0x97, 0x44, 0x6f, 0xdc, 0x14, 0x3d, 0x70, 0x31, 0x64, 0x14, 0x63, 0xbc, 0x4c, 0x7a,
	0x7b, 0x8b, 0x65, 0x64, 0x66, 0x7a, 0xec, 0xe9, 0x31, 0xf1, 0xea, 0xc5, 0xc4, 0xc4, 0xc4, 0xc4,
	0x07, 0xf0, 0xe6, 0xd1, 0x93, 0x3e, 0x89, 0xaf, 0xe0, 0x83, 0x98, 0xe9, 0xee, 0x61, 0x67, 0x64,
	0x00, 0x5d, 0xe3, 0x6d, 0xeb, 0xcf, 0x56, 0x7d, 0xaa, 0xb6, 0xbe, 0xbd, 0x30, 0x2b, 0x78, 0x2c,
	0xb1, 0x97, 0x28, 0xa9, 0x25, 0x6d, 0xf4, 0x15, 0xf2, 0x41, 0x67, 0x65, 0x28, 0xe5, 0x30, 0xc4,
	0x6d, 0x9e, 0x04, 0xdb, 0x3c, 0x8e, 0xa5, 0xe6, 0x3a, 0x90, 0x71, 0x6a, 0x93, 0xd8, 0x67, 0x02,
	0xf3, 0xbb, 0x0a, 0xb9, 0xc6, 0x07, 0x98, 0x84, 0xf2, 0xb5, 0x87, 0x2f, 0x33, 0x4c, 0x35, 0x5d,
	0x05, 0xc8, 0x52, 0x54, 0x3e, 0x46, 0x3c, 0x08, 0xdb, 0x64, 0x8d, 0x6c, 0xb6, 0xbc, 0x56, 0xee,
	0x79, 0x98, 0x3b, 0x68, 0x1b, 0x66, 0x12, 0x25, 0x5f, 0xa0, 0xd0, 0xed, 0x09, 0x13, 0x2b, 0x4c,
	0x7a, 0x19, 0x66, 0x35, 0x57, 0x43, 0xd4, 0x7e, 0xcc, 0x23, 0x6c, 0x4f, 0x9a, 0x28, 0x58, 0xd7,
	0x23, 0x1e, 0x21, 0x5d, 0x87, 0x39, 0xae, 0x74, 0x70, 0xc8, 0x85, 0xf6, 0x33, 0x15, 0xb6, 0xa7,
	0x4c, 0xc6, 0x6c, 0xe1, 0x3b, 0x50, 0x21, 0xa5, 0x30, 0x15, 0x4a, 0x71, 0xdc, 0x6e, 0xac, 0x91,
	0xcd, 0xa6, 0x67, 0x3e, 0x33, 0x01, 0x0b, 0x55, 0xce, 0x34, 0x91, 0x71, 0x8a, 0x74, 0x01, 0x1a,
	0xa8, 0x94, 0x54, 0x86, 0xb1, 0xe9, 0x59, 0x23, 0xe7, 0x8b, 0x30, 0x4d, 0xf9, 0x10, 0x0b, 0x3e,
	0x67, 0xd2, 0x65, 0x68, 0x0d, 0x4c, 0x05, 0x3f, 0x18, 0x18, 0xba, 0x49, 0xaf, 0x69, 0x1d, 0x7b,
	0x03, 0xf6, 0x95, 0xc0, 0x8a, 0xed, 0xf2, 0x04, 0x95, 0xe2, 0x87, 0x52, 0x45, 0x7f, 0xb4, 0x96,
	0x45, 0x98, 0xee, 0x2b, 0x1e, 0x8b, 0x23, 0xd7, 0xd5, 0x59, 0xb9, 0x5f, 0xc8, 0x28, 0x0a, 0xb4,
	0xdb, 0x87, 0xb3, 0xca, 0x6b, 0x9c, 0xaa, 0xae, 0xf1, 0x26, 0xfc, 0xaf, 0x0b, 0x04, 0xff, 0x15,
	0xaa, 0x34, 0x90, 0xb1, 0xd9, 0x47, 0xcb, 0xfb, 0xef, 0x24, 0xf0, 0xd4, 0xfa, 0xd9, 0x37, 0x02,
	0xdd, 0x5d, 0x19, 0x25, 0x21, 0x8e, 0x0b, 0x5e, 0xd9, 0xca, 0x44, 0x75, 0x2b, 0xb4, 0x0b, 0x90,
	0x66, 0x42, 0x60, 0x9a, 0x1e, 0x66, 0xa1, 0x99, 0xa0, 0xe9, 0x95, 0x3c, 0x79, 0x6d, 0x65, 0xdb,
	0xe4, 0xdf, 0xb6, 0x83, 0xb4, 0x9c, 0x67, 0x6f, 0x50, 0x1e, 0xb2, 0x51, 0x19, 0x92, 0x3d, 0x83,
	0xd5, 0x83, 0x38, 0xff, 0x75, 0x4f, 0xa0, 0xf7, 0x6d, 0xe4, 0x6f, 0xaf, 0x90, 0x7d, 0x22, 0xb0,
	0x74, 0x6a, 0x13, 0xff, 0xe0, 0x62, 0xc6, 0x9f, 0xfd, 0x31, 0x74, 0xf6, 0x8f, 0x64, 0x8c, 0xf7,
	0x32, 0x7d, 0x84, 0xb1, 0x0e, 0x84, 0x91, 0xe5, 0xef, 0xdf, 0x19, 0x17, 0x79, 0x7e, 0x71, 0x67,
	0xd6, 0x62, 0x21, 0x2c, 0xd7, 0x16, 0x1d, 0x73, 0xf2, 0x2a, 0xc5, 0xe4, 0x2f, 0x14, 0x77, 0xbe,
	0x34, 0xa0, 0xb1, 0x9b, 0x3f, 0x38, 0xf4, 0x18, 0xe6, 0xca, 0xe2, 0xa4, 0x9d, 0x9e, 0x79, 0x7b,
	0x7a, 0x35, 0x2f, 0x4b, 0x67, 0xb9, 0x36, 0x66, 0x09, 0x19, 0x7b, 0xf3, 0xfd, 0xc7, 0xc7, 0x89,
	0x15, 0xb6, 0x64, 0x9e, 0xab, 0xa1, 0x4a, 0xc4, 0xb6, 0x30, 0x79, 0xbe, 0xdd, 0xf8, 0x0e, 0xd9,
	0xa2, 0xef, 0x08, 0x5c, 0xaa, 0x15, 0x29, 0xdd, 0xa8, 0x94, 0xae, 0x57, 0x42, 0xa7, 0xeb, 0x92,
	0xce, 0x38, 0x0f, 0x76, 0xcb, 0x20, 0x5c, 0x63, 0xeb, 0xa7, 0x10, 0x46, 0x82, 0x1c, 0xc1, 0x7c,
	0x20, 0xb0, 0x74, 0x86, 0xf4, 0xe8, 0xd5, 0x02, 0xe7, 0x5c, 0x69, 0x5e, 0x08, 0xd4, 0x33, 0x40,
	0x9b, 0x6c, 0xa3, 0x04, 0xe4, 0x2a, 0xd6, 0x22, 0xbd, 0x27, 0xb0, 0x58, 0x2f, 0x2b, 0x7a, 0xc5,
	0xb5, 0x3a, 0x57, 0x75, 0x17, 0x02, 0xdd, 0x36, 0x40, 0xd7, 0x19, 0x1b, 0x01, 0x65, 0xa6, 0x60,
	0x09, 0xc7, 0x9d, 0x79, 0xce, 0xf3, 0x96, 0xc0, 0x7c, 0xcd, 0x55, 0xd2, 0x75, 0xd7, 0xe6, 0x6c,
	0x19, 0x74, 0xd8, 0x79, 0x29, 0x8e, 0xe6, 0x86, 0xa1, 0xd9, 0x60, 0xdd, 0x11, 0x4d, 0x92, 0xa7,
	0xfb, 0xbc, 0x92, 0xbf, 0x43, 0xb6, 0xee, 0xb7, 0x9e, 0xcf, 0x98, 0x7a, 0x49, 0xbf, 0x3f, 0x6d,
	0xfe, 0xfe, 0xee, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x00, 0x4d, 0x2d, 0xa0, 0x32, 0x07, 0x00,
	0x00,
}
