// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qingcloud_sdk_rule/rule.proto

/*
Package qingcloud_sdk_rule is a generated protocol buffer package.

It is generated from these files:
	qingcloud_sdk_rule/rule.proto

It has these top-level messages:
	ServiceRule
	MethodRule
	MethodInputRule
*/
package qingcloud_sdk_rule

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 服务规则
// 有主服务和子服务之分, 子服务隶属于某个主服务
type ServiceRule struct {
	DocUrl         string `protobuf:"bytes,1,opt,name=doc_url,json=docUrl" json:"doc_url,omitempty"`
	ServiceName    string `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	SubServiceName string `protobuf:"bytes,3,opt,name=sub_service_name,json=subServiceName" json:"sub_service_name,omitempty"`
}

func (m *ServiceRule) Reset()                    { *m = ServiceRule{} }
func (m *ServiceRule) String() string            { return proto.CompactTextString(m) }
func (*ServiceRule) ProtoMessage()               {}
func (*ServiceRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceRule) GetDocUrl() string {
	if m != nil {
		return m.DocUrl
	}
	return ""
}

func (m *ServiceRule) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceRule) GetSubServiceName() string {
	if m != nil {
		return m.SubServiceName
	}
	return ""
}

// 方法规则
type MethodRule struct {
	DocUrl     string `protobuf:"bytes,1,opt,name=doc_url,json=docUrl" json:"doc_url,omitempty"`
	HttpMethod string `protobuf:"bytes,2,opt,name=http_method,json=httpMethod" json:"http_method,omitempty"`
}

func (m *MethodRule) Reset()                    { *m = MethodRule{} }
func (m *MethodRule) String() string            { return proto.CompactTextString(m) }
func (*MethodRule) ProtoMessage()               {}
func (*MethodRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MethodRule) GetDocUrl() string {
	if m != nil {
		return m.DocUrl
	}
	return ""
}

func (m *MethodRule) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

// 输入参数规则
// 输入参数成员只有数值类型和字符串, 以及对应的数组, 不含嵌套结构
type MethodInputRule struct {
	RequiredFileds string `protobuf:"bytes,1,opt,name=required_fileds,json=requiredFileds" json:"required_fileds,omitempty"`
	DefaultValue   string `protobuf:"bytes,2,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	EnumValue      string `protobuf:"bytes,3,opt,name=enum_value,json=enumValue" json:"enum_value,omitempty"`
	Minimum        string `protobuf:"bytes,4,opt,name=minimum" json:"minimum,omitempty"`
	Maximum        string `protobuf:"bytes,5,opt,name=maximum" json:"maximum,omitempty"`
	MultipleOf     string `protobuf:"bytes,6,opt,name=multipleOf" json:"multipleOf,omitempty"`
}

func (m *MethodInputRule) Reset()                    { *m = MethodInputRule{} }
func (m *MethodInputRule) String() string            { return proto.CompactTextString(m) }
func (*MethodInputRule) ProtoMessage()               {}
func (*MethodInputRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MethodInputRule) GetRequiredFileds() string {
	if m != nil {
		return m.RequiredFileds
	}
	return ""
}

func (m *MethodInputRule) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func (m *MethodInputRule) GetEnumValue() string {
	if m != nil {
		return m.EnumValue
	}
	return ""
}

func (m *MethodInputRule) GetMinimum() string {
	if m != nil {
		return m.Minimum
	}
	return ""
}

func (m *MethodInputRule) GetMaximum() string {
	if m != nil {
		return m.Maximum
	}
	return ""
}

func (m *MethodInputRule) GetMultipleOf() string {
	if m != nil {
		return m.MultipleOf
	}
	return ""
}

var E_ServiceRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ServiceRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.service_rule",
	Tag:           "bytes,10001,opt,name=service_rule,json=serviceRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

var E_MethodRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.method_rule",
	Tag:           "bytes,10001,opt,name=method_rule,json=methodRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

var E_MethodInputRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MethodInputRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.method_input_rule",
	Tag:           "bytes,10001,opt,name=method_input_rule,json=methodInputRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

func init() {
	proto.RegisterType((*ServiceRule)(nil), "qingcloud.sdk.rule.ServiceRule")
	proto.RegisterType((*MethodRule)(nil), "qingcloud.sdk.rule.MethodRule")
	proto.RegisterType((*MethodInputRule)(nil), "qingcloud.sdk.rule.MethodInputRule")
	proto.RegisterExtension(E_ServiceRule)
	proto.RegisterExtension(E_MethodRule)
	proto.RegisterExtension(E_MethodInputRule)
}

func init() { proto.RegisterFile("qingcloud_sdk_rule/rule.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xf8, 0x48, 0xd5, 0x71, 0x69, 0x60, 0x2f, 0x58, 0x48, 0x4d, 0x42, 0x7a, 0x20, 0x17,
	0xd6, 0xa5, 0xdc, 0xc2, 0x8d, 0x43, 0x25, 0x0e, 0x4d, 0x51, 0x2a, 0x38, 0x70, 0x31, 0xb6, 0x77,
	0xe2, 0xac, 0xba, 0xeb, 0x75, 0xf6, 0xa3, 0xe2, 0x6f, 0xf0, 0xd7, 0xf8, 0x45, 0xc8, 0xbb, 0x8e,
	0x43, 0x6b, 0xda, 0x8b, 0xb5, 0xf3, 0xde, 0xdb, 0x79, 0x9e, 0x79, 0x36, 0x9c, 0x6c, 0x79, 0x55,
	0x16, 0x42, 0x39, 0x96, 0x1a, 0x76, 0x93, 0x6a, 0x27, 0x30, 0x69, 0x1e, 0xb4, 0xd6, 0xca, 0x2a,
	0x42, 0x3a, 0x9a, 0x1a, 0x76, 0x43, 0x1b, 0xe6, 0xcd, 0xb4, 0x54, 0xaa, 0x14, 0x98, 0x78, 0x45,
	0xee, 0xd6, 0x09, 0x43, 0x53, 0x68, 0x5e, 0x5b, 0xa5, 0xc3, 0xad, 0x99, 0x81, 0xe8, 0x1a, 0xf5,
	0x2d, 0x2f, 0x70, 0xe5, 0x04, 0x92, 0xd7, 0x70, 0xc0, 0x54, 0x91, 0x3a, 0x2d, 0xe2, 0xc1, 0x74,
	0x30, 0x3f, 0x5c, 0x0d, 0x99, 0x2a, 0xbe, 0x69, 0x41, 0xde, 0xc2, 0x91, 0x09, 0xba, 0xb4, 0xca,
	0x24, 0xc6, 0x4f, 0x3c, 0x1b, 0xb5, 0xd8, 0x32, 0x93, 0x48, 0xe6, 0xf0, 0xd2, 0xb8, 0x3c, 0xbd,
	0x23, 0x7b, 0xea, 0x65, 0xc7, 0xc6, 0xe5, 0xd7, 0x7b, 0xe5, 0xec, 0x02, 0xe0, 0x12, 0xed, 0x46,
	0xb1, 0xc7, 0x3d, 0x27, 0x10, 0x6d, 0xac, 0xad, 0x53, 0xe9, 0xb5, 0xad, 0x25, 0x34, 0x50, 0xb8,
	0x3d, 0xfb, 0x33, 0x80, 0x51, 0x38, 0x7e, 0xa9, 0x6a, 0x67, 0x7d, 0xb7, 0x77, 0x30, 0xd2, 0xb8,
	0x75, 0x5c, 0x23, 0x4b, 0xd7, 0x5c, 0x20, 0x33, 0x6d, 0xd7, 0xe3, 0x1d, 0x7c, 0xe1, 0x51, 0x72,
	0x0a, 0x2f, 0x18, 0xae, 0x33, 0x27, 0x6c, 0x7a, 0x9b, 0x09, 0xb7, 0x1b, 0xe9, 0xa8, 0x05, 0xbf,
	0x37, 0x18, 0x39, 0x01, 0xc0, 0xca, 0xc9, 0x56, 0x11, 0xa6, 0x39, 0x6c, 0x90, 0x40, 0xc7, 0x70,
	0x20, 0x79, 0xc5, 0xa5, 0x93, 0xf1, 0x33, 0xcf, 0xed, 0x4a, 0xcf, 0x64, 0xbf, 0x3c, 0xf3, 0xbc,
	0x65, 0x42, 0x49, 0xc6, 0x00, 0xd2, 0x09, 0xcb, 0x6b, 0x81, 0x57, 0xeb, 0x78, 0x18, 0x86, 0xda,
	0x23, 0x0b, 0xb6, 0xdf, 0x74, 0x93, 0x21, 0x99, 0xd0, 0x10, 0x22, 0xdd, 0x85, 0x48, 0xdb, 0x55,
	0x5e, 0xd5, 0x96, 0xab, 0xca, 0xc4, 0xbf, 0x97, 0xd3, 0xc1, 0x3c, 0x3a, 0x9f, 0xd0, 0xfe, 0x07,
	0x40, 0xff, 0xc9, 0xb6, 0x0b, 0xab, 0x29, 0x16, 0x3f, 0x21, 0x0a, 0x6b, 0x0d, 0x26, 0xe3, 0x9e,
	0x49, 0xd8, 0xeb, 0x3d, 0x8f, 0xf1, 0xff, 0x3c, 0xf6, 0x51, 0xae, 0x40, 0x76, 0xe7, 0xc5, 0x16,
	0x5e, 0xb5, 0x0e, 0xbc, 0x09, 0xe7, 0xa1, 0x61, 0x2e, 0xd1, 0x98, 0xac, 0xbc, 0x3f, 0xcc, 0xe9,
	0xc3, 0x46, 0x5d, 0xd4, 0xab, 0x91, 0xbc, 0x0b, 0x7c, 0xfe, 0xfa, 0x63, 0x59, 0x72, 0xbb, 0x71,
	0x39, 0x2d, 0x94, 0x4c, 0x8a, 0x4d, 0xc6, 0xcf, 0xcf, 0x3e, 0x9c, 0x25, 0x5d, 0xab, 0xf7, 0xa5,
	0x4a, 0x4c, 0x8d, 0x05, 0xad, 0xf3, 0xa4, 0xff, 0x33, 0x7d, 0xea, 0x43, 0xf9, 0xd0, 0xbf, 0xe7,
	0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x04, 0x8f, 0x1f, 0x85, 0x7c, 0x03, 0x00, 0x00,
}
