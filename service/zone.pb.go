// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zone.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

import "regexp"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/logger"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.SetLevel
var _ = request.Request{}
var _ = data.Operation{}

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
func (*ZoneServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *ZoneServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeZonesInput struct {
	Zones  []string `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	Status []string `protobuf:"bytes,2,rep,name=status" json:"status,omitempty"`
}

func (m *DescribeZonesInput) Reset()                    { *m = DescribeZonesInput{} }
func (m *DescribeZonesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeZonesInput) ProtoMessage()               {}
func (*DescribeZonesInput) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{1} }

func (m *DescribeZonesInput) GetZones() []string {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *DescribeZonesInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type DescribeZonesOutput struct {
	Action     string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ZoneSet    []*Zone `protobuf:"bytes,4,rep,name=zone_set,json=zoneSet" json:"zone_set,omitempty"`
	TotalCount int32   `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeZonesOutput) Reset()                    { *m = DescribeZonesOutput{} }
func (m *DescribeZonesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeZonesOutput) ProtoMessage()               {}
func (*DescribeZonesOutput) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{2} }

func (m *DescribeZonesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeZonesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeZonesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeZonesOutput) GetZoneSet() []*Zone {
	if m != nil {
		return m.ZoneSet
	}
	return nil
}

func (m *DescribeZonesOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ZoneServiceProperties)(nil), "service.ZoneServiceProperties")
	proto.RegisterType((*DescribeZonesInput)(nil), "service.DescribeZonesInput")
	proto.RegisterType((*DescribeZonesOutput)(nil), "service.DescribeZonesOutput")
}

// See https://docs.qingcloud.com/api/zone/index.html
type ZoneServiceInterface interface {
	DescribeZones(in *DescribeZonesInput) (out *DescribeZonesOutput, err error)
}

// See https://docs.qingcloud.com/api/zone/index.html
type ZoneService struct {
	Config           *config.Config
	Properties       *ZoneServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/zone/index.html
func NewZoneService(conf *config.Config, zone string) (p *ZoneService) {
	return &ZoneService{
		Config:     conf,
		Properties: &ZoneServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/zone/index.html
func (s *QingCloudService) Zone(zone string) (*ZoneService, error) {
	properties := &ZoneServiceProperties{
		Zone: zone,
	}

	return &ZoneService{Config: s.Config, Properties: properties}, nil
}

func (p *ZoneService) DescribeZones(in *DescribeZonesInput) (out *DescribeZonesOutput, err error) {
	if in == nil {
		in = &DescribeZonesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeZones",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeZonesOutput{}
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

func (p *ZoneServiceProperties) Validate() error {
	return nil
}

func (p *DescribeZonesInput) Validate() error {
	return nil
}

func (p *DescribeZonesOutput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("zone.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcb, 0xae, 0xd3, 0x30,
	0x10, 0x55, 0xfa, 0x4a, 0x3b, 0x51, 0x59, 0x18, 0xa8, 0x4c, 0x00, 0x11, 0x75, 0x15, 0x09, 0x29,
	0x91, 0xca, 0x8e, 0x15, 0xa2, 0xdd, 0xc0, 0x06, 0x08, 0x3b, 0x36, 0x51, 0xea, 0x8c, 0x5a, 0x8b,
	0xd4, 0x0e, 0xf6, 0x04, 0x41, 0x3f, 0x81, 0x9f, 0xe0, 0x0b, 0xf8, 0x38, 0xfe, 0xe0, 0xca, 0x4e,
	0x6e, 0x75, 0xab, 0x7b, 0x37, 0x51, 0xce, 0x19, 0xcf, 0x99, 0x33, 0x67, 0x00, 0xce, 0x5a, 0x61,
	0xd6, 0x1a, 0x4d, 0x9a, 0x85, 0x16, 0xcd, 0x4f, 0x29, 0x30, 0x8e, 0xe8, 0x77, 0x8b, 0xb6, 0x67,
	0xe3, 0x97, 0x3f, 0xa4, 0x3a, 0x88, 0x46, 0x77, 0x75, 0x69, 0xeb, 0xef, 0xa5, 0xe9, 0x1a, 0xcc,
	0xdd, 0xa7, 0x2f, 0xaf, 0x5f, 0xc3, 0xd3, 0x6f, 0x5a, 0xe1, 0xd7, 0xbe, 0xf5, 0xb3, 0xd1, 0x2d,
	0x1a, 0x92, 0x68, 0x19, 0x83, 0x89, 0xd3, 0xe6, 0x41, 0x12, 0xa4, 0x8b, 0xc2, 0xff, 0xaf, 0xdf,
	0x03, 0xdb, 0xa1, 0x15, 0x46, 0xee, 0xd1, 0x35, 0xd9, 0x0f, 0xaa, 0xed, 0x88, 0x3d, 0x81, 0xa9,
	0xab, 0x5a, 0x1e, 0x24, 0xe3, 0x74, 0x51, 0xf4, 0x80, 0xad, 0x60, 0x66, 0xa9, 0xa2, 0xce, 0xf2,
	0x91, 0xa7, 0x07, 0xb4, 0xfe, 0x17, 0xc0, 0xe3, 0x2b, 0x91, 0x4f, 0x1d, 0x39, 0x95, 0x15, 0xcc,
	0x2a, 0x41, 0x52, 0xab, 0x61, 0xe2, 0x80, 0xd8, 0x33, 0x98, 0x1b, 0xa4, 0x52, 0xe8, 0x1a, 0xf9,
	0x28, 0x09, 0xd2, 0x69, 0x11, 0x1a, 0xa4, 0xad, 0xae, 0x91, 0x71, 0x08, 0x4f, 0x68, 0x6d, 0x75,
	0x40, 0x3e, 0xf6, 0x3d, 0xb7, 0x90, 0xa5, 0x30, 0x77, 0x2e, 0x4a, 0x8b, 0xc4, 0x27, 0xc9, 0x38,
	0x8d, 0x36, 0xcb, 0x6c, 0x48, 0x27, 0x73, 0x43, 0x8b, 0xf0, 0xec, 0x97, 0x26, 0xf6, 0x0a, 0x22,
	0xd2, 0x54, 0x35, 0xa5, 0xd0, 0x9d, 0x22, 0x3e, 0xf5, 0x13, 0xc0, 0x53, 0x5b, 0xc7, 0x6c, 0xfe,
	0x06, 0x10, 0xdd, 0x49, 0x88, 0x7d, 0x84, 0xe5, 0x95, 0x7d, 0xf6, 0xfc, 0xa2, 0x7c, 0x3f, 0x9b,
	0xf8, 0xc5, 0xc3, 0xc5, 0x7e, 0xe7, 0x78, 0xf7, 0xe7, 0xff, 0xe4, 0x1d, 0x64, 0x47, 0xa2, 0xd6,
	0xbe, 0xcd, 0xf3, 0x5a, 0x0b, 0x9b, 0x5d, 0xce, 0x95, 0x09, 0x7d, 0xca, 0xab, 0x56, 0xe6, 0xce,
	0x6a, 0x2e, 0x55, 0x8d, 0xbf, 0xb2, 0x23, 0x9d, 0x1a, 0xf6, 0xe8, 0x8b, 0x54, 0x87, 0xad, 0x7f,
	0xe2, 0xc4, 0xf6, 0x33, 0x7f, 0xc9, 0x37, 0x37, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xd5, 0x42,
	0xca, 0x0c, 0x02, 0x00, 0x00,
}
