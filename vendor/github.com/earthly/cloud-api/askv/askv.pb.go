// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: askv.proto

package askv

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName      string `protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	ProjectName  string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	PipelineName string `protobuf:"bytes,3,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	Hash         []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ExistsRequest) Reset() {
	*x = ExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsRequest) ProtoMessage() {}

func (x *ExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsRequest.ProtoReflect.Descriptor instead.
func (*ExistsRequest) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{0}
}

func (x *ExistsRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *ExistsRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ExistsRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *ExistsRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type ExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *ExistsResponse) Reset() {
	*x = ExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsResponse) ProtoMessage() {}

func (x *ExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsResponse.ProtoReflect.Descriptor instead.
func (*ExistsResponse) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{1}
}

func (x *ExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName      string `protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	ProjectName  string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	PipelineName string `protobuf:"bytes,3,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	Hash         []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *AddRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *AddRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *AddRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{3}
}

type PrunePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName      string `protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	ProjectName  string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	PipelineName string `protobuf:"bytes,3,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
}

func (x *PrunePipelineRequest) Reset() {
	*x = PrunePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrunePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrunePipelineRequest) ProtoMessage() {}

func (x *PrunePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrunePipelineRequest.ProtoReflect.Descriptor instead.
func (*PrunePipelineRequest) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{4}
}

func (x *PrunePipelineRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *PrunePipelineRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *PrunePipelineRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

type PrunePipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrunePipelineResponse) Reset() {
	*x = PrunePipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrunePipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrunePipelineResponse) ProtoMessage() {}

func (x *PrunePipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrunePipelineResponse.ProtoReflect.Descriptor instead.
func (*PrunePipelineResponse) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{5}
}

type PruneProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName     string `protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	ProjectName string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *PruneProjectRequest) Reset() {
	*x = PruneProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneProjectRequest) ProtoMessage() {}

func (x *PruneProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneProjectRequest.ProtoReflect.Descriptor instead.
func (*PruneProjectRequest) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{6}
}

func (x *PruneProjectRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *PruneProjectRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type PruneProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PruneProjectResponse) Reset() {
	*x = PruneProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneProjectResponse) ProtoMessage() {}

func (x *PruneProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneProjectResponse.ProtoReflect.Descriptor instead.
func (*PruneProjectResponse) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{7}
}

type PruneOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName string `protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
}

func (x *PruneOrgRequest) Reset() {
	*x = PruneOrgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneOrgRequest) ProtoMessage() {}

func (x *PruneOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneOrgRequest.ProtoReflect.Descriptor instead.
func (*PruneOrgRequest) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{8}
}

func (x *PruneOrgRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

type PruneOrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PruneOrgResponse) Reset() {
	*x = PruneOrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_askv_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneOrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneOrgResponse) ProtoMessage() {}

func (x *PruneOrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_askv_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneOrgResponse.ProtoReflect.Descriptor instead.
func (*PruneOrgResponse) Descriptor() ([]byte, []int) {
	return file_askv_proto_rawDescGZIP(), []int{9}
}

var File_askv_proto protoreflect.FileDescriptor

var file_askv_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a,
	0x0d, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x28, 0x0a, 0x0e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22,
	0x83, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x0a, 0x14, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x17, 0x0a, 0x15, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x50, 0x72, 0x75, 0x6e,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a,
	0x14, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x4f, 0x72,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x4f, 0x72, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xab, 0x03, 0x0a, 0x04, 0x41, 0x73, 0x6b, 0x76,
	0x12, 0x4b, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x03, 0x41, 0x64, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61,
	0x73, 0x6b, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x60, 0x0a, 0x0d, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x61, 0x73, 0x6b, 0x76, 0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x50, 0x72, 0x75, 0x6e,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76, 0x2e, 0x50, 0x72, 0x75, 0x6e,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73, 0x6b, 0x76,
	0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x61, 0x73,
	0x6b, 0x76, 0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x61, 0x73, 0x6b, 0x76, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_askv_proto_rawDescOnce sync.Once
	file_askv_proto_rawDescData = file_askv_proto_rawDesc
)

func file_askv_proto_rawDescGZIP() []byte {
	file_askv_proto_rawDescOnce.Do(func() {
		file_askv_proto_rawDescData = protoimpl.X.CompressGZIP(file_askv_proto_rawDescData)
	})
	return file_askv_proto_rawDescData
}

var file_askv_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_askv_proto_goTypes = []interface{}{
	(*ExistsRequest)(nil),         // 0: api.public.askv.ExistsRequest
	(*ExistsResponse)(nil),        // 1: api.public.askv.ExistsResponse
	(*AddRequest)(nil),            // 2: api.public.askv.AddRequest
	(*AddResponse)(nil),           // 3: api.public.askv.AddResponse
	(*PrunePipelineRequest)(nil),  // 4: api.public.askv.PrunePipelineRequest
	(*PrunePipelineResponse)(nil), // 5: api.public.askv.PrunePipelineResponse
	(*PruneProjectRequest)(nil),   // 6: api.public.askv.PruneProjectRequest
	(*PruneProjectResponse)(nil),  // 7: api.public.askv.PruneProjectResponse
	(*PruneOrgRequest)(nil),       // 8: api.public.askv.PruneOrgRequest
	(*PruneOrgResponse)(nil),      // 9: api.public.askv.PruneOrgResponse
}
var file_askv_proto_depIdxs = []int32{
	0, // 0: api.public.askv.Askv.Exists:input_type -> api.public.askv.ExistsRequest
	2, // 1: api.public.askv.Askv.Add:input_type -> api.public.askv.AddRequest
	4, // 2: api.public.askv.Askv.PrunePipeline:input_type -> api.public.askv.PrunePipelineRequest
	6, // 3: api.public.askv.Askv.PruneProject:input_type -> api.public.askv.PruneProjectRequest
	8, // 4: api.public.askv.Askv.PruneOrg:input_type -> api.public.askv.PruneOrgRequest
	1, // 5: api.public.askv.Askv.Exists:output_type -> api.public.askv.ExistsResponse
	3, // 6: api.public.askv.Askv.Add:output_type -> api.public.askv.AddResponse
	5, // 7: api.public.askv.Askv.PrunePipeline:output_type -> api.public.askv.PrunePipelineResponse
	7, // 8: api.public.askv.Askv.PruneProject:output_type -> api.public.askv.PruneProjectResponse
	9, // 9: api.public.askv.Askv.PruneOrg:output_type -> api.public.askv.PruneOrgResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_askv_proto_init() }
func file_askv_proto_init() {
	if File_askv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_askv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrunePipelineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrunePipelineResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneOrgRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_askv_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneOrgResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_askv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_askv_proto_goTypes,
		DependencyIndexes: file_askv_proto_depIdxs,
		MessageInfos:      file_askv_proto_msgTypes,
	}.Build()
	File_askv_proto = out.File
	file_askv_proto_rawDesc = nil
	file_askv_proto_goTypes = nil
	file_askv_proto_depIdxs = nil
}
