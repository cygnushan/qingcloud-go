// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: s2.proto

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
var _ = logger.SetLevel
var _ = request.Request{}
var _ = data.Operation{}

// See https://docs.qingcloud.com/api/s2/index.html
type S2ServiceInterface interface {
	CreateS2Server(in *CreateS2ServerInput) (out *CreateS2ServerOutput, err error)
	DescribeS2Servers(in *DescribeS2ServersInput) (out *DescribeS2ServersOutput, err error)
	ModifyS2Server(in *ModifyS2ServerInput) (out *ModifyS2ServerOutput, err error)
	ResizeS2Servers(in *ResizeS2ServersInput) (out *ResizeS2ServersOutput, err error)
	DeleteS2Servers(in *DeleteS2ServersInput) (out *DeleteS2ServersOutput, err error)
	PowerOnS2Servers(in *PowerOnS2ServersInput) (out *PowerOnS2ServersOutput, err error)
	PowerOffS2Servers(in *PowerOffS2ServersInput) (out *PowerOffS2ServersOutput, err error)
	UpdateS2Servers(in *UpdateS2ServersInput) (out *UpdateS2ServersOutput, err error)
	ChangeS2ServerVxnet(in *ChangeS2ServerVxnetInput) (out *ChangeS2ServerVxnetOutput, err error)
	CreateS2SharedTarget(in *CreateS2SharedTargetInput) (out *CreateS2SharedTargetOutput, err error)
	DescribeS2SharedTargets(in *DescribeS2SharedTargetsInput) (out *DescribeS2SharedTargetsOutput, err error)
	DeleteS2SharedTargets(in *DeleteS2SharedTargetsInput) (out *DeleteS2SharedTargetsOutput, err error)
	EnableS2SharedTargets(in *EnableS2SharedTargetsInput) (out *EnableS2SharedTargetsOutput, err error)
	DisableS2SharedTargets(in *DisableS2SharedTargetsInput) (out *DisableS2SharedTargetsOutput, err error)
	ModifyS2SharedTargetAttributes(in *ModifyS2SharedTargetAttributesInput) (out *ModifyS2SharedTargetAttributesOutput, err error)
	AttachToS2SharedTarget(in *AttachToS2SharedTargetInput) (out *AttachToS2SharedTargetOutput, err error)
	DetachFromS2SharedTarget(in *DetachFromS2SharedTargetInput) (out *DetachFromS2SharedTargetOutput, err error)
	DescribeS2DefaultParameters(in *DescribeS2DefaultParametersInput) (out *DescribeS2DefaultParametersOutput, err error)
	CreateS2Group(in *CreateS2GroupInput) (out *CreateS2GroupOutput, err error)
	DescribeS2Groups(in *DescribeS2GroupsInput) (out *DescribeS2GroupsOutput, err error)
	ModifyS2Group(in *ModifyS2GroupInput) (out *ModifyS2GroupOutput, err error)
	DeleteS2Groups(in *DeleteS2GroupsInput) (out *DeleteS2GroupsOutput, err error)
	CreateS2Account(in *CreateS2AccountInput) (out *CreateS2AccountOutput, err error)
	DescribeS2Accounts(in *DescribeS2AccountsInput) (out *DescribeS2AccountsOutput, err error)
	ModifyS2Account(in *ModifyS2AccountInput) (out *ModifyS2AccountOutput, err error)
	DeleteS2Accounts(in *DeleteS2AccountsInput) (out *DeleteS2AccountsOutput, err error)
	AssociateS2AccountGroup(in *AssociateS2AccountGroupInput) (out *AssociateS2AccountGroupOutput, err error)
	DissociateS2AccountGroup(in *DissociateS2AccountGroupInput) (out *DissociateS2AccountGroupOutput, err error)
}

// See https://docs.qingcloud.com/api/s2/index.html
type S2Service struct {
	Config           *config.Config
	Properties       *S2ServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/s2/index.html
func NewS2Service(conf *config.Config, zone string) (p *S2Service) {
	return &S2Service{
		Config:     conf,
		Properties: &S2ServiceProperties{Zone: proto.String(zone)},
	}
}

// See https://docs.qingcloud.com/api/s2/index.html
func (s *QingCloudService) S2(zone string) (*S2Service, error) {
	properties := &S2ServiceProperties{
		Zone: proto.String(zone),
	}

	return &S2Service{Config: s.Config, Properties: properties}, nil
}

func (p *S2Service) CreateS2Server(in *CreateS2ServerInput) (out *CreateS2ServerOutput, err error) {
	if in == nil {
		in = &CreateS2ServerInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2Server",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateS2ServerOutput{}
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

func (p *S2Service) DescribeS2Servers(in *DescribeS2ServersInput) (out *DescribeS2ServersOutput, err error) {
	if in == nil {
		in = &DescribeS2ServersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2Servers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeS2ServersOutput{}
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

func (p *S2Service) ModifyS2Server(in *ModifyS2ServerInput) (out *ModifyS2ServerOutput, err error) {
	if in == nil {
		in = &ModifyS2ServerInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2Server",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyS2ServerOutput{}
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

func (p *S2Service) ResizeS2Servers(in *ResizeS2ServersInput) (out *ResizeS2ServersOutput, err error) {
	if in == nil {
		in = &ResizeS2ServersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeS2Servers",
		RequestMethod: "GET", // GET or POST
	}

	x := &ResizeS2ServersOutput{}
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

func (p *S2Service) DeleteS2Servers(in *DeleteS2ServersInput) (out *DeleteS2ServersOutput, err error) {
	if in == nil {
		in = &DeleteS2ServersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2Servers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteS2ServersOutput{}
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

func (p *S2Service) PowerOnS2Servers(in *PowerOnS2ServersInput) (out *PowerOnS2ServersOutput, err error) {
	if in == nil {
		in = &PowerOnS2ServersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOnS2Servers",
		RequestMethod: "GET", // GET or POST
	}

	x := &PowerOnS2ServersOutput{}
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

func (p *S2Service) PowerOffS2Servers(in *PowerOffS2ServersInput) (out *PowerOffS2ServersOutput, err error) {
	if in == nil {
		in = &PowerOffS2ServersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOffS2Servers",
		RequestMethod: "GET", // GET or POST
	}

	x := &PowerOffS2ServersOutput{}
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

func (p *S2Service) UpdateS2Servers(in *UpdateS2ServersInput) (out *UpdateS2ServersOutput, err error) {
	if in == nil {
		in = &UpdateS2ServersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateS2Servers",
		RequestMethod: "GET", // GET or POST
	}

	x := &UpdateS2ServersOutput{}
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

func (p *S2Service) ChangeS2ServerVxnet(in *ChangeS2ServerVxnetInput) (out *ChangeS2ServerVxnetOutput, err error) {
	if in == nil {
		in = &ChangeS2ServerVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeS2ServerVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &ChangeS2ServerVxnetOutput{}
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

func (p *S2Service) CreateS2SharedTarget(in *CreateS2SharedTargetInput) (out *CreateS2SharedTargetOutput, err error) {
	if in == nil {
		in = &CreateS2SharedTargetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2SharedTarget",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateS2SharedTargetOutput{}
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

func (p *S2Service) DescribeS2SharedTargets(in *DescribeS2SharedTargetsInput) (out *DescribeS2SharedTargetsOutput, err error) {
	if in == nil {
		in = &DescribeS2SharedTargetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2SharedTargets",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeS2SharedTargetsOutput{}
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

func (p *S2Service) DeleteS2SharedTargets(in *DeleteS2SharedTargetsInput) (out *DeleteS2SharedTargetsOutput, err error) {
	if in == nil {
		in = &DeleteS2SharedTargetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2SharedTargets",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteS2SharedTargetsOutput{}
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

func (p *S2Service) EnableS2SharedTargets(in *EnableS2SharedTargetsInput) (out *EnableS2SharedTargetsOutput, err error) {
	if in == nil {
		in = &EnableS2SharedTargetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "EnableS2SharedTargets",
		RequestMethod: "GET", // GET or POST
	}

	x := &EnableS2SharedTargetsOutput{}
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

func (p *S2Service) DisableS2SharedTargets(in *DisableS2SharedTargetsInput) (out *DisableS2SharedTargetsOutput, err error) {
	if in == nil {
		in = &DisableS2SharedTargetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DisableS2SharedTargets",
		RequestMethod: "GET", // GET or POST
	}

	x := &DisableS2SharedTargetsOutput{}
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

func (p *S2Service) ModifyS2SharedTargetAttributes(in *ModifyS2SharedTargetAttributesInput) (out *ModifyS2SharedTargetAttributesOutput, err error) {
	if in == nil {
		in = &ModifyS2SharedTargetAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2SharedTargetAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyS2SharedTargetAttributesOutput{}
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

func (p *S2Service) AttachToS2SharedTarget(in *AttachToS2SharedTargetInput) (out *AttachToS2SharedTargetOutput, err error) {
	if in == nil {
		in = &AttachToS2SharedTargetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachToS2SharedTarget",
		RequestMethod: "GET", // GET or POST
	}

	x := &AttachToS2SharedTargetOutput{}
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

func (p *S2Service) DetachFromS2SharedTarget(in *DetachFromS2SharedTargetInput) (out *DetachFromS2SharedTargetOutput, err error) {
	if in == nil {
		in = &DetachFromS2SharedTargetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachFromS2SharedTarget",
		RequestMethod: "GET", // GET or POST
	}

	x := &DetachFromS2SharedTargetOutput{}
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

func (p *S2Service) DescribeS2DefaultParameters(in *DescribeS2DefaultParametersInput) (out *DescribeS2DefaultParametersOutput, err error) {
	if in == nil {
		in = &DescribeS2DefaultParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2DefaultParameters",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeS2DefaultParametersOutput{}
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

func (p *S2Service) CreateS2Group(in *CreateS2GroupInput) (out *CreateS2GroupOutput, err error) {
	if in == nil {
		in = &CreateS2GroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2Group",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateS2GroupOutput{}
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

func (p *S2Service) DescribeS2Groups(in *DescribeS2GroupsInput) (out *DescribeS2GroupsOutput, err error) {
	if in == nil {
		in = &DescribeS2GroupsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2Groups",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeS2GroupsOutput{}
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

func (p *S2Service) ModifyS2Group(in *ModifyS2GroupInput) (out *ModifyS2GroupOutput, err error) {
	if in == nil {
		in = &ModifyS2GroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2Group",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyS2GroupOutput{}
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

func (p *S2Service) DeleteS2Groups(in *DeleteS2GroupsInput) (out *DeleteS2GroupsOutput, err error) {
	if in == nil {
		in = &DeleteS2GroupsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2Groups",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteS2GroupsOutput{}
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

func (p *S2Service) CreateS2Account(in *CreateS2AccountInput) (out *CreateS2AccountOutput, err error) {
	if in == nil {
		in = &CreateS2AccountInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateS2Account",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateS2AccountOutput{}
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

func (p *S2Service) DescribeS2Accounts(in *DescribeS2AccountsInput) (out *DescribeS2AccountsOutput, err error) {
	if in == nil {
		in = &DescribeS2AccountsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeS2Accounts",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeS2AccountsOutput{}
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

func (p *S2Service) ModifyS2Account(in *ModifyS2AccountInput) (out *ModifyS2AccountOutput, err error) {
	if in == nil {
		in = &ModifyS2AccountInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyS2Account",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyS2AccountOutput{}
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

func (p *S2Service) DeleteS2Accounts(in *DeleteS2AccountsInput) (out *DeleteS2AccountsOutput, err error) {
	if in == nil {
		in = &DeleteS2AccountsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteS2Accounts",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteS2AccountsOutput{}
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

func (p *S2Service) AssociateS2AccountGroup(in *AssociateS2AccountGroupInput) (out *AssociateS2AccountGroupOutput, err error) {
	if in == nil {
		in = &AssociateS2AccountGroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateS2AccountGroup",
		RequestMethod: "GET", // GET or POST
	}

	x := &AssociateS2AccountGroupOutput{}
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

func (p *S2Service) DissociateS2AccountGroup(in *DissociateS2AccountGroupInput) (out *DissociateS2AccountGroupOutput, err error) {
	if in == nil {
		in = &DissociateS2AccountGroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateS2AccountGroup",
		RequestMethod: "GET", // GET or POST
	}

	x := &DissociateS2AccountGroupOutput{}
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

func (p *S2ServiceProperties) Validate() error {
	return nil
}

func (p *CreateS2ServerInput) Validate() error {
	return nil
}

func (p *CreateS2ServerOutput) Validate() error {
	return nil
}

func (p *DescribeS2ServersInput) Validate() error {
	return nil
}

func (p *DescribeS2ServersOutput) Validate() error {
	return nil
}

func (p *ModifyS2ServerInput) Validate() error {
	return nil
}

func (p *ModifyS2ServerOutput) Validate() error {
	return nil
}

func (p *ResizeS2ServersInput) Validate() error {
	return nil
}

func (p *ResizeS2ServersOutput) Validate() error {
	return nil
}

func (p *DeleteS2ServersInput) Validate() error {
	return nil
}

func (p *DeleteS2ServersOutput) Validate() error {
	return nil
}

func (p *PowerOnS2ServersInput) Validate() error {
	return nil
}

func (p *PowerOnS2ServersOutput) Validate() error {
	return nil
}

func (p *PowerOffS2ServersInput) Validate() error {
	return nil
}

func (p *PowerOffS2ServersOutput) Validate() error {
	return nil
}

func (p *UpdateS2ServersInput) Validate() error {
	return nil
}

func (p *UpdateS2ServersOutput) Validate() error {
	return nil
}

func (p *ChangeS2ServerVxnetInput) Validate() error {
	return nil
}

func (p *ChangeS2ServerVxnetOutput) Validate() error {
	return nil
}

func (p *CreateS2SharedTargetInput) Validate() error {
	return nil
}

func (p *CreateS2SharedTargetOutput) Validate() error {
	return nil
}

func (p *DescribeS2SharedTargetsInput) Validate() error {
	return nil
}

func (p *DescribeS2SharedTargetsOutput) Validate() error {
	return nil
}

func (p *DeleteS2SharedTargetsInput) Validate() error {
	return nil
}

func (p *DeleteS2SharedTargetsOutput) Validate() error {
	return nil
}

func (p *EnableS2SharedTargetsInput) Validate() error {
	return nil
}

func (p *EnableS2SharedTargetsOutput) Validate() error {
	return nil
}

func (p *DisableS2SharedTargetsInput) Validate() error {
	return nil
}

func (p *DisableS2SharedTargetsOutput) Validate() error {
	return nil
}

func (p *ModifyS2SharedTargetAttributesInput) Validate() error {
	return nil
}

func (p *ModifyS2SharedTargetAttributesOutput) Validate() error {
	return nil
}

func (p *AttachToS2SharedTargetInput) Validate() error {
	return nil
}

func (p *AttachToS2SharedTargetOutput) Validate() error {
	return nil
}

func (p *DetachFromS2SharedTargetInput) Validate() error {
	return nil
}

func (p *DetachFromS2SharedTargetOutput) Validate() error {
	return nil
}

func (p *DescribeS2DefaultParametersInput) Validate() error {
	return nil
}

func (p *DescribeS2DefaultParametersOutput) Validate() error {
	return nil
}

func (p *CreateS2GroupInput) Validate() error {
	return nil
}

func (p *CreateS2GroupOutput) Validate() error {
	return nil
}

func (p *DescribeS2GroupsInput) Validate() error {
	return nil
}

func (p *DescribeS2GroupsOutput) Validate() error {
	return nil
}

func (p *ModifyS2GroupInput) Validate() error {
	return nil
}

func (p *ModifyS2GroupOutput) Validate() error {
	return nil
}

func (p *DeleteS2GroupsInput) Validate() error {
	return nil
}

func (p *DeleteS2GroupsOutput) Validate() error {
	return nil
}

func (p *CreateS2AccountInput) Validate() error {
	return nil
}

func (p *CreateS2AccountOutput) Validate() error {
	return nil
}

func (p *DescribeS2AccountsInput) Validate() error {
	return nil
}

func (p *DescribeS2AccountsOutput) Validate() error {
	return nil
}

func (p *ModifyS2AccountInput) Validate() error {
	return nil
}

func (p *ModifyS2AccountOutput) Validate() error {
	return nil
}

func (p *DeleteS2AccountsInput) Validate() error {
	return nil
}

func (p *DeleteS2AccountsOutput) Validate() error {
	return nil
}

func (p *AssociateS2AccountGroupInput) Validate() error {
	return nil
}

func (p *AssociateS2AccountGroupOutput) Validate() error {
	return nil
}

func (p *DissociateS2AccountGroupInput) Validate() error {
	return nil
}

func (p *DissociateS2AccountGroupOutput) Validate() error {
	return nil
}
