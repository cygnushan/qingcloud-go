// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eip.proto

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

type EipServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *EipServiceProperties) Reset()                    { *m = EipServiceProperties{} }
func (m *EipServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*EipServiceProperties) ProtoMessage()               {}
func (*EipServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *EipServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeEipsInput struct {
}

func (m *DescribeEipsInput) Reset()                    { *m = DescribeEipsInput{} }
func (m *DescribeEipsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeEipsInput) ProtoMessage()               {}
func (*DescribeEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type DescribeEipsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DescribeEipsOutput) Reset()                    { *m = DescribeEipsOutput{} }
func (m *DescribeEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeEipsOutput) ProtoMessage()               {}
func (*DescribeEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DescribeEipsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeEipsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeEipsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AllocateEipsInput struct {
}

func (m *AllocateEipsInput) Reset()                    { *m = AllocateEipsInput{} }
func (m *AllocateEipsInput) String() string            { return proto.CompactTextString(m) }
func (*AllocateEipsInput) ProtoMessage()               {}
func (*AllocateEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type AllocateEipsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AllocateEipsOutput) Reset()                    { *m = AllocateEipsOutput{} }
func (m *AllocateEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*AllocateEipsOutput) ProtoMessage()               {}
func (*AllocateEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *AllocateEipsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AllocateEipsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AllocateEipsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReleaseEipsInput struct {
}

func (m *ReleaseEipsInput) Reset()                    { *m = ReleaseEipsInput{} }
func (m *ReleaseEipsInput) String() string            { return proto.CompactTextString(m) }
func (*ReleaseEipsInput) ProtoMessage()               {}
func (*ReleaseEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type ReleaseEipsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ReleaseEipsOutput) Reset()                    { *m = ReleaseEipsOutput{} }
func (m *ReleaseEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*ReleaseEipsOutput) ProtoMessage()               {}
func (*ReleaseEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ReleaseEipsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ReleaseEipsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ReleaseEipsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AssociateEipInput struct {
}

func (m *AssociateEipInput) Reset()                    { *m = AssociateEipInput{} }
func (m *AssociateEipInput) String() string            { return proto.CompactTextString(m) }
func (*AssociateEipInput) ProtoMessage()               {}
func (*AssociateEipInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

type AssociateEipOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AssociateEipOutput) Reset()                    { *m = AssociateEipOutput{} }
func (m *AssociateEipOutput) String() string            { return proto.CompactTextString(m) }
func (*AssociateEipOutput) ProtoMessage()               {}
func (*AssociateEipOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *AssociateEipOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AssociateEipOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AssociateEipOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DissociateEipsInput struct {
}

func (m *DissociateEipsInput) Reset()                    { *m = DissociateEipsInput{} }
func (m *DissociateEipsInput) String() string            { return proto.CompactTextString(m) }
func (*DissociateEipsInput) ProtoMessage()               {}
func (*DissociateEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

type DissociateEipsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DissociateEipsOutput) Reset()                    { *m = DissociateEipsOutput{} }
func (m *DissociateEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*DissociateEipsOutput) ProtoMessage()               {}
func (*DissociateEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *DissociateEipsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DissociateEipsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DissociateEipsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ChangeEipsBandwidthInput struct {
}

func (m *ChangeEipsBandwidthInput) Reset()                    { *m = ChangeEipsBandwidthInput{} }
func (m *ChangeEipsBandwidthInput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBandwidthInput) ProtoMessage()               {}
func (*ChangeEipsBandwidthInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

type ChangeEipsBandwidthOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ChangeEipsBandwidthOutput) Reset()                    { *m = ChangeEipsBandwidthOutput{} }
func (m *ChangeEipsBandwidthOutput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBandwidthOutput) ProtoMessage()               {}
func (*ChangeEipsBandwidthOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *ChangeEipsBandwidthOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ChangeEipsBandwidthOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ChangeEipsBandwidthOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ChangeEipsBillingModeInput struct {
}

func (m *ChangeEipsBillingModeInput) Reset()                    { *m = ChangeEipsBillingModeInput{} }
func (m *ChangeEipsBillingModeInput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBillingModeInput) ProtoMessage()               {}
func (*ChangeEipsBillingModeInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

type ChangeEipsBillingModeOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ChangeEipsBillingModeOutput) Reset()                    { *m = ChangeEipsBillingModeOutput{} }
func (m *ChangeEipsBillingModeOutput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBillingModeOutput) ProtoMessage()               {}
func (*ChangeEipsBillingModeOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *ChangeEipsBillingModeOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ChangeEipsBillingModeOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ChangeEipsBillingModeOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ModifyEipAttributesInput struct {
}

func (m *ModifyEipAttributesInput) Reset()                    { *m = ModifyEipAttributesInput{} }
func (m *ModifyEipAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyEipAttributesInput) ProtoMessage()               {}
func (*ModifyEipAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

type ModifyEipAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifyEipAttributesOutput) Reset()                    { *m = ModifyEipAttributesOutput{} }
func (m *ModifyEipAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyEipAttributesOutput) ProtoMessage()               {}
func (*ModifyEipAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

func (m *ModifyEipAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifyEipAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifyEipAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EipServiceProperties)(nil), "service.EipServiceProperties")
	proto.RegisterType((*DescribeEipsInput)(nil), "service.DescribeEipsInput")
	proto.RegisterType((*DescribeEipsOutput)(nil), "service.DescribeEipsOutput")
	proto.RegisterType((*AllocateEipsInput)(nil), "service.AllocateEipsInput")
	proto.RegisterType((*AllocateEipsOutput)(nil), "service.AllocateEipsOutput")
	proto.RegisterType((*ReleaseEipsInput)(nil), "service.ReleaseEipsInput")
	proto.RegisterType((*ReleaseEipsOutput)(nil), "service.ReleaseEipsOutput")
	proto.RegisterType((*AssociateEipInput)(nil), "service.AssociateEipInput")
	proto.RegisterType((*AssociateEipOutput)(nil), "service.AssociateEipOutput")
	proto.RegisterType((*DissociateEipsInput)(nil), "service.DissociateEipsInput")
	proto.RegisterType((*DissociateEipsOutput)(nil), "service.DissociateEipsOutput")
	proto.RegisterType((*ChangeEipsBandwidthInput)(nil), "service.ChangeEipsBandwidthInput")
	proto.RegisterType((*ChangeEipsBandwidthOutput)(nil), "service.ChangeEipsBandwidthOutput")
	proto.RegisterType((*ChangeEipsBillingModeInput)(nil), "service.ChangeEipsBillingModeInput")
	proto.RegisterType((*ChangeEipsBillingModeOutput)(nil), "service.ChangeEipsBillingModeOutput")
	proto.RegisterType((*ModifyEipAttributesInput)(nil), "service.ModifyEipAttributesInput")
	proto.RegisterType((*ModifyEipAttributesOutput)(nil), "service.ModifyEipAttributesOutput")
}

type EipServiceInterface interface {
	DescribeEips(in *DescribeEipsInput) (out *DescribeEipsOutput, err error)
	AllocateEips(in *AllocateEipsInput) (out *AllocateEipsOutput, err error)
	ReleaseEips(in *ReleaseEipsInput) (out *ReleaseEipsOutput, err error)
	AssociateEip(in *AssociateEipInput) (out *AssociateEipOutput, err error)
	DissociateEips(in *DissociateEipsInput) (out *DissociateEipsOutput, err error)
	ChangeEipsBandwidth(in *ChangeEipsBandwidthInput) (out *ChangeEipsBandwidthOutput, err error)
	ChangeEipsBillingMode(in *ChangeEipsBillingModeInput) (out *ChangeEipsBillingModeOutput, err error)
	ModifyEipAttributes(in *ModifyEipAttributesInput) (out *ModifyEipAttributesOutput, err error)
}

type EipService struct {
	Config           *config.Config
	Properties       *EipServiceProperties
	LastResponseBody string
}

func NewEipService(conf *config.Config, zone string) (p *EipService) {
	return &EipService{
		Config:     conf,
		Properties: &EipServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Eip(zone string) (*EipService, error) {
	properties := &EipServiceProperties{
		Zone: zone,
	}

	return &EipService{Config: s.Config, Properties: properties}, nil
}

func (p *EipService) DescribeEips(in *DescribeEipsInput) (out *DescribeEipsOutput, err error) {
	if in == nil {
		in = &DescribeEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeEipsOutput{}
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

func (p *DescribeEipsInput) Validate() error {
	return nil
}

func (p *EipService) AllocateEips(in *AllocateEipsInput) (out *AllocateEipsOutput, err error) {
	if in == nil {
		in = &AllocateEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AllocateEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &AllocateEipsOutput{}
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

func (p *AllocateEipsInput) Validate() error {
	return nil
}

func (p *EipService) ReleaseEips(in *ReleaseEipsInput) (out *ReleaseEipsOutput, err error) {
	if in == nil {
		in = &ReleaseEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ReleaseEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &ReleaseEipsOutput{}
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

func (p *ReleaseEipsInput) Validate() error {
	return nil
}

func (p *EipService) AssociateEip(in *AssociateEipInput) (out *AssociateEipOutput, err error) {
	if in == nil {
		in = &AssociateEipInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEip",
		RequestMethod: "GET", // GET or POST
	}

	x := &AssociateEipOutput{}
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

func (p *AssociateEipInput) Validate() error {
	return nil
}

func (p *EipService) DissociateEips(in *DissociateEipsInput) (out *DissociateEipsOutput, err error) {
	if in == nil {
		in = &DissociateEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &DissociateEipsOutput{}
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

func (p *DissociateEipsInput) Validate() error {
	return nil
}

func (p *EipService) ChangeEipsBandwidth(in *ChangeEipsBandwidthInput) (out *ChangeEipsBandwidthOutput, err error) {
	if in == nil {
		in = &ChangeEipsBandwidthInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeEipsBandwidth",
		RequestMethod: "GET", // GET or POST
	}

	x := &ChangeEipsBandwidthOutput{}
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

func (p *ChangeEipsBandwidthInput) Validate() error {
	return nil
}

func (p *EipService) ChangeEipsBillingMode(in *ChangeEipsBillingModeInput) (out *ChangeEipsBillingModeOutput, err error) {
	if in == nil {
		in = &ChangeEipsBillingModeInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeEipsBillingMode",
		RequestMethod: "GET", // GET or POST
	}

	x := &ChangeEipsBillingModeOutput{}
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

func (p *ChangeEipsBillingModeInput) Validate() error {
	return nil
}

func (p *EipService) ModifyEipAttributes(in *ModifyEipAttributesInput) (out *ModifyEipAttributesOutput, err error) {
	if in == nil {
		in = &ModifyEipAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyEipAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyEipAttributesOutput{}
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

func (p *ModifyEipAttributesInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("eip.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xd9, 0x9f, 0xd2, 0x61, 0x85, 0x76, 0xdd, 0x2e, 0x4a, 0xbd, 0xbb, 0xd2, 0x52, 0x38,
	0xac, 0x38, 0x14, 0x09, 0x9e, 0x60, 0x69, 0xab, 0x15, 0x87, 0x0a, 0x28, 0x57, 0xa4, 0x92, 0x26,
	0x43, 0x3a, 0x10, 0x62, 0x63, 0x3b, 0x20, 0x78, 0x04, 0x1e, 0x87, 0x37, 0xe3, 0x0d, 0x50, 0x9b,
	0xd0, 0x4e, 0x5b, 0xa7, 0xb7, 0x5c, 0xa2, 0x78, 0x66, 0x3c, 0xdf, 0xf7, 0x59, 0xf3, 0xd9, 0xd0,
	0x42, 0xd2, 0x7d, 0x6d, 0x94, 0x53, 0xa2, 0x69, 0xd1, 0x7c, 0xa7, 0x08, 0xe5, 0xd5, 0x37, 0xca,
	0x92, 0x28, 0x55, 0x79, 0x3c, 0xb5, 0xf1, 0x97, 0xa9, 0xc9, 0x53, 0x7c, 0xbe, 0xf8, 0x14, 0x75,
	0xbd, 0x67, 0xd0, 0x19, 0x91, 0x7e, 0x5f, 0x14, 0xbf, 0x35, 0x4a, 0xa3, 0x71, 0x84, 0x56, 0x08,
	0x38, 0xfc, 0xa5, 0x32, 0x0c, 0x1a, 0xd7, 0x8d, 0x9b, 0xd6, 0x64, 0xf9, 0xdf, 0x6b, 0xc3, 0xd9,
	0x10, 0x6d, 0x64, 0x68, 0x86, 0x23, 0xd2, 0xf6, 0x75, 0xa6, 0x73, 0xd7, 0x0b, 0x41, 0xf0, 0xe0,
	0x9b, 0xdc, 0xe9, 0xdc, 0x89, 0x47, 0x70, 0x1c, 0x46, 0x8e, 0x54, 0x56, 0x36, 0x28, 0x57, 0xa2,
	0x0b, 0xf7, 0x0d, 0xba, 0x69, 0xa4, 0x62, 0x0c, 0xee, 0x5d, 0x37, 0x6e, 0x8e, 0x26, 0x4d, 0x83,
	0x6e, 0xa0, 0x62, 0x14, 0x01, 0x34, 0xbf, 0xa2, 0xb5, 0x61, 0x82, 0xc1, 0xc1, 0x72, 0xcf, 0xff,
	0xe5, 0x02, 0xf7, 0x36, 0x4d, 0x55, 0x14, 0xba, 0x4d, 0x5c, 0x1e, 0xac, 0x03, 0x57, 0xc0, 0xe9,
	0x04, 0x53, 0x0c, 0x2d, 0x83, 0xfd, 0x08, 0x67, 0x2c, 0x56, 0x97, 0x5a, 0x6b, 0x55, 0x44, 0x85,
	0xb2, 0xb5, 0x5a, 0x16, 0xac, 0x03, 0xf7, 0x1c, 0xda, 0x43, 0x62, 0x18, 0xa5, 0xe0, 0x08, 0x3a,
	0x9b, 0xe1, 0x3a, 0xb0, 0x25, 0x04, 0x83, 0x79, 0x98, 0x25, 0x4b, 0x80, 0x57, 0x61, 0x16, 0xff,
	0xa0, 0xd8, 0xcd, 0x0b, 0x02, 0x73, 0xe8, 0x7a, 0x72, 0x75, 0xb0, 0xb8, 0x04, 0xc9, 0x90, 0x28,
	0x4d, 0x29, 0x4b, 0xc6, 0x2a, 0xc6, 0x82, 0xc7, 0x67, 0xb8, 0xf0, 0x66, 0x6b, 0x3a, 0x8f, 0xb1,
	0x8a, 0xe9, 0xd3, 0xcf, 0x11, 0xe9, 0x5b, 0xe7, 0x0c, 0xcd, 0x72, 0x87, 0x76, 0x75, 0x1e, 0x9e,
	0x5c, 0x0d, 0x2c, 0x5e, 0xfc, 0x39, 0x02, 0x58, 0x5f, 0x0e, 0xe2, 0x0e, 0x4e, 0xb8, 0xd3, 0x85,
	0xec, 0x97, 0x77, 0x4c, 0x7f, 0xe7, 0x56, 0x90, 0x17, 0xde, 0x5c, 0x49, 0xf2, 0x0e, 0x4e, 0xb8,
	0x75, 0x59, 0xa3, 0x1d, 0x9b, 0xb3, 0x46, 0x1e, 0xb7, 0x0f, 0xe1, 0x01, 0x33, 0xa3, 0xe8, 0xae,
	0x6a, 0xb7, 0x6d, 0x2b, 0xa5, 0x2f, 0xc5, 0xe8, 0xb0, 0x01, 0xe7, 0x74, 0xb6, 0x7d, 0xc8, 0xe9,
	0xec, 0xda, 0x71, 0x0c, 0x0f, 0x37, 0xad, 0x22, 0x2e, 0xd7, 0xc7, 0xb0, 0x6b, 0x2d, 0x79, 0x55,
	0x91, 0x2d, 0xdb, 0x7d, 0x80, 0xb6, 0x67, 0xf0, 0xc5, 0xe3, 0xd5, 0xae, 0x2a, 0xcb, 0xc8, 0xde,
	0xbe, 0x92, 0xb2, 0xfb, 0x0c, 0xce, 0xbd, 0xe3, 0x2c, 0x9e, 0xf8, 0x36, 0x6f, 0x99, 0x41, 0x3e,
	0xdd, 0x5f, 0xb4, 0x56, 0xe0, 0x19, 0x55, 0xa6, 0xa0, 0x6a, 0xc8, 0x99, 0x82, 0xca, 0x59, 0x97,
	0x9d, 0xdf, 0x7f, 0x0f, 0x4f, 0x45, 0xeb, 0x1d, 0x65, 0xc9, 0x60, 0xf1, 0xbe, 0xc9, 0x83, 0x11,
	0xe9, 0xd9, 0xf1, 0xf2, 0x5d, 0x7b, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x42, 0xb2, 0x9f,
	0x0c, 0x07, 0x00, 0x00,
}
