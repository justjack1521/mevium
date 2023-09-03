// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: challenge.response.proto

package protosc

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

type GetActivePlayerChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysId         string `protobuf:"bytes,1,opt,name=sys_id,json=sysId,proto3" json:"sys_id,omitempty"`
	ChallengeId   string `protobuf:"bytes,2,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ChallengeFlag int32  `protobuf:"varint,3,opt,name=challenge_flag,json=challengeFlag,proto3" json:"challenge_flag,omitempty"`
}

func (x *GetActivePlayerChallengeResponse) Reset() {
	*x = GetActivePlayerChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_challenge_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActivePlayerChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActivePlayerChallengeResponse) ProtoMessage() {}

func (x *GetActivePlayerChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_challenge_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActivePlayerChallengeResponse.ProtoReflect.Descriptor instead.
func (*GetActivePlayerChallengeResponse) Descriptor() ([]byte, []int) {
	return file_challenge_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetActivePlayerChallengeResponse) GetSysId() string {
	if x != nil {
		return x.SysId
	}
	return ""
}

func (x *GetActivePlayerChallengeResponse) GetChallengeId() string {
	if x != nil {
		return x.ChallengeId
	}
	return ""
}

func (x *GetActivePlayerChallengeResponse) GetChallengeFlag() int32 {
	if x != nil {
		return x.ChallengeFlag
	}
	return 0
}

type JoinSocialChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance *ProtoSocialChallengeInstance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *JoinSocialChallengeResponse) Reset() {
	*x = JoinSocialChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_challenge_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSocialChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSocialChallengeResponse) ProtoMessage() {}

func (x *JoinSocialChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_challenge_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSocialChallengeResponse.ProtoReflect.Descriptor instead.
func (*JoinSocialChallengeResponse) Descriptor() ([]byte, []int) {
	return file_challenge_response_proto_rawDescGZIP(), []int{1}
}

func (x *JoinSocialChallengeResponse) GetInstance() *ProtoSocialChallengeInstance {
	if x != nil {
		return x.Instance
	}
	return nil
}

var File_challenge_response_proto protoreflect.FileDescriptor

var file_challenge_response_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x1a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x79, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x79, 0x73,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x62, 0x0a, 0x1b,
	0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x44, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69,
	0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x63, 0xaa, 0x02, 0x0c, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_challenge_response_proto_rawDescOnce sync.Once
	file_challenge_response_proto_rawDescData = file_challenge_response_proto_rawDesc
)

func file_challenge_response_proto_rawDescGZIP() []byte {
	file_challenge_response_proto_rawDescOnce.Do(func() {
		file_challenge_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_challenge_response_proto_rawDescData)
	})
	return file_challenge_response_proto_rawDescData
}

var file_challenge_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_challenge_response_proto_goTypes = []interface{}{
	(*GetActivePlayerChallengeResponse)(nil), // 0: challenge.GetActivePlayerChallengeResponse
	(*JoinSocialChallengeResponse)(nil),      // 1: challenge.JoinSocialChallengeResponse
	(*ProtoSocialChallengeInstance)(nil),     // 2: challenge.ProtoSocialChallengeInstance
}
var file_challenge_response_proto_depIdxs = []int32{
	2, // 0: challenge.JoinSocialChallengeResponse.instance:type_name -> challenge.ProtoSocialChallengeInstance
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_challenge_response_proto_init() }
func file_challenge_response_proto_init() {
	if File_challenge_response_proto != nil {
		return
	}
	file_challenge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_challenge_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActivePlayerChallengeResponse); i {
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
		file_challenge_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSocialChallengeResponse); i {
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
			RawDescriptor: file_challenge_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_challenge_response_proto_goTypes,
		DependencyIndexes: file_challenge_response_proto_depIdxs,
		MessageInfos:      file_challenge_response_proto_msgTypes,
	}.Build()
	File_challenge_response_proto = out.File
	file_challenge_response_proto_rawDesc = nil
	file_challenge_response_proto_goTypes = nil
	file_challenge_response_proto_depIdxs = nil
}
