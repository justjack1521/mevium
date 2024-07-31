// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protocommon/common.message.proto

package protocommon

import (
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

type ClientConnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	RemoteAddress string `protobuf:"bytes,2,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *ClientConnected) Reset() {
	*x = ClientConnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocommon_common_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnected) ProtoMessage() {}

func (x *ClientConnected) ProtoReflect() protoreflect.Message {
	mi := &file_protocommon_common_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnected.ProtoReflect.Descriptor instead.
func (*ClientConnected) Descriptor() ([]byte, []int) {
	return file_protocommon_common_message_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConnected) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ClientConnected) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

type ClientHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	RemoteAddress string `protobuf:"bytes,2,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *ClientHeartbeat) Reset() {
	*x = ClientHeartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocommon_common_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientHeartbeat) ProtoMessage() {}

func (x *ClientHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_protocommon_common_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientHeartbeat.ProtoReflect.Descriptor instead.
func (*ClientHeartbeat) Descriptor() ([]byte, []int) {
	return file_protocommon_common_message_proto_rawDescGZIP(), []int{1}
}

func (x *ClientHeartbeat) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ClientHeartbeat) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

type ClientDisconnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	RemoteAddress string `protobuf:"bytes,2,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *ClientDisconnected) Reset() {
	*x = ClientDisconnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocommon_common_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDisconnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDisconnected) ProtoMessage() {}

func (x *ClientDisconnected) ProtoReflect() protoreflect.Message {
	mi := &file_protocommon_common_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDisconnected.ProtoReflect.Descriptor instead.
func (*ClientDisconnected) Descriptor() ([]byte, []int) {
	return file_protocommon_common_message_proto_rawDescGZIP(), []int{2}
}

func (x *ClientDisconnected) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ClientDisconnected) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

var File_protocommon_common_message_proto protoreflect.FileDescriptor

var file_protocommon_common_message_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5a, 0x0a, 0x12,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31,
	0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocommon_common_message_proto_rawDescOnce sync.Once
	file_protocommon_common_message_proto_rawDescData = file_protocommon_common_message_proto_rawDesc
)

func file_protocommon_common_message_proto_rawDescGZIP() []byte {
	file_protocommon_common_message_proto_rawDescOnce.Do(func() {
		file_protocommon_common_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocommon_common_message_proto_rawDescData)
	})
	return file_protocommon_common_message_proto_rawDescData
}

var file_protocommon_common_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocommon_common_message_proto_goTypes = []interface{}{
	(*ClientConnected)(nil),    // 0: common.ClientConnected
	(*ClientHeartbeat)(nil),    // 1: common.ClientHeartbeat
	(*ClientDisconnected)(nil), // 2: common.ClientDisconnected
}
var file_protocommon_common_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocommon_common_message_proto_init() }
func file_protocommon_common_message_proto_init() {
	if File_protocommon_common_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocommon_common_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConnected); i {
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
		file_protocommon_common_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientHeartbeat); i {
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
		file_protocommon_common_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDisconnected); i {
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
			RawDescriptor: file_protocommon_common_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocommon_common_message_proto_goTypes,
		DependencyIndexes: file_protocommon_common_message_proto_depIdxs,
		MessageInfos:      file_protocommon_common_message_proto_msgTypes,
	}.Build()
	File_protocommon_common_message_proto = out.File
	file_protocommon_common_message_proto_rawDesc = nil
	file_protocommon_common_message_proto_goTypes = nil
	file_protocommon_common_message_proto_depIdxs = nil
}
