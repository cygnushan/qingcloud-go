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

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// qingcloudPlugin produce the Service interface.
type qingcloudPlugin struct {
	*generator.Generator
	ServiceGenerator
}

// Name returns the name of the plugin.
func (p *qingcloudPlugin) Name() string { return "qingcloud" }

// Init is called once after data structures are built but before
// code generation begins.
func (p *qingcloudPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *qingcloudPlugin) InitService(g ServiceGenerator) {
	p.ServiceGenerator = g
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
	fmt.Fprintln(&buf, p.HeaderCode(p.Generator, file))

	for _, msg := range file.MessageType {
		fmt.Fprintln(&buf, p.MessageCode(p.Generator, file, msg))
	}

	for _, svc := range file.Service {
		fmt.Fprintln(&buf, p.ServiceCode(p.Generator, file, svc))
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

func (p *qingcloudPlugin) _buildFileSpec(file *generator.FileDescriptor) *spec_metadata.FileSpec {
	return &spec_metadata.FileSpec{
		FileName:    proto.String(file.GetName()),
		PackageName: proto.String(file.PackageName()),
	}
}

func (p *qingcloudPlugin) _buildServiceSpec(file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceSpec {
	spec := new(spec_metadata.ServiceSpec)

	spec.ServiceName = proto.String(generator.CamelCase(svc.GetName()))

	if opt := p._getFileOption(file.FileDescriptorProto); opt != nil {
		_ = opt
	}

	for _, m := range svc.Method {
		methodSpec := &spec_metadata.ServiceMethodSpec{
			MethodName:     proto.String(generator.CamelCase(m.GetName())),
			InputTypeName:  proto.String(p.TypeName(p.ObjectNamed(m.GetInputType()))),
			OutputTypeName: proto.String(p.TypeName(p.ObjectNamed(m.GetOutputType()))),
		}

		if opt := p._getServiceMethodOption(m); opt != nil {
			methodSpec.HttpMethod = proto.String(opt.GetHttpMethod())
		}

		spec.MethodList = append(spec.MethodList, methodSpec)
	}

	return spec
}

func (p *qingcloudPlugin) _getFileOption(file *descriptor.FileDescriptorProto) *spec_metadata.FileOption {
	if file.Options != nil && proto.HasExtension(file.Options, spec_metadata.E_FileOption) {
		if ext, _ := proto.GetExtension(file.Options, spec_metadata.E_FileOption); ext != nil {
			if x, _ := ext.(*spec_metadata.FileOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) _getServiceOption(svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceOption {
	if svc.Options != nil && proto.HasExtension(svc.Options, spec_metadata.E_FileOption) {
		if ext, _ := proto.GetExtension(svc.Options, spec_metadata.E_FileOption); ext != nil {
			if x, _ := ext.(*spec_metadata.ServiceOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) _getServiceMethodOption(m *descriptor.MethodDescriptorProto) *spec_metadata.MethodOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_MethodOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_MethodOption); ext != nil {
			if x, _ := ext.(*spec_metadata.MethodOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) _getMessageOption(m *descriptor.DescriptorProto) *spec_metadata.MessageOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_MessageOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_MessageOption); ext != nil {
			if x, _ := ext.(*spec_metadata.MessageOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) _getMessageFieldOption(m *descriptor.FieldDescriptorProto) *spec_metadata.FieldOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_FieldOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_FieldOption); ext != nil {
			if x, _ := ext.(*spec_metadata.FieldOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func (p *qingcloudPlugin) _getMethodInputDescriptor(m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	return p.ObjectNamed(m.GetInputType()).(*generator.Descriptor).DescriptorProto
}
func (p *qingcloudPlugin) _getMethodOutputDescriptor(m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	return p.ObjectNamed(m.GetOutputType()).(*generator.Descriptor).DescriptorProto
}
