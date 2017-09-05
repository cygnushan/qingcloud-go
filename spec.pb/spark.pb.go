// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spark.proto

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

type SparkService interface {
	CreateSpark(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeSparks(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddSparkNodes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteSparkNodes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartSparks(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopSparks(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteSparks(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type SparkServiceClient struct{}

// NewSparkServiceClient returns a SparkService stub to handle
// requests to the set of SparkService at the other end of the connection.
func NewSparkServiceClient(opt *Options) *SparkServiceClient {
	return &SparkServiceClient{}
}

func (c *SparkServiceClient) CreateSpark(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SparkServiceClient) DescribeSparks(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SparkServiceClient) AddSparkNodes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SparkServiceClient) DeleteSparkNodes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SparkServiceClient) StartSparks(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SparkServiceClient) StopSparks(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *SparkServiceClient) DeleteSparks(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("spark.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8f, 0xcb, 0x4e, 0x85, 0x30,
	0x10, 0x86, 0x17, 0x1a, 0x17, 0x03, 0x1a, 0xd3, 0x85, 0x89, 0x98, 0xe8, 0x1b, 0x94, 0x44, 0xb7,
	0xe2, 0x15, 0xb7, 0x6e, 0xfa, 0x04, 0x05, 0x46, 0x42, 0x04, 0x3b, 0x99, 0x8e, 0x1a, 0x1f, 0xd3,
	0x37, 0x3a, 0xa1, 0x85, 0xe4, 0x84, 0xe4, 0x2c, 0x60, 0xd9, 0xff, 0xf2, 0x75, 0x7e, 0x48, 0x3c,
	0x59, 0xfe, 0xd4, 0xc4, 0x4e, 0x9c, 0x3a, 0xf6, 0x84, 0x75, 0x76, 0xd9, 0x3a, 0xd7, 0xf6, 0x98,
	0x07, 0xad, 0xfa, 0xfe, 0xc8, 0xed, 0xd7, 0x5f, 0x0c, 0x64, 0x57, 0x4b, 0x0b, 0x07, 0x92, 0xd9,
	0xbc, 0x59, 0x9a, 0xd2, 0x0d, 0xe8, 0xc5, 0x0e, 0x34, 0x05, 0xae, 0x97, 0x81, 0x5f, 0xb6, 0x44,
	0xc8, 0x3e, 0xfa, 0xb7, 0xff, 0x47, 0x90, 0x9a, 0xf1, 0x1c, 0x83, 0xfc, 0xd3, 0xd5, 0xa8, 0x0a,
	0x48, 0x5e, 0x19, 0xad, 0x60, 0x50, 0xd5, 0x85, 0x8e, 0x00, 0x3d, 0x03, 0xf4, 0xdb, 0xf8, 0x7d,
	0x76, 0x40, 0x57, 0x4f, 0x70, 0x56, 0xa2, 0xaf, 0xb9, 0xab, 0x22, 0xc0, 0xaf, 0x26, 0x3c, 0xc2,
	0xe9, 0x73, 0xd3, 0x84, 0xf2, 0xbb, 0x6b, 0x70, 0x3d, 0xe0, 0x05, 0xce, 0x4b, 0xec, 0x71, 0x5a,
	0xb0, 0x8d, 0x51, 0x40, 0x62, 0xc4, 0xb2, 0x6c, 0xdc, 0x70, 0x0f, 0x60, 0xc4, 0xd1, 0xc6, 0xf6,
	0x03, 0xa4, 0x7b, 0x03, 0x56, 0xf7, 0xab, 0x93, 0xf0, 0xbe, 0xdb, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xfb, 0x04, 0x62, 0x68, 0x02, 0x00, 0x00,
}
