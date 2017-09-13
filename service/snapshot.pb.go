// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snapshot.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type SnapshotServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SnapshotServiceProperties) Reset()                    { *m = SnapshotServiceProperties{} }
func (m *SnapshotServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SnapshotServiceProperties) ProtoMessage()               {}
func (*SnapshotServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *SnapshotServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeSnapshotsInput struct {
	Snapshots    []string `protobuf:"bytes,1,rep,name=snapshots" json:"snapshots,omitempty"`
	ResourceId   string   `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	SnapshotType int32    `protobuf:"varint,3,opt,name=snapshot_type,json=snapshotType" json:"snapshot_type,omitempty"`
	Status       []string `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	SearchWord   string   `protobuf:"bytes,5,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Tags         []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Verbose      int32    `protobuf:"varint,7,opt,name=verbose" json:"verbose,omitempty"`
	Offset       int32    `protobuf:"varint,8,opt,name=offset" json:"offset,omitempty"`
	Limit        int32    `protobuf:"varint,9,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeSnapshotsInput) Reset()                    { *m = DescribeSnapshotsInput{} }
func (m *DescribeSnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSnapshotsInput) ProtoMessage()               {}
func (*DescribeSnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

func (m *DescribeSnapshotsInput) GetSnapshots() []string {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func (m *DescribeSnapshotsInput) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *DescribeSnapshotsInput) GetSnapshotType() int32 {
	if m != nil {
		return m.SnapshotType
	}
	return 0
}

func (m *DescribeSnapshotsInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeSnapshotsInput) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *DescribeSnapshotsInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeSnapshotsInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

func (m *DescribeSnapshotsInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeSnapshotsInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeSnapshotsOutput struct {
	Action      string                                  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode     int32                                   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message     string                                  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SnapshotSet []*DescribeSnapshotsOutput_ResponseItem `protobuf:"bytes,4,rep,name=snapshot_set,json=snapshotSet" json:"snapshot_set,omitempty"`
}

func (m *DescribeSnapshotsOutput) Reset()                    { *m = DescribeSnapshotsOutput{} }
func (m *DescribeSnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSnapshotsOutput) ProtoMessage()               {}
func (*DescribeSnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func (m *DescribeSnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeSnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeSnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeSnapshotsOutput) GetSnapshotSet() []*DescribeSnapshotsOutput_ResponseItem {
	if m != nil {
		return m.SnapshotSet
	}
	return nil
}

type DescribeSnapshotsOutput_ResponseItem struct {
	SnapshotId          string                      `protobuf:"bytes,1,opt,name=snapshot_id,json=snapshotId" json:"snapshot_id,omitempty"`
	SnapshotName        string                      `protobuf:"bytes,2,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	Description         string                      `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	SnapshotType        string                      `protobuf:"bytes,4,opt,name=snapshot_type,json=snapshotType" json:"snapshot_type,omitempty"`
	Status              string                      `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	TransitionStatus    string                      `protobuf:"bytes,6,opt,name=transition_status,json=transitionStatus" json:"transition_status,omitempty"`
	CreateTime          *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime          *google_protobuf1.Timestamp `protobuf:"bytes,8,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	SnapshotTime        *google_protobuf1.Timestamp `protobuf:"bytes,9,opt,name=snapshot_time,json=snapshotTime" json:"snapshot_time,omitempty"`
	IsTaken             int32                       `protobuf:"varint,10,opt,name=is_taken,json=isTaken" json:"is_taken,omitempty"`
	IsHead              int32                       `protobuf:"varint,11,opt,name=is_head,json=isHead" json:"is_head,omitempty"`
	RootId              string                      `protobuf:"bytes,12,opt,name=root_id,json=rootId" json:"root_id,omitempty"`
	ParentId            string                      `protobuf:"bytes,13,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	Size                int32                       `protobuf:"varint,14,opt,name=size" json:"size,omitempty"`
	TotalSize           int32                       `protobuf:"varint,15,opt,name=total_size,json=totalSize" json:"total_size,omitempty"`
	TotalCount          int32                       `protobuf:"varint,16,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	LastestSnapshotTime *google_protobuf1.Timestamp `protobuf:"bytes,17,opt,name=lastest_snapshot_time,json=lastestSnapshotTime" json:"lastest_snapshot_time,omitempty"`
}

func (m *DescribeSnapshotsOutput_ResponseItem) Reset()         { *m = DescribeSnapshotsOutput_ResponseItem{} }
func (m *DescribeSnapshotsOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeSnapshotsOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeSnapshotsOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{2, 0}
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetSnapshotType() string {
	if m != nil {
		return m.SnapshotType
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetTransitionStatus() string {
	if m != nil {
		return m.TransitionStatus
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetSnapshotTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.SnapshotTime
	}
	return nil
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetIsTaken() int32 {
	if m != nil {
		return m.IsTaken
	}
	return 0
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetIsHead() int32 {
	if m != nil {
		return m.IsHead
	}
	return 0
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetRootId() string {
	if m != nil {
		return m.RootId
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeSnapshotsOutput_ResponseItem) GetLastestSnapshotTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastestSnapshotTime
	}
	return nil
}

type CreateSnapshotsInput struct {
	Resources    []string `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
	SnapshotName string   `protobuf:"bytes,2,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	IsFull       int32    `protobuf:"varint,3,opt,name=is_full,json=isFull" json:"is_full,omitempty"`
}

func (m *CreateSnapshotsInput) Reset()                    { *m = CreateSnapshotsInput{} }
func (m *CreateSnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotsInput) ProtoMessage()               {}
func (*CreateSnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{3} }

func (m *CreateSnapshotsInput) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *CreateSnapshotsInput) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *CreateSnapshotsInput) GetIsFull() int32 {
	if m != nil {
		return m.IsFull
	}
	return 0
}

type CreateSnapshotsOutput struct {
	Action    string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode   int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message   string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId     string   `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Snapshots []string `protobuf:"bytes,5,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *CreateSnapshotsOutput) Reset()                    { *m = CreateSnapshotsOutput{} }
func (m *CreateSnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotsOutput) ProtoMessage()               {}
func (*CreateSnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{4} }

func (m *CreateSnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateSnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateSnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateSnapshotsOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *CreateSnapshotsOutput) GetSnapshots() []string {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type DeleteSnapshotsInput struct {
	Snapshots []string `protobuf:"bytes,1,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *DeleteSnapshotsInput) Reset()                    { *m = DeleteSnapshotsInput{} }
func (m *DeleteSnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSnapshotsInput) ProtoMessage()               {}
func (*DeleteSnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{5} }

func (m *DeleteSnapshotsInput) GetSnapshots() []string {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type DeleteSnapshotsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *DeleteSnapshotsOutput) Reset()                    { *m = DeleteSnapshotsOutput{} }
func (m *DeleteSnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSnapshotsOutput) ProtoMessage()               {}
func (*DeleteSnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{6} }

func (m *DeleteSnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteSnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteSnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteSnapshotsOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ApplySnapshotsInput struct {
	Snapshots []string `protobuf:"bytes,1,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *ApplySnapshotsInput) Reset()                    { *m = ApplySnapshotsInput{} }
func (m *ApplySnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*ApplySnapshotsInput) ProtoMessage()               {}
func (*ApplySnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{7} }

func (m *ApplySnapshotsInput) GetSnapshots() []string {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type ApplySnapshotsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *ApplySnapshotsOutput) Reset()                    { *m = ApplySnapshotsOutput{} }
func (m *ApplySnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*ApplySnapshotsOutput) ProtoMessage()               {}
func (*ApplySnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{8} }

func (m *ApplySnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ApplySnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ApplySnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApplySnapshotsOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ModifySnapshotAttributesInput struct {
	Snapshot     string `protobuf:"bytes,1,opt,name=snapshot" json:"snapshot,omitempty"`
	SnapshotName string `protobuf:"bytes,2,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *ModifySnapshotAttributesInput) Reset()                    { *m = ModifySnapshotAttributesInput{} }
func (m *ModifySnapshotAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySnapshotAttributesInput) ProtoMessage()               {}
func (*ModifySnapshotAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{9} }

func (m *ModifySnapshotAttributesInput) GetSnapshot() string {
	if m != nil {
		return m.Snapshot
	}
	return ""
}

func (m *ModifySnapshotAttributesInput) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *ModifySnapshotAttributesInput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ModifySnapshotAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifySnapshotAttributesOutput) Reset()         { *m = ModifySnapshotAttributesOutput{} }
func (m *ModifySnapshotAttributesOutput) String() string { return proto.CompactTextString(m) }
func (*ModifySnapshotAttributesOutput) ProtoMessage()    {}
func (*ModifySnapshotAttributesOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{10}
}

func (m *ModifySnapshotAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifySnapshotAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifySnapshotAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CaptureInstanceFromSnapshotInput struct {
	Snapshot  string `protobuf:"bytes,1,opt,name=snapshot" json:"snapshot,omitempty"`
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
}

func (m *CaptureInstanceFromSnapshotInput) Reset()         { *m = CaptureInstanceFromSnapshotInput{} }
func (m *CaptureInstanceFromSnapshotInput) String() string { return proto.CompactTextString(m) }
func (*CaptureInstanceFromSnapshotInput) ProtoMessage()    {}
func (*CaptureInstanceFromSnapshotInput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{11}
}

func (m *CaptureInstanceFromSnapshotInput) GetSnapshot() string {
	if m != nil {
		return m.Snapshot
	}
	return ""
}

func (m *CaptureInstanceFromSnapshotInput) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

type CaptureInstanceFromSnapshotOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	ImageId string `protobuf:"bytes,5,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
}

func (m *CaptureInstanceFromSnapshotOutput) Reset()         { *m = CaptureInstanceFromSnapshotOutput{} }
func (m *CaptureInstanceFromSnapshotOutput) String() string { return proto.CompactTextString(m) }
func (*CaptureInstanceFromSnapshotOutput) ProtoMessage()    {}
func (*CaptureInstanceFromSnapshotOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{12}
}

func (m *CaptureInstanceFromSnapshotOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CaptureInstanceFromSnapshotOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CaptureInstanceFromSnapshotOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CaptureInstanceFromSnapshotOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *CaptureInstanceFromSnapshotOutput) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

type CreateVolumeFromSnapshotInput struct {
	Snapshot   string `protobuf:"bytes,1,opt,name=snapshot" json:"snapshot,omitempty"`
	VolumeName string `protobuf:"bytes,2,opt,name=volume_name,json=volumeName" json:"volume_name,omitempty"`
}

func (m *CreateVolumeFromSnapshotInput) Reset()                    { *m = CreateVolumeFromSnapshotInput{} }
func (m *CreateVolumeFromSnapshotInput) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumeFromSnapshotInput) ProtoMessage()               {}
func (*CreateVolumeFromSnapshotInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{13} }

func (m *CreateVolumeFromSnapshotInput) GetSnapshot() string {
	if m != nil {
		return m.Snapshot
	}
	return ""
}

func (m *CreateVolumeFromSnapshotInput) GetVolumeName() string {
	if m != nil {
		return m.VolumeName
	}
	return ""
}

type CreateVolumeFromSnapshotOutput struct {
	Action   string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode  int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId    string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	VolumeId string `protobuf:"bytes,5,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
}

func (m *CreateVolumeFromSnapshotOutput) Reset()         { *m = CreateVolumeFromSnapshotOutput{} }
func (m *CreateVolumeFromSnapshotOutput) String() string { return proto.CompactTextString(m) }
func (*CreateVolumeFromSnapshotOutput) ProtoMessage()    {}
func (*CreateVolumeFromSnapshotOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{14}
}

func (m *CreateVolumeFromSnapshotOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateVolumeFromSnapshotOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateVolumeFromSnapshotOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateVolumeFromSnapshotOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *CreateVolumeFromSnapshotOutput) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func init() {
	proto.RegisterType((*SnapshotServiceProperties)(nil), "service.SnapshotServiceProperties")
	proto.RegisterType((*DescribeSnapshotsInput)(nil), "service.DescribeSnapshotsInput")
	proto.RegisterType((*DescribeSnapshotsOutput)(nil), "service.DescribeSnapshotsOutput")
	proto.RegisterType((*DescribeSnapshotsOutput_ResponseItem)(nil), "service.DescribeSnapshotsOutput.ResponseItem")
	proto.RegisterType((*CreateSnapshotsInput)(nil), "service.CreateSnapshotsInput")
	proto.RegisterType((*CreateSnapshotsOutput)(nil), "service.CreateSnapshotsOutput")
	proto.RegisterType((*DeleteSnapshotsInput)(nil), "service.DeleteSnapshotsInput")
	proto.RegisterType((*DeleteSnapshotsOutput)(nil), "service.DeleteSnapshotsOutput")
	proto.RegisterType((*ApplySnapshotsInput)(nil), "service.ApplySnapshotsInput")
	proto.RegisterType((*ApplySnapshotsOutput)(nil), "service.ApplySnapshotsOutput")
	proto.RegisterType((*ModifySnapshotAttributesInput)(nil), "service.ModifySnapshotAttributesInput")
	proto.RegisterType((*ModifySnapshotAttributesOutput)(nil), "service.ModifySnapshotAttributesOutput")
	proto.RegisterType((*CaptureInstanceFromSnapshotInput)(nil), "service.CaptureInstanceFromSnapshotInput")
	proto.RegisterType((*CaptureInstanceFromSnapshotOutput)(nil), "service.CaptureInstanceFromSnapshotOutput")
	proto.RegisterType((*CreateVolumeFromSnapshotInput)(nil), "service.CreateVolumeFromSnapshotInput")
	proto.RegisterType((*CreateVolumeFromSnapshotOutput)(nil), "service.CreateVolumeFromSnapshotOutput")
}

type SnapshotServiceInterface interface {
	DescribeSnapshots(in *DescribeSnapshotsInput) (out *DescribeSnapshotsOutput, err error)
	CreateSnapshots(in *CreateSnapshotsInput) (out *CreateSnapshotsOutput, err error)
	DeleteSnapshots(in *DeleteSnapshotsInput) (out *DeleteSnapshotsOutput, err error)
	ApplySnapshots(in *ApplySnapshotsInput) (out *ApplySnapshotsOutput, err error)
	ModifySnapshotAttributes(in *ModifySnapshotAttributesInput) (out *ModifySnapshotAttributesOutput, err error)
	CaptureInstanceFromSnapshot(in *CaptureInstanceFromSnapshotInput) (out *CaptureInstanceFromSnapshotOutput, err error)
	CreateVolumeFromSnapshot(in *CreateVolumeFromSnapshotInput) (out *CreateVolumeFromSnapshotOutput, err error)
}

type SnapshotService struct {
	Config           *config.Config
	Properties       *SnapshotServiceProperties
	LastResponseBody string
}

func NewSnapshotService(conf *config.Config, zone string) (p *SnapshotService) {
	return &SnapshotService{
		Config:     conf,
		Properties: &SnapshotServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Snapshot(zone string) (*SnapshotService, error) {
	properties := &SnapshotServiceProperties{
		Zone: zone,
	}

	return &SnapshotService{Config: s.Config, Properties: properties}, nil
}

func (p *SnapshotService) DescribeSnapshots(in *DescribeSnapshotsInput) (out *DescribeSnapshotsOutput, err error) {
	if in == nil {
		in = &DescribeSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeSnapshotsOutput{}
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

func (p *DescribeSnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) CreateSnapshots(in *CreateSnapshotsInput) (out *CreateSnapshotsOutput, err error) {
	if in == nil {
		in = &CreateSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateSnapshotsOutput{}
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

func (p *CreateSnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) DeleteSnapshots(in *DeleteSnapshotsInput) (out *DeleteSnapshotsOutput, err error) {
	if in == nil {
		in = &DeleteSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteSnapshotsOutput{}
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

func (p *DeleteSnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) ApplySnapshots(in *ApplySnapshotsInput) (out *ApplySnapshotsOutput, err error) {
	if in == nil {
		in = &ApplySnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplySnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &ApplySnapshotsOutput{}
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

func (p *ApplySnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) ModifySnapshotAttributes(in *ModifySnapshotAttributesInput) (out *ModifySnapshotAttributesOutput, err error) {
	if in == nil {
		in = &ModifySnapshotAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySnapshotAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifySnapshotAttributesOutput{}
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

func (p *ModifySnapshotAttributesInput) Validate() error {
	return nil
}

func (p *SnapshotService) CaptureInstanceFromSnapshot(in *CaptureInstanceFromSnapshotInput) (out *CaptureInstanceFromSnapshotOutput, err error) {
	if in == nil {
		in = &CaptureInstanceFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CaptureInstanceFromSnapshot",
		RequestMethod: "GET", // GET or POST
	}

	x := &CaptureInstanceFromSnapshotOutput{}
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

func (p *CaptureInstanceFromSnapshotInput) Validate() error {
	return nil
}

func (p *SnapshotService) CreateVolumeFromSnapshot(in *CreateVolumeFromSnapshotInput) (out *CreateVolumeFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateVolumeFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVolumeFromSnapshot",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateVolumeFromSnapshotOutput{}
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

func (p *CreateVolumeFromSnapshotInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("snapshot.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 1059 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0x97, 0xdb, 0xf8, 0xdf, 0x73, 0x9b, 0x34, 0xd3, 0x24, 0xdd, 0x6c, 0xea, 0xc4, 0x2c, 0x12,
	0x04, 0x10, 0xb6, 0x94, 0x72, 0x82, 0x03, 0xaa, 0x52, 0x55, 0xb5, 0x44, 0x4b, 0x70, 0xa2, 0x72,
	0x01, 0xad, 0xc6, 0xbb, 0x2f, 0xce, 0x34, 0xbb, 0x3b, 0xcb, 0xcc, 0x6c, 0x20, 0xe1, 0xc6, 0x91,
	0x1b, 0x47, 0x8e, 0x70, 0xe3, 0x7b, 0xf1, 0x05, 0xb8, 0x72, 0x42, 0x33, 0xb3, 0xbb, 0x76, 0x9c,
	0xd8, 0x6e, 0x24, 0x9a, 0x8b, 0xb5, 0xef, 0xef, 0xbc, 0xf7, 0x7b, 0xf3, 0x7e, 0x1a, 0xc3, 0xb2,
	0x4c, 0x68, 0x2a, 0x4f, 0xb8, 0xea, 0xa6, 0x82, 0x2b, 0x4e, 0xea, 0x12, 0xc5, 0x19, 0x0b, 0xd0,
	0x6d, 0xff, 0xc0, 0x92, 0x51, 0x10, 0xf1, 0x2c, 0xf4, 0x65, 0x78, 0xea, 0x8b, 0x2c, 0xc2, 0x9e,
	0xfe, 0xb1, 0x7e, 0xee, 0xce, 0x88, 0xf3, 0x51, 0x84, 0x3d, 0x23, 0x0d, 0xb3, 0xe3, 0x9e, 0x62,
	0x31, 0x4a, 0x45, 0xe3, 0xd4, 0x3a, 0x78, 0x3d, 0xd8, 0x3c, 0xcc, 0x53, 0x1f, 0xda, 0x94, 0x07,
	0x82, 0xa7, 0x28, 0x14, 0x43, 0x49, 0x08, 0x2c, 0x5d, 0xf0, 0x04, 0x9d, 0x4a, 0xa7, 0xb2, 0xdb,
	0x1c, 0x98, 0x6f, 0xef, 0xb7, 0x3b, 0xb0, 0xf1, 0x0c, 0x65, 0x20, 0xd8, 0x10, 0x8b, 0x48, 0xd9,
	0x4f, 0xd2, 0x4c, 0x91, 0xc7, 0xd0, 0x2c, 0xca, 0x94, 0x4e, 0xa5, 0x73, 0x77, 0xb7, 0x39, 0x18,
	0x2b, 0xc8, 0x0e, 0xb4, 0x04, 0x4a, 0x9e, 0x89, 0x00, 0x7d, 0x16, 0x3a, 0x77, 0x4c, 0x4e, 0x28,
	0x54, 0xfd, 0x90, 0xbc, 0x0f, 0xf7, 0x0b, 0x6f, 0x5f, 0x9d, 0xa7, 0xe8, 0xdc, 0xed, 0x54, 0x76,
	0xab, 0x83, 0x7b, 0x85, 0xf2, 0xe8, 0x3c, 0x45, 0xb2, 0x01, 0x35, 0xa9, 0xa8, 0xca, 0xa4, 0xb3,
	0x64, 0x0e, 0xc8, 0x25, 0x9d, 0x5d, 0x22, 0x15, 0xc1, 0x89, 0xff, 0x23, 0x17, 0xa1, 0x53, 0xb5,
	0xd9, 0xad, 0xea, 0x5b, 0x2e, 0x42, 0xdd, 0x8b, 0xa2, 0x23, 0xe9, 0xd4, 0x4c, 0x98, 0xf9, 0x26,
	0x0e, 0xd4, 0xcf, 0x50, 0x0c, 0xb9, 0x44, 0xa7, 0x6e, 0xce, 0x2a, 0x44, 0x7d, 0x0c, 0x3f, 0x3e,
	0x96, 0xa8, 0x9c, 0x86, 0x31, 0xe4, 0x12, 0x59, 0x83, 0x6a, 0xc4, 0x62, 0xa6, 0x9c, 0xa6, 0x51,
	0x5b, 0xc1, 0xfb, 0xbb, 0x06, 0x8f, 0xae, 0x60, 0xf2, 0x75, 0xa6, 0x34, 0x28, 0x1b, 0x50, 0xa3,
	0x81, 0x62, 0x3c, 0xc9, 0x51, 0xcc, 0x25, 0xb2, 0x09, 0x0d, 0x81, 0xca, 0x0f, 0x78, 0x88, 0x06,
	0x8b, 0xea, 0xa0, 0x2e, 0x50, 0xed, 0xf3, 0x10, 0x75, 0x59, 0x31, 0x4a, 0x49, 0x47, 0x16, 0x82,
	0xe6, 0xa0, 0x10, 0xc9, 0x01, 0x94, 0x68, 0xf8, 0xba, 0x38, 0x8d, 0x41, 0x6b, 0xef, 0xd3, 0x6e,
	0x7e, 0x1b, 0xba, 0x33, 0x8a, 0xe8, 0x0e, 0x50, 0xa6, 0x3c, 0x91, 0xd8, 0x57, 0x18, 0x0f, 0x5a,
	0xb2, 0x1c, 0xb8, 0x72, 0xff, 0xaa, 0xc2, 0xbd, 0x49, 0xab, 0x01, 0xb2, 0x38, 0x82, 0x85, 0x79,
	0xd1, 0x50, 0xa8, 0xa6, 0xc6, 0x94, 0xd0, 0x18, 0xf3, 0x49, 0x96, 0x85, 0xbd, 0xa2, 0x31, 0x92,
	0x0e, 0xb4, 0x42, 0x53, 0x4b, 0x6a, 0x5a, 0xb7, 0x6d, 0x4c, 0xaa, 0xae, 0x4e, 0x7b, 0xe9, 0x72,
	0x9a, 0xa9, 0x69, 0xdb, 0x81, 0x16, 0xd3, 0xfe, 0x04, 0x56, 0x95, 0xa0, 0x89, 0x64, 0x3a, 0x95,
	0x9f, 0xbb, 0xd4, 0x8c, 0xcb, 0x83, 0xb1, 0xe1, 0xd0, 0x3a, 0x7f, 0x01, 0xad, 0x40, 0x20, 0x55,
	0xe8, 0xeb, 0xcb, 0x6f, 0x26, 0xdd, 0xda, 0x73, 0xbb, 0x76, 0x33, 0xba, 0xc5, 0x66, 0x74, 0x8f,
	0x8a, 0xcd, 0x18, 0x80, 0x75, 0xd7, 0x0a, 0x1d, 0x6c, 0xd3, 0xdb, 0xe0, 0xc6, 0xe2, 0x60, 0xeb,
	0x6e, 0x82, 0xbf, 0x9c, 0xec, 0x51, 0x87, 0x37, 0x17, 0x86, 0x8f, 0xfb, 0xd7, 0x09, 0x36, 0xa1,
	0xc1, 0xa4, 0xaf, 0xe8, 0x29, 0x26, 0x0e, 0xd8, 0x4b, 0xc2, 0xe4, 0x91, 0x16, 0xc9, 0x23, 0xa8,
	0x33, 0xe9, 0x9f, 0x20, 0x0d, 0x9d, 0x96, 0xbd, 0xa2, 0x4c, 0xbe, 0x40, 0x1a, 0x6a, 0x83, 0xe0,
	0x76, 0x78, 0xf7, 0x2c, 0x68, 0x5a, 0xec, 0x87, 0x64, 0x0b, 0x9a, 0x29, 0x15, 0x98, 0x18, 0xd3,
	0x7d, 0x63, 0x6a, 0x58, 0x45, 0xdf, 0xac, 0x87, 0x64, 0x17, 0xe8, 0x2c, 0x9b, 0x5c, 0xe6, 0x9b,
	0xb4, 0x01, 0x14, 0x57, 0x34, 0xf2, 0x8d, 0x65, 0xc5, 0x58, 0x9a, 0x46, 0x73, 0xa8, 0xcd, 0x3b,
	0xd0, 0xb2, 0xe6, 0x80, 0x67, 0x89, 0x72, 0x1e, 0x18, 0xbb, 0x8d, 0xd8, 0xd7, 0x1a, 0xf2, 0x0a,
	0xd6, 0x23, 0x2a, 0x15, 0x4a, 0xe5, 0x5f, 0x86, 0x61, 0x75, 0x21, 0x0c, 0x0f, 0xf3, 0xc0, 0xc3,
	0x09, 0x34, 0x3c, 0x01, 0x6b, 0xfb, 0x66, 0x32, 0x57, 0x79, 0xa7, 0xa0, 0x91, 0x92, 0x77, 0x4a,
	0xc5, 0xdb, 0xdd, 0x57, 0x8b, 0xe6, 0x71, 0x16, 0x45, 0x39, 0xeb, 0xd4, 0x98, 0x7c, 0x9e, 0x45,
	0x91, 0xf7, 0x7b, 0x05, 0xd6, 0xa7, 0x0e, 0x7d, 0x17, 0x8b, 0xbd, 0x0e, 0xb5, 0x37, 0x7c, 0xa8,
	0x07, 0x63, 0xd7, 0xa0, 0xfa, 0x86, 0x0f, 0xfb, 0xe1, 0x65, 0x46, 0xad, 0x4e, 0x31, 0xaa, 0xf7,
	0x19, 0xac, 0x3d, 0xc3, 0x08, 0xd5, 0x8d, 0x78, 0xd8, 0xfb, 0x19, 0xd6, 0xa7, 0xa2, 0x6e, 0xaf,
	0x21, 0xef, 0x09, 0x3c, 0x7c, 0x9a, 0xa6, 0xd1, 0xf9, 0x8d, 0x2a, 0xbe, 0x80, 0xb5, 0xcb, 0x41,
	0xb7, 0x58, 0xf0, 0x2f, 0x15, 0x68, 0xbf, 0xe4, 0x21, 0x3b, 0x2e, 0x4f, 0x7f, 0xaa, 0x94, 0x60,
	0xc3, 0x4c, 0x61, 0x5e, 0xbb, 0x0b, 0x8d, 0xa2, 0xd4, 0xbc, 0x8e, 0x52, 0xfe, 0x9f, 0xb8, 0xd2,
	0x8b, 0x61, 0x7b, 0x56, 0x0d, 0xef, 0x00, 0x0a, 0xef, 0x7b, 0xe8, 0xec, 0xd3, 0x54, 0x65, 0x02,
	0xfb, 0x89, 0x54, 0x34, 0x09, 0xf0, 0xb9, 0xe0, 0x71, 0x71, 0xf6, 0xe2, 0xae, 0xdb, 0x00, 0x2c,
	0xa6, 0x23, 0x9c, 0x6c, 0xb9, 0x69, 0x34, 0xba, 0x5f, 0xef, 0xcf, 0x0a, 0xbc, 0x37, 0x27, 0xff,
	0x2d, 0xae, 0x97, 0xa6, 0x57, 0x53, 0x28, 0x2b, 0x5e, 0x0c, 0x75, 0x23, 0xf7, 0x43, 0xef, 0x3b,
	0x68, 0xdb, 0xb5, 0x7f, 0xcd, 0xa3, 0x2c, 0xbe, 0x21, 0x00, 0x3b, 0xd0, 0x3a, 0x33, 0x61, 0x93,
	0x08, 0x80, 0x55, 0x19, 0x08, 0xfe, 0xa8, 0xc0, 0xf6, 0xac, 0xf4, 0xb7, 0xd8, 0xff, 0x16, 0x34,
	0xf3, 0x3a, 0x4b, 0x00, 0x1a, 0x56, 0xd1, 0x0f, 0xf7, 0xfe, 0xad, 0xc2, 0xca, 0xd4, 0xd3, 0x90,
	0xbc, 0x86, 0xd5, 0x2b, 0x4f, 0x0c, 0xb2, 0x33, 0xfb, 0xf9, 0x61, 0xa0, 0x72, 0x3b, 0x8b, 0xde,
	0x27, 0xe4, 0x00, 0x56, 0xa6, 0x48, 0x96, 0xb4, 0xcb, 0xa0, 0xeb, 0x38, 0xdf, 0xdd, 0x9e, 0x65,
	0x1e, 0x67, 0x9c, 0x62, 0xb9, 0x89, 0x8c, 0xd7, 0xb1, 0xe6, 0x44, 0xc6, 0xeb, 0xe9, 0xf1, 0x25,
	0x2c, 0x5f, 0x66, 0x21, 0xf2, 0xb8, 0x8c, 0xb8, 0x86, 0xd3, 0xdc, 0xf6, 0x0c, 0x6b, 0x9e, 0xee,
	0x14, 0x9c, 0x59, 0x3b, 0x4d, 0x3e, 0x28, 0x43, 0xe7, 0x52, 0x8f, 0xfb, 0xe1, 0x42, 0xbf, 0xfc,
	0x30, 0x05, 0x5b, 0x73, 0x36, 0x8e, 0x7c, 0x34, 0x06, 0x73, 0xc1, 0xde, 0xbb, 0x1f, 0xbf, 0x8d,
	0xeb, 0xb8, 0xc5, 0x59, 0x97, 0x7c, 0xa2, 0xc5, 0xb9, 0x6b, 0x36, 0xd1, 0xe2, 0xfc, 0x7d, 0x71,
	0xbf, 0xfa, 0xf5, 0x9f, 0xa5, 0x17, 0xb0, 0x77, 0xa2, 0x54, 0x2a, 0x3f, 0xef, 0xf5, 0x42, 0x1e,
	0xc8, 0x6e, 0xf9, 0xdf, 0xa8, 0x1b, 0xf0, 0xb8, 0x47, 0x53, 0xd6, 0x2b, 0x56, 0xb4, 0xc7, 0x92,
	0x10, 0x7f, 0xea, 0x9e, 0xa8, 0x38, 0x22, 0xe4, 0x1b, 0x96, 0x8c, 0xf6, 0x8d, 0x5b, 0x91, 0x73,
	0x58, 0x33, 0x6f, 0x92, 0x27, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x81, 0xce, 0x5d, 0xf3, 0x78,
	0x0d, 0x00, 0x00,
}
