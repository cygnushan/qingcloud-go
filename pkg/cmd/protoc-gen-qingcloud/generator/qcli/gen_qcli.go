// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package gen_qcli

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud"
	"github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/utils"
)

func init() {
	plugin.RegisterServiceGenerator(new(pkgGenerator))
}

type pkgGenerator struct{}

func (p pkgGenerator) Name() string        { return "qcli" }
func (p pkgGenerator) FileNameExt() string { return ".pb.qcli.go" }

func (pkgGenerator) HeaderCode(g *generator.Generator, file *generator.FileDescriptor) string {
	spec := utils.BuildFileSpec(g, file)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplFileHeader))
	err := t.Execute(&buf,
		struct {
			G    *generator.Generator
			File *generator.FileDescriptor
			Spec *spec_metadata.FileSpec
		}{
			G:    g,
			File: file,
			Spec: spec,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func (pkgGenerator) ServiceCode(g *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string {
	spec := utils.BuildServiceSpec(g, file, svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmplService))
	err := t.Execute(&buf,
		struct {
			G       *generator.Generator
			File    *generator.FileDescriptor
			Service *descriptor.ServiceDescriptorProto
			Spec    *spec_metadata.ServiceSpec
		}{
			G:       g,
			File:    file,
			Service: svc,
			Spec:    spec,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func (pkgGenerator) MessageCode(p *generator.Generator, file *generator.FileDescriptor, msg *descriptor.DescriptorProto) string {
	return ""
}

var tmplFuncMap = template.FuncMap{
	"strings_Title":   strings.Title,
	"strings_ToLower": strings.ToLower,
	"strings_ToUpper": strings.ToUpper,

	"strings_TrimPrefix": strings.TrimPrefix,
	"strings_TrimSuffix": strings.TrimSuffix,

	"generator_CamelCase": generator.CamelCase,

	"utils_GetMethodInputDescriptor":  utils.GetMethodInputDescriptor,
	"utils_GetMethodOutputDescriptor": utils.GetMethodOutputDescriptor,

	"pkgGenCmdFlagName":   pkgGenCmdFlagName,
	"pkgGenCmdFlagUsage":  pkgGenCmdFlagUsage,
	"pkgGenCmdFlagParser": pkgGenCmdFlagParser,
}

func pkgGenCmdFlagName(field *descriptor.FieldDescriptorProto) string {
	return ""
}
func pkgGenCmdFlagUsage(field *descriptor.FieldDescriptorProto) string {
	return ""
}

func pkgGenCmdFlagParser(field *descriptor.FieldDescriptorProto) string {
	return ""
}

const tmplFileHeader = `
{{- $root := . -}}

// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: {{.Spec.GetFileName}}

package qcli_pb

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
	"github.com/chai2010/qingcloud-go/pkg/logger"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

`

const tmplService = `
{{$root := .}}
{{$service := .Spec}}

func init() {
	AllCommands = append(AllCommands, Cmd{{.Spec.GetServiceName}})
}

var Cmd{{.Spec.GetServiceName}} = cli.Command{
	Name:    "{{strings_TrimSuffix .Spec.GetServiceName "Service" | strings_ToLower}}",
	Aliases: []string{},
	Usage:   "manage {{.Spec.GetServiceName}}",
	Subcommands: []cli.Command{
		{{range $_, $m := .Spec.GetMethodList -}}
		{
			Name:    "{{$m.GetMethodName}}",
			Aliases: []string{},
			Usage:   "{{$m.GetMethodName}}",
			Flags:   _flag_{{$service.GetServiceName}}_{{$m.GetMethodName}},
			Action:  _cmd_{{$service.GetServiceName}}_{{$m.GetMethodName}},
		},
		{{end}}
	},
}

{{/* generate flags */}}
{{range $_, $m := $root.Service.Method}}
	{{$in := utils_GetMethodInputDescriptor $root.G $m}}
	{{range $_, $field := $in.Field}}
		{{/* todo */}}
	{{end}}
{{end}}

{{/* generate command functions */}}
{{range $_, $m := .Spec.GetMethodList}}
var _flag_{{$service.GetServiceName}}_{{$m.GetMethodName}} = []cli.Flag{ /* fields */ }

func _cmd_{{$service.GetServiceName}}_{{$m.GetMethodName}}(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.New{{$service.GetServiceName}}(conf, zone)

	in := new(pb.{{$m.GetInputTypeName}})

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.{{$m.GetMethodName}}(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}
{{end}}
`
