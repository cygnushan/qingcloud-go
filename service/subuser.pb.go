// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subuser.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = config.Config{}
var _ = request.Request{}
var _ = data.Operation{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SubuserServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SubuserServiceProperties) Reset()                    { *m = SubuserServiceProperties{} }
func (m *SubuserServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SubuserServiceProperties) ProtoMessage()               {}
func (*SubuserServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *SubuserServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeSubUsersInput struct {
}

func (m *DescribeSubUsersInput) Reset()                    { *m = DescribeSubUsersInput{} }
func (m *DescribeSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSubUsersInput) ProtoMessage()               {}
func (*DescribeSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

type DescribeSubUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DescribeSubUsersOutput) Reset()                    { *m = DescribeSubUsersOutput{} }
func (m *DescribeSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSubUsersOutput) ProtoMessage()               {}
func (*DescribeSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{2} }

func (m *DescribeSubUsersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeSubUsersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeSubUsersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateSubUserInput struct {
}

func (m *CreateSubUserInput) Reset()                    { *m = CreateSubUserInput{} }
func (m *CreateSubUserInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserInput) ProtoMessage()               {}
func (*CreateSubUserInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

type CreateSubUserOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CreateSubUserOutput) Reset()                    { *m = CreateSubUserOutput{} }
func (m *CreateSubUserOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserOutput) ProtoMessage()               {}
func (*CreateSubUserOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{4} }

func (m *CreateSubUserOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateSubUserOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateSubUserOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ModifySubUserAttributesInput struct {
}

func (m *ModifySubUserAttributesInput) Reset()                    { *m = ModifySubUserAttributesInput{} }
func (m *ModifySubUserAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesInput) ProtoMessage()               {}
func (*ModifySubUserAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

type ModifySubUserAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifySubUserAttributesOutput) Reset()                    { *m = ModifySubUserAttributesOutput{} }
func (m *ModifySubUserAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesOutput) ProtoMessage()               {}
func (*ModifySubUserAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{6} }

func (m *ModifySubUserAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifySubUserAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifySubUserAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteSubUsersInput struct {
}

func (m *DeleteSubUsersInput) Reset()                    { *m = DeleteSubUsersInput{} }
func (m *DeleteSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersInput) ProtoMessage()               {}
func (*DeleteSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

type DeleteSubUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteSubUsersOutput) Reset()                    { *m = DeleteSubUsersOutput{} }
func (m *DeleteSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersOutput) ProtoMessage()               {}
func (*DeleteSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{8} }

func (m *DeleteSubUsersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteSubUsersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteSubUsersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RestoreSubUsersInput struct {
}

func (m *RestoreSubUsersInput) Reset()                    { *m = RestoreSubUsersInput{} }
func (m *RestoreSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersInput) ProtoMessage()               {}
func (*RestoreSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

type RestoreSubUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *RestoreSubUsersOutput) Reset()                    { *m = RestoreSubUsersOutput{} }
func (m *RestoreSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersOutput) ProtoMessage()               {}
func (*RestoreSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{10} }

func (m *RestoreSubUsersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RestoreSubUsersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RestoreSubUsersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SubuserServiceProperties)(nil), "service.SubuserServiceProperties")
	proto.RegisterType((*DescribeSubUsersInput)(nil), "service.DescribeSubUsersInput")
	proto.RegisterType((*DescribeSubUsersOutput)(nil), "service.DescribeSubUsersOutput")
	proto.RegisterType((*CreateSubUserInput)(nil), "service.CreateSubUserInput")
	proto.RegisterType((*CreateSubUserOutput)(nil), "service.CreateSubUserOutput")
	proto.RegisterType((*ModifySubUserAttributesInput)(nil), "service.ModifySubUserAttributesInput")
	proto.RegisterType((*ModifySubUserAttributesOutput)(nil), "service.ModifySubUserAttributesOutput")
	proto.RegisterType((*DeleteSubUsersInput)(nil), "service.DeleteSubUsersInput")
	proto.RegisterType((*DeleteSubUsersOutput)(nil), "service.DeleteSubUsersOutput")
	proto.RegisterType((*RestoreSubUsersInput)(nil), "service.RestoreSubUsersInput")
	proto.RegisterType((*RestoreSubUsersOutput)(nil), "service.RestoreSubUsersOutput")
}

type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

type SubuserService struct {
	Config           *config.Config
	Properties       *SubuserServiceProperties
	LastResponseBody string
}

func NewSubuserService(conf *config.Config, zone string) (p *SubuserService) {
	return &SubuserService{
		Config:     conf,
		Properties: &SubuserServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Subuser(zone string) (*SubuserService, error) {
	properties := &SubuserServiceProperties{
		Zone: zone,
	}

	return &SubuserService{Config: s.Config, Properties: properties}, nil
}

func (p *SubuserService) DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error) {
	if in == nil {
		in = &DescribeSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSubUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeSubUsersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DescribeSubUsersInput) Validate() error {
	return nil
}

func (p *SubuserService) CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error) {
	if in == nil {
		in = &CreateSubUserInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSubUser",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateSubUserOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *CreateSubUserInput) Validate() error {
	return nil
}

func (p *SubuserService) ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error) {
	if in == nil {
		in = &ModifySubUserAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySubUserAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifySubUserAttributesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *ModifySubUserAttributesInput) Validate() error {
	return nil
}

func (p *SubuserService) DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error) {
	if in == nil {
		in = &DeleteSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSubUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteSubUsersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DeleteSubUsersInput) Validate() error {
	return nil
}

func (p *SubuserService) RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error) {
	if in == nil {
		in = &RestoreSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestoreSubUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &RestoreSubUsersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *RestoreSubUsersInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("subuser.proto", fileDescriptor25) }

var fileDescriptor25 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6e, 0xda, 0x40,
	0x14, 0xc6, 0x45, 0xa1, 0xb8, 0x3c, 0x09, 0x5a, 0x0d, 0x06, 0x5c, 0x17, 0x28, 0xb2, 0xd4, 0x8a,
	0x95, 0x2b, 0xb5, 0x27, 0xa8, 0x60, 0x93, 0x48, 0x28, 0x04, 0x2b, 0x6b, 0x84, 0xed, 0x17, 0x32,
	0x8a, 0xe3, 0x71, 0xe6, 0x4f, 0xa4, 0xe4, 0x08, 0x39, 0x60, 0xee, 0x91, 0x1b, 0x44, 0xe0, 0x81,
	0xc4, 0x06, 0x27, 0x2b, 0x6f, 0x10, 0xf3, 0xbe, 0xf9, 0xde, 0x37, 0xf6, 0xfb, 0x79, 0xa0, 0x29,
	0x94, 0xaf, 0x04, 0x72, 0x37, 0xe1, 0x4c, 0x32, 0x62, 0x08, 0xe4, 0x77, 0x34, 0x40, 0x7b, 0x70,
	0x4b, 0xe3, 0x75, 0x10, 0x31, 0x15, 0x2e, 0x45, 0x78, 0xbd, 0xe4, 0x2a, 0xc2, 0x3f, 0x9b, 0x9f,
	0x74, 0x9f, 0xe3, 0x82, 0xe5, 0xa5, 0x46, 0x2f, 0x35, 0xcc, 0x39, 0x4b, 0x90, 0x4b, 0x8a, 0x82,
	0x10, 0xa8, 0x3d, 0xb0, 0x18, 0xad, 0xca, 0xa8, 0x32, 0x6e, 0x2c, 0xb6, 0xff, 0x9d, 0x1e, 0x74,
	0xa6, 0x28, 0x02, 0x4e, 0x7d, 0xf4, 0x94, 0x7f, 0x21, 0x90, 0x8b, 0x93, 0x38, 0x51, 0xd2, 0x41,
	0xe8, 0xe6, 0x85, 0x33, 0x25, 0x13, 0x25, 0x49, 0x17, 0xea, 0xab, 0x40, 0x52, 0x16, 0xeb, 0x46,
	0x7a, 0x45, 0xbe, 0xc3, 0x17, 0x8e, 0x72, 0x19, 0xb0, 0x10, 0xad, 0x4f, 0xa3, 0xca, 0xf8, 0xf3,
	0xc2, 0xe0, 0x28, 0x27, 0x2c, 0x44, 0x62, 0x81, 0x71, 0x83, 0x42, 0xac, 0xd6, 0x68, 0x55, 0xb7,
	0x9e, 0xdd, 0xd2, 0x31, 0x81, 0x4c, 0x38, 0xae, 0xe4, 0x2e, 0x24, 0x0d, 0xf7, 0xa1, 0x9d, 0xa9,
	0x96, 0x91, 0x3c, 0x84, 0xfe, 0x8c, 0x85, 0xf4, 0xf2, 0x5e, 0x67, 0xfc, 0x97, 0x92, 0x53, 0x5f,
	0x49, 0xd4, 0x2f, 0x20, 0x82, 0x41, 0x81, 0x5e, 0xc6, 0x69, 0x3a, 0xd0, 0x9e, 0x62, 0x84, 0x32,
	0x37, 0x85, 0x00, 0xcc, 0x6c, 0xb9, 0x8c, 0xec, 0x2e, 0x98, 0x0b, 0x14, 0x92, 0xf1, 0x5c, 0x78,
	0x08, 0x9d, 0x5c, 0xbd, 0x84, 0xf4, 0xbf, 0x4f, 0x55, 0x68, 0x65, 0x91, 0x25, 0x1e, 0x7c, 0xcb,
	0xb3, 0x47, 0x86, 0xae, 0xfe, 0x02, 0xdc, 0xa3, 0xbc, 0xda, 0x3f, 0x0b, 0x75, 0x7d, 0xe8, 0x53,
	0x68, 0x66, 0x98, 0x22, 0x3f, 0xf6, 0x8e, 0x43, 0x02, 0xed, 0xfe, 0x71, 0x51, 0xf7, 0xba, 0x82,
	0x5e, 0x01, 0x1b, 0xe4, 0xd7, 0xde, 0xf8, 0x1e, 0x5d, 0xf6, 0xef, 0x8f, 0xb6, 0xe9, 0xa4, 0x19,
	0xb4, 0xb2, 0x00, 0x90, 0xfe, 0x9b, 0x07, 0x3d, 0x00, 0xc6, 0x1e, 0x14, 0xa8, 0xba, 0xdd, 0x1c,
	0xbe, 0xe6, 0x46, 0x4a, 0x5e, 0x1d, 0xc7, 0x20, 0xb0, 0x87, 0x45, 0x72, 0xda, 0xd1, 0xb6, 0x1e,
	0x9f, 0x6b, 0x26, 0x34, 0xce, 0x69, 0xbc, 0x9e, 0x6c, 0x6e, 0x25, 0x62, 0xe8, 0x61, 0xfa, 0xf5,
	0xed, 0x8d, 0xf4, 0xef, 0x25, 0x00, 0x00, 0xff, 0xff, 0x59, 0x18, 0xd4, 0xd1, 0xca, 0x04, 0x00,
	0x00,
}
