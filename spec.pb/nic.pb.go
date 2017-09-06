// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nic.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import request_data_pkg "github.com/chai2010/qingcloud-go/request/data"
import "github.com/chai2010/qingcloud-go/request/errors"

var _ = config.Config{}
var _ = request.Request{}
var _ = request_data_pkg.Operation{}
var _ = errors.ParameterRequiredError{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NicServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *NicServiceProperties) Reset()                    { *m = NicServiceProperties{} }
func (m *NicServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*NicServiceProperties) ProtoMessage()               {}
func (*NicServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *NicServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type CreateNicsInput struct {
}

func (m *CreateNicsInput) Reset()                    { *m = CreateNicsInput{} }
func (m *CreateNicsInput) String() string            { return proto.CompactTextString(m) }
func (*CreateNicsInput) ProtoMessage()               {}
func (*CreateNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

type CreateNicsOutput struct {
}

func (m *CreateNicsOutput) Reset()                    { *m = CreateNicsOutput{} }
func (m *CreateNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateNicsOutput) ProtoMessage()               {}
func (*CreateNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

type DescribeNicsInput struct {
	Instances []string `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	Limit     int32    `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	NicName   string   `protobuf:"bytes,3,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Nics      []string `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
	Offset    int32    `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Status    string   `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	VxnetType int32    `protobuf:"varint,7,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	Vxnets    []string `protobuf:"bytes,8,rep,name=vxnets" json:"vxnets,omitempty"`
}

func (m *DescribeNicsInput) Reset()                    { *m = DescribeNicsInput{} }
func (m *DescribeNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeNicsInput) ProtoMessage()               {}
func (*DescribeNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

func (m *DescribeNicsInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *DescribeNicsInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeNicsInput) GetNicName() string {
	if m != nil {
		return m.NicName
	}
	return ""
}

func (m *DescribeNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *DescribeNicsInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeNicsInput) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DescribeNicsInput) GetVxnetType() int32 {
	if m != nil {
		return m.VxnetType
	}
	return 0
}

func (m *DescribeNicsInput) GetVxnets() []string {
	if m != nil {
		return m.Vxnets
	}
	return nil
}

type DescribeNicsOutput struct {
	Message    string  `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Action     string  `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	NicSet     []int32 `protobuf:"varint,3,rep,packed,name=nic_set,json=nicSet" json:"nic_set,omitempty"`
	RetCode    int32   `protobuf:"varint,4,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	TotalCount int32   `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeNicsOutput) Reset()                    { *m = DescribeNicsOutput{} }
func (m *DescribeNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeNicsOutput) ProtoMessage()               {}
func (*DescribeNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

func (m *DescribeNicsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeNicsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeNicsOutput) GetNicSet() []int32 {
	if m != nil {
		return m.NicSet
	}
	return nil
}

func (m *DescribeNicsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeNicsOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type NIC struct {
	CreateTime    *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	InstanceId    string                      `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	NicId         string                      `protobuf:"bytes,3,opt,name=nic_id,json=nicId" json:"nic_id,omitempty"`
	NicName       string                      `protobuf:"bytes,4,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Owner         string                      `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	PrivateIp     string                      `protobuf:"bytes,6,opt,name=private_ip,json=privateIp" json:"private_ip,omitempty"`
	Role          int32                       `protobuf:"varint,7,opt,name=role" json:"role,omitempty"`
	RootUserId    string                      `protobuf:"bytes,8,opt,name=root_user_id,json=rootUserId" json:"root_user_id,omitempty"`
	SecurityGroup string                      `protobuf:"bytes,9,opt,name=security_group,json=securityGroup" json:"security_group,omitempty"`
	Sequence      int32                       `protobuf:"varint,10,opt,name=sequence" json:"sequence,omitempty"`
	Status        string                      `protobuf:"bytes,11,opt,name=status" json:"status,omitempty"`
	StatusTime    *google_protobuf2.Timestamp `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	Tags          []*Tag                      `protobuf:"bytes,13,rep,name=tags" json:"tags,omitempty"`
	VxnetId       string                      `protobuf:"bytes,14,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
}

func (m *NIC) Reset()                    { *m = NIC{} }
func (m *NIC) String() string            { return proto.CompactTextString(m) }
func (*NIC) ProtoMessage()               {}
func (*NIC) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{5} }

func (m *NIC) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *NIC) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *NIC) GetNicId() string {
	if m != nil {
		return m.NicId
	}
	return ""
}

func (m *NIC) GetNicName() string {
	if m != nil {
		return m.NicName
	}
	return ""
}

func (m *NIC) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NIC) GetPrivateIp() string {
	if m != nil {
		return m.PrivateIp
	}
	return ""
}

func (m *NIC) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *NIC) GetRootUserId() string {
	if m != nil {
		return m.RootUserId
	}
	return ""
}

func (m *NIC) GetSecurityGroup() string {
	if m != nil {
		return m.SecurityGroup
	}
	return ""
}

func (m *NIC) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *NIC) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NIC) GetStatusTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *NIC) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *NIC) GetVxnetId() string {
	if m != nil {
		return m.VxnetId
	}
	return ""
}

type Tag struct {
	Color             string                      `protobuf:"bytes,1,opt,name=color" json:"color,omitempty"`
	CreateTime        *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	Description       string                      `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Owner             string                      `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	ResourceCount     int32                       `protobuf:"varint,5,opt,name=resource_count,json=resourceCount" json:"resource_count,omitempty"`
	ResourceTagPairs  []*ResourceTagPair          `protobuf:"bytes,6,rep,name=resource_tag_pairs,json=resourceTagPairs" json:"resource_tag_pairs,omitempty"`
	ResourceTypeCount []*ResourceTypeCount        `protobuf:"bytes,7,rep,name=resource_type_count,json=resourceTypeCount" json:"resource_type_count,omitempty"`
	TagId             string                      `protobuf:"bytes,8,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	TagKey            string                      `protobuf:"bytes,9,opt,name=tag_key,json=tagKey" json:"tag_key,omitempty"`
	TagName           string                      `protobuf:"bytes,10,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{6} }

func (m *Tag) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Tag) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Tag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tag) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Tag) GetResourceCount() int32 {
	if m != nil {
		return m.ResourceCount
	}
	return 0
}

func (m *Tag) GetResourceTagPairs() []*ResourceTagPair {
	if m != nil {
		return m.ResourceTagPairs
	}
	return nil
}

func (m *Tag) GetResourceTypeCount() []*ResourceTypeCount {
	if m != nil {
		return m.ResourceTypeCount
	}
	return nil
}

func (m *Tag) GetTagId() string {
	if m != nil {
		return m.TagId
	}
	return ""
}

func (m *Tag) GetTagKey() string {
	if m != nil {
		return m.TagKey
	}
	return ""
}

func (m *Tag) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

type ResourceTagPair struct {
	ResourceId   string                      `protobuf:"bytes,1,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	ResourceType string                      `protobuf:"bytes,2,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	Status       string                      `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	StatusTime   *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	TagId        string                      `protobuf:"bytes,5,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
}

func (m *ResourceTagPair) Reset()                    { *m = ResourceTagPair{} }
func (m *ResourceTagPair) String() string            { return proto.CompactTextString(m) }
func (*ResourceTagPair) ProtoMessage()               {}
func (*ResourceTagPair) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{7} }

func (m *ResourceTagPair) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ResourceTagPair) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceTagPair) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ResourceTagPair) GetStatusTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *ResourceTagPair) GetTagId() string {
	if m != nil {
		return m.TagId
	}
	return ""
}

type ResourceTypeCount struct {
	Count        int32  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
}

func (m *ResourceTypeCount) Reset()                    { *m = ResourceTypeCount{} }
func (m *ResourceTypeCount) String() string            { return proto.CompactTextString(m) }
func (*ResourceTypeCount) ProtoMessage()               {}
func (*ResourceTypeCount) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{8} }

func (m *ResourceTypeCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ResourceTypeCount) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

type AttachNicsInput struct {
}

func (m *AttachNicsInput) Reset()                    { *m = AttachNicsInput{} }
func (m *AttachNicsInput) String() string            { return proto.CompactTextString(m) }
func (*AttachNicsInput) ProtoMessage()               {}
func (*AttachNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{9} }

type AttachNicsOutput struct {
}

func (m *AttachNicsOutput) Reset()                    { *m = AttachNicsOutput{} }
func (m *AttachNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachNicsOutput) ProtoMessage()               {}
func (*AttachNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{10} }

type DetachNicsInput struct {
}

func (m *DetachNicsInput) Reset()                    { *m = DetachNicsInput{} }
func (m *DetachNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DetachNicsInput) ProtoMessage()               {}
func (*DetachNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{11} }

type DetachNicsOutput struct {
}

func (m *DetachNicsOutput) Reset()                    { *m = DetachNicsOutput{} }
func (m *DetachNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachNicsOutput) ProtoMessage()               {}
func (*DetachNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{12} }

type ModifyNicAttributesInput struct {
}

func (m *ModifyNicAttributesInput) Reset()                    { *m = ModifyNicAttributesInput{} }
func (m *ModifyNicAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyNicAttributesInput) ProtoMessage()               {}
func (*ModifyNicAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{13} }

type ModifyNicAttributesOutput struct {
}

func (m *ModifyNicAttributesOutput) Reset()                    { *m = ModifyNicAttributesOutput{} }
func (m *ModifyNicAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyNicAttributesOutput) ProtoMessage()               {}
func (*ModifyNicAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{14} }

type DeleteNicsInput struct {
}

func (m *DeleteNicsInput) Reset()                    { *m = DeleteNicsInput{} }
func (m *DeleteNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteNicsInput) ProtoMessage()               {}
func (*DeleteNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{15} }

type DeleteNicsOutput struct {
}

func (m *DeleteNicsOutput) Reset()                    { *m = DeleteNicsOutput{} }
func (m *DeleteNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteNicsOutput) ProtoMessage()               {}
func (*DeleteNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{16} }

func init() {
	proto.RegisterType((*NicServiceProperties)(nil), "spec.NicServiceProperties")
	proto.RegisterType((*CreateNicsInput)(nil), "spec.CreateNicsInput")
	proto.RegisterType((*CreateNicsOutput)(nil), "spec.CreateNicsOutput")
	proto.RegisterType((*DescribeNicsInput)(nil), "spec.DescribeNicsInput")
	proto.RegisterType((*DescribeNicsOutput)(nil), "spec.DescribeNicsOutput")
	proto.RegisterType((*NIC)(nil), "spec.NIC")
	proto.RegisterType((*Tag)(nil), "spec.Tag")
	proto.RegisterType((*ResourceTagPair)(nil), "spec.ResourceTagPair")
	proto.RegisterType((*ResourceTypeCount)(nil), "spec.ResourceTypeCount")
	proto.RegisterType((*AttachNicsInput)(nil), "spec.AttachNicsInput")
	proto.RegisterType((*AttachNicsOutput)(nil), "spec.AttachNicsOutput")
	proto.RegisterType((*DetachNicsInput)(nil), "spec.DetachNicsInput")
	proto.RegisterType((*DetachNicsOutput)(nil), "spec.DetachNicsOutput")
	proto.RegisterType((*ModifyNicAttributesInput)(nil), "spec.ModifyNicAttributesInput")
	proto.RegisterType((*ModifyNicAttributesOutput)(nil), "spec.ModifyNicAttributesOutput")
	proto.RegisterType((*DeleteNicsInput)(nil), "spec.DeleteNicsInput")
	proto.RegisterType((*DeleteNicsOutput)(nil), "spec.DeleteNicsOutput")
}

type NicServiceInterface interface {
	CreateNics(in *CreateNicsInput) (out *CreateNicsOutput, err error)
	DescribeNics(in *DescribeNicsInput) (out *DescribeNicsOutput, err error)
	AttachNics(in *AttachNicsInput) (out *AttachNicsOutput, err error)
	DetachNics(in *DetachNicsInput) (out *DetachNicsOutput, err error)
	ModifyNicAttributes(in *ModifyNicAttributesInput) (out *ModifyNicAttributesOutput, err error)
	DeleteNics(in *DeleteNicsInput) (out *DeleteNicsOutput, err error)
}

type NicService struct {
	Config     *config.Config
	Properties *NicServiceProperties
}

func NewNicService(conf *config.Config, zone string) (p *NicService, err error) {
	return &NicService{
		Config:     conf,
		Properties: &NicServiceProperties{Zone: zone},
	}, nil
}

func (p *NicService) CreateNics(in *CreateNicsInput) (out *CreateNicsOutput, err error) {
	if in == nil {
		in = &CreateNicsInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateNics",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateNicsOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
func (p *NicService) DescribeNics(in *DescribeNicsInput) (out *DescribeNicsOutput, err error) {
	if in == nil {
		in = &DescribeNicsInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeNics",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeNicsOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
func (p *NicService) AttachNics(in *AttachNicsInput) (out *AttachNicsOutput, err error) {
	if in == nil {
		in = &AttachNicsInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachNics",
		RequestMethod: "GET", // GET or POST
	}

	x := &AttachNicsOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
func (p *NicService) DetachNics(in *DetachNicsInput) (out *DetachNicsOutput, err error) {
	if in == nil {
		in = &DetachNicsInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachNics",
		RequestMethod: "GET", // GET or POST
	}

	x := &DetachNicsOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
func (p *NicService) ModifyNicAttributes(in *ModifyNicAttributesInput) (out *ModifyNicAttributesOutput, err error) {
	if in == nil {
		in = &ModifyNicAttributesInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyNicAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyNicAttributesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
func (p *NicService) DeleteNics(in *DeleteNicsInput) (out *DeleteNicsOutput, err error) {
	if in == nil {
		in = &DeleteNicsInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteNics",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteNicsOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func init() { proto.RegisterFile("nic.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 950 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x85, 0x42, 0xc9, 0x32, 0xaf, 0xfc, 0x88, 0x26, 0x4e, 0x42, 0x2b, 0x4d, 0x2c, 0xa8, 0x08,
	0x60, 0x74, 0xa1, 0x00, 0xee, 0xd2, 0x2b, 0x43, 0x01, 0x02, 0xa1, 0xa8, 0x1a, 0xb0, 0x6a, 0xb7,
	0xc4, 0x78, 0x78, 0xcd, 0x0e, 0x2a, 0x71, 0xd8, 0x99, 0xa1, 0x53, 0xf5, 0x2b, 0xba, 0xcf, 0xe7,
	0xf4, 0x0f, 0xba, 0xef, 0xbf, 0x14, 0xf3, 0xa0, 0xf9, 0x70, 0xd2, 0xc6, 0x3b, 0xde, 0x73, 0x2f,
	0x67, 0xce, 0x9c, 0x73, 0xc8, 0x81, 0x30, 0xe7, 0x6c, 0x5e, 0x48, 0xa1, 0x05, 0xe9, 0xab, 0x02,
	0xd9, 0xe4, 0x34, 0x13, 0x22, 0xdb, 0xe0, 0x1b, 0x8b, 0x5d, 0x97, 0x37, 0x6f, 0x68, 0xbe, 0x73,
	0x03, 0x93, 0x17, 0xdd, 0x16, 0x6e, 0x0b, 0x5d, 0x35, 0xcf, 0xba, 0x4d, 0xcd, 0xb7, 0xa8, 0x34,
	0xdd, 0x16, 0x7e, 0xe0, 0x55, 0x77, 0xe0, 0x83, 0xa4, 0x45, 0x81, 0x52, 0xb9, 0xfe, 0xec, 0x1b,
	0x38, 0x59, 0x71, 0xf6, 0x23, 0xca, 0x5b, 0xce, 0xf0, 0xbd, 0x14, 0x05, 0x4a, 0xcd, 0x51, 0x11,
	0x02, 0xfd, 0x3f, 0x44, 0x8e, 0x51, 0x6f, 0xda, 0x3b, 0x0f, 0x63, 0xfb, 0x3c, 0x1b, 0xc3, 0xf1,
	0x42, 0x22, 0xd5, 0xb8, 0xe2, 0x4c, 0x2d, 0xf3, 0xa2, 0xd4, 0x33, 0x02, 0x8f, 0x6b, 0xe8, 0x87,
	0x52, 0x1b, 0xec, 0x9f, 0x1e, 0x8c, 0xdf, 0xa2, 0x62, 0x92, 0x5f, 0xd7, 0x93, 0xe4, 0x2b, 0x08,
	0x79, 0xae, 0x34, 0xcd, 0x19, 0xaa, 0xa8, 0x37, 0x0d, 0xce, 0xc3, 0xb8, 0x06, 0xc8, 0x09, 0x0c,
	0x36, 0x7c, 0xcb, 0x75, 0xf4, 0x68, 0xda, 0x3b, 0x1f, 0xc4, 0xae, 0x20, 0xa7, 0xb0, 0x9f, 0x73,
	0x96, 0xe4, 0x74, 0x8b, 0x51, 0x60, 0x89, 0x0c, 0x73, 0xce, 0x56, 0x74, 0x8b, 0x86, 0x5f, 0xce,
	0x99, 0x8a, 0xfa, 0x76, 0x25, 0xfb, 0x4c, 0x9e, 0xc1, 0x9e, 0xb8, 0xb9, 0x51, 0xa8, 0xa3, 0x81,
	0x5d, 0xc5, 0x57, 0x06, 0x57, 0x9a, 0xea, 0x52, 0x45, 0x7b, 0x76, 0x11, 0x5f, 0x91, 0x97, 0x00,
	0xb7, 0xbf, 0xe7, 0xa8, 0x13, 0xbd, 0x2b, 0x30, 0x1a, 0xda, 0x77, 0x42, 0x8b, 0xac, 0x77, 0x05,
	0x9a, 0xd7, 0x6c, 0xa1, 0xa2, 0x7d, 0xbb, 0x89, 0xaf, 0x66, 0x1f, 0x7b, 0x40, 0x9a, 0xe7, 0x73,
	0xc7, 0x26, 0x11, 0x0c, 0xb7, 0xa8, 0x14, 0xcd, 0x2a, 0xd1, 0xaa, 0xd2, 0x2c, 0x44, 0x99, 0xe6,
	0x22, 0xb7, 0xa7, 0x0b, 0x63, 0x5f, 0x91, 0xe7, 0x60, 0x8e, 0x93, 0x18, 0xc2, 0xc1, 0x34, 0x30,
	0x84, 0x73, 0x63, 0x85, 0x3d, 0xb7, 0x44, 0x9d, 0x30, 0x91, 0x62, 0xd4, 0xb7, 0xb4, 0x86, 0x12,
	0xf5, 0x42, 0xa4, 0x48, 0xce, 0x60, 0xa4, 0x85, 0xa6, 0x9b, 0x84, 0x89, 0x32, 0xaf, 0x0e, 0x0a,
	0x16, 0x5a, 0x18, 0x64, 0xf6, 0x77, 0x00, 0xc1, 0x6a, 0xb9, 0x20, 0x97, 0x30, 0x62, 0xd6, 0x99,
	0xc4, 0x44, 0xc2, 0x52, 0x1a, 0x5d, 0x4c, 0xe6, 0x2e, 0x0e, 0xf3, 0x2a, 0x0e, 0xf3, 0x75, 0x95,
	0x97, 0x18, 0xdc, 0xb8, 0x01, 0xcc, 0x2e, 0x95, 0x37, 0x09, 0x4f, 0x3d, 0x6d, 0xa8, 0xa0, 0x65,
	0x4a, 0x9e, 0x82, 0xe1, 0x6a, 0x7a, 0xce, 0x97, 0x41, 0xce, 0xd9, 0x32, 0x6d, 0x19, 0xd6, 0x6f,
	0x1b, 0x76, 0x02, 0x03, 0xf1, 0x21, 0x47, 0x69, 0x29, 0x87, 0xb1, 0x2b, 0x8c, 0x05, 0x85, 0xe4,
	0xb7, 0x86, 0x26, 0x2f, 0xbc, 0x3d, 0xa1, 0x47, 0x96, 0x85, 0x71, 0x59, 0x8a, 0x4d, 0xe5, 0x8d,
	0x7d, 0x26, 0x53, 0x38, 0x90, 0x42, 0xe8, 0xa4, 0x54, 0x28, 0x0d, 0x81, 0x7d, 0x47, 0xce, 0x60,
	0x3f, 0x29, 0x94, 0xcb, 0x94, 0xbc, 0x86, 0x23, 0x85, 0xac, 0x94, 0x5c, 0xef, 0x92, 0x4c, 0x8a,
	0xb2, 0x88, 0x42, 0x3b, 0x73, 0x58, 0xa1, 0xef, 0x0c, 0x48, 0x26, 0xb0, 0xaf, 0xf0, 0xb7, 0x12,
	0x73, 0x86, 0x11, 0xd8, 0x0d, 0xee, 0xea, 0x46, 0x64, 0x46, 0xad, 0xc8, 0x5c, 0xc2, 0xc8, 0x3d,
	0x39, 0x55, 0x0f, 0xfe, 0x5f, 0x55, 0x37, 0x6e, 0x55, 0x7d, 0x09, 0x7d, 0x4d, 0x33, 0x15, 0x1d,
	0x4e, 0x83, 0xf3, 0xd1, 0x45, 0x38, 0x37, 0x5f, 0xfe, 0x7c, 0x4d, 0xb3, 0xd8, 0xc2, 0x46, 0x3c,
	0x17, 0x47, 0x9e, 0x46, 0x47, 0x4e, 0x3c, 0x5b, 0x2f, 0xd3, 0xd9, 0x9f, 0x01, 0x04, 0x6b, 0x9a,
	0x19, 0x11, 0x99, 0xd8, 0x08, 0xe9, 0x13, 0xe6, 0x8a, 0xae, 0xd5, 0x8f, 0x1e, 0x64, 0xf5, 0x14,
	0x46, 0xa9, 0x0d, 0x73, 0x61, 0x13, 0xea, 0xec, 0x6c, 0x42, 0xb5, 0x73, 0xfd, 0xa6, 0x73, 0xaf,
	0xe1, 0x48, 0xa2, 0x12, 0xa5, 0x64, 0xd8, 0xca, 0xe2, 0x61, 0x85, 0xda, 0x38, 0x92, 0x05, 0x90,
	0xbb, 0x31, 0x4d, 0xb3, 0xa4, 0xa0, 0x5c, 0x9a, 0xef, 0xd0, 0x28, 0xf0, 0xd4, 0x29, 0x10, 0xfb,
	0xfe, 0x9a, 0x66, 0xef, 0x29, 0x97, 0xf1, 0x63, 0xd9, 0x06, 0x14, 0x79, 0x07, 0x4f, 0xea, 0x45,
	0x76, 0x45, 0xb5, 0xe1, 0xd0, 0xae, 0xf2, 0xbc, 0xb3, 0xca, 0xae, 0x70, 0x5b, 0xc7, 0x63, 0xd9,
	0x85, 0x4c, 0x6c, 0x0d, 0x89, 0xbb, 0xd4, 0x0c, 0x34, 0xcd, 0x96, 0xa9, 0xf9, 0x10, 0x0d, 0xfc,
	0x2b, 0xee, 0x7c, 0x52, 0xcc, 0xd4, 0x77, 0xb8, 0x33, 0x96, 0x98, 0x86, 0xcd, 0x33, 0x38, 0x4b,
	0x34, 0xcd, 0x4c, 0x9e, 0x67, 0x7f, 0xf5, 0xe0, 0xb8, 0xc3, 0xdc, 0x7c, 0x36, 0x77, 0x3c, 0x79,
	0xea, 0x4d, 0x82, 0x0a, 0x5a, 0xa6, 0xe4, 0x6b, 0x38, 0x6c, 0x1d, 0xc4, 0x7f, 0x59, 0x07, 0x4d,
	0xa6, 0x8d, 0xec, 0x05, 0xff, 0x95, 0xbd, 0xfe, 0x83, 0xb2, 0x57, 0x9f, 0x7c, 0xd0, 0x38, 0xf9,
	0x6c, 0x05, 0xe3, 0x7b, 0xc2, 0xb9, 0x94, 0x19, 0x81, 0x7b, 0xee, 0x67, 0x6c, 0x8b, 0x2f, 0xe2,
	0x6e, 0xae, 0x88, 0x2b, 0xad, 0x29, 0xfb, 0xa5, 0x75, 0x45, 0xd4, 0x90, 0xbf, 0x22, 0xc6, 0x70,
	0xfc, 0x16, 0xef, 0x8d, 0xd5, 0x90, 0x1f, 0x9b, 0x40, 0xf4, 0xbd, 0x48, 0xf9, 0xcd, 0x6e, 0xc5,
	0xd9, 0x95, 0xd6, 0x92, 0x5f, 0x97, 0x1a, 0xfd, 0xfc, 0x0b, 0x38, 0xfd, 0x44, 0xaf, 0xb9, 0xfe,
	0x06, 0x3b, 0x37, 0x55, 0x0d, 0xb9, 0xb1, 0x8b, 0x8f, 0x01, 0x40, 0x7d, 0xfb, 0x91, 0x4b, 0x80,
	0xfa, 0x32, 0x23, 0x3e, 0x9d, 0x9d, 0x1b, 0x6f, 0xf2, 0xac, 0x0b, 0xfb, 0xdf, 0xff, 0x15, 0x1c,
	0x34, 0x2f, 0x05, 0xe2, 0x63, 0x79, 0xef, 0x22, 0x9c, 0x44, 0xf7, 0x1b, 0x7e, 0x89, 0x4b, 0x80,
	0x5a, 0xa9, 0x6a, 0xff, 0x8e, 0x9c, 0xd5, 0xfe, 0x5d, 0x49, 0xcd, 0xcb, 0xb5, 0x7e, 0xd5, 0xcb,
	0x1d, 0x91, 0xab, 0x97, 0xbb, 0x42, 0x93, 0x9f, 0xe1, 0xc9, 0x27, 0xc4, 0x24, 0xaf, 0xdc, 0xf8,
	0xe7, 0x3c, 0x98, 0x9c, 0x7d, 0xb6, 0xdf, 0x24, 0x55, 0x89, 0x5e, 0x93, 0x6a, 0x39, 0x53, 0x93,
	0x6a, 0xbb, 0x73, 0xbd, 0x67, 0x23, 0xfd, 0xed, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x33,
	0x67, 0x63, 0x2d, 0x09, 0x00, 0x00,
}
