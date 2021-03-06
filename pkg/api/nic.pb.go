// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nic.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NicServiceProperties struct {
	Zone             *string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NicServiceProperties) Reset()                    { *m = NicServiceProperties{} }
func (m *NicServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*NicServiceProperties) ProtoMessage()               {}
func (*NicServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *NicServiceProperties) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

type CreateNicsInput struct {
	Vxnet            *string  `protobuf:"bytes,1,opt,name=vxnet" json:"vxnet,omitempty"`
	NicName          *string  `protobuf:"bytes,2,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Count            *int32   `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	PrivateIps       []string `protobuf:"bytes,4,rep,name=private_ips,json=privateIps" json:"private_ips,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CreateNicsInput) Reset()                    { *m = CreateNicsInput{} }
func (m *CreateNicsInput) String() string            { return proto.CompactTextString(m) }
func (*CreateNicsInput) ProtoMessage()               {}
func (*CreateNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *CreateNicsInput) GetVxnet() string {
	if m != nil && m.Vxnet != nil {
		return *m.Vxnet
	}
	return ""
}

func (m *CreateNicsInput) GetNicName() string {
	if m != nil && m.NicName != nil {
		return *m.NicName
	}
	return ""
}

func (m *CreateNicsInput) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CreateNicsInput) GetPrivateIps() []string {
	if m != nil {
		return m.PrivateIps
	}
	return nil
}

type CreateNicsOutput struct {
	Action           *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Nics             []*NICIP `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CreateNicsOutput) Reset()                    { *m = CreateNicsOutput{} }
func (m *CreateNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateNicsOutput) ProtoMessage()               {}
func (*CreateNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *CreateNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *CreateNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *CreateNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *CreateNicsOutput) GetNics() []*NICIP {
	if m != nil {
		return m.Nics
	}
	return nil
}

type DescribeNicsInput struct {
	Instances        []string `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	Limit            *int32   `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	NicName          *string  `protobuf:"bytes,3,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Nics             []string `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
	Offset           *int32   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Status           *string  `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	VxnetType        *int32   `protobuf:"varint,7,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	Vxnets           []string `protobuf:"bytes,8,rep,name=vxnets" json:"vxnets,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DescribeNicsInput) Reset()                    { *m = DescribeNicsInput{} }
func (m *DescribeNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeNicsInput) ProtoMessage()               {}
func (*DescribeNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *DescribeNicsInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *DescribeNicsInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *DescribeNicsInput) GetNicName() string {
	if m != nil && m.NicName != nil {
		return *m.NicName
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
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeNicsInput) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *DescribeNicsInput) GetVxnetType() int32 {
	if m != nil && m.VxnetType != nil {
		return *m.VxnetType
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
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	NicSet           []*NIC  `protobuf:"bytes,4,rep,name=nic_set,json=nicSet" json:"nic_set,omitempty"`
	TotalCount       *int32  `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DescribeNicsOutput) Reset()                    { *m = DescribeNicsOutput{} }
func (m *DescribeNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeNicsOutput) ProtoMessage()               {}
func (*DescribeNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *DescribeNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeNicsOutput) GetNicSet() []*NIC {
	if m != nil {
		return m.NicSet
	}
	return nil
}

func (m *DescribeNicsOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type AttachNicsInput struct {
	Nics             []string `protobuf:"bytes,1,rep,name=nics" json:"nics,omitempty"`
	Instance         *string  `protobuf:"bytes,2,opt,name=instance" json:"instance,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AttachNicsInput) Reset()                    { *m = AttachNicsInput{} }
func (m *AttachNicsInput) String() string            { return proto.CompactTextString(m) }
func (*AttachNicsInput) ProtoMessage()               {}
func (*AttachNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *AttachNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *AttachNicsInput) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

type AttachNicsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AttachNicsOutput) Reset()                    { *m = AttachNicsOutput{} }
func (m *AttachNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachNicsOutput) ProtoMessage()               {}
func (*AttachNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *AttachNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AttachNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AttachNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AttachNicsOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type DetachNicsInput struct {
	Nics             []string `protobuf:"bytes,1,rep,name=nics" json:"nics,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DetachNicsInput) Reset()                    { *m = DetachNicsInput{} }
func (m *DetachNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DetachNicsInput) ProtoMessage()               {}
func (*DetachNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *DetachNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

type DetachNicsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DetachNicsOutput) Reset()                    { *m = DetachNicsOutput{} }
func (m *DetachNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachNicsOutput) ProtoMessage()               {}
func (*DetachNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *DetachNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DetachNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DetachNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DetachNicsOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type ModifyNicAttributesInput struct {
	Nic              *string `protobuf:"bytes,1,opt,name=nic" json:"nic,omitempty"`
	NicName          *string `protobuf:"bytes,2,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	PrivateIp        *string `protobuf:"bytes,3,opt,name=private_ip,json=privateIp" json:"private_ip,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifyNicAttributesInput) Reset()                    { *m = ModifyNicAttributesInput{} }
func (m *ModifyNicAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyNicAttributesInput) ProtoMessage()               {}
func (*ModifyNicAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *ModifyNicAttributesInput) GetNic() string {
	if m != nil && m.Nic != nil {
		return *m.Nic
	}
	return ""
}

func (m *ModifyNicAttributesInput) GetNicName() string {
	if m != nil && m.NicName != nil {
		return *m.NicName
	}
	return ""
}

func (m *ModifyNicAttributesInput) GetPrivateIp() string {
	if m != nil && m.PrivateIp != nil {
		return *m.PrivateIp
	}
	return ""
}

type ModifyNicAttributesOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifyNicAttributesOutput) Reset()                    { *m = ModifyNicAttributesOutput{} }
func (m *ModifyNicAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyNicAttributesOutput) ProtoMessage()               {}
func (*ModifyNicAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *ModifyNicAttributesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ModifyNicAttributesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ModifyNicAttributesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type DeleteNicsInput struct {
	Nics             []string `protobuf:"bytes,1,rep,name=nics" json:"nics,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeleteNicsInput) Reset()                    { *m = DeleteNicsInput{} }
func (m *DeleteNicsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteNicsInput) ProtoMessage()               {}
func (*DeleteNicsInput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

func (m *DeleteNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

type DeleteNicsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteNicsOutput) Reset()                    { *m = DeleteNicsOutput{} }
func (m *DeleteNicsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteNicsOutput) ProtoMessage()               {}
func (*DeleteNicsOutput) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{12} }

func (m *DeleteNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DeleteNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DeleteNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*NicServiceProperties)(nil), "service.NicServiceProperties")
	proto.RegisterType((*CreateNicsInput)(nil), "service.CreateNicsInput")
	proto.RegisterType((*CreateNicsOutput)(nil), "service.CreateNicsOutput")
	proto.RegisterType((*DescribeNicsInput)(nil), "service.DescribeNicsInput")
	proto.RegisterType((*DescribeNicsOutput)(nil), "service.DescribeNicsOutput")
	proto.RegisterType((*AttachNicsInput)(nil), "service.AttachNicsInput")
	proto.RegisterType((*AttachNicsOutput)(nil), "service.AttachNicsOutput")
	proto.RegisterType((*DetachNicsInput)(nil), "service.DetachNicsInput")
	proto.RegisterType((*DetachNicsOutput)(nil), "service.DetachNicsOutput")
	proto.RegisterType((*ModifyNicAttributesInput)(nil), "service.ModifyNicAttributesInput")
	proto.RegisterType((*ModifyNicAttributesOutput)(nil), "service.ModifyNicAttributesOutput")
	proto.RegisterType((*DeleteNicsInput)(nil), "service.DeleteNicsInput")
	proto.RegisterType((*DeleteNicsOutput)(nil), "service.DeleteNicsOutput")
}

func init() { proto.RegisterFile("nic.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x6f, 0xfe, 0xea, 0xd3, 0xea, 0xb6, 0x77, 0x6e, 0xef, 0x95, 0x6b, 0xa8, 0x28, 0x96,
	0x2a, 0x55, 0x48, 0x24, 0xa2, 0x0b, 0x16, 0xec, 0xa2, 0x54, 0x42, 0x91, 0x20, 0x54, 0x81, 0x25,
	0x92, 0x35, 0x19, 0x4f, 0x9a, 0xa9, 0xe2, 0x19, 0xe3, 0x39, 0x89, 0x1a, 0xb6, 0xbc, 0x03, 0xef,
	0xc0, 0x2b, 0xf0, 0x2e, 0x3c, 0x07, 0x62, 0x87, 0xc6, 0x33, 0x89, 0x9d, 0x36, 0x94, 0x55, 0xd8,
	0xf9, 0x3b, 0xf3, 0xf3, 0x7d, 0xe7, 0xf3, 0x77, 0x06, 0x7c, 0x29, 0x58, 0x3b, 0xcb, 0x15, 0x2a,
	0xd2, 0xd2, 0x3c, 0x9f, 0x0b, 0xc6, 0xc3, 0x5d, 0x5c, 0x64, 0x5c, 0xdb, 0x6a, 0x18, 0xe8, 0x8c,
	0xb3, 0x38, 0xe5, 0x48, 0x13, 0x8a, 0xb4, 0x63, 0x90, 0x5d, 0x89, 0x9e, 0xc0, 0xe1, 0x40, 0xb0,
	0xb7, 0xf6, 0xd0, 0x65, 0xae, 0x32, 0x9e, 0xa3, 0xe0, 0x9a, 0x10, 0xa8, 0x7f, 0x54, 0x92, 0x07,
	0xde, 0x89, 0x77, 0xe6, 0x0f, 0x8b, 0xef, 0x68, 0x01, 0xfb, 0xbd, 0x9c, 0x53, 0xe4, 0x03, 0xc1,
	0x74, 0x5f, 0x66, 0x33, 0x24, 0x87, 0xd0, 0x98, 0xdf, 0x48, 0x8e, 0x6e, 0x9f, 0x05, 0xe4, 0x08,
	0x76, 0xa4, 0x60, 0xb1, 0xa4, 0x29, 0x0f, 0xfe, 0x2a, 0x16, 0x5a, 0x52, 0xb0, 0x01, 0x4d, 0xb9,
	0x39, 0xc0, 0xd4, 0x4c, 0x62, 0x50, 0x3b, 0xf1, 0xce, 0x1a, 0x43, 0x0b, 0xc8, 0x23, 0xd8, 0xcd,
	0x72, 0x31, 0xa7, 0xc8, 0x63, 0x91, 0xe9, 0xa0, 0x7e, 0x52, 0x3b, 0xf3, 0x87, 0xe0, 0x4a, 0xfd,
	0x4c, 0x47, 0x9f, 0x3c, 0x38, 0x28, 0xb9, 0xdf, 0xcc, 0xd0, 0x90, 0xff, 0x0f, 0x4d, 0xca, 0x50,
	0x28, 0xe9, 0xd8, 0x1d, 0x32, 0xf4, 0x39, 0xc7, 0x98, 0xa9, 0xc4, 0xd2, 0x37, 0x86, 0xad, 0x9c,
	0x63, 0x4f, 0x25, 0x9c, 0x04, 0xd0, 0x4a, 0xb9, 0xd6, 0xf4, 0x8a, 0x17, 0x02, 0xfc, 0xe1, 0x12,
	0x92, 0x08, 0xea, 0x52, 0x30, 0xcb, 0xbd, 0x7b, 0xfe, 0x77, 0xdb, 0xf9, 0xd8, 0x1e, 0xf4, 0x7b,
	0xfd, 0xcb, 0x61, 0xb1, 0x16, 0x7d, 0xf3, 0xe0, 0x9f, 0x0b, 0xae, 0x59, 0x2e, 0x46, 0x15, 0x0f,
	0x1e, 0x82, 0x2f, 0xa4, 0x46, 0x2a, 0x19, 0xd7, 0x81, 0x57, 0x48, 0x2f, 0x0b, 0xa6, 0xe1, 0xa9,
	0x48, 0x05, 0x3a, 0x25, 0x16, 0xac, 0x39, 0x54, 0x5b, 0x77, 0x88, 0x54, 0x84, 0xf8, 0x96, 0xd8,
	0x74, 0xaa, 0xc6, 0x63, 0xcd, 0x31, 0x68, 0x14, 0xb7, 0x38, 0x64, 0xea, 0x1a, 0x29, 0xce, 0x74,
	0xd0, 0xb4, 0x0e, 0x58, 0x44, 0x8e, 0x01, 0x8a, 0x3f, 0x11, 0x9b, 0x10, 0x04, 0xad, 0xe2, 0x8c,
	0x5f, 0x54, 0xde, 0x2d, 0x32, 0x6e, 0x8e, 0x15, 0x40, 0x07, 0x3b, 0x05, 0x89, 0x43, 0xd1, 0x17,
	0x0f, 0x48, 0xb5, 0xbf, 0x6d, 0xf8, 0x7c, 0x0a, 0xa6, 0xd3, 0xd8, 0xf4, 0x62, 0xad, 0xde, 0xab,
	0x5a, 0x3d, 0x6c, 0x4a, 0x93, 0xc6, 0x22, 0x11, 0xa8, 0x90, 0x4e, 0x63, 0x9b, 0x16, 0xdb, 0x36,
	0x14, 0xa5, 0x9e, 0xa9, 0x44, 0x5d, 0xd8, 0xef, 0x22, 0x52, 0x36, 0x29, 0x7f, 0xc4, 0xd2, 0x39,
	0xaf, 0xe2, 0x5c, 0x08, 0x3b, 0xcb, 0x7f, 0xe1, 0xa2, 0xb8, 0xc2, 0xd1, 0x1c, 0x0e, 0xca, 0x2b,
	0xb6, 0xd1, 0xeb, 0x7f, 0xd0, 0xbc, 0x56, 0xa3, 0x58, 0x24, 0x41, 0xdd, 0x8e, 0xc7, 0xb5, 0x1a,
	0xf5, 0x93, 0xe8, 0x14, 0xf6, 0x2f, 0xf8, 0x6f, 0xa5, 0x1b, 0x79, 0xe5, 0xb6, 0x3f, 0x28, 0x6f,
	0x0c, 0xc1, 0x6b, 0x95, 0x88, 0xf1, 0x62, 0x20, 0x58, 0x17, 0x31, 0x17, 0xa3, 0x19, 0x72, 0xa7,
	0xf3, 0x00, 0x6a, 0x52, 0x30, 0x47, 0x6e, 0x3e, 0xef, 0x9b, 0xf5, 0x63, 0x80, 0x72, 0xaa, 0x1d,
	0xb9, 0xbf, 0x1a, 0xea, 0x68, 0x02, 0x47, 0x1b, 0x78, 0xb6, 0xd0, 0xa8, 0x35, 0x7c, 0xca, 0xab,
	0x0f, 0xd7, 0x26, 0xc3, 0x63, 0x63, 0xf8, 0x72, 0xdb, 0x16, 0x74, 0x9c, 0x7f, 0xaf, 0x01, 0x94,
	0xaf, 0x2d, 0xe9, 0x02, 0x94, 0x6f, 0x1a, 0x09, 0x56, 0x73, 0x70, 0xeb, 0x91, 0x0d, 0x8f, 0x36,
	0xac, 0x38, 0x79, 0x2f, 0x61, 0xaf, 0x3a, 0xb0, 0x24, 0x5c, 0x6d, 0xbd, 0xf3, 0x4e, 0x85, 0x0f,
	0x36, 0xae, 0xb9, 0x8b, 0xba, 0x00, 0xe5, 0x2c, 0x54, 0xb4, 0xdc, 0x9a, 0xb1, 0x8a, 0x96, 0x3b,
	0xa3, 0xd3, 0x05, 0x28, 0xf3, 0x5a, 0xb9, 0xe2, 0x56, 0xd6, 0x2b, 0x57, 0xdc, 0x89, 0xf7, 0x7b,
	0xf8, 0x77, 0x43, 0x24, 0xc8, 0xe3, 0xd5, 0x89, 0x5f, 0x05, 0x33, 0x8c, 0xee, 0xdb, 0x52, 0x15,
	0xb8, 0xfc, 0xbf, 0x6b, 0x02, 0xd7, 0xb2, 0xb1, 0x26, 0x70, 0x3d, 0x0e, 0xe1, 0xf3, 0xaf, 0x9f,
	0x7f, 0xbc, 0x7a, 0x06, 0x9d, 0xf0, 0xe9, 0x04, 0x31, 0xd3, 0x2f, 0x3a, 0x9d, 0x44, 0x31, 0xdd,
	0xfe, 0x20, 0xe4, 0x15, 0x9b, 0xaa, 0x59, 0xd2, 0x66, 0x2a, 0xed, 0xd0, 0x4c, 0x74, 0xa4, 0x60,
	0x1d, 0x21, 0x13, 0x7e, 0xd3, 0x9e, 0x60, 0x3a, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x4f,
	0xb4, 0x4a, 0xa2, 0x07, 0x00, 0x00,
}
