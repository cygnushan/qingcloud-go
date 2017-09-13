// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qingcloud_sdk_rule/rule.proto

/*
Package qingcloud_sdk_rule is a generated protocol buffer package.

It is generated from these files:
	qingcloud_sdk_rule/rule.proto

It has these top-level messages:
	ServiceOptionsRule
	MethodOptionsRule
	MessageOptionsRule
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
type ServiceOptionsRule struct {
	DocUrl         string `protobuf:"bytes,1,opt,name=doc_url,json=docUrl" json:"doc_url,omitempty"`
	ServiceName    string `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	SubServiceName string `protobuf:"bytes,3,opt,name=sub_service_name,json=subServiceName" json:"sub_service_name,omitempty"`
}

func (m *ServiceOptionsRule) Reset()                    { *m = ServiceOptionsRule{} }
func (m *ServiceOptionsRule) String() string            { return proto.CompactTextString(m) }
func (*ServiceOptionsRule) ProtoMessage()               {}
func (*ServiceOptionsRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceOptionsRule) GetDocUrl() string {
	if m != nil {
		return m.DocUrl
	}
	return ""
}

func (m *ServiceOptionsRule) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceOptionsRule) GetSubServiceName() string {
	if m != nil {
		return m.SubServiceName
	}
	return ""
}

// 方法规则
type MethodOptionsRule struct {
	DocUrl     string `protobuf:"bytes,1,opt,name=doc_url,json=docUrl" json:"doc_url,omitempty"`
	HttpMethod string `protobuf:"bytes,2,opt,name=http_method,json=httpMethod" json:"http_method,omitempty"`
	InputType  string `protobuf:"bytes,3,opt,name=input_type,json=inputType" json:"input_type,omitempty"`
	OutputType string `protobuf:"bytes,4,opt,name=output_type,json=outputType" json:"output_type,omitempty"`
}

func (m *MethodOptionsRule) Reset()                    { *m = MethodOptionsRule{} }
func (m *MethodOptionsRule) String() string            { return proto.CompactTextString(m) }
func (*MethodOptionsRule) ProtoMessage()               {}
func (*MethodOptionsRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MethodOptionsRule) GetDocUrl() string {
	if m != nil {
		return m.DocUrl
	}
	return ""
}

func (m *MethodOptionsRule) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *MethodOptionsRule) GetInputType() string {
	if m != nil {
		return m.InputType
	}
	return ""
}

func (m *MethodOptionsRule) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

// 输入参数规则
// 输入参数成员只有数值类型和字符串, 以及对应的数组, 不含嵌套结构
// 元信息部分只能包含 字符串/数字/链接符号 等普通的字符
type MessageOptionsRule struct {
	RequiredFileds  string `protobuf:"bytes,1,opt,name=required_fileds,json=requiredFileds" json:"required_fileds,omitempty"`
	DefaultValue    string `protobuf:"bytes,2,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	EnumValue       string `protobuf:"bytes,3,opt,name=enum_value,json=enumValue" json:"enum_value,omitempty"`
	MinValue        string `protobuf:"bytes,4,opt,name=min_value,json=minValue" json:"min_value,omitempty"`
	MaxValue        string `protobuf:"bytes,5,opt,name=max_value,json=maxValue" json:"max_value,omitempty"`
	MultipleOfValue string `protobuf:"bytes,6,opt,name=multiple_of_value,json=multipleOfValue" json:"multiple_of_value,omitempty"`
	RegexpValue     string `protobuf:"bytes,7,opt,name=regexp_value,json=regexpValue" json:"regexp_value,omitempty"`
}

func (m *MessageOptionsRule) Reset()                    { *m = MessageOptionsRule{} }
func (m *MessageOptionsRule) String() string            { return proto.CompactTextString(m) }
func (*MessageOptionsRule) ProtoMessage()               {}
func (*MessageOptionsRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MessageOptionsRule) GetRequiredFileds() string {
	if m != nil {
		return m.RequiredFileds
	}
	return ""
}

func (m *MessageOptionsRule) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func (m *MessageOptionsRule) GetEnumValue() string {
	if m != nil {
		return m.EnumValue
	}
	return ""
}

func (m *MessageOptionsRule) GetMinValue() string {
	if m != nil {
		return m.MinValue
	}
	return ""
}

func (m *MessageOptionsRule) GetMaxValue() string {
	if m != nil {
		return m.MaxValue
	}
	return ""
}

func (m *MessageOptionsRule) GetMultipleOfValue() string {
	if m != nil {
		return m.MultipleOfValue
	}
	return ""
}

func (m *MessageOptionsRule) GetRegexpValue() string {
	if m != nil {
		return m.RegexpValue
	}
	return ""
}

var E_ServiceRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ServiceOptionsRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.service_rule",
	Tag:           "bytes,10001,opt,name=service_rule,json=serviceRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

var E_MethodRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodOptionsRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.method_rule",
	Tag:           "bytes,10001,opt,name=method_rule,json=methodRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

var E_MessageRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MessageOptionsRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.message_rule",
	Tag:           "bytes,10001,opt,name=message_rule,json=messageRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

func init() {
	proto.RegisterType((*ServiceOptionsRule)(nil), "qingcloud.sdk.rule.ServiceOptionsRule")
	proto.RegisterType((*MethodOptionsRule)(nil), "qingcloud.sdk.rule.MethodOptionsRule")
	proto.RegisterType((*MessageOptionsRule)(nil), "qingcloud.sdk.rule.MessageOptionsRule")
	proto.RegisterExtension(E_ServiceRule)
	proto.RegisterExtension(E_MethodRule)
	proto.RegisterExtension(E_MessageRule)
}

func init() { proto.RegisterFile("qingcloud_sdk_rule/rule.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x55, 0xfa, 0x7d, 0xa4, 0x64, 0x12, 0x5a, 0xea, 0x0d, 0x11, 0xa8, 0x34, 0x0d, 0x02, 0x22,
	0x24, 0xc6, 0xa5, 0xec, 0xca, 0x8e, 0x05, 0xbb, 0xa6, 0x28, 0x05, 0x16, 0x6c, 0x46, 0xb6, 0xe7,
	0xc6, 0x19, 0x65, 0xec, 0x99, 0xce, 0x4f, 0x95, 0xf2, 0x08, 0x48, 0x2c, 0x78, 0x63, 0x34, 0x3f,
	0x4e, 0x64, 0x19, 0x10, 0x1b, 0x4b, 0x73, 0xce, 0xb9, 0xf7, 0xdc, 0x99, 0x7b, 0x8c, 0x8e, 0x6f,
	0x58, 0x5d, 0x16, 0x5c, 0x58, 0x4a, 0x34, 0x5d, 0x13, 0x65, 0x39, 0xa4, 0xee, 0x83, 0xa5, 0x12,
	0x46, 0x24, 0xc9, 0x96, 0xc6, 0x9a, 0xae, 0xb1, 0x63, 0x1e, 0x4f, 0x4a, 0x21, 0x4a, 0x0e, 0xa9,
	0x57, 0xe4, 0x76, 0x99, 0x52, 0xd0, 0x85, 0x62, 0xd2, 0x08, 0x15, 0xaa, 0xa6, 0xdf, 0x50, 0x72,
	0x0d, 0xea, 0x96, 0x15, 0x70, 0x25, 0x0d, 0x13, 0xb5, 0x5e, 0x58, 0x0e, 0xc9, 0x23, 0xb4, 0x4f,
	0x45, 0x41, 0xac, 0xe2, 0xe3, 0xde, 0xa4, 0x37, 0x1b, 0x2c, 0xfa, 0x54, 0x14, 0x9f, 0x15, 0x4f,
	0x4e, 0xd1, 0x48, 0x07, 0x39, 0xa9, 0xb3, 0x0a, 0xc6, 0x7b, 0x9e, 0x1d, 0x46, 0x6c, 0x9e, 0x55,
	0x90, 0xcc, 0xd0, 0x43, 0x6d, 0x73, 0xd2, 0x92, 0xfd, 0xe7, 0x65, 0x07, 0xda, 0xe6, 0xd7, 0x3b,
	0xe5, 0xf4, 0x47, 0x0f, 0x1d, 0x5d, 0x82, 0x59, 0x09, 0xfa, 0x4f, 0xde, 0x27, 0x68, 0xb8, 0x32,
	0x46, 0x92, 0xca, 0x97, 0x44, 0x6b, 0xe4, 0xa0, 0xd0, 0x24, 0x39, 0x46, 0x88, 0xd5, 0xd2, 0x1a,
	0x62, 0xee, 0x64, 0xe3, 0x39, 0xf0, 0xc8, 0xa7, 0x3b, 0x09, 0xae, 0x5e, 0x58, 0xb3, 0xe5, 0xff,
	0x0f, 0xf5, 0x01, 0x72, 0x82, 0xe9, 0xf7, 0x3d, 0x94, 0x5c, 0x82, 0xd6, 0x59, 0xd9, 0x7a, 0x8c,
	0x97, 0xe8, 0x50, 0xc1, 0x8d, 0x65, 0x0a, 0x28, 0x59, 0x32, 0x0e, 0x54, 0xc7, 0xc1, 0x0e, 0x1a,
	0xf8, 0x83, 0x47, 0x93, 0x67, 0xe8, 0x01, 0x85, 0x65, 0x66, 0xb9, 0x21, 0xb7, 0x19, 0xb7, 0xcd,
	0xeb, 0x8c, 0x22, 0xf8, 0xc5, 0x61, 0x6e, 0x48, 0xa8, 0x6d, 0x15, 0x15, 0x71, 0x48, 0x87, 0x04,
	0xfa, 0x09, 0x1a, 0x54, 0xac, 0x8e, 0x6c, 0x18, 0xf1, 0x7e, 0xc5, 0xea, 0x1d, 0x99, 0x6d, 0x22,
	0x79, 0x2f, 0x92, 0xd9, 0x26, 0x90, 0xaf, 0xd0, 0x51, 0x65, 0xb9, 0x61, 0x92, 0x03, 0x11, 0xcb,
	0x28, 0xea, 0x7b, 0xd1, 0x61, 0x43, 0x5c, 0x2d, 0x83, 0xf6, 0x14, 0x8d, 0x14, 0x94, 0xb0, 0x91,
	0x51, 0xb6, 0x1f, 0xd6, 0x18, 0x30, 0x2f, 0xb9, 0x58, 0xef, 0x36, 0xed, 0xa2, 0x94, 0x9c, 0xe0,
	0x90, 0x25, 0xdc, 0x64, 0x09, 0xb7, 0x73, 0x33, 0xfe, 0x39, 0x9f, 0xf4, 0x66, 0xc3, 0xf3, 0x17,
	0xb8, 0x9b, 0x43, 0xdc, 0x8d, 0xd8, 0x36, 0x33, 0xee, 0x70, 0x51, 0xa2, 0x61, 0xd8, 0x6a, 0xf0,
	0x7a, 0xda, 0xf1, 0x6a, 0xc5, 0xa4, 0xb1, 0x7a, 0xfe, 0x3b, 0xab, 0x4e, 0xa0, 0x16, 0x28, 0xb4,
	0xf6, 0x46, 0x6b, 0x34, 0xaa, 0xc2, 0x86, 0xff, 0x74, 0xab, 0x76, 0x00, 0xfe, 0x7a, 0xab, 0x6e,
	0x56, 0x16, 0xc3, 0xd8, 0xdd, 0x1d, 0xde, 0x7f, 0xfc, 0x3a, 0x2f, 0x99, 0x59, 0xd9, 0x1c, 0x17,
	0xa2, 0x4a, 0x8b, 0x55, 0xc6, 0xce, 0xcf, 0xde, 0x9c, 0xa5, 0xdb, 0x4e, 0xaf, 0x4b, 0x91, 0x6a,
	0x09, 0x05, 0x96, 0x79, 0xda, 0xfd, 0xb7, 0xdf, 0x75, 0xa1, 0xbc, 0xef, 0xc7, 0x7c, 0xfb, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x63, 0xe6, 0x56, 0x0b, 0x04, 0x00, 0x00,
}
