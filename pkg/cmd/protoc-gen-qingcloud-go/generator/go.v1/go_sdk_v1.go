// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package go_sdk_v1

import (
	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud-go"
)

func init() {
	plugin.RegisterGenerater(new(pkgGenerator))
}

type pkgGenerator struct{}

func (p pkgGenerator) Name() string        { return "go.v1" }
func (p pkgGenerator) FileNameExt() string { return ".pb.qingcloud.go" }

func (p pkgGenerator) GenFileHeaderCode(spec *spec_metadata.FileSpec) (code string, err error) {
	panic("")
}
func (p pkgGenerator) GenFileTailCode(spec *spec_metadata.FileSpec) (code string, err error) {
	return "", nil
}

func (p pkgGenerator) GenServiceCode(spec *spec_metadata.ServiceSpec) (code string, err error) {
	panic("")
}

const FileHeaderTemplate = `
// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: {{.GetFileName}}

package {{.GetPackageName}}

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

`

const ServiceTemplate = `
{{$service := .}}

type {{.GetServiceName}}ServiceInterface interface {
	{{- range $_, $m := .GetMethodList}}
		{{$m.GetMethodName}}(in *{{$m.GetInputTypeName}}) (out *{{$m.GetOutputTypeName}}, err error)
	{{- end}}
}

type {{.GetServiceName}}Service struct {
	Config           *config.Config
	Properties       *{{.GetServiceName}}ServiceProperties
	LastResponseBody string
}

func New{{.GetServiceName}}Service(conf *config.Config, zone string) (p *{{.GetServiceName}}Service) {
	return &{{.GetServiceName}}Service{
		Config:     conf,
		Properties: &{{.GetServiceName}}ServiceProperties{ Zone: proto.String(zone) },
	}
}

{{range $_, $m := .GetMethodList}}
func (p *{{$service.GetServiceName}}Service) {{$m.GetMethodName}}(in *{{$m.GetInputTypeName}}) (out *{{$m.OutputTypeName}}, err error) {
	if in == nil {
		in = &{{$m.GetInputTypeName}}{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "{{$m.GetMethodName}}",
		RequestMethod: "{{$m.GetHttpMethod}}",
	}

	x := &{{$m.GetOutputTypeName}}{}
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
{{end}}
`
