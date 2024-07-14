// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protomulti/multi.request.proto

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

type MultiRequestType int32

const (
	MultiRequestType_NONE           MultiRequestType = 0
	MultiRequestType_CREATE_SESSION MultiRequestType = 100
	MultiRequestType_END_SESSION    MultiRequestType = 200
	MultiRequestType_SEARCH_LOBBY   MultiRequestType = 300
	MultiRequestType_CREATE_LOBBY   MultiRequestType = 400
	MultiRequestType_JOIN_LOBBY     MultiRequestType = 500
	MultiRequestType_READY_LOBBY    MultiRequestType = 600
	MultiRequestType_WATCH_LOBBY    MultiRequestType = 700
	MultiRequestType_DISCARD_LOBBY  MultiRequestType = 800
	MultiRequestType_CANCEL_LOBBY   MultiRequestType = 900
	MultiRequestType_UNREADY_LOBBY  MultiRequestType = 1000
	MultiRequestType_SEARCH_PLAYER  MultiRequestType = 1100
	MultiRequestType_SEND_STAMP     MultiRequestType = 1200
)

// Enum value maps for MultiRequestType.
var (
	MultiRequestType_name = map[int32]string{
		0:    "NONE",
		100:  "CREATE_SESSION",
		200:  "END_SESSION",
		300:  "SEARCH_LOBBY",
		400:  "CREATE_LOBBY",
		500:  "JOIN_LOBBY",
		600:  "READY_LOBBY",
		700:  "WATCH_LOBBY",
		800:  "DISCARD_LOBBY",
		900:  "CANCEL_LOBBY",
		1000: "UNREADY_LOBBY",
		1100: "SEARCH_PLAYER",
		1200: "SEND_STAMP",
	}
	MultiRequestType_value = map[string]int32{
		"NONE":           0,
		"CREATE_SESSION": 100,
		"END_SESSION":    200,
		"SEARCH_LOBBY":   300,
		"CREATE_LOBBY":   400,
		"JOIN_LOBBY":     500,
		"READY_LOBBY":    600,
		"WATCH_LOBBY":    700,
		"DISCARD_LOBBY":  800,
		"CANCEL_LOBBY":   900,
		"UNREADY_LOBBY":  1000,
		"SEARCH_PLAYER":  1100,
		"SEND_STAMP":     1200,
	}
)

func (x MultiRequestType) Enum() *MultiRequestType {
	p := new(MultiRequestType)
	*p = x
	return p
}

func (x MultiRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MultiRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_protomulti_multi_request_proto_enumTypes[0].Descriptor()
}

func (MultiRequestType) Type() protoreflect.EnumType {
	return &file_protomulti_multi_request_proto_enumTypes[0]
}

func (x MultiRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MultiRequestType.Descriptor instead.
func (MultiRequestType) EnumDescriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{0}
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId       string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	ModeIdentifier string `protobuf:"bytes,2,opt,name=mode_identifier,json=modeIdentifier,proto3" json:"mode_identifier,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *CreateSessionRequest) GetModeIdentifier() string {
	if x != nil {
		return x.ModeIdentifier
	}
	return ""
}

type EndSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId       string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	ModeIdentifier string `protobuf:"bytes,3,opt,name=mode_identifier,json=modeIdentifier,proto3" json:"mode_identifier,omitempty"`
}

func (x *EndSessionRequest) Reset() {
	*x = EndSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionRequest) ProtoMessage() {}

func (x *EndSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionRequest.ProtoReflect.Descriptor instead.
func (*EndSessionRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{1}
}

func (x *EndSessionRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *EndSessionRequest) GetModeIdentifier() string {
	if x != nil {
		return x.ModeIdentifier
	}
	return ""
}

type SearchLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckIndex      int32    `protobuf:"varint,2,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
	ModeIdentifier string   `protobuf:"bytes,3,opt,name=mode_identifier,json=modeIdentifier,proto3" json:"mode_identifier,omitempty"`
	Levels         []int32  `protobuf:"varint,4,rep,packed,name=levels,proto3" json:"levels,omitempty"`
	Categories     []string `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
	MinPlayerLevel int32    `protobuf:"varint,6,opt,name=min_player_level,json=minPlayerLevel,proto3" json:"min_player_level,omitempty"`
	PartyId        string   `protobuf:"bytes,7,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *SearchLobbyRequest) Reset() {
	*x = SearchLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLobbyRequest) ProtoMessage() {}

func (x *SearchLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLobbyRequest.ProtoReflect.Descriptor instead.
func (*SearchLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{2}
}

func (x *SearchLobbyRequest) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

func (x *SearchLobbyRequest) GetModeIdentifier() string {
	if x != nil {
		return x.ModeIdentifier
	}
	return ""
}

func (x *SearchLobbyRequest) GetLevels() []int32 {
	if x != nil {
		return x.Levels
	}
	return nil
}

func (x *SearchLobbyRequest) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *SearchLobbyRequest) GetMinPlayerLevel() int32 {
	if x != nil {
		return x.MinPlayerLevel
	}
	return 0
}

func (x *SearchLobbyRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type WatchLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LobbyId string `protobuf:"bytes,2,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
}

func (x *WatchLobbyRequest) Reset() {
	*x = WatchLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchLobbyRequest) ProtoMessage() {}

func (x *WatchLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchLobbyRequest.ProtoReflect.Descriptor instead.
func (*WatchLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{3}
}

func (x *WatchLobbyRequest) GetLobbyId() string {
	if x != nil {
		return x.LobbyId
	}
	return ""
}

type UnwatchLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnwatchLobbyRequest) Reset() {
	*x = UnwatchLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnwatchLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnwatchLobbyRequest) ProtoMessage() {}

func (x *UnwatchLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnwatchLobbyRequest.ProtoReflect.Descriptor instead.
func (*UnwatchLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{4}
}

type CreateLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckIndex      int32                              `protobuf:"varint,2,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
	QuestId        string                             `protobuf:"bytes,3,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	Comment        string                             `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	MinPlayerLevel int32                              `protobuf:"varint,5,opt,name=min_player_level,json=minPlayerLevel,proto3" json:"min_player_level,omitempty"`
	Restrictions   []*ProtoLobbyPlayerSlotRestriction `protobuf:"bytes,6,rep,name=restrictions,proto3" json:"restrictions,omitempty"`
}

func (x *CreateLobbyRequest) Reset() {
	*x = CreateLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLobbyRequest) ProtoMessage() {}

func (x *CreateLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLobbyRequest.ProtoReflect.Descriptor instead.
func (*CreateLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{5}
}

func (x *CreateLobbyRequest) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

func (x *CreateLobbyRequest) GetQuestId() string {
	if x != nil {
		return x.QuestId
	}
	return ""
}

func (x *CreateLobbyRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CreateLobbyRequest) GetMinPlayerLevel() int32 {
	if x != nil {
		return x.MinPlayerLevel
	}
	return 0
}

func (x *CreateLobbyRequest) GetRestrictions() []*ProtoLobbyPlayerSlotRestriction {
	if x != nil {
		return x.Restrictions
	}
	return nil
}

type JoinLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LobbyId    string `protobuf:"bytes,2,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
	SlotIndex  int32  `protobuf:"varint,3,opt,name=slot_index,json=slotIndex,proto3" json:"slot_index,omitempty"`
	DeckIndex  int32  `protobuf:"varint,4,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
	UseStamina bool   `protobuf:"varint,5,opt,name=use_stamina,json=useStamina,proto3" json:"use_stamina,omitempty"`
	FromInvite bool   `protobuf:"varint,6,opt,name=from_invite,json=fromInvite,proto3" json:"from_invite,omitempty"`
}

func (x *JoinLobbyRequest) Reset() {
	*x = JoinLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLobbyRequest) ProtoMessage() {}

func (x *JoinLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLobbyRequest.ProtoReflect.Descriptor instead.
func (*JoinLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{6}
}

func (x *JoinLobbyRequest) GetLobbyId() string {
	if x != nil {
		return x.LobbyId
	}
	return ""
}

func (x *JoinLobbyRequest) GetSlotIndex() int32 {
	if x != nil {
		return x.SlotIndex
	}
	return 0
}

func (x *JoinLobbyRequest) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

func (x *JoinLobbyRequest) GetUseStamina() bool {
	if x != nil {
		return x.UseStamina
	}
	return false
}

func (x *JoinLobbyRequest) GetFromInvite() bool {
	if x != nil {
		return x.FromInvite
	}
	return false
}

type LeaveLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveLobbyRequest) Reset() {
	*x = LeaveLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveLobbyRequest) ProtoMessage() {}

func (x *LeaveLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveLobbyRequest.ProtoReflect.Descriptor instead.
func (*LeaveLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{7}
}

type StartLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartLobbyRequest) Reset() {
	*x = StartLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartLobbyRequest) ProtoMessage() {}

func (x *StartLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartLobbyRequest.ProtoReflect.Descriptor instead.
func (*StartLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{8}
}

type ReadyLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LobbyId   string `protobuf:"bytes,2,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
	DeckIndex int32  `protobuf:"varint,4,opt,name=deck_index,json=deckIndex,proto3" json:"deck_index,omitempty"`
}

func (x *ReadyLobbyRequest) Reset() {
	*x = ReadyLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyLobbyRequest) ProtoMessage() {}

func (x *ReadyLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyLobbyRequest.ProtoReflect.Descriptor instead.
func (*ReadyLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{9}
}

func (x *ReadyLobbyRequest) GetLobbyId() string {
	if x != nil {
		return x.LobbyId
	}
	return ""
}

func (x *ReadyLobbyRequest) GetDeckIndex() int32 {
	if x != nil {
		return x.DeckIndex
	}
	return 0
}

type UnreadyLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LobbyId string `protobuf:"bytes,2,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
}

func (x *UnreadyLobbyRequest) Reset() {
	*x = UnreadyLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnreadyLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnreadyLobbyRequest) ProtoMessage() {}

func (x *UnreadyLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnreadyLobbyRequest.ProtoReflect.Descriptor instead.
func (*UnreadyLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{10}
}

func (x *UnreadyLobbyRequest) GetLobbyId() string {
	if x != nil {
		return x.LobbyId
	}
	return ""
}

type CancelLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelLobbyRequest) Reset() {
	*x = CancelLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelLobbyRequest) ProtoMessage() {}

func (x *CancelLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelLobbyRequest.ProtoReflect.Descriptor instead.
func (*CancelLobbyRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{11}
}

type SendStampRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StampId string `protobuf:"bytes,1,opt,name=stamp_id,json=stampId,proto3" json:"stamp_id,omitempty"`
}

func (x *SendStampRequest) Reset() {
	*x = SendStampRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_request_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendStampRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStampRequest) ProtoMessage() {}

func (x *SendStampRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_request_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStampRequest.ProtoReflect.Descriptor instead.
func (*SendStampRequest) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_request_proto_rawDescGZIP(), []int{12}
}

func (x *SendStampRequest) GetStampId() string {
	if x != nil {
		return x.StampId
	}
	return ""
}

var File_protomulti_multi_request_proto protoreflect.FileDescriptor

var file_protomulti_multi_request_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x59, 0x0a,
	0x11, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xd9, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x27,
	0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x28, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x62,
	0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x62,
	0x62, 0x79, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x6e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xad, 0x01, 0x0a,
	0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x6f, 0x62, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x6f, 0x62, 0x62, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x62, 0x62, 0x79, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x64, 0x2a, 0xfd, 0x01, 0x0a,
	0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x64, 0x12,
	0x10, 0x0a, 0x0b, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0xc8,
	0x01, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4c, 0x4f, 0x42, 0x42,
	0x59, 0x10, 0xac, 0x02, 0x12, 0x11, 0x0a, 0x0c, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x4c,
	0x4f, 0x42, 0x42, 0x59, 0x10, 0x90, 0x03, 0x12, 0x0f, 0x0a, 0x0a, 0x4a, 0x4f, 0x49, 0x4e, 0x5f,
	0x4c, 0x4f, 0x42, 0x42, 0x59, 0x10, 0xf4, 0x03, 0x12, 0x10, 0x0a, 0x0b, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x4c, 0x4f, 0x42, 0x42, 0x59, 0x10, 0xd8, 0x04, 0x12, 0x10, 0x0a, 0x0b, 0x57, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x4c, 0x4f, 0x42, 0x42, 0x59, 0x10, 0xbc, 0x05, 0x12, 0x12, 0x0a, 0x0d,
	0x44, 0x49, 0x53, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x4c, 0x4f, 0x42, 0x42, 0x59, 0x10, 0xa0, 0x06,
	0x12, 0x11, 0x0a, 0x0c, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x4c, 0x4f, 0x42, 0x42, 0x59,
	0x10, 0x84, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x4c,
	0x4f, 0x42, 0x42, 0x59, 0x10, 0xe8, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x53, 0x45, 0x41, 0x52, 0x43,
	0x48, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0xcc, 0x08, 0x12, 0x0f, 0x0a, 0x0a, 0x53,
	0x45, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0xb0, 0x09, 0x42, 0x4d, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a,
	0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0xaa, 0x02, 0x12, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protomulti_multi_request_proto_rawDescOnce sync.Once
	file_protomulti_multi_request_proto_rawDescData = file_protomulti_multi_request_proto_rawDesc
)

func file_protomulti_multi_request_proto_rawDescGZIP() []byte {
	file_protomulti_multi_request_proto_rawDescOnce.Do(func() {
		file_protomulti_multi_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_protomulti_multi_request_proto_rawDescData)
	})
	return file_protomulti_multi_request_proto_rawDescData
}

var file_protomulti_multi_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protomulti_multi_request_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protomulti_multi_request_proto_goTypes = []interface{}{
	(MultiRequestType)(0),                   // 0: multi.MultiRequestType
	(*CreateSessionRequest)(nil),            // 1: multi.CreateSessionRequest
	(*EndSessionRequest)(nil),               // 2: multi.EndSessionRequest
	(*SearchLobbyRequest)(nil),              // 3: multi.SearchLobbyRequest
	(*WatchLobbyRequest)(nil),               // 4: multi.WatchLobbyRequest
	(*UnwatchLobbyRequest)(nil),             // 5: multi.UnwatchLobbyRequest
	(*CreateLobbyRequest)(nil),              // 6: multi.CreateLobbyRequest
	(*JoinLobbyRequest)(nil),                // 7: multi.JoinLobbyRequest
	(*LeaveLobbyRequest)(nil),               // 8: multi.LeaveLobbyRequest
	(*StartLobbyRequest)(nil),               // 9: multi.StartLobbyRequest
	(*ReadyLobbyRequest)(nil),               // 10: multi.ReadyLobbyRequest
	(*UnreadyLobbyRequest)(nil),             // 11: multi.UnreadyLobbyRequest
	(*CancelLobbyRequest)(nil),              // 12: multi.CancelLobbyRequest
	(*SendStampRequest)(nil),                // 13: multi.SendStampRequest
	(*ProtoLobbyPlayerSlotRestriction)(nil), // 14: multi.ProtoLobbyPlayerSlotRestriction
}
var file_protomulti_multi_request_proto_depIdxs = []int32{
	14, // 0: multi.CreateLobbyRequest.restrictions:type_name -> multi.ProtoLobbyPlayerSlotRestriction
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_protomulti_multi_request_proto_init() }
func file_protomulti_multi_request_proto_init() {
	if File_protomulti_multi_request_proto != nil {
		return
	}
	file_protomulti_multi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protomulti_multi_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnwatchLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnreadyLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelLobbyRequest); i {
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
		file_protomulti_multi_request_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendStampRequest); i {
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
			RawDescriptor: file_protomulti_multi_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protomulti_multi_request_proto_goTypes,
		DependencyIndexes: file_protomulti_multi_request_proto_depIdxs,
		EnumInfos:         file_protomulti_multi_request_proto_enumTypes,
		MessageInfos:      file_protomulti_multi_request_proto_msgTypes,
	}.Build()
	File_protomulti_multi_request_proto = out.File
	file_protomulti_multi_request_proto_rawDesc = nil
	file_protomulti_multi_request_proto_goTypes = nil
	file_protomulti_multi_request_proto_depIdxs = nil
}
