// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import request_data_pkg "github.com/chai2010/qingcloud-go/request/data"
import "github.com/chai2010/qingcloud-go/request/errors"

var _ = config.Config{}
var _ = request.Request{}
var _ = request_data_pkg.Operation{}
var _ = errors.ParameterRequiredError{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ImageServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ImageServiceProperties) Reset()                    { *m = ImageServiceProperties{} }
func (m *ImageServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*ImageServiceProperties) ProtoMessage()               {}
func (*ImageServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ImageServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageServiceProperties)(nil), "spec.ImageServiceProperties")
}

type ImageServiceInterface interface {
	DescribeImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CaptureInstance(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyImageAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	GrantImageToUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RevokeImageFromUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeImageUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CloneImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type ImageService struct {
	Config     *config.Config
	Properties *ImageServiceProperties
}

func NewImageService(conf *config.Config, zone string) (p *ImageService, err error) {
	return &ImageService{
		Config:     conf,
		Properties: &ImageServiceProperties{Zone: zone},
	}, nil
}

func (p *ImageService) DescribeImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeImages",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) CaptureInstance(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CaptureInstance",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) DeleteImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteImages",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) ModifyImageAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyImageAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) GrantImageToUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GrantImageToUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) RevokeImageFromUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RevokeImageFromUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) DescribeImageUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeImageUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *ImageService) CloneImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CloneImages",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func init() { proto.RegisterFile("image.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0x19, 0x8c, 0x3f, 0xfc, 0xef, 0x86, 0x62, 0xd0, 0xa1, 0x13, 0x54, 0x7c, 0xf2, 0x41,
	0x3a, 0xd0, 0x67, 0xc5, 0xd1, 0xba, 0xb1, 0x07, 0x41, 0xa6, 0x7e, 0x80, 0xb4, 0xde, 0x95, 0x60,
	0x9b, 0x1b, 0x6e, 0x6e, 0x27, 0xf3, 0x43, 0xf9, 0x19, 0x65, 0xa9, 0x03, 0x2d, 0xf8, 0xb0, 0xbe,
	0x25, 0xe7, 0x9c, 0xfc, 0xee, 0x49, 0x02, 0x3d, 0x53, 0xea, 0x1c, 0x23, 0xc7, 0x24, 0xa4, 0xba,
	0xde, 0x61, 0x36, 0x3c, 0xca, 0x89, 0xf2, 0x02, 0x47, 0x41, 0x4b, 0xab, 0xc5, 0x48, 0xdb, 0x55,
	0x1d, 0x18, 0x1e, 0x37, 0x2d, 0x2c, 0x9d, 0x6c, 0xcc, 0xd3, 0xa6, 0x29, 0xa6, 0x44, 0x2f, 0xba,
	0x74, 0xdf, 0x81, 0x93, 0x66, 0xe0, 0x9d, 0xb5, 0x73, 0xc8, 0xbe, 0xf6, 0xcf, 0x2f, 0x61, 0x30,
	0x5b, 0xb7, 0x79, 0x42, 0x5e, 0x9a, 0x0c, 0x1f, 0x99, 0x1c, 0xb2, 0x18, 0xf4, 0x4a, 0x41, 0xf7,
	0x83, 0x2c, 0x1e, 0x76, 0xce, 0x3a, 0x17, 0xff, 0xe7, 0x61, 0x7d, 0xf5, 0xd9, 0x85, 0xfe, 0xcf,
	0xb8, 0xba, 0x83, 0x9d, 0x04, 0x7d, 0xc6, 0x26, 0xc5, 0xa0, 0x7b, 0x35, 0x88, 0xea, 0x89, 0xd1,
	0x66, 0x62, 0x74, 0xbf, 0xee, 0x3b, 0xfc, 0x43, 0x57, 0x63, 0xd8, 0x8d, 0xb5, 0x93, 0x8a, 0x71,
	0x66, 0xbd, 0x68, 0x9b, 0xe1, 0xd6, 0x88, 0x5b, 0xe8, 0x27, 0x58, 0xa0, 0xb4, 0xad, 0x30, 0x85,
	0x83, 0x07, 0x7a, 0x35, 0x8b, 0x55, 0x38, 0x3f, 0x16, 0x61, 0x93, 0x56, 0xd2, 0x02, 0x14, 0xc3,
	0xde, 0x94, 0xb5, 0x95, 0xc0, 0x79, 0xa6, 0x17, 0x8f, 0xbc, 0x3d, 0x64, 0x02, 0xfb, 0x73, 0x5c,
	0xd2, 0x5b, 0x7d, 0x9b, 0x09, 0x53, 0xd9, 0x8e, 0x93, 0x80, 0xfa, 0xf5, 0x35, 0xed, 0x28, 0x37,
	0xd0, 0x8b, 0x0b, 0xb2, 0x2d, 0x9f, 0x36, 0xfd, 0x17, 0xf6, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0xbc, 0xce, 0xc5, 0xf3, 0x02, 0x00, 0x00,
}
