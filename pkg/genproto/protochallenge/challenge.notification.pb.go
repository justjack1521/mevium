// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protochallenge/challenge.notification.proto

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

type NotificationType int32

const (
	NotificationType_BASE                           NotificationType = 0
	NotificationType_SOCIAL_CHALLENGE_PLAYER_JOINED NotificationType = 100
	NotificationType_SOCIAL_CHALLENGE_STARTED       NotificationType = 200
	NotificationType_SOCIAL_CHALLENGE_ENDED         NotificationType = 300
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0:   "BASE",
		100: "SOCIAL_CHALLENGE_PLAYER_JOINED",
		200: "SOCIAL_CHALLENGE_STARTED",
		300: "SOCIAL_CHALLENGE_ENDED",
	}
	NotificationType_value = map[string]int32{
		"BASE":                           0,
		"SOCIAL_CHALLENGE_PLAYER_JOINED": 100,
		"SOCIAL_CHALLENGE_STARTED":       200,
		"SOCIAL_CHALLENGE_ENDED":         300,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_protochallenge_challenge_notification_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_protochallenge_challenge_notification_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_protochallenge_challenge_notification_proto_rawDescGZIP(), []int{0}
}

type SocialChallengePlayerJoinedNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId  string                           `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	PlayerEntry *ProtoPlayerSocialChallengeEntry `protobuf:"bytes,2,opt,name=player_entry,json=playerEntry,proto3" json:"player_entry,omitempty"`
}

func (x *SocialChallengePlayerJoinedNotification) Reset() {
	*x = SocialChallengePlayerJoinedNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protochallenge_challenge_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialChallengePlayerJoinedNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialChallengePlayerJoinedNotification) ProtoMessage() {}

func (x *SocialChallengePlayerJoinedNotification) ProtoReflect() protoreflect.Message {
	mi := &file_protochallenge_challenge_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialChallengePlayerJoinedNotification.ProtoReflect.Descriptor instead.
func (*SocialChallengePlayerJoinedNotification) Descriptor() ([]byte, []int) {
	return file_protochallenge_challenge_notification_proto_rawDescGZIP(), []int{0}
}

func (x *SocialChallengePlayerJoinedNotification) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *SocialChallengePlayerJoinedNotification) GetPlayerEntry() *ProtoPlayerSocialChallengeEntry {
	if x != nil {
		return x.PlayerEntry
	}
	return nil
}

type SocialChallengeStartedNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *SocialChallengeStartedNotification) Reset() {
	*x = SocialChallengeStartedNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protochallenge_challenge_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialChallengeStartedNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialChallengeStartedNotification) ProtoMessage() {}

func (x *SocialChallengeStartedNotification) ProtoReflect() protoreflect.Message {
	mi := &file_protochallenge_challenge_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialChallengeStartedNotification.ProtoReflect.Descriptor instead.
func (*SocialChallengeStartedNotification) Descriptor() ([]byte, []int) {
	return file_protochallenge_challenge_notification_proto_rawDescGZIP(), []int{1}
}

func (x *SocialChallengeStartedNotification) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type SocialChallengeEndedNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *SocialChallengeEndedNotification) Reset() {
	*x = SocialChallengeEndedNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protochallenge_challenge_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialChallengeEndedNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialChallengeEndedNotification) ProtoMessage() {}

func (x *SocialChallengeEndedNotification) ProtoReflect() protoreflect.Message {
	mi := &file_protochallenge_challenge_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialChallengeEndedNotification.ProtoReflect.Descriptor instead.
func (*SocialChallengeEndedNotification) Descriptor() ([]byte, []int) {
	return file_protochallenge_challenge_notification_proto_rawDescGZIP(), []int{2}
}

func (x *SocialChallengeEndedNotification) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

var File_protochallenge_challenge_notification_proto protoreflect.FileDescriptor

var file_protochallenge_challenge_notification_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x27, 0x53, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0x45, 0x0a, 0x22, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x20, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x45, 0x6e,
	0x64, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x2a, 0x7c, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12, 0x22,
	0x0a, 0x1e, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e,
	0x47, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x64, 0x12, 0x1d, 0x0a, 0x18, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x48, 0x41,
	0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0xc8,
	0x01, 0x12, 0x1b, 0x0a, 0x16, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x48, 0x41, 0x4c,
	0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0xac, 0x02, 0x42, 0x55,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73,
	0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0xaa, 0x02, 0x16, 0x4d,
	0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protochallenge_challenge_notification_proto_rawDescOnce sync.Once
	file_protochallenge_challenge_notification_proto_rawDescData = file_protochallenge_challenge_notification_proto_rawDesc
)

func file_protochallenge_challenge_notification_proto_rawDescGZIP() []byte {
	file_protochallenge_challenge_notification_proto_rawDescOnce.Do(func() {
		file_protochallenge_challenge_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_protochallenge_challenge_notification_proto_rawDescData)
	})
	return file_protochallenge_challenge_notification_proto_rawDescData
}

var file_protochallenge_challenge_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protochallenge_challenge_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protochallenge_challenge_notification_proto_goTypes = []interface{}{
	(NotificationType)(0),                           // 0: challenge.NotificationType
	(*SocialChallengePlayerJoinedNotification)(nil), // 1: challenge.SocialChallengePlayerJoinedNotification
	(*SocialChallengeStartedNotification)(nil),      // 2: challenge.SocialChallengeStartedNotification
	(*SocialChallengeEndedNotification)(nil),        // 3: challenge.SocialChallengeEndedNotification
	(*ProtoPlayerSocialChallengeEntry)(nil),         // 4: challenge.ProtoPlayerSocialChallengeEntry
}
var file_protochallenge_challenge_notification_proto_depIdxs = []int32{
	4, // 0: challenge.SocialChallengePlayerJoinedNotification.player_entry:type_name -> challenge.ProtoPlayerSocialChallengeEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protochallenge_challenge_notification_proto_init() }
func file_protochallenge_challenge_notification_proto_init() {
	if File_protochallenge_challenge_notification_proto != nil {
		return
	}
	file_protochallenge_challenge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protochallenge_challenge_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialChallengePlayerJoinedNotification); i {
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
		file_protochallenge_challenge_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialChallengeStartedNotification); i {
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
		file_protochallenge_challenge_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialChallengeEndedNotification); i {
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
			RawDescriptor: file_protochallenge_challenge_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protochallenge_challenge_notification_proto_goTypes,
		DependencyIndexes: file_protochallenge_challenge_notification_proto_depIdxs,
		EnumInfos:         file_protochallenge_challenge_notification_proto_enumTypes,
		MessageInfos:      file_protochallenge_challenge_notification_proto_msgTypes,
	}.Build()
	File_protochallenge_challenge_notification_proto = out.File
	file_protochallenge_challenge_notification_proto_rawDesc = nil
	file_protochallenge_challenge_notification_proto_goTypes = nil
	file_protochallenge_challenge_notification_proto_depIdxs = nil
}