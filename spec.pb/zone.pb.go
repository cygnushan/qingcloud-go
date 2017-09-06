// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zone.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
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

type ZoneServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ZoneServiceProperties) Reset()                    { *m = ZoneServiceProperties{} }
func (m *ZoneServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*ZoneServiceProperties) ProtoMessage()               {}
func (*ZoneServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *ZoneServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*ZoneServiceProperties)(nil), "spec.ZoneServiceProperties")
}

type ZoneServiceInterface interface {
	DescribeZones(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type ZoneService struct {
	Config     *config.Config
	Properties *ZoneServiceProperties
}

func NewZoneService(conf *config.Config, zone string) (p *ZoneService, err error) {
	return &ZoneService{
		Config:     conf,
		Properties: &ZoneServiceProperties{Zone: zone},
	}, nil
}

func (p *ZoneService) DescribeZones(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeZones",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func init() { proto.RegisterFile("zone.proto", fileDescriptor30) }

var fileDescriptor30 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xaa, 0xca, 0xcf, 0x4b,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x96, 0x92, 0x4c, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42, 0x14,
	0x48, 0x49, 0xa3, 0x4b, 0xa5, 0xe6, 0x16, 0x94, 0xc0, 0x24, 0xe5, 0xd1, 0x25, 0x4b, 0x32, 0x73,
	0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0xa0, 0x0a, 0xe4, 0xd0, 0x15, 0x94, 0x17, 0x25, 0x16, 0x14,
	0xa4, 0x16, 0x15, 0x43, 0xe4, 0x95, 0xb4, 0xb9, 0x44, 0xa3, 0xf2, 0xf3, 0x52, 0x83, 0x53, 0x8b,
	0xca, 0x32, 0x93, 0x53, 0x03, 0x8a, 0xf2, 0x0b, 0x52, 0x8b, 0x4a, 0x32, 0x53, 0x8b, 0x85, 0x84,
	0xb8, 0x58, 0x40, 0xae, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0xfc, 0xb8,
	0xb8, 0x91, 0x14, 0x0b, 0xd9, 0x73, 0xf1, 0xba, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x26, 0xa5, 0x82,
	0x84, 0x8b, 0x85, 0xc4, 0xf4, 0x20, 0xb6, 0xe9, 0xc1, 0x6c, 0xd3, 0x73, 0x05, 0xb9, 0x55, 0x0a,
	0x87, 0x78, 0x12, 0x1b, 0x98, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x55, 0x83, 0xf6, 0xe1,
	0x10, 0x01, 0x00, 0x00,
}
