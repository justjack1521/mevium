// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protochallenge/challenge.request.proto

package protochallenge

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

type ChallengeRequestType int32

const (
	ChallengeRequestType_NONE                  ChallengeRequestType = 0
	ChallengeRequestType_JOIN_SOCIAL_CHALLENGE ChallengeRequestType = 3500
	ChallengeRequestType_GET_SOCIAL_CHALLENGE  ChallengeRequestType = 3600
)

// Enum value maps for ChallengeRequestType.
var (
	ChallengeRequestType_name = map[int32]string{
		0:    "NONE",
		3500: "JOIN_SOCIAL_CHALLENGE",
		3600: "GET_SOCIAL_CHALLENGE",
	}
	ChallengeRequestType_value = map[string]int32{
		"NONE":                  0,
		"JOIN_SOCIAL_CHALLENGE": 3500,
		"GET_SOCIAL_CHALLENGE":  3600,
	}
)

func (x ChallengeRequestType) Enum() *ChallengeRequestType {
	p := new(ChallengeRequestType)
	*p = x
	return p
}

func (x ChallengeRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChallengeRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_protochallenge_challenge_request_proto_enumTypes[0].Descriptor()
}

func (ChallengeRequestType) Type() protoreflect.EnumType {
	return &file_protochallenge_challenge_request_proto_enumTypes[0]
}

func (x ChallengeRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChallengeRequestType.Descriptor instead.
func (ChallengeRequestType) EnumDescriptor() ([]byte, []int) {
	return file_protochallenge_challenge_request_proto_rawDescGZIP(), []int{0}
}

type GetActivePlayerChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetActivePlayerChallengeRequest) Reset() {
	*x = GetActivePlayerChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protochallenge_challenge_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActivePlayerChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActivePlayerChallengeRequest) ProtoMessage() {}

func (x *GetActivePlayerChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protochallenge_challenge_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActivePlayerChallengeRequest.ProtoReflect.Descriptor instead.
func (*GetActivePlayerChallengeRequest) Descriptor() ([]byte, []int) {
	return file_protochallenge_challenge_request_proto_rawDescGZIP(), []int{0}
}

type JoinSocialChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinSocialChallengeRequest) Reset() {
	*x = JoinSocialChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protochallenge_challenge_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSocialChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSocialChallengeRequest) ProtoMessage() {}

func (x *JoinSocialChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protochallenge_challenge_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSocialChallengeRequest.ProtoReflect.Descriptor instead.
func (*JoinSocialChallengeRequest) Descriptor() ([]byte, []int) {
	return file_protochallenge_challenge_request_proto_rawDescGZIP(), []int{1}
}

type GetPlayerChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlayerChallengeRequest) Reset() {
	*x = GetPlayerChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protochallenge_challenge_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerChallengeRequest) ProtoMessage() {}

func (x *GetPlayerChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protochallenge_challenge_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerChallengeRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerChallengeRequest) Descriptor() ([]byte, []int) {
	return file_protochallenge_challenge_request_proto_rawDescGZIP(), []int{2}
}

var File_protochallenge_challenge_request_proto protoreflect.FileDescriptor

var file_protochallenge_challenge_request_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2a, 0x57, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x15, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x53, 0x4f, 0x43, 0x49,
	0x41, 0x4c, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10, 0xac, 0x1b, 0x12,
	0x19, 0x0a, 0x14, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x48,
	0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10, 0x90, 0x1c, 0x42, 0x55, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63,
	0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0xaa, 0x02, 0x16, 0x4d, 0x6f, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protochallenge_challenge_request_proto_rawDescOnce sync.Once
	file_protochallenge_challenge_request_proto_rawDescData = file_protochallenge_challenge_request_proto_rawDesc
)

func file_protochallenge_challenge_request_proto_rawDescGZIP() []byte {
	file_protochallenge_challenge_request_proto_rawDescOnce.Do(func() {
		file_protochallenge_challenge_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_protochallenge_challenge_request_proto_rawDescData)
	})
	return file_protochallenge_challenge_request_proto_rawDescData
}

var file_protochallenge_challenge_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protochallenge_challenge_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protochallenge_challenge_request_proto_goTypes = []interface{}{
	(ChallengeRequestType)(0),               // 0: challenge.ChallengeRequestType
	(*GetActivePlayerChallengeRequest)(nil), // 1: challenge.GetActivePlayerChallengeRequest
	(*JoinSocialChallengeRequest)(nil),      // 2: challenge.JoinSocialChallengeRequest
	(*GetPlayerChallengeRequest)(nil),       // 3: challenge.GetPlayerChallengeRequest
}
var file_protochallenge_challenge_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protochallenge_challenge_request_proto_init() }
func file_protochallenge_challenge_request_proto_init() {
	if File_protochallenge_challenge_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protochallenge_challenge_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActivePlayerChallengeRequest); i {
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
		file_protochallenge_challenge_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSocialChallengeRequest); i {
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
		file_protochallenge_challenge_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerChallengeRequest); i {
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
			RawDescriptor: file_protochallenge_challenge_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protochallenge_challenge_request_proto_goTypes,
		DependencyIndexes: file_protochallenge_challenge_request_proto_depIdxs,
		EnumInfos:         file_protochallenge_challenge_request_proto_enumTypes,
		MessageInfos:      file_protochallenge_challenge_request_proto_msgTypes,
	}.Build()
	File_protochallenge_challenge_request_proto = out.File
	file_protochallenge_challenge_request_proto_rawDesc = nil
	file_protochallenge_challenge_request_proto_goTypes = nil
	file_protochallenge_challenge_request_proto_depIdxs = nil
}
