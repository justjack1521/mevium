// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protoadmin/response.proto

package protoadmin

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

type GrantItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GrantItemResponse) Reset() {
	*x = GrantItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoadmin_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantItemResponse) ProtoMessage() {}

func (x *GrantItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoadmin_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantItemResponse.ProtoReflect.Descriptor instead.
func (*GrantItemResponse) Descriptor() ([]byte, []int) {
	return file_protoadmin_response_proto_rawDescGZIP(), []int{0}
}

type CreateSkillPanelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created bool `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *CreateSkillPanelResponse) Reset() {
	*x = CreateSkillPanelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoadmin_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSkillPanelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSkillPanelResponse) ProtoMessage() {}

func (x *CreateSkillPanelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoadmin_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSkillPanelResponse.ProtoReflect.Descriptor instead.
func (*CreateSkillPanelResponse) Descriptor() ([]byte, []int) {
	return file_protoadmin_response_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSkillPanelResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

type CreateBaseJobCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created bool `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *CreateBaseJobCardResponse) Reset() {
	*x = CreateBaseJobCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoadmin_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBaseJobCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBaseJobCardResponse) ProtoMessage() {}

func (x *CreateBaseJobCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protoadmin_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBaseJobCardResponse.ProtoReflect.Descriptor instead.
func (*CreateBaseJobCardResponse) Descriptor() ([]byte, []int) {
	return file_protoadmin_response_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBaseJobCardResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

var File_protoadmin_response_proto protoreflect.FileDescriptor

var file_protoadmin_response_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x35, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f,
	0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoadmin_response_proto_rawDescOnce sync.Once
	file_protoadmin_response_proto_rawDescData = file_protoadmin_response_proto_rawDesc
)

func file_protoadmin_response_proto_rawDescGZIP() []byte {
	file_protoadmin_response_proto_rawDescOnce.Do(func() {
		file_protoadmin_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoadmin_response_proto_rawDescData)
	})
	return file_protoadmin_response_proto_rawDescData
}

var file_protoadmin_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protoadmin_response_proto_goTypes = []interface{}{
	(*GrantItemResponse)(nil),         // 0: admin.GrantItemResponse
	(*CreateSkillPanelResponse)(nil),  // 1: admin.CreateSkillPanelResponse
	(*CreateBaseJobCardResponse)(nil), // 2: admin.CreateBaseJobCardResponse
}
var file_protoadmin_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoadmin_response_proto_init() }
func file_protoadmin_response_proto_init() {
	if File_protoadmin_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoadmin_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantItemResponse); i {
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
		file_protoadmin_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSkillPanelResponse); i {
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
		file_protoadmin_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBaseJobCardResponse); i {
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
			RawDescriptor: file_protoadmin_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoadmin_response_proto_goTypes,
		DependencyIndexes: file_protoadmin_response_proto_depIdxs,
		MessageInfos:      file_protoadmin_response_proto_msgTypes,
	}.Build()
	File_protoadmin_response_proto = out.File
	file_protoadmin_response_proto_rawDesc = nil
	file_protoadmin_response_proto_goTypes = nil
	file_protoadmin_response_proto_depIdxs = nil
}
