// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vxnet.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/errors"

var _ = config.Config{}
var _ = request.Request{}
var _ = errors.ParameterRequiredError{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VxnetServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *VxnetServiceProperties) Reset()                    { *m = VxnetServiceProperties{} }
func (m *VxnetServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*VxnetServiceProperties) ProtoMessage()               {}
func (*VxnetServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{0} }

func (m *VxnetServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*VxnetServiceProperties)(nil), "spec.VxnetServiceProperties")
}

type VxnetServiceInterface interface {
	DescribeVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	JoinVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	LeaveVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyVxnetAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeVxnetInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type VxnetService struct {
	Config     *config.Config
	Properties *VxnetServiceProperties
}

func NewVxnetService(conf *config.Config, zone string) (p *VxnetService, err error) {
	return &VxnetService{
		Config:     conf,
		Properties: &VxnetServiceProperties{Zone: zone},
	}, nil
}

func (p *VxnetService) DescribeVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVxnets",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *VxnetService) CreateVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVxnets",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *VxnetService) DeleteVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteVxnets",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *VxnetService) JoinVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "JoinVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *VxnetService) LeaveVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "LeaveVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *VxnetService) ModifyVxnetAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyVxnetAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *VxnetService) DescribeVxnetInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVxnetInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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

func init() { proto.RegisterFile("vxnet.proto", fileDescriptor29) }

var fileDescriptor29 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xab, 0xc8, 0x4b,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x96, 0x92, 0x4e,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94,
	0x54, 0x42, 0x94, 0x28, 0xe9, 0x70, 0x89, 0x85, 0x81, 0x74, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x06, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x64, 0xa6, 0x16, 0x0b, 0x09, 0x71, 0xb1, 0x54,
	0xe5, 0xe7, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0xf7, 0x98, 0xb9,
	0x78, 0x90, 0x95, 0x0b, 0x39, 0x70, 0xf1, 0xb9, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x26, 0xa5, 0x82,
	0xc5, 0x8b, 0x85, 0xc4, 0xf4, 0x20, 0xd6, 0xe9, 0xc1, 0xac, 0xd3, 0x73, 0x05, 0x59, 0x27, 0x85,
	0x43, 0x5c, 0xc8, 0x8e, 0x8b, 0xc7, 0xb9, 0x28, 0x35, 0xb1, 0x84, 0x02, 0xfd, 0x2e, 0xa9, 0x39,
	0xa9, 0x64, 0xeb, 0xb7, 0xe6, 0xe2, 0xf4, 0xca, 0xcf, 0xcc, 0x03, 0xeb, 0x26, 0x59, 0xb3, 0x0d,
	0x17, 0x97, 0x4f, 0x6a, 0x62, 0x59, 0x2a, 0x79, 0xba, 0xdd, 0xb9, 0x44, 0x7d, 0xf3, 0x53, 0x32,
	0xd3, 0x2a, 0xc1, 0xda, 0x1d, 0x4b, 0x4a, 0x8a, 0x32, 0x93, 0x4a, 0x4b, 0x52, 0x49, 0xf7, 0x83,
	0x07, 0x97, 0x18, 0x4a, 0x2c, 0x78, 0xe6, 0x15, 0x97, 0x24, 0xe6, 0x25, 0x93, 0x6e, 0x52, 0x12,
	0x1b, 0x98, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x50, 0x33, 0x5d, 0x18, 0x47, 0x02, 0x00,
	0x00,
}
