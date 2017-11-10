// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_data.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/service.pb/qingcloud_sdk_rule"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserDataServiceProperties struct {
	Zone             *string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserDataServiceProperties) Reset()                    { *m = UserDataServiceProperties{} }
func (m *UserDataServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*UserDataServiceProperties) ProtoMessage()               {}
func (*UserDataServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

func (m *UserDataServiceProperties) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

type UploadUserDataAttachmentInput struct {
	AttachmentContent []byte  `protobuf:"bytes,2,opt,name=attachment_content,json=attachmentContent" json:"attachment_content,omitempty"`
	AttachmentName    *string `protobuf:"bytes,1,opt,name=attachment_name,json=attachmentName" json:"attachment_name,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *UploadUserDataAttachmentInput) Reset()                    { *m = UploadUserDataAttachmentInput{} }
func (m *UploadUserDataAttachmentInput) String() string            { return proto.CompactTextString(m) }
func (*UploadUserDataAttachmentInput) ProtoMessage()               {}
func (*UploadUserDataAttachmentInput) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{1} }

func (m *UploadUserDataAttachmentInput) GetAttachmentContent() []byte {
	if m != nil {
		return m.AttachmentContent
	}
	return nil
}

func (m *UploadUserDataAttachmentInput) GetAttachmentName() string {
	if m != nil && m.AttachmentName != nil {
		return *m.AttachmentName
	}
	return ""
}

type UploadUserDataAttachmentOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	AttachmentId     *string `protobuf:"bytes,4,opt,name=attachment_id,json=attachmentId" json:"attachment_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UploadUserDataAttachmentOutput) Reset()                    { *m = UploadUserDataAttachmentOutput{} }
func (m *UploadUserDataAttachmentOutput) String() string            { return proto.CompactTextString(m) }
func (*UploadUserDataAttachmentOutput) ProtoMessage()               {}
func (*UploadUserDataAttachmentOutput) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{2} }

func (m *UploadUserDataAttachmentOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *UploadUserDataAttachmentOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *UploadUserDataAttachmentOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *UploadUserDataAttachmentOutput) GetAttachmentId() string {
	if m != nil && m.AttachmentId != nil {
		return *m.AttachmentId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDataServiceProperties)(nil), "service.UserDataServiceProperties")
	proto.RegisterType((*UploadUserDataAttachmentInput)(nil), "service.UploadUserDataAttachmentInput")
	proto.RegisterType((*UploadUserDataAttachmentOutput)(nil), "service.UploadUserDataAttachmentOutput")
}

func init() { proto.RegisterFile("user_data.proto", fileDescriptor28) }

var fileDescriptor28 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0xc5, 0xad, 0x0b, 0xed, 0x8a, 0x16, 0x75, 0x0f, 0x95, 0xb1, 0x4a, 0x85, 0x5c, 0xa9, 0x70,
	0xa9, 0x2d, 0x71, 0xe4, 0x16, 0x91, 0x43, 0x90, 0xa2, 0x40, 0x20, 0x9c, 0xad, 0x95, 0x77, 0x04,
	0xab, 0xd8, 0xbb, 0xce, 0xee, 0x38, 0x89, 0xf2, 0x09, 0x1c, 0x23, 0xe5, 0xdf, 0xf2, 0x1b, 0xfc,
	0x41, 0x64, 0x1b, 0x03, 0x8a, 0x42, 0x72, 0x59, 0xed, 0xcc, 0xbc, 0x99, 0xf7, 0x76, 0xf6, 0x91,
	0x56, 0x66, 0x40, 0x87, 0x9c, 0x21, 0xf3, 0x53, 0xad, 0x50, 0xd1, 0x86, 0x01, 0x7d, 0x2b, 0x22,
	0x70, 0x3b, 0x37, 0x42, 0x2e, 0xa3, 0x58, 0x65, 0x3c, 0x34, 0xfc, 0x3a, 0xd4, 0x59, 0x0c, 0x41,
	0x7e, 0x94, 0x38, 0x2f, 0x20, 0xed, 0x85, 0x01, 0x7d, 0xca, 0x90, 0xcd, 0xcb, 0x8e, 0xa9, 0x56,
	0x29, 0x68, 0x14, 0x60, 0x28, 0x25, 0xf6, 0x83, 0x92, 0xe0, 0x58, 0x5d, 0xab, 0xff, 0x6d, 0x56,
	0xdc, 0xbd, 0x27, 0x8b, 0x74, 0x16, 0x69, 0xac, 0x18, 0xaf, 0xfa, 0x4e, 0x10, 0x59, 0xb4, 0x4a,
	0x40, 0xe2, 0x58, 0xa6, 0x19, 0xd2, 0xff, 0x84, 0xb2, 0x5d, 0x2a, 0x8c, 0x94, 0x44, 0x90, 0xe8,
	0x7c, 0xea, 0x5a, 0xfd, 0xe6, 0xec, 0xe7, 0xbe, 0x32, 0x2a, 0x0b, 0xb4, 0x47, 0x5a, 0x07, 0x70,
	0xc9, 0x92, 0x8a, 0xef, 0xc7, 0x3e, 0x7d, 0xc1, 0x12, 0x18, 0xfe, 0x5e, 0x6f, 0x6c, 0xe7, 0xad,
	0xd9, 0xb4, 0xe6, 0xd6, 0xbc, 0x47, 0x8b, 0xfc, 0x39, 0xa6, 0x6b, 0x92, 0x61, 0x2e, 0xec, 0x17,
	0xa9, 0xb3, 0x08, 0x85, 0x92, 0x5b, 0x82, 0x6d, 0x44, 0xdb, 0xe4, 0xab, 0x86, 0x7c, 0x1a, 0x87,
	0x42, 0xe6, 0x97, 0x59, 0x43, 0x03, 0x8e, 0x14, 0x07, 0xea, 0x90, 0x46, 0x02, 0xc6, 0xb0, 0x25,
	0x38, 0x9f, 0x8b, 0x9e, 0x2a, 0xa4, 0x7f, 0xc9, 0xf7, 0x03, 0x25, 0x82, 0x3b, 0x76, 0x51, 0x6f,
	0xee, 0x93, 0x63, 0x3e, 0x78, 0xb6, 0x48, 0xeb, 0xd5, 0x7a, 0xe9, 0x1d, 0x71, 0x8e, 0xe9, 0xa4,
	0xff, 0xfc, 0xed, 0xb7, 0xf9, 0xef, 0xae, 0xd8, 0xed, 0x7d, 0x88, 0x2b, 0x9f, 0xec, 0x91, 0xf5,
	0xc6, 0xae, 0x53, 0x7b, 0x3a, 0x99, 0x5f, 0xb9, 0xe7, 0xeb, 0x8d, 0x7d, 0x46, 0x06, 0x2b, 0xc4,
	0xd4, 0x0c, 0x83, 0x80, 0xab, 0xc8, 0xf8, 0x3b, 0x73, 0xf8, 0x91, 0x4a, 0x02, 0x96, 0x8a, 0x20,
	0x37, 0x52, 0xee, 0xa3, 0x40, 0x48, 0x0e, 0xf7, 0xfe, 0x0a, 0x93, 0x98, 0xd2, 0x4b, 0x21, 0x97,
	0xa3, 0x02, 0x56, 0x31, 0xbd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x24, 0x3f, 0x41, 0x72, 0x02,
	0x00, 0x00,
}
