// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"path"
	"strings"

	rule_pb "github.com/chai2010/qingcloud-go/pkg/service.pb/qingcloud_sdk_rule"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// qingcloudPlugin produce the Service interface.
type qingcloudPlugin struct {
	*generator.Generator
	GeneratorInterface
}

// Name returns the name of the plugin.
func (p *qingcloudPlugin) Name() string { return "qingcloud" }

// Init is called once after data structures are built but before
// code generation begins.
func (p *qingcloudPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *qingcloudPlugin) InitService(g GeneratorInterface) {
	p.GeneratorInterface = g
}

// Generate produces the code generated by the plugin for this file.
func (p *qingcloudPlugin) GenerateImports(file *generator.FileDescriptor) {
	// skip
}

// Generate generates the Service interface.
// rpc service can't handle other proto message!!!
func (p *qingcloudPlugin) Generate(file *generator.FileDescriptor) {
	if !p.isFileNeedGenerate(file) {
		return
	}

	var buf bytes.Buffer
	fmt.Fprintln(&buf, p.buildFileSpec(file).HeaderCode())

	for _, v := range file.Service {
		fmt.Fprintln(&buf, p.buildServiceSpec(file, v).Code())
	}

	fileContent := buf.String()
	if code, err := format.Source(buf.Bytes()); err != nil {
		log.Printf("qingcloudPlugin.Generate: format %q failed, err = %v", file.GetName(), err)
	} else {
		fileContent = string(code)
	}

	p.Generator.Response.File = append(p.Generator.Response.File, &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(p.goFileName(file)),
		Content: proto.String(fileContent),
	})
}

func (p *qingcloudPlugin) isFileNeedGenerate(file *generator.FileDescriptor) bool {
	for _, v := range p.Generator.Request.FileToGenerate {
		if v == file.GetName() {
			return true
		}
	}
	return false
}

func (p *qingcloudPlugin) goFileName(file *generator.FileDescriptor) string {
	name := *file.Name
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += p.FileNameExt()

	// Does the file have a "go_package" option?
	// If it does, it may override the filename.
	if impPath, _, ok := p.goPackageOption(file); ok && impPath != "" {
		// Replace the existing dirname with the declared import path.
		_, name = path.Split(name)
		name = path.Join(impPath, name)
		return name
	}

	return name
}

func (p *qingcloudPlugin) goPackageOption(file *generator.FileDescriptor) (impPath, pkg string, ok bool) {
	pkg = file.GetOptions().GetGoPackage()
	if pkg == "" {
		return
	}
	ok = true
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(pkg, "/")
	if slash < 0 {
		return
	}
	impPath, pkg = pkg, pkg[slash+1:]
	// A semicolon-delimited suffix overrides the package name.
	sc := strings.IndexByte(impPath, ';')
	if sc < 0 {
		return
	}
	impPath, pkg = impPath[:sc], impPath[sc+1:]
	return
}

func (p *qingcloudPlugin) buildFileSpec(file *generator.FileDescriptor) *FileSpec {
	return &FileSpec{
		IsProto3:    file.GetSyntax() == "proto3",
		FileName:    file.GetName(),
		PackageName: file.PackageName(),
	}
}

func (p *qingcloudPlugin) buildServiceSpec(file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := new(ServiceSpec)
	spec.IsProto3 = file.GetSyntax() == "proto3"

	if rule := p.getServiceRule(svc); rule != nil {
		spec.DocUrl = rule.GetDocUrl()
		spec.ServiceName = rule.GetServiceName()
		spec.MainServiceName = rule.GetMainServiceName()
	} else {
		spec.ServiceName = generator.CamelCase(svc.GetName())
	}

	for _, m := range svc.Method {
		methodSpec := MethodSpec{
			MethodName: generator.CamelCase(m.GetName()),
		}

		if rule := p.getMethodRule(m); rule != nil {
			methodSpec.DocUrl = rule.GetDocUrl()
			methodSpec.HttpMethod = rule.GetHttpMethod()
			methodSpec.InputTypeName = rule.GetInputType()
			methodSpec.OutputTypeName = rule.GetOutputType()
		}

		if methodSpec.InputTypeName == "" {
			methodSpec.InputTypeName = p.TypeName(p.ObjectNamed(m.GetInputType()))
		}
		if methodSpec.OutputTypeName == "" {
			methodSpec.OutputTypeName = p.TypeName(p.ObjectNamed(m.GetOutputType()))
		}
		if methodSpec.HttpMethod == "" {
			methodSpec.HttpMethod = "GET"
		}

		spec.MethodList = append(spec.MethodList, methodSpec)
	}

	return spec
}

func (p *qingcloudPlugin) buildMessageOptionsSpec(file *generator.FileDescriptor, msg *descriptor.DescriptorProto) *MessageOptionsSpec {
	spec := NewMessageOptionsSpec()
	spec.IsProto3 = file.GetSyntax() == "proto3"
	spec.MessageName = generator.CamelCase(msg.GetName())

	for _, field := range msg.Field {
		if rule := p.getMessageFieldRule(field); rule != nil {
			_ = rule
		}

		name := field.GetName()
		typeName := field.GetType().String()
		fixedName := generator.CamelCase(field.GetName())

		spec.FieldNameMap[name] = fixedName
		spec.FiledTypeMap[name] = typeName

		if field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
			spec.RequiredFieldMap[name] = true
		}
		if field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			spec.RepeatedFieldMap[name] = true
		}

		// fix type name
		switch {
		case typeName == "TYPE_BOOL":
			spec.FiledTypeMap[name] = "bool"
		case strings.Contains(typeName, "INT") || strings.Contains(typeName, "FIXED"):
			spec.FiledTypeMap[name] = "int"
		case strings.Contains(typeName, "DOUBLE") || strings.Contains(typeName, "FLOAT"):
			spec.FiledTypeMap[name] = "float"
		case typeName == "TYPE_ENUM":
			spec.FiledTypeMap[name] = "int"
		case typeName == "TYPE_STRING":
			spec.FiledTypeMap[name] = "string"
		case typeName == "TYPE_MESSAGE":
			spec.FiledTypeMap[name] = "message"
		case typeName == "TYPE_BYTES":
			spec.FiledTypeMap[name] = "bytes"
		default:
			spec.FiledTypeMap[name] = "?"
		}
	}

	if rule := p.getMessageRule(msg); rule != nil {
		for name, _ := range spec.FieldNameMap {
			if rule.IsRequired(name) {
				spec.RequiredFieldMap[name] = true
			}

			if rule.HasDefaultValue(name) {
				spec.DefaultValueMap[name] = rule.GetDefaultValue(name)
			}
			if rule.HasEnumValue(name) {
				spec.EnumValueListMap[name] = rule.GetEnumValueList(name)
			}

			if rule.HasMinValue(name) {
				spec.MinValueMap[name] = rule.GetMinValue(name)
			}
			if rule.HasMaxValue(name) {
				spec.MaxValueMap[name] = rule.GetMaxValue(name)
			}
			if rule.HasMultipleOfValue(name) {
				spec.MultipleOfValueMap[name] = rule.GetMultipleOfValue(name)
			}
			if rule.HasRegexpValue(name) {
				spec.RegexpValueMap[name] = fmt.Sprintf("%q", rule.GetRegexpValue(name))
			}
		}
	}

	return spec
}

func (p *qingcloudPlugin) getServiceRule(svc *descriptor.ServiceDescriptorProto) (svcRule *ServiceRule) {
	if svc.Options != nil && proto.HasExtension(svc.Options, rule_pb.E_ServiceRule) {
		if ext, _ := proto.GetExtension(svc.Options, rule_pb.E_ServiceRule); ext != nil {
			if x, _ := ext.(*rule_pb.ServiceOptionsRule); x != nil {
				svcRule = &ServiceRule{x}
			}
		}
	}
	return
}

func (p *qingcloudPlugin) getMethodRule(m *descriptor.MethodDescriptorProto) *MethodRule {
	if m.Options != nil && proto.HasExtension(m.Options, rule_pb.E_MethodRule) {
		if ext, _ := proto.GetExtension(m.Options, rule_pb.E_MethodRule); ext != nil {
			if x, _ := ext.(*rule_pb.MethodOptionsRule); x != nil {
				return &MethodRule{x}
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) getMessageRule(m *descriptor.DescriptorProto) *MessageRule {
	if m.Options != nil && proto.HasExtension(m.Options, rule_pb.E_MessageRule) {
		if ext, _ := proto.GetExtension(m.Options, rule_pb.E_MessageRule); ext != nil {
			if x, _ := ext.(*rule_pb.MessageOptionsRule); x != nil {
				return &MessageRule{x}
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) getMessageFieldRule(m *descriptor.FieldDescriptorProto) *FieldRule {
	if m.Options != nil && proto.HasExtension(m.Options, rule_pb.E_FieldRule) {
		if ext, _ := proto.GetExtension(m.Options, rule_pb.E_FieldRule); ext != nil {
			if x, _ := ext.(*rule_pb.FieldOptionsRule); x != nil {
				return &FieldRule{x}
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) getMethodInputDescriptor(m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	return p.ObjectNamed(m.GetInputType()).(*generator.Descriptor).DescriptorProto
}
func (p *qingcloudPlugin) getMethodOutputDescriptor(m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	return p.ObjectNamed(m.GetOutputType()).(*generator.Descriptor).DescriptorProto
}
