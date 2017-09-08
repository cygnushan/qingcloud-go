// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package service

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

type RouterServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *RouterServiceProperties) Reset()                    { *m = RouterServiceProperties{} }
func (m *RouterServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*RouterServiceProperties) ProtoMessage()               {}
func (*RouterServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *RouterServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*RouterServiceProperties)(nil), "service.RouterServiceProperties")
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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

func init() { proto.RegisterFile("router.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0x43, 0xf2, 0xe6, 0x35, 0x8e, 0x22, 0xba, 0x2a, 0x2a, 0x78, 0x50, 0x4f, 0x5e, 0x2c,
	0x89, 0x5e, 0x35, 0x5a, 0xa1, 0x6a, 0x54, 0x94, 0xd0, 0xe8, 0xbd, 0x7f, 0xa6, 0x64, 0x13, 0xec,
	0x36, 0xdb, 0x29, 0x5a, 0xbf, 0xa3, 0xdf, 0xc9, 0xd0, 0xd5, 0x84, 0x8a, 0x1e, 0xba, 0x7b, 0x6b,
	0x67, 0xe7, 0xf9, 0xcd, 0xcc, 0x33, 0x03, 0xcb, 0x52, 0x64, 0x84, 0xd2, 0x4a, 0xa4, 0x20, 0xc1,
	0x16, 0x52, 0x94, 0x13, 0x1e, 0x60, 0xab, 0x3d, 0x12, 0x62, 0x34, 0xc6, 0x4e, 0x11, 0xf6, 0xb3,
	0xa8, 0x83, 0x2f, 0x09, 0xe5, 0x2a, 0xeb, 0xe0, 0x08, 0xb6, 0x86, 0x85, 0xca, 0x55, 0xd9, 0x03,
	0x29, 0x12, 0x94, 0xc4, 0x31, 0x65, 0x0c, 0xfe, 0xbd, 0x8b, 0x18, 0xb7, 0x6b, 0x7b, 0xb5, 0xc3,
	0xc5, 0x61, 0xf1, 0x7d, 0xfc, 0x01, 0x50, 0x2f, 0xe5, 0x33, 0x1b, 0x1a, 0x3d, 0x4c, 0x03, 0xc9,
	0x7d, 0x54, 0x0f, 0x29, 0x6b, 0x5a, 0xaa, 0xa2, 0xf5, 0x5d, 0xd1, 0x72, 0xa6, 0x15, 0x5b, 0x7f,
	0xc4, 0xd9, 0x39, 0xd4, 0xbb, 0x12, 0x3d, 0x32, 0x01, 0xf4, 0x70, 0x8c, 0x46, 0x80, 0xa7, 0x24,
	0x34, 0xe8, 0xc0, 0x86, 0xc6, 0x40, 0xbc, 0xa2, 0x7c, 0x8c, 0x22, 0x5d, 0xc4, 0x05, 0xac, 0x28,
	0x44, 0xac, 0x4b, 0x38, 0x05, 0xb8, 0x15, 0xfc, 0x4b, 0x5e, 0x59, 0x7d, 0x06, 0x4b, 0xf7, 0xe8,
	0x4d, 0x50, 0x53, 0x7e, 0x03, 0xcd, 0xbe, 0x08, 0x79, 0x94, 0x2b, 0xbd, 0x4d, 0x24, 0xb9, 0x9f,
	0x11, 0x56, 0x1f, 0xe3, 0x1a, 0x36, 0xcb, 0x17, 0xe5, 0x92, 0x47, 0x3c, 0xa8, 0x0e, 0xba, 0x84,
	0x55, 0x3b, 0x0c, 0xcd, 0x18, 0x0f, 0xb0, 0x3b, 0x3b, 0x96, 0xc2, 0x18, 0x0c, 0xe7, 0xc0, 0xfa,
	0xec, 0xa9, 0xea, 0xb6, 0xd5, 0x85, 0xb5, 0xae, 0x48, 0x72, 0x33, 0xc8, 0x15, 0x6c, 0x94, 0x8d,
	0x7e, 0x7e, 0x8b, 0x91, 0x52, 0x9d, 0xd5, 0xff, 0xf0, 0xd9, 0x89, 0x49, 0x72, 0x0d, 0x77, 0xee,
	0x60, 0x67, 0xde, 0x1d, 0x5d, 0x98, 0x0b, 0xfb, 0xf3, 0xab, 0x9b, 0xc2, 0x72, 0x83, 0xfd, 0xf5,
	0xa1, 0xfd, 0xdb, 0x71, 0x6a, 0xf6, 0xe8, 0xff, 0x2f, 0xfe, 0x4f, 0x3e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xe7, 0x72, 0xf5, 0x2c, 0xbb, 0x05, 0x00, 0x00,
}
