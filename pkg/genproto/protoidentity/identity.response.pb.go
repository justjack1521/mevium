// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protoidentity/identity.response.proto

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

type GetSinglePlayerIdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *ProtoPlayerIdentity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *GetSinglePlayerIdentityResponse) Reset() {
	*x = GetSinglePlayerIdentityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinglePlayerIdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinglePlayerIdentityResponse) ProtoMessage() {}

func (x *GetSinglePlayerIdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinglePlayerIdentityResponse.ProtoReflect.Descriptor instead.
func (*GetSinglePlayerIdentityResponse) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetSinglePlayerIdentityResponse) GetIdentity() *ProtoPlayerIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

type GetSinglePlayerLoadoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *ProtoPlayerIdentity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Loadout  *ProtoPlayerLoadout  `protobuf:"bytes,2,opt,name=loadout,proto3" json:"loadout,omitempty"`
}

func (x *GetSinglePlayerLoadoutResponse) Reset() {
	*x = GetSinglePlayerLoadoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSinglePlayerLoadoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSinglePlayerLoadoutResponse) ProtoMessage() {}

func (x *GetSinglePlayerLoadoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSinglePlayerLoadoutResponse.ProtoReflect.Descriptor instead.
func (*GetSinglePlayerLoadoutResponse) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_response_proto_rawDescGZIP(), []int{1}
}

func (x *GetSinglePlayerLoadoutResponse) GetIdentity() *ProtoPlayerIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *GetSinglePlayerLoadoutResponse) GetLoadout() *ProtoPlayerLoadout {
	if x != nil {
		return x.Loadout
	}
	return nil
}

type GetMultiPlayerLoadoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *ProtoPlayerIdentity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Loadout  *ProtoPlayerLoadout  `protobuf:"bytes,2,opt,name=loadout,proto3" json:"loadout,omitempty"`
}

func (x *GetMultiPlayerLoadoutResponse) Reset() {
	*x = GetMultiPlayerLoadoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoidentity_identity_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultiPlayerLoadoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultiPlayerLoadoutResponse) ProtoMessage() {}

func (x *GetMultiPlayerLoadoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoidentity_identity_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultiPlayerLoadoutResponse.ProtoReflect.Descriptor instead.
func (*GetMultiPlayerLoadoutResponse) Descriptor() ([]byte, []int) {
	return file_protoidentity_identity_response_proto_rawDescGZIP(), []int{2}
}

func (x *GetMultiPlayerLoadoutResponse) GetIdentity() *ProtoPlayerIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *GetMultiPlayerLoadoutResponse) GetLoadout() *ProtoPlayerLoadout {
	if x != nil {
		return x.Loadout
	}
	return nil
}

var File_protoidentity_identity_response_proto protoreflect.FileDescriptor

var file_protoidentity_identity_response_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x93, 0x01, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c,
	0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x6c, 0x6f, 0x61,
	0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75,
	0x74, 0x22, 0x92, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36,
	0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x6c,
	0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x42, 0x53, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32,
	0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0xaa, 0x02, 0x15, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protoidentity_identity_response_proto_rawDescOnce sync.Once
	file_protoidentity_identity_response_proto_rawDescData = file_protoidentity_identity_response_proto_rawDesc
)

func file_protoidentity_identity_response_proto_rawDescGZIP() []byte {
	file_protoidentity_identity_response_proto_rawDescOnce.Do(func() {
		file_protoidentity_identity_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoidentity_identity_response_proto_rawDescData)
	})
	return file_protoidentity_identity_response_proto_rawDescData
}

var file_protoidentity_identity_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protoidentity_identity_response_proto_goTypes = []interface{}{
	(*GetSinglePlayerIdentityResponse)(nil), // 0: identity.GetSinglePlayerIdentityResponse
	(*GetSinglePlayerLoadoutResponse)(nil),  // 1: identity.GetSinglePlayerLoadoutResponse
	(*GetMultiPlayerLoadoutResponse)(nil),   // 2: identity.GetMultiPlayerLoadoutResponse
	(*ProtoPlayerIdentity)(nil),             // 3: identity.ProtoPlayerIdentity
	(*ProtoPlayerLoadout)(nil),              // 4: identity.ProtoPlayerLoadout
}
var file_protoidentity_identity_response_proto_depIdxs = []int32{
	3, // 0: identity.GetSinglePlayerIdentityResponse.identity:type_name -> identity.ProtoPlayerIdentity
	3, // 1: identity.GetSinglePlayerLoadoutResponse.identity:type_name -> identity.ProtoPlayerIdentity
	4, // 2: identity.GetSinglePlayerLoadoutResponse.loadout:type_name -> identity.ProtoPlayerLoadout
	3, // 3: identity.GetMultiPlayerLoadoutResponse.identity:type_name -> identity.ProtoPlayerIdentity
	4, // 4: identity.GetMultiPlayerLoadoutResponse.loadout:type_name -> identity.ProtoPlayerLoadout
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protoidentity_identity_response_proto_init() }
func file_protoidentity_identity_response_proto_init() {
	if File_protoidentity_identity_response_proto != nil {
		return
	}
	file_protoidentity_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protoidentity_identity_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinglePlayerIdentityResponse); i {
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
		file_protoidentity_identity_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSinglePlayerLoadoutResponse); i {
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
		file_protoidentity_identity_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultiPlayerLoadoutResponse); i {
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
			RawDescriptor: file_protoidentity_identity_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoidentity_identity_response_proto_goTypes,
		DependencyIndexes: file_protoidentity_identity_response_proto_depIdxs,
		MessageInfos:      file_protoidentity_identity_response_proto_msgTypes,
	}.Build()
	File_protoidentity_identity_response_proto = out.File
	file_protoidentity_identity_response_proto_rawDesc = nil
	file_protoidentity_identity_response_proto_goTypes = nil
	file_protoidentity_identity_response_proto_depIdxs = nil
}
