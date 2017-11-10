// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qingcloud.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/service.pb/qingcloud_sdk_rule"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type QingCloudServiceProperties struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *QingCloudServiceProperties) Reset()                    { *m = QingCloudServiceProperties{} }
func (m *QingCloudServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*QingCloudServiceProperties) ProtoMessage()               {}
func (*QingCloudServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func init() {
	proto.RegisterType((*QingCloudServiceProperties)(nil), "service.QingCloudServiceProperties")
}

func init() { proto.RegisterFile("qingcloud.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xcc, 0xcc, 0x4b,
	0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x95, 0x92, 0x85, 0xcb, 0xc4, 0x17, 0xa7, 0x64, 0xc7, 0x17, 0x95, 0xe6,
	0xa4, 0xea, 0x83, 0x08, 0x88, 0x3a, 0x29, 0xae, 0xaa, 0xfc, 0x3c, 0x28, 0x5b, 0x49, 0x86, 0x4b,
	0x2a, 0x30, 0x33, 0x2f, 0xdd, 0x19, 0xa4, 0x38, 0x18, 0xa2, 0x3d, 0xa0, 0x28, 0xbf, 0x20, 0xb5,
	0xa8, 0x24, 0x33, 0xb5, 0xd8, 0x68, 0x16, 0x23, 0x97, 0x00, 0xba, 0xb4, 0x90, 0x17, 0x17, 0xaf,
	0x4b, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x52, 0x6a, 0x54, 0x7e, 0x5e, 0x6a, 0xb1, 0x90, 0xb4, 0x1e,
	0xd4, 0x62, 0x3d, 0x14, 0x71, 0xcf, 0xbc, 0x82, 0xd2, 0x12, 0x29, 0x19, 0xec, 0x92, 0xfe, 0xa5,
	0x25, 0x20, 0x59, 0xab, 0xae, 0x8f, 0x2c, 0x66, 0x5c, 0x9a, 0x19, 0x25, 0x25, 0x05, 0xc5, 0x56,
	0xfa, 0xfa, 0x29, 0xf9, 0xc9, 0xc5, 0x7a, 0x08, 0x6f, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x16, 0x64,
	0xea, 0x67, 0xe6, 0xa5, 0xa4, 0x56, 0xe8, 0x65, 0x94, 0xe4, 0xe6, 0x08, 0x71, 0xc2, 0x9d, 0x03,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x79, 0xb0, 0x37, 0xd2, 0x00, 0x01, 0x00, 0x00,
}
