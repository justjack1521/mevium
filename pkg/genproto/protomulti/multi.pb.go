// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protomulti/multi.proto

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

	PlayerId      string                   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName    string                   `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerComment string                   `protobuf:"bytes,3,opt,name=player_comment,json=playerComment,proto3" json:"player_comment,omitempty"`
	PlayerLevel   int32                    `protobuf:"varint,4,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	JobCard       *ProtoLobbyJobCard       `protobuf:"bytes,5,opt,name=job_card,json=jobCard,proto3" json:"job_card,omitempty"`
	Weapon        *ProtoLobbyWeapon        `protobuf:"bytes,6,opt,name=weapon,proto3" json:"weapon,omitempty"`
	AbilityCards  []*ProtoLobbyAbilityCard `protobuf:"bytes,7,rep,name=ability_cards,json=abilityCards,proto3" json:"ability_cards,omitempty"`
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

func (x *ProtoLobbyPlayer) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *ProtoLobbyPlayer) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *ProtoLobbyPlayer) GetPlayerComment() string {
	if x != nil {
		return x.PlayerComment
	}
	return ""
}

func (x *ProtoLobbyPlayer) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *ProtoLobbyPlayer) GetJobCard() *ProtoLobbyJobCard {
	if x != nil {
		return x.JobCard
	}
	return nil
}

func (x *ProtoLobbyPlayer) GetWeapon() *ProtoLobbyWeapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

func (x *ProtoLobbyPlayer) GetAbilityCards() []*ProtoLobbyAbilityCard {
	if x != nil {
		return x.AbilityCards
	}
	return nil
}

type ProtoLobbyJobCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobCardId      string `protobuf:"bytes,1,opt,name=job_card_id,json=jobCardId,proto3" json:"job_card_id,omitempty"`
	SubJobIndex    int32  `protobuf:"varint,2,opt,name=sub_job_index,json=subJobIndex,proto3" json:"sub_job_index,omitempty"`
	OverBoostLevel int32  `protobuf:"varint,3,opt,name=over_boost_level,json=overBoostLevel,proto3" json:"over_boost_level,omitempty"`
}

func (x *ProtoLobbyJobCard) Reset() {
	*x = ProtoLobbyJobCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbyJobCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbyJobCard) ProtoMessage() {}

func (x *ProtoLobbyJobCard) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProtoLobbyJobCard.ProtoReflect.Descriptor instead.
func (*ProtoLobbyJobCard) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoLobbyJobCard) GetJobCardId() string {
	if x != nil {
		return x.JobCardId
	}
	return ""
}

func (x *ProtoLobbyJobCard) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

func (x *ProtoLobbyJobCard) GetOverBoostLevel() int32 {
	if x != nil {
		return x.OverBoostLevel
	}
	return 0
}

type ProtoLobbyWeapon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeaponId        string `protobuf:"bytes,1,opt,name=weapon_id,json=weaponId,proto3" json:"weapon_id,omitempty"`
	SubWeaponUnlock int32  `protobuf:"varint,2,opt,name=sub_weapon_unlock,json=subWeaponUnlock,proto3" json:"sub_weapon_unlock,omitempty"`
}

func (x *ProtoLobbyWeapon) Reset() {
	*x = ProtoLobbyWeapon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbyWeapon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbyWeapon) ProtoMessage() {}

func (x *ProtoLobbyWeapon) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProtoLobbyWeapon.ProtoReflect.Descriptor instead.
func (*ProtoLobbyWeapon) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{4}
}

func (x *ProtoLobbyWeapon) GetWeaponId() string {
	if x != nil {
		return x.WeaponId
	}
	return ""
}

func (x *ProtoLobbyWeapon) GetSubWeaponUnlock() int32 {
	if x != nil {
		return x.SubWeaponUnlock
	}
	return 0
}

type ProtoLobbyAbilityCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityCardId    string `protobuf:"bytes,1,opt,name=ability_card_id,json=abilityCardId,proto3" json:"ability_card_id,omitempty"`
	AbilityCardLevel int32  `protobuf:"varint,2,opt,name=ability_card_level,json=abilityCardLevel,proto3" json:"ability_card_level,omitempty"`
	AbilityLevel     int32  `protobuf:"varint,3,opt,name=ability_level,json=abilityLevel,proto3" json:"ability_level,omitempty"`
	OverBoostLevel   int32  `protobuf:"varint,4,opt,name=over_boost_level,json=overBoostLevel,proto3" json:"over_boost_level,omitempty"`
	SlotIndex        int32  `protobuf:"varint,20,opt,name=slot_index,json=slotIndex,proto3" json:"slot_index,omitempty"`
}

func (x *ProtoLobbyAbilityCard) Reset() {
	*x = ProtoLobbyAbilityCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoLobbyAbilityCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoLobbyAbilityCard) ProtoMessage() {}

func (x *ProtoLobbyAbilityCard) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProtoLobbyAbilityCard.ProtoReflect.Descriptor instead.
func (*ProtoLobbyAbilityCard) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_proto_rawDescGZIP(), []int{5}
}

func (x *ProtoLobbyAbilityCard) GetAbilityCardId() string {
	if x != nil {
		return x.AbilityCardId
	}
	return ""
}

func (x *ProtoLobbyAbilityCard) GetAbilityCardLevel() int32 {
	if x != nil {
		return x.AbilityCardLevel
	}
	return 0
}

func (x *ProtoLobbyAbilityCard) GetAbilityLevel() int32 {
	if x != nil {
		return x.AbilityLevel
	}
	return 0
}

func (x *ProtoLobbyAbilityCard) GetOverBoostLevel() int32 {
	if x != nil {
		return x.OverBoostLevel
	}
	return 0
}

func (x *ProtoLobbyAbilityCard) GetSlotIndex() int32 {
	if x != nil {
		return x.SlotIndex
	}
	return 0
}

var File_protomulti_multi_proto protoreflect.FileDescriptor

var file_protomulti_multi_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x22,
	0xf7, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x7c, 0x0a, 0x14, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0xc3, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x07, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x77, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x57, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x0c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x73, 0x22, 0x81, 0x01,
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4a, 0x6f, 0x62, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x5b, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x57,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73,
	0x75, 0x62, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xdb,
	0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x73,
	0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f,
	0x76, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x4d, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a,
	0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0xaa, 0x02, 0x12, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_protomulti_multi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protomulti_multi_proto_goTypes = []interface{}{
	(*ProtoLobbySummary)(nil),     // 0: multi.ProtoLobbySummary
	(*ProtoLobbyPlayerSlot)(nil),  // 1: multi.ProtoLobbyPlayerSlot
	(*ProtoLobbyPlayer)(nil),      // 2: multi.ProtoLobbyPlayer
	(*ProtoLobbyJobCard)(nil),     // 3: multi.ProtoLobbyJobCard
	(*ProtoLobbyWeapon)(nil),      // 4: multi.ProtoLobbyWeapon
	(*ProtoLobbyAbilityCard)(nil), // 5: multi.ProtoLobbyAbilityCard
}
var file_protomulti_multi_proto_depIdxs = []int32{
	1, // 0: multi.ProtoLobbySummary.players:type_name -> multi.ProtoLobbyPlayerSlot
	2, // 1: multi.ProtoLobbyPlayerSlot.player:type_name -> multi.ProtoLobbyPlayer
	3, // 2: multi.ProtoLobbyPlayer.job_card:type_name -> multi.ProtoLobbyJobCard
	4, // 3: multi.ProtoLobbyPlayer.weapon:type_name -> multi.ProtoLobbyWeapon
	5, // 4: multi.ProtoLobbyPlayer.ability_cards:type_name -> multi.ProtoLobbyAbilityCard
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
			switch v := v.(*ProtoLobbyJobCard); i {
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
			switch v := v.(*ProtoLobbyWeapon); i {
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
			switch v := v.(*ProtoLobbyAbilityCard); i {
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
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protomulti_multi_proto_goTypes,
		DependencyIndexes: file_protomulti_multi_proto_depIdxs,
		MessageInfos:      file_protomulti_multi_proto_msgTypes,
	}.Build()
	File_protomulti_multi_proto = out.File
	file_protomulti_multi_proto_rawDesc = nil
	file_protomulti_multi_proto_goTypes = nil
	file_protomulti_multi_proto_depIdxs = nil
}
