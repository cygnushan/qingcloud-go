// Code generated by protoc-gen-go. DO NOT EDIT.
// source: span.proto

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

type SpanServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SpanServiceProperties) Reset()                    { *m = SpanServiceProperties{} }
func (m *SpanServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SpanServiceProperties) ProtoMessage()               {}
func (*SpanServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *SpanServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type CreateSpanInput struct {
}

func (m *CreateSpanInput) Reset()                    { *m = CreateSpanInput{} }
func (m *CreateSpanInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSpanInput) ProtoMessage()               {}
func (*CreateSpanInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{1} }

type CreateSpanOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CreateSpanOutput) Reset()                    { *m = CreateSpanOutput{} }
func (m *CreateSpanOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSpanOutput) ProtoMessage()               {}
func (*CreateSpanOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{2} }

func (m *CreateSpanOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateSpanOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateSpanOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DescribeSpansInput struct {
}

func (m *DescribeSpansInput) Reset()                    { *m = DescribeSpansInput{} }
func (m *DescribeSpansInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSpansInput) ProtoMessage()               {}
func (*DescribeSpansInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{3} }

type DescribeSpansOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DescribeSpansOutput) Reset()                    { *m = DescribeSpansOutput{} }
func (m *DescribeSpansOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSpansOutput) ProtoMessage()               {}
func (*DescribeSpansOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{4} }

func (m *DescribeSpansOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeSpansOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeSpansOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteSpansInput struct {
}

func (m *DeleteSpansInput) Reset()                    { *m = DeleteSpansInput{} }
func (m *DeleteSpansInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSpansInput) ProtoMessage()               {}
func (*DeleteSpansInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{5} }

type DeleteSpansOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteSpansOutput) Reset()                    { *m = DeleteSpansOutput{} }
func (m *DeleteSpansOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSpansOutput) ProtoMessage()               {}
func (*DeleteSpansOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{6} }

func (m *DeleteSpansOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteSpansOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteSpansOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddSpanMembersInput struct {
}

func (m *AddSpanMembersInput) Reset()                    { *m = AddSpanMembersInput{} }
func (m *AddSpanMembersInput) String() string            { return proto.CompactTextString(m) }
func (*AddSpanMembersInput) ProtoMessage()               {}
func (*AddSpanMembersInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{7} }

type AddSpanMembersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AddSpanMembersOutput) Reset()                    { *m = AddSpanMembersOutput{} }
func (m *AddSpanMembersOutput) String() string            { return proto.CompactTextString(m) }
func (*AddSpanMembersOutput) ProtoMessage()               {}
func (*AddSpanMembersOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{8} }

func (m *AddSpanMembersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AddSpanMembersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AddSpanMembersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RemoveSpanMembersInput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *RemoveSpanMembersInput) Reset()                    { *m = RemoveSpanMembersInput{} }
func (m *RemoveSpanMembersInput) String() string            { return proto.CompactTextString(m) }
func (*RemoveSpanMembersInput) ProtoMessage()               {}
func (*RemoveSpanMembersInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{9} }

func (m *RemoveSpanMembersInput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RemoveSpanMembersInput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RemoveSpanMembersInput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RemoveSpanMembersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *RemoveSpanMembersOutput) Reset()                    { *m = RemoveSpanMembersOutput{} }
func (m *RemoveSpanMembersOutput) String() string            { return proto.CompactTextString(m) }
func (*RemoveSpanMembersOutput) ProtoMessage()               {}
func (*RemoveSpanMembersOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{10} }

func (m *RemoveSpanMembersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RemoveSpanMembersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RemoveSpanMembersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ModifySpanAttributesInput struct {
}

func (m *ModifySpanAttributesInput) Reset()                    { *m = ModifySpanAttributesInput{} }
func (m *ModifySpanAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySpanAttributesInput) ProtoMessage()               {}
func (*ModifySpanAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{11} }

type ModifySpanAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifySpanAttributesOutput) Reset()                    { *m = ModifySpanAttributesOutput{} }
func (m *ModifySpanAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifySpanAttributesOutput) ProtoMessage()               {}
func (*ModifySpanAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{12} }

func (m *ModifySpanAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifySpanAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifySpanAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateSpanInput struct {
}

func (m *UpdateSpanInput) Reset()                    { *m = UpdateSpanInput{} }
func (m *UpdateSpanInput) String() string            { return proto.CompactTextString(m) }
func (*UpdateSpanInput) ProtoMessage()               {}
func (*UpdateSpanInput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{13} }

type UpdateSpanOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateSpanOutput) Reset()                    { *m = UpdateSpanOutput{} }
func (m *UpdateSpanOutput) String() string            { return proto.CompactTextString(m) }
func (*UpdateSpanOutput) ProtoMessage()               {}
func (*UpdateSpanOutput) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{14} }

func (m *UpdateSpanOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *UpdateSpanOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *UpdateSpanOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SpanServiceProperties)(nil), "service.SpanServiceProperties")
	proto.RegisterType((*CreateSpanInput)(nil), "service.CreateSpanInput")
	proto.RegisterType((*CreateSpanOutput)(nil), "service.CreateSpanOutput")
	proto.RegisterType((*DescribeSpansInput)(nil), "service.DescribeSpansInput")
	proto.RegisterType((*DescribeSpansOutput)(nil), "service.DescribeSpansOutput")
	proto.RegisterType((*DeleteSpansInput)(nil), "service.DeleteSpansInput")
	proto.RegisterType((*DeleteSpansOutput)(nil), "service.DeleteSpansOutput")
	proto.RegisterType((*AddSpanMembersInput)(nil), "service.AddSpanMembersInput")
	proto.RegisterType((*AddSpanMembersOutput)(nil), "service.AddSpanMembersOutput")
	proto.RegisterType((*RemoveSpanMembersInput)(nil), "service.RemoveSpanMembersInput")
	proto.RegisterType((*RemoveSpanMembersOutput)(nil), "service.RemoveSpanMembersOutput")
	proto.RegisterType((*ModifySpanAttributesInput)(nil), "service.ModifySpanAttributesInput")
	proto.RegisterType((*ModifySpanAttributesOutput)(nil), "service.ModifySpanAttributesOutput")
	proto.RegisterType((*UpdateSpanInput)(nil), "service.UpdateSpanInput")
	proto.RegisterType((*UpdateSpanOutput)(nil), "service.UpdateSpanOutput")
}

type SpanServiceInterface interface {
	CreateSpan(in *CreateSpanInput) (out *CreateSpanOutput, err error)
	DescribeSpans(in *DescribeSpansInput) (out *DescribeSpansOutput, err error)
	DeleteSpans(in *DeleteSpansInput) (out *DeleteSpansOutput, err error)
	AddSpanMembers(in *AddSpanMembersInput) (out *AddSpanMembersOutput, err error)
	RemoveSpanMembers(in *RemoveSpanMembersInput) (out *RemoveSpanMembersOutput, err error)
	ModifySpanAttributes(in *ModifySpanAttributesInput) (out *ModifySpanAttributesOutput, err error)
	UpdateSpan(in *UpdateSpanInput) (out *UpdateSpanOutput, err error)
}

type SpanService struct {
	Config           *config.Config
	Properties       *SpanServiceProperties
	LastResponseBody string
}

func NewSpanService(conf *config.Config, zone string) (p *SpanService) {
	return &SpanService{
		Config:     conf,
		Properties: &SpanServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Span(zone string) (*SpanService, error) {
	properties := &SpanServiceProperties{
		Zone: zone,
	}

	return &SpanService{Config: s.Config, Properties: properties}, nil
}

func (p *SpanService) CreateSpan(in *CreateSpanInput) (out *CreateSpanOutput, err error) {
	if in == nil {
		in = &CreateSpanInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSpan",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateSpanOutput{}
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

func (p *CreateSpanInput) Validate() error {
	return nil
}

func (p *SpanService) DescribeSpans(in *DescribeSpansInput) (out *DescribeSpansOutput, err error) {
	if in == nil {
		in = &DescribeSpansInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSpans",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeSpansOutput{}
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

func (p *DescribeSpansInput) Validate() error {
	return nil
}

func (p *SpanService) DeleteSpans(in *DeleteSpansInput) (out *DeleteSpansOutput, err error) {
	if in == nil {
		in = &DeleteSpansInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSpans",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteSpansOutput{}
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

func (p *DeleteSpansInput) Validate() error {
	return nil
}

func (p *SpanService) AddSpanMembers(in *AddSpanMembersInput) (out *AddSpanMembersOutput, err error) {
	if in == nil {
		in = &AddSpanMembersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddSpanMembers",
		RequestMethod: "GET", // GET or POST
	}

	x := &AddSpanMembersOutput{}
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

func (p *AddSpanMembersInput) Validate() error {
	return nil
}

func (p *SpanService) RemoveSpanMembers(in *RemoveSpanMembersInput) (out *RemoveSpanMembersOutput, err error) {
	if in == nil {
		in = &RemoveSpanMembersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RemoveSpanMembers",
		RequestMethod: "GET", // GET or POST
	}

	x := &RemoveSpanMembersOutput{}
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

func (p *RemoveSpanMembersInput) Validate() error {
	return nil
}

func (p *SpanService) ModifySpanAttributes(in *ModifySpanAttributesInput) (out *ModifySpanAttributesOutput, err error) {
	if in == nil {
		in = &ModifySpanAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySpanAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifySpanAttributesOutput{}
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

func (p *ModifySpanAttributesInput) Validate() error {
	return nil
}

func (p *SpanService) UpdateSpan(in *UpdateSpanInput) (out *UpdateSpanOutput, err error) {
	if in == nil {
		in = &UpdateSpanInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateSpan",
		RequestMethod: "GET", // GET or POST
	}

	x := &UpdateSpanOutput{}
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

func (p *UpdateSpanInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("span.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x55, 0x28, 0x2b, 0x3b, 0x13, 0x6c, 0x39, 0x6b, 0x47, 0xea, 0x31, 0x51, 0x85, 0x9b,
	0x49, 0x48, 0x45, 0x82, 0x27, 0x88, 0xda, 0x1b, 0x90, 0x2a, 0xa0, 0x13, 0xdc, 0x86, 0x24, 0x3e,
	0xab, 0x2c, 0xda, 0x38, 0xd8, 0xce, 0x24, 0x78, 0x04, 0x9e, 0x8b, 0x17, 0xe2, 0x0d, 0x50, 0x52,
	0xd3, 0x7c, 0x34, 0xe5, 0x6a, 0xb9, 0xa9, 0xea, 0xf3, 0xf5, 0x73, 0xe2, 0xff, 0x3f, 0x06, 0xd0,
	0x69, 0x98, 0x4c, 0x53, 0x25, 0x8d, 0xc4, 0x81, 0x26, 0x75, 0x27, 0x62, 0x62, 0x57, 0xdf, 0x45,
	0xb2, 0x8a, 0xd7, 0x32, 0xe3, 0x81, 0xe6, 0xdf, 0x02, 0x95, 0xad, 0xe9, 0x75, 0xfe, 0xb3, 0xad,
	0xf3, 0x5e, 0xc1, 0xe8, 0x26, 0x0d, 0x93, 0x9b, 0x6d, 0xf5, 0x47, 0x25, 0x53, 0x52, 0x46, 0x90,
	0x46, 0x84, 0xfe, 0x4f, 0x99, 0x90, 0xdb, 0x9b, 0xf4, 0xae, 0x8f, 0x97, 0xc5, 0x7f, 0xcf, 0x81,
	0xd3, 0x99, 0xa2, 0xd0, 0x50, 0xde, 0xf2, 0x2e, 0x49, 0x33, 0xe3, 0x05, 0x70, 0x56, 0x86, 0x3e,
	0x64, 0x26, 0xcd, 0x0c, 0x5e, 0xc0, 0x51, 0x18, 0x1b, 0x21, 0x13, 0xdb, 0x6c, 0x57, 0x38, 0x86,
	0xc7, 0x8a, 0x4c, 0x10, 0x4b, 0x4e, 0xee, 0x83, 0x49, 0xef, 0xfa, 0xd1, 0x72, 0xa0, 0xc8, 0xcc,
	0x24, 0x27, 0x74, 0x61, 0xb0, 0x21, 0xad, 0xc3, 0x15, 0xb9, 0x0f, 0x8b, 0x9e, 0x7f, 0x4b, 0x6f,
	0x08, 0x38, 0x27, 0x1d, 0x2b, 0x11, 0x15, 0x08, 0xbd, 0xc5, 0x46, 0x70, 0x5e, 0x8b, 0x76, 0x41,
	0x46, 0x38, 0x9b, 0xd3, 0x9a, 0x4c, 0x95, 0xfb, 0x15, 0x9c, 0x4a, 0xac, 0x0b, 0xea, 0x08, 0xce,
	0x7d, 0xce, 0xf3, 0xf1, 0x0b, 0xda, 0x44, 0xa4, 0x2c, 0x38, 0x86, 0x61, 0x3d, 0xdc, 0x05, 0x9b,
	0xe0, 0x62, 0x49, 0x1b, 0x79, 0x47, 0x4d, 0xfc, 0xfd, 0x62, 0x6e, 0xe1, 0xd9, 0x1e, 0xa6, 0x8b,
	0xc7, 0xb9, 0x84, 0xf1, 0x42, 0x72, 0x71, 0xfb, 0x23, 0xe7, 0xf8, 0xc6, 0x28, 0x11, 0x65, 0x86,
	0xec, 0x0b, 0x15, 0xc0, 0xda, 0x92, 0x5d, 0xec, 0xc3, 0x81, 0xd3, 0xcf, 0x29, 0x6f, 0xda, 0xa6,
	0x0c, 0x75, 0xc0, 0x7c, 0xf3, 0xbb, 0x0f, 0x27, 0x15, 0x63, 0xa3, 0x0f, 0x50, 0xfa, 0x14, 0xdd,
	0xa9, 0xfd, 0x3c, 0x4c, 0x1b, 0x7e, 0x66, 0xe3, 0x96, 0x8c, 0xdd, 0xdf, 0x7b, 0x78, 0x52, 0xf3,
	0x1c, 0x5e, 0xee, 0x6a, 0xf7, 0x1d, 0xca, 0x9e, 0xb7, 0x27, 0xed, 0xac, 0x39, 0x9c, 0x54, 0x7c,
	0x84, 0xe3, 0x4a, 0x71, 0xdd, 0x71, 0x8c, 0xb5, 0xa5, 0xec, 0x94, 0x05, 0x3c, 0xad, 0x9b, 0x02,
	0x4b, 0x6a, 0x8b, 0x89, 0xd8, 0xd5, 0x81, 0xac, 0x1d, 0xf7, 0x05, 0x9c, 0x3d, 0x5d, 0xe2, 0x8b,
	0x5d, 0x4f, 0xbb, 0x35, 0xd8, 0xe4, 0x70, 0x81, 0x9d, 0x1b, 0xc0, 0xb0, 0x4d, 0x6a, 0xe8, 0xed,
	0x3a, 0x0f, 0xca, 0x94, 0xbd, 0xfc, 0x6f, 0x8d, 0x05, 0xf8, 0x00, 0xa5, 0x9a, 0x2a, 0x87, 0xdb,
	0x50, 0x5d, 0xe5, 0x70, 0x9b, 0xe2, 0x63, 0xa3, 0x5f, 0x7f, 0xfa, 0x0e, 0x1e, 0x7f, 0x12, 0xc9,
	0x6a, 0x96, 0x5f, 0x16, 0xac, 0x9f, 0x27, 0xa3, 0xa3, 0xe2, 0x96, 0x78, 0xfb, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0x21, 0x5a, 0xe0, 0x5b, 0x06, 0x00, 0x00,
}
