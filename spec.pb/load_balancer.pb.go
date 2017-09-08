// Code generated by protoc-gen-go. DO NOT EDIT.
// source: load_balancer.proto

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

type LoadBalancerServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *LoadBalancerServiceProperties) Reset()                    { *m = LoadBalancerServiceProperties{} }
func (m *LoadBalancerServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerServiceProperties) ProtoMessage()               {}
func (*LoadBalancerServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *LoadBalancerServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*LoadBalancerServiceProperties)(nil), "spec.LoadBalancerServiceProperties")
}

type LoadBalancerServiceInterface interface {
	CreateLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AssociateEipsToLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DissociateEipsFromLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerListenerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerBackendAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerPolicyAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ApplyLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerPolicyRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateServerCertificate(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyServerCertificateAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type LoadBalancerService struct {
	Config     *config.Config
	Properties *LoadBalancerServiceProperties
}

func NewLoadBalancerService(conf *config.Config, zone string) (p *LoadBalancerService, err error) {
	return &LoadBalancerService{
		Config:     conf,
		Properties: &LoadBalancerServiceProperties{Zone: zone},
	}, nil
}

func (p *LoadBalancerService) CreateLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancer",
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
func (p *LoadBalancerService) DescribeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancers",
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
func (p *LoadBalancerService) DeleteLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancers",
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
func (p *LoadBalancerService) ModifyLoadBalancerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerAttributes",
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
func (p *LoadBalancerService) StartLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartLoadBalancers",
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
func (p *LoadBalancerService) StopLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopLoadBalancers",
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
func (p *LoadBalancerService) UpdateLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateLoadBalancers",
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
func (p *LoadBalancerService) ResizeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeLoadBalancers",
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
func (p *LoadBalancerService) AssociateEipsToLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEipsToLoadBalancer",
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
func (p *LoadBalancerService) DissociateEipsFromLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEipsFromLoadBalancer",
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
func (p *LoadBalancerService) AddLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerListeners",
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
func (p *LoadBalancerService) DescribeLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerListeners",
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
func (p *LoadBalancerService) DeleteLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerListeners",
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
func (p *LoadBalancerService) ModifyLoadBalancerListenerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerListenerAttributes",
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
func (p *LoadBalancerService) AddLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerBackends",
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
func (p *LoadBalancerService) DescribeLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerBackends",
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
func (p *LoadBalancerService) DeleteLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerBackends",
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
func (p *LoadBalancerService) ModifyLoadBalancerBackendAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerBackendAttributes",
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
func (p *LoadBalancerService) CreateLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancerPolicy",
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
func (p *LoadBalancerService) DescribeLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicies",
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
func (p *LoadBalancerService) ModifyLoadBalancerPolicyAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyAttributes",
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
func (p *LoadBalancerService) ApplyLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyLoadBalancerPolicy",
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
func (p *LoadBalancerService) DeleteLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicies",
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
func (p *LoadBalancerService) AddLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerPolicyRules",
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
func (p *LoadBalancerService) DescribeLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicyRules",
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
func (p *LoadBalancerService) ModifyLoadBalancerPolicyRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyRuleAttributes",
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
func (p *LoadBalancerService) DeleteLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicyRules",
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
func (p *LoadBalancerService) CreateServerCertificate(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateServerCertificate",
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
func (p *LoadBalancerService) DescribeServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeServerCertificates",
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
func (p *LoadBalancerService) ModifyServerCertificateAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyServerCertificateAttributes",
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
func (p *LoadBalancerService) DeleteServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteServerCertificates",
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

func init() { proto.RegisterFile("load_balancer.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x29, 0x94, 0xc1, 0xce, 0xdd, 0x64, 0xb6, 0x06, 0x67, 0xdd, 0x9f, 0x6e, 0x8c, 0x5d,
	0xb9, 0xb0, 0x3e, 0x41, 0x1a, 0x67, 0xa3, 0xc5, 0xd9, 0xb2, 0xb8, 0x1d, 0xbb, 0x1b, 0xb2, 0x7c,
	0x52, 0xc4, 0x5c, 0x4b, 0x48, 0xca, 0xc0, 0x7d, 0x84, 0x3d, 0xf5, 0xb0, 0xd5, 0x80, 0x17, 0x39,
	0x17, 0x91, 0x72, 0x67, 0xcb, 0xd6, 0x8f, 0xef, 0x7c, 0xe7, 0x93, 0x38, 0x10, 0x55, 0x82, 0x96,
	0xbf, 0x0a, 0x5a, 0xd1, 0x9a, 0xa1, 0x4a, 0xa4, 0x12, 0x46, 0x90, 0x63, 0x2d, 0x91, 0xc5, 0xe3,
	0x3b, 0x21, 0xee, 0x2a, 0x3c, 0xef, 0xd6, 0x8a, 0xf5, 0xea, 0x1c, 0xef, 0xa5, 0x69, 0xec, 0x2f,
	0x67, 0x17, 0x70, 0x9a, 0x09, 0x5a, 0x5e, 0x3e, 0x6e, 0xcc, 0x51, 0xfd, 0xe1, 0x0c, 0x17, 0x4a,
	0x48, 0x54, 0x86, 0xa3, 0x26, 0x04, 0x8e, 0x1f, 0x44, 0x8d, 0xa3, 0xa3, 0x37, 0x47, 0x1f, 0x9f,
	0x2e, 0xbb, 0xe7, 0x4f, 0x7f, 0x23, 0x88, 0x06, 0x76, 0x91, 0x14, 0xc8, 0x54, 0x21, 0x35, 0xd8,
	0xff, 0x48, 0x5e, 0x24, 0x56, 0x40, 0xb2, 0x11, 0x90, 0xcc, 0x5a, 0x01, 0xf1, 0x8e, 0x75, 0xf2,
	0x05, 0x9e, 0xa7, 0xa8, 0x99, 0xe2, 0xc5, 0x7f, 0x1c, 0xbd, 0x37, 0x68, 0x06, 0x51, 0x8a, 0x15,
	0x9a, 0x40, 0xcc, 0x57, 0x78, 0x39, 0x17, 0x25, 0x5f, 0x35, 0x7d, 0xcc, 0xc4, 0x18, 0xc5, 0x8b,
	0xb5, 0xc1, 0xfd, 0x79, 0x29, 0x90, 0xdc, 0x50, 0x65, 0xc2, 0x54, 0x4d, 0xe1, 0x59, 0x6e, 0x84,
	0x0c, 0x76, 0xe8, 0x56, 0x96, 0xd4, 0x84, 0x1b, 0xbd, 0x44, 0xcd, 0x1f, 0x02, 0x31, 0x73, 0x18,
	0x4f, 0xb4, 0x16, 0x8c, 0x53, 0x83, 0x33, 0x2e, 0xf5, 0x8d, 0x08, 0xca, 0xd1, 0x02, 0x5e, 0xa5,
	0xbc, 0xcf, 0xfb, 0xac, 0xc4, 0x7d, 0x10, 0xf1, 0x1a, 0x46, 0x93, 0xb2, 0xec, 0x23, 0x32, 0xae,
	0x0d, 0xd6, 0x3e, 0xc5, 0x7e, 0x83, 0xd3, 0xa1, 0x94, 0xfb, 0x03, 0xe7, 0x30, 0x76, 0xd3, 0xee,
	0x8f, 0xfb, 0x01, 0xef, 0xdd, 0xd4, 0x6f, 0x70, 0x01, 0xe9, 0xbf, 0x82, 0x93, 0x2d, 0x0f, 0x2f,
	0x29, 0xfb, 0x8d, 0x75, 0xe9, 0x75, 0x30, 0x87, 0x2c, 0xf4, 0xe6, 0x65, 0x10, 0xbb, 0x0e, 0x7a,
	0xd3, 0x6e, 0xe1, 0x9d, 0x6b, 0xe0, 0x23, 0x2d, 0xc0, 0xbf, 0x6b, 0x18, 0xb9, 0x77, 0xec, 0x42,
	0x54, 0x9c, 0x35, 0x87, 0x32, 0xb0, 0xa3, 0x71, 0x0f, 0x6d, 0x37, 0x70, 0xe6, 0x96, 0x6c, 0xb5,
	0x05, 0x26, 0x46, 0xca, 0xaa, 0x39, 0x40, 0xc1, 0x83, 0x1d, 0xf6, 0x2e, 0x37, 0x83, 0x78, 0x2b,
	0xca, 0x56, 0xd6, 0x72, 0x5d, 0x79, 0xd0, 0xbe, 0xc3, 0xeb, 0x9d, 0xcd, 0xf0, 0x44, 0xfe, 0x84,
	0x0f, 0xbb, 0xfa, 0xd1, 0x02, 0x03, 0x7a, 0xd2, 0xdd, 0x5e, 0x83, 0x46, 0x7a, 0x4a, 0xbd, 0x82,
	0x13, 0x1b, 0xeb, 0x76, 0x96, 0x40, 0x35, 0x6d, 0xa7, 0x8f, 0x15, 0x67, 0xd4, 0xa0, 0x5f, 0x93,
	0xad, 0x91, 0x0e, 0x6c, 0x7f, 0x61, 0x39, 0xbc, 0xb5, 0x1e, 0x3a, 0xac, 0xb0, 0x43, 0x6c, 0xed,
	0x0b, 0x17, 0x58, 0x3c, 0xe9, 0xde, 0x2f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x44, 0x84,
	0xbe, 0x02, 0x0a, 0x00, 0x00,
}
