// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protoidentity/identity.request.proto

package protoidentity

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

type GetSinglePlayerIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *GetSinglePlayerIdentityRequest) Reset() {
	*x = GetSinglePlayerIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinglePlayerIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinglePlayerIdentityRequest) ProtoMessage() {}

func (x *GetSinglePlayerIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinglePlayerIdentityRequest.ProtoReflect.Descriptor instead.
func (*GetSinglePlayerIdentityRequest) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetSinglePlayerIdentityRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetSinglePlayerLoadoutIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *GetSinglePlayerLoadoutIdentityRequest) Reset() {
	*x = GetSinglePlayerLoadoutIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinglePlayerLoadoutIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinglePlayerLoadoutIdentityRequest) ProtoMessage() {}

func (x *GetSinglePlayerLoadoutIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinglePlayerLoadoutIdentityRequest.ProtoReflect.Descriptor instead.
func (*GetSinglePlayerLoadoutIdentityRequest) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_request_proto_rawDescGZIP(), []int{1}
}

func (x *GetSinglePlayerLoadoutIdentityRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetSinglePlayerLoadoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	DeckIndex int32  `protobuf:"varint,2,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
}

func (x *GetSinglePlayerLoadoutRequest) Reset() {
	*x = GetSinglePlayerLoadoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinglePlayerLoadoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinglePlayerLoadoutRequest) ProtoMessage() {}

func (x *GetSinglePlayerLoadoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinglePlayerLoadoutRequest.ProtoReflect.Descriptor instead.
func (*GetSinglePlayerLoadoutRequest) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_request_proto_rawDescGZIP(), []int{2}
}

func (x *GetSinglePlayerLoadoutRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *GetSinglePlayerLoadoutRequest) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

type GetMultiPlayerLoadoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	DeckIndex int32  `protobuf:"varint,2,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
}

func (x *GetMultiPlayerLoadoutRequest) Reset() {
	*x = GetMultiPlayerLoadoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultiPlayerLoadoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultiPlayerLoadoutRequest) ProtoMessage() {}

func (x *GetMultiPlayerLoadoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultiPlayerLoadoutRequest.ProtoReflect.Descriptor instead.
func (*GetMultiPlayerLoadoutRequest) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_request_proto_rawDescGZIP(), []int{3}
}

func (x *GetMultiPlayerLoadoutRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *GetMultiPlayerLoadoutRequest) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

var File_protoidentity_identity_request_proto protoreflect.FileDescriptor

var file_protoidentity_identity_request_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x3d, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x44, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0x5a, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x53,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73,
	0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0xaa, 0x02, 0x15, 0x4d, 0x6f,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoidentity_identity_request_proto_rawDescOnce sync.Once
	file_protoidentity_identity_request_proto_rawDescData = file_protoidentity_identity_request_proto_rawDesc
)

func file_protoidentity_identity_request_proto_rawDescGZIP() []byte {
	file_protoidentity_identity_request_proto_rawDescOnce.Do(func() {
		file_protoidentity_identity_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoidentity_identity_request_proto_rawDescData)
	})
	return file_protoidentity_identity_request_proto_rawDescData
}

var file_protoidentity_identity_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protoidentity_identity_request_proto_goTypes = []interface{}{
	(*GetSinglePlayerIdentityRequest)(nil),        // 0: identity.GetSinglePlayerIdentityRequest
	(*GetSinglePlayerLoadoutIdentityRequest)(nil), // 1: identity.GetSinglePlayerLoadoutIdentityRequest
	(*GetSinglePlayerLoadoutRequest)(nil),         // 2: identity.GetSinglePlayerLoadoutRequest
	(*GetMultiPlayerLoadoutRequest)(nil),          // 3: identity.GetMultiPlayerLoadoutRequest
}
var file_protoidentity_identity_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoidentity_identity_request_proto_init() }
func file_protoidentity_identity_request_proto_init() {
	if File_protoidentity_identity_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoidentity_identity_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinglePlayerIdentityRequest); i {
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
		file_protoidentity_identity_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinglePlayerLoadoutIdentityRequest); i {
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
		file_protoidentity_identity_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinglePlayerLoadoutRequest); i {
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
		file_protoidentity_identity_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultiPlayerLoadoutRequest); i {
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
			RawDescriptor: file_protoidentity_identity_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoidentity_identity_request_proto_goTypes,
		DependencyIndexes: file_protoidentity_identity_request_proto_depIdxs,
		MessageInfos:      file_protoidentity_identity_request_proto_msgTypes,
	}.Build()
	File_protoidentity_identity_request_proto = out.File
	file_protoidentity_identity_request_proto_rawDesc = nil
	file_protoidentity_identity_request_proto_goTypes = nil
	file_protoidentity_identity_request_proto_depIdxs = nil
}
