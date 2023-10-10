// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: challenge/challenge.proto

package protochallenge

import (
	protogame "github.com/justjack1521/mevium/pkg/genproto/protogame"
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

type ProtoSocialChallengeInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysId         string                             `protobuf:"bytes,1,opt,name=sys_id,json=sysId,proto3" json:"sys_id,omitempty"`
	ChallengeId   string                             `protobuf:"bytes,2,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	MissionId     string                             `protobuf:"bytes,3,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	RegisterTime  int64                              `protobuf:"varint,4,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"`
	StartTime     int64                              `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       int64                              `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	State         int32                              `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	PlayerEntries []*ProtoPlayerSocialChallengeEntry `protobuf:"bytes,8,rep,name=player_entries,json=playerEntries,proto3" json:"player_entries,omitempty"`
}

func (x *ProtoSocialChallengeInstance) Reset() {
	*x = ProtoSocialChallengeInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_challenge_challenge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoSocialChallengeInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoSocialChallengeInstance) ProtoMessage() {}

func (x *ProtoSocialChallengeInstance) ProtoReflect() protoreflect.Message {
	mi := &file_challenge_challenge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoSocialChallengeInstance.ProtoReflect.Descriptor instead.
func (*ProtoSocialChallengeInstance) Descriptor() ([]byte, []int) {
	return file_challenge_challenge_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoSocialChallengeInstance) GetSysId() string {
	if x != nil {
		return x.SysId
	}
	return ""
}

func (x *ProtoSocialChallengeInstance) GetChallengeId() string {
	if x != nil {
		return x.ChallengeId
	}
	return ""
}

func (x *ProtoSocialChallengeInstance) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

func (x *ProtoSocialChallengeInstance) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

func (x *ProtoSocialChallengeInstance) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *ProtoSocialChallengeInstance) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ProtoSocialChallengeInstance) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *ProtoSocialChallengeInstance) GetPlayerEntries() []*ProtoPlayerSocialChallengeEntry {
	if x != nil {
		return x.PlayerEntries
	}
	return nil
}

type ProtoPlayerSocialChallengeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId       string                        `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName     string                        `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel    int32                         `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	Score          int64                         `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Rank           int64                         `protobuf:"varint,5,opt,name=rank,proto3" json:"rank,omitempty"`
	PrimaryLoadout *protogame.ProtoPlayerLoadout `protobuf:"bytes,6,opt,name=primary_loadout,json=primaryLoadout,proto3" json:"primary_loadout,omitempty"`
	IsPlayer       bool                          `protobuf:"varint,7,opt,name=is_player,json=isPlayer,proto3" json:"is_player,omitempty"`
}

func (x *ProtoPlayerSocialChallengeEntry) Reset() {
	*x = ProtoPlayerSocialChallengeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_challenge_challenge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerSocialChallengeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerSocialChallengeEntry) ProtoMessage() {}

func (x *ProtoPlayerSocialChallengeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_challenge_challenge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerSocialChallengeEntry.ProtoReflect.Descriptor instead.
func (*ProtoPlayerSocialChallengeEntry) Descriptor() ([]byte, []int) {
	return file_challenge_challenge_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoPlayerSocialChallengeEntry) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *ProtoPlayerSocialChallengeEntry) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *ProtoPlayerSocialChallengeEntry) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *ProtoPlayerSocialChallengeEntry) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ProtoPlayerSocialChallengeEntry) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *ProtoPlayerSocialChallengeEntry) GetPrimaryLoadout() *protogame.ProtoPlayerLoadout {
	if x != nil {
		return x.PrimaryLoadout
	}
	return nil
}

func (x *ProtoPlayerSocialChallengeEntry) GetIsPlayer() bool {
	if x != nil {
		return x.IsPlayer
	}
	return false
}

var File_challenge_challenge_proto protoreflect.FileDescriptor

var file_challenge_challenge_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x1a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x79, 0x73, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x8c, 0x02, 0x0a, 0x1f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x41, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x0e, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x4b, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31,
	0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0xaa, 0x02, 0x0c, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_challenge_challenge_proto_rawDescOnce sync.Once
	file_challenge_challenge_proto_rawDescData = file_challenge_challenge_proto_rawDesc
)

func file_challenge_challenge_proto_rawDescGZIP() []byte {
	file_challenge_challenge_proto_rawDescOnce.Do(func() {
		file_challenge_challenge_proto_rawDescData = protoimpl.X.CompressGZIP(file_challenge_challenge_proto_rawDescData)
	})
	return file_challenge_challenge_proto_rawDescData
}

var file_challenge_challenge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_challenge_challenge_proto_goTypes = []interface{}{
	(*ProtoSocialChallengeInstance)(nil),    // 0: challenge.ProtoSocialChallengeInstance
	(*ProtoPlayerSocialChallengeEntry)(nil), // 1: challenge.ProtoPlayerSocialChallengeEntry
	(*protogame.ProtoPlayerLoadout)(nil),    // 2: game.ProtoPlayerLoadout
}
var file_challenge_challenge_proto_depIdxs = []int32{
	1, // 0: challenge.ProtoSocialChallengeInstance.player_entries:type_name -> challenge.ProtoPlayerSocialChallengeEntry
	2, // 1: challenge.ProtoPlayerSocialChallengeEntry.primary_loadout:type_name -> game.ProtoPlayerLoadout
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_challenge_challenge_proto_init() }
func file_challenge_challenge_proto_init() {
	if File_challenge_challenge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_challenge_challenge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoSocialChallengeInstance); i {
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
		file_challenge_challenge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerSocialChallengeEntry); i {
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
			RawDescriptor: file_challenge_challenge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_challenge_challenge_proto_goTypes,
		DependencyIndexes: file_challenge_challenge_proto_depIdxs,
		MessageInfos:      file_challenge_challenge_proto_msgTypes,
	}.Build()
	File_challenge_challenge_proto = out.File
	file_challenge_challenge_proto_rawDesc = nil
	file_challenge_challenge_proto_goTypes = nil
	file_challenge_challenge_proto_depIdxs = nil
}
