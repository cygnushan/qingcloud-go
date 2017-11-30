// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package golang_validate_v1

import (
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/generator"
)

var tmplFuncMap = template.FuncMap{
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,

	"camel_case": generator.CamelCase,
}

const tmplFileHeader = `// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud-validate
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud-validate/generator/golang
// source: {{.File.GetName}}

package {{.File.PackageName}}

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

`

const tmplValidate = `
func (p *{{camel_case .Message.GetName}}) Validate() error {
	{{- range $_, $m := .FieldValidate}}
		{{$m}}
	{{- end}}

	return nil
}
`