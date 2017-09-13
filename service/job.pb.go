// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

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

type JobServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *JobServiceProperties) Reset()                    { *m = JobServiceProperties{} }
func (m *JobServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*JobServiceProperties) ProtoMessage()               {}
func (*JobServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *JobServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeJobsInput struct {
	Jobs    []string `protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
	Limit   int32    `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Offset  int32    `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Status  []string `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	Verbose int32    `protobuf:"varint,5,opt,name=verbose" json:"verbose,omitempty"`
}

func (m *DescribeJobsInput) Reset()                    { *m = DescribeJobsInput{} }
func (m *DescribeJobsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeJobsInput) ProtoMessage()               {}
func (*DescribeJobsInput) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *DescribeJobsInput) GetJobs() []string {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *DescribeJobsInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeJobsInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeJobsInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeJobsInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

type DescribeJobsOutput struct {
	Action     string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobSet     []*Job `protobuf:"bytes,4,rep,name=job_set,json=jobSet" json:"job_set,omitempty"`
	TotalCount int32  `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeJobsOutput) Reset()                    { *m = DescribeJobsOutput{} }
func (m *DescribeJobsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeJobsOutput) ProtoMessage()               {}
func (*DescribeJobsOutput) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *DescribeJobsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeJobsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeJobsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeJobsOutput) GetJobSet() []*Job {
	if m != nil {
		return m.JobSet
	}
	return nil
}

func (m *DescribeJobsOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*JobServiceProperties)(nil), "service.JobServiceProperties")
	proto.RegisterType((*DescribeJobsInput)(nil), "service.DescribeJobsInput")
	proto.RegisterType((*DescribeJobsOutput)(nil), "service.DescribeJobsOutput")
}

type JobServiceInterface interface {
	DescribeJobs(in *DescribeJobsInput) (out *DescribeJobsOutput, err error)
}

type JobService struct {
	Config           *config.Config
	Properties       *JobServiceProperties
	LastResponseBody string
}

func NewJobService(conf *config.Config, zone string) (p *JobService) {
	return &JobService{
		Config:     conf,
		Properties: &JobServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Job(zone string) (*JobService, error) {
	properties := &JobServiceProperties{
		Zone: zone,
	}

	return &JobService{Config: s.Config, Properties: properties}, nil
}

func (p *JobService) DescribeJobs(in *DescribeJobsInput) (out *DescribeJobsOutput, err error) {
	if in == nil {
		in = &DescribeJobsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeJobs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeJobsOutput{}
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

func (p *DescribeJobsInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("job.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x8a, 0xd4, 0x40,
	0x14, 0xc6, 0x89, 0xd3, 0x9d, 0xd8, 0xaf, 0xc7, 0x85, 0xc5, 0x30, 0xc4, 0x88, 0xd8, 0x34, 0x08,
	0x83, 0x60, 0x02, 0xe3, 0xce, 0x85, 0x0b, 0x5b, 0x10, 0xb3, 0x51, 0xe3, 0x01, 0x42, 0xaa, 0xf2,
	0xa6, 0xa7, 0x62, 0x92, 0x17, 0xab, 0x5e, 0x06, 0xf5, 0x02, 0x82, 0x37, 0xf0, 0x08, 0x1e, 0xcf,
	0x1b, 0x48, 0x55, 0x32, 0x6d, 0x83, 0x6e, 0x42, 0x7d, 0xdf, 0xfb, 0x93, 0x5f, 0xbe, 0x14, 0xac,
	0x1a, 0x92, 0xe9, 0x60, 0x88, 0x49, 0x44, 0x16, 0xcd, 0x8d, 0x56, 0x98, 0xac, 0xf9, 0xeb, 0x80,
	0x76, 0x72, 0x93, 0x47, 0x9f, 0x75, 0xbf, 0x57, 0x2d, 0x8d, 0x75, 0x69, 0xeb, 0x4f, 0xa5, 0x19,
	0x5b, 0xcc, 0xdc, 0x63, 0x2a, 0x6f, 0x9f, 0xc2, 0x59, 0x4e, 0xf2, 0xe3, 0x34, 0xf9, 0xde, 0xd0,
	0x80, 0x86, 0x35, 0x5a, 0x21, 0x60, 0xf1, 0x8d, 0x7a, 0x8c, 0x83, 0x4d, 0x70, 0xb1, 0x2a, 0xfc,
	0x79, 0xfb, 0x3d, 0x80, 0xfb, 0xaf, 0xd1, 0x2a, 0xa3, 0x25, 0xe6, 0x24, 0xed, 0xdb, 0x7e, 0x18,
	0xd9, 0x75, 0x36, 0x24, 0x6d, 0x1c, 0x6c, 0x4e, 0x5c, 0xa7, 0x3b, 0x8b, 0x33, 0x58, 0xb6, 0xba,
	0xd3, 0x1c, 0xdf, 0xd9, 0x04, 0x17, 0xcb, 0x62, 0x12, 0xe2, 0x1c, 0x42, 0xba, 0xba, 0xb2, 0xc8,
	0xf1, 0x89, 0xb7, 0x67, 0xe5, 0x7c, 0xcb, 0x15, 0x8f, 0x36, 0x5e, 0xf8, 0x1d, 0xb3, 0x12, 0x31,
	0x44, 0x37, 0x68, 0x24, 0x59, 0x8c, 0x97, 0x7e, 0xe0, 0x56, 0x6e, 0x7f, 0x05, 0x20, 0x8e, 0x49,
	0xde, 0x8d, 0xec, 0x50, 0xce, 0x21, 0xac, 0x14, 0x6b, 0xea, 0x67, 0xec, 0x59, 0x89, 0x07, 0x70,
	0xd7, 0x20, 0x97, 0x8a, 0x6a, 0x9c, 0x89, 0x22, 0x83, 0xbc, 0xa3, 0x1a, 0xdd, 0x3b, 0x3a, 0xb4,
	0xb6, 0xda, 0xa3, 0x87, 0x5a, 0x15, 0xb7, 0x52, 0x3c, 0x81, 0xa8, 0x21, 0x59, 0x3a, 0x5c, 0x87,
	0xb5, 0xbe, 0x3c, 0x4d, 0xe7, 0x80, 0xd3, 0x9c, 0x64, 0x11, 0x36, 0x2e, 0x36, 0x16, 0x8f, 0x61,
	0xcd, 0xc4, 0x55, 0x5b, 0x2a, 0x1a, 0x7b, 0x9e, 0x41, 0xc1, 0x5b, 0x3b, 0xe7, 0x5c, 0xfe, 0x0c,
	0x00, 0xfe, 0x46, 0x2c, 0xde, 0xc0, 0xe9, 0x31, 0xb9, 0x48, 0x0e, 0x5b, 0xff, 0x89, 0x36, 0x79,
	0xf8, 0xdf, 0xda, 0xf4, 0xb1, 0xc9, 0xab, 0x1f, 0xbf, 0x17, 0x2f, 0xe1, 0xd9, 0x35, 0xf3, 0x60,
	0x5f, 0x64, 0x59, 0x4d, 0xca, 0xa6, 0x87, 0x7f, 0x9d, 0x2a, 0xea, 0xb2, 0x6a, 0xd0, 0x59, 0x43,
	0x32, 0xd3, 0x7d, 0x8d, 0x5f, 0xd2, 0x6b, 0xee, 0x5a, 0x71, 0xef, 0x83, 0xee, 0xf7, 0x3b, 0xdf,
	0x91, 0x93, 0x94, 0xa1, 0xbf, 0x04, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x78, 0xc1,
	0x76, 0x46, 0x02, 0x00, 0x00,
}
