// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = config.Config{}
var _ = request.Request{}
var _ = data.Operation{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DNSAliasServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *DNSAliasServiceProperties) Reset()                    { *m = DNSAliasServiceProperties{} }
func (m *DNSAliasServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*DNSAliasServiceProperties) ProtoMessage()               {}
func (*DNSAliasServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *DNSAliasServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeDNSAliasesInput struct {
}

func (m *DescribeDNSAliasesInput) Reset()                    { *m = DescribeDNSAliasesInput{} }
func (m *DescribeDNSAliasesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesInput) ProtoMessage()               {}
func (*DescribeDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type DescribeDNSAliasesOutput struct {
}

func (m *DescribeDNSAliasesOutput) Reset()                    { *m = DescribeDNSAliasesOutput{} }
func (m *DescribeDNSAliasesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesOutput) ProtoMessage()               {}
func (*DescribeDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type AssociateDNSAliasInput struct {
}

func (m *AssociateDNSAliasInput) Reset()                    { *m = AssociateDNSAliasInput{} }
func (m *AssociateDNSAliasInput) String() string            { return proto.CompactTextString(m) }
func (*AssociateDNSAliasInput) ProtoMessage()               {}
func (*AssociateDNSAliasInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type AssociateDNSAliasOutput struct {
}

func (m *AssociateDNSAliasOutput) Reset()                    { *m = AssociateDNSAliasOutput{} }
func (m *AssociateDNSAliasOutput) String() string            { return proto.CompactTextString(m) }
func (*AssociateDNSAliasOutput) ProtoMessage()               {}
func (*AssociateDNSAliasOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

type DissociateDNSAliasesInput struct {
}

func (m *DissociateDNSAliasesInput) Reset()                    { *m = DissociateDNSAliasesInput{} }
func (m *DissociateDNSAliasesInput) String() string            { return proto.CompactTextString(m) }
func (*DissociateDNSAliasesInput) ProtoMessage()               {}
func (*DissociateDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type DissociateDNSAliasesOutput struct {
}

func (m *DissociateDNSAliasesOutput) Reset()                    { *m = DissociateDNSAliasesOutput{} }
func (m *DissociateDNSAliasesOutput) String() string            { return proto.CompactTextString(m) }
func (*DissociateDNSAliasesOutput) ProtoMessage()               {}
func (*DissociateDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

type GetDNSLabelInput struct {
}

func (m *GetDNSLabelInput) Reset()                    { *m = GetDNSLabelInput{} }
func (m *GetDNSLabelInput) String() string            { return proto.CompactTextString(m) }
func (*GetDNSLabelInput) ProtoMessage()               {}
func (*GetDNSLabelInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

type GetDNSLabelOutput struct {
}

func (m *GetDNSLabelOutput) Reset()                    { *m = GetDNSLabelOutput{} }
func (m *GetDNSLabelOutput) String() string            { return proto.CompactTextString(m) }
func (*GetDNSLabelOutput) ProtoMessage()               {}
func (*GetDNSLabelOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func init() {
	proto.RegisterType((*DNSAliasServiceProperties)(nil), "spec.DNSAliasServiceProperties")
	proto.RegisterType((*DescribeDNSAliasesInput)(nil), "spec.DescribeDNSAliasesInput")
	proto.RegisterType((*DescribeDNSAliasesOutput)(nil), "spec.DescribeDNSAliasesOutput")
	proto.RegisterType((*AssociateDNSAliasInput)(nil), "spec.AssociateDNSAliasInput")
	proto.RegisterType((*AssociateDNSAliasOutput)(nil), "spec.AssociateDNSAliasOutput")
	proto.RegisterType((*DissociateDNSAliasesInput)(nil), "spec.DissociateDNSAliasesInput")
	proto.RegisterType((*DissociateDNSAliasesOutput)(nil), "spec.DissociateDNSAliasesOutput")
	proto.RegisterType((*GetDNSLabelInput)(nil), "spec.GetDNSLabelInput")
	proto.RegisterType((*GetDNSLabelOutput)(nil), "spec.GetDNSLabelOutput")
}

type DNSAliasServiceInterface interface {
	DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error)
	AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error)
	DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error)
	GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error)
}

type DNSAliasService struct {
	Config     *config.Config
	Properties *DNSAliasServiceProperties
}

func NewDNSAliasService(conf *config.Config, zone string) (p *DNSAliasService, err error) {
	return &DNSAliasService{
		Config:     conf,
		Properties: &DNSAliasServiceProperties{Zone: zone},
	}, nil
}

func (p *DNSAliasService) DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error) {
	if in == nil {
		in = &DescribeDNSAliasesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeDNSAliases",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeDNSAliasesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DescribeDNSAliasesInput) Validate() error {
	return nil
}

func (p *DNSAliasService) AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error) {
	if in == nil {
		in = &AssociateDNSAliasInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateDNSAlias",
		RequestMethod: "GET", // GET or POST
	}

	x := &AssociateDNSAliasOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *AssociateDNSAliasInput) Validate() error {
	return nil
}

func (p *DNSAliasService) DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error) {
	if in == nil {
		in = &DissociateDNSAliasesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateDNSAliases",
		RequestMethod: "GET", // GET or POST
	}

	x := &DissociateDNSAliasesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DissociateDNSAliasesInput) Validate() error {
	return nil
}

func (p *DNSAliasService) GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error) {
	if in == nil {
		in = &GetDNSLabelInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetDNSLabel",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetDNSLabelOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *GetDNSLabelInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x14, 0xa1, 0xe3, 0x41, 0x3b, 0x4a, 0x93, 0xac, 0xad, 0x96, 0x9c, 0x3c, 0x45,
	0xd0, 0xbb, 0x50, 0x08, 0x88, 0x20, 0xb5, 0x98, 0x93, 0xc7, 0x24, 0xce, 0x61, 0xa1, 0x64, 0xc3,
	0xce, 0xd6, 0x83, 0xaf, 0xe9, 0x0b, 0x49, 0xbb, 0xdb, 0x12, 0xb2, 0xd9, 0x5b, 0x98, 0x6f, 0xbe,
	0x99, 0xcc, 0xcf, 0xc2, 0xe4, 0xbb, 0xe1, 0xac, 0xd5, 0xca, 0x28, 0x1c, 0x73, 0x4b, 0x75, 0xfa,
	0x08, 0x49, 0xbe, 0x2e, 0x56, 0x5b, 0x59, 0x72, 0x41, 0xfa, 0x47, 0xd6, 0xb4, 0xd1, 0xaa, 0x25,
	0x6d, 0x24, 0x31, 0x22, 0x8c, 0x7f, 0x55, 0x43, 0xf1, 0x68, 0x39, 0x7a, 0x98, 0x7c, 0x1e, 0xbe,
	0xd3, 0x04, 0xa2, 0x9c, 0xb8, 0xd6, 0xb2, 0xa2, 0xa3, 0x48, 0xfc, 0xd6, 0xb4, 0x3b, 0x93, 0x0a,
	0x88, 0x7d, 0xf4, 0xb1, 0x33, 0x7b, 0x16, 0xc3, 0x6c, 0xc5, 0xac, 0x6a, 0x59, 0x9a, 0x13, 0xb4,
	0x56, 0x02, 0x91, 0x47, 0x9c, 0x74, 0x0b, 0x49, 0x2e, 0xfb, 0xec, 0xb8, 0x6d, 0x0e, 0x62, 0x08,
	0x3a, 0x15, 0xe1, 0xea, 0x95, 0x4c, 0xbe, 0x2e, 0xde, 0xcb, 0x8a, 0xb6, 0xd6, 0xb8, 0x86, 0x69,
	0xa7, 0x66, 0x1b, 0x9f, 0xfe, 0xce, 0xe0, 0xb2, 0x97, 0x00, 0x16, 0x80, 0xfe, 0x21, 0xb8, 0xc8,
	0xf6, 0x89, 0x65, 0x81, 0xeb, 0xc5, 0x5d, 0x08, 0xdb, 0x45, 0xb8, 0x81, 0xa9, 0x77, 0x27, 0xce,
	0xad, 0x34, 0x1c, 0x8d, 0x58, 0x04, 0xa8, 0x9b, 0xf8, 0x05, 0x37, 0x43, 0x09, 0xe0, 0xbd, 0xfb,
	0x93, 0x50, 0x74, 0x62, 0x19, 0x6e, 0x70, 0xa3, 0x5f, 0xe0, 0xa2, 0x13, 0x15, 0xce, 0xac, 0xd0,
	0x4f, 0x54, 0x44, 0x5e, 0xdd, 0xfa, 0xd5, 0xf9, 0xe1, 0x8d, 0x3d, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x50, 0x35, 0xe3, 0x70, 0x02, 0x00, 0x00,
}
