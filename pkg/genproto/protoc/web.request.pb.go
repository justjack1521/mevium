// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: web.request.proto

package protoc

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

type RequestType int32

const (
	RequestType_BASE                  RequestType = 0
	RequestType_GET_PROFILE           RequestType = 200
	RequestType_CREATE_PROFILE        RequestType = 300
	RequestType_CARD_SALE             RequestType = 400
	RequestType_CARD_FILTER_SORT      RequestType = 500
	RequestType_CARD_FAVOURITE        RequestType = 600
	RequestType_SKILL_PANEL           RequestType = 700
	RequestType_DECK_EDIT             RequestType = 800
	RequestType_DECK_EDIT_ALL         RequestType = 900
	RequestType_TELEPORT              RequestType = 1000
	RequestType_REGION_MAP_OPEN       RequestType = 1100
	RequestType_BATTLE_REVIVE         RequestType = 1200
	RequestType_BATTLE_COMPLETE       RequestType = 1300
	RequestType_FOLLOW_PLAYER         RequestType = 1400
	RequestType_UNFOLLOW_PLAYER       RequestType = 1500
	RequestType_CONFIRM_DAILY_MISSION RequestType = 1600
	RequestType_CLAIM_EVENT_RANKING   RequestType = 1700
	RequestType_CLAIM_MAILBOX         RequestType = 1800
	RequestType_BATTLE_START          RequestType = 1900
	RequestType_CARD_TRANSFER         RequestType = 2000
	RequestType_CARD_FUSION           RequestType = 2100
	RequestType_CARD_FUSION_BOOST     RequestType = 2200
	RequestType_STAMINA_RESTORE       RequestType = 2300
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0:    "BASE",
		200:  "GET_PROFILE",
		300:  "CREATE_PROFILE",
		400:  "CARD_SALE",
		500:  "CARD_FILTER_SORT",
		600:  "CARD_FAVOURITE",
		700:  "SKILL_PANEL",
		800:  "DECK_EDIT",
		900:  "DECK_EDIT_ALL",
		1000: "TELEPORT",
		1100: "REGION_MAP_OPEN",
		1200: "BATTLE_REVIVE",
		1300: "BATTLE_COMPLETE",
		1400: "FOLLOW_PLAYER",
		1500: "UNFOLLOW_PLAYER",
		1600: "CONFIRM_DAILY_MISSION",
		1700: "CLAIM_EVENT_RANKING",
		1800: "CLAIM_MAILBOX",
		1900: "BATTLE_START",
		2000: "CARD_TRANSFER",
		2100: "CARD_FUSION",
		2200: "CARD_FUSION_BOOST",
		2300: "STAMINA_RESTORE",
	}
	RequestType_value = map[string]int32{
		"BASE":                  0,
		"GET_PROFILE":           200,
		"CREATE_PROFILE":        300,
		"CARD_SALE":             400,
		"CARD_FILTER_SORT":      500,
		"CARD_FAVOURITE":        600,
		"SKILL_PANEL":           700,
		"DECK_EDIT":             800,
		"DECK_EDIT_ALL":         900,
		"TELEPORT":              1000,
		"REGION_MAP_OPEN":       1100,
		"BATTLE_REVIVE":         1200,
		"BATTLE_COMPLETE":       1300,
		"FOLLOW_PLAYER":         1400,
		"UNFOLLOW_PLAYER":       1500,
		"CONFIRM_DAILY_MISSION": 1600,
		"CLAIM_EVENT_RANKING":   1700,
		"CLAIM_MAILBOX":         1800,
		"BATTLE_START":          1900,
		"CARD_TRANSFER":         2000,
		"CARD_FUSION":           2100,
		"CARD_FUSION_BOOST":     2200,
		"STAMINA_RESTORE":       2300,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_web_request_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_web_request_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_web_request_proto_rawDescGZIP(), []int{0}
}

type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	CommandId    string `protobuf:"bytes,3,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_web_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_web_request_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeader) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RequestHeader) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *RequestHeader) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

type BaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operation int32          `protobuf:"varint,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Data      []byte         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BaseRequest) Reset() {
	*x = BaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseRequest) ProtoMessage() {}

func (x *BaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseRequest.ProtoReflect.Descriptor instead.
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return file_web_request_proto_rawDescGZIP(), []int{1}
}

func (x *BaseRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BaseRequest) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *BaseRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_web_request_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProfileRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type CardFavouriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card  string `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CardFavouriteRequest) Reset() {
	*x = CardFavouriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardFavouriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardFavouriteRequest) ProtoMessage() {}

func (x *CardFavouriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardFavouriteRequest.ProtoReflect.Descriptor instead.
func (*CardFavouriteRequest) Descriptor() ([]byte, []int) {
	return file_web_request_proto_rawDescGZIP(), []int{3}
}

func (x *CardFavouriteRequest) GetCard() string {
	if x != nil {
		return x.Card
	}
	return ""
}

func (x *CardFavouriteRequest) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_web_request_proto protoreflect.FileDescriptor

var file_web_request_proto_rawDesc = []byte{
	0x0a, 0x11, 0x77, 0x65, 0x62, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x22, 0x70, 0x0a, 0x0d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x6e, 0x0a,
	0x0b, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x61, 0x72, 0x64, 0x46, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0xda, 0x03, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x53, 0x45,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0b, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c,
	0x45, 0x10, 0xc8, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0xac, 0x02, 0x12, 0x0e, 0x0a, 0x09, 0x43, 0x41, 0x52,
	0x44, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x10, 0x90, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x41, 0x52,
	0x44, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x4f, 0x52, 0x54, 0x10, 0xf4, 0x03,
	0x12, 0x13, 0x0a, 0x0e, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x46, 0x41, 0x56, 0x4f, 0x55, 0x52, 0x49,
	0x54, 0x45, 0x10, 0xd8, 0x04, 0x12, 0x10, 0x0a, 0x0b, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x50,
	0x41, 0x4e, 0x45, 0x4c, 0x10, 0xbc, 0x05, 0x12, 0x0e, 0x0a, 0x09, 0x44, 0x45, 0x43, 0x4b, 0x5f,
	0x45, 0x44, 0x49, 0x54, 0x10, 0xa0, 0x06, 0x12, 0x12, 0x0a, 0x0d, 0x44, 0x45, 0x43, 0x4b, 0x5f,
	0x45, 0x44, 0x49, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x84, 0x07, 0x12, 0x0d, 0x0a, 0x08, 0x54,
	0x45, 0x4c, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0xe8, 0x07, 0x12, 0x14, 0x0a, 0x0f, 0x52, 0x45,
	0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0xcc, 0x08,
	0x12, 0x12, 0x0a, 0x0d, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x56,
	0x45, 0x10, 0xb0, 0x09, 0x12, 0x14, 0x0a, 0x0f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x94, 0x0a, 0x12, 0x12, 0x0a, 0x0d, 0x46, 0x4f,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0xf8, 0x0a, 0x12, 0x14,
	0x0a, 0x0f, 0x55, 0x4e, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x10, 0xdc, 0x0b, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x5f,
	0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0xc0, 0x0c,
	0x12, 0x18, 0x0a, 0x13, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x52, 0x41, 0x4e, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0xa4, 0x0d, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x4c,
	0x41, 0x49, 0x4d, 0x5f, 0x4d, 0x41, 0x49, 0x4c, 0x42, 0x4f, 0x58, 0x10, 0x88, 0x0e, 0x12, 0x11,
	0x0a, 0x0c, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0xec,
	0x0e, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46,
	0x45, 0x52, 0x10, 0xd0, 0x0f, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x46, 0x55,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0xb4, 0x10, 0x12, 0x16, 0x0a, 0x11, 0x43, 0x41, 0x52, 0x44, 0x5f,
	0x46, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x4f, 0x4f, 0x53, 0x54, 0x10, 0x98, 0x11, 0x12,
	0x14, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x4d, 0x49, 0x4e, 0x41, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x4f,
	0x52, 0x45, 0x10, 0xfc, 0x11, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31,
	0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_web_request_proto_rawDescOnce sync.Once
	file_web_request_proto_rawDescData = file_web_request_proto_rawDesc
)

func file_web_request_proto_rawDescGZIP() []byte {
	file_web_request_proto_rawDescOnce.Do(func() {
		file_web_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_request_proto_rawDescData)
	})
	return file_web_request_proto_rawDescData
}

var file_web_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_web_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_request_proto_goTypes = []interface{}{
	(RequestType)(0),             // 0: protoc.RequestType
	(*RequestHeader)(nil),        // 1: protoc.RequestHeader
	(*BaseRequest)(nil),          // 2: protoc.BaseRequest
	(*CreateProfileRequest)(nil), // 3: protoc.CreateProfileRequest
	(*CardFavouriteRequest)(nil), // 4: protoc.CardFavouriteRequest
}
var file_web_request_proto_depIdxs = []int32{
	1, // 0: protoc.BaseRequest.header:type_name -> protoc.RequestHeader
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_request_proto_init() }
func file_web_request_proto_init() {
	if File_web_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_web_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseRequest); i {
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
		file_web_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileRequest); i {
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
		file_web_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardFavouriteRequest); i {
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
			RawDescriptor: file_web_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_request_proto_goTypes,
		DependencyIndexes: file_web_request_proto_depIdxs,
		EnumInfos:         file_web_request_proto_enumTypes,
		MessageInfos:      file_web_request_proto_msgTypes,
	}.Build()
	File_web_request_proto = out.File
	file_web_request_proto_rawDesc = nil
	file_web_request_proto_goTypes = nil
	file_web_request_proto_depIdxs = nil
}
