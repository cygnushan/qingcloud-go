// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"

var _ = config.Config{}
var _ = request.Request{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type S2ServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *S2ServiceProperties) Reset()                    { *m = S2ServiceProperties{} }
func (m *S2ServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*S2ServiceProperties) ProtoMessage()               {}
func (*S2ServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *S2ServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*S2ServiceProperties)(nil), "spec.S2ServiceProperties")
}

type S2ServiceInterface interface {
	CreateS2Server(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyS2Server(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	PowerOnS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	PowerOffS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ChangeS2ServerVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateS2SharedTarget(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	EnableS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DisableS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyS2SharedTargetAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AttachToS2SharedTarget(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DetachFromS2SharedTarget(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeS2DefaultParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateS2Group(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeS2Groups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyS2Group(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteS2Groups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateS2Account(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeS2Accounts(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyS2Account(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteS2Accounts(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AssociateS2AccountGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DissociateS2AccountGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type S2Service struct {
	Config     *config.Config
	Properties *S2ServiceProperties
}

func NewS2Service(conf *config.Config, zone string) (p *S2Service, err error) {
	return &S2Service{
		Config:     conf,
		Properties: &S2ServiceProperties{Zone: zone},
	}, nil
}

func (p *S2Service) CreateS2Server(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2Server",
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
func (p *S2Service) DescribeS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2Servers",
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
func (p *S2Service) ModifyS2Server(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2Server",
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
func (p *S2Service) ResizeS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeS2Servers",
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
func (p *S2Service) DeleteS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2Servers",
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
func (p *S2Service) PowerOnS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOnS2Servers",
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
func (p *S2Service) PowerOffS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOffS2Servers",
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
func (p *S2Service) UpdateS2Servers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateS2Servers",
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
func (p *S2Service) ChangeS2ServerVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeS2ServerVxnet",
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
func (p *S2Service) CreateS2SharedTarget(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2SharedTarget",
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
func (p *S2Service) DescribeS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2SharedTargets",
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
func (p *S2Service) DeleteS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2SharedTargets",
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
func (p *S2Service) EnableS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "EnableS2SharedTargets",
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
func (p *S2Service) DisableS2SharedTargets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DisableS2SharedTargets",
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
func (p *S2Service) ModifyS2SharedTargetAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2SharedTargetAttributes",
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
func (p *S2Service) AttachToS2SharedTarget(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachToS2SharedTarget",
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
func (p *S2Service) DetachFromS2SharedTarget(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachFromS2SharedTarget",
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
func (p *S2Service) DescribeS2DefaultParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2DefaultParameters",
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
func (p *S2Service) CreateS2Group(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2Group",
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
func (p *S2Service) DescribeS2Groups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2Groups",
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
func (p *S2Service) ModifyS2Group(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2Group",
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
func (p *S2Service) DeleteS2Groups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2Groups",
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
func (p *S2Service) CreateS2Account(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2Account",
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
func (p *S2Service) DescribeS2Accounts(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2Accounts",
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
func (p *S2Service) ModifyS2Account(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2Account",
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
func (p *S2Service) DeleteS2Accounts(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2Accounts",
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
func (p *S2Service) AssociateS2AccountGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateS2AccountGroup",
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
func (p *S2Service) DissociateS2AccountGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateS2AccountGroup",
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

func init() { proto.RegisterFile("s2.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0xab, 0xda, 0x40,
	0x14, 0x85, 0x11, 0xa4, 0xd4, 0x0b, 0x35, 0x35, 0xb6, 0x2a, 0x0a, 0xa5, 0x74, 0xd5, 0x6e, 0x22,
	0xd8, 0x1f, 0x50, 0xd3, 0x44, 0xad, 0x05, 0x69, 0x50, 0xdb, 0xfd, 0x24, 0xde, 0xc4, 0x81, 0x98,
	0x09, 0x33, 0x93, 0xbe, 0xa7, 0x3f, 0xef, 0xfd, 0xb2, 0x87, 0xc9, 0x4b, 0xcc, 0xe6, 0x2d, 0x32,
	0xe3, 0x2e, 0x19, 0x86, 0xef, 0x9e, 0x7b, 0xce, 0x19, 0x78, 0x2b, 0x66, 0x56, 0xca, 0x99, 0x64,
	0x66, 0x5b, 0xa4, 0x18, 0x8c, 0x27, 0x11, 0x63, 0x51, 0x8c, 0xd3, 0xfc, 0xcc, 0xcf, 0xc2, 0x29,
	0x9e, 0x52, 0x79, 0x2e, 0xae, 0x7c, 0xf9, 0x06, 0xfd, 0xdd, 0x6c, 0x87, 0xfc, 0x3f, 0x0d, 0xd0,
	0xe3, 0x2c, 0x45, 0x2e, 0x29, 0x0a, 0xd3, 0x84, 0xf6, 0x85, 0x25, 0x38, 0x6a, 0x7d, 0x6e, 0x7d,
	0xed, 0x6c, 0xf3, 0xef, 0xd9, 0x93, 0x01, 0x9d, 0xea, 0xae, 0x39, 0x87, 0xae, 0xc3, 0x91, 0x48,
	0x2c, 0x8e, 0x90, 0x9b, 0x03, 0xab, 0x18, 0x64, 0x95, 0x83, 0xac, 0xc5, 0x75, 0xd0, 0xf8, 0x95,
	0x73, 0xd3, 0x81, 0x9e, 0x8b, 0x22, 0xe0, 0xd4, 0xaf, 0x18, 0xa2, 0x31, 0x64, 0x0e, 0xdd, 0x0d,
	0x3b, 0xd0, 0xf0, 0xac, 0x2c, 0xc3, 0x06, 0x63, 0x8b, 0x82, 0x5e, 0x34, 0x44, 0xd8, 0x60, 0xb8,
	0x18, 0xa3, 0xd4, 0x40, 0xfc, 0x84, 0xf7, 0x1e, 0x7b, 0x40, 0xfe, 0x27, 0x51, 0x67, 0x38, 0xd0,
	0x2b, 0x18, 0x61, 0xa8, 0xb5, 0xcb, 0xdf, 0xf4, 0x40, 0x74, 0x76, 0x59, 0x40, 0xdf, 0x39, 0x92,
	0x24, 0xaa, 0x10, 0xff, 0x1e, 0x13, 0x94, 0x8d, 0x31, 0x4b, 0xf8, 0x50, 0x35, 0xec, 0x48, 0x38,
	0x1e, 0xf6, 0x84, 0x47, 0x0a, 0x9c, 0x35, 0x0c, 0x6b, 0x3d, 0xab, 0x91, 0x9a, 0x6f, 0xb6, 0x82,
	0x8f, 0x55, 0xd0, 0xba, 0xa0, 0x45, 0x42, 0xfc, 0x58, 0x1b, 0xf4, 0x0b, 0x06, 0x2e, 0x15, 0xf7,
	0x20, 0x79, 0xf0, 0xa9, 0x7a, 0x49, 0x35, 0x90, 0x2d, 0x25, 0xa7, 0x7e, 0x26, 0x51, 0x49, 0x9b,
	0x2d, 0x25, 0x09, 0x8e, 0x7b, 0xa6, 0x19, 0xe1, 0x6f, 0x18, 0xb9, 0x78, 0x25, 0x2d, 0x39, 0x3b,
	0x69, 0xb2, 0x36, 0x30, 0xb9, 0xd5, 0xc1, 0xc5, 0x90, 0x64, 0xb1, 0xf4, 0x08, 0x27, 0x27, 0x94,
	0x2a, 0x65, 0xff, 0x01, 0xef, 0xca, 0x96, 0xae, 0x38, 0xcb, 0x52, 0x95, 0x97, 0x7f, 0xd3, 0x93,
	0x23, 0x94, 0x44, 0x94, 0xd9, 0xa9, 0x89, 0x98, 0x43, 0xb7, 0x2c, 0xb6, 0xa2, 0x04, 0x1b, 0x8c,
	0xd2, 0x07, 0x3b, 0x08, 0x58, 0x96, 0x34, 0x4f, 0xc6, 0x05, 0xf3, 0xe6, 0xc4, 0x0b, 0x44, 0x49,
	0x48, 0xe9, 0x85, 0xaa, 0x90, 0x3c, 0x92, 0xc2, 0x0d, 0x65, 0x19, 0x6b, 0x18, 0xda, 0x42, 0xb0,
	0x80, 0xd6, 0x2d, 0x51, 0x0b, 0xe7, 0xda, 0x7e, 0x7a, 0x1f, 0x96, 0xff, 0x26, 0xff, 0xff, 0xfe,
	0x1c, 0x00, 0x00, 0xff, 0xff, 0x59, 0x00, 0x28, 0x31, 0x25, 0x08, 0x00, 0x00,
}
