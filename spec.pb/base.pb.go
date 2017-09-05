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
}

func (m *Options) Reset()                    { *m = Options{} }
func (m *Options) String() string            { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()               {}
func (*Options) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4d, 0x6e, 0xd4, 0x30,
	0x1c, 0xc5, 0x49, 0x99, 0xaf, 0xfc, 0x47, 0x15, 0xc6, 0x48, 0x28, 0x1d, 0x10, 0x8c, 0xb2, 0x21,
	0x05, 0x29, 0x75, 0xa7, 0x74, 0x81, 0xc4, 0xae, 0xea, 0x02, 0x51, 0x40, 0x0a, 0x1f, 0xdb, 0xc8,
	0x93, 0xfc, 0x09, 0x11, 0x24, 0x0e, 0xb6, 0x07, 0x34, 0x9c, 0x80, 0x03, 0x70, 0x06, 0x16, 0xac,
	0x58, 0xf4, 0x08, 0x2c, 0x59, 0x70, 0x04, 0x0e, 0xc0, 0x01, 0x58, 0x56, 0x76, 0xec, 0x8e, 0x34,
	0x3b, 0xff, 0xde, 0x7b, 0x7e, 0xcf, 0x4a, 0x00, 0x96, 0x5c, 0x61, 0xda, 0x49, 0xa1, 0x05, 0x1d,
	0xa8, 0x0e, 0x8b, 0xd9, 0x5e, 0x25, 0x44, 0xf5, 0x01, 0x0f, 0xac, 0xb6, 0x5c, 0xbd, 0x3d, 0xe0,
	0xed, 0xba, 0x0f, 0xcc, 0x6e, 0x6d, 0x5b, 0xd8, 0x74, 0xda, 0x9b, 0x77, 0xb7, 0x4d, 0x5d, 0x37,
	0xa8, 0x34, 0x6f, 0x3a, 0x17, 0xb8, 0xb3, 0x1d, 0xf8, 0x2c, 0x79, 0xd7, 0xa1, 0x54, 0xbd, 0x1f,
	0x87, 0x30, 0x7e, 0xd1, 0xe9, 0x5a, 0xb4, 0x2a, 0xfe, 0xb1, 0x03, 0xbb, 0x27, 0xa2, 0x69, 0x44,
	0x9b, 0xe1, 0xc7, 0x15, 0x2a, 0x4d, 0x6f, 0xc2, 0x88, 0x17, 0xc6, 0x8c, 0x82, 0x79, 0x90, 0x84,
	0x99, 0x23, 0x4a, 0x61, 0xf0, 0x45, 0xb4, 0x18, 0xed, 0x58, 0xd5, 0x9e, 0xe9, 0x23, 0x00, 0xb3,
	0x9d, 0xdb, 0xf1, 0xe8, 0xea, 0x3c, 0x48, 0xa6, 0x8b, 0x59, 0xda, 0xaf, 0xa7, 0x7e, 0x3d, 0x7d,
	0xe5, 0x9f, 0x97, 0x85, 0x26, 0xfd, 0xd2, 0x1c, 0x69, 0x0c, 0xbb, 0xbc, 0x28, 0x50, 0xa9, 0xfc,
	0x3d, 0xae, 0xf3, 0xba, 0x8c, 0x06, 0xb6, 0x77, 0xda, 0x8b, 0x4f, 0x71, 0xfd, 0xa4, 0xa4, 0x11,
	0x8c, 0x3f, 0xa1, 0x54, 0xe6, 0x2d, 0xc3, 0x79, 0x90, 0x0c, 0x33, 0x8f, 0x74, 0x1f, 0x88, 0xaa,
	0xab, 0x96, 0xeb, 0x95, 0xc4, 0xbc, 0x41, 0xfd, 0x4e, 0x94, 0xd1, 0xc8, 0x16, 0x5c, 0xbb, 0xd4,
	0x9f, 0x59, 0x99, 0x3e, 0x80, 0xeb, 0x9b, 0xa8, 0xaf, 0x1b, 0xdb, 0xba, 0x4d, 0xc7, 0x1b, 0xd7,
	0x7b, 0x1b, 0xc2, 0x4b, 0x2d, 0x9a, 0xd8, 0xc2, 0x8d, 0x10, 0x3f, 0x86, 0xe1, 0xa9, 0x94, 0x42,
	0xd2, 0x3d, 0x98, 0x48, 0xd4, 0x79, 0x21, 0x4a, 0xb4, 0x5f, 0x69, 0x98, 0x8d, 0x25, 0xea, 0x13,
	0x51, 0xa2, 0x79, 0x73, 0x83, 0x4a, 0xf1, 0xca, 0x7f, 0x29, 0x8f, 0xf7, 0xbf, 0x07, 0x10, 0xda,
	0xeb, 0x36, 0x07, 0x30, 0x3a, 0xcd, 0x9f, 0xbf, 0x3e, 0x3b, 0x23, 0x57, 0xe8, 0xd4, 0x9c, 0x0f,
	0x0f, 0x19, 0x23, 0xbf, 0x27, 0x0e, 0x16, 0x8c, 0x91, 0x9f, 0xa1, 0x83, 0x23, 0xc6, 0xc8, 0x37,
	0x70, 0xf0, 0x90, 0x31, 0xf2, 0xdf, 0xc1, 0xc2, 0xdc, 0x39, 0x27, 0x0e, 0x8c, 0xf3, 0x97, 0x3a,
	0x38, 0x66, 0x8c, 0xfc, 0xba, 0xd1, 0xc3, 0x31, 0x63, 0x8c, 0x7c, 0xbd, 0xe7, 0xc0, 0xdc, 0xf9,
	0xe7, 0xc1, 0xec, 0xfc, 0x49, 0x1c, 0x98, 0x9d, 0xf3, 0xfd, 0xe5, 0xc8, 0xfe, 0xb9, 0xa3, 0x8b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x58, 0x45, 0x2d, 0xb2, 0x02, 0x00, 0x00,
}
