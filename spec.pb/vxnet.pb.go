// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vxnet.proto

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

type VxnetService interface {
	DescribeVxnets(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateVxnets(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteVxnets(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	JoinVxnet(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	LeaveVxnet(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyVxnetAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeVxnetInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type VxnetServiceClient struct{}

// NewVxnetServiceClient returns a VxnetService stub to handle
// requests to the set of VxnetService at the other end of the connection.
func NewVxnetServiceClient(opt *Options) *VxnetServiceClient {
	return &VxnetServiceClient{}
}

func (c *VxnetServiceClient) DescribeVxnets(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *VxnetServiceClient) CreateVxnets(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *VxnetServiceClient) DeleteVxnets(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *VxnetServiceClient) JoinVxnet(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *VxnetServiceClient) LeaveVxnet(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *VxnetServiceClient) ModifyVxnetAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *VxnetServiceClient) DescribeVxnetInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("vxnet.proto", fileDescriptor29) }

var fileDescriptor29 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xab, 0xc8, 0x4b,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x96, 0x92, 0x4c,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42,
	0x14, 0x48, 0x49, 0xa3, 0x4b, 0xa5, 0xe6, 0x16, 0x94, 0xc0, 0x24, 0xe5, 0xd1, 0x25, 0x4b, 0x32,
	0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0xa0, 0x0a, 0xe4, 0xd0, 0x15, 0x94, 0x17, 0x25, 0x16,
	0x14, 0xa4, 0x16, 0x15, 0x43, 0xe4, 0x8d, 0xee, 0x31, 0x73, 0xf1, 0x84, 0x81, 0x9c, 0x13, 0x9c,
	0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc0, 0xc5, 0xe7, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x99,
	0x94, 0x0a, 0x16, 0x2f, 0x16, 0x12, 0xd3, 0x83, 0x98, 0xa1, 0x07, 0x33, 0x43, 0xcf, 0x15, 0xe4,
	0x02, 0x29, 0x1c, 0xe2, 0x42, 0x76, 0x5c, 0x3c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0x14, 0xe8, 0x77,
	0x49, 0xcd, 0x49, 0x25, 0x5b, 0xbf, 0x35, 0x17, 0xa7, 0x57, 0x7e, 0x66, 0x1e, 0x58, 0x37, 0xc9,
	0x9a, 0x6d, 0xb8, 0xb8, 0x7c, 0x52, 0x13, 0xcb, 0x52, 0xc9, 0xd3, 0xed, 0xce, 0x25, 0xea, 0x9b,
	0x9f, 0x92, 0x99, 0x56, 0x09, 0xd6, 0xee, 0x58, 0x52, 0x52, 0x94, 0x99, 0x54, 0x5a, 0x92, 0x4a,
	0xba, 0x1f, 0x3c, 0xb8, 0xc4, 0x50, 0x62, 0xc1, 0x33, 0xaf, 0xb8, 0x24, 0x31, 0x2f, 0x99, 0x74,
	0x93, 0x92, 0xd8, 0xc0, 0x7c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x81, 0x0f, 0x24,
	0x75, 0x02, 0x00, 0x00,
}
