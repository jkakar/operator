// Code generated by protoc-gen-go.
// source: repository.proto
// DO NOT EDIT!

package breadpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RepositoryConfig struct {
	// List of automated test statuses that are required to pass for this
	// repository's build to be considered successful
	RequiredTestingStatuses []string `protobuf:"bytes,1,rep,name=required_testing_statuses,json=requiredTestingStatuses" json:"required_testing_statuses,omitempty"`
	// Changing files matching glob patterns configured in this watchlist
	// requires additional peer review and approval from a member of the
	// @Pardot/security team
	SecurityWatchlist *RepositoryWatchlist `protobuf:"bytes,2,opt,name=security_watchlist,json=securityWatchlist" json:"security_watchlist,omitempty"`
	// Watchlists attached to this repository
	Watchlists []*RepositoryWatchlist `protobuf:"bytes,3,rep,name=watchlists" json:"watchlists,omitempty"`
}

func (m *RepositoryConfig) Reset()                    { *m = RepositoryConfig{} }
func (m *RepositoryConfig) String() string            { return proto.CompactTextString(m) }
func (*RepositoryConfig) ProtoMessage()               {}
func (*RepositoryConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *RepositoryConfig) GetRequiredTestingStatuses() []string {
	if m != nil {
		return m.RequiredTestingStatuses
	}
	return nil
}

func (m *RepositoryConfig) GetSecurityWatchlist() *RepositoryWatchlist {
	if m != nil {
		return m.SecurityWatchlist
	}
	return nil
}

func (m *RepositoryConfig) GetWatchlists() []*RepositoryWatchlist {
	if m != nil {
		return m.Watchlists
	}
	return nil
}

type RepositoryWatchlist struct {
	// Slug of the team that should be notified when a file matching one of the
	// globs listed bellow matches. Example: Pardot/build-and-automate
	Team string `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	// List of file glob patterns
	Globs []string `protobuf:"bytes,2,rep,name=globs" json:"globs,omitempty"`
}

func (m *RepositoryWatchlist) Reset()                    { *m = RepositoryWatchlist{} }
func (m *RepositoryWatchlist) String() string            { return proto.CompactTextString(m) }
func (*RepositoryWatchlist) ProtoMessage()               {}
func (*RepositoryWatchlist) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *RepositoryWatchlist) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *RepositoryWatchlist) GetGlobs() []string {
	if m != nil {
		return m.Globs
	}
	return nil
}

func init() {
	proto.RegisterType((*RepositoryConfig)(nil), "bread.RepositoryConfig")
	proto.RegisterType((*RepositoryWatchlist)(nil), "bread.RepositoryWatchlist")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2d, 0xc8,
	0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2a,
	0x4a, 0x4d, 0x4c, 0x51, 0xba, 0xca, 0xc8, 0x25, 0x10, 0x04, 0x97, 0x73, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0x17, 0xb2, 0xe2, 0x92, 0x2c, 0x4a, 0x2d, 0x2c, 0xcd, 0x2c, 0x4a, 0x4d, 0x89, 0x2f, 0x49,
	0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x4e, 0x2d, 0x96, 0x60,
	0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x12, 0x87, 0x29, 0x08, 0x81, 0xc8, 0x07, 0x43, 0xa5, 0x85, 0x3c,
	0xb9, 0x84, 0x8a, 0x53, 0x93, 0x4b, 0x8b, 0x32, 0x4b, 0x2a, 0xe3, 0xcb, 0x13, 0x4b, 0x92, 0x33,
	0x72, 0x32, 0x8b, 0x4b, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0xc0, 0x96, 0xea,
	0x21, 0x2c, 0x0c, 0x87, 0xa9, 0x08, 0x12, 0x84, 0xe9, 0x82, 0x0b, 0x09, 0x59, 0x71, 0x71, 0xc1,
	0x4d, 0x28, 0x96, 0x60, 0x56, 0x60, 0x26, 0x60, 0x04, 0x92, 0x6a, 0x25, 0x7b, 0x2e, 0x61, 0x2c,
	0x4a, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x13, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xf4, 0x9c, 0xfc, 0xa4, 0x62, 0x09, 0x26, 0xb0, 0xcf, 0x20,
	0x9c, 0x24, 0x36, 0x70, 0x30, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x03, 0xfe, 0x1d,
	0x3a, 0x01, 0x00, 0x00,
}
