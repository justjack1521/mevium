// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: game.proto

package game

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

type ProtoMailBoxItem_ReferenceSource int32

const (
	ProtoMailBoxItem_NONE            ProtoMailBoxItem_ReferenceSource = 0
	ProtoMailBoxItem_ABILITY_CARD    ProtoMailBoxItem_ReferenceSource = 1
	ProtoMailBoxItem_ITEM            ProtoMailBoxItem_ReferenceSource = 2
	ProtoMailBoxItem_JOB_CARD        ProtoMailBoxItem_ReferenceSource = 3
	ProtoMailBoxItem_WEAPON          ProtoMailBoxItem_ReferenceSource = 4
	ProtoMailBoxItem_FAIRY_COMPANION ProtoMailBoxItem_ReferenceSource = 5
	ProtoMailBoxItem_MP_STAMP        ProtoMailBoxItem_ReferenceSource = 6
)

// Enum value maps for ProtoMailBoxItem_ReferenceSource.
var (
	ProtoMailBoxItem_ReferenceSource_name = map[int32]string{
		0: "NONE",
		1: "ABILITY_CARD",
		2: "ITEM",
		3: "JOB_CARD",
		4: "WEAPON",
		5: "FAIRY_COMPANION",
		6: "MP_STAMP",
	}
	ProtoMailBoxItem_ReferenceSource_value = map[string]int32{
		"NONE":            0,
		"ABILITY_CARD":    1,
		"ITEM":            2,
		"JOB_CARD":        3,
		"WEAPON":          4,
		"FAIRY_COMPANION": 5,
		"MP_STAMP":        6,
	}
)

func (x ProtoMailBoxItem_ReferenceSource) Enum() *ProtoMailBoxItem_ReferenceSource {
	p := new(ProtoMailBoxItem_ReferenceSource)
	*p = x
	return p
}

func (x ProtoMailBoxItem_ReferenceSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoMailBoxItem_ReferenceSource) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (ProtoMailBoxItem_ReferenceSource) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x ProtoMailBoxItem_ReferenceSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtoMailBoxItem_ReferenceSource.Descriptor instead.
func (ProtoMailBoxItem_ReferenceSource) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5, 0}
}

type ProtoPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID      string               `protobuf:"bytes,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlayerName    string               `protobuf:"bytes,2,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerComment string               `protobuf:"bytes,3,opt,name=PlayerComment,proto3" json:"PlayerComment,omitempty"`
	PlayerLevel   int32                `protobuf:"varint,4,opt,name=PlayerLevel,proto3" json:"PlayerLevel,omitempty"`
	LastOnline    int64                `protobuf:"varint,5,opt,name=LastOnline,proto3" json:"LastOnline,omitempty"`
	JobCardID     string               `protobuf:"bytes,6,opt,name=JobCardID,proto3" json:"JobCardID,omitempty"`
	SubJobIndex   int32                `protobuf:"varint,7,opt,name=SubJobIndex,proto3" json:"SubJobIndex,omitempty"`
	CompanionID   string               `protobuf:"bytes,8,opt,name=CompanionID,proto3" json:"CompanionID,omitempty"`
	RentalCard    *ProtoPlayerCardInfo `protobuf:"bytes,9,opt,name=RentalCard,proto3" json:"RentalCard,omitempty"`
}

func (x *ProtoPlayerInfo) Reset() {
	*x = ProtoPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerInfo) ProtoMessage() {}

func (x *ProtoPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
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
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoPlayerInfo) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
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

func (x *ProtoPlayerInfo) GetLastOnline() int64 {
	if x != nil {
		return x.LastOnline
	}
	return 0
}

func (x *ProtoPlayerInfo) GetJobCardID() string {
	if x != nil {
		return x.JobCardID
	}
	return ""
}

func (x *ProtoPlayerInfo) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

func (x *ProtoPlayerInfo) GetCompanionID() string {
	if x != nil {
		return x.CompanionID
	}
	return ""
}

func (x *ProtoPlayerInfo) GetRentalCard() *ProtoPlayerCardInfo {
	if x != nil {
		return x.RentalCard
	}
	return nil
}

type ProtoPlayerCardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityCardID    string `protobuf:"bytes,1,opt,name=AbilityCardID,proto3" json:"AbilityCardID,omitempty"`
	AbilityLevel     int32  `protobuf:"varint,3,opt,name=AbilityLevel,proto3" json:"AbilityLevel,omitempty"`
	AbilityCardLevel int32  `protobuf:"varint,2,opt,name=AbilityCardLevel,proto3" json:"AbilityCardLevel,omitempty"`
	ExtraSkillUnlock int32  `protobuf:"varint,4,opt,name=ExtraSkillUnlock,proto3" json:"ExtraSkillUnlock,omitempty"`
	OverBoostLevel   int32  `protobuf:"varint,5,opt,name=OverBoostLevel,proto3" json:"OverBoostLevel,omitempty"`
	SlotIndex        int32  `protobuf:"varint,6,opt,name=SlotIndex,proto3" json:"SlotIndex,omitempty"`
}

func (x *ProtoPlayerCardInfo) Reset() {
	*x = ProtoPlayerCardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerCardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerCardInfo) ProtoMessage() {}

func (x *ProtoPlayerCardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
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
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoPlayerCardInfo) GetAbilityCardID() string {
	if x != nil {
		return x.AbilityCardID
	}
	return ""
}

func (x *ProtoPlayerCardInfo) GetAbilityLevel() int32 {
	if x != nil {
		return x.AbilityLevel
	}
	return 0
}

func (x *ProtoPlayerCardInfo) GetAbilityCardLevel() int32 {
	if x != nil {
		return x.AbilityCardLevel
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

type ProtoAbilityCardInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysId            string `protobuf:"bytes,1,opt,name=SysId,proto3" json:"SysId,omitempty"`
	CardLevel        int32  `protobuf:"varint,2,opt,name=CardLevel,proto3" json:"CardLevel,omitempty"`
	AccumExp         uint64 `protobuf:"varint,3,opt,name=AccumExp,proto3" json:"AccumExp,omitempty"`
	CardLock         bool   `protobuf:"varint,4,opt,name=CardLock,proto3" json:"CardLock,omitempty"`
	AbilityLevel     int32  `protobuf:"varint,5,opt,name=AbilityLevel,proto3" json:"AbilityLevel,omitempty"`
	SkillSeedUnlock  int32  `protobuf:"varint,6,opt,name=SkillSeedUnlock,proto3" json:"SkillSeedUnlock,omitempty"`
	AbilityCardId    string `protobuf:"bytes,7,opt,name=AbilityCardId,proto3" json:"AbilityCardId,omitempty"`
	CreatedAt        int64  `protobuf:"varint,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	TotalAccumExp    uint64 `protobuf:"varint,9,opt,name=TotalAccumExp,proto3" json:"TotalAccumExp,omitempty"`
	ExtraSkillUnlock int32  `protobuf:"varint,10,opt,name=ExtraSkillUnlock,proto3" json:"ExtraSkillUnlock,omitempty"`
	OverBoostLevel   int32  `protobuf:"varint,11,opt,name=OverBoostLevel,proto3" json:"OverBoostLevel,omitempty"`
}

func (x *ProtoAbilityCardInstance) Reset() {
	*x = ProtoAbilityCardInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoAbilityCardInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoAbilityCardInstance) ProtoMessage() {}

func (x *ProtoAbilityCardInstance) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoAbilityCardInstance.ProtoReflect.Descriptor instead.
func (*ProtoAbilityCardInstance) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoAbilityCardInstance) GetSysId() string {
	if x != nil {
		return x.SysId
	}
	return ""
}

func (x *ProtoAbilityCardInstance) GetCardLevel() int32 {
	if x != nil {
		return x.CardLevel
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetAccumExp() uint64 {
	if x != nil {
		return x.AccumExp
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetCardLock() bool {
	if x != nil {
		return x.CardLock
	}
	return false
}

func (x *ProtoAbilityCardInstance) GetAbilityLevel() int32 {
	if x != nil {
		return x.AbilityLevel
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetSkillSeedUnlock() int32 {
	if x != nil {
		return x.SkillSeedUnlock
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetAbilityCardId() string {
	if x != nil {
		return x.AbilityCardId
	}
	return ""
}

func (x *ProtoAbilityCardInstance) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetTotalAccumExp() uint64 {
	if x != nil {
		return x.TotalAccumExp
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetExtraSkillUnlock() int32 {
	if x != nil {
		return x.ExtraSkillUnlock
	}
	return 0
}

func (x *ProtoAbilityCardInstance) GetOverBoostLevel() int32 {
	if x != nil {
		return x.OverBoostLevel
	}
	return 0
}

type ProtoItemInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysId     string `protobuf:"bytes,1,opt,name=SysId,proto3" json:"SysId,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	ItemId    string `protobuf:"bytes,3,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *ProtoItemInstance) Reset() {
	*x = ProtoItemInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoItemInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoItemInstance) ProtoMessage() {}

func (x *ProtoItemInstance) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoItemInstance.ProtoReflect.Descriptor instead.
func (*ProtoItemInstance) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoItemInstance) GetSysId() string {
	if x != nil {
		return x.SysId
	}
	return ""
}

func (x *ProtoItemInstance) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProtoItemInstance) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *ProtoItemInstance) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ItemValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID string `protobuf:"bytes,1,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	Value  int32  `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ItemValuePair) Reset() {
	*x = ItemValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemValuePair) ProtoMessage() {}

func (x *ItemValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemValuePair.ProtoReflect.Descriptor instead.
func (*ItemValuePair) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *ItemValuePair) GetItemID() string {
	if x != nil {
		return x.ItemID
	}
	return ""
}

func (x *ItemValuePair) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ProtoMailBoxItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceID string                           `protobuf:"bytes,1,opt,name=ReferenceID,proto3" json:"ReferenceID,omitempty"`
	Source      ProtoMailBoxItem_ReferenceSource `protobuf:"varint,2,opt,name=Source,proto3,enum=game.ProtoMailBoxItem_ReferenceSource" json:"Source,omitempty"`
	Quantity    int32                            `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	ReceivedAt  int64                            `protobuf:"varint,4,opt,name=ReceivedAt,proto3" json:"ReceivedAt,omitempty"`
}

func (x *ProtoMailBoxItem) Reset() {
	*x = ProtoMailBoxItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMailBoxItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMailBoxItem) ProtoMessage() {}

func (x *ProtoMailBoxItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMailBoxItem.ProtoReflect.Descriptor instead.
func (*ProtoMailBoxItem) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *ProtoMailBoxItem) GetReferenceID() string {
	if x != nil {
		return x.ReferenceID
	}
	return ""
}

func (x *ProtoMailBoxItem) GetSource() ProtoMailBoxItem_ReferenceSource {
	if x != nil {
		return x.Source
	}
	return ProtoMailBoxItem_NONE
}

func (x *ProtoMailBoxItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProtoMailBoxItem) GetReceivedAt() int64 {
	if x != nil {
		return x.ReceivedAt
	}
	return 0
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0xd2, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61,
	0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x4c, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4a, 0x6f,
	0x62, 0x43, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4a,
	0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53,
	0x75, 0x62, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0a,
	0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x52, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x22, 0xfd, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x24, 0x0a, 0x0d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x42,
	0x6f, 0x6f, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6c, 0x6f,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x6c,
	0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x92, 0x03, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x79, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x79, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61,
	0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43,
	0x61, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x75,
	0x6d, 0x45, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x41, 0x63, 0x63, 0x75,
	0x6d, 0x45, 0x78, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x6f, 0x63, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x6f, 0x63, 0x6b,
	0x12, 0x22, 0x0a, 0x0c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x65,
	0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x65, 0x64, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x24,
	0x0a, 0x0d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x75, 0x6d,
	0x45, 0x78, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x63, 0x63, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73,
	0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x4f, 0x76,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x7b, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x79, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x79, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x49, 0x74, 0x65,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa6, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a,
	0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x3e, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6c,
	0x42, 0x6f, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x0f, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x54,
	0x45, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x41, 0x52, 0x44,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x41, 0x50, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x46, 0x41, 0x49, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x49, 0x4f,
	0x4e, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10,
	0x06, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76,
	0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_game_proto_goTypes = []interface{}{
	(ProtoMailBoxItem_ReferenceSource)(0), // 0: game.ProtoMailBoxItem.ReferenceSource
	(*ProtoPlayerInfo)(nil),               // 1: game.ProtoPlayerInfo
	(*ProtoPlayerCardInfo)(nil),           // 2: game.ProtoPlayerCardInfo
	(*ProtoAbilityCardInstance)(nil),      // 3: game.ProtoAbilityCardInstance
	(*ProtoItemInstance)(nil),             // 4: game.ProtoItemInstance
	(*ItemValuePair)(nil),                 // 5: game.ItemValuePair
	(*ProtoMailBoxItem)(nil),              // 6: game.ProtoMailBoxItem
}
var file_game_proto_depIdxs = []int32{
	2, // 0: game.ProtoPlayerInfo.RentalCard:type_name -> game.ProtoPlayerCardInfo
	0, // 1: game.ProtoMailBoxItem.Source:type_name -> game.ProtoMailBoxItem.ReferenceSource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoAbilityCardInstance); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoItemInstance); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemValuePair); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMailBoxItem); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
