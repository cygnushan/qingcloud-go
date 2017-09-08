// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// Package qingcloud outputs qingcloud stub code.
package qingcloud

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	options "google.golang.org/genproto/googleapis/api/annotations"
)

// qingcloudPlugin produce the Service interface.
type qingcloudPlugin struct {
	*generator.Generator
}

// Name returns the name of the plugin.
func (p *qingcloudPlugin) Name() string { return "qingcloud" }

// Init is called once after data structures are built but before
// code generation begins.
func (p *qingcloudPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

// Generate produces the code generated by the plugin for this file.
func (p *qingcloudPlugin) GenerateImports(file *generator.FileDescriptor) {
	if !p.getGenericServicesOptions(file) {
		return
	}
	if len(file.Service) > 0 {
		p.P(`import "github.com/chai2010/qingcloud-go/config"`)
		p.P(`import "github.com/chai2010/qingcloud-go/request"`)
		p.P(`import "github.com/chai2010/qingcloud-go/request/errors"`)
		p.P(``)
		p.P(`var _ = config.Config{}`)
		p.P(`var _ = request.Request{}`)
		p.P(`var _ = errors.ParameterRequiredError{}`)
	}
}

// Generate generates the Service interface.
// rpc service can't handle other proto message!!!
func (p *qingcloudPlugin) Generate(file *generator.FileDescriptor) {
	if !p.getGenericServicesOptions(file) {
		return
	}
	for _, svc := range file.Service {
		p.genServiceInterface(file, svc)
		p.genServiceServer(file, svc)
		p.genServiceClient(file, svc)
	}
}

func (p *qingcloudPlugin) getGenericServicesOptions(
	file *generator.FileDescriptor,
) bool {
	const env = "qingcloud_services"

	// try command line first
	// protoc --go_out=qingcloud_services=true:. xxx.proto
	if value, ok := p.Generator.Param[env]; ok {
		if value == "1" || strings.ToLower(value) == "true" {
			return true
		}
		if value == "0" || strings.ToLower(value) == "false" {
			return false
		}
	}

	// default is ture
	return true
}

func (p *qingcloudPlugin) genServiceInterface(
	file *generator.FileDescriptor,
	svc *descriptor.ServiceDescriptorProto,
) {
	const serviceInterfaceTmpl = `
type {{.ServiceName}}Interface interface {
	{{.CallMethodList}}
}
`
	const callMethodTmpl = `
{{.MethodName}}(in *{{.ArgsType}}) (out *{{.ReplyType}}, err error)`

	// gen call method list
	var callMethodList string
	for _, m := range svc.Method {
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(callMethodTmpl))
		t.Execute(out, &struct{ ServiceName, MethodName, ArgsType, ReplyType string }{
			ServiceName: generator.CamelCase(svc.GetName()),
			MethodName:  generator.CamelCase(m.GetName()),
			ArgsType:    p.TypeName(p.ObjectNamed(m.GetInputType())),
			ReplyType:   p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
		callMethodList += out.String()

		p.RecordTypeUse(m.GetInputType())
		p.RecordTypeUse(m.GetOutputType())
	}

	// gen all interface code
	{
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(serviceInterfaceTmpl))
		t.Execute(out, &struct{ ServiceName, CallMethodList string }{
			ServiceName:    generator.CamelCase(svc.GetName()),
			CallMethodList: callMethodList,
		})
		p.P(out.String())
	}
}

func (p *qingcloudPlugin) genServiceServer(
	file *generator.FileDescriptor,
	svc *descriptor.ServiceDescriptorProto,
) {
	const serviceHelperFunTmpl = `
{{/* TODO */}}
`
	{
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(serviceHelperFunTmpl))
		t.Execute(out, &struct{ PackageName, ServiceName, ServiceRegisterName string }{
			PackageName:         file.GetPackage(),
			ServiceName:         generator.CamelCase(svc.GetName()),
			ServiceRegisterName: generator.CamelCase(svc.GetName()),
		})
		p.P(out.String())
	}
}

func (p *qingcloudPlugin) genServiceClient(
	file *generator.FileDescriptor,
	svc *descriptor.ServiceDescriptorProto,
) {
	const clientHelperFuncTmpl = `
type {{.ServiceName}} struct {
	Config     *config.Config
	Properties *{{.ServiceName}}Properties
}

func New{{.ServiceName}}(conf *config.Config, zone string) (p *{{.ServiceName}}, err error) {
	return &{{.ServiceName}}{
		Config:     conf,
		Properties: &{{.ServiceName}}Properties{Zone: zone},
	}, nil
}

{{.MethodList}}
`
	const clientMethodTmpl = `
func (p *{{.ServiceName}}) {{.MethodName}}(in *{{.ArgsType}}) (out *{{.ReplyType}}, err error) {
	if in == nil {
		in = &{{.ArgsType}}{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "{{.MethodName}}",
		RequestMethod: "{{.RequestMethod}}", // GET or POST
	}

	x := &{{.ReplyType}}{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}`

	// gen client method list
	var methodList string
	for _, m := range svc.Method {
		var RequestMethod = "GET" // default is GET
		if m.Options != nil && proto.HasExtension(m.Options, options.E_Http) {
			if ext, _ := proto.GetExtension(m.Options, options.E_Http); ext != nil {
				if extHttp, _ := ext.(*options.HttpRule); extHttp != nil {
					if kind := extHttp.GetCustom().GetKind(); kind != "" {
						RequestMethod = kind
					}
				}
			}
		}

		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(clientMethodTmpl))
		t.Execute(out, &struct{ ServiceName, ServiceRegisterName, MethodName, ArgsType, ReplyType, RequestMethod string }{
			ServiceName:         generator.CamelCase(svc.GetName()),
			ServiceRegisterName: generator.CamelCase(svc.GetName()),
			MethodName:          generator.CamelCase(m.GetName()),
			ArgsType:            p.TypeName(p.ObjectNamed(m.GetInputType())),
			ReplyType:           p.TypeName(p.ObjectNamed(m.GetOutputType())),
			RequestMethod:       RequestMethod,
		})
		methodList += out.String()
	}

	// gen all client code
	{
		out := bytes.NewBuffer([]byte{})
		t := template.Must(template.New("").Parse(clientHelperFuncTmpl))
		t.Execute(out, &struct{ PackageName, ServiceName, MethodList string }{
			PackageName: file.GetPackage(),
			ServiceName: generator.CamelCase(svc.GetName()),
			MethodList:  methodList,
		})
		p.P(out.String())
	}
}

func init() {
	generator.RegisterPlugin(new(qingcloudPlugin))
}
