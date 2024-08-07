// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protomulti/multi.response.proto

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

type GetGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameData     *ProtoGameInstance      `protobuf:"bytes,1,opt,name=game_data,json=gameData,proto3" json:"game_data,omitempty"`
	Participants []*ProtoGameParticipant `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *GetGameResponse) Reset() {
	*x = GetGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameResponse) ProtoMessage() {}

func (x *GetGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameResponse.ProtoReflect.Descriptor instead.
func (*GetGameResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetGameResponse) GetGameData() *ProtoGameInstance {
	if x != nil {
		return x.GameData
	}
	return nil
}

func (x *GetGameResponse) GetParticipants() []*ProtoGameParticipant {
	if x != nil {
		return x.Participants
	}
	return nil
}

type GameReadyPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameReadyPlayerResponse) Reset() {
	*x = GameReadyPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameReadyPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameReadyPlayerResponse) ProtoMessage() {}

func (x *GameReadyPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameReadyPlayerResponse.ProtoReflect.Descriptor instead.
func (*GameReadyPlayerResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{1}
}

type GameEnqueueActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameEnqueueActionResponse) Reset() {
	*x = GameEnqueueActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameEnqueueActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEnqueueActionResponse) ProtoMessage() {}

func (x *GameEnqueueActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEnqueueActionResponse.ProtoReflect.Descriptor instead.
func (*GameEnqueueActionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{2}
}

type GameDequeueActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameDequeueActionResponse) Reset() {
	*x = GameDequeueActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameDequeueActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameDequeueActionResponse) ProtoMessage() {}

func (x *GameDequeueActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameDequeueActionResponse.ProtoReflect.Descriptor instead.
func (*GameDequeueActionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{3}
}

type GameLockActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameLockActionResponse) Reset() {
	*x = GameLockActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameLockActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameLockActionResponse) ProtoMessage() {}

func (x *GameLockActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameLockActionResponse.ProtoReflect.Descriptor instead.
func (*GameLockActionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{4}
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{5}
}

type EndSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndSessionResponse) Reset() {
	*x = EndSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionResponse) ProtoMessage() {}

func (x *EndSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionResponse.ProtoReflect.Descriptor instead.
func (*EndSessionResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{6}
}

type SearchLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lobbies []*ProtoLobbySummary `protobuf:"bytes,1,rep,name=lobbies,proto3" json:"lobbies,omitempty"`
}

func (x *SearchLobbyResponse) Reset() {
	*x = SearchLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLobbyResponse) ProtoMessage() {}

func (x *SearchLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLobbyResponse.ProtoReflect.Descriptor instead.
func (*SearchLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{7}
}

func (x *SearchLobbyResponse) GetLobbies() []*ProtoLobbySummary {
	if x != nil {
		return x.Lobbies
	}
	return nil
}

type WatchLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchLobbyResponse) Reset() {
	*x = WatchLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchLobbyResponse) ProtoMessage() {}

func (x *WatchLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchLobbyResponse.ProtoReflect.Descriptor instead.
func (*WatchLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{8}
}

type UnwatchLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnwatchLobbyResponse) Reset() {
	*x = UnwatchLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnwatchLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnwatchLobbyResponse) ProtoMessage() {}

func (x *UnwatchLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnwatchLobbyResponse.ProtoReflect.Descriptor instead.
func (*UnwatchLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{9}
}

type CreateLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LobbyId string `protobuf:"bytes,1,opt,name=lobby_id,json=lobbyId,proto3" json:"lobby_id,omitempty"`
	PartyId string `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *CreateLobbyResponse) Reset() {
	*x = CreateLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLobbyResponse) ProtoMessage() {}

func (x *CreateLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLobbyResponse.ProtoReflect.Descriptor instead.
func (*CreateLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{10}
}

func (x *CreateLobbyResponse) GetLobbyId() string {
	if x != nil {
		return x.LobbyId
	}
	return ""
}

func (x *CreateLobbyResponse) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type JoinLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinLobbyResponse) Reset() {
	*x = JoinLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLobbyResponse) ProtoMessage() {}

func (x *JoinLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLobbyResponse.ProtoReflect.Descriptor instead.
func (*JoinLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{11}
}

type LeaveLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveLobbyResponse) Reset() {
	*x = LeaveLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveLobbyResponse) ProtoMessage() {}

func (x *LeaveLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveLobbyResponse.ProtoReflect.Descriptor instead.
func (*LeaveLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{12}
}

type ReadyLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadyLobbyResponse) Reset() {
	*x = ReadyLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyLobbyResponse) ProtoMessage() {}

func (x *ReadyLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyLobbyResponse.ProtoReflect.Descriptor instead.
func (*ReadyLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{13}
}

type UnreadyLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnreadyLobbyResponse) Reset() {
	*x = UnreadyLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnreadyLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnreadyLobbyResponse) ProtoMessage() {}

func (x *UnreadyLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnreadyLobbyResponse.ProtoReflect.Descriptor instead.
func (*UnreadyLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{14}
}

type CancelLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelLobbyResponse) Reset() {
	*x = CancelLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelLobbyResponse) ProtoMessage() {}

func (x *CancelLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelLobbyResponse.ProtoReflect.Descriptor instead.
func (*CancelLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{15}
}

type SendStampResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendStampResponse) Reset() {
	*x = SendStampResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendStampResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStampResponse) ProtoMessage() {}

func (x *SendStampResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStampResponse.ProtoReflect.Descriptor instead.
func (*SendStampResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{16}
}

type StartLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartLobbyResponse) Reset() {
	*x = StartLobbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protomulti_multi_response_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartLobbyResponse) ProtoMessage() {}

func (x *StartLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protomulti_multi_response_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartLobbyResponse.ProtoReflect.Descriptor instead.
func (*StartLobbyResponse) Descriptor() ([]byte, []int) {
	return file_protomulti_multi_response_proto_rawDescGZIP(), []int{17}
}

var File_protomulti_multi_response_proto protoreflect.FileDescriptor

var file_protomulti_multi_response_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0c, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x47,
	0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x17,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x61, 0x6d, 0x65, 0x45,
	0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x65, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x07, 0x6c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x6f,
	0x62, 0x62, 0x69, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x55,
	0x6e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f,
	0x62, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f,
	0x62, 0x62, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64,
	0x22, 0x13, 0x0a, 0x11, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4d, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61,
	0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0xaa, 0x02, 0x12, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protomulti_multi_response_proto_rawDescOnce sync.Once
	file_protomulti_multi_response_proto_rawDescData = file_protomulti_multi_response_proto_rawDesc
)

func file_protomulti_multi_response_proto_rawDescGZIP() []byte {
	file_protomulti_multi_response_proto_rawDescOnce.Do(func() {
		file_protomulti_multi_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_protomulti_multi_response_proto_rawDescData)
	})
	return file_protomulti_multi_response_proto_rawDescData
}

var file_protomulti_multi_response_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_protomulti_multi_response_proto_goTypes = []interface{}{
	(*GetGameResponse)(nil),           // 0: multi.GetGameResponse
	(*GameReadyPlayerResponse)(nil),   // 1: multi.GameReadyPlayerResponse
	(*GameEnqueueActionResponse)(nil), // 2: multi.GameEnqueueActionResponse
	(*GameDequeueActionResponse)(nil), // 3: multi.GameDequeueActionResponse
	(*GameLockActionResponse)(nil),    // 4: multi.GameLockActionResponse
	(*CreateSessionResponse)(nil),     // 5: multi.CreateSessionResponse
	(*EndSessionResponse)(nil),        // 6: multi.EndSessionResponse
	(*SearchLobbyResponse)(nil),       // 7: multi.SearchLobbyResponse
	(*WatchLobbyResponse)(nil),        // 8: multi.WatchLobbyResponse
	(*UnwatchLobbyResponse)(nil),      // 9: multi.UnwatchLobbyResponse
	(*CreateLobbyResponse)(nil),       // 10: multi.CreateLobbyResponse
	(*JoinLobbyResponse)(nil),         // 11: multi.JoinLobbyResponse
	(*LeaveLobbyResponse)(nil),        // 12: multi.LeaveLobbyResponse
	(*ReadyLobbyResponse)(nil),        // 13: multi.ReadyLobbyResponse
	(*UnreadyLobbyResponse)(nil),      // 14: multi.UnreadyLobbyResponse
	(*CancelLobbyResponse)(nil),       // 15: multi.CancelLobbyResponse
	(*SendStampResponse)(nil),         // 16: multi.SendStampResponse
	(*StartLobbyResponse)(nil),        // 17: multi.StartLobbyResponse
	(*ProtoGameInstance)(nil),         // 18: multi.ProtoGameInstance
	(*ProtoGameParticipant)(nil),      // 19: multi.ProtoGameParticipant
	(*ProtoLobbySummary)(nil),         // 20: multi.ProtoLobbySummary
}
var file_protomulti_multi_response_proto_depIdxs = []int32{
	18, // 0: multi.GetGameResponse.game_data:type_name -> multi.ProtoGameInstance
	19, // 1: multi.GetGameResponse.participants:type_name -> multi.ProtoGameParticipant
	20, // 2: multi.SearchLobbyResponse.lobbies:type_name -> multi.ProtoLobbySummary
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protomulti_multi_response_proto_init() }
func file_protomulti_multi_response_proto_init() {
	if File_protomulti_multi_response_proto != nil {
		return
	}
	file_protomulti_multi_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protomulti_multi_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameReadyPlayerResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameEnqueueActionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameDequeueActionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameLockActionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnwatchLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnreadyLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelLobbyResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendStampResponse); i {
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
		file_protomulti_multi_response_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartLobbyResponse); i {
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
			RawDescriptor: file_protomulti_multi_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protomulti_multi_response_proto_goTypes,
		DependencyIndexes: file_protomulti_multi_response_proto_depIdxs,
		MessageInfos:      file_protomulti_multi_response_proto_msgTypes,
	}.Build()
	File_protomulti_multi_response_proto = out.File
	file_protomulti_multi_response_proto_rawDesc = nil
	file_protomulti_multi_response_proto_goTypes = nil
	file_protomulti_multi_response_proto_depIdxs = nil
}
