// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: logstream.proto

package logstream

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string   `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Deltas  []*Delta `protobuf:"bytes,2,rep,name=deltas,proto3" json:"deltas,omitempty"`
	Eof     bool     `protobuf:"varint,3,opt,name=eof,proto3" json:"eof,omitempty"`
}

func (x *StreamLogRequest) Reset() {
	*x = StreamLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamLogRequest) ProtoMessage() {}

func (x *StreamLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamLogRequest.ProtoReflect.Descriptor instead.
func (*StreamLogRequest) Descriptor() ([]byte, []int) {
	return file_logstream_proto_rawDescGZIP(), []int{0}
}

func (x *StreamLogRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *StreamLogRequest) GetDeltas() []*Delta {
	if x != nil {
		return x.Deltas
	}
	return nil
}

func (x *StreamLogRequest) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

type StreamLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EofAck bool `protobuf:"varint,1,opt,name=eof_ack,json=eofAck,proto3" json:"eof_ack,omitempty"`
}

func (x *StreamLogResponse) Reset() {
	*x = StreamLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamLogResponse) ProtoMessage() {}

func (x *StreamLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamLogResponse.ProtoReflect.Descriptor instead.
func (*StreamLogResponse) Descriptor() ([]byte, []int) {
	return file_logstream_proto_rawDescGZIP(), []int{1}
}

func (x *StreamLogResponse) GetEofAck() bool {
	if x != nil {
		return x.EofAck
	}
	return false
}

type GetFirebaseAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFirebaseAuthTokenRequest) Reset() {
	*x = GetFirebaseAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFirebaseAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFirebaseAuthTokenRequest) ProtoMessage() {}

func (x *GetFirebaseAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFirebaseAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*GetFirebaseAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_logstream_proto_rawDescGZIP(), []int{2}
}

type GetFirebaseAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetFirebaseAuthTokenResponse) Reset() {
	*x = GetFirebaseAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logstream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFirebaseAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFirebaseAuthTokenResponse) ProtoMessage() {}

func (x *GetFirebaseAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logstream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFirebaseAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*GetFirebaseAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_logstream_proto_rawDescGZIP(), []int{3}
}

func (x *GetFirebaseAuthTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_logstream_proto protoreflect.FileDescriptor

var file_logstream_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x06,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x66, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x22, 0x2c, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x6f, 0x66, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x65, 0x6f, 0x66, 0x41, 0x63, 0x6b, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x9f, 0x02, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x63, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0xac,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logstream_proto_rawDescOnce sync.Once
	file_logstream_proto_rawDescData = file_logstream_proto_rawDesc
)

func file_logstream_proto_rawDescGZIP() []byte {
	file_logstream_proto_rawDescOnce.Do(func() {
		file_logstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_logstream_proto_rawDescData)
	})
	return file_logstream_proto_rawDescData
}

var file_logstream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logstream_proto_goTypes = []interface{}{
	(*StreamLogRequest)(nil),             // 0: api.public.logstream.StreamLogRequest
	(*StreamLogResponse)(nil),            // 1: api.public.logstream.StreamLogResponse
	(*GetFirebaseAuthTokenRequest)(nil),  // 2: api.public.logstream.GetFirebaseAuthTokenRequest
	(*GetFirebaseAuthTokenResponse)(nil), // 3: api.public.logstream.GetFirebaseAuthTokenResponse
	(*Delta)(nil),                        // 4: api.public.logstream.Delta
}
var file_logstream_proto_depIdxs = []int32{
	4, // 0: api.public.logstream.StreamLogRequest.deltas:type_name -> api.public.logstream.Delta
	0, // 1: api.public.logstream.LogStream.StreamLogs:input_type -> api.public.logstream.StreamLogRequest
	2, // 2: api.public.logstream.LogStream.GetFirebaseAuthToken:input_type -> api.public.logstream.GetFirebaseAuthTokenRequest
	1, // 3: api.public.logstream.LogStream.StreamLogs:output_type -> api.public.logstream.StreamLogResponse
	3, // 4: api.public.logstream.LogStream.GetFirebaseAuthToken:output_type -> api.public.logstream.GetFirebaseAuthTokenResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_logstream_proto_init() }
func file_logstream_proto_init() {
	if File_logstream_proto != nil {
		return
	}
	file_manifest_proto_init()
	file_delta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_logstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamLogRequest); i {
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
		file_logstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamLogResponse); i {
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
		file_logstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFirebaseAuthTokenRequest); i {
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
		file_logstream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFirebaseAuthTokenResponse); i {
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
			RawDescriptor: file_logstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logstream_proto_goTypes,
		DependencyIndexes: file_logstream_proto_depIdxs,
		MessageInfos:      file_logstream_proto_msgTypes,
	}.Build()
	File_logstream_proto = out.File
	file_logstream_proto_rawDesc = nil
	file_logstream_proto_goTypes = nil
	file_logstream_proto_depIdxs = nil
}
