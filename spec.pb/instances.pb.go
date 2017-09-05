// Code generated by protoc-gen-go. DO NOT EDIT.
// source: instances.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "context"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DescribeInstanceTypes_Request struct {
	Action        string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Zone          string   `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	InstanceTypes []string `protobuf:"bytes,3,rep,name=instance_types,json=instanceTypes" json:"instance_types,omitempty"`
}

func (m *DescribeInstanceTypes_Request) Reset()                    { *m = DescribeInstanceTypes_Request{} }
func (m *DescribeInstanceTypes_Request) String() string            { return proto.CompactTextString(m) }
func (*DescribeInstanceTypes_Request) ProtoMessage()               {}
func (*DescribeInstanceTypes_Request) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DescribeInstanceTypes_Request) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeInstanceTypes_Request) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *DescribeInstanceTypes_Request) GetInstanceTypes() []string {
	if m != nil {
		return m.InstanceTypes
	}
	return nil
}

type DescribeInstanceTypes_Reponse struct {
	Action          string                 `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode         int32                  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	InstanceTypeSet []*InstanceTypeSetElem `protobuf:"bytes,3,rep,name=instance_type_set,json=instanceTypeSet" json:"instance_type_set,omitempty"`
	TotalCount      int32                  `protobuf:"varint,4,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeInstanceTypes_Reponse) Reset()                    { *m = DescribeInstanceTypes_Reponse{} }
func (m *DescribeInstanceTypes_Reponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeInstanceTypes_Reponse) ProtoMessage()               {}
func (*DescribeInstanceTypes_Reponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *DescribeInstanceTypes_Reponse) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeInstanceTypes_Reponse) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeInstanceTypes_Reponse) GetInstanceTypeSet() []*InstanceTypeSetElem {
	if m != nil {
		return m.InstanceTypeSet
	}
	return nil
}

func (m *DescribeInstanceTypes_Reponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type InstanceTypeSetElem struct {
	InstanceTypeId   string `protobuf:"bytes,1,opt,name=instance_type_id,json=instanceTypeId" json:"instance_type_id,omitempty"`
	InstanceTypeName string `protobuf:"bytes,2,opt,name=instance_type_name,json=instanceTypeName" json:"instance_type_name,omitempty"`
	Description      string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	VcpusCurrent     int32  `protobuf:"varint,4,opt,name=vcpus_current,json=vcpusCurrent" json:"vcpus_current,omitempty"`
	MemoryCurrent    int32  `protobuf:"varint,5,opt,name=memory_current,json=memoryCurrent" json:"memory_current,omitempty"`
	Status           string `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *InstanceTypeSetElem) Reset()                    { *m = InstanceTypeSetElem{} }
func (m *InstanceTypeSetElem) String() string            { return proto.CompactTextString(m) }
func (*InstanceTypeSetElem) ProtoMessage()               {}
func (*InstanceTypeSetElem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *InstanceTypeSetElem) GetInstanceTypeId() string {
	if m != nil {
		return m.InstanceTypeId
	}
	return ""
}

func (m *InstanceTypeSetElem) GetInstanceTypeName() string {
	if m != nil {
		return m.InstanceTypeName
	}
	return ""
}

func (m *InstanceTypeSetElem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InstanceTypeSetElem) GetVcpusCurrent() int32 {
	if m != nil {
		return m.VcpusCurrent
	}
	return 0
}

func (m *InstanceTypeSetElem) GetMemoryCurrent() int32 {
	if m != nil {
		return m.MemoryCurrent
	}
	return 0
}

func (m *InstanceTypeSetElem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*DescribeInstanceTypes_Request)(nil), "spec.DescribeInstanceTypes_Request")
	proto.RegisterType((*DescribeInstanceTypes_Reponse)(nil), "spec.DescribeInstanceTypes_Reponse")
	proto.RegisterType((*InstanceTypeSetElem)(nil), "spec.InstanceTypeSetElem")
}

type InstancesService interface {
	DescribeInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RunInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	TerminateInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RestartInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResetInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResizeInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyInstanceAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeInstanceTypes(ctx context.Context, in *DescribeInstanceTypes_Request) (out *DescribeInstanceTypes_Reponse, err error)
	CreateBrokers(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteBrokers(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type InstancesServiceClient struct{}

// NewInstancesServiceClient returns a InstancesService stub to handle
// requests to the set of InstancesService at the other end of the connection.
func NewInstancesServiceClient(opt *Options) *InstancesServiceClient {
	return &InstancesServiceClient{}
}

func (c *InstancesServiceClient) DescribeInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) RunInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) TerminateInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) StartInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) StopInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) RestartInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) ResetInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) ResizeInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) ModifyInstanceAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) DescribeInstanceTypes(ctx context.Context, in *DescribeInstanceTypes_Request, opt ...*Options) (out *DescribeInstanceTypes_Reponse, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) CreateBrokers(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) DeleteBrokers(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("instances.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x15, 0x83, 0xb2, 0xf5, 0x52, 0x3e, 0xea, 0x69, 0x28, 0x30, 0x6d, 0x45, 0x54, 0x93, 0x78,
	0x98, 0x52, 0xa9, 0x7b, 0x5f, 0x3f, 0x80, 0x87, 0x4e, 0xda, 0x1e, 0x42, 0xdf, 0x23, 0x93, 0xdc,
	0x56, 0xd6, 0x88, 0xed, 0xd9, 0x37, 0x9d, 0xe8, 0x4f, 0xdb, 0x4f, 0xd9, 0xdf, 0xd8, 0x1f, 0x98,
	0xe2, 0x90, 0x0a, 0x50, 0xd9, 0xa4, 0xf0, 0x16, 0x1f, 0x9f, 0x9c, 0x73, 0x7d, 0x7c, 0x7d, 0xa1,
	0x2d, 0xa4, 0x25, 0x2e, 0x23, 0xb4, 0xbe, 0x36, 0x8a, 0x14, 0xab, 0x59, 0x8d, 0x51, 0xbf, 0x77,
	0xaf, 0xd4, 0xfd, 0x02, 0xcf, 0x1c, 0x36, 0x4f, 0xef, 0xce, 0xb8, 0x5c, 0xe6, 0x84, 0xfe, 0xdb,
	0xed, 0x2d, 0x4c, 0x34, 0x15, 0x9b, 0x27, 0xdb, 0x9b, 0x24, 0x12, 0xb4, 0xc4, 0x13, 0xbd, 0x22,
	0xbc, 0xdf, 0x26, 0xfc, 0x34, 0x5c, 0x6b, 0x34, 0x2b, 0xfb, 0xa1, 0x81, 0x77, 0x13, 0xb4, 0x91,
	0x11, 0x73, 0xbc, 0x59, 0x55, 0x76, 0xbb, 0xd4, 0x68, 0xc3, 0x00, 0x7f, 0xa4, 0x68, 0x89, 0x75,
	0xa1, 0xce, 0x23, 0x12, 0x4a, 0x7a, 0x95, 0x41, 0x65, 0x74, 0x18, 0xac, 0x56, 0x8c, 0x41, 0xed,
	0x51, 0x49, 0xf4, 0x5e, 0x38, 0xd4, 0x7d, 0xb3, 0x0f, 0xd0, 0x2a, 0x8e, 0x17, 0x52, 0xa6, 0xe2,
	0x55, 0x07, 0xd5, 0xd1, 0x61, 0xd0, 0x14, 0xeb, 0xd2, 0xc3, 0x5f, 0x95, 0xdd, 0xa6, 0x5a, 0x49,
	0x8b, 0x3b, 0x4d, 0x7b, 0xf0, 0xca, 0x20, 0x85, 0x91, 0x8a, 0x73, 0xe3, 0x83, 0xe0, 0xa5, 0x41,
	0x1a, 0xab, 0x18, 0xd9, 0x14, 0x8e, 0x37, 0xbc, 0x43, 0x8b, 0xe4, 0xec, 0x1b, 0xe7, 0x3d, 0x3f,
	0xcb, 0xd8, 0x5f, 0xb7, 0x9a, 0x21, 0x4d, 0x17, 0x98, 0x04, 0x6d, 0xb1, 0x09, 0xb2, 0x13, 0x68,
	0x90, 0x22, 0xbe, 0x08, 0x23, 0x95, 0x4a, 0xf2, 0x6a, 0xce, 0x04, 0x1c, 0x34, 0xce, 0x90, 0xe1,
	0x9f, 0x0a, 0xbc, 0x7e, 0x46, 0x89, 0x8d, 0xa0, 0xb3, 0xe9, 0x2f, 0xe2, 0x55, 0xf1, 0xad, 0x75,
	0x8f, 0x9b, 0x98, 0x7d, 0x04, 0xb6, 0xc9, 0x94, 0x3c, 0x29, 0x72, 0xec, 0xac, 0x73, 0xbf, 0xf1,
	0x04, 0xd9, 0x00, 0x1a, 0xb1, 0xcb, 0x4a, 0xbb, 0x3c, 0xaa, 0x8e, 0xb6, 0x0e, 0xb1, 0x53, 0x68,
	0x3e, 0x44, 0x3a, 0xb5, 0x61, 0x94, 0x1a, 0x83, 0x4f, 0x45, 0x1f, 0x39, 0x70, 0x9c, 0x63, 0xd9,
	0xd5, 0x24, 0x98, 0x28, 0xb3, 0x7c, 0x62, 0x1d, 0x38, 0x56, 0x33, 0x47, 0x0b, 0x5a, 0x17, 0xea,
	0x96, 0x38, 0xa5, 0xd6, 0xab, 0xe7, 0xc1, 0xe7, 0xab, 0xf3, 0xdf, 0x75, 0xe8, 0x14, 0xa7, 0xb6,
	0x33, 0x34, 0x0f, 0x22, 0x42, 0x36, 0x86, 0xe3, 0xed, 0x6b, 0xb4, 0xac, 0xeb, 0xe7, 0x1d, 0xe7,
	0x17, 0x1d, 0xe7, 0x4f, 0xb3, 0x7e, 0xed, 0xef, 0xc0, 0xd9, 0x67, 0x38, 0x0a, 0x52, 0x59, 0xfe,
	0xff, 0x09, 0xb0, 0x5b, 0x34, 0x89, 0x90, 0x9c, 0xf6, 0xa8, 0xe2, 0x12, 0x5a, 0x33, 0xe2, 0x86,
	0xca, 0x2b, 0x5c, 0x40, 0x73, 0x46, 0x4a, 0x97, 0x17, 0xb8, 0x86, 0x4e, 0x90, 0x3d, 0xdd, 0x7d,
	0x8a, 0xb8, 0x84, 0x56, 0x80, 0x16, 0xf7, 0x50, 0xb8, 0x82, 0x76, 0x80, 0x56, 0x3c, 0xee, 0x91,
	0xe5, 0x17, 0xf0, 0xbe, 0xaa, 0x58, 0xdc, 0x2d, 0x0b, 0x89, 0x2b, 0x22, 0x23, 0xe6, 0x29, 0x95,
	0xd0, 0xe2, 0xf0, 0xe6, 0xd9, 0x49, 0xc1, 0x4e, 0xf3, 0x37, 0xfd, 0xcf, 0xd9, 0xd5, 0xff, 0x0f,
	0x29, 0x9f, 0x35, 0x17, 0xd0, 0x1c, 0x1b, 0xe4, 0x84, 0xd7, 0x46, 0x7d, 0x47, 0x53, 0xea, 0xe6,
	0x27, 0xb8, 0xc0, 0xd2, 0x02, 0xf3, 0xba, 0x5b, 0x7f, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xeb, 0xd7, 0xa8, 0x1c, 0x06, 0x00, 0x00,
}
