// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: cluster.proto

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

type ClusterServiceInterface interface {
	CreateCluster(in *CreateClusterInput) (out *CreateClusterOutput, err error)
	DescribeClusters(in *DescribeClustersInput) (out *DescribeClustersOutput, err error)
	DescribeClusterNodes(in *DescribeClusterNodesInput) (out *DescribeClusterNodesOutput, err error)
	StopClusters(in *StopClustersInput) (out *StopClustersOutput, err error)
	StartClusters(in *StartClustersInput) (out *StartClustersOutput, err error)
	DeleteClusters(in *DeleteClustersInput) (out *DeleteClustersOutput, err error)
	Lease(in *LeaseInput) (out *LeaseOutput, err error)
	AddClusterNodes(in *AddClusterNodesInput) (out *AddClusterNodesOutput, err error)
	DeleteClusterNodes(in *DeleteClusterNodesInput) (out *DeleteClusterNodesOutput, err error)
	ResizeCluster(in *ResizeClusterInput) (out *ResizeClusterOutput, err error)
	ChangeClusterVxnet(in *ChangeClusterVxnetInput) (out *ChangeClusterVxnetOutput, err error)
	SuspendClusters(in *SuspendClustersInput) (out *SuspendClustersOutput, err error)
	UpdateClusterEnvironment(in *UpdateClusterEnvironmentInput) (out *UpdateClusterEnvironmentOutput, err error)
	ModifyClusterAttributes(in *ModifyClusterAttributesInput) (out *ModifyClusterAttributesOutput, err error)
	ModifyClusterNodeAttributes(in *ModifyClusterNodeAttributesInput) (out *ModifyClusterNodeAttributesOutput, err error)
	GetClustersStats(in *GetClustersStatsInput) (out *GetClustersStatsOutput, err error)
	DescribeClusterUsers(in *DescribeClusterUsersInput) (out *DescribeClusterUsersOutput, err error)
	RestartClusterService(in *RestartClusterServiceInput) (out *RestartClusterServiceOutput, err error)
	UpgradeClusters(in *UpgradeClustersInput) (out *UpgradeClustersOutput, err error)
	AuthorizeClustersBrokerToDeveloper(in *AuthorizeClustersBrokerToDeveloperInput) (out *AuthorizeClustersBrokerToDeveloperOutput, err error)
	RevokeClustersBrokerFromDeveloper(in *RevokeClustersBrokerFromDeveloperInput) (out *RevokeClustersBrokerFromDeveloperOutput, err error)
}

type ClusterService struct {
	Config           *config.Config
	Properties       *ClusterServiceProperties
	LastResponseBody string
}

func NewClusterService(conf *config.Config, zone string) (p *ClusterService) {
	return &ClusterService{
		Config:     conf,
		Properties: &ClusterServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *ClusterService) CreateCluster(in *CreateClusterInput) (out *CreateClusterOutput, err error) {
	if in == nil {
		in = &CreateClusterInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCluster",
		RequestMethod: "GET",
	}

	x := &CreateClusterOutput{}
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

func (p *ClusterService) DescribeClusters(in *DescribeClustersInput) (out *DescribeClustersOutput, err error) {
	if in == nil {
		in = &DescribeClustersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeClusters",
		RequestMethod: "GET",
	}

	x := &DescribeClustersOutput{}
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

func (p *ClusterService) DescribeClusterNodes(in *DescribeClusterNodesInput) (out *DescribeClusterNodesOutput, err error) {
	if in == nil {
		in = &DescribeClusterNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeClusterNodes",
		RequestMethod: "GET",
	}

	x := &DescribeClusterNodesOutput{}
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

func (p *ClusterService) StopClusters(in *StopClustersInput) (out *StopClustersOutput, err error) {
	if in == nil {
		in = &StopClustersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopClusters",
		RequestMethod: "GET",
	}

	x := &StopClustersOutput{}
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

func (p *ClusterService) StartClusters(in *StartClustersInput) (out *StartClustersOutput, err error) {
	if in == nil {
		in = &StartClustersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartClusters",
		RequestMethod: "GET",
	}

	x := &StartClustersOutput{}
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

func (p *ClusterService) DeleteClusters(in *DeleteClustersInput) (out *DeleteClustersOutput, err error) {
	if in == nil {
		in = &DeleteClustersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteClusters",
		RequestMethod: "GET",
	}

	x := &DeleteClustersOutput{}
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

func (p *ClusterService) Lease(in *LeaseInput) (out *LeaseOutput, err error) {
	if in == nil {
		in = &LeaseInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "Lease",
		RequestMethod: "GET",
	}

	x := &LeaseOutput{}
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

func (p *ClusterService) AddClusterNodes(in *AddClusterNodesInput) (out *AddClusterNodesOutput, err error) {
	if in == nil {
		in = &AddClusterNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddClusterNodes",
		RequestMethod: "GET",
	}

	x := &AddClusterNodesOutput{}
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

func (p *ClusterService) DeleteClusterNodes(in *DeleteClusterNodesInput) (out *DeleteClusterNodesOutput, err error) {
	if in == nil {
		in = &DeleteClusterNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteClusterNodes",
		RequestMethod: "GET",
	}

	x := &DeleteClusterNodesOutput{}
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

func (p *ClusterService) ResizeCluster(in *ResizeClusterInput) (out *ResizeClusterOutput, err error) {
	if in == nil {
		in = &ResizeClusterInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeCluster",
		RequestMethod: "GET",
	}

	x := &ResizeClusterOutput{}
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

func (p *ClusterService) ChangeClusterVxnet(in *ChangeClusterVxnetInput) (out *ChangeClusterVxnetOutput, err error) {
	if in == nil {
		in = &ChangeClusterVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeClusterVxnet",
		RequestMethod: "GET",
	}

	x := &ChangeClusterVxnetOutput{}
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

func (p *ClusterService) SuspendClusters(in *SuspendClustersInput) (out *SuspendClustersOutput, err error) {
	if in == nil {
		in = &SuspendClustersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "SuspendClusters",
		RequestMethod: "GET",
	}

	x := &SuspendClustersOutput{}
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

func (p *ClusterService) UpdateClusterEnvironment(in *UpdateClusterEnvironmentInput) (out *UpdateClusterEnvironmentOutput, err error) {
	if in == nil {
		in = &UpdateClusterEnvironmentInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateClusterEnvironment",
		RequestMethod: "GET",
	}

	x := &UpdateClusterEnvironmentOutput{}
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

func (p *ClusterService) ModifyClusterAttributes(in *ModifyClusterAttributesInput) (out *ModifyClusterAttributesOutput, err error) {
	if in == nil {
		in = &ModifyClusterAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyClusterAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyClusterAttributesOutput{}
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

func (p *ClusterService) ModifyClusterNodeAttributes(in *ModifyClusterNodeAttributesInput) (out *ModifyClusterNodeAttributesOutput, err error) {
	if in == nil {
		in = &ModifyClusterNodeAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyClusterNodeAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyClusterNodeAttributesOutput{}
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

func (p *ClusterService) GetClustersStats(in *GetClustersStatsInput) (out *GetClustersStatsOutput, err error) {
	if in == nil {
		in = &GetClustersStatsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetClustersStats",
		RequestMethod: "GET",
	}

	x := &GetClustersStatsOutput{}
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

func (p *ClusterService) DescribeClusterUsers(in *DescribeClusterUsersInput) (out *DescribeClusterUsersOutput, err error) {
	if in == nil {
		in = &DescribeClusterUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeClusterUsers",
		RequestMethod: "GET",
	}

	x := &DescribeClusterUsersOutput{}
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

func (p *ClusterService) RestartClusterService(in *RestartClusterServiceInput) (out *RestartClusterServiceOutput, err error) {
	if in == nil {
		in = &RestartClusterServiceInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartClusterService",
		RequestMethod: "GET",
	}

	x := &RestartClusterServiceOutput{}
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

func (p *ClusterService) UpgradeClusters(in *UpgradeClustersInput) (out *UpgradeClustersOutput, err error) {
	if in == nil {
		in = &UpgradeClustersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpgradeClusters",
		RequestMethod: "GET",
	}

	x := &UpgradeClustersOutput{}
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

func (p *ClusterService) AuthorizeClustersBrokerToDeveloper(in *AuthorizeClustersBrokerToDeveloperInput) (out *AuthorizeClustersBrokerToDeveloperOutput, err error) {
	if in == nil {
		in = &AuthorizeClustersBrokerToDeveloperInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AuthorizeClustersBrokerToDeveloper",
		RequestMethod: "GET",
	}

	x := &AuthorizeClustersBrokerToDeveloperOutput{}
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

func (p *ClusterService) RevokeClustersBrokerFromDeveloper(in *RevokeClustersBrokerFromDeveloperInput) (out *RevokeClustersBrokerFromDeveloperOutput, err error) {
	if in == nil {
		in = &RevokeClustersBrokerFromDeveloperInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RevokeClustersBrokerFromDeveloper",
		RequestMethod: "GET",
	}

	x := &RevokeClustersBrokerFromDeveloperOutput{}
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
