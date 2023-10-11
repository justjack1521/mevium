// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protosocial/social.proto

package protosocial

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

type ProtoPlayerIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId    string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName  string `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel int32  `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	IsPlayer    bool   `protobuf:"varint,4,opt,name=is_player,json=isPlayer,proto3" json:"is_player,omitempty"`
}

func (x *ProtoPlayerIdentity) Reset() {
	*x = ProtoPlayerIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosocial_social_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerIdentity) ProtoMessage() {}

func (x *ProtoPlayerIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_protosocial_social_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerIdentity.ProtoReflect.Descriptor instead.
func (*ProtoPlayerIdentity) Descriptor() ([]byte, []int) {
	return file_protosocial_social_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoPlayerIdentity) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *ProtoPlayerIdentity) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *ProtoPlayerIdentity) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *ProtoPlayerIdentity) GetIsPlayer() bool {
	if x != nil {
		return x.IsPlayer
	}
	return false
}

type ProtoPlayerSocialInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerInfo      *ProtoPlayerInfo `protobuf:"bytes,1,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	LastOnline      int64            `protobuf:"varint,2,opt,name=last_online,json=lastOnline,proto3" json:"last_online,omitempty"`
	RegionMapId     string           `protobuf:"bytes,3,opt,name=region_map_id,json=regionMapId,proto3" json:"region_map_id,omitempty"`
	RegionNodeIndex int32            `protobuf:"varint,4,opt,name=region_node_index,json=regionNodeIndex,proto3" json:"region_node_index,omitempty"`
}

func (x *ProtoPlayerSocialInfo) Reset() {
	*x = ProtoPlayerSocialInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosocial_social_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerSocialInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerSocialInfo) ProtoMessage() {}

func (x *ProtoPlayerSocialInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protosocial_social_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerSocialInfo.ProtoReflect.Descriptor instead.
func (*ProtoPlayerSocialInfo) Descriptor() ([]byte, []int) {
	return file_protosocial_social_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoPlayerSocialInfo) GetPlayerInfo() *ProtoPlayerInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *ProtoPlayerSocialInfo) GetLastOnline() int64 {
	if x != nil {
		return x.LastOnline
	}
	return 0
}

func (x *ProtoPlayerSocialInfo) GetRegionMapId() string {
	if x != nil {
		return x.RegionMapId
	}
	return ""
}

func (x *ProtoPlayerSocialInfo) GetRegionNodeIndex() int32 {
	if x != nil {
		return x.RegionNodeIndex
	}
	return 0
}

type ProtoPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId      string               `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName    string               `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerComment string               `protobuf:"bytes,3,opt,name=player_comment,json=playerComment,proto3" json:"player_comment,omitempty"`
	PlayerLevel   int32                `protobuf:"varint,4,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	JobCardId     string               `protobuf:"bytes,6,opt,name=job_card_id,json=jobCardId,proto3" json:"job_card_id,omitempty"`
	SubJobIndex   int32                `protobuf:"varint,7,opt,name=sub_job_index,json=subJobIndex,proto3" json:"sub_job_index,omitempty"`
	CompanionId   string               `protobuf:"bytes,8,opt,name=companion_id,json=companionId,proto3" json:"companion_id,omitempty"`
	RentalCard    *ProtoPlayerCardInfo `protobuf:"bytes,9,opt,name=rental_card,json=rentalCard,proto3" json:"rental_card,omitempty"`
}

func (x *ProtoPlayerInfo) Reset() {
	*x = ProtoPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosocial_social_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerInfo) ProtoMessage() {}

func (x *ProtoPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protosocial_social_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerInfo.ProtoReflect.Descriptor instead.
func (*ProtoPlayerInfo) Descriptor() ([]byte, []int) {
	return file_protosocial_social_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoPlayerInfo) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *ProtoPlayerInfo) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *ProtoPlayerInfo) GetPlayerComment() string {
	if x != nil {
		return x.PlayerComment
	}
	return ""
}

func (x *ProtoPlayerInfo) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *ProtoPlayerInfo) GetJobCardId() string {
	if x != nil {
		return x.JobCardId
	}
	return ""
}

func (x *ProtoPlayerInfo) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

func (x *ProtoPlayerInfo) GetCompanionId() string {
	if x != nil {
		return x.CompanionId
	}
	return ""
}

func (x *ProtoPlayerInfo) GetRentalCard() *ProtoPlayerCardInfo {
	if x != nil {
		return x.RentalCard
	}
	return nil
}

type ProtoPlayerLoadout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobCardId       string                 `protobuf:"bytes,1,opt,name=job_card_id,json=jobCardId,proto3" json:"job_card_id,omitempty"`
	SubJobIndex     int32                  `protobuf:"varint,2,opt,name=sub_job_index,json=subJobIndex,proto3" json:"sub_job_index,omitempty"`
	WeaponId        string                 `protobuf:"bytes,3,opt,name=weapon_id,json=weaponId,proto3" json:"weapon_id,omitempty"`
	SubWeaponUnlock int32                  `protobuf:"varint,4,opt,name=sub_weapon_unlock,json=subWeaponUnlock,proto3" json:"sub_weapon_unlock,omitempty"`
	AbilityCards    []*ProtoPlayerCardInfo `protobuf:"bytes,5,rep,name=ability_cards,json=abilityCards,proto3" json:"ability_cards,omitempty"`
}

func (x *ProtoPlayerLoadout) Reset() {
	*x = ProtoPlayerLoadout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosocial_social_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerLoadout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerLoadout) ProtoMessage() {}

func (x *ProtoPlayerLoadout) ProtoReflect() protoreflect.Message {
	mi := &file_protosocial_social_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerLoadout.ProtoReflect.Descriptor instead.
func (*ProtoPlayerLoadout) Descriptor() ([]byte, []int) {
	return file_protosocial_social_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoPlayerLoadout) GetJobCardId() string {
	if x != nil {
		return x.JobCardId
	}
	return ""
}

func (x *ProtoPlayerLoadout) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

func (x *ProtoPlayerLoadout) GetWeaponId() string {
	if x != nil {
		return x.WeaponId
	}
	return ""
}

func (x *ProtoPlayerLoadout) GetSubWeaponUnlock() int32 {
	if x != nil {
		return x.SubWeaponUnlock
	}
	return 0
}

func (x *ProtoPlayerLoadout) GetAbilityCards() []*ProtoPlayerCardInfo {
	if x != nil {
		return x.AbilityCards
	}
	return nil
}

type ProtoPlayerCardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityCardId    string `protobuf:"bytes,1,opt,name=ability_card_id,json=abilityCardId,proto3" json:"ability_card_id,omitempty"`
	AbilityCardLevel int32  `protobuf:"varint,2,opt,name=ability_card_level,json=abilityCardLevel,proto3" json:"ability_card_level,omitempty"`
	AbilityLevel     int32  `protobuf:"varint,3,opt,name=ability_level,json=abilityLevel,proto3" json:"ability_level,omitempty"`
	ExtraSkillUnlock int32  `protobuf:"varint,4,opt,name=extra_skill_unlock,json=extraSkillUnlock,proto3" json:"extra_skill_unlock,omitempty"`
	OverBoostLevel   int32  `protobuf:"varint,5,opt,name=over_boost_level,json=overBoostLevel,proto3" json:"over_boost_level,omitempty"`
	SlotIndex        int32  `protobuf:"varint,6,opt,name=slot_index,json=slotIndex,proto3" json:"slot_index,omitempty"`
}

func (x *ProtoPlayerCardInfo) Reset() {
	*x = ProtoPlayerCardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosocial_social_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerCardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerCardInfo) ProtoMessage() {}

func (x *ProtoPlayerCardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protosocial_social_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerCardInfo.ProtoReflect.Descriptor instead.
func (*ProtoPlayerCardInfo) Descriptor() ([]byte, []int) {
	return file_protosocial_social_proto_rawDescGZIP(), []int{4}
}

func (x *ProtoPlayerCardInfo) GetAbilityCardId() string {
	if x != nil {
		return x.AbilityCardId
	}
	return ""
}

func (x *ProtoPlayerCardInfo) GetAbilityCardLevel() int32 {
	if x != nil {
		return x.AbilityCardLevel
	}
	return 0
}

func (x *ProtoPlayerCardInfo) GetAbilityLevel() int32 {
	if x != nil {
		return x.AbilityLevel
	}
	return 0
}

func (x *ProtoPlayerCardInfo) GetExtraSkillUnlock() int32 {
	if x != nil {
		return x.ExtraSkillUnlock
	}
	return 0
}

func (x *ProtoPlayerCardInfo) GetOverBoostLevel() int32 {
	if x != nil {
		return x.OverBoostLevel
	}
	return 0
}

func (x *ProtoPlayerCardInfo) GetSlotIndex() int32 {
	if x != nil {
		return x.SlotIndex
	}
	return 0
}

var File_protosocial_social_proto protoreflect.FileDescriptor

var file_protosocial_social_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0xc4, 0x01, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x4d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0xc0, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a,
	0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x73, 0x75, 0x62, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x43, 0x61, 0x72, 0x64, 0x22, 0xe5, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6a,
	0x6f, 0x62, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73,
	0x75, 0x62, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1b, 0x0a, 0x09, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x57, 0x65, 0x61, 0x70,
	0x6f, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x73, 0x22, 0x87, 0x02, 0x0a,
	0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x2c, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a,
	0x10, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6c, 0x6f,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x4f, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32,
	0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0xaa, 0x02, 0x13, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protosocial_social_proto_rawDescOnce sync.Once
	file_protosocial_social_proto_rawDescData = file_protosocial_social_proto_rawDesc
)

func file_protosocial_social_proto_rawDescGZIP() []byte {
	file_protosocial_social_proto_rawDescOnce.Do(func() {
		file_protosocial_social_proto_rawDescData = protoimpl.X.CompressGZIP(file_protosocial_social_proto_rawDescData)
	})
	return file_protosocial_social_proto_rawDescData
}

var file_protosocial_social_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protosocial_social_proto_goTypes = []interface{}{
	(*ProtoPlayerIdentity)(nil),   // 0: presence.ProtoPlayerIdentity
	(*ProtoPlayerSocialInfo)(nil), // 1: presence.ProtoPlayerSocialInfo
	(*ProtoPlayerInfo)(nil),       // 2: presence.ProtoPlayerInfo
	(*ProtoPlayerLoadout)(nil),    // 3: presence.ProtoPlayerLoadout
	(*ProtoPlayerCardInfo)(nil),   // 4: presence.ProtoPlayerCardInfo
}
var file_protosocial_social_proto_depIdxs = []int32{
	2, // 0: presence.ProtoPlayerSocialInfo.player_info:type_name -> presence.ProtoPlayerInfo
	4, // 1: presence.ProtoPlayerInfo.rental_card:type_name -> presence.ProtoPlayerCardInfo
	4, // 2: presence.ProtoPlayerLoadout.ability_cards:type_name -> presence.ProtoPlayerCardInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protosocial_social_proto_init() }
func file_protosocial_social_proto_init() {
	if File_protosocial_social_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protosocial_social_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerIdentity); i {
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
		file_protosocial_social_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerSocialInfo); i {
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
		file_protosocial_social_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerInfo); i {
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
		file_protosocial_social_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerLoadout); i {
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
		file_protosocial_social_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerCardInfo); i {
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
			RawDescriptor: file_protosocial_social_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protosocial_social_proto_goTypes,
		DependencyIndexes: file_protosocial_social_proto_depIdxs,
		MessageInfos:      file_protosocial_social_proto_msgTypes,
	}.Build()
	File_protosocial_social_proto = out.File
	file_protosocial_social_proto_rawDesc = nil
	file_protosocial_social_proto_goTypes = nil
	file_protosocial_social_proto_depIdxs = nil
}