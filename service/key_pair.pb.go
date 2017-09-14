// Code generated by protoc-gen-go. DO NOT EDIT.
// source: key_pair.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

import "regexp"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/logger"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.SetLevel
var _ = request.Request{}
var _ = data.Operation{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type KeyPairServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *KeyPairServiceProperties) Reset()                    { *m = KeyPairServiceProperties{} }
func (m *KeyPairServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*KeyPairServiceProperties) ProtoMessage()               {}
func (*KeyPairServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *KeyPairServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeKeyPairsInput struct {
	Keypairs      []string `protobuf:"bytes,1,rep,name=keypairs" json:"keypairs,omitempty"`
	InstanceId    string   `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Owner         string   `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	EncryptMethod string   `protobuf:"bytes,4,opt,name=encrypt_method,json=encryptMethod" json:"encrypt_method,omitempty"`
	SearchWord    string   `protobuf:"bytes,5,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Tags          []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Verbose       int32    `protobuf:"varint,7,opt,name=verbose" json:"verbose,omitempty"`
	Offset        int32    `protobuf:"varint,8,opt,name=offset" json:"offset,omitempty"`
	Limit         int32    `protobuf:"varint,9,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeKeyPairsInput) Reset()                    { *m = DescribeKeyPairsInput{} }
func (m *DescribeKeyPairsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeKeyPairsInput) ProtoMessage()               {}
func (*DescribeKeyPairsInput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *DescribeKeyPairsInput) GetKeypairs() []string {
	if m != nil {
		return m.Keypairs
	}
	return nil
}

func (m *DescribeKeyPairsInput) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *DescribeKeyPairsInput) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DescribeKeyPairsInput) GetEncryptMethod() string {
	if m != nil {
		return m.EncryptMethod
	}
	return ""
}

func (m *DescribeKeyPairsInput) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *DescribeKeyPairsInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeKeyPairsInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

func (m *DescribeKeyPairsInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeKeyPairsInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeKeyPairsOutput struct {
	Action     string                                 `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                                  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                                 `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	KeypairSet []*DescribeKeyPairsOutput_ResponseItem `protobuf:"bytes,4,rep,name=keypair_set,json=keypairSet" json:"keypair_set,omitempty"`
	TotalCount int32                                  `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeKeyPairsOutput) Reset()                    { *m = DescribeKeyPairsOutput{} }
func (m *DescribeKeyPairsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeKeyPairsOutput) ProtoMessage()               {}
func (*DescribeKeyPairsOutput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *DescribeKeyPairsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeKeyPairsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeKeyPairsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeKeyPairsOutput) GetKeypairSet() []*DescribeKeyPairsOutput_ResponseItem {
	if m != nil {
		return m.KeypairSet
	}
	return nil
}

func (m *DescribeKeyPairsOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeKeyPairsOutput_ResponseItem struct {
	KeypairId     string   `protobuf:"bytes,1,opt,name=keypair_id,json=keypairId" json:"keypair_id,omitempty"`
	KeypairName   string   `protobuf:"bytes,2,opt,name=keypair_name,json=keypairName" json:"keypair_name,omitempty"`
	Description   string   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	EncryptMethod string   `protobuf:"bytes,4,opt,name=encrypt_method,json=encryptMethod" json:"encrypt_method,omitempty"`
	PubKey        string   `protobuf:"bytes,5,opt,name=pub_key,json=pubKey" json:"pub_key,omitempty"`
	InstanceIds   []string `protobuf:"bytes,6,rep,name=instance_ids,json=instanceIds" json:"instance_ids,omitempty"`
}

func (m *DescribeKeyPairsOutput_ResponseItem) Reset()         { *m = DescribeKeyPairsOutput_ResponseItem{} }
func (m *DescribeKeyPairsOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeKeyPairsOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeKeyPairsOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{2, 0}
}

func (m *DescribeKeyPairsOutput_ResponseItem) GetKeypairId() string {
	if m != nil {
		return m.KeypairId
	}
	return ""
}

func (m *DescribeKeyPairsOutput_ResponseItem) GetKeypairName() string {
	if m != nil {
		return m.KeypairName
	}
	return ""
}

func (m *DescribeKeyPairsOutput_ResponseItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DescribeKeyPairsOutput_ResponseItem) GetEncryptMethod() string {
	if m != nil {
		return m.EncryptMethod
	}
	return ""
}

func (m *DescribeKeyPairsOutput_ResponseItem) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *DescribeKeyPairsOutput_ResponseItem) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

type CreateKeyPairInput struct {
	KeypairName   string `protobuf:"bytes,1,opt,name=keypair_name,json=keypairName" json:"keypair_name,omitempty"`
	Mode          string `protobuf:"bytes,2,opt,name=mode" json:"mode,omitempty"`
	EncryptMethod string `protobuf:"bytes,3,opt,name=encrypt_method,json=encryptMethod" json:"encrypt_method,omitempty"`
	PublicKey     string `protobuf:"bytes,4,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *CreateKeyPairInput) Reset()                    { *m = CreateKeyPairInput{} }
func (m *CreateKeyPairInput) String() string            { return proto.CompactTextString(m) }
func (*CreateKeyPairInput) ProtoMessage()               {}
func (*CreateKeyPairInput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *CreateKeyPairInput) GetKeypairName() string {
	if m != nil {
		return m.KeypairName
	}
	return ""
}

func (m *CreateKeyPairInput) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *CreateKeyPairInput) GetEncryptMethod() string {
	if m != nil {
		return m.EncryptMethod
	}
	return ""
}

func (m *CreateKeyPairInput) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type CreateKeyPairOutput struct {
	Action     string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	PrivateKey string `protobuf:"bytes,4,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	KeypairId  string `protobuf:"bytes,5,opt,name=keypair_id,json=keypairId" json:"keypair_id,omitempty"`
}

func (m *CreateKeyPairOutput) Reset()                    { *m = CreateKeyPairOutput{} }
func (m *CreateKeyPairOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateKeyPairOutput) ProtoMessage()               {}
func (*CreateKeyPairOutput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *CreateKeyPairOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateKeyPairOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateKeyPairOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateKeyPairOutput) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *CreateKeyPairOutput) GetKeypairId() string {
	if m != nil {
		return m.KeypairId
	}
	return ""
}

type DeleteKeyPairsInput struct {
	Keypairs []string `protobuf:"bytes,1,rep,name=keypairs" json:"keypairs,omitempty"`
}

func (m *DeleteKeyPairsInput) Reset()                    { *m = DeleteKeyPairsInput{} }
func (m *DeleteKeyPairsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteKeyPairsInput) ProtoMessage()               {}
func (*DeleteKeyPairsInput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *DeleteKeyPairsInput) GetKeypairs() []string {
	if m != nil {
		return m.Keypairs
	}
	return nil
}

type DeleteKeyPairsOutput struct {
	Action   string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode  int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message  string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Keypairs []string `protobuf:"bytes,4,rep,name=keypairs" json:"keypairs,omitempty"`
}

func (m *DeleteKeyPairsOutput) Reset()                    { *m = DeleteKeyPairsOutput{} }
func (m *DeleteKeyPairsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteKeyPairsOutput) ProtoMessage()               {}
func (*DeleteKeyPairsOutput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *DeleteKeyPairsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteKeyPairsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteKeyPairsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteKeyPairsOutput) GetKeypairs() []string {
	if m != nil {
		return m.Keypairs
	}
	return nil
}

type AttachKeyPairsInput struct {
	Keypairs  []string `protobuf:"bytes,1,rep,name=keypairs" json:"keypairs,omitempty"`
	Instances []string `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
}

func (m *AttachKeyPairsInput) Reset()                    { *m = AttachKeyPairsInput{} }
func (m *AttachKeyPairsInput) String() string            { return proto.CompactTextString(m) }
func (*AttachKeyPairsInput) ProtoMessage()               {}
func (*AttachKeyPairsInput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *AttachKeyPairsInput) GetKeypairs() []string {
	if m != nil {
		return m.Keypairs
	}
	return nil
}

func (m *AttachKeyPairsInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

type AttachKeyPairsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *AttachKeyPairsOutput) Reset()                    { *m = AttachKeyPairsOutput{} }
func (m *AttachKeyPairsOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachKeyPairsOutput) ProtoMessage()               {}
func (*AttachKeyPairsOutput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *AttachKeyPairsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AttachKeyPairsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AttachKeyPairsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AttachKeyPairsOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type DetachKeyPairsInput struct {
	Keypairs  []string `protobuf:"bytes,1,rep,name=keypairs" json:"keypairs,omitempty"`
	Instances []string `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
}

func (m *DetachKeyPairsInput) Reset()                    { *m = DetachKeyPairsInput{} }
func (m *DetachKeyPairsInput) String() string            { return proto.CompactTextString(m) }
func (*DetachKeyPairsInput) ProtoMessage()               {}
func (*DetachKeyPairsInput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *DetachKeyPairsInput) GetKeypairs() []string {
	if m != nil {
		return m.Keypairs
	}
	return nil
}

func (m *DetachKeyPairsInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

type DetachKeyPairsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *DetachKeyPairsOutput) Reset()                    { *m = DetachKeyPairsOutput{} }
func (m *DetachKeyPairsOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachKeyPairsOutput) ProtoMessage()               {}
func (*DetachKeyPairsOutput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *DetachKeyPairsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DetachKeyPairsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DetachKeyPairsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DetachKeyPairsOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ModifyKeyPairAttributesInput struct {
	Keypair     string `protobuf:"bytes,1,opt,name=keypair" json:"keypair,omitempty"`
	KeypairName string `protobuf:"bytes,2,opt,name=keypair_name,json=keypairName" json:"keypair_name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *ModifyKeyPairAttributesInput) Reset()                    { *m = ModifyKeyPairAttributesInput{} }
func (m *ModifyKeyPairAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyKeyPairAttributesInput) ProtoMessage()               {}
func (*ModifyKeyPairAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *ModifyKeyPairAttributesInput) GetKeypair() string {
	if m != nil {
		return m.Keypair
	}
	return ""
}

func (m *ModifyKeyPairAttributesInput) GetKeypairName() string {
	if m != nil {
		return m.KeypairName
	}
	return ""
}

func (m *ModifyKeyPairAttributesInput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ModifyKeyPairAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifyKeyPairAttributesOutput) Reset()                    { *m = ModifyKeyPairAttributesOutput{} }
func (m *ModifyKeyPairAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyKeyPairAttributesOutput) ProtoMessage()               {}
func (*ModifyKeyPairAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *ModifyKeyPairAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifyKeyPairAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifyKeyPairAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyPairServiceProperties)(nil), "service.KeyPairServiceProperties")
	proto.RegisterType((*DescribeKeyPairsInput)(nil), "service.DescribeKeyPairsInput")
	proto.RegisterType((*DescribeKeyPairsOutput)(nil), "service.DescribeKeyPairsOutput")
	proto.RegisterType((*DescribeKeyPairsOutput_ResponseItem)(nil), "service.DescribeKeyPairsOutput.ResponseItem")
	proto.RegisterType((*CreateKeyPairInput)(nil), "service.CreateKeyPairInput")
	proto.RegisterType((*CreateKeyPairOutput)(nil), "service.CreateKeyPairOutput")
	proto.RegisterType((*DeleteKeyPairsInput)(nil), "service.DeleteKeyPairsInput")
	proto.RegisterType((*DeleteKeyPairsOutput)(nil), "service.DeleteKeyPairsOutput")
	proto.RegisterType((*AttachKeyPairsInput)(nil), "service.AttachKeyPairsInput")
	proto.RegisterType((*AttachKeyPairsOutput)(nil), "service.AttachKeyPairsOutput")
	proto.RegisterType((*DetachKeyPairsInput)(nil), "service.DetachKeyPairsInput")
	proto.RegisterType((*DetachKeyPairsOutput)(nil), "service.DetachKeyPairsOutput")
	proto.RegisterType((*ModifyKeyPairAttributesInput)(nil), "service.ModifyKeyPairAttributesInput")
	proto.RegisterType((*ModifyKeyPairAttributesOutput)(nil), "service.ModifyKeyPairAttributesOutput")
}

// See https://docs.qingcloud.com/api/keypair/index.html
type KeyPairServiceInterface interface {
	DescribeKeyPairs(in *DescribeKeyPairsInput) (out *DescribeKeyPairsOutput, err error)
	CreateKeyPair(in *CreateKeyPairInput) (out *CreateKeyPairOutput, err error)
	DeleteKeyPairs(in *DeleteKeyPairsInput) (out *DeleteKeyPairsOutput, err error)
	AttachKeyPairs(in *AttachKeyPairsInput) (out *AttachKeyPairsOutput, err error)
	DetachKeyPairs(in *DetachKeyPairsInput) (out *DetachKeyPairsOutput, err error)
	ModifyKeyPairAttributes(in *ModifyKeyPairAttributesInput) (out *ModifyKeyPairAttributesOutput, err error)
}

// See https://docs.qingcloud.com/api/keypair/index.html
type KeyPairService struct {
	Config           *config.Config
	Properties       *KeyPairServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/keypair/index.html
func NewKeyPairService(conf *config.Config, zone string) (p *KeyPairService) {
	return &KeyPairService{
		Config:     conf,
		Properties: &KeyPairServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/keypair/index.html
func (s *QingCloudService) KeyPair(zone string) (*KeyPairService, error) {
	properties := &KeyPairServiceProperties{
		Zone: zone,
	}

	return &KeyPairService{Config: s.Config, Properties: properties}, nil
}

func (p *KeyPairService) DescribeKeyPairs(in *DescribeKeyPairsInput) (out *DescribeKeyPairsOutput, err error) {
	if in == nil {
		in = &DescribeKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeKeyPairs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeKeyPairsOutput{}
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

func (p *KeyPairService) CreateKeyPair(in *CreateKeyPairInput) (out *CreateKeyPairOutput, err error) {
	if in == nil {
		in = &CreateKeyPairInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateKeyPair",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateKeyPairOutput{}
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

func (p *KeyPairService) DeleteKeyPairs(in *DeleteKeyPairsInput) (out *DeleteKeyPairsOutput, err error) {
	if in == nil {
		in = &DeleteKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteKeyPairs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteKeyPairsOutput{}
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

func (p *KeyPairService) AttachKeyPairs(in *AttachKeyPairsInput) (out *AttachKeyPairsOutput, err error) {
	if in == nil {
		in = &AttachKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachKeyPairs",
		RequestMethod: "GET", // GET or POST
	}

	x := &AttachKeyPairsOutput{}
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

func (p *KeyPairService) DetachKeyPairs(in *DetachKeyPairsInput) (out *DetachKeyPairsOutput, err error) {
	if in == nil {
		in = &DetachKeyPairsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachKeyPairs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DetachKeyPairsOutput{}
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

func (p *KeyPairService) ModifyKeyPairAttributes(in *ModifyKeyPairAttributesInput) (out *ModifyKeyPairAttributesOutput, err error) {
	if in == nil {
		in = &ModifyKeyPairAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyKeyPairAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyKeyPairAttributesOutput{}
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

func (p *KeyPairServiceProperties) Validate() error {
	return nil
}

func (p *DescribeKeyPairsInput) Validate() error {
	return nil
}

func (p *DescribeKeyPairsOutput) Validate() error {
	return nil
}

func (p *CreateKeyPairInput) Validate() error {
	return nil
}

func (p *CreateKeyPairOutput) Validate() error {
	return nil
}

func (p *DeleteKeyPairsInput) Validate() error {
	return nil
}

func (p *DeleteKeyPairsOutput) Validate() error {
	return nil
}

func (p *AttachKeyPairsInput) Validate() error {
	return nil
}

func (p *AttachKeyPairsOutput) Validate() error {
	return nil
}

func (p *DetachKeyPairsInput) Validate() error {
	return nil
}

func (p *DetachKeyPairsOutput) Validate() error {
	return nil
}

func (p *ModifyKeyPairAttributesInput) Validate() error {
	return nil
}

func (p *ModifyKeyPairAttributesOutput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("key_pair.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0xad, 0x3f, 0x6b, 0x94, 0x18, 0xed, 0xda, 0x71, 0x58, 0xd6, 0xaa, 0x55, 0x02, 0x29,
	0x7c, 0x28, 0x28, 0x24, 0xbd, 0xf5, 0x16, 0xd8, 0x28, 0xa0, 0x18, 0x6e, 0x52, 0xfa, 0xd0, 0x23,
	0xc1, 0x9f, 0xb1, 0xb5, 0x31, 0xc9, 0x65, 0x77, 0x97, 0x4e, 0x19, 0xa0, 0xe8, 0xbd, 0x87, 0x02,
	0x7d, 0x85, 0xde, 0xfb, 0x36, 0xed, 0x83, 0xf4, 0x0d, 0x0a, 0x2e, 0x97, 0x36, 0x29, 0x53, 0x8d,
	0x0f, 0x56, 0x2e, 0x02, 0x67, 0x86, 0xf3, 0xed, 0x37, 0xdf, 0xcc, 0x8e, 0x08, 0x3b, 0x57, 0x58,
	0x78, 0x99, 0x4f, 0xb9, 0x93, 0x71, 0x26, 0x19, 0x19, 0x09, 0xe4, 0xd7, 0x34, 0x44, 0x6b, 0xfa,
	0x13, 0x4d, 0x2f, 0xc3, 0x98, 0xe5, 0x91, 0x27, 0xa2, 0x2b, 0x8f, 0xe7, 0x31, 0xce, 0xcb, 0x9f,
	0xea, 0x3d, 0xdb, 0x01, 0xf3, 0x14, 0x8b, 0x37, 0x3e, 0xe5, 0xe7, 0x55, 0xc2, 0x1b, 0xce, 0x32,
	0xe4, 0x92, 0xa2, 0x20, 0x04, 0xfa, 0xef, 0x59, 0x8a, 0xa6, 0x31, 0x33, 0x8e, 0xc6, 0xae, 0x7a,
	0xb6, 0x7f, 0xdf, 0x82, 0x27, 0x27, 0x28, 0x42, 0x4e, 0x03, 0xd4, 0x89, 0x62, 0x91, 0x66, 0xb9,
	0x24, 0x16, 0x6c, 0x5f, 0x61, 0x51, 0x52, 0x10, 0xa6, 0x31, 0xeb, 0x1d, 0x8d, 0xdd, 0x1b, 0x9b,
	0x1c, 0xc2, 0x84, 0xa6, 0x42, 0xfa, 0x69, 0x88, 0x1e, 0x8d, 0xcc, 0x2d, 0x05, 0x08, 0xb5, 0x6b,
	0x11, 0x91, 0x3d, 0x18, 0xb0, 0x77, 0x29, 0x72, 0xb3, 0xa7, 0x42, 0x95, 0x41, 0x9e, 0xc1, 0x0e,
	0xa6, 0x21, 0x2f, 0x32, 0xe9, 0x25, 0x28, 0x97, 0x2c, 0x32, 0xfb, 0x2a, 0xfc, 0x58, 0x7b, 0xcf,
	0x94, 0xb3, 0x44, 0x17, 0xe8, 0xf3, 0x70, 0xe9, 0xbd, 0x63, 0x3c, 0x32, 0x07, 0x15, 0x7a, 0xe5,
	0xfa, 0x91, 0xf1, 0xa8, 0x2c, 0x44, 0xfa, 0x97, 0xc2, 0x1c, 0x2a, 0x5a, 0xea, 0x99, 0x98, 0x30,
	0xba, 0x46, 0x1e, 0x30, 0x81, 0xe6, 0x68, 0x66, 0x1c, 0x0d, 0xdc, 0xda, 0x24, 0xfb, 0x30, 0x64,
	0x17, 0x17, 0x02, 0xa5, 0xb9, 0xad, 0x02, 0xda, 0x2a, 0x39, 0xc6, 0x34, 0xa1, 0xd2, 0x1c, 0x2b,
	0x77, 0x65, 0xd8, 0x7f, 0xf5, 0x60, 0x7f, 0x55, 0x90, 0xd7, 0xb9, 0x2c, 0x15, 0xd9, 0x87, 0xa1,
	0x1f, 0x4a, 0xca, 0x52, 0xad, 0xa0, 0xb6, 0xc8, 0x67, 0xb0, 0xcd, 0x51, 0x7a, 0x21, 0x8b, 0x50,
	0x49, 0x31, 0x70, 0x47, 0x1c, 0xe5, 0x31, 0x8b, 0xb0, 0x64, 0x95, 0xa0, 0x10, 0xfe, 0x25, 0x6a,
	0x25, 0x6a, 0x93, 0x9c, 0xc1, 0x44, 0xcb, 0xe9, 0x95, 0xd4, 0xfa, 0xb3, 0xde, 0xd1, 0xe4, 0xc5,
	0xd7, 0x8e, 0x6e, 0xb3, 0xd3, 0x4d, 0xc1, 0x71, 0x51, 0x64, 0x2c, 0x15, 0xb8, 0x90, 0x98, 0xb8,
	0xa0, 0x01, 0xce, 0x51, 0x96, 0x9a, 0x49, 0x26, 0xfd, 0xd8, 0x0b, 0x59, 0x9e, 0x4a, 0xa5, 0xd9,
	0xc0, 0x05, 0xe5, 0x3a, 0x2e, 0x3d, 0xd6, 0x3f, 0x06, 0x3c, 0x6a, 0x66, 0x93, 0x29, 0xd4, 0xf9,
	0x65, 0x0b, 0xab, 0x8a, 0xc6, 0xda, 0xb3, 0x88, 0xc8, 0x97, 0xf0, 0xa8, 0x0e, 0xa7, 0x7e, 0x82,
	0xba, 0xc7, 0x35, 0xe7, 0xef, 0xfd, 0x04, 0xc9, 0x0c, 0x26, 0x91, 0xa2, 0x99, 0x29, 0x51, 0xaa,
	0x02, 0x9b, 0xae, 0xfb, 0x36, 0xfc, 0x29, 0x8c, 0xb2, 0x3c, 0xf0, 0xae, 0xb0, 0xd0, 0xcd, 0x1e,
	0x66, 0x79, 0x70, 0x8a, 0x45, 0x49, 0xa2, 0x31, 0x67, 0x75, 0xc3, 0x27, 0xb7, 0x83, 0x26, 0xec,
	0x3f, 0x0c, 0x20, 0xc7, 0x1c, 0x7d, 0x59, 0x4b, 0x55, 0x4d, 0xef, 0x2a, 0x7d, 0xe3, 0x2e, 0x7d,
	0x02, 0xfd, 0xa4, 0x6e, 0xd9, 0xd8, 0x55, 0xcf, 0x1d, 0x84, 0x7b, 0x5d, 0x84, 0xa7, 0x00, 0x59,
	0x1e, 0xc4, 0x34, 0x54, 0x9c, 0xab, 0x9a, 0xc6, 0x95, 0xe7, 0x14, 0x0b, 0xfb, 0x4f, 0x03, 0x76,
	0x5b, 0x9c, 0x36, 0x31, 0x40, 0x87, 0x30, 0xc9, 0x38, 0xbd, 0xf6, 0x25, 0x36, 0x48, 0x80, 0x76,
	0x95, 0xe2, 0xb5, 0x1b, 0x3c, 0x58, 0x69, 0xb0, 0xfd, 0x1c, 0x76, 0x4f, 0x30, 0x46, 0x79, 0xff,
	0x6b, 0x6f, 0xff, 0x0a, 0x7b, 0xed, 0x94, 0x4d, 0xd4, 0xd5, 0x24, 0xd0, 0x5f, 0x21, 0xf0, 0x1a,
	0x76, 0x5f, 0x4a, 0xe9, 0x87, 0xcb, 0xfb, 0xaf, 0xaa, 0x03, 0x18, 0xd7, 0xe3, 0x22, 0xcc, 0x2d,
	0x15, 0xbc, 0x75, 0xd8, 0xef, 0x61, 0xaf, 0x0d, 0xb8, 0x89, 0x8a, 0x9e, 0xc0, 0xf0, 0x2d, 0x0b,
	0xca, 0x26, 0x54, 0x4d, 0x1a, 0xbc, 0x65, 0xc1, 0x22, 0x2a, 0x8b, 0x39, 0xc1, 0x07, 0x2e, 0xa6,
	0x0d, 0xf8, 0x11, 0x8b, 0xf9, 0x05, 0x0e, 0xce, 0x58, 0x44, 0x2f, 0x0a, 0x7d, 0xf6, 0x4b, 0x29,
	0x39, 0x0d, 0x72, 0x89, 0xba, 0x2a, 0x13, 0x46, 0xba, 0x0a, 0x4d, 0xa2, 0x36, 0x1f, 0x64, 0xd1,
	0xd8, 0x31, 0x4c, 0xd7, 0x1c, 0xbf, 0x01, 0x0d, 0x5e, 0xfc, 0xdd, 0x87, 0x9d, 0xf6, 0xbf, 0x2c,
	0x39, 0x87, 0x4f, 0x56, 0x57, 0x36, 0xf9, 0x62, 0xed, 0x36, 0x57, 0x9a, 0x58, 0x87, 0x1f, 0xd8,
	0xf6, 0xe4, 0x15, 0x3c, 0x6e, 0xad, 0x11, 0xf2, 0xf9, 0x4d, 0xc6, 0xdd, 0x95, 0x67, 0x1d, 0x74,
	0x07, 0x35, 0xd6, 0x19, 0xec, 0xb4, 0xef, 0x2e, 0x39, 0x68, 0x1c, 0x7f, 0x67, 0x0f, 0x58, 0xd3,
	0x35, 0xd1, 0x5b, 0xb8, 0xf6, 0xc5, 0x69, 0xc0, 0x75, 0x5c, 0xd1, 0x06, 0x5c, 0xe7, 0x7d, 0x53,
	0xec, 0xd6, 0xc0, 0x75, 0x5c, 0x92, 0x16, 0xbb, 0x0e, 0xb8, 0x25, 0x3c, 0x5d, 0x33, 0x0e, 0xe4,
	0xd9, 0x4d, 0xe6, 0xff, 0xcd, 0xab, 0xf5, 0xd5, 0x87, 0x5e, 0xab, 0x4e, 0xb2, 0x5e, 0xfd, 0xf6,
	0x6f, 0xff, 0x3b, 0x78, 0xbe, 0x94, 0x32, 0x13, 0xdf, 0xce, 0xe7, 0x11, 0x0b, 0x85, 0x73, 0xf3,
	0x85, 0xe6, 0x84, 0x2c, 0x99, 0xfb, 0x19, 0x9d, 0xeb, 0x61, 0x9e, 0xd3, 0x34, 0xc2, 0x9f, 0x9d,
	0xa5, 0x4c, 0x62, 0xf2, 0xe9, 0x0f, 0x34, 0xbd, 0x3c, 0x56, 0x6f, 0x69, 0xec, 0x60, 0xa8, 0x3e,
	0xe1, 0xbe, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x10, 0x83, 0xa0, 0xfc, 0x09, 0x00, 0x00,
}
