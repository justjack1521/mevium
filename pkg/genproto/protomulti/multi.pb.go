// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protomulti/multi.proto

package protomulti

import (
	protoidentity "github.com/justjack1521/mevium/pkg/genproto/protoidentity"
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

type GamePlayerActionType int32

const (
	GamePlayerActionType_PLAYER_ACTION_TYPE_NONE GamePlayerActionType = 0
	GamePlayerActionType_NORMAL_ATTACK           GamePlayerActionType = 1
	GamePlayerActionType_ABILITY_CAST            GamePlayerActionType = 2
	GamePlayerActionType_ELEMENT_DRIVE           GamePlayerActionType = 3
)

// Enum value maps for GamePlayerActionType.
var (
	GamePlayerActionType_name = map[int32]string{
		0: "PLAYER_ACTION_TYPE_NONE",
		1: "NORMAL_ATTACK",
		2: "ABILITY_CAST",
		3: "ELEMENT_DRIVE",
	}
	GamePlayerActionType_value = map[string]int32{
		"PLAYER_ACTION_TYPE_NONE": 0,
		"NORMAL_ATTACK":           1,
		"ABILITY_CAST":            2,
		"ELEMENT_DRIVE":           3,
	}
)

func (x GamePlayerActionType) Enum() *GamePlayerActionType {
	p := new(GamePlayerActionType)
	*p = x
	return p
}

func (x GamePlayerActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GamePlayerActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_protomulti_multi_proto_enumTypes[0].Descriptor()
}

func (GamePlayerActionType) Type() protoreflect.EnumType {
	return &file_protomulti_multi_proto_enumTypes[0]
}

func (x GamePlayerActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GamePlayerActionType.Descriptor instead.
func (GamePlayerActionType) EnumDescriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{0}
}

type ProtoLobbySummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId         string                  `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	QuestId            string                  `protobuf:"bytes,2,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	Comment            string                  `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	MinimumPlayerLevel int32                   `protobuf:"varint,4,opt,name=minimum_player_level,json=minimumPlayerLevel,proto3" json:"minimum_player_level,omitempty"`
	RegisteredAt       int32                   `protobuf:"varint,5,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at,omitempty"`
	Players            []*ProtoLobbyPlayerSlot `protobuf:"bytes,6,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *ProtoLobbySummary) Reset() {
	*x = ProtoLobbySummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbySummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbySummary) ProtoMessage() {}

func (x *ProtoLobbySummary) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoLobbySummary.ProtoReflect.Descriptor instead.
func (*ProtoLobbySummary) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoLobbySummary) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *ProtoLobbySummary) GetQuestId() string {
	if x != nil {
		return x.QuestId
	}
	return ""
}

func (x *ProtoLobbySummary) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *ProtoLobbySummary) GetMinimumPlayerLevel() int32 {
	if x != nil {
		return x.MinimumPlayerLevel
	}
	return 0
}

func (x *ProtoLobbySummary) GetRegisteredAt() int32 {
	if x != nil {
		return x.RegisteredAt
	}
	return 0
}

func (x *ProtoLobbySummary) GetPlayers() []*ProtoLobbyPlayerSlot {
	if x != nil {
		return x.Players
	}
	return nil
}

type ProtoLobbyPlayerSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotIndex int32             `protobuf:"varint,2,opt,name=slot_index,json=slotIndex,proto3" json:"slot_index,omitempty"`
	Ready     bool              `protobuf:"varint,3,opt,name=ready,proto3" json:"ready,omitempty"`
	Player    *ProtoLobbyPlayer `protobuf:"bytes,4,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *ProtoLobbyPlayerSlot) Reset() {
	*x = ProtoLobbyPlayerSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbyPlayerSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbyPlayerSlot) ProtoMessage() {}

func (x *ProtoLobbyPlayerSlot) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoLobbyPlayerSlot.ProtoReflect.Descriptor instead.
func (*ProtoLobbyPlayerSlot) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoLobbyPlayerSlot) GetSlotIndex() int32 {
	if x != nil {
		return x.SlotIndex
	}
	return 0
}

func (x *ProtoLobbyPlayerSlot) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *ProtoLobbyPlayerSlot) GetPlayer() *ProtoLobbyPlayer {
	if x != nil {
		return x.Player
	}
	return nil
}

type ProtoLobbyPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *protoidentity.ProtoPlayerIdentity        `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Loadout  *protoidentity.ProtoPlayerLoadoutIdentity `protobuf:"bytes,2,opt,name=loadout,proto3" json:"loadout,omitempty"`
}

func (x *ProtoLobbyPlayer) Reset() {
	*x = ProtoLobbyPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbyPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbyPlayer) ProtoMessage() {}

func (x *ProtoLobbyPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoLobbyPlayer.ProtoReflect.Descriptor instead.
func (*ProtoLobbyPlayer) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoLobbyPlayer) GetIdentity() *protoidentity.ProtoPlayerIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *ProtoLobbyPlayer) GetLoadout() *protoidentity.ProtoPlayerLoadoutIdentity {
	if x != nil {
		return x.Loadout
	}
	return nil
}

type ProtoLobbyPlayerSlotRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotIndex       int32  `protobuf:"varint,1,opt,name=slot_index,json=slotIndex,proto3" json:"slot_index,omitempty"`
	RoleRestriction string `protobuf:"bytes,2,opt,name=role_restriction,json=roleRestriction,proto3" json:"role_restriction,omitempty"`
	Locked          bool   `protobuf:"varint,3,opt,name=locked,proto3" json:"locked,omitempty"`
	Bot             bool   `protobuf:"varint,4,opt,name=bot,proto3" json:"bot,omitempty"`
	InviteOnly      bool   `protobuf:"varint,5,opt,name=invite_only,json=inviteOnly,proto3" json:"invite_only,omitempty"`
}

func (x *ProtoLobbyPlayerSlotRestriction) Reset() {
	*x = ProtoLobbyPlayerSlotRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbyPlayerSlotRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbyPlayerSlotRestriction) ProtoMessage() {}

func (x *ProtoLobbyPlayerSlotRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoLobbyPlayerSlotRestriction.ProtoReflect.Descriptor instead.
func (*ProtoLobbyPlayerSlotRestriction) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoLobbyPlayerSlotRestriction) GetSlotIndex() int32 {
	if x != nil {
		return x.SlotIndex
	}
	return 0
}

func (x *ProtoLobbyPlayerSlotRestriction) GetRoleRestriction() string {
	if x != nil {
		return x.RoleRestriction
	}
	return ""
}

func (x *ProtoLobbyPlayerSlotRestriction) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *ProtoLobbyPlayerSlotRestriction) GetBot() bool {
	if x != nil {
		return x.Bot
	}
	return false
}

func (x *ProtoLobbyPlayerSlotRestriction) GetInviteOnly() bool {
	if x != nil {
		return x.InviteOnly
	}
	return false
}

type ProtoGameInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysId        string `protobuf:"bytes,1,opt,name=sys_id,json=sysId,proto3" json:"sys_id,omitempty"`
	PartyId      string `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Seed         int64  `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
	State        int32  `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	StartedAt    int64  `protobuf:"varint,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	RegisteredAt int64  `protobuf:"varint,20,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at,omitempty"`
}

func (x *ProtoGameInstance) Reset() {
	*x = ProtoGameInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoGameInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoGameInstance) ProtoMessage() {}

func (x *ProtoGameInstance) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoGameInstance.ProtoReflect.Descriptor instead.
func (*ProtoGameInstance) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{4}
}

func (x *ProtoGameInstance) GetSysId() string {
	if x != nil {
		return x.SysId
	}
	return ""
}

func (x *ProtoGameInstance) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *ProtoGameInstance) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *ProtoGameInstance) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *ProtoGameInstance) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *ProtoGameInstance) GetRegisteredAt() int64 {
	if x != nil {
		return x.RegisteredAt
	}
	return 0
}

type ProtoGameParticipant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartySlot  int32                             `protobuf:"varint,1,opt,name=party_slot,json=partySlot,proto3" json:"party_slot,omitempty"`
	BotControl bool                              `protobuf:"varint,2,opt,name=bot_control,json=botControl,proto3" json:"bot_control,omitempty"`
	Loadout    *protoidentity.ProtoPlayerLoadout `protobuf:"bytes,3,opt,name=loadout,proto3" json:"loadout,omitempty"`
}

func (x *ProtoGameParticipant) Reset() {
	*x = ProtoGameParticipant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoGameParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoGameParticipant) ProtoMessage() {}

func (x *ProtoGameParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoGameParticipant.ProtoReflect.Descriptor instead.
func (*ProtoGameParticipant) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{5}
}

func (x *ProtoGameParticipant) GetPartySlot() int32 {
	if x != nil {
		return x.PartySlot
	}
	return 0
}

func (x *ProtoGameParticipant) GetBotControl() bool {
	if x != nil {
		return x.BotControl
	}
	return false
}

func (x *ProtoGameParticipant) GetLoadout() *protoidentity.ProtoPlayerLoadout {
	if x != nil {
		return x.Loadout
	}
	return nil
}

var File_protomulti_multi_proto protoreflect.FileDescriptor

var file_protomulti_multi_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x7c, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f,
	0x75, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64,
	0x6f, 0x75, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6c, 0x6f,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xb3, 0x01, 0x0a,
	0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x79, 0x73, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x61, 0x6d, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x36, 0x0a, 0x07, 0x6c,
	0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64,
	0x6f, 0x75, 0x74, 0x2a, 0x6b, 0x0a, 0x14, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x50,
	0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x52, 0x4d,
	0x41, 0x4c, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x41, 0x53, 0x54, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x10, 0x03,
	0x42, 0x4d, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69,
	0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0xaa, 0x02, 0x12, 0x4d, 0x6f, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protomulti_multi_proto_rawDescOnce sync.Once
	file_protomulti_multi_proto_rawDescData = file_protomulti_multi_proto_rawDesc
)

func file_protomulti_multi_proto_rawDescGZIP() []byte {
	file_protomulti_multi_proto_rawDescOnce.Do(func() {
		file_protomulti_multi_proto_rawDescData = protoimpl.X.CompressGZIP(file_protomulti_multi_proto_rawDescData)
	})
	return file_protomulti_multi_proto_rawDescData
}

var file_protomulti_multi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protomulti_multi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protomulti_multi_proto_goTypes = []interface{}{
	(GamePlayerActionType)(0),                        // 0: multi.GamePlayerActionType
	(*ProtoLobbySummary)(nil),                        // 1: multi.ProtoLobbySummary
	(*ProtoLobbyPlayerSlot)(nil),                     // 2: multi.ProtoLobbyPlayerSlot
	(*ProtoLobbyPlayer)(nil),                         // 3: multi.ProtoLobbyPlayer
	(*ProtoLobbyPlayerSlotRestriction)(nil),          // 4: multi.ProtoLobbyPlayerSlotRestriction
	(*ProtoGameInstance)(nil),                        // 5: multi.ProtoGameInstance
	(*ProtoGameParticipant)(nil),                     // 6: multi.ProtoGameParticipant
	(*protoidentity.ProtoPlayerIdentity)(nil),        // 7: identity.ProtoPlayerIdentity
	(*protoidentity.ProtoPlayerLoadoutIdentity)(nil), // 8: identity.ProtoPlayerLoadoutIdentity
	(*protoidentity.ProtoPlayerLoadout)(nil),         // 9: identity.ProtoPlayerLoadout
}
var file_protomulti_multi_proto_depIdxs = []int32{
	2, // 0: multi.ProtoLobbySummary.players:type_name -> multi.ProtoLobbyPlayerSlot
	3, // 1: multi.ProtoLobbyPlayerSlot.player:type_name -> multi.ProtoLobbyPlayer
	7, // 2: multi.ProtoLobbyPlayer.identity:type_name -> identity.ProtoPlayerIdentity
	8, // 3: multi.ProtoLobbyPlayer.loadout:type_name -> identity.ProtoPlayerLoadoutIdentity
	9, // 4: multi.ProtoGameParticipant.loadout:type_name -> identity.ProtoPlayerLoadout
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protomulti_multi_proto_init() }
func file_protomulti_multi_proto_init() {
	if File_protomulti_multi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protomulti_multi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoLobbySummary); i {
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
		file_protomulti_multi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoLobbyPlayerSlot); i {
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
		file_protomulti_multi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoLobbyPlayer); i {
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
		file_protomulti_multi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoLobbyPlayerSlotRestriction); i {
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
		file_protomulti_multi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoGameInstance); i {
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
		file_protomulti_multi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoGameParticipant); i {
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
			RawDescriptor: file_protomulti_multi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protomulti_multi_proto_goTypes,
		DependencyIndexes: file_protomulti_multi_proto_depIdxs,
		EnumInfos:         file_protomulti_multi_proto_enumTypes,
		MessageInfos:      file_protomulti_multi_proto_msgTypes,
	}.Build()
	File_protomulti_multi_proto = out.File
	file_protomulti_multi_proto_rawDesc = nil
	file_protomulti_multi_proto_goTypes = nil
	file_protomulti_multi_proto_depIdxs = nil
}
