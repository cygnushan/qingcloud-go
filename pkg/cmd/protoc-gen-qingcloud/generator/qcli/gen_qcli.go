// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package gen_qcli

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

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
	t.Execute(&buf, spec)
	return buf.String()
}

func (pkgGenerator) ServiceCode(g *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string {
	spec := utils.BuildServiceSpec(g, file, svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmplService))
	t.Execute(&buf, spec)
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
}

const tmplFileHeader = `// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: {{.GetFileName}}

package qcli_pb

import (
	"fmt"

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

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

`

const tmplService = `
{{$service := .}}

func init() {
	AllCommands = append(AllCommands, Cmd{{.GetServiceName}})
}

var Cmd{{.GetServiceName}} = cli.Command{
	Name:    "{{strings_TrimSuffix .GetServiceName "Service" | strings_ToLower}}",
	Aliases: []string{},
	Usage:   "manage {{.GetServiceName}}",
	Subcommands: []cli.Command{
		{{range $_, $m := .GetMethodList -}}
		{
			Name:    "{{$m.GetMethodName}}",
			Aliases: []string{},
			Usage:   "{{$m.GetMethodName}}",
			Action: cmd{{$m.GetMethodName}},
		},
		{{end}}
	},
}

{{range $_, $m := .GetMethodList}}
func cmd{{$m.GetMethodName}}(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.New{{$service.GetServiceName}}(conf, zone)

	in := new(pb.{{$m.GetInputTypeName}})

	// TODO: fill field from flags

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
