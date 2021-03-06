// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: nic.proto

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

type NicServiceInterface interface {
	CreateNics(in *CreateNicsInput) (out *CreateNicsOutput, err error)
	DescribeNics(in *DescribeNicsInput) (out *DescribeNicsOutput, err error)
	AttachNics(in *AttachNicsInput) (out *AttachNicsOutput, err error)
	DetachNics(in *DetachNicsInput) (out *DetachNicsOutput, err error)
	ModifyNicAttributes(in *ModifyNicAttributesInput) (out *ModifyNicAttributesOutput, err error)
	DeleteNics(in *DeleteNicsInput) (out *DeleteNicsOutput, err error)
}

type NicService struct {
	Config           *config.Config
	Properties       *NicServiceProperties
	LastResponseBody string
}

func NewNicService(conf *config.Config, zone string) (p *NicService) {
	return &NicService{
		Config:     conf,
		Properties: &NicServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Nic(zone string) (*NicService, error) {
	properties := &NicServiceProperties{
		Zone: proto.String(zone),
	}

	return &NicService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *NicService) CreateNics(in *CreateNicsInput) (out *CreateNicsOutput, err error) {
	if in == nil {
		in = &CreateNicsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateNics",
		RequestMethod: "GET",
	}

	x := &CreateNicsOutput{}
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

func (p *NicService) DescribeNics(in *DescribeNicsInput) (out *DescribeNicsOutput, err error) {
	if in == nil {
		in = &DescribeNicsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeNics",
		RequestMethod: "GET",
	}

	x := &DescribeNicsOutput{}
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

func (p *NicService) AttachNics(in *AttachNicsInput) (out *AttachNicsOutput, err error) {
	if in == nil {
		in = &AttachNicsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachNics",
		RequestMethod: "GET",
	}

	x := &AttachNicsOutput{}
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

func (p *NicService) DetachNics(in *DetachNicsInput) (out *DetachNicsOutput, err error) {
	if in == nil {
		in = &DetachNicsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachNics",
		RequestMethod: "GET",
	}

	x := &DetachNicsOutput{}
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

func (p *NicService) ModifyNicAttributes(in *ModifyNicAttributesInput) (out *ModifyNicAttributesOutput, err error) {
	if in == nil {
		in = &ModifyNicAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyNicAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyNicAttributesOutput{}
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

func (p *NicService) DeleteNics(in *DeleteNicsInput) (out *DeleteNicsOutput, err error) {
	if in == nil {
		in = &DeleteNicsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteNics",
		RequestMethod: "GET",
	}

	x := &DeleteNicsOutput{}
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
