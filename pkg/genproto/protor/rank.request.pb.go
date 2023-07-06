// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: rank.request.proto

package protor

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

type SubmitScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId    string                     `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName  string                     `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel int32                      `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	Score       uint64                     `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	RegionId    string                     `protobuf:"bytes,5,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	NodeIndex   int32                      `protobuf:"varint,6,opt,name=node_index,json=nodeIndex,proto3" json:"node_index,omitempty"`
	Loadout     *protog.ProtoPlayerLoadout `protobuf:"bytes,7,opt,name=loadout,proto3" json:"loadout,omitempty"`
}

func (x *SubmitScoreRequest) Reset() {
	*x = SubmitScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitScoreRequest) ProtoMessage() {}

func (x *SubmitScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rank_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitScoreRequest.ProtoReflect.Descriptor instead.
func (*SubmitScoreRequest) Descriptor() ([]byte, []int) {
	return file_rank_request_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitScoreRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *SubmitScoreRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *SubmitScoreRequest) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *SubmitScoreRequest) GetScore() uint64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SubmitScoreRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *SubmitScoreRequest) GetNodeIndex() int32 {
	if x != nil {
		return x.NodeIndex
	}
	return 0
}

func (x *SubmitScoreRequest) GetLoadout() *protog.ProtoPlayerLoadout {
	if x != nil {
		return x.Loadout
	}
	return nil
}

type RefreshWeeklyRankRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RegionId string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *RefreshWeeklyRankRangeRequest) Reset() {
	*x = RefreshWeeklyRankRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshWeeklyRankRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshWeeklyRankRangeRequest) ProtoMessage() {}

func (x *RefreshWeeklyRankRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rank_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshWeeklyRankRangeRequest.ProtoReflect.Descriptor instead.
func (*RefreshWeeklyRankRangeRequest) Descriptor() ([]byte, []int) {
	return file_rank_request_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshWeeklyRankRangeRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *RefreshWeeklyRankRangeRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type RefreshSpecialRankRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RegionId string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *RefreshSpecialRankRangeRequest) Reset() {
	*x = RefreshSpecialRankRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshSpecialRankRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshSpecialRankRangeRequest) ProtoMessage() {}

func (x *RefreshSpecialRankRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rank_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshSpecialRankRangeRequest.ProtoReflect.Descriptor instead.
func (*RefreshSpecialRankRangeRequest) Descriptor() ([]byte, []int) {
	return file_rank_request_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshSpecialRankRangeRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *RefreshSpecialRankRangeRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type FetchPlayerRankingInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *FetchPlayerRankingInfoRequest) Reset() {
	*x = FetchPlayerRankingInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPlayerRankingInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPlayerRankingInfoRequest) ProtoMessage() {}

func (x *FetchPlayerRankingInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rank_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPlayerRankingInfoRequest.ProtoReflect.Descriptor instead.
func (*FetchPlayerRankingInfoRequest) Descriptor() ([]byte, []int) {
	return file_rank_request_proto_rawDescGZIP(), []int{3}
}

func (x *FetchPlayerRankingInfoRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type FetchPlayerRankingInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active []*ProtoRankingInfo       `protobuf:"bytes,1,rep,name=active,proto3" json:"active,omitempty"`
	Claims []*ProtoRankingEventClaim `protobuf:"bytes,2,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (x *FetchPlayerRankingInfoResponse) Reset() {
	*x = FetchPlayerRankingInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPlayerRankingInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPlayerRankingInfoResponse) ProtoMessage() {}

func (x *FetchPlayerRankingInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rank_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPlayerRankingInfoResponse.ProtoReflect.Descriptor instead.
func (*FetchPlayerRankingInfoResponse) Descriptor() ([]byte, []int) {
	return file_rank_request_proto_rawDescGZIP(), []int{4}
}

func (x *FetchPlayerRankingInfoResponse) GetActive() []*ProtoRankingInfo {
	if x != nil {
		return x.Active
	}
	return nil
}

func (x *FetchPlayerRankingInfoResponse) GetClaims() []*ProtoRankingEventClaim {
	if x != nil {
		return x.Claims
	}
	return nil
}

var File_rank_request_proto protoreflect.FileDescriptor

var file_rank_request_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x34, 0x0a, 0x07,
	0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x6f,
	0x75, 0x74, 0x22, 0x59, 0x0a, 0x1d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x57, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5a, 0x0a,
	0x1e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x1d, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x1e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x42, 0x43, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69,
	0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0xaa, 0x02, 0x0c, 0x4d, 0x6f, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rank_request_proto_rawDescOnce sync.Once
	file_rank_request_proto_rawDescData = file_rank_request_proto_rawDesc
)

func file_rank_request_proto_rawDescGZIP() []byte {
	file_rank_request_proto_rawDescOnce.Do(func() {
		file_rank_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_rank_request_proto_rawDescData)
	})
	return file_rank_request_proto_rawDescData
}

var file_rank_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rank_request_proto_goTypes = []interface{}{
	(*SubmitScoreRequest)(nil),             // 0: rank.SubmitScoreRequest
	(*RefreshWeeklyRankRangeRequest)(nil),  // 1: rank.RefreshWeeklyRankRangeRequest
	(*RefreshSpecialRankRangeRequest)(nil), // 2: rank.RefreshSpecialRankRangeRequest
	(*FetchPlayerRankingInfoRequest)(nil),  // 3: rank.FetchPlayerRankingInfoRequest
	(*FetchPlayerRankingInfoResponse)(nil), // 4: rank.FetchPlayerRankingInfoResponse
	(*protog.ProtoPlayerLoadout)(nil),      // 5: protog.ProtoPlayerLoadout
	(*ProtoRankingInfo)(nil),               // 6: rank.ProtoRankingInfo
	(*ProtoRankingEventClaim)(nil),         // 7: rank.ProtoRankingEventClaim
}
var file_rank_request_proto_depIdxs = []int32{
	5, // 0: rank.SubmitScoreRequest.loadout:type_name -> protog.ProtoPlayerLoadout
	6, // 1: rank.FetchPlayerRankingInfoResponse.active:type_name -> rank.ProtoRankingInfo
	7, // 2: rank.FetchPlayerRankingInfoResponse.claims:type_name -> rank.ProtoRankingEventClaim
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rank_request_proto_init() }
func file_rank_request_proto_init() {
	if File_rank_request_proto != nil {
		return
	}
	file_rank_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rank_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitScoreRequest); i {
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
		file_rank_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshWeeklyRankRangeRequest); i {
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
		file_rank_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshSpecialRankRangeRequest); i {
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
		file_rank_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPlayerRankingInfoRequest); i {
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
		file_rank_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPlayerRankingInfoResponse); i {
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
			RawDescriptor: file_rank_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rank_request_proto_goTypes,
		DependencyIndexes: file_rank_request_proto_depIdxs,
		MessageInfos:      file_rank_request_proto_msgTypes,
	}.Build()
	File_rank_request_proto = out.File
	file_rank_request_proto_rawDesc = nil
	file_rank_request_proto_goTypes = nil
	file_rank_request_proto_depIdxs = nil
}
