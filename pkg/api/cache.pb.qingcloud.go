// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: cache.proto

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

type CacheServiceInterface interface {
	DescribeCaches(in *DescribeCachesInput) (out *DescribeCachesOutput, err error)
	CreateCache(in *CreateCacheInput) (out *CreateCacheOutput, err error)
	StopCaches(in *StopCachesInput) (out *StopCachesOutput, err error)
	StartCaches(in *StartCachesInput) (out *StartCachesOutput, err error)
	RestartCaches(in *RestartCachesInput) (out *RestartCachesOutput, err error)
	DeleteCaches(in *DeleteCachesInput) (out *DeleteCachesOutput, err error)
	ResizeCaches(in *ResizeCachesInput) (out *ResizeCachesOutput, err error)
	UpdateCache(in *UpdateCacheInput) (out *UpdateCacheOutput, err error)
	ChangeCacheVxnet(in *ChangeCacheVxnetInput) (out *ChangeCacheVxnetOutput, err error)
	ModifyCacheAttributes(in *ModifyCacheAttributesInput) (out *ModifyCacheAttributesOutput, err error)
	DescribeCacheNodes(in *DescribeCacheNodesInput) (out *DescribeCacheNodesOutput, err error)
	AddCacheNodes(in *AddCacheNodesInput) (out *AddCacheNodesOutput, err error)
	DeleteCacheNodes(in *DeleteCacheNodesInput) (out *DeleteCacheNodesOutput, err error)
	RestartCacheNodes(in *RestartCacheNodesInput) (out *RestartCacheNodesOutput, err error)
	ModifyCacheNodeAttributes(in *ModifyCacheNodeAttributesInput) (out *ModifyCacheNodeAttributesOutput, err error)
	CreateCacheFromSnapshot(in *CreateCacheFromSnapshotInput) (out *CreateCacheFromSnapshotOutput, err error)
	DescribeCacheParameterGroups(in *DescribeCacheParameterGroupsInput) (out *DescribeCacheParameterGroupsOutput, err error)
	CreateCacheParameterGroup(in *CreateCacheParameterGroupInput) (out *CreateCacheParameterGroupOutput, err error)
	ApplyCacheParameterGroup(in *ApplyCacheParameterGroupInput) (out *ApplyCacheParameterGroupOutput, err error)
	DeleteCacheParameterGroups(in *DeleteCacheParameterGroupsInput) (out *DeleteCacheParameterGroupsOutput, err error)
	ModifyCacheParameterGroupAttributes(in *ModifyCacheParameterGroupAttributesInput) (out *ModifyCacheParameterGroupAttributesOutput, err error)
	DescribeCacheParameters(in *DescribeCacheParametersInput) (out *DescribeCacheParametersOutput, err error)
	UpdateCacheParameters(in *UpdateCacheParametersInput) (out *UpdateCacheParametersOutput, err error)
	ResetCacheParameters(in *ResetCacheParametersInput) (out *ResetCacheParametersOutput, err error)
}

type CacheService struct {
	Config           *config.Config
	Properties       *CacheServiceProperties
	LastResponseBody string
}

func NewCacheService(conf *config.Config, zone string) (p *CacheService) {
	return &CacheService{
		Config:     conf,
		Properties: &CacheServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Cache(zone string) (*CacheService, error) {
	properties := &CacheServiceProperties{
		Zone: proto.String(zone),
	}

	return &CacheService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *CacheService) DescribeCaches(in *DescribeCachesInput) (out *DescribeCachesOutput, err error) {
	if in == nil {
		in = &DescribeCachesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCaches",
		RequestMethod: "GET",
	}

	x := &DescribeCachesOutput{}
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

func (p *CacheService) CreateCache(in *CreateCacheInput) (out *CreateCacheOutput, err error) {
	if in == nil {
		in = &CreateCacheInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCache",
		RequestMethod: "GET",
	}

	x := &CreateCacheOutput{}
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

func (p *CacheService) StopCaches(in *StopCachesInput) (out *StopCachesOutput, err error) {
	if in == nil {
		in = &StopCachesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopCaches",
		RequestMethod: "GET",
	}

	x := &StopCachesOutput{}
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

func (p *CacheService) StartCaches(in *StartCachesInput) (out *StartCachesOutput, err error) {
	if in == nil {
		in = &StartCachesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartCaches",
		RequestMethod: "GET",
	}

	x := &StartCachesOutput{}
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

func (p *CacheService) RestartCaches(in *RestartCachesInput) (out *RestartCachesOutput, err error) {
	if in == nil {
		in = &RestartCachesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartCaches",
		RequestMethod: "GET",
	}

	x := &RestartCachesOutput{}
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

func (p *CacheService) DeleteCaches(in *DeleteCachesInput) (out *DeleteCachesOutput, err error) {
	if in == nil {
		in = &DeleteCachesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteCaches",
		RequestMethod: "GET",
	}

	x := &DeleteCachesOutput{}
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

func (p *CacheService) ResizeCaches(in *ResizeCachesInput) (out *ResizeCachesOutput, err error) {
	if in == nil {
		in = &ResizeCachesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeCaches",
		RequestMethod: "GET",
	}

	x := &ResizeCachesOutput{}
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

func (p *CacheService) UpdateCache(in *UpdateCacheInput) (out *UpdateCacheOutput, err error) {
	if in == nil {
		in = &UpdateCacheInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateCache",
		RequestMethod: "GET",
	}

	x := &UpdateCacheOutput{}
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

func (p *CacheService) ChangeCacheVxnet(in *ChangeCacheVxnetInput) (out *ChangeCacheVxnetOutput, err error) {
	if in == nil {
		in = &ChangeCacheVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeCacheVxnet",
		RequestMethod: "GET",
	}

	x := &ChangeCacheVxnetOutput{}
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

func (p *CacheService) ModifyCacheAttributes(in *ModifyCacheAttributesInput) (out *ModifyCacheAttributesOutput, err error) {
	if in == nil {
		in = &ModifyCacheAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyCacheAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyCacheAttributesOutput{}
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

func (p *CacheService) DescribeCacheNodes(in *DescribeCacheNodesInput) (out *DescribeCacheNodesOutput, err error) {
	if in == nil {
		in = &DescribeCacheNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCacheNodes",
		RequestMethod: "GET",
	}

	x := &DescribeCacheNodesOutput{}
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

func (p *CacheService) AddCacheNodes(in *AddCacheNodesInput) (out *AddCacheNodesOutput, err error) {
	if in == nil {
		in = &AddCacheNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddCacheNodes",
		RequestMethod: "GET",
	}

	x := &AddCacheNodesOutput{}
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

func (p *CacheService) DeleteCacheNodes(in *DeleteCacheNodesInput) (out *DeleteCacheNodesOutput, err error) {
	if in == nil {
		in = &DeleteCacheNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteCacheNodes",
		RequestMethod: "GET",
	}

	x := &DeleteCacheNodesOutput{}
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

func (p *CacheService) RestartCacheNodes(in *RestartCacheNodesInput) (out *RestartCacheNodesOutput, err error) {
	if in == nil {
		in = &RestartCacheNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartCacheNodes",
		RequestMethod: "GET",
	}

	x := &RestartCacheNodesOutput{}
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

func (p *CacheService) ModifyCacheNodeAttributes(in *ModifyCacheNodeAttributesInput) (out *ModifyCacheNodeAttributesOutput, err error) {
	if in == nil {
		in = &ModifyCacheNodeAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyCacheNodeAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyCacheNodeAttributesOutput{}
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

func (p *CacheService) CreateCacheFromSnapshot(in *CreateCacheFromSnapshotInput) (out *CreateCacheFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateCacheFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCacheFromSnapshot",
		RequestMethod: "GET",
	}

	x := &CreateCacheFromSnapshotOutput{}
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

func (p *CacheService) DescribeCacheParameterGroups(in *DescribeCacheParameterGroupsInput) (out *DescribeCacheParameterGroupsOutput, err error) {
	if in == nil {
		in = &DescribeCacheParameterGroupsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCacheParameterGroups",
		RequestMethod: "GET",
	}

	x := &DescribeCacheParameterGroupsOutput{}
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

func (p *CacheService) CreateCacheParameterGroup(in *CreateCacheParameterGroupInput) (out *CreateCacheParameterGroupOutput, err error) {
	if in == nil {
		in = &CreateCacheParameterGroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCacheParameterGroup",
		RequestMethod: "GET",
	}

	x := &CreateCacheParameterGroupOutput{}
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

func (p *CacheService) ApplyCacheParameterGroup(in *ApplyCacheParameterGroupInput) (out *ApplyCacheParameterGroupOutput, err error) {
	if in == nil {
		in = &ApplyCacheParameterGroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyCacheParameterGroup",
		RequestMethod: "GET",
	}

	x := &ApplyCacheParameterGroupOutput{}
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

func (p *CacheService) DeleteCacheParameterGroups(in *DeleteCacheParameterGroupsInput) (out *DeleteCacheParameterGroupsOutput, err error) {
	if in == nil {
		in = &DeleteCacheParameterGroupsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteCacheParameterGroups",
		RequestMethod: "GET",
	}

	x := &DeleteCacheParameterGroupsOutput{}
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

func (p *CacheService) ModifyCacheParameterGroupAttributes(in *ModifyCacheParameterGroupAttributesInput) (out *ModifyCacheParameterGroupAttributesOutput, err error) {
	if in == nil {
		in = &ModifyCacheParameterGroupAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyCacheParameterGroupAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyCacheParameterGroupAttributesOutput{}
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

func (p *CacheService) DescribeCacheParameters(in *DescribeCacheParametersInput) (out *DescribeCacheParametersOutput, err error) {
	if in == nil {
		in = &DescribeCacheParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCacheParameters",
		RequestMethod: "GET",
	}

	x := &DescribeCacheParametersOutput{}
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

func (p *CacheService) UpdateCacheParameters(in *UpdateCacheParametersInput) (out *UpdateCacheParametersOutput, err error) {
	if in == nil {
		in = &UpdateCacheParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateCacheParameters",
		RequestMethod: "GET",
	}

	x := &UpdateCacheParametersOutput{}
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

func (p *CacheService) ResetCacheParameters(in *ResetCacheParametersInput) (out *ResetCacheParametersOutput, err error) {
	if in == nil {
		in = &ResetCacheParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResetCacheParameters",
		RequestMethod: "GET",
	}

	x := &ResetCacheParametersOutput{}
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
