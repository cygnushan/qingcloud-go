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
	Message    string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Action     string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	NicSet     []*NIC `protobuf:"bytes,3,rep,name=nic_set,json=nicSet" json:"nic_set,omitempty"`
	RetCode    int32  `protobuf:"varint,4,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	TotalCount int32  `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
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

func (m *DescribeNicsOutput) GetNicSet() []*NIC {
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
	CreateTime       *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	InstanceId       string                      `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	NicId            string                      `protobuf:"bytes,3,opt,name=nic_id,json=nicId" json:"nic_id,omitempty"`
	NicName          string                      `protobuf:"bytes,4,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Owner            string                      `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	PrivateIp        string                      `protobuf:"bytes,6,opt,name=private_ip,json=privateIp" json:"private_ip,omitempty"`
	Role             int32                       `protobuf:"varint,7,opt,name=role" json:"role,omitempty"`
	RootUserId       string                      `protobuf:"bytes,8,opt,name=root_user_id,json=rootUserId" json:"root_user_id,omitempty"`
	SecurityGroup    string                      `protobuf:"bytes,9,opt,name=security_group,json=securityGroup" json:"security_group,omitempty"`
	Sequence         int32                       `protobuf:"varint,10,opt,name=sequence" json:"sequence,omitempty"`
	Status           string                      `protobuf:"bytes,11,opt,name=status" json:"status,omitempty"`
	StatusTime       *google_protobuf2.Timestamp `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	Tags             []*Tag                      `protobuf:"bytes,13,rep,name=tags" json:"tags,omitempty"`
	VxnetId          string                      `protobuf:"bytes,14,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	Eip              *Eip                        `protobuf:"bytes,15,opt,name=eip" json:"eip,omitempty"`
	TransitionStatus string                      `protobuf:"bytes,16,opt,name=transition_status,json=transitionStatus" json:"transition_status,omitempty"`
	Controller       string                      `protobuf:"bytes,17,opt,name=controller" json:"controller,omitempty"`
	VxnetType        int32                       `protobuf:"varint,18,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	ConsoleId        string                      `protobuf:"bytes,19,opt,name=console_id,json=consoleId" json:"console_id,omitempty"`
	ResourceId       string                      `protobuf:"bytes,20,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
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

func (m *NIC) GetEip() *Eip {
	if m != nil {
		return m.Eip
	}
	return nil
}

func (m *NIC) GetTransitionStatus() string {
	if m != nil {
		return m.TransitionStatus
	}
	return ""
}

func (m *NIC) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *NIC) GetVxnetType() int32 {
	if m != nil {
		return m.VxnetType
	}
	return 0
}

func (m *NIC) GetConsoleId() string {
	if m != nil {
		return m.ConsoleId
	}
	return ""
}

func (m *NIC) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

type Eip struct {
	EipId     string `protobuf:"bytes,1,opt,name=eip_id,json=eipId" json:"eip_id,omitempty"`
	EipAddr   string `protobuf:"bytes,2,opt,name=eip_addr,json=eipAddr" json:"eip_addr,omitempty"`
	Bandwidth int32  `protobuf:"varint,3,opt,name=bandwidth" json:"bandwidth,omitempty"`
}

func (m *Eip) Reset()                    { *m = Eip{} }
func (m *Eip) String() string            { return proto.CompactTextString(m) }
func (*Eip) ProtoMessage()               {}
func (*Eip) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{6} }

func (m *Eip) GetEipId() string {
	if m != nil {
		return m.EipId
	}
	return ""
}

func (m *Eip) GetEipAddr() string {
	if m != nil {
		return m.EipAddr
	}
	return ""
}

func (m *Eip) GetBandwidth() int32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
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
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{7} }

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
func (*ResourceTagPair) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{8} }

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
func (*ResourceTypeCount) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{9} }

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
func (*AttachNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{10} }

type AttachNicsOutput struct {
}

func (m *AttachNicsOutput) Reset()                    { *m = AttachNicsOutput{} }
func (m *AttachNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachNicsOutput) ProtoMessage()               {}
func (*AttachNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{11} }

type DetachNicsInput struct {
}

func (m *DetachNicsInput) Reset()                    { *m = DetachNicsInput{} }
func (m *DetachNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DetachNicsInput) ProtoMessage()               {}
func (*DetachNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{12} }

type DetachNicsOutput struct {
}

func (m *DetachNicsOutput) Reset()                    { *m = DetachNicsOutput{} }
func (m *DetachNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachNicsOutput) ProtoMessage()               {}
func (*DetachNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{13} }

type ModifyNicAttributesInput struct {
}

func (m *ModifyNicAttributesInput) Reset()                    { *m = ModifyNicAttributesInput{} }
func (m *ModifyNicAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyNicAttributesInput) ProtoMessage()               {}
func (*ModifyNicAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{14} }

type ModifyNicAttributesOutput struct {
}

func (m *ModifyNicAttributesOutput) Reset()                    { *m = ModifyNicAttributesOutput{} }
func (m *ModifyNicAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyNicAttributesOutput) ProtoMessage()               {}
func (*ModifyNicAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{15} }

type DeleteNicsInput struct {
}

func (m *DeleteNicsInput) Reset()                    { *m = DeleteNicsInput{} }
func (m *DeleteNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteNicsInput) ProtoMessage()               {}
func (*DeleteNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{16} }

type DeleteNicsOutput struct {
}

func (m *DeleteNicsOutput) Reset()                    { *m = DeleteNicsOutput{} }
func (m *DeleteNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteNicsOutput) ProtoMessage()               {}
func (*DeleteNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{17} }

func init() {
	proto.RegisterType((*NicServiceProperties)(nil), "spec.NicServiceProperties")
	proto.RegisterType((*CreateNicsInput)(nil), "spec.CreateNicsInput")
	proto.RegisterType((*CreateNicsOutput)(nil), "spec.CreateNicsOutput")
	proto.RegisterType((*DescribeNicsInput)(nil), "spec.DescribeNicsInput")
	proto.RegisterType((*DescribeNicsOutput)(nil), "spec.DescribeNicsOutput")
	proto.RegisterType((*NIC)(nil), "spec.NIC")
	proto.RegisterType((*Eip)(nil), "spec.Eip")
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
	// 1077 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x86, 0x42, 0xc9, 0xb2, 0x46, 0x76, 0x6c, 0xad, 0x9d, 0x84, 0x96, 0x7f, 0xb6, 0x05, 0xfd,
	0x10, 0xc0, 0x68, 0x01, 0x05, 0x70, 0x8f, 0x3e, 0x19, 0x4a, 0x10, 0x10, 0x45, 0xd5, 0x80, 0x71,
	0x7a, 0x25, 0xd6, 0xe4, 0x98, 0x59, 0x54, 0xe2, 0x6e, 0x77, 0x97, 0x76, 0xd5, 0xa7, 0xe8, 0x3d,
	0xf7, 0xbe, 0x48, 0x9f, 0xa3, 0xef, 0x52, 0xec, 0x1f, 0x9a, 0x14, 0x9d, 0xb4, 0xcd, 0x8d, 0xf3,
	0xcd, 0xec, 0xec, 0xcc, 0x7c, 0xdf, 0x68, 0x05, 0x83, 0x82, 0xa5, 0x33, 0x21, 0xb9, 0xe6, 0xa4,
	0xab, 0x04, 0xa6, 0xe3, 0xa3, 0x9c, 0xf3, 0x7c, 0x89, 0xaf, 0x2c, 0x76, 0x53, 0xde, 0xbe, 0xa2,
	0xc5, 0xda, 0x05, 0x8c, 0x8f, 0xdb, 0x2e, 0x5c, 0x09, 0x5d, 0x39, 0xcf, 0xda, 0x4e, 0xcd, 0x56,
	0xa8, 0x34, 0x5d, 0x09, 0x1f, 0x70, 0xda, 0x0e, 0xb8, 0x97, 0x54, 0x08, 0x94, 0xca, 0xf9, 0xa7,
	0xdf, 0xc0, 0xe1, 0x82, 0xa5, 0xef, 0x51, 0xde, 0xb1, 0x14, 0xdf, 0x49, 0x2e, 0x50, 0x6a, 0x86,
	0x8a, 0x10, 0xe8, 0xfe, 0xc6, 0x0b, 0x0c, 0x3b, 0x93, 0xce, 0xf9, 0x20, 0xb6, 0xdf, 0xd3, 0x11,
	0xec, 0xcd, 0x25, 0x52, 0x8d, 0x0b, 0x96, 0xaa, 0xa8, 0x10, 0xa5, 0x9e, 0x12, 0xd8, 0xaf, 0xa1,
	0x1f, 0x4b, 0x6d, 0xb0, 0xbf, 0x3a, 0x30, 0x7a, 0x8d, 0x2a, 0x95, 0xec, 0xa6, 0x8e, 0x24, 0xff,
	0x83, 0x01, 0x2b, 0x94, 0xa6, 0x45, 0x8a, 0x2a, 0xec, 0x4c, 0x82, 0xf3, 0x41, 0x5c, 0x03, 0xe4,
	0x10, 0x7a, 0x4b, 0xb6, 0x62, 0x3a, 0x7c, 0x32, 0xe9, 0x9c, 0xf7, 0x62, 0x67, 0x90, 0x23, 0xd8,
	0x2e, 0x58, 0x9a, 0x14, 0x74, 0x85, 0x61, 0x60, 0x0b, 0xe9, 0x17, 0x2c, 0x5d, 0xd0, 0x15, 0x9a,
	0xfa, 0x0a, 0x96, 0xaa, 0xb0, 0x6b, 0x33, 0xd9, 0x6f, 0xf2, 0x1c, 0xb6, 0xf8, 0xed, 0xad, 0x42,
	0x1d, 0xf6, 0x6c, 0x16, 0x6f, 0x19, 0x5c, 0x69, 0xaa, 0x4b, 0x15, 0x6e, 0xd9, 0x24, 0xde, 0x22,
	0x27, 0x00, 0x77, 0xbf, 0x16, 0xa8, 0x13, 0xbd, 0x16, 0x18, 0xf6, 0xed, 0x99, 0x81, 0x45, 0xae,
	0xd7, 0x02, 0xcd, 0x31, 0x6b, 0xa8, 0x70, 0xdb, 0x5e, 0xe2, 0xad, 0xe9, 0x1f, 0x1d, 0x20, 0xcd,
	0xfe, 0x5c, 0xdb, 0x24, 0x84, 0xfe, 0x0a, 0x95, 0xa2, 0x79, 0x35, 0xb4, 0xca, 0x34, 0x89, 0x68,
	0xaa, 0x19, 0x2f, 0x6c, 0x77, 0x83, 0xd8, 0x5b, 0x64, 0x0a, 0xa6, 0x9d, 0xc4, 0x14, 0x1c, 0x4c,
	0x82, 0xf3, 0xe1, 0xc5, 0x60, 0x66, 0xc4, 0x30, 0x5b, 0x44, 0xf3, 0x78, 0xab, 0x30, 0xac, 0xd8,
	0x11, 0x48, 0xd4, 0x49, 0xca, 0x33, 0x0c, 0xbb, 0xb6, 0xc2, 0xbe, 0x44, 0x3d, 0xe7, 0x19, 0x92,
	0x33, 0x18, 0x6a, 0xae, 0xe9, 0x32, 0x49, 0x79, 0x59, 0x54, 0x3d, 0x83, 0x85, 0xe6, 0x06, 0x99,
	0x7e, 0xea, 0x41, 0xb0, 0x88, 0xe6, 0xe4, 0x12, 0x86, 0xa9, 0x25, 0x29, 0x31, 0xea, 0xb0, 0xd5,
	0x0d, 0x2f, 0xc6, 0x33, 0xa7, 0x8c, 0x59, 0xa5, 0x8c, 0xd9, 0x75, 0x25, 0x9d, 0x18, 0x5c, 0xb8,
	0x01, 0xcc, 0x2d, 0x15, 0x4d, 0x09, 0xcb, 0x7c, 0x07, 0x50, 0x41, 0x51, 0x46, 0x9e, 0x81, 0xa9,
	0xd5, 0xf8, 0x1c, 0x45, 0xbd, 0x82, 0xa5, 0x51, 0xb6, 0xc1, 0x5d, 0x77, 0x93, 0xbb, 0x43, 0xe8,
	0xf1, 0xfb, 0x02, 0xa5, 0x2d, 0x79, 0x10, 0x3b, 0xc3, 0xb0, 0x21, 0x24, 0xbb, 0x33, 0x65, 0x32,
	0xe1, 0x99, 0x1a, 0x78, 0x24, 0x12, 0x86, 0x70, 0xc9, 0x97, 0x15, 0x4d, 0xf6, 0x9b, 0x4c, 0x60,
	0x47, 0x72, 0xae, 0x93, 0x52, 0xa1, 0x34, 0x05, 0x6c, 0xbb, 0xe2, 0x0c, 0xf6, 0x41, 0xa1, 0x8c,
	0x32, 0xf2, 0x12, 0x9e, 0x2a, 0x4c, 0x4b, 0xc9, 0xf4, 0x3a, 0xc9, 0x25, 0x2f, 0x45, 0x38, 0xb0,
	0x31, 0xbb, 0x15, 0xfa, 0xd6, 0x80, 0x64, 0x0c, 0xdb, 0x0a, 0x7f, 0x29, 0xb1, 0x48, 0x31, 0x04,
	0x7b, 0xc1, 0x83, 0xdd, 0x50, 0xcf, 0x70, 0x43, 0x3d, 0x97, 0x30, 0x74, 0x5f, 0x6e, 0xaa, 0x3b,
	0xff, 0x3e, 0x55, 0x17, 0x6e, 0xa7, 0x7a, 0x02, 0x5d, 0x4d, 0x73, 0x15, 0xee, 0x36, 0x79, 0xbf,
	0xa6, 0x79, 0x6c, 0x61, 0x33, 0x3c, 0xa7, 0x4c, 0x96, 0x85, 0x4f, 0xdd, 0xf0, 0xac, 0x1d, 0x65,
	0xe4, 0x18, 0x02, 0x64, 0x22, 0xdc, 0xb3, 0xd7, 0xf9, 0x83, 0x6f, 0x98, 0x88, 0x0d, 0x4a, 0xbe,
	0x85, 0x91, 0x96, 0xb4, 0x50, 0xcc, 0xe8, 0x2b, 0xf1, 0x65, 0xef, 0xdb, 0x04, 0xfb, 0xb5, 0xe3,
	0xbd, 0x6b, 0xe0, 0x14, 0x20, 0xe5, 0x85, 0x96, 0x7c, 0xb9, 0x44, 0x19, 0x8e, 0xdc, 0xec, 0x6a,
	0xa4, 0xb5, 0x1e, 0xa4, 0xbd, 0x1e, 0x27, 0xf6, 0xb8, 0xe2, 0x4b, 0xab, 0x8b, 0x03, 0xc7, 0x97,
	0x47, 0xa2, 0xcc, 0xe8, 0x46, 0xa2, 0xe2, 0xa5, 0x74, 0xba, 0x39, 0xf4, 0xd4, 0x78, 0x28, 0xca,
	0xa6, 0x1f, 0x20, 0x78, 0xc3, 0x84, 0x91, 0x0f, 0x32, 0x61, 0x42, 0xdc, 0xd6, 0xf4, 0x90, 0x09,
	0x27, 0x1f, 0x03, 0xd3, 0x2c, 0x93, 0x5e, 0x73, 0x7d, 0x64, 0xe2, 0x2a, 0xcb, 0xa4, 0xf9, 0x25,
	0xb9, 0xa1, 0x45, 0x76, 0xcf, 0x32, 0xfd, 0xd1, 0x6a, 0xae, 0x17, 0xd7, 0xc0, 0xf4, 0xf7, 0x00,
	0x82, 0x6b, 0x9a, 0x1b, 0x91, 0xa5, 0x7c, 0xc9, 0x65, 0x95, 0xd6, 0x1a, 0xed, 0x55, 0x78, 0xf2,
	0x55, 0xab, 0x30, 0x81, 0x61, 0x66, 0xf7, 0x5e, 0xd8, 0x65, 0x76, 0x72, 0x6f, 0x42, 0xb5, 0xb2,
	0xbb, 0x4d, 0x65, 0xbf, 0x84, 0xa7, 0x0f, 0xa3, 0x68, 0xee, 0xea, 0x6e, 0x85, 0xda, 0x75, 0x25,
	0x73, 0x20, 0x0f, 0x61, 0x9a, 0xe6, 0x89, 0xa0, 0x4c, 0x9a, 0x9f, 0x2c, 0xa3, 0x90, 0x67, 0x8e,
	0xe8, 0xd8, 0xfb, 0xaf, 0x69, 0xfe, 0x8e, 0x32, 0x19, 0xef, 0xcb, 0x4d, 0x40, 0x91, 0xb7, 0x70,
	0x50, 0x27, 0x59, 0x8b, 0xea, 0xc2, 0xbe, 0xcd, 0xf2, 0xa2, 0x95, 0x65, 0x2d, 0xdc, 0xd5, 0xf1,
	0x48, 0xb6, 0x21, 0xc3, 0x8b, 0x29, 0xe2, 0x61, 0xab, 0x7a, 0x9a, 0xe6, 0x51, 0x46, 0x5e, 0x40,
	0xdf, 0xc0, 0x3f, 0xe3, 0xda, 0x6f, 0x92, 0x89, 0xfa, 0x1e, 0xd7, 0x86, 0x30, 0xe3, 0xb0, 0xfb,
	0x0e, 0x8e, 0x30, 0x4d, 0x73, 0xb3, 0xef, 0xd3, 0x3f, 0x3b, 0xb0, 0xd7, 0xaa, 0xbc, 0x2d, 0x8f,
	0x4e, 0x5b, 0x1e, 0xe4, 0xff, 0xb0, 0xbb, 0xd1, 0x88, 0x57, 0xc1, 0x4e, 0xb3, 0xd2, 0xc6, 0x6e,
	0x06, 0xff, 0xb4, 0x9b, 0xdd, 0xaf, 0xda, 0xcd, 0xba, 0xf3, 0x5e, 0xa3, 0xf3, 0xe9, 0x02, 0x46,
	0x8f, 0x06, 0xe7, 0x54, 0x66, 0x06, 0xdc, 0x71, 0xef, 0x96, 0x35, 0xfe, 0x53, 0xed, 0xe6, 0x35,
	0xbd, 0xd2, 0x9a, 0xa6, 0x1f, 0x37, 0x5e, 0xd3, 0x1a, 0xf2, 0xaf, 0xe9, 0x08, 0xf6, 0x5e, 0xe3,
	0xa3, 0xb0, 0x1a, 0xf2, 0x61, 0x63, 0x08, 0x7f, 0xe0, 0x19, 0xbb, 0x5d, 0x2f, 0x58, 0x7a, 0xa5,
	0xb5, 0x64, 0x37, 0xa5, 0x46, 0x1f, 0x7f, 0x0c, 0x47, 0x9f, 0xf1, 0x35, 0xf3, 0x2f, 0xb1, 0xf5,
	0xa8, 0xd7, 0x90, 0x0b, 0xbb, 0xf8, 0x14, 0x00, 0xd4, 0x7f, 0x14, 0xc8, 0x25, 0x40, 0xfd, 0xee,
	0x13, 0xaf, 0xce, 0xd6, 0x9f, 0x83, 0xf1, 0xf3, 0x36, 0xec, 0x5f, 0xca, 0x2b, 0xd8, 0x69, 0xbe,
	0x9f, 0xc4, 0xcb, 0xf2, 0xd1, 0x7f, 0x86, 0x71, 0xf8, 0xd8, 0xe1, 0x53, 0x5c, 0x02, 0xd4, 0x93,
	0xaa, 0xee, 0x6f, 0x8d, 0xb3, 0xba, 0xbf, 0x3d, 0x52, 0x73, 0xb8, 0x9e, 0x5f, 0x75, 0xb8, 0x35,
	0xe4, 0xea, 0x70, 0x7b, 0xd0, 0xe4, 0x27, 0x38, 0xf8, 0xcc, 0x30, 0xc9, 0xa9, 0x0b, 0xff, 0x12,
	0x07, 0xe3, 0xb3, 0x2f, 0xfa, 0x9b, 0x45, 0x55, 0x43, 0xaf, 0x8b, 0xda, 0x60, 0xa6, 0x2e, 0x6a,
	0x93, 0x9d, 0x9b, 0x2d, 0x2b, 0xe9, 0xef, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x45, 0x39, 0xc7,
	0x79, 0x58, 0x0a, 0x00, 0x00,
}
