// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: zone.proto

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

// See https://docs.qingcloud.com/api/zone/index.html
type ZoneServiceInterface interface {
	DescribeZones(in *DescribeZonesInput) (out *DescribeZonesOutput, err error)
}

// See https://docs.qingcloud.com/api/zone/index.html
type ZoneService struct {
	Config           *config.Config
	Properties       *ZoneServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/zone/index.html
func NewZoneService(conf *config.Config, zone string) (p *ZoneService) {
	return &ZoneService{
		Config:     conf,
		Properties: &ZoneServiceProperties{Zone: proto.String(zone)},
	}
}

// See https://docs.qingcloud.com/api/zone/index.html
func (s *QingCloudService) Zone(zone string) (*ZoneService, error) {
	properties := &ZoneServiceProperties{
		Zone: proto.String(zone),
	}

	return &ZoneService{Config: s.Config, Properties: properties}, nil
}

func (p *ZoneService) DescribeZones(in *DescribeZonesInput) (out *DescribeZonesOutput, err error) {
	if in == nil {
		in = &DescribeZonesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeZones",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeZonesOutput{}
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

func (p *ZoneServiceProperties) Validate() error {
	return nil
}

func (p *DescribeZonesInput) Validate() error {
	return nil
}

func (p *DescribeZonesOutput) Validate() error {
	return nil
}
