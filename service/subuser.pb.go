// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subuser.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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
	Users  []string `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	Status string   `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Offset int32    `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Limit  int32    `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeSubUsersInput) Reset()                    { *m = DescribeSubUsersInput{} }
func (m *DescribeSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSubUsersInput) ProtoMessage()               {}
func (*DescribeSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

func (m *DescribeSubUsersInput) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *DescribeSubUsersInput) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DescribeSubUsersInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeSubUsersInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeSubUsersOutput struct {
	Action     string                                 `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                                  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                                 `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	UserSet    []*DescribeSubUsersOutput_ResponseItem `protobuf:"bytes,4,rep,name=user_set,json=userSet" json:"user_set,omitempty"`
	TotalCount int32                                  `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
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

func (m *DescribeSubUsersOutput) GetUserSet() []*DescribeSubUsersOutput_ResponseItem {
	if m != nil {
		return m.UserSet
	}
	return nil
}

func (m *DescribeSubUsersOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeSubUsersOutput_ResponseItem struct {
	UserId     string                      `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName   string                      `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	CreateTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	Email      string                      `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Status     string                      `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *DescribeSubUsersOutput_ResponseItem) Reset()         { *m = DescribeSubUsersOutput_ResponseItem{} }
func (m *DescribeSubUsersOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeSubUsersOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeSubUsersOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor25, []int{2, 0}
}

func (m *DescribeSubUsersOutput_ResponseItem) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DescribeSubUsersOutput_ResponseItem) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *DescribeSubUsersOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeSubUsersOutput_ResponseItem) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *DescribeSubUsersOutput_ResponseItem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CreateSubUserInput struct {
	Email       string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Passwd      string `protobuf:"bytes,3,opt,name=passwd" json:"passwd,omitempty"`
	NotifyEmail string `protobuf:"bytes,4,opt,name=notify_email,json=notifyEmail" json:"notify_email,omitempty"`
}

func (m *CreateSubUserInput) Reset()                    { *m = CreateSubUserInput{} }
func (m *CreateSubUserInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserInput) ProtoMessage()               {}
func (*CreateSubUserInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

func (m *CreateSubUserInput) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateSubUserInput) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateSubUserInput) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *CreateSubUserInput) GetNotifyEmail() string {
	if m != nil {
		return m.NotifyEmail
	}
	return ""
}

type CreateSubUserOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	UserId  string `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
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

func (m *CreateSubUserOutput) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateSubUserOutput) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ModifySubUserAttributesInput struct {
	User        string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Nologin     int32  `protobuf:"varint,3,opt,name=nologin" json:"nologin,omitempty"`
	NotifyEmail string `protobuf:"bytes,4,opt,name=notify_email,json=notifyEmail" json:"notify_email,omitempty"`
}

func (m *ModifySubUserAttributesInput) Reset()                    { *m = ModifySubUserAttributesInput{} }
func (m *ModifySubUserAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesInput) ProtoMessage()               {}
func (*ModifySubUserAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

func (m *ModifySubUserAttributesInput) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ModifySubUserAttributesInput) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ModifySubUserAttributesInput) GetNologin() int32 {
	if m != nil {
		return m.Nologin
	}
	return 0
}

func (m *ModifySubUserAttributesInput) GetNotifyEmail() string {
	if m != nil {
		return m.NotifyEmail
	}
	return ""
}

type ModifySubUserAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	UserId  string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
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

func (m *ModifySubUserAttributesOutput) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DeleteSubUsersInput struct {
	Users []string `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *DeleteSubUsersInput) Reset()                    { *m = DeleteSubUsersInput{} }
func (m *DeleteSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersInput) ProtoMessage()               {}
func (*DeleteSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

func (m *DeleteSubUsersInput) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeleteSubUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Users   string `protobuf:"bytes,4,opt,name=users" json:"users,omitempty"`
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

func (m *DeleteSubUsersOutput) GetUsers() string {
	if m != nil {
		return m.Users
	}
	return ""
}

type RestoreSubUsersInput struct {
	Users []string `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *RestoreSubUsersInput) Reset()                    { *m = RestoreSubUsersInput{} }
func (m *RestoreSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersInput) ProtoMessage()               {}
func (*RestoreSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

func (m *RestoreSubUsersInput) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type RestoreSubUsersOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Users   string `protobuf:"bytes,4,opt,name=users" json:"users,omitempty"`
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

func (m *RestoreSubUsersOutput) GetUsers() string {
	if m != nil {
		return m.Users
	}
	return ""
}

func init() {
	proto.RegisterType((*SubuserServiceProperties)(nil), "service.SubuserServiceProperties")
	proto.RegisterType((*DescribeSubUsersInput)(nil), "service.DescribeSubUsersInput")
	proto.RegisterType((*DescribeSubUsersOutput)(nil), "service.DescribeSubUsersOutput")
	proto.RegisterType((*DescribeSubUsersOutput_ResponseItem)(nil), "service.DescribeSubUsersOutput.ResponseItem")
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
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x05, 0xab, 0x2f, 0x6b, 0x64, 0xbb, 0xed, 0xda, 0x96, 0x59, 0xfa, 0x4b, 0x15, 0xd0, 0x42,
	0x40, 0x0d, 0x12, 0x75, 0x6f, 0xed, 0xa9, 0x90, 0x93, 0xc0, 0x06, 0x9c, 0x38, 0x54, 0x72, 0x26,
	0x28, 0x72, 0x24, 0x2f, 0x42, 0x72, 0x19, 0xee, 0x32, 0x89, 0x7d, 0x4b, 0x80, 0x00, 0x41, 0x90,
	0x4b, 0x7e, 0x49, 0x7e, 0x5b, 0xfe, 0x41, 0xc0, 0xdd, 0x95, 0x42, 0xc9, 0x92, 0xed, 0x8b, 0x73,
	0x11, 0xf4, 0x66, 0x67, 0x67, 0xde, 0xbe, 0xb7, 0x3b, 0x84, 0x35, 0x9e, 0x0f, 0x73, 0x8e, 0x99,
	0x9d, 0x66, 0x4c, 0x30, 0xd2, 0xe0, 0x98, 0xbd, 0xa2, 0x01, 0x5a, 0x7b, 0x2f, 0x69, 0x32, 0x0e,
	0x22, 0x96, 0x87, 0x1e, 0x0f, 0x5f, 0x78, 0x59, 0x1e, 0xa1, 0x53, 0xfc, 0xa8, 0x3c, 0xeb, 0x60,
	0xcc, 0xd8, 0x38, 0x42, 0x47, 0xa2, 0x61, 0x3e, 0x72, 0x04, 0x8d, 0x91, 0x0b, 0x3f, 0x4e, 0x55,
	0x42, 0xd7, 0x06, 0x73, 0xa0, 0x2a, 0x0f, 0x54, 0xc5, 0xf3, 0x8c, 0xa5, 0x98, 0x09, 0x8a, 0x9c,
	0x10, 0xa8, 0x5e, 0xb1, 0x04, 0x4d, 0xa3, 0x63, 0xf4, 0x9a, 0xae, 0xfc, 0xdf, 0xe5, 0xb0, 0x75,
	0x8c, 0x3c, 0xc8, 0xe8, 0x10, 0x07, 0xf9, 0xf0, 0x39, 0xc7, 0x8c, 0x9f, 0x24, 0x69, 0x2e, 0xc8,
	0x26, 0xd4, 0x8a, 0x2a, 0xdc, 0x34, 0x3a, 0x95, 0x5e, 0xd3, 0x55, 0x80, 0xb4, 0xa1, 0xce, 0x85,
	0x2f, 0x72, 0x6e, 0xfe, 0x24, 0x8b, 0x68, 0x54, 0xc4, 0xd9, 0x68, 0xc4, 0x51, 0x98, 0x95, 0x8e,
	0xd1, 0xab, 0xb9, 0x1a, 0x15, 0x55, 0x22, 0x1a, 0x53, 0x61, 0x56, 0x65, 0x58, 0x81, 0xee, 0xdb,
	0x0a, 0xb4, 0xe7, 0xbb, 0x3e, 0xc9, 0x45, 0xd1, 0xb6, 0x0d, 0x75, 0x3f, 0x10, 0x94, 0x25, 0x9a,
	0xa5, 0x46, 0xe4, 0x37, 0x58, 0xc9, 0x50, 0x78, 0x01, 0x0b, 0x51, 0xb6, 0xae, 0xb9, 0x8d, 0x0c,
	0x45, 0x9f, 0x85, 0x48, 0x4c, 0x68, 0xc4, 0xc8, 0xb9, 0x3f, 0x46, 0xd9, 0xbc, 0xe9, 0x4e, 0x20,
	0x79, 0x04, 0x2b, 0x05, 0x6d, 0xaf, 0xe0, 0x55, 0xed, 0x54, 0x7a, 0xad, 0xa3, 0x43, 0x5b, 0x0b,
	0x6d, 0x2f, 0xee, 0x6f, 0xbb, 0xc8, 0x53, 0x96, 0x70, 0x3c, 0x11, 0x18, 0xbb, 0x0d, 0xa5, 0xa3,
	0x20, 0x07, 0xd0, 0x12, 0x4c, 0xf8, 0x91, 0x17, 0xb0, 0x3c, 0x11, 0x66, 0x4d, 0x12, 0x00, 0x19,
	0xea, 0x17, 0x11, 0xeb, 0x8b, 0x01, 0xab, 0xe5, 0xad, 0x64, 0x1b, 0xe4, 0x66, 0x8f, 0x86, 0x93,
	0x83, 0x14, 0xf0, 0x24, 0x24, 0x3b, 0xd0, 0x94, 0x0b, 0x89, 0x1f, 0xa3, 0x16, 0x51, 0x92, 0x7c,
	0xec, 0xc7, 0x48, 0xfe, 0x83, 0x56, 0x90, 0xa1, 0x2f, 0xd0, 0x2b, 0x7c, 0x95, 0xc7, 0x69, 0x1d,
	0x59, 0xb6, 0x32, 0xdd, 0x9e, 0x98, 0x6e, 0x3f, 0x9b, 0x98, 0xee, 0x82, 0x4a, 0x2f, 0x02, 0x85,
	0xd6, 0x18, 0xfb, 0x34, 0x92, 0x5a, 0x37, 0x5d, 0x05, 0x4a, 0x8e, 0xd5, 0xca, 0x8e, 0x75, 0xdf,
	0x19, 0x40, 0xfa, 0x72, 0xb3, 0x56, 0x60, 0x6a, 0xbb, 0x2a, 0x62, 0x94, 0x8b, 0xdc, 0x48, 0xba,
	0x0d, 0xf5, 0xd4, 0xe7, 0xfc, 0x75, 0xa8, 0xe5, 0xd7, 0x88, 0xfc, 0x0e, 0xab, 0x09, 0x13, 0x74,
	0x74, 0xe9, 0x95, 0x69, 0xb5, 0x54, 0xec, 0x41, 0x11, 0xea, 0x7e, 0x36, 0x60, 0x63, 0x86, 0xc4,
	0x7d, 0xdc, 0x82, 0xc5, 0xba, 0x94, 0x0c, 0xaa, 0x95, 0x0d, 0xea, 0x7e, 0x32, 0x60, 0xf7, 0x8c,
	0x85, 0x74, 0x74, 0xa9, 0x39, 0xfd, 0x2f, 0x44, 0x46, 0x87, 0xb9, 0x40, 0xfd, 0x32, 0x08, 0x54,
	0x8b, 0xd4, 0xc9, 0x33, 0x2a, 0xfe, 0xdf, 0x2c, 0x90, 0x09, 0x8d, 0x84, 0x45, 0x6c, 0x4c, 0x13,
	0xfd, 0x3a, 0x26, 0xf0, 0x2e, 0x12, 0xbd, 0x37, 0x60, 0x6f, 0x09, 0x9d, 0xfb, 0x10, 0xab, 0x24,
	0x4b, 0x75, 0x46, 0x96, 0xbf, 0x60, 0xe3, 0x18, 0x23, 0x14, 0x77, 0x19, 0x13, 0xdd, 0x4b, 0xd8,
	0x9c, 0x4d, 0xbe, 0x27, 0x5f, 0x55, 0x6b, 0xed, 0xab, 0x6a, 0x7d, 0x08, 0x9b, 0x2e, 0x72, 0xc1,
	0xb2, 0x3b, 0x11, 0xbd, 0x82, 0xad, 0xb9, 0xec, 0x1f, 0xc6, 0xf4, 0xe8, 0x43, 0x15, 0xd6, 0x67,
	0x67, 0x35, 0x19, 0xc0, 0x2f, 0xf3, 0x73, 0x89, 0xec, 0x2f, 0x1d, 0x59, 0xf2, 0x60, 0xd6, 0xc1,
	0x2d, 0x23, 0x8d, 0x9c, 0xc2, 0xda, 0xcc, 0x1b, 0x23, 0x3b, 0xd3, 0x1d, 0xd7, 0x07, 0x80, 0xb5,
	0xbb, 0x78, 0x51, 0xd7, 0xba, 0x80, 0xed, 0x25, 0x97, 0x91, 0xfc, 0x31, 0xdd, 0x78, 0xd3, 0xeb,
	0xb1, 0xfe, 0xbc, 0x2d, 0x4d, 0x77, 0x3a, 0x83, 0xf5, 0xd9, 0x2b, 0x44, 0x76, 0x4b, 0x07, 0xbd,
	0x76, 0x11, 0xad, 0xbd, 0x25, 0xab, 0xba, 0xdc, 0x39, 0xfc, 0x3c, 0x67, 0x34, 0xf9, 0xbe, 0x63,
	0xd1, 0x85, 0xb1, 0xf6, 0x97, 0x2d, 0xab, 0x8a, 0xd6, 0xe9, 0xc7, 0xaf, 0xd5, 0x87, 0xf0, 0xf7,
	0x85, 0x10, 0x29, 0xff, 0xd7, 0x71, 0x42, 0x16, 0x70, 0x7b, 0xfa, 0xf1, 0xb6, 0x03, 0x16, 0x3b,
	0x7e, 0x4a, 0x1d, 0xfd, 0x99, 0x77, 0x68, 0x12, 0xe2, 0x1b, 0xfb, 0x42, 0xc4, 0x11, 0xf9, 0xf5,
	0x29, 0x4d, 0xc6, 0x7d, 0x99, 0xa5, 0xed, 0x1f, 0xd6, 0xe5, 0x68, 0xff, 0xe7, 0x5b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x69, 0xa4, 0xd4, 0x6f, 0x16, 0x08, 0x00, 0x00,
}
