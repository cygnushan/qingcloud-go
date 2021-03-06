// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: load_balancer.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"
import "github.com/chai2010/qingcloud-go/pkg/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}
var _ = data.Operation{}

type LoadBalancerServiceInterface interface {
	CreateLoadBalancer(in *CreateLoadBalancerInput) (out *CreateLoadBalancerOutput, err error)
	DescribeLoadBalancers(in *DescribeLoadBalancersInput) (out *DescribeLoadBalancersOutput, err error)
	DeleteLoadBalancers(in *DeleteLoadBalancersInput) (out *DeleteLoadBalancersOutput, err error)
	ModifyLoadBalancerAttributes(in *ModifyLoadBalancerAttributesInput) (out *ModifyLoadBalancerAttributesOutput, err error)
	StartLoadBalancers(in *StartLoadBalancersInput) (out *StartLoadBalancersOutput, err error)
	StopLoadBalancers(in *StopLoadBalancersInput) (out *StopLoadBalancersOutput, err error)
	UpdateLoadBalancers(in *UpdateLoadBalancersInput) (out *UpdateLoadBalancersOutput, err error)
	ResizeLoadBalancers(in *ResizeLoadBalancersInput) (out *ResizeLoadBalancersOutput, err error)
	AssociateEipsToLoadBalancer(in *AssociateEipsToLoadBalancerInput) (out *AssociateEipsToLoadBalancerOutput, err error)
	DissociateEipsFromLoadBalancer(in *DissociateEipsFromLoadBalancerInput) (out *DissociateEipsFromLoadBalancerOutput, err error)
	AddLoadBalancerListeners(in *AddLoadBalancerListenersInput) (out *AddLoadBalancerListenersOutput, err error)
	DescribeLoadBalancerListeners(in *DescribeLoadBalancerListenersInput) (out *DescribeLoadBalancerListenersOutput, err error)
	DeleteLoadBalancerListeners(in *DeleteLoadBalancerListenersInput) (out *DeleteLoadBalancerListenersOutput, err error)
	ModifyLoadBalancerListenerAttributes(in *ModifyLoadBalancerListenerAttributesInput) (out *ModifyLoadBalancerListenerAttributesOutput, err error)
	AddLoadBalancerBackends(in *AddLoadBalancerBackendsInput) (out *AddLoadBalancerBackendsOutput, err error)
	DescribeLoadBalancerBackends(in *DescribeLoadBalancerBackendsInput) (out *DescribeLoadBalancerBackendsOutput, err error)
	DeleteLoadBalancerBackends(in *DeleteLoadBalancerBackendsInput) (out *DeleteLoadBalancerBackendsOutput, err error)
	ModifyLoadBalancerBackendAttributes(in *ModifyLoadBalancerBackendAttributesInput) (out *ModifyLoadBalancerBackendAttributesOutput, err error)
	CreateLoadBalancerPolicy(in *CreateLoadBalancerPolicyInput) (out *CreateLoadBalancerPolicyOutput, err error)
	DescribeLoadBalancerPolicies(in *DescribeLoadBalancerPoliciesInput) (out *DescribeLoadBalancerPoliciesOutput, err error)
	ModifyLoadBalancerPolicyAttributes(in *ModifyLoadBalancerPolicyAttributesInput) (out *ModifyLoadBalancerPolicyAttributesOutput, err error)
	ApplyLoadBalancerPolicy(in *ApplyLoadBalancerPolicyInput) (out *ApplyLoadBalancerPolicyOutput, err error)
	DeleteLoadBalancerPolicies(in *DeleteLoadBalancerPoliciesInput) (out *DeleteLoadBalancerPoliciesOutput, err error)
	AddLoadBalancerPolicyRules(in *AddLoadBalancerPolicyRulesInput) (out *AddLoadBalancerPolicyRulesOutput, err error)
	DescribeLoadBalancerPolicyRules(in *DescribeLoadBalancerPolicyRulesInput) (out *DescribeLoadBalancerPolicyRulesOutput, err error)
	ModifyLoadBalancerPolicyRuleAttributes(in *ModifyLoadBalancerPolicyRuleAttributesInput) (out *ModifyLoadBalancerPolicyRuleAttributesOutput, err error)
	DeleteLoadBalancerPolicyRules(in *DeleteLoadBalancerPolicyRulesInput) (out *DeleteLoadBalancerPolicyRulesOutput, err error)
	CreateServerCertificate(in *CreateServerCertificateInput) (out *CreateServerCertificateOutput, err error)
	DescribeServerCertificates(in *DescribeServerCertificatesInput) (out *DescribeServerCertificatesOutput, err error)
	ModifyServerCertificateAttributes(in *ModifyServerCertificateAttributesInput) (out *ModifyServerCertificateAttributesOutput, err error)
	DeleteServerCertificates(in *DeleteServerCertificatesInput) (out *DeleteServerCertificatesOutput, err error)
}

type LoadBalancerService struct {
	Config           *config.Config
	Properties       *LoadBalancerServiceProperties
	LastResponseBody string
}

func NewLoadBalancerService(conf *config.Config, zone string) (p *LoadBalancerService) {
	return &LoadBalancerService{
		Config:     conf,
		Properties: &LoadBalancerServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) LoadBalancer(zone string) (*LoadBalancerService, error) {
	properties := &LoadBalancerServiceProperties{
		Zone: proto.String(zone),
	}

	return &LoadBalancerService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *LoadBalancerService) CreateLoadBalancer(in *CreateLoadBalancerInput) (out *CreateLoadBalancerOutput, err error) {
	if in == nil {
		in = &CreateLoadBalancerInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancer",
		RequestMethod: "GET",
	}

	x := &CreateLoadBalancerOutput{}
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

func (p *LoadBalancerService) DescribeLoadBalancers(in *DescribeLoadBalancersInput) (out *DescribeLoadBalancersOutput, err error) {
	if in == nil {
		in = &DescribeLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancers",
		RequestMethod: "GET",
	}

	x := &DescribeLoadBalancersOutput{}
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

func (p *LoadBalancerService) DeleteLoadBalancers(in *DeleteLoadBalancersInput) (out *DeleteLoadBalancersOutput, err error) {
	if in == nil {
		in = &DeleteLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancers",
		RequestMethod: "GET",
	}

	x := &DeleteLoadBalancersOutput{}
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

func (p *LoadBalancerService) ModifyLoadBalancerAttributes(in *ModifyLoadBalancerAttributesInput) (out *ModifyLoadBalancerAttributesOutput, err error) {
	if in == nil {
		in = &ModifyLoadBalancerAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyLoadBalancerAttributesOutput{}
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

func (p *LoadBalancerService) StartLoadBalancers(in *StartLoadBalancersInput) (out *StartLoadBalancersOutput, err error) {
	if in == nil {
		in = &StartLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartLoadBalancers",
		RequestMethod: "GET",
	}

	x := &StartLoadBalancersOutput{}
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

func (p *LoadBalancerService) StopLoadBalancers(in *StopLoadBalancersInput) (out *StopLoadBalancersOutput, err error) {
	if in == nil {
		in = &StopLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopLoadBalancers",
		RequestMethod: "GET",
	}

	x := &StopLoadBalancersOutput{}
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

func (p *LoadBalancerService) UpdateLoadBalancers(in *UpdateLoadBalancersInput) (out *UpdateLoadBalancersOutput, err error) {
	if in == nil {
		in = &UpdateLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateLoadBalancers",
		RequestMethod: "GET",
	}

	x := &UpdateLoadBalancersOutput{}
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

func (p *LoadBalancerService) ResizeLoadBalancers(in *ResizeLoadBalancersInput) (out *ResizeLoadBalancersOutput, err error) {
	if in == nil {
		in = &ResizeLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeLoadBalancers",
		RequestMethod: "GET",
	}

	x := &ResizeLoadBalancersOutput{}
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

func (p *LoadBalancerService) AssociateEipsToLoadBalancer(in *AssociateEipsToLoadBalancerInput) (out *AssociateEipsToLoadBalancerOutput, err error) {
	if in == nil {
		in = &AssociateEipsToLoadBalancerInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEipsToLoadBalancer",
		RequestMethod: "GET",
	}

	x := &AssociateEipsToLoadBalancerOutput{}
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

func (p *LoadBalancerService) DissociateEipsFromLoadBalancer(in *DissociateEipsFromLoadBalancerInput) (out *DissociateEipsFromLoadBalancerOutput, err error) {
	if in == nil {
		in = &DissociateEipsFromLoadBalancerInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEipsFromLoadBalancer",
		RequestMethod: "GET",
	}

	x := &DissociateEipsFromLoadBalancerOutput{}
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

func (p *LoadBalancerService) AddLoadBalancerListeners(in *AddLoadBalancerListenersInput) (out *AddLoadBalancerListenersOutput, err error) {
	if in == nil {
		in = &AddLoadBalancerListenersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerListeners",
		RequestMethod: "GET",
	}

	x := &AddLoadBalancerListenersOutput{}
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

func (p *LoadBalancerService) DescribeLoadBalancerListeners(in *DescribeLoadBalancerListenersInput) (out *DescribeLoadBalancerListenersOutput, err error) {
	if in == nil {
		in = &DescribeLoadBalancerListenersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerListeners",
		RequestMethod: "GET",
	}

	x := &DescribeLoadBalancerListenersOutput{}
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

func (p *LoadBalancerService) DeleteLoadBalancerListeners(in *DeleteLoadBalancerListenersInput) (out *DeleteLoadBalancerListenersOutput, err error) {
	if in == nil {
		in = &DeleteLoadBalancerListenersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerListeners",
		RequestMethod: "GET",
	}

	x := &DeleteLoadBalancerListenersOutput{}
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

func (p *LoadBalancerService) ModifyLoadBalancerListenerAttributes(in *ModifyLoadBalancerListenerAttributesInput) (out *ModifyLoadBalancerListenerAttributesOutput, err error) {
	if in == nil {
		in = &ModifyLoadBalancerListenerAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerListenerAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyLoadBalancerListenerAttributesOutput{}
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

func (p *LoadBalancerService) AddLoadBalancerBackends(in *AddLoadBalancerBackendsInput) (out *AddLoadBalancerBackendsOutput, err error) {
	if in == nil {
		in = &AddLoadBalancerBackendsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerBackends",
		RequestMethod: "GET",
	}

	x := &AddLoadBalancerBackendsOutput{}
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

func (p *LoadBalancerService) DescribeLoadBalancerBackends(in *DescribeLoadBalancerBackendsInput) (out *DescribeLoadBalancerBackendsOutput, err error) {
	if in == nil {
		in = &DescribeLoadBalancerBackendsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerBackends",
		RequestMethod: "GET",
	}

	x := &DescribeLoadBalancerBackendsOutput{}
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

func (p *LoadBalancerService) DeleteLoadBalancerBackends(in *DeleteLoadBalancerBackendsInput) (out *DeleteLoadBalancerBackendsOutput, err error) {
	if in == nil {
		in = &DeleteLoadBalancerBackendsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerBackends",
		RequestMethod: "GET",
	}

	x := &DeleteLoadBalancerBackendsOutput{}
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

func (p *LoadBalancerService) ModifyLoadBalancerBackendAttributes(in *ModifyLoadBalancerBackendAttributesInput) (out *ModifyLoadBalancerBackendAttributesOutput, err error) {
	if in == nil {
		in = &ModifyLoadBalancerBackendAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerBackendAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyLoadBalancerBackendAttributesOutput{}
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

func (p *LoadBalancerService) CreateLoadBalancerPolicy(in *CreateLoadBalancerPolicyInput) (out *CreateLoadBalancerPolicyOutput, err error) {
	if in == nil {
		in = &CreateLoadBalancerPolicyInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancerPolicy",
		RequestMethod: "GET",
	}

	x := &CreateLoadBalancerPolicyOutput{}
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

func (p *LoadBalancerService) DescribeLoadBalancerPolicies(in *DescribeLoadBalancerPoliciesInput) (out *DescribeLoadBalancerPoliciesOutput, err error) {
	if in == nil {
		in = &DescribeLoadBalancerPoliciesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicies",
		RequestMethod: "GET",
	}

	x := &DescribeLoadBalancerPoliciesOutput{}
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

func (p *LoadBalancerService) ModifyLoadBalancerPolicyAttributes(in *ModifyLoadBalancerPolicyAttributesInput) (out *ModifyLoadBalancerPolicyAttributesOutput, err error) {
	if in == nil {
		in = &ModifyLoadBalancerPolicyAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyLoadBalancerPolicyAttributesOutput{}
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

func (p *LoadBalancerService) ApplyLoadBalancerPolicy(in *ApplyLoadBalancerPolicyInput) (out *ApplyLoadBalancerPolicyOutput, err error) {
	if in == nil {
		in = &ApplyLoadBalancerPolicyInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyLoadBalancerPolicy",
		RequestMethod: "GET",
	}

	x := &ApplyLoadBalancerPolicyOutput{}
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

func (p *LoadBalancerService) DeleteLoadBalancerPolicies(in *DeleteLoadBalancerPoliciesInput) (out *DeleteLoadBalancerPoliciesOutput, err error) {
	if in == nil {
		in = &DeleteLoadBalancerPoliciesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicies",
		RequestMethod: "GET",
	}

	x := &DeleteLoadBalancerPoliciesOutput{}
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

func (p *LoadBalancerService) AddLoadBalancerPolicyRules(in *AddLoadBalancerPolicyRulesInput) (out *AddLoadBalancerPolicyRulesOutput, err error) {
	if in == nil {
		in = &AddLoadBalancerPolicyRulesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerPolicyRules",
		RequestMethod: "GET",
	}

	x := &AddLoadBalancerPolicyRulesOutput{}
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

func (p *LoadBalancerService) DescribeLoadBalancerPolicyRules(in *DescribeLoadBalancerPolicyRulesInput) (out *DescribeLoadBalancerPolicyRulesOutput, err error) {
	if in == nil {
		in = &DescribeLoadBalancerPolicyRulesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicyRules",
		RequestMethod: "GET",
	}

	x := &DescribeLoadBalancerPolicyRulesOutput{}
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

func (p *LoadBalancerService) ModifyLoadBalancerPolicyRuleAttributes(in *ModifyLoadBalancerPolicyRuleAttributesInput) (out *ModifyLoadBalancerPolicyRuleAttributesOutput, err error) {
	if in == nil {
		in = &ModifyLoadBalancerPolicyRuleAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyRuleAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyLoadBalancerPolicyRuleAttributesOutput{}
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

func (p *LoadBalancerService) DeleteLoadBalancerPolicyRules(in *DeleteLoadBalancerPolicyRulesInput) (out *DeleteLoadBalancerPolicyRulesOutput, err error) {
	if in == nil {
		in = &DeleteLoadBalancerPolicyRulesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicyRules",
		RequestMethod: "GET",
	}

	x := &DeleteLoadBalancerPolicyRulesOutput{}
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

func (p *LoadBalancerService) CreateServerCertificate(in *CreateServerCertificateInput) (out *CreateServerCertificateOutput, err error) {
	if in == nil {
		in = &CreateServerCertificateInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateServerCertificate",
		RequestMethod: "GET",
	}

	x := &CreateServerCertificateOutput{}
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

func (p *LoadBalancerService) DescribeServerCertificates(in *DescribeServerCertificatesInput) (out *DescribeServerCertificatesOutput, err error) {
	if in == nil {
		in = &DescribeServerCertificatesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeServerCertificates",
		RequestMethod: "GET",
	}

	x := &DescribeServerCertificatesOutput{}
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

func (p *LoadBalancerService) ModifyServerCertificateAttributes(in *ModifyServerCertificateAttributesInput) (out *ModifyServerCertificateAttributesOutput, err error) {
	if in == nil {
		in = &ModifyServerCertificateAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyServerCertificateAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyServerCertificateAttributesOutput{}
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

func (p *LoadBalancerService) DeleteServerCertificates(in *DeleteServerCertificatesInput) (out *DeleteServerCertificatesOutput, err error) {
	if in == nil {
		in = &DeleteServerCertificatesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteServerCertificates",
		RequestMethod: "GET",
	}

	x := &DeleteServerCertificatesOutput{}
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
