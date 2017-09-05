// Code generated by protoc-gen-go. DO NOT EDIT.
// source: security_group.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "context"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SecurityGroupsService interface {
	DescribeSecurityGroups(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateSecurityGroup(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteSecurityGroups(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ApplySecurityGroup(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifySecurityGroupAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeSecurityGroupRules(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddSecurityGroupRules(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteSecurityGroupRules(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifySecurityGroupRuleAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateSecurityGroupSnapshot(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeSecurityGroupSnapshots(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteSecurityGroupSnapshots(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RollbackSecurityGroup(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeSecurityGroupIPSets(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateSecurityGroupIPSet(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteSecurityGroupIPSets(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifySecurityGroupIPSetAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CopySecurityGroupIPSets(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type SecurityGroupsServiceClient struct{}

// NewSecurityGroupsServiceClient returns a SecurityGroupsService stub to handle
// requests to the set of SecurityGroupsService at the other end of the connection.
func NewSecurityGroupsServiceClient(opt *Options) *SecurityGroupsServiceClient {
	return &SecurityGroupsServiceClient{}
}

func (c *SecurityGroupsServiceClient) DescribeSecurityGroups(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) CreateSecurityGroup(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DeleteSecurityGroups(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) ApplySecurityGroup(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) ModifySecurityGroupAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DescribeSecurityGroupRules(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) AddSecurityGroupRules(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DeleteSecurityGroupRules(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) ModifySecurityGroupRuleAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) CreateSecurityGroupSnapshot(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DescribeSecurityGroupSnapshots(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DeleteSecurityGroupSnapshots(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) RollbackSecurityGroup(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DescribeSecurityGroupIPSets(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) CreateSecurityGroupIPSet(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) DeleteSecurityGroupIPSets(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) ModifySecurityGroupIPSetAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SecurityGroupsServiceClient) CopySecurityGroupIPSets(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("security_group.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x86, 0x2f, 0x5f, 0xbe, 0xc3, 0x78, 0xab, 0x80, 0x52, 0x14, 0xa3, 0x3f, 0xa0, 0x24, 0xfa,
	0x0b, 0x08, 0x20, 0xa2, 0xa2, 0x84, 0x7a, 0x37, 0xdb, 0x32, 0xe0, 0xc6, 0x2d, 0xbb, 0xd9, 0x9d,
	0x6a, 0x7a, 0xf5, 0x97, 0x9b, 0xb6, 0xf4, 0xd0, 0x66, 0x3d, 0xb4, 0xe5, 0xd8, 0x9d, 0xc9, 0x93,
	0xa7, 0xef, 0xbc, 0xd0, 0x31, 0x18, 0xc6, 0x9a, 0x53, 0xf2, 0xbe, 0xd3, 0x32, 0x56, 0x9e, 0xd2,
	0x92, 0xa4, 0xf3, 0xcf, 0x28, 0x0c, 0xdd, 0xfe, 0x4e, 0xca, 0x9d, 0xc0, 0x51, 0xf6, 0x16, 0xc4,
	0xdb, 0x11, 0xdb, 0x27, 0xf9, 0x82, 0x3b, 0xa8, 0x8e, 0x30, 0x52, 0x54, 0x0c, 0xaf, 0xaa, 0x43,
	0xe2, 0x11, 0x1a, 0x62, 0xd1, 0x01, 0xef, 0x0e, 0xab, 0x0b, 0xdf, 0x9a, 0x29, 0x85, 0xda, 0xe4,
	0xf3, 0xdb, 0x9f, 0x13, 0xe8, 0xfa, 0x07, 0xaf, 0x79, 0xaa, 0x65, 0x7c, 0xd4, 0x5f, 0x3c, 0x44,
	0xe7, 0x01, 0x7a, 0x53, 0x34, 0xa1, 0xe6, 0x01, 0x96, 0x17, 0x9c, 0x9e, 0x97, 0x43, 0xbd, 0x02,
	0xea, 0xcd, 0x52, 0x25, 0xf7, 0x8f, 0x77, 0x67, 0x06, 0xa7, 0x13, 0x8d, 0x8c, 0xca, 0x9c, 0xda,
	0x98, 0x7b, 0xe8, 0x4c, 0x51, 0x20, 0xb5, 0xd5, 0x99, 0x82, 0x33, 0x56, 0x4a, 0x24, 0xed, 0x6c,
	0x5e, 0xe1, 0x72, 0x29, 0x37, 0x7c, 0x5b, 0xc6, 0x8c, 0x89, 0x34, 0x0f, 0x62, 0xc2, 0xfa, 0x5a,
	0xcf, 0xe0, 0x5a, 0xf3, 0x5e, 0xc7, 0xa2, 0x01, 0x6d, 0x0e, 0xdd, 0xf1, 0x66, 0x73, 0x04, 0xd0,
	0x23, 0x9c, 0x5b, 0x52, 0x6f, 0xc6, 0xf2, 0xe1, 0xda, 0x92, 0x59, 0xca, 0x6a, 0x91, 0xdb, 0x12,
	0x06, 0x96, 0x76, 0xf9, 0x7b, 0xa6, 0xcc, 0x87, 0xa4, 0xda, 0xb8, 0x15, 0x0c, 0xad, 0x67, 0x28,
	0x80, 0xf5, 0x05, 0x5f, 0xe0, 0xc2, 0x92, 0x60, 0x73, 0xde, 0x1c, 0xba, 0x6b, 0x29, 0x44, 0xc0,
	0xc2, 0xcf, 0x76, 0x15, 0x5e, 0xc2, 0xc0, 0xfa, 0xab, 0x8b, 0x95, 0x8f, 0xd4, 0xa8, 0x29, 0x96,
	0x43, 0x64, 0xb0, 0xda, 0xac, 0x27, 0xe8, 0x5b, 0x32, 0x6b, 0x28, 0xf6, 0x06, 0x37, 0x96, 0xda,
	0x65, 0xb0, 0x16, 0xbd, 0x5b, 0xc0, 0xd9, 0x44, 0xaa, 0xe4, 0x08, 0x82, 0xc1, 0xff, 0xec, 0xfb,
	0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x60, 0xea, 0x04, 0x37, 0x22, 0x06, 0x00, 0x00,
}
