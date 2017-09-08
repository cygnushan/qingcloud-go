// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spark.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type SparkServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SparkServiceProperties) Reset()                    { *m = SparkServiceProperties{} }
func (m *SparkServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SparkServiceProperties) ProtoMessage()               {}
func (*SparkServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *SparkServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*SparkServiceProperties)(nil), "spec.SparkServiceProperties")
}

type SparkServiceInterface interface {
	CreateSpark(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddSparkNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSparkNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type SparkService struct {
	Config     *config.Config
	Properties *SparkServiceProperties
}

func NewSparkService(conf *config.Config, zone string) (p *SparkService, err error) {
	return &SparkService{
		Config:     conf,
		Properties: &SparkServiceProperties{Zone: zone},
	}, nil
}

func (p *SparkService) CreateSpark(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSpark",
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

func (p *SparkService) DescribeSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSparks",
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

func (p *SparkService) AddSparkNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddSparkNodes",
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

func (p *SparkService) DeleteSparkNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSparkNodes",
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

func (p *SparkService) StartSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartSparks",
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

func (p *SparkService) StopSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopSparks",
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

func (p *SparkService) DeleteSparks(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSparks",
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

func init() { proto.RegisterFile("spark.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0x48, 0x2c,
	0xca, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x96, 0x92, 0x4e,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94,
	0x54, 0x42, 0x94, 0x28, 0xe9, 0x70, 0x89, 0x05, 0x83, 0x74, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x06, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x64, 0xa6, 0x16, 0x0b, 0x09, 0x71, 0xb1, 0x54,
	0xe5, 0xe7, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0x17, 0x99, 0xb9,
	0x78, 0x90, 0x95, 0x0b, 0xd9, 0x72, 0x71, 0x3b, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x82, 0x45, 0x85,
	0xc4, 0xf4, 0x20, 0x76, 0xe9, 0xc1, 0xec, 0xd2, 0x73, 0x05, 0xd9, 0x25, 0x85, 0x43, 0x5c, 0xc8,
	0x81, 0x8b, 0xcf, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0x33, 0x09, 0x62, 0x40, 0x31, 0xc9, 0x26, 0xd8,
	0x73, 0xf1, 0x3a, 0xa6, 0xa4, 0x80, 0x35, 0xfb, 0xe5, 0xa7, 0xa4, 0x92, 0x6e, 0x80, 0x13, 0x97,
	0x80, 0x4b, 0x6a, 0x4e, 0x2a, 0xd4, 0x07, 0xe4, 0x99, 0x61, 0xcb, 0xc5, 0x1d, 0x5c, 0x92, 0x58,
	0x54, 0x42, 0xa6, 0x1f, 0x6c, 0xb8, 0xb8, 0x82, 0x4b, 0xf2, 0x0b, 0xc8, 0xd4, 0x6d, 0xc7, 0xc5,
	0x83, 0xe4, 0x01, 0x92, 0xf5, 0x27, 0xb1, 0x81, 0xf9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x58, 0xa4, 0xc1, 0x57, 0x3a, 0x02, 0x00, 0x00,
}