// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: message.notification.proto

package protom

import (
	protog "github.com/justjack1521/mevium/pkg/genproto/protog"
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
	NotificationType_BASE            NotificationType = 0
	NotificationType_SOCIAL_DATA     NotificationType = 100
	NotificationType_PLAYER_PRESENCE NotificationType = 101
	NotificationType_PLAYER_POSITION NotificationType = 102
	NotificationType_PLAYER_LOADOUT  NotificationType = 103
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0:   "BASE",
		100: "SOCIAL_DATA",
		101: "PLAYER_PRESENCE",
		102: "PLAYER_POSITION",
		103: "PLAYER_LOADOUT",
	}
	NotificationType_value = map[string]int32{
		"BASE":            0,
		"SOCIAL_DATA":     100,
		"PLAYER_PRESENCE": 101,
		"PLAYER_POSITION": 102,
		"PLAYER_LOADOUT":  103,
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
	return file_message_notification_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_message_notification_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_message_notification_proto_rawDescGZIP(), []int{0}
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type NotificationType `protobuf:"varint,1,opt,name=type,proto3,enum=notification.NotificationType" json:"type,omitempty"`
	Data []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_message_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_message_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_BASE
}

func (x *Notification) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SocialDataNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowingList []*ProtoPlayerSocialInfo `protobuf:"bytes,1,rep,name=following_list,json=followingList,proto3" json:"following_list,omitempty"`
	FollowerList  []*ProtoPlayerSocialInfo `protobuf:"bytes,2,rep,name=follower_list,json=followerList,proto3" json:"follower_list,omitempty"`
}

func (x *SocialDataNotification) Reset() {
	*x = SocialDataNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialDataNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialDataNotification) ProtoMessage() {}

func (x *SocialDataNotification) ProtoReflect() protoreflect.Message {
	mi := &file_message_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialDataNotification.ProtoReflect.Descriptor instead.
func (*SocialDataNotification) Descriptor() ([]byte, []int) {
	return file_message_notification_proto_rawDescGZIP(), []int{1}
}

func (x *SocialDataNotification) GetFollowingList() []*ProtoPlayerSocialInfo {
	if x != nil {
		return x.FollowingList
	}
	return nil
}

func (x *SocialDataNotification) GetFollowerList() []*ProtoPlayerSocialInfo {
	if x != nil {
		return x.FollowerList
	}
	return nil
}

type PlayerPresenceNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	LastOnline int64  `protobuf:"varint,2,opt,name=last_online,json=lastOnline,proto3" json:"last_online,omitempty"`
}

func (x *PlayerPresenceNotification) Reset() {
	*x = PlayerPresenceNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPresenceNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPresenceNotification) ProtoMessage() {}

func (x *PlayerPresenceNotification) ProtoReflect() protoreflect.Message {
	mi := &file_message_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPresenceNotification.ProtoReflect.Descriptor instead.
func (*PlayerPresenceNotification) Descriptor() ([]byte, []int) {
	return file_message_notification_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerPresenceNotification) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerPresenceNotification) GetLastOnline() int64 {
	if x != nil {
		return x.LastOnline
	}
	return 0
}

type PlayerLoadoutNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId        string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	JobCardId       string `protobuf:"bytes,2,opt,name=job_card_id,json=jobCardId,proto3" json:"job_card_id,omitempty"`
	SubJobIndex     int32  `protobuf:"varint,3,opt,name=sub_job_index,json=subJobIndex,proto3" json:"sub_job_index,omitempty"`
	WeaponId        string `protobuf:"bytes,4,opt,name=weapon_id,json=weaponId,proto3" json:"weapon_id,omitempty"`
	SubWeaponUnlock int32  `protobuf:"varint,5,opt,name=sub_weapon_unlock,json=subWeaponUnlock,proto3" json:"sub_weapon_unlock,omitempty"`
}

func (x *PlayerLoadoutNotification) Reset() {
	*x = PlayerLoadoutNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoadoutNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoadoutNotification) ProtoMessage() {}

func (x *PlayerLoadoutNotification) ProtoReflect() protoreflect.Message {
	mi := &file_message_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoadoutNotification.ProtoReflect.Descriptor instead.
func (*PlayerLoadoutNotification) Descriptor() ([]byte, []int) {
	return file_message_notification_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerLoadoutNotification) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerLoadoutNotification) GetJobCardId() string {
	if x != nil {
		return x.JobCardId
	}
	return ""
}

func (x *PlayerLoadoutNotification) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

func (x *PlayerLoadoutNotification) GetWeaponId() string {
	if x != nil {
		return x.WeaponId
	}
	return ""
}

func (x *PlayerLoadoutNotification) GetSubWeaponUnlock() int32 {
	if x != nil {
		return x.SubWeaponUnlock
	}
	return 0
}

type PlayerPositionNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RegionId  string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	NodeIndex int32  `protobuf:"varint,3,opt,name=node_index,json=nodeIndex,proto3" json:"node_index,omitempty"`
}

func (x *PlayerPositionNotification) Reset() {
	*x = PlayerPositionNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPositionNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPositionNotification) ProtoMessage() {}

func (x *PlayerPositionNotification) ProtoReflect() protoreflect.Message {
	mi := &file_message_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPositionNotification.ProtoReflect.Descriptor instead.
func (*PlayerPositionNotification) Descriptor() ([]byte, []int) {
	return file_message_notification_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerPositionNotification) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerPositionNotification) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *PlayerPositionNotification) GetNodeIndex() int32 {
	if x != nil {
		return x.NodeIndex
	}
	return 0
}

type ProtoPlayerSocialInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerInfo      *protog.ProtoPlayerInfo `protobuf:"bytes,1,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	LastOnline      int64                   `protobuf:"varint,2,opt,name=last_online,json=lastOnline,proto3" json:"last_online,omitempty"`
	RegionMapId     string                  `protobuf:"bytes,3,opt,name=region_map_id,json=regionMapId,proto3" json:"region_map_id,omitempty"`
	RegionNodeIndex int32                   `protobuf:"varint,4,opt,name=region_node_index,json=regionNodeIndex,proto3" json:"region_node_index,omitempty"`
}

func (x *ProtoPlayerSocialInfo) Reset() {
	*x = ProtoPlayerSocialInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_notification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerSocialInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerSocialInfo) ProtoMessage() {}

func (x *ProtoPlayerSocialInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_notification_proto_msgTypes[5]
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
	return file_message_notification_proto_rawDescGZIP(), []int{5}
}

func (x *ProtoPlayerSocialInfo) GetPlayerInfo() *protog.ProtoPlayerInfo {
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

var File_message_notification_proto protoreflect.FileDescriptor

var file_message_notification_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xae,
	0x01, 0x0a, 0x16, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x5a, 0x0a, 0x1a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x19,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62,
	0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65,
	0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x5f, 0x77,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0x75, 0x0a, 0x1a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xc2, 0x01, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61,
	0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2a,
	0x6b, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x64, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x43,
	0x45, 0x10, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x66, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x4f, 0x55, 0x54, 0x10, 0x67, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a,
	0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_notification_proto_rawDescOnce sync.Once
	file_message_notification_proto_rawDescData = file_message_notification_proto_rawDesc
)

func file_message_notification_proto_rawDescGZIP() []byte {
	file_message_notification_proto_rawDescOnce.Do(func() {
		file_message_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_notification_proto_rawDescData)
	})
	return file_message_notification_proto_rawDescData
}

var file_message_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_message_notification_proto_goTypes = []interface{}{
	(NotificationType)(0),              // 0: notification.NotificationType
	(*Notification)(nil),               // 1: notification.Notification
	(*SocialDataNotification)(nil),     // 2: notification.SocialDataNotification
	(*PlayerPresenceNotification)(nil), // 3: notification.PlayerPresenceNotification
	(*PlayerLoadoutNotification)(nil),  // 4: notification.PlayerLoadoutNotification
	(*PlayerPositionNotification)(nil), // 5: notification.PlayerPositionNotification
	(*ProtoPlayerSocialInfo)(nil),      // 6: notification.ProtoPlayerSocialInfo
	(*protog.ProtoPlayerInfo)(nil),     // 7: protog.ProtoPlayerInfo
}
var file_message_notification_proto_depIdxs = []int32{
	0, // 0: notification.Notification.type:type_name -> notification.NotificationType
	6, // 1: notification.SocialDataNotification.following_list:type_name -> notification.ProtoPlayerSocialInfo
	6, // 2: notification.SocialDataNotification.follower_list:type_name -> notification.ProtoPlayerSocialInfo
	7, // 3: notification.ProtoPlayerSocialInfo.player_info:type_name -> protog.ProtoPlayerInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_notification_proto_init() }
func file_message_notification_proto_init() {
	if File_message_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_message_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialDataNotification); i {
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
		file_message_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPresenceNotification); i {
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
		file_message_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoadoutNotification); i {
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
		file_message_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPositionNotification); i {
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
		file_message_notification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_notification_proto_goTypes,
		DependencyIndexes: file_message_notification_proto_depIdxs,
		EnumInfos:         file_message_notification_proto_enumTypes,
		MessageInfos:      file_message_notification_proto_msgTypes,
	}.Build()
	File_message_notification_proto = out.File
	file_message_notification_proto_rawDesc = nil
	file_message_notification_proto_goTypes = nil
	file_message_notification_proto_depIdxs = nil
}
