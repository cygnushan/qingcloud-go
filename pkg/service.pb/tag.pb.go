// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tag.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/service.pb/qingcloud_sdk_rule"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TagServiceProperties struct {
	Zone             *string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TagServiceProperties) Reset()                    { *m = TagServiceProperties{} }
func (m *TagServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*TagServiceProperties) ProtoMessage()               {}
func (*TagServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *TagServiceProperties) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

type DescribeTagsInput struct {
	Tags             []string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	SearchWord       *string  `protobuf:"bytes,2,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Verbose          *int32   `protobuf:"varint,3,opt,name=verbose" json:"verbose,omitempty"`
	Offset           *int32   `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Limit            *int32   `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DescribeTagsInput) Reset()                    { *m = DescribeTagsInput{} }
func (m *DescribeTagsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeTagsInput) ProtoMessage()               {}
func (*DescribeTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

func (m *DescribeTagsInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeTagsInput) GetSearchWord() string {
	if m != nil && m.SearchWord != nil {
		return *m.SearchWord
	}
	return ""
}

func (m *DescribeTagsInput) GetVerbose() int32 {
	if m != nil && m.Verbose != nil {
		return *m.Verbose
	}
	return 0
}

func (m *DescribeTagsInput) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeTagsInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type DescribeTagsOutput struct {
	Action           *string                            `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32                             `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string                            `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	TagSet           []*DescribeTagsOutput_ResponseItem `protobuf:"bytes,4,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	TotalCount       *int32                             `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *DescribeTagsOutput) Reset()                    { *m = DescribeTagsOutput{} }
func (m *DescribeTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeTagsOutput) ProtoMessage()               {}
func (*DescribeTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{2} }

func (m *DescribeTagsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeTagsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeTagsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeTagsOutput) GetTagSet() []*DescribeTagsOutput_ResponseItem {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *DescribeTagsOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type DescribeTagsOutput_ResponseItem struct {
	TagId             *string  `protobuf:"bytes,1,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	TagName           *string  `protobuf:"bytes,2,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
	Description       *string  `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	ResourceCount     *int32   `protobuf:"varint,4,opt,name=resource_count,json=resourceCount" json:"resource_count,omitempty"`
	ResourceTypeCount *int32   `protobuf:"varint,5,opt,name=resource_type_count,json=resourceTypeCount" json:"resource_type_count,omitempty"`
	ResourceTagPairs  []string `protobuf:"bytes,6,rep,name=resource_tag_pairs,json=resourceTagPairs" json:"resource_tag_pairs,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *DescribeTagsOutput_ResponseItem) Reset()         { *m = DescribeTagsOutput_ResponseItem{} }
func (m *DescribeTagsOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeTagsOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeTagsOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor26, []int{2, 0}
}

func (m *DescribeTagsOutput_ResponseItem) GetTagId() string {
	if m != nil && m.TagId != nil {
		return *m.TagId
	}
	return ""
}

func (m *DescribeTagsOutput_ResponseItem) GetTagName() string {
	if m != nil && m.TagName != nil {
		return *m.TagName
	}
	return ""
}

func (m *DescribeTagsOutput_ResponseItem) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *DescribeTagsOutput_ResponseItem) GetResourceCount() int32 {
	if m != nil && m.ResourceCount != nil {
		return *m.ResourceCount
	}
	return 0
}

func (m *DescribeTagsOutput_ResponseItem) GetResourceTypeCount() int32 {
	if m != nil && m.ResourceTypeCount != nil {
		return *m.ResourceTypeCount
	}
	return 0
}

func (m *DescribeTagsOutput_ResponseItem) GetResourceTagPairs() []string {
	if m != nil {
		return m.ResourceTagPairs
	}
	return nil
}

type CreateTagInput struct {
	TagName          *string `protobuf:"bytes,1,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateTagInput) Reset()                    { *m = CreateTagInput{} }
func (m *CreateTagInput) String() string            { return proto.CompactTextString(m) }
func (*CreateTagInput) ProtoMessage()               {}
func (*CreateTagInput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{3} }

func (m *CreateTagInput) GetTagName() string {
	if m != nil && m.TagName != nil {
		return *m.TagName
	}
	return ""
}

type CreateTagOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	TagId            *string `protobuf:"bytes,4,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateTagOutput) Reset()                    { *m = CreateTagOutput{} }
func (m *CreateTagOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateTagOutput) ProtoMessage()               {}
func (*CreateTagOutput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{4} }

func (m *CreateTagOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *CreateTagOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *CreateTagOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *CreateTagOutput) GetTagId() string {
	if m != nil && m.TagId != nil {
		return *m.TagId
	}
	return ""
}

type DeleteTagsInput struct {
	Tags             []string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeleteTagsInput) Reset()                    { *m = DeleteTagsInput{} }
func (m *DeleteTagsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteTagsInput) ProtoMessage()               {}
func (*DeleteTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{5} }

func (m *DeleteTagsInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DeleteTagsOutput struct {
	Action           *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Tags             []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeleteTagsOutput) Reset()                    { *m = DeleteTagsOutput{} }
func (m *DeleteTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteTagsOutput) ProtoMessage()               {}
func (*DeleteTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{6} }

func (m *DeleteTagsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DeleteTagsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DeleteTagsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DeleteTagsOutput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ModifyTagAttributesInput struct {
	Tag              *string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	TagName          *string `protobuf:"bytes,2,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifyTagAttributesInput) Reset()                    { *m = ModifyTagAttributesInput{} }
func (m *ModifyTagAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyTagAttributesInput) ProtoMessage()               {}
func (*ModifyTagAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{7} }

func (m *ModifyTagAttributesInput) GetTag() string {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return ""
}

func (m *ModifyTagAttributesInput) GetTagName() string {
	if m != nil && m.TagName != nil {
		return *m.TagName
	}
	return ""
}

func (m *ModifyTagAttributesInput) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

type ModifyTagAttributesOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifyTagAttributesOutput) Reset()                    { *m = ModifyTagAttributesOutput{} }
func (m *ModifyTagAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyTagAttributesOutput) ProtoMessage()               {}
func (*ModifyTagAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{8} }

func (m *ModifyTagAttributesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ModifyTagAttributesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ModifyTagAttributesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type AttachTagsInput struct {
	ResourceTagPairs []*ResourceTagPair `protobuf:"bytes,1,rep,name=resource_tag_pairs,json=resourceTagPairs" json:"resource_tag_pairs,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *AttachTagsInput) Reset()                    { *m = AttachTagsInput{} }
func (m *AttachTagsInput) String() string            { return proto.CompactTextString(m) }
func (*AttachTagsInput) ProtoMessage()               {}
func (*AttachTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{9} }

func (m *AttachTagsInput) GetResourceTagPairs() []*ResourceTagPair {
	if m != nil {
		return m.ResourceTagPairs
	}
	return nil
}

type AttachTagsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AttachTagsOutput) Reset()                    { *m = AttachTagsOutput{} }
func (m *AttachTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachTagsOutput) ProtoMessage()               {}
func (*AttachTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{10} }

func (m *AttachTagsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AttachTagsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AttachTagsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type DetachTagsInput struct {
	ResourceTagPairs []*ResourceTagPair `protobuf:"bytes,1,rep,name=resource_tag_pairs,json=resourceTagPairs" json:"resource_tag_pairs,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *DetachTagsInput) Reset()                    { *m = DetachTagsInput{} }
func (m *DetachTagsInput) String() string            { return proto.CompactTextString(m) }
func (*DetachTagsInput) ProtoMessage()               {}
func (*DetachTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{11} }

func (m *DetachTagsInput) GetResourceTagPairs() []*ResourceTagPair {
	if m != nil {
		return m.ResourceTagPairs
	}
	return nil
}

type DetachTagsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DetachTagsOutput) Reset()                    { *m = DetachTagsOutput{} }
func (m *DetachTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachTagsOutput) ProtoMessage()               {}
func (*DetachTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{12} }

func (m *DetachTagsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DetachTagsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DetachTagsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TagServiceProperties)(nil), "service.TagServiceProperties")
	proto.RegisterType((*DescribeTagsInput)(nil), "service.DescribeTagsInput")
	proto.RegisterType((*DescribeTagsOutput)(nil), "service.DescribeTagsOutput")
	proto.RegisterType((*DescribeTagsOutput_ResponseItem)(nil), "service.DescribeTagsOutput.ResponseItem")
	proto.RegisterType((*CreateTagInput)(nil), "service.CreateTagInput")
	proto.RegisterType((*CreateTagOutput)(nil), "service.CreateTagOutput")
	proto.RegisterType((*DeleteTagsInput)(nil), "service.DeleteTagsInput")
	proto.RegisterType((*DeleteTagsOutput)(nil), "service.DeleteTagsOutput")
	proto.RegisterType((*ModifyTagAttributesInput)(nil), "service.ModifyTagAttributesInput")
	proto.RegisterType((*ModifyTagAttributesOutput)(nil), "service.ModifyTagAttributesOutput")
	proto.RegisterType((*AttachTagsInput)(nil), "service.AttachTagsInput")
	proto.RegisterType((*AttachTagsOutput)(nil), "service.AttachTagsOutput")
	proto.RegisterType((*DetachTagsInput)(nil), "service.DetachTagsInput")
	proto.RegisterType((*DetachTagsOutput)(nil), "service.DetachTagsOutput")
}

func init() { proto.RegisterFile("tag.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0xdf, 0x7c, 0xbd, 0x99, 0xf4, 0x23, 0xdd, 0xf6, 0xed, 0xeb, 0x1a, 0x21, 0x82, 0xa5,
	0x4a, 0x11, 0x1f, 0x8e, 0xd4, 0x23, 0x87, 0x4a, 0x21, 0x15, 0xa8, 0x07, 0xa0, 0x84, 0x48, 0x08,
	0x09, 0xc9, 0xda, 0xda, 0x53, 0xc7, 0x22, 0xf1, 0x9a, 0xdd, 0x71, 0xa1, 0x9c, 0x38, 0x73, 0xe2,
	0x9f, 0xf0, 0xb3, 0x38, 0xf3, 0x0f, 0xd0, 0xda, 0x8e, 0xed, 0xb4, 0xa1, 0x17, 0x94, 0x4b, 0xb4,
	0x3b, 0xfb, 0xec, 0xcc, 0xf3, 0x4c, 0x9e, 0x59, 0x43, 0x9b, 0x78, 0xe0, 0xc4, 0x52, 0x90, 0x60,
	0x2d, 0x85, 0xf2, 0x32, 0xf4, 0xd0, 0xea, 0xd0, 0x55, 0x8c, 0x2a, 0x8b, 0x5a, 0x77, 0x3f, 0x86,
	0x51, 0xe0, 0xcd, 0x44, 0xe2, 0xbb, 0xca, 0xff, 0xe0, 0xca, 0x64, 0x86, 0x03, 0xfd, 0x93, 0x1d,
	0xdb, 0x0f, 0x60, 0x6f, 0xc2, 0x83, 0x37, 0xd9, 0xcd, 0x33, 0x29, 0x62, 0x94, 0x14, 0xa2, 0x62,
	0x0c, 0xea, 0x5f, 0x44, 0x84, 0xa6, 0xd1, 0x33, 0xfa, 0xed, 0x71, 0xba, 0xb6, 0xbf, 0x1b, 0xb0,
	0x73, 0x82, 0xca, 0x93, 0xe1, 0x39, 0x4e, 0x78, 0xa0, 0x4e, 0xa3, 0x38, 0x21, 0x8d, 0x24, 0x1e,
	0x28, 0xd3, 0xe8, 0xd5, 0x34, 0x52, 0xaf, 0xd9, 0x3d, 0xe8, 0x28, 0xe4, 0xd2, 0x9b, 0xba, 0x9f,
	0x84, 0xf4, 0xcd, 0x7f, 0xd2, 0x24, 0x90, 0x85, 0xde, 0x0a, 0xe9, 0x33, 0x13, 0x5a, 0x97, 0x28,
	0xcf, 0x85, 0x42, 0xb3, 0xd6, 0x33, 0xfa, 0x8d, 0xf1, 0x62, 0xcb, 0xf6, 0xa1, 0x29, 0x2e, 0x2e,
	0x14, 0x92, 0x59, 0x4f, 0x0f, 0xf2, 0x1d, 0xdb, 0x83, 0xc6, 0x2c, 0x9c, 0x87, 0x64, 0x36, 0xd2,
	0x70, 0xb6, 0xb1, 0x7f, 0xd4, 0x80, 0x55, 0x29, 0xbd, 0x4a, 0x48, 0x73, 0xda, 0x87, 0x26, 0xf7,
	0x28, 0x14, 0x51, 0xce, 0x3f, 0xdf, 0xb1, 0x03, 0xf8, 0x57, 0x22, 0xb9, 0x9e, 0xf0, 0x31, 0x25,
	0xd5, 0x18, 0xb7, 0x24, 0xd2, 0x48, 0xf8, 0xa8, 0x19, 0xcd, 0x51, 0x29, 0x1e, 0x64, 0x8c, 0xda,
	0xe3, 0xc5, 0x96, 0x0d, 0xa1, 0x45, 0x3c, 0x70, 0x33, 0x4a, 0xb5, 0x7e, 0xe7, 0xa8, 0xef, 0xe4,
	0x9d, 0x76, 0x6e, 0x96, 0x76, 0xc6, 0xa8, 0x62, 0x11, 0x29, 0x3c, 0x25, 0x9c, 0x8f, 0x9b, 0xa4,
	0x7b, 0x4b, 0xba, 0x1f, 0x24, 0x88, 0xcf, 0x5c, 0x4f, 0x24, 0xd1, 0x42, 0x02, 0xa4, 0xa1, 0x91,
	0x8e, 0x58, 0x3f, 0x0d, 0xd8, 0xa8, 0xde, 0x64, 0xff, 0x81, 0xbe, 0xeb, 0x86, 0x7e, 0xae, 0xa0,
	0x41, 0x3c, 0x38, 0xf5, 0xb5, 0x00, 0x1d, 0x8e, 0xf8, 0x1c, 0xf3, 0xae, 0x6a, 0x6e, 0x2f, 0xf9,
	0x1c, 0x59, 0x0f, 0x3a, 0x7e, 0x4a, 0x27, 0x4e, 0x85, 0x67, 0x22, 0xaa, 0x21, 0x76, 0x08, 0x5b,
	0x12, 0x95, 0x48, 0xa4, 0x87, 0x39, 0x91, 0xac, 0xc5, 0x9b, 0x8b, 0x68, 0xca, 0x85, 0x39, 0xb0,
	0x5b, 0xc0, 0xb4, 0x93, 0x96, 0x48, 0xef, 0x2c, 0x8e, 0x26, 0x57, 0x71, 0x8e, 0x7f, 0x04, 0xac,
	0xc4, 0xf3, 0xc0, 0x8d, 0x79, 0x28, 0x95, 0xd9, 0x4c, 0xed, 0xd0, 0x2d, 0xe0, 0x3c, 0x38, 0xd3,
	0x71, 0xfb, 0x21, 0x6c, 0x8d, 0x24, 0x72, 0xd2, 0x91, 0xcc, 0x40, 0x55, 0x4d, 0xc6, 0x92, 0x26,
	0x3b, 0x81, 0xed, 0x02, 0xbc, 0x8e, 0xbf, 0xb6, 0xec, 0x72, 0xbd, 0xd2, 0x65, 0xfb, 0x10, 0xb6,
	0x4f, 0x70, 0x86, 0x74, 0xbb, 0xcb, 0x6d, 0x05, 0xdd, 0x12, 0xb6, 0x0e, 0x7a, 0x8b, 0xa2, 0xf5,
	0x4a, 0xd1, 0x10, 0xcc, 0x17, 0xc2, 0x0f, 0x2f, 0xae, 0x26, 0x3c, 0x18, 0x12, 0xc9, 0xf0, 0x3c,
	0x21, 0xcc, 0x49, 0x76, 0xa1, 0x46, 0x3c, 0xc8, 0x2b, 0xeb, 0xe5, 0x5f, 0xf9, 0xc5, 0x9e, 0xc2,
	0xc1, 0x8a, 0x52, 0x6b, 0x10, 0x6a, 0xbf, 0x83, 0xed, 0x21, 0x11, 0xf7, 0xa6, 0x65, 0xc3, 0x9f,
	0xad, 0x74, 0x95, 0x91, 0x0e, 0xa0, 0x59, 0x0c, 0xe0, 0x78, 0xd9, 0x5e, 0x2b, 0xfc, 0xe6, 0x42,
	0xb7, 0x4c, 0xbd, 0x26, 0xee, 0x27, 0xb8, 0x36, 0xee, 0x65, 0xea, 0x35, 0x70, 0x3f, 0xfa, 0x5a,
	0x07, 0x28, 0x9f, 0x7f, 0xf6, 0x1c, 0x36, 0xaa, 0x2f, 0x1a, 0xb3, 0x56, 0x3e, 0x74, 0xa9, 0x46,
	0xeb, 0xce, 0x2d, 0x8f, 0x20, 0x3b, 0x86, 0x76, 0x31, 0xb7, 0xec, 0xff, 0x02, 0xb9, 0x3c, 0xf8,
	0x96, 0x79, 0xf3, 0x20, 0xbf, 0x3f, 0x04, 0x28, 0x27, 0x8b, 0x99, 0x95, 0x52, 0x4b, 0x53, 0x69,
	0x1d, 0xac, 0x38, 0xc9, 0x53, 0xbc, 0x87, 0xdd, 0x15, 0xe6, 0x65, 0xf7, 0x8b, 0x1b, 0x7f, 0x9a,
	0x22, 0xcb, 0xbe, 0x0d, 0x52, 0x12, 0x2c, 0x5d, 0x55, 0x21, 0x78, 0xcd, 0xc5, 0x15, 0x82, 0x37,
	0x4c, 0x98, 0x6a, 0x5c, 0x91, 0xe2, 0x9a, 0x99, 0x96, 0x34, 0x2e, 0xa7, 0xb0, 0x9e, 0x7e, 0xfb,
	0x55, 0x3f, 0x86, 0xc7, 0x53, 0xa2, 0x58, 0x3d, 0x19, 0x0c, 0x7c, 0xe1, 0x29, 0xa7, 0xf8, 0xdc,
	0x3b, 0x9e, 0x98, 0x0f, 0x78, 0x1c, 0x0e, 0x88, 0x07, 0x83, 0x30, 0xf2, 0xf1, 0xb3, 0x33, 0xa5,
	0xf9, 0x8c, 0x6d, 0xbe, 0x0e, 0xa3, 0x60, 0x94, 0x22, 0x26, 0x3c, 0xf8, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x55, 0x0a, 0x32, 0xcf, 0x41, 0x08, 0x00, 0x00,
}
