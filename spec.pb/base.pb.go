// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

/*
Package spec is a generated protocol buffer package.

It is generated from these files:
	base.proto
	instances.proto

It has these top-level messages:
	Message2
	Message
*/
package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message2 struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Message2) Reset()                    { *m = Message2{} }
func (m *Message2) String() string            { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()               {}
func (*Message2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message2) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message2)(nil), "spec.Message2")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x56, 0x92, 0xe1, 0xe2,
	0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x35, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x8d, 0x4c, 0xb9, 0xb8, 0x5d, 0x93, 0x33, 0xf2,
	0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xd4, 0xb8, 0x58, 0x40, 0x5c, 0x21, 0x3e, 0x3d,
	0x90, 0x5e, 0x3d, 0x98, 0x46, 0x29, 0x34, 0x7e, 0x12, 0x1b, 0xd8, 0x06, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x9c, 0x6b, 0x75, 0x6f, 0x00, 0x00, 0x00,
}
