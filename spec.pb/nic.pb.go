// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nic.proto

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

type NicService interface {
	CreateNics(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeNics(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AttachNics(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DetachNics(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyNicAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteNics(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type NicServiceClient struct{}

// NewNicServiceClient returns a NicService stub to handle
// requests to the set of NicService at the other end of the connection.
func NewNicServiceClient(opt *Options) *NicServiceClient {
	return &NicServiceClient{}
}

func (c *NicServiceClient) CreateNics(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *NicServiceClient) DescribeNics(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *NicServiceClient) AttachNics(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *NicServiceClient) DetachNics(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *NicServiceClient) ModifyNicAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *NicServiceClient) DeleteNics(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("nic.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x0f, 0x8a, 0xe0, 0xe0, 0x29, 0x82, 0x60, 0x05, 0x7d, 0x83, 0x14, 0xf4, 0x5a, 0x84,
	0x62, 0x7b, 0xb4, 0x17, 0x9f, 0x20, 0x19, 0xa7, 0x75, 0xa0, 0x6d, 0x42, 0x32, 0x75, 0xe9, 0x33,
	0xec, 0x4b, 0x2f, 0xdb, 0x6e, 0x2f, 0x85, 0xbd, 0xe4, 0x38, 0xff, 0xf7, 0x7f, 0xcc, 0x30, 0x70,
	0x3f, 0x32, 0x6a, 0x1f, 0x9c, 0x38, 0x75, 0x1b, 0x3d, 0x61, 0xf6, 0xdc, 0x39, 0xd7, 0xf5, 0x94,
	0x2f, 0x99, 0x9d, 0xda, 0xdc, 0x8c, 0xf3, 0x5a, 0xc8, 0x5e, 0xf6, 0x88, 0x06, 0x2f, 0x1b, 0x7c,
	0xdb, 0x43, 0xe1, 0x81, 0xa2, 0x98, 0xc1, 0x5f, 0x0a, 0xaf, 0xfb, 0xc2, 0x21, 0x18, 0xef, 0x29,
	0xc4, 0x95, 0xbf, 0x1f, 0x6f, 0x00, 0x1a, 0xc6, 0x1f, 0x0a, 0xff, 0x8c, 0xa4, 0x0a, 0x80, 0xaf,
	0x40, 0x46, 0xa8, 0x61, 0x8c, 0xea, 0x49, 0xaf, 0xb6, 0xde, 0x6c, 0x5d, 0x9f, 0x77, 0x67, 0x57,
	0x72, 0xf5, 0x09, 0x0f, 0x15, 0x45, 0x0c, 0x6c, 0xd3, 0xfc, 0x02, 0xa0, 0x14, 0x31, 0xf8, 0x97,
	0x6a, 0x57, 0x94, 0x6c, 0xd7, 0xf0, 0xf8, 0xed, 0x7e, 0xb9, 0x9d, 0x1b, 0xc6, 0x52, 0x24, 0xb0,
	0x9d, 0x84, 0x12, 0x8f, 0xe8, 0x29, 0xed, 0x81, 0xf6, 0x6e, 0x99, 0x3f, 0x4e, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x9e, 0xc6, 0xda, 0x20, 0x02, 0x00, 0x00,
}
