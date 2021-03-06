// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: router.proto

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

type RouterServiceInterface interface {
	DescribeRouters(in *DescribeRoutersInput) (out *DescribeRoutersOutput, err error)
	CreateRouters(in *CreateRoutersInput) (out *CreateRoutersOutput, err error)
	DeleteRouters(in *DeleteRoutersInput) (out *DeleteRoutersOutput, err error)
	UpdateRouters(in *UpdateRoutersInput) (out *UpdateRoutersOutput, err error)
	PowerOffRouters(in *PowerOffRoutersInput) (out *PowerOffRoutersOutput, err error)
	PowerOnRouters(in *PowerOnRoutersInput) (out *PowerOnRoutersOutput, err error)
	JoinRouter(in *JoinRouterInput) (out *JoinRouterOutput, err error)
	LeaveRouter(in *LeaveRouterInput) (out *LeaveRouterOutput, err error)
	ModifyRouterAttributes(in *ModifyRouterAttributesInput) (out *ModifyRouterAttributesOutput, err error)
	DescribeRouterStatics(in *DescribeRouterStaticsInput) (out *DescribeRouterStaticsOutput, err error)
	AddRouterStatics(in *AddRouterStaticsInput) (out *AddRouterStaticsOutput, err error)
	ModifyRouterStaticAttributes(in *ModifyRouterStaticAttributesInput) (out *ModifyRouterStaticAttributesOutput, err error)
	DeleteRouterStatics(in *DeleteRouterStaticsInput) (out *DeleteRouterStaticsOutput, err error)
	CopyRouterStatics(in *CopyRouterStaticsInput) (out *CopyRouterStaticsOutput, err error)
	DescribeRouterVxnets(in *DescribeRouterVxnetsInput) (out *DescribeRouterVxnetsOutput, err error)
	AddRouterStaticEntries(in *AddRouterStaticEntriesInput) (out *AddRouterStaticEntriesOutput, err error)
	DeleteRouterStaticEntries(in *DeleteRouterStaticEntriesInput) (out *DeleteRouterStaticEntriesOutput, err error)
	ModifyRouterStaticEntryAttributes(in *ModifyRouterStaticEntryAttributesInput) (out *ModifyRouterStaticEntryAttributesOutput, err error)
	DescribeRouterStaticEntries(in *DescribeRouterStaticEntriesInput) (out *DescribeRouterStaticEntriesOutput, err error)
}

type RouterService struct {
	Config           *config.Config
	Properties       *RouterServiceProperties
	LastResponseBody string
}

func NewRouterService(conf *config.Config, zone string) (p *RouterService) {
	return &RouterService{
		Config:     conf,
		Properties: &RouterServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Router(zone string) (*RouterService, error) {
	properties := &RouterServiceProperties{
		Zone: proto.String(zone),
	}

	return &RouterService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *RouterService) DescribeRouters(in *DescribeRoutersInput) (out *DescribeRoutersOutput, err error) {
	if in == nil {
		in = &DescribeRoutersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouters",
		RequestMethod: "GET",
	}

	x := &DescribeRoutersOutput{}
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

func (p *RouterService) CreateRouters(in *CreateRoutersInput) (out *CreateRoutersOutput, err error) {
	if in == nil {
		in = &CreateRoutersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRouters",
		RequestMethod: "GET",
	}

	x := &CreateRoutersOutput{}
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

func (p *RouterService) DeleteRouters(in *DeleteRoutersInput) (out *DeleteRoutersOutput, err error) {
	if in == nil {
		in = &DeleteRoutersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRouters",
		RequestMethod: "GET",
	}

	x := &DeleteRoutersOutput{}
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

func (p *RouterService) UpdateRouters(in *UpdateRoutersInput) (out *UpdateRoutersOutput, err error) {
	if in == nil {
		in = &UpdateRoutersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateRouters",
		RequestMethod: "GET",
	}

	x := &UpdateRoutersOutput{}
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

func (p *RouterService) PowerOffRouters(in *PowerOffRoutersInput) (out *PowerOffRoutersOutput, err error) {
	if in == nil {
		in = &PowerOffRoutersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOffRouters",
		RequestMethod: "GET",
	}

	x := &PowerOffRoutersOutput{}
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

func (p *RouterService) PowerOnRouters(in *PowerOnRoutersInput) (out *PowerOnRoutersOutput, err error) {
	if in == nil {
		in = &PowerOnRoutersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PowerOnRouters",
		RequestMethod: "GET",
	}

	x := &PowerOnRoutersOutput{}
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

func (p *RouterService) JoinRouter(in *JoinRouterInput) (out *JoinRouterOutput, err error) {
	if in == nil {
		in = &JoinRouterInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "JoinRouter",
		RequestMethod: "GET",
	}

	x := &JoinRouterOutput{}
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

func (p *RouterService) LeaveRouter(in *LeaveRouterInput) (out *LeaveRouterOutput, err error) {
	if in == nil {
		in = &LeaveRouterInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "LeaveRouter",
		RequestMethod: "GET",
	}

	x := &LeaveRouterOutput{}
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

func (p *RouterService) ModifyRouterAttributes(in *ModifyRouterAttributesInput) (out *ModifyRouterAttributesOutput, err error) {
	if in == nil {
		in = &ModifyRouterAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRouterAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyRouterAttributesOutput{}
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

func (p *RouterService) DescribeRouterStatics(in *DescribeRouterStaticsInput) (out *DescribeRouterStaticsOutput, err error) {
	if in == nil {
		in = &DescribeRouterStaticsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouterStatics",
		RequestMethod: "GET",
	}

	x := &DescribeRouterStaticsOutput{}
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

func (p *RouterService) AddRouterStatics(in *AddRouterStaticsInput) (out *AddRouterStaticsOutput, err error) {
	if in == nil {
		in = &AddRouterStaticsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddRouterStatics",
		RequestMethod: "GET",
	}

	x := &AddRouterStaticsOutput{}
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

func (p *RouterService) ModifyRouterStaticAttributes(in *ModifyRouterStaticAttributesInput) (out *ModifyRouterStaticAttributesOutput, err error) {
	if in == nil {
		in = &ModifyRouterStaticAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRouterStaticAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyRouterStaticAttributesOutput{}
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

func (p *RouterService) DeleteRouterStatics(in *DeleteRouterStaticsInput) (out *DeleteRouterStaticsOutput, err error) {
	if in == nil {
		in = &DeleteRouterStaticsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRouterStatics",
		RequestMethod: "GET",
	}

	x := &DeleteRouterStaticsOutput{}
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

func (p *RouterService) CopyRouterStatics(in *CopyRouterStaticsInput) (out *CopyRouterStaticsOutput, err error) {
	if in == nil {
		in = &CopyRouterStaticsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CopyRouterStatics",
		RequestMethod: "GET",
	}

	x := &CopyRouterStaticsOutput{}
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

func (p *RouterService) DescribeRouterVxnets(in *DescribeRouterVxnetsInput) (out *DescribeRouterVxnetsOutput, err error) {
	if in == nil {
		in = &DescribeRouterVxnetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouterVxnets",
		RequestMethod: "GET",
	}

	x := &DescribeRouterVxnetsOutput{}
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

func (p *RouterService) AddRouterStaticEntries(in *AddRouterStaticEntriesInput) (out *AddRouterStaticEntriesOutput, err error) {
	if in == nil {
		in = &AddRouterStaticEntriesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddRouterStaticEntries",
		RequestMethod: "GET",
	}

	x := &AddRouterStaticEntriesOutput{}
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

func (p *RouterService) DeleteRouterStaticEntries(in *DeleteRouterStaticEntriesInput) (out *DeleteRouterStaticEntriesOutput, err error) {
	if in == nil {
		in = &DeleteRouterStaticEntriesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRouterStaticEntries",
		RequestMethod: "GET",
	}

	x := &DeleteRouterStaticEntriesOutput{}
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

func (p *RouterService) ModifyRouterStaticEntryAttributes(in *ModifyRouterStaticEntryAttributesInput) (out *ModifyRouterStaticEntryAttributesOutput, err error) {
	if in == nil {
		in = &ModifyRouterStaticEntryAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRouterStaticEntryAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyRouterStaticEntryAttributesOutput{}
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

func (p *RouterService) DescribeRouterStaticEntries(in *DescribeRouterStaticEntriesInput) (out *DescribeRouterStaticEntriesOutput, err error) {
	if in == nil {
		in = &DescribeRouterStaticEntriesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRouterStaticEntries",
		RequestMethod: "GET",
	}

	x := &DescribeRouterStaticEntriesOutput{}
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
