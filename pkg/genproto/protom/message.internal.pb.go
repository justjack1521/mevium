// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: message.internal.proto

package protom

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

type ClientConnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteAddress string `protobuf:"bytes,1,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *ClientConnected) Reset() {
	*x = ClientConnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnected) ProtoMessage() {}

func (x *ClientConnected) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnected.ProtoReflect.Descriptor instead.
func (*ClientConnected) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConnected) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

type ClientHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteAddress string `protobuf:"bytes,1,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *ClientHeartbeat) Reset() {
	*x = ClientHeartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientHeartbeat) ProtoMessage() {}

func (x *ClientHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientHeartbeat.ProtoReflect.Descriptor instead.
func (*ClientHeartbeat) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{1}
}

func (x *ClientHeartbeat) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

type ClientDisconnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteAddress string `protobuf:"bytes,1,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
}

func (x *ClientDisconnected) Reset() {
	*x = ClientDisconnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientDisconnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDisconnected) ProtoMessage() {}

func (x *ClientDisconnected) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDisconnected.ProtoReflect.Descriptor instead.
func (*ClientDisconnected) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{2}
}

func (x *ClientDisconnected) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

type PlayerProfileCreatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId      string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName    string `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel   int32  `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	PlayerComment string `protobuf:"bytes,4,opt,name=player_comment,json=playerComment,proto3" json:"player_comment,omitempty"`
	CompanionId   string `protobuf:"bytes,5,opt,name=companion_id,json=companionId,proto3" json:"companion_id,omitempty"`
}

func (x *PlayerProfileCreatedMessage) Reset() {
	*x = PlayerProfileCreatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerProfileCreatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerProfileCreatedMessage) ProtoMessage() {}

func (x *PlayerProfileCreatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerProfileCreatedMessage.ProtoReflect.Descriptor instead.
func (*PlayerProfileCreatedMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerProfileCreatedMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerProfileCreatedMessage) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerProfileCreatedMessage) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *PlayerProfileCreatedMessage) GetPlayerComment() string {
	if x != nil {
		return x.PlayerComment
	}
	return ""
}

func (x *PlayerProfileCreatedMessage) GetCompanionId() string {
	if x != nil {
		return x.CompanionId
	}
	return ""
}

type PlayerPresenceUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Online   uint64 `protobuf:"varint,2,opt,name=online,proto3" json:"online,omitempty"`
}

func (x *PlayerPresenceUpdateMessage) Reset() {
	*x = PlayerPresenceUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPresenceUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPresenceUpdateMessage) ProtoMessage() {}

func (x *PlayerPresenceUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPresenceUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerPresenceUpdateMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerPresenceUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerPresenceUpdateMessage) GetOnline() uint64 {
	if x != nil {
		return x.Online
	}
	return 0
}

type PlayerPositionUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId      string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RegionMap     string `protobuf:"bytes,2,opt,name=region_map,json=regionMap,proto3" json:"region_map,omitempty"`
	RegionMapNode int32  `protobuf:"varint,3,opt,name=region_map_node,json=regionMapNode,proto3" json:"region_map_node,omitempty"`
}

func (x *PlayerPositionUpdateMessage) Reset() {
	*x = PlayerPositionUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPositionUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPositionUpdateMessage) ProtoMessage() {}

func (x *PlayerPositionUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPositionUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerPositionUpdateMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerPositionUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerPositionUpdateMessage) GetRegionMap() string {
	if x != nil {
		return x.RegionMap
	}
	return ""
}

func (x *PlayerPositionUpdateMessage) GetRegionMapNode() int32 {
	if x != nil {
		return x.RegionMapNode
	}
	return 0
}

type PlayerCompanionUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId    string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	CompanionId string `protobuf:"bytes,2,opt,name=companion_id,json=companionId,proto3" json:"companion_id,omitempty"`
}

func (x *PlayerCompanionUpdateMessage) Reset() {
	*x = PlayerCompanionUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerCompanionUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerCompanionUpdateMessage) ProtoMessage() {}

func (x *PlayerCompanionUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerCompanionUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerCompanionUpdateMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{6}
}

func (x *PlayerCompanionUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerCompanionUpdateMessage) GetCompanionId() string {
	if x != nil {
		return x.CompanionId
	}
	return ""
}

type PlayerCommentUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Comment  string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *PlayerCommentUpdateMessage) Reset() {
	*x = PlayerCommentUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerCommentUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerCommentUpdateMessage) ProtoMessage() {}

func (x *PlayerCommentUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerCommentUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerCommentUpdateMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{7}
}

func (x *PlayerCommentUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerCommentUpdateMessage) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type PlayerRentalCardUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RentalCard string `protobuf:"bytes,2,opt,name=rental_card,json=rentalCard,proto3" json:"rental_card,omitempty"`
}

func (x *PlayerRentalCardUpdateMessage) Reset() {
	*x = PlayerRentalCardUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRentalCardUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRentalCardUpdateMessage) ProtoMessage() {}

func (x *PlayerRentalCardUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRentalCardUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerRentalCardUpdateMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{8}
}

func (x *PlayerRentalCardUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerRentalCardUpdateMessage) GetRentalCard() string {
	if x != nil {
		return x.RentalCard
	}
	return ""
}

type PlayerLevelUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Level    int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *PlayerLevelUpdateMessage) Reset() {
	*x = PlayerLevelUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_internal_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLevelUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLevelUpdateMessage) ProtoMessage() {}

func (x *PlayerLevelUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_internal_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLevelUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerLevelUpdateMessage) Descriptor() ([]byte, []int) {
	return file_message_internal_proto_rawDescGZIP(), []int{9}
}

func (x *PlayerLevelUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerLevelUpdateMessage) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_message_internal_proto protoreflect.FileDescriptor

var file_message_internal_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x38, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x12, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x52, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x5e, 0x0a, 0x1c, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x1a, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x5d, 0x0a, 0x1d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x43,
	0x61, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x22, 0x4d,
	0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74,
	0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_internal_proto_rawDescOnce sync.Once
	file_message_internal_proto_rawDescData = file_message_internal_proto_rawDesc
)

func file_message_internal_proto_rawDescGZIP() []byte {
	file_message_internal_proto_rawDescOnce.Do(func() {
		file_message_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_internal_proto_rawDescData)
	})
	return file_message_internal_proto_rawDescData
}

var file_message_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_message_internal_proto_goTypes = []interface{}{
	(*ClientConnected)(nil),               // 0: notification.ClientConnected
	(*ClientHeartbeat)(nil),               // 1: notification.ClientHeartbeat
	(*ClientDisconnected)(nil),            // 2: notification.ClientDisconnected
	(*PlayerProfileCreatedMessage)(nil),   // 3: notification.PlayerProfileCreatedMessage
	(*PlayerPresenceUpdateMessage)(nil),   // 4: notification.PlayerPresenceUpdateMessage
	(*PlayerPositionUpdateMessage)(nil),   // 5: notification.PlayerPositionUpdateMessage
	(*PlayerCompanionUpdateMessage)(nil),  // 6: notification.PlayerCompanionUpdateMessage
	(*PlayerCommentUpdateMessage)(nil),    // 7: notification.PlayerCommentUpdateMessage
	(*PlayerRentalCardUpdateMessage)(nil), // 8: notification.PlayerRentalCardUpdateMessage
	(*PlayerLevelUpdateMessage)(nil),      // 9: notification.PlayerLevelUpdateMessage
}
var file_message_internal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_internal_proto_init() }
func file_message_internal_proto_init() {
	if File_message_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConnected); i {
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
		file_message_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientHeartbeat); i {
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
		file_message_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientDisconnected); i {
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
		file_message_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerProfileCreatedMessage); i {
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
		file_message_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPresenceUpdateMessage); i {
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
		file_message_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPositionUpdateMessage); i {
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
		file_message_internal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerCompanionUpdateMessage); i {
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
		file_message_internal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerCommentUpdateMessage); i {
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
		file_message_internal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRentalCardUpdateMessage); i {
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
		file_message_internal_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLevelUpdateMessage); i {
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
			RawDescriptor: file_message_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_internal_proto_goTypes,
		DependencyIndexes: file_message_internal_proto_depIdxs,
		MessageInfos:      file_message_internal_proto_msgTypes,
	}.Build()
	File_message_internal_proto = out.File
	file_message_internal_proto_rawDesc = nil
	file_message_internal_proto_goTypes = nil
	file_message_internal_proto_depIdxs = nil
}
