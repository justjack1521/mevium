// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protoplayer/player.proto

package protoplayer

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

type ProtoPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId    string              `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName  string              `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel int32               `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	Loadout     *ProtoPlayerLoadout `protobuf:"bytes,4,opt,name=loadout,proto3" json:"loadout,omitempty"`
}

func (x *ProtoPlayerInfo) Reset() {
	*x = ProtoPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoplayer_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerInfo) ProtoMessage() {}

func (x *ProtoPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protoplayer_player_proto_msgTypes[0]
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
	return file_protoplayer_player_proto_rawDescGZIP(), []int{0}
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

func (x *ProtoPlayerInfo) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *ProtoPlayerInfo) GetLoadout() *ProtoPlayerLoadout {
	if x != nil {
		return x.Loadout
	}
	return nil
}

type ProtoPlayerLoadout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckIndex    int32               `protobuf:"varint,1,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
	JobCard      *ProtoJobCard       `protobuf:"bytes,2,opt,name=job_card,json=jobCard,proto3" json:"job_card,omitempty"`
	Weapon       *ProtoWeapon        `protobuf:"bytes,3,opt,name=weapon,proto3" json:"weapon,omitempty"`
	AbilityCards []*ProtoAbilityCard `protobuf:"bytes,5,rep,name=ability_cards,json=abilityCards,proto3" json:"ability_cards,omitempty"`
}

func (x *ProtoPlayerLoadout) Reset() {
	*x = ProtoPlayerLoadout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoplayer_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerLoadout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerLoadout) ProtoMessage() {}

func (x *ProtoPlayerLoadout) ProtoReflect() protoreflect.Message {
	mi := &file_protoplayer_player_proto_msgTypes[1]
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
	return file_protoplayer_player_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoPlayerLoadout) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

func (x *ProtoPlayerLoadout) GetJobCard() *ProtoJobCard {
	if x != nil {
		return x.JobCard
	}
	return nil
}

func (x *ProtoPlayerLoadout) GetWeapon() *ProtoWeapon {
	if x != nil {
		return x.Weapon
	}
	return nil
}

func (x *ProtoPlayerLoadout) GetAbilityCards() []*ProtoAbilityCard {
	if x != nil {
		return x.AbilityCards
	}
	return nil
}

type ProtoJobCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseJobId         string           `protobuf:"bytes,1,opt,name=base_job_id,json=baseJobId,proto3" json:"base_job_id,omitempty"`
	SubJobIndex       int32            `protobuf:"varint,2,opt,name=sub_job_index,json=subJobIndex,proto3" json:"sub_job_index,omitempty"`
	HpStatMod         int32            `protobuf:"varint,3,opt,name=hp_stat_mod,json=hpStatMod,proto3" json:"hp_stat_mod,omitempty"`
	AttackStatMod     int32            `protobuf:"varint,4,opt,name=attack_stat_mod,json=attackStatMod,proto3" json:"attack_stat_mod,omitempty"`
	BreakStatMod      int32            `protobuf:"varint,5,opt,name=break_stat_mod,json=breakStatMod,proto3" json:"break_stat_mod,omitempty"`
	MagicStatMod      int32            `protobuf:"varint,6,opt,name=magic_stat_mod,json=magicStatMod,proto3" json:"magic_stat_mod,omitempty"`
	SpeedStatMod      int32            `protobuf:"varint,7,opt,name=speed_stat_mod,json=speedStatMod,proto3" json:"speed_stat_mod,omitempty"`
	DefenseStatMod    int32            `protobuf:"varint,8,opt,name=defense_stat_mod,json=defenseStatMod,proto3" json:"defense_stat_mod,omitempty"`
	CritChanceStatMod int32            `protobuf:"varint,9,opt,name=crit_chance_stat_mod,json=critChanceStatMod,proto3" json:"crit_chance_stat_mod,omitempty"`
	UltimateBoost     int32            `protobuf:"varint,10,opt,name=ultimate_boost,json=ultimateBoost,proto3" json:"ultimate_boost,omitempty"`
	OverBoostLevel    int32            `protobuf:"varint,11,opt,name=over_boost_level,json=overBoostLevel,proto3" json:"over_boost_level,omitempty"`
	AutoAbilities     map[string]int32 `protobuf:"bytes,12,rep,name=auto_abilities,json=autoAbilities,proto3" json:"auto_abilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ProtoJobCard) Reset() {
	*x = ProtoJobCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoplayer_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoJobCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoJobCard) ProtoMessage() {}

func (x *ProtoJobCard) ProtoReflect() protoreflect.Message {
	mi := &file_protoplayer_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoJobCard.ProtoReflect.Descriptor instead.
func (*ProtoJobCard) Descriptor() ([]byte, []int) {
	return file_protoplayer_player_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoJobCard) GetBaseJobId() string {
	if x != nil {
		return x.BaseJobId
	}
	return ""
}

func (x *ProtoJobCard) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

func (x *ProtoJobCard) GetHpStatMod() int32 {
	if x != nil {
		return x.HpStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetAttackStatMod() int32 {
	if x != nil {
		return x.AttackStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetBreakStatMod() int32 {
	if x != nil {
		return x.BreakStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetMagicStatMod() int32 {
	if x != nil {
		return x.MagicStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetSpeedStatMod() int32 {
	if x != nil {
		return x.SpeedStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetDefenseStatMod() int32 {
	if x != nil {
		return x.DefenseStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetCritChanceStatMod() int32 {
	if x != nil {
		return x.CritChanceStatMod
	}
	return 0
}

func (x *ProtoJobCard) GetUltimateBoost() int32 {
	if x != nil {
		return x.UltimateBoost
	}
	return 0
}

func (x *ProtoJobCard) GetOverBoostLevel() int32 {
	if x != nil {
		return x.OverBoostLevel
	}
	return 0
}

func (x *ProtoJobCard) GetAutoAbilities() map[string]int32 {
	if x != nil {
		return x.AutoAbilities
	}
	return nil
}

type ProtoWeapon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseWeaponId      string           `protobuf:"bytes,1,opt,name=base_weapon_id,json=baseWeaponId,proto3" json:"base_weapon_id,omitempty"`
	SubWeaponUnlock   int32            `protobuf:"varint,2,opt,name=sub_weapon_unlock,json=subWeaponUnlock,proto3" json:"sub_weapon_unlock,omitempty"`
	HpStatMod         int32            `protobuf:"varint,3,opt,name=hp_stat_mod,json=hpStatMod,proto3" json:"hp_stat_mod,omitempty"`
	AttackStatMod     int32            `protobuf:"varint,4,opt,name=attack_stat_mod,json=attackStatMod,proto3" json:"attack_stat_mod,omitempty"`
	BreakStatMod      int32            `protobuf:"varint,5,opt,name=break_stat_mod,json=breakStatMod,proto3" json:"break_stat_mod,omitempty"`
	MagicStatMod      int32            `protobuf:"varint,6,opt,name=magic_stat_mod,json=magicStatMod,proto3" json:"magic_stat_mod,omitempty"`
	SpeedStatMod      int32            `protobuf:"varint,7,opt,name=speed_stat_mod,json=speedStatMod,proto3" json:"speed_stat_mod,omitempty"`
	DefenseStatMod    int32            `protobuf:"varint,8,opt,name=defense_stat_mod,json=defenseStatMod,proto3" json:"defense_stat_mod,omitempty"`
	CritChanceStatMod int32            `protobuf:"varint,9,opt,name=crit_chance_stat_mod,json=critChanceStatMod,proto3" json:"crit_chance_stat_mod,omitempty"`
	UltimateBoost     int32            `protobuf:"varint,10,opt,name=ultimate_boost,json=ultimateBoost,proto3" json:"ultimate_boost,omitempty"`
	AutoAbilities     map[string]int32 `protobuf:"bytes,11,rep,name=auto_abilities,json=autoAbilities,proto3" json:"auto_abilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ProtoWeapon) Reset() {
	*x = ProtoWeapon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoplayer_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoWeapon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoWeapon) ProtoMessage() {}

func (x *ProtoWeapon) ProtoReflect() protoreflect.Message {
	mi := &file_protoplayer_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoWeapon.ProtoReflect.Descriptor instead.
func (*ProtoWeapon) Descriptor() ([]byte, []int) {
	return file_protoplayer_player_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoWeapon) GetBaseWeaponId() string {
	if x != nil {
		return x.BaseWeaponId
	}
	return ""
}

func (x *ProtoWeapon) GetSubWeaponUnlock() int32 {
	if x != nil {
		return x.SubWeaponUnlock
	}
	return 0
}

func (x *ProtoWeapon) GetHpStatMod() int32 {
	if x != nil {
		return x.HpStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetAttackStatMod() int32 {
	if x != nil {
		return x.AttackStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetBreakStatMod() int32 {
	if x != nil {
		return x.BreakStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetMagicStatMod() int32 {
	if x != nil {
		return x.MagicStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetSpeedStatMod() int32 {
	if x != nil {
		return x.SpeedStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetDefenseStatMod() int32 {
	if x != nil {
		return x.DefenseStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetCritChanceStatMod() int32 {
	if x != nil {
		return x.CritChanceStatMod
	}
	return 0
}

func (x *ProtoWeapon) GetUltimateBoost() int32 {
	if x != nil {
		return x.UltimateBoost
	}
	return 0
}

func (x *ProtoWeapon) GetAutoAbilities() map[string]int32 {
	if x != nil {
		return x.AutoAbilities
	}
	return nil
}

type ProtoAbilityCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityCardId    string           `protobuf:"bytes,1,opt,name=ability_card_id,json=abilityCardId,proto3" json:"ability_card_id,omitempty"`
	AbilityCardLevel int32            `protobuf:"varint,2,opt,name=ability_card_level,json=abilityCardLevel,proto3" json:"ability_card_level,omitempty"`
	AbilityLevel     int32            `protobuf:"varint,3,opt,name=ability_level,json=abilityLevel,proto3" json:"ability_level,omitempty"`
	ExtraSkillUnlock int32            `protobuf:"varint,4,opt,name=extra_skill_unlock,json=extraSkillUnlock,proto3" json:"extra_skill_unlock,omitempty"`
	OverBoostLevel   int32            `protobuf:"varint,5,opt,name=over_boost_level,json=overBoostLevel,proto3" json:"over_boost_level,omitempty"`
	AutoAbilities    map[string]int32 `protobuf:"bytes,6,rep,name=auto_abilities,json=autoAbilities,proto3" json:"auto_abilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SlotIndex        int32            `protobuf:"varint,20,opt,name=slot_index,json=slotIndex,proto3" json:"slot_index,omitempty"`
}

func (x *ProtoAbilityCard) Reset() {
	*x = ProtoAbilityCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoplayer_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoAbilityCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoAbilityCard) ProtoMessage() {}

func (x *ProtoAbilityCard) ProtoReflect() protoreflect.Message {
	mi := &file_protoplayer_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoAbilityCard.ProtoReflect.Descriptor instead.
func (*ProtoAbilityCard) Descriptor() ([]byte, []int) {
	return file_protoplayer_player_proto_rawDescGZIP(), []int{4}
}

func (x *ProtoAbilityCard) GetAbilityCardId() string {
	if x != nil {
		return x.AbilityCardId
	}
	return ""
}

func (x *ProtoAbilityCard) GetAbilityCardLevel() int32 {
	if x != nil {
		return x.AbilityCardLevel
	}
	return 0
}

func (x *ProtoAbilityCard) GetAbilityLevel() int32 {
	if x != nil {
		return x.AbilityLevel
	}
	return 0
}

func (x *ProtoAbilityCard) GetExtraSkillUnlock() int32 {
	if x != nil {
		return x.ExtraSkillUnlock
	}
	return 0
}

func (x *ProtoAbilityCard) GetOverBoostLevel() int32 {
	if x != nil {
		return x.OverBoostLevel
	}
	return 0
}

func (x *ProtoAbilityCard) GetAutoAbilities() map[string]int32 {
	if x != nil {
		return x.AutoAbilities
	}
	return nil
}

func (x *ProtoAbilityCard) GetSlotIndex() int32 {
	if x != nil {
		return x.SlotIndex
	}
	return 0
}

var File_protoplayer_player_proto protoreflect.FileDescriptor

var file_protoplayer_player_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x34, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x6f,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61,
	0x64, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x22, 0xd0, 0x01,
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61,
	0x64, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x2f, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x52, 0x07, 0x6a, 0x6f, 0x62,
	0x43, 0x61, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x0c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x22, 0xca, 0x04, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4a, 0x6f, 0x62, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0b, 0x68, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x4d, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x67,
	0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x65, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x72, 0x69,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6c,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x76, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x4e, 0x0a, 0x0e, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x75,
	0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x41,
	0x75, 0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xac, 0x04,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x57, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x73, 0x75, 0x62, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x1e, 0x0a, 0x0b, 0x68, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x70, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x4d, 0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x4d, 0x6f, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x72, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x63, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x4d, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x75, 0x6c,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0e, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x75, 0x74,
	0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x41, 0x75,
	0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9a, 0x03, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61,
	0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x12,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x52, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6c,
	0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x40, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x4f, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b,
	0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0xaa, 0x02, 0x13, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protoplayer_player_proto_rawDescOnce sync.Once
	file_protoplayer_player_proto_rawDescData = file_protoplayer_player_proto_rawDesc
)

func file_protoplayer_player_proto_rawDescGZIP() []byte {
	file_protoplayer_player_proto_rawDescOnce.Do(func() {
		file_protoplayer_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoplayer_player_proto_rawDescData)
	})
	return file_protoplayer_player_proto_rawDescData
}

var file_protoplayer_player_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protoplayer_player_proto_goTypes = []interface{}{
	(*ProtoPlayerInfo)(nil),    // 0: player.ProtoPlayerInfo
	(*ProtoPlayerLoadout)(nil), // 1: player.ProtoPlayerLoadout
	(*ProtoJobCard)(nil),       // 2: player.ProtoJobCard
	(*ProtoWeapon)(nil),        // 3: player.ProtoWeapon
	(*ProtoAbilityCard)(nil),   // 4: player.ProtoAbilityCard
	nil,                        // 5: player.ProtoJobCard.AutoAbilitiesEntry
	nil,                        // 6: player.ProtoWeapon.AutoAbilitiesEntry
	nil,                        // 7: player.ProtoAbilityCard.AutoAbilitiesEntry
}
var file_protoplayer_player_proto_depIdxs = []int32{
	1, // 0: player.ProtoPlayerInfo.loadout:type_name -> player.ProtoPlayerLoadout
	2, // 1: player.ProtoPlayerLoadout.job_card:type_name -> player.ProtoJobCard
	3, // 2: player.ProtoPlayerLoadout.weapon:type_name -> player.ProtoWeapon
	4, // 3: player.ProtoPlayerLoadout.ability_cards:type_name -> player.ProtoAbilityCard
	5, // 4: player.ProtoJobCard.auto_abilities:type_name -> player.ProtoJobCard.AutoAbilitiesEntry
	6, // 5: player.ProtoWeapon.auto_abilities:type_name -> player.ProtoWeapon.AutoAbilitiesEntry
	7, // 6: player.ProtoAbilityCard.auto_abilities:type_name -> player.ProtoAbilityCard.AutoAbilitiesEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protoplayer_player_proto_init() }
func file_protoplayer_player_proto_init() {
	if File_protoplayer_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoplayer_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protoplayer_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_protoplayer_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoJobCard); i {
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
		file_protoplayer_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoWeapon); i {
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
		file_protoplayer_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoAbilityCard); i {
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
			RawDescriptor: file_protoplayer_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoplayer_player_proto_goTypes,
		DependencyIndexes: file_protoplayer_player_proto_depIdxs,
		MessageInfos:      file_protoplayer_player_proto_msgTypes,
	}.Build()
	File_protoplayer_player_proto = out.File
	file_protoplayer_player_proto_rawDesc = nil
	file_protoplayer_player_proto_goTypes = nil
	file_protoplayer_player_proto_depIdxs = nil
}
