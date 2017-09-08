// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hadoop.proto

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

type HadoopServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *HadoopServiceProperties) Reset()                    { *m = HadoopServiceProperties{} }
func (m *HadoopServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*HadoopServiceProperties) ProtoMessage()               {}
func (*HadoopServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *HadoopServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*HadoopServiceProperties)(nil), "spec.HadoopServiceProperties")
}

type HadoopServiceInterface interface {
	AddHadoopNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteHadoopNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartHadoops(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopHadoops(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type HadoopService struct {
	Config     *config.Config
	Properties *HadoopServiceProperties
}

func NewHadoopService(conf *config.Config, zone string) (p *HadoopService, err error) {
	return &HadoopService{
		Config:     conf,
		Properties: &HadoopServiceProperties{Zone: zone},
	}, nil
}

func (p *HadoopService) AddHadoopNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddHadoopNodes",
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
func (p *HadoopService) DeleteHadoopNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteHadoopNodes",
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
func (p *HadoopService) StartHadoops(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartHadoops",
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
func (p *HadoopService) StopHadoops(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopHadoops",
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

func init() { proto.RegisterFile("hadoop.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4c, 0xc9,
	0xcf, 0x2f, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x96, 0x92,
	0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16,
	0x94, 0x54, 0x42, 0x94, 0x28, 0xe9, 0x72, 0x89, 0x7b, 0x80, 0xb5, 0x04, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0x06, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x64, 0xa6, 0x16, 0x0b, 0x09, 0x71, 0xb1,
	0x54, 0xe5, 0xe7, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0x53, 0x99,
	0xb8, 0x78, 0x51, 0xd4, 0x0b, 0x39, 0x70, 0xf1, 0x39, 0xa6, 0xa4, 0x40, 0xc4, 0xfc, 0xf2, 0x53,
	0x52, 0x8b, 0x85, 0xc4, 0xf4, 0x20, 0x16, 0xea, 0xc1, 0x2c, 0xd4, 0x73, 0x05, 0x59, 0x28, 0x85,
	0x43, 0x5c, 0xc8, 0x99, 0x4b, 0xd0, 0x25, 0x35, 0x27, 0xb5, 0x24, 0x95, 0x12, 0x43, 0xec, 0xb8,
	0x78, 0x82, 0x4b, 0x12, 0x8b, 0x4a, 0x20, 0x66, 0x90, 0xae, 0xdf, 0x96, 0x8b, 0x3b, 0xb8, 0x24,
	0xbf, 0x80, 0x4c, 0xed, 0x49, 0x6c, 0x60, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x83, 0x10,
	0x42, 0x78, 0x80, 0x01, 0x00, 0x00,
}
