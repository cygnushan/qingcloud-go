// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

/*
Package spec is a generated protocol buffer package.

It is generated from these files:
	base.proto
	instances.proto
	nic.proto
	snapshot.proto
	zone.proto
	cache.proto
	job.proto
	rdb.proto
	tag.proto
	dns.proto
	key_pair.proto
	router.proto
	user_data.proto
	eip.proto
	load_balancer.proto
	security_group.proto
	volume.proto
	image.proto
	mongo.proto
	shared_storage.proto
	vxnet.proto

It has these top-level messages:
	Options
	CommonRequest
	Error
	DescribeInstanceTypes_Request
	DescribeInstanceTypes_Reponse
	InstanceTypeSetElem
*/
package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_E_NULL ErrorCode = 0
	// Client
	ErrorCode_E_1100 ErrorCode = 1100
	ErrorCode_E_1200 ErrorCode = 1200
	ErrorCode_E_1300 ErrorCode = 1300
	ErrorCode_E_1400 ErrorCode = 1400
	ErrorCode_E_2100 ErrorCode = 2100
	ErrorCode_E_2400 ErrorCode = 2400
	ErrorCode_E_2500 ErrorCode = 2500
	// Server
	ErrorCode_E_5000 ErrorCode = 5000
	ErrorCode_E_5100 ErrorCode = 5100
	ErrorCode_E_5200 ErrorCode = 5200
	ErrorCode_E_5300 ErrorCode = 5300
)

var ErrorCode_name = map[int32]string{
	0:    "E_NULL",
	1100: "E_1100",
	1200: "E_1200",
	1300: "E_1300",
	1400: "E_1400",
	2100: "E_2100",
	2400: "E_2400",
	2500: "E_2500",
	5000: "E_5000",
	5100: "E_5100",
	5200: "E_5200",
	5300: "E_5300",
}
var ErrorCode_value = map[string]int32{
	"E_NULL": 0,
	"E_1100": 1100,
	"E_1200": 1200,
	"E_1300": 1300,
	"E_1400": 1400,
	"E_2100": 2100,
	"E_2400": 2400,
	"E_2500": 2500,
	"E_5000": 5000,
	"E_5100": 5100,
	"E_5200": 5200,
	"E_5300": 5300,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Options struct {
	Foo int32 `protobuf:"varint,1,opt,name=foo" json:"foo,omitempty"`
}

func (m *Options) Reset()                    { *m = Options{} }
func (m *Options) String() string            { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()               {}
func (*Options) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Options) GetFoo() int32 {
	if m != nil {
		return m.Foo
	}
	return 0
}

type CommonRequest struct {
	Action           string                      `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Zone             string                      `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	TimeStamp        *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=time_stamp,json=timeStamp" json:"time_stamp,omitempty"`
	AccessKeyId      string                      `protobuf:"bytes,4,opt,name=access_key_id,json=accessKeyId" json:"access_key_id,omitempty"`
	Version          int32                       `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	SignatureMethod  string                      `protobuf:"bytes,6,opt,name=signature_method,json=signatureMethod" json:"signature_method,omitempty"`
	SignatureVersion int32                       `protobuf:"varint,7,opt,name=signature_version,json=signatureVersion" json:"signature_version,omitempty"`
	Signature        string                      `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
}

func (m *CommonRequest) Reset()                    { *m = CommonRequest{} }
func (m *CommonRequest) String() string            { return proto.CompactTextString(m) }
func (*CommonRequest) ProtoMessage()               {}
func (*CommonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommonRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CommonRequest) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *CommonRequest) GetTimeStamp() *google_protobuf2.Timestamp {
	if m != nil {
		return m.TimeStamp
	}
	return nil
}

func (m *CommonRequest) GetAccessKeyId() string {
	if m != nil {
		return m.AccessKeyId
	}
	return ""
}

func (m *CommonRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CommonRequest) GetSignatureMethod() string {
	if m != nil {
		return m.SignatureMethod
	}
	return ""
}

func (m *CommonRequest) GetSignatureVersion() int32 {
	if m != nil {
		return m.SignatureVersion
	}
	return 0
}

func (m *CommonRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type Error struct {
	RetCode int32  `protobuf:"varint,1,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Error) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Options)(nil), "spec.Options")
	proto.RegisterType((*CommonRequest)(nil), "spec.CommonRequest")
	proto.RegisterType((*Error)(nil), "spec.Error")
	proto.RegisterEnum("spec.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0x9b, 0xd8, 0xf1, 0x8b, 0x2a, 0x86, 0x41, 0x42, 0x6e, 0x8a, 0xa0, 0xf2, 0x86,
	0x14, 0x24, 0x77, 0x9a, 0xd2, 0x05, 0x12, 0xbb, 0xaa, 0x0b, 0x44, 0x01, 0xc9, 0xfc, 0xd9, 0x5a,
	0x13, 0xfb, 0xd5, 0x58, 0x60, 0x8f, 0x99, 0x99, 0x80, 0xc2, 0x09, 0x38, 0x00, 0x67, 0x60, 0xc1,
	0x8a, 0x45, 0x8f, 0xc0, 0x92, 0x05, 0x47, 0xe0, 0x00, 0x1c, 0x80, 0x25, 0x9a, 0xf1, 0x4c, 0x23,
	0x65, 0xf7, 0x7e, 0xdf, 0xfb, 0xde, 0xe7, 0x4f, 0x63, 0x80, 0x05, 0x57, 0x98, 0x76, 0x52, 0x68,
	0x41, 0x87, 0xaa, 0xc3, 0x62, 0xba, 0x5b, 0x09, 0x51, 0xbd, 0xc7, 0x43, 0xab, 0x2d, 0x96, 0x17,
	0x87, 0xbc, 0x5d, 0xf5, 0x86, 0xe9, 0xde, 0xe6, 0x0a, 0x9b, 0x4e, 0xfb, 0xe5, 0xdd, 0xcd, 0xa5,
	0xae, 0x1b, 0x54, 0x9a, 0x37, 0x9d, 0x33, 0xdc, 0xd9, 0x34, 0x7c, 0x92, 0xbc, 0xeb, 0x50, 0xaa,
	0x7e, 0x9f, 0xec, 0x41, 0xf8, 0xa2, 0xd3, 0xb5, 0x68, 0x15, 0x25, 0xb0, 0x7d, 0x21, 0x44, 0x3c,
	0xd8, 0x1f, 0xcc, 0x46, 0x99, 0x19, 0x93, 0xef, 0x5b, 0xb0, 0x73, 0x2a, 0x9a, 0x46, 0xb4, 0x19,
	0x7e, 0x58, 0xa2, 0xd2, 0xf4, 0x16, 0x04, 0xbc, 0x30, 0x76, 0x6b, 0x8b, 0x32, 0x47, 0x94, 0xc2,
	0xf0, 0xb3, 0x68, 0x31, 0xde, 0xb2, 0xaa, 0x9d, 0xe9, 0x23, 0x00, 0xd3, 0x26, 0xb7, 0x75, 0xe2,
	0xed, 0xfd, 0xc1, 0x6c, 0x32, 0x9f, 0xa6, 0x7d, 0x9f, 0xd4, 0xf7, 0x49, 0x5f, 0xf9, 0xc2, 0x59,
	0x64, 0xdc, 0x2f, 0xcd, 0x48, 0x13, 0xd8, 0xe1, 0x45, 0x81, 0x4a, 0xe5, 0xef, 0x70, 0x95, 0xd7,
	0x65, 0x3c, 0xb4, 0xb9, 0x93, 0x5e, 0x7c, 0x8a, 0xab, 0x27, 0x25, 0x8d, 0x21, 0xfc, 0x88, 0x52,
	0x99, 0x2e, 0x23, 0x5b, 0xd9, 0x23, 0x3d, 0x00, 0xa2, 0xea, 0xaa, 0xe5, 0x7a, 0x29, 0x31, 0x6f,
	0x50, 0xbf, 0x15, 0x65, 0x1c, 0xd8, 0x80, 0xeb, 0x57, 0xfa, 0x33, 0x2b, 0xd3, 0x07, 0x70, 0x63,
	0x6d, 0xf5, 0x71, 0xa1, 0x8d, 0x5b, 0x67, 0xbc, 0x71, 0xb9, 0xb7, 0x21, 0xba, 0xd2, 0xe2, 0xb1,
	0x0d, 0x5c, 0x0b, 0xc9, 0x63, 0x18, 0x9d, 0x49, 0x29, 0x24, 0xdd, 0x85, 0xb1, 0x44, 0x9d, 0x17,
	0xa2, 0x44, 0xf7, 0x98, 0xa1, 0x44, 0x7d, 0x2a, 0x4a, 0x34, 0x9d, 0x1b, 0x54, 0x8a, 0x57, 0xfe,
	0xa5, 0x3c, 0xde, 0xff, 0x36, 0x80, 0xc8, 0x9e, 0x5b, 0x1f, 0x40, 0x70, 0x96, 0x3f, 0x7f, 0x7d,
	0x7e, 0x4e, 0xae, 0xd1, 0x89, 0x99, 0x8f, 0x8e, 0x18, 0x23, 0xbf, 0xc6, 0x0e, 0xe6, 0x8c, 0x91,
	0x1f, 0x91, 0x83, 0x63, 0xc6, 0xc8, 0x57, 0x70, 0xf0, 0x90, 0x31, 0xf2, 0xcf, 0xc1, 0xdc, 0xdc,
	0x5c, 0x12, 0x07, 0x66, 0xf3, 0x87, 0x3a, 0x38, 0x61, 0x8c, 0xfc, 0xbc, 0xd9, 0xc3, 0x09, 0x63,
	0x8c, 0x7c, 0xb9, 0xe7, 0xc0, 0xdc, 0xfc, 0xf5, 0x60, 0xbe, 0xf3, 0x7b, 0xe6, 0xc0, 0x7c, 0xe7,
	0xf2, 0x60, 0x11, 0xd8, 0x3f, 0x77, 0xfc, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x04, 0xa6, 0x20, 0xef,
	0xc4, 0x02, 0x00, 0x00,
}
