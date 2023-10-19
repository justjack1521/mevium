// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protomulti/multi.response.proto

package protomulti

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

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type EndSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndSessionResponse) Reset() {
	*x = EndSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionResponse) ProtoMessage() {}

func (x *EndSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionResponse.ProtoReflect.Descriptor instead.
func (*EndSessionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{1}
}

type SearchLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lobbies []*ProtoLobbySummary `protobuf:"bytes,1,rep,name=lobbies,proto3" json:"lobbies,omitempty"`
}

func (x *SearchLobbyResponse) Reset() {
	*x = SearchLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLobbyResponse) ProtoMessage() {}

func (x *SearchLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLobbyResponse.ProtoReflect.Descriptor instead.
func (*SearchLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{2}
}

func (x *SearchLobbyResponse) GetLobbies() []*ProtoLobbySummary {
	if x != nil {
		return x.Lobbies
	}
	return nil
}

type WatchLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchLobbyResponse) Reset() {
	*x = WatchLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchLobbyResponse) ProtoMessage() {}

func (x *WatchLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchLobbyResponse.ProtoReflect.Descriptor instead.
func (*WatchLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{3}
}

type DiscardLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscardLobbyResponse) Reset() {
	*x = DiscardLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscardLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscardLobbyResponse) ProtoMessage() {}

func (x *DiscardLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscardLobbyResponse.ProtoReflect.Descriptor instead.
func (*DiscardLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{4}
}

type JoinLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinLobbyResponse) Reset() {
	*x = JoinLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLobbyResponse) ProtoMessage() {}

func (x *JoinLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLobbyResponse.ProtoReflect.Descriptor instead.
func (*JoinLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{5}
}

type ReadyLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadyLobbyResponse) Reset() {
	*x = ReadyLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyLobbyResponse) ProtoMessage() {}

func (x *ReadyLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyLobbyResponse.ProtoReflect.Descriptor instead.
func (*ReadyLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{6}
}

var File_protomulti_multi_response_proto protoreflect.FileDescriptor

var file_protomulti_multi_response_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49,
	0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x6c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x07, 0x6c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4a, 0x6f, 0x69, 0x6e, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x52, 0x65, 0x61, 0x64, 0x79, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x4d, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65,
	0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0xaa, 0x02, 0x12, 0x4d,
	0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protomulti_multi_response_proto_rawDescOnce sync.Once
	file_protomulti_multi_response_proto_rawDescData = file_protomulti_multi_response_proto_rawDesc
)

func file_protomulti_multi_response_proto_rawDescGZIP() []byte {
	file_protomulti_multi_response_proto_rawDescOnce.Do(func() {
		file_protomulti_multi_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_protomulti_multi_response_proto_rawDescData)
	})
	return file_protomulti_multi_response_proto_rawDescData
}

var file_protomulti_multi_response_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protomulti_multi_response_proto_goTypes = []interface{}{
	(*CreateSessionResponse)(nil), // 0: multi.CreateSessionResponse
	(*EndSessionResponse)(nil),    // 1: multi.EndSessionResponse
	(*SearchLobbyResponse)(nil),   // 2: multi.SearchLobbyResponse
	(*WatchLobbyResponse)(nil),    // 3: multi.WatchLobbyResponse
	(*DiscardLobbyResponse)(nil),  // 4: multi.DiscardLobbyResponse
	(*JoinLobbyResponse)(nil),     // 5: multi.JoinLobbyResponse
	(*ReadyLobbyResponse)(nil),    // 6: multi.ReadyLobbyResponse
	(*ProtoLobbySummary)(nil),     // 7: multi.ProtoLobbySummary
}
var file_protomulti_multi_response_proto_depIdxs = []int32{
	7, // 0: multi.SearchLobbyResponse.lobbies:type_name -> multi.ProtoLobbySummary
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protomulti_multi_response_proto_init() }
func file_protomulti_multi_response_proto_init() {
	if File_protomulti_multi_response_proto != nil {
		return
	}
	file_protomulti_multi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protomulti_multi_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscardLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyLobbyResponse); i {
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
			RawDescriptor: file_protomulti_multi_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protomulti_multi_response_proto_goTypes,
		DependencyIndexes: file_protomulti_multi_response_proto_depIdxs,
		MessageInfos:      file_protomulti_multi_response_proto_msgTypes,
	}.Build()
	File_protomulti_multi_response_proto = out.File
	file_protomulti_multi_response_proto_rawDesc = nil
	file_protomulti_multi_response_proto_goTypes = nil
	file_protomulti_multi_response_proto_depIdxs = nil
}
