// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: tag.proto

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

type TagServiceInterface interface {
	DescribeTags(in *DescribeTagsInput) (out *DescribeTagsOutput, err error)
	CreateTag(in *CreateTagInput) (out *CreateTagOutput, err error)
	DeleteTags(in *DeleteTagsInput) (out *DeleteTagsOutput, err error)
	ModifyTagAttributes(in *ModifyTagAttributesInput) (out *ModifyTagAttributesOutput, err error)
	AttachTags(in *AttachTagsInput) (out *AttachTagsOutput, err error)
	DetachTags(in *DetachTagsInput) (out *DetachTagsOutput, err error)
}

type TagService struct {
	Config           *config.Config
	Properties       *TagServiceProperties
	LastResponseBody string
}

func NewTagService(conf *config.Config, zone string) (p *TagService) {
	return &TagService{
		Config:     conf,
		Properties: &TagServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Tag(zone string) (*TagService, error) {
	properties := &TagServiceProperties{
		Zone: proto.String(zone),
	}

	return &TagService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *TagService) DescribeTags(in *DescribeTagsInput) (out *DescribeTagsOutput, err error) {
	if in == nil {
		in = &DescribeTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeTags",
		RequestMethod: "GET",
	}

	x := &DescribeTagsOutput{}
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

func (p *TagService) CreateTag(in *CreateTagInput) (out *CreateTagOutput, err error) {
	if in == nil {
		in = &CreateTagInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateTag",
		RequestMethod: "GET",
	}

	x := &CreateTagOutput{}
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

func (p *TagService) DeleteTags(in *DeleteTagsInput) (out *DeleteTagsOutput, err error) {
	if in == nil {
		in = &DeleteTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteTags",
		RequestMethod: "GET",
	}

	x := &DeleteTagsOutput{}
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

func (p *TagService) ModifyTagAttributes(in *ModifyTagAttributesInput) (out *ModifyTagAttributesOutput, err error) {
	if in == nil {
		in = &ModifyTagAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyTagAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyTagAttributesOutput{}
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

func (p *TagService) AttachTags(in *AttachTagsInput) (out *AttachTagsOutput, err error) {
	if in == nil {
		in = &AttachTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachTags",
		RequestMethod: "GET",
	}

	x := &AttachTagsOutput{}
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

func (p *TagService) DetachTags(in *DetachTagsInput) (out *DetachTagsOutput, err error) {
	if in == nil {
		in = &DetachTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachTags",
		RequestMethod: "GET",
	}

	x := &DetachTagsOutput{}
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
