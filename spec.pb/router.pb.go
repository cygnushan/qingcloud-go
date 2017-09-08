// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

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

type RouterServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *RouterServiceProperties) Reset()                    { *m = RouterServiceProperties{} }
func (m *RouterServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*RouterServiceProperties) ProtoMessage()               {}
func (*RouterServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *RouterServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*RouterServiceProperties)(nil), "spec.RouterServiceProperties")
}

type RouterServiceInterface interface {
	DescribeRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	PowerOffRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	PowerOnRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	JoinRouter(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	LeaveRouter(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyRouterAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyRouterStaticAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CopyRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeRouterVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddRouterStaticEntries(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteRouterStaticEntries(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyRouterStaticEntryAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeRouterStaticEntries(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type RouterService struct {
	Config     *config.Config
	Properties *RouterServiceProperties
}

func NewRouterService(conf *config.Config, zone string) (p *RouterService, err error) {
	return &RouterService{
		Config:     conf,
		Properties: &RouterServiceProperties{Zone: zone},
	}, nil
}

func (p *RouterService) DescribeRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouters",
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
func (p *RouterService) CreateRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRouters",
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
func (p *RouterService) DeleteRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRouters",
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
func (p *RouterService) UpdateRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateRouters",
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
func (p *RouterService) PowerOffRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOffRouters",
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
func (p *RouterService) PowerOnRouters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOnRouters",
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
func (p *RouterService) JoinRouter(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "JoinRouter",
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
func (p *RouterService) LeaveRouter(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "LeaveRouter",
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
func (p *RouterService) ModifyRouterAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRouterAttributes",
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
func (p *RouterService) DescribeRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouterStatics",
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
func (p *RouterService) AddRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddRouterStatics",
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
func (p *RouterService) ModifyRouterStaticAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRouterStaticAttributes",
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
func (p *RouterService) DeleteRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRouterStatics",
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
func (p *RouterService) CopyRouterStatics(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CopyRouterStatics",
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
func (p *RouterService) DescribeRouterVxnets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouterVxnets",
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
func (p *RouterService) AddRouterStaticEntries(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddRouterStaticEntries",
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
func (p *RouterService) DeleteRouterStaticEntries(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRouterStaticEntries",
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
func (p *RouterService) ModifyRouterStaticEntryAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRouterStaticEntryAttributes",
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
func (p *RouterService) DescribeRouterStaticEntries(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouterStaticEntries",
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

func init() { proto.RegisterFile("router.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0x43, 0x42, 0xde, 0xe4, 0x1d, 0x45, 0x74, 0x55, 0x54, 0xf0, 0xa0, 0x9e, 0xbc, 0x58,
	0x12, 0xbd, 0x6a, 0xb4, 0x02, 0x6a, 0x54, 0x94, 0x40, 0xf4, 0xde, 0x3f, 0x53, 0xb2, 0x09, 0x76,
	0x37, 0xdb, 0x29, 0x5a, 0xbf, 0xa3, 0xdf, 0xc9, 0xd0, 0xd5, 0x84, 0x8a, 0x1e, 0xba, 0x7b, 0x6b,
	0x67, 0xe7, 0xf9, 0xcd, 0xcc, 0x33, 0x03, 0xcb, 0x4a, 0xa4, 0x84, 0xca, 0x91, 0x4a, 0x90, 0x60,
	0xd5, 0x44, 0x62, 0xd0, 0x6c, 0x8d, 0x85, 0x18, 0x4f, 0xb0, 0x9d, 0xc7, 0xfc, 0x34, 0x6a, 0xe3,
	0x8b, 0xa4, 0x4c, 0xa7, 0x1c, 0x1c, 0xc1, 0xd6, 0x30, 0x97, 0x8c, 0x50, 0x4d, 0x79, 0x80, 0x03,
	0x25, 0x24, 0x2a, 0xe2, 0x98, 0x30, 0x06, 0xd5, 0x77, 0x11, 0xe3, 0x76, 0x65, 0xaf, 0x72, 0xf8,
	0x7f, 0x98, 0x7f, 0x1f, 0x7f, 0x00, 0xd4, 0x0a, 0xf9, 0xcc, 0x85, 0x7a, 0x17, 0x93, 0x40, 0x71,
	0x1f, 0xf5, 0x43, 0xc2, 0x1a, 0x8e, 0xae, 0xe8, 0x7c, 0x57, 0x74, 0x7a, 0xb3, 0x8a, 0xcd, 0x3f,
	0xe2, 0xec, 0x1c, 0x6a, 0x1d, 0x85, 0x1e, 0xd9, 0x00, 0xba, 0x38, 0x41, 0x2b, 0xc0, 0x93, 0x0c,
	0x2d, 0x3a, 0x70, 0xa1, 0x3e, 0x10, 0xaf, 0xa8, 0x1e, 0xa3, 0xc8, 0x14, 0x71, 0x01, 0x2b, 0x1a,
	0x11, 0x9b, 0x12, 0x4e, 0x01, 0x6e, 0x05, 0xff, 0x92, 0x97, 0x56, 0x9f, 0xc1, 0xd2, 0x3d, 0x7a,
	0x53, 0x34, 0x94, 0xdf, 0x40, 0xa3, 0x2f, 0x42, 0x1e, 0x65, 0x5a, 0xef, 0x12, 0x29, 0xee, 0xa7,
	0x84, 0xe5, 0xc7, 0xb8, 0x86, 0xcd, 0xe2, 0x45, 0x8d, 0xc8, 0x23, 0x1e, 0x94, 0x07, 0x5d, 0xc2,
	0xaa, 0x1b, 0x86, 0x76, 0x8c, 0x07, 0xd8, 0x9d, 0x1f, 0x4b, 0x63, 0x2c, 0x86, 0xeb, 0xc1, 0xfa,
	0xfc, 0xa9, 0x9a, 0xb6, 0xd5, 0x81, 0xb5, 0x8e, 0x90, 0x99, 0x1d, 0xe4, 0x0a, 0x36, 0x8a, 0x46,
	0x3f, 0xbf, 0xc5, 0x48, 0x89, 0xc9, 0xea, 0x7f, 0xf8, 0xdc, 0x8b, 0x49, 0x71, 0x03, 0x77, 0xee,
	0x60, 0x67, 0xd1, 0x1d, 0x53, 0xd8, 0x08, 0xf6, 0x17, 0x57, 0x37, 0x83, 0x65, 0x16, 0xfb, 0xeb,
	0x43, 0xeb, 0xb7, 0xe3, 0x34, 0xec, 0xd1, 0xff, 0x97, 0xff, 0x9f, 0x7c, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x86, 0x74, 0xca, 0x59, 0xb8, 0x05, 0x00, 0x00,
}
