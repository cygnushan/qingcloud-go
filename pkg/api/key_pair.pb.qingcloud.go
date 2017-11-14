// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: key_pair.proto

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

type KeyPairServiceInterface interface {
	DescribeKeyPairs(in *DescribeKeyPairsInput) (out *DescribeKeyPairsOutput, err error)
	CreateKeyPair(in *CreateKeyPairInput) (out *CreateKeyPairOutput, err error)
	DeleteKeyPairs(in *DeleteKeyPairsInput) (out *DeleteKeyPairsOutput, err error)
	AttachKeyPairs(in *AttachKeyPairsInput) (out *AttachKeyPairsOutput, err error)
	DetachKeyPairs(in *DetachKeyPairsInput) (out *DetachKeyPairsOutput, err error)
	ModifyKeyPairAttributes(in *ModifyKeyPairAttributesInput) (out *ModifyKeyPairAttributesOutput, err error)
}

type KeyPairService struct {
	Config           *config.Config
	Properties       *KeyPairServiceProperties
	LastResponseBody string
}

func NewKeyPairService(conf *config.Config, zone string) (p *KeyPairService) {
	return &KeyPairService{
		Config:     conf,
		Properties: &KeyPairServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *KeyPairService) DescribeKeyPairs(in *DescribeKeyPairsInput) (out *DescribeKeyPairsOutput, err error) {
	if in == nil {
		in = &DescribeKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeKeyPairs",
		RequestMethod: "GET",
	}

	x := &DescribeKeyPairsOutput{}
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

func (p *KeyPairService) CreateKeyPair(in *CreateKeyPairInput) (out *CreateKeyPairOutput, err error) {
	if in == nil {
		in = &CreateKeyPairInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateKeyPair",
		RequestMethod: "GET",
	}

	x := &CreateKeyPairOutput{}
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

func (p *KeyPairService) DeleteKeyPairs(in *DeleteKeyPairsInput) (out *DeleteKeyPairsOutput, err error) {
	if in == nil {
		in = &DeleteKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteKeyPairs",
		RequestMethod: "GET",
	}

	x := &DeleteKeyPairsOutput{}
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

func (p *KeyPairService) AttachKeyPairs(in *AttachKeyPairsInput) (out *AttachKeyPairsOutput, err error) {
	if in == nil {
		in = &AttachKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachKeyPairs",
		RequestMethod: "GET",
	}

	x := &AttachKeyPairsOutput{}
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

func (p *KeyPairService) DetachKeyPairs(in *DetachKeyPairsInput) (out *DetachKeyPairsOutput, err error) {
	if in == nil {
		in = &DetachKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachKeyPairs",
		RequestMethod: "GET",
	}

	x := &DetachKeyPairsOutput{}
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

func (p *KeyPairService) ModifyKeyPairAttributes(in *ModifyKeyPairAttributesInput) (out *ModifyKeyPairAttributesOutput, err error) {
	if in == nil {
		in = &ModifyKeyPairAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyKeyPairAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyKeyPairAttributesOutput{}
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
