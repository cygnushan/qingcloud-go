// Code generated by protoc-gen-go. DO NOT EDIT.
// source: key_pair.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type KeyPairServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *KeyPairServiceProperties) Reset()                    { *m = KeyPairServiceProperties{} }
func (m *KeyPairServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*KeyPairServiceProperties) ProtoMessage()               {}
func (*KeyPairServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *KeyPairServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyPairServiceProperties)(nil), "service.KeyPairServiceProperties")
}

type KeyPairServiceInterface interface {
	DescribeKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateKeyPair(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AttachKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DetachKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyKeyPairAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type KeyPairService struct {
	Config     *config.Config
	Properties *KeyPairServiceProperties
}

func NewKeyPairService(conf *config.Config, zone string) (p *KeyPairService) {
	return &KeyPairService{
		Config:     conf,
		Properties: &KeyPairServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) KeyPair(zone string) (*KeyPairService, error) {
	properties := &KeyPairServiceProperties{
		Zone: zone,
	}

	return &KeyPairService{Config: s.Config, Properties: properties}, nil
}

func (p *KeyPairService) DescribeKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeKeyPairs",
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

func (p *KeyPairService) CreateKeyPair(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateKeyPair",
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

func (p *KeyPairService) DeleteKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteKeyPairs",
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

func (p *KeyPairService) AttachKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachKeyPairs",
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

func (p *KeyPairService) DetachKeyPairs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachKeyPairs",
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

func (p *KeyPairService) ModifyKeyPairAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyKeyPairAttributes",
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

func init() { proto.RegisterFile("key_pair.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4e, 0xad, 0x8c,
	0x2f, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x95, 0x92, 0x2d, 0xcc, 0xcc, 0x4b, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0x89, 0x2f,
	0x4e, 0xc9, 0x8e, 0x2f, 0x2a, 0xcd, 0x49, 0xd5, 0x07, 0x11, 0x10, 0x75, 0x52, 0xd2, 0xe9, 0xf9,
	0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25,
	0x44, 0x52, 0x49, 0x8f, 0x4b, 0xc2, 0x3b, 0xb5, 0x32, 0x20, 0x31, 0xb3, 0x28, 0x18, 0x62, 0x5a,
	0x40, 0x51, 0x7e, 0x41, 0x6a, 0x51, 0x49, 0x66, 0x6a, 0xb1, 0x90, 0x10, 0x17, 0x4b, 0x55, 0x7e,
	0x5e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x6d, 0x74, 0x80, 0x99, 0x8b, 0x0f,
	0x55, 0x83, 0x90, 0x13, 0x97, 0x80, 0x4b, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x52, 0x2a, 0x54, 0xa6,
	0x58, 0x48, 0x4c, 0x0f, 0x62, 0xa9, 0x1e, 0xcc, 0x52, 0x3d, 0x57, 0x90, 0xa5, 0x52, 0x38, 0xc4,
	0x85, 0xec, 0xb9, 0x78, 0x9d, 0x8b, 0x52, 0x13, 0x4b, 0x60, 0x26, 0x90, 0x6c, 0x80, 0x03, 0x17,
	0x9f, 0x4b, 0x6a, 0x4e, 0x6a, 0x09, 0xf9, 0x4e, 0x70, 0xe0, 0xe2, 0x73, 0x2c, 0x29, 0x49, 0x4c,
	0xce, 0xa0, 0xc4, 0x04, 0x97, 0x54, 0x8a, 0x4c, 0xf0, 0xe4, 0x12, 0xf7, 0xcd, 0x4f, 0xc9, 0x4c,
	0xab, 0x84, 0x9a, 0xe0, 0x58, 0x52, 0x52, 0x94, 0x99, 0x54, 0x5a, 0x92, 0x4a, 0xb2, 0x51, 0x52,
	0x12, 0x5d, 0x1f, 0x59, 0x44, 0xb8, 0x38, 0x03, 0x33, 0xf3, 0xd2, 0x9d, 0x41, 0x49, 0x43, 0x88,
	0x1d, 0x6a, 0x66, 0x12, 0x1b, 0x58, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x15, 0x5c,
	0x7b, 0x50, 0x02, 0x00, 0x00,
}
