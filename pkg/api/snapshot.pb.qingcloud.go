// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: snapshot.proto

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

type SnapshotServiceInterface interface {
	DescribeSnapshots(in *DescribeSnapshotsInput) (out *DescribeSnapshotsOutput, err error)
	CreateSnapshots(in *CreateSnapshotsInput) (out *CreateSnapshotsOutput, err error)
	DeleteSnapshots(in *DeleteSnapshotsInput) (out *DeleteSnapshotsOutput, err error)
	ApplySnapshots(in *ApplySnapshotsInput) (out *ApplySnapshotsOutput, err error)
	ModifySnapshotAttributes(in *ModifySnapshotAttributesInput) (out *ModifySnapshotAttributesOutput, err error)
	CaptureInstanceFromSnapshot(in *CaptureInstanceFromSnapshotInput) (out *CaptureInstanceFromSnapshotOutput, err error)
	CreateVolumeFromSnapshot(in *CreateVolumeFromSnapshotInput) (out *CreateVolumeFromSnapshotOutput, err error)
}

type SnapshotService struct {
	Config           *config.Config
	Properties       *SnapshotServiceProperties
	LastResponseBody string
}

func NewSnapshotService(conf *config.Config, zone string) (p *SnapshotService) {
	return &SnapshotService{
		Config:     conf,
		Properties: &SnapshotServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Snapshot(zone string) (*SnapshotService, error) {
	properties := &SnapshotServiceProperties{
		Zone: proto.String(zone),
	}

	return &SnapshotService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *SnapshotService) DescribeSnapshots(in *DescribeSnapshotsInput) (out *DescribeSnapshotsOutput, err error) {
	if in == nil {
		in = &DescribeSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSnapshots",
		RequestMethod: "GET",
	}

	x := &DescribeSnapshotsOutput{}
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

func (p *SnapshotService) CreateSnapshots(in *CreateSnapshotsInput) (out *CreateSnapshotsOutput, err error) {
	if in == nil {
		in = &CreateSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSnapshots",
		RequestMethod: "GET",
	}

	x := &CreateSnapshotsOutput{}
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

func (p *SnapshotService) DeleteSnapshots(in *DeleteSnapshotsInput) (out *DeleteSnapshotsOutput, err error) {
	if in == nil {
		in = &DeleteSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSnapshots",
		RequestMethod: "GET",
	}

	x := &DeleteSnapshotsOutput{}
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

func (p *SnapshotService) ApplySnapshots(in *ApplySnapshotsInput) (out *ApplySnapshotsOutput, err error) {
	if in == nil {
		in = &ApplySnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplySnapshots",
		RequestMethod: "GET",
	}

	x := &ApplySnapshotsOutput{}
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

func (p *SnapshotService) ModifySnapshotAttributes(in *ModifySnapshotAttributesInput) (out *ModifySnapshotAttributesOutput, err error) {
	if in == nil {
		in = &ModifySnapshotAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySnapshotAttributes",
		RequestMethod: "GET",
	}

	x := &ModifySnapshotAttributesOutput{}
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

func (p *SnapshotService) CaptureInstanceFromSnapshot(in *CaptureInstanceFromSnapshotInput) (out *CaptureInstanceFromSnapshotOutput, err error) {
	if in == nil {
		in = &CaptureInstanceFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CaptureInstanceFromSnapshot",
		RequestMethod: "GET",
	}

	x := &CaptureInstanceFromSnapshotOutput{}
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

func (p *SnapshotService) CreateVolumeFromSnapshot(in *CreateVolumeFromSnapshotInput) (out *CreateVolumeFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateVolumeFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVolumeFromSnapshot",
		RequestMethod: "GET",
	}

	x := &CreateVolumeFromSnapshotOutput{}
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
