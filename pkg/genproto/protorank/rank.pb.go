// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protorank/rank.proto

package protorank

import (
	protoidentity "github.com/justjack1521/mevium/pkg/genproto/protoidentity"
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

type ProtoRankingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShortCode       string                       `protobuf:"bytes,2,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	PlayerRankings  []*ProtoPlayerRankSetDetails `protobuf:"bytes,3,rep,name=player_rankings,json=playerRankings,proto3" json:"player_rankings,omitempty"`
	TopRankings     []*ProtoPlayerRankSetDetails `protobuf:"bytes,6,rep,name=top_rankings,json=topRankings,proto3" json:"top_rankings,omitempty"`
	RankRangeScores map[int64]float64            `protobuf:"bytes,7,rep,name=rank_range_scores,json=rankRangeScores,proto3" json:"rank_range_scores,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *ProtoRankingInfo) Reset() {
	*x = ProtoRankingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingInfo) ProtoMessage() {}

func (x *ProtoRankingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingInfo.ProtoReflect.Descriptor instead.
func (*ProtoRankingInfo) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoRankingInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProtoRankingInfo) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

func (x *ProtoRankingInfo) GetPlayerRankings() []*ProtoPlayerRankSetDetails {
	if x != nil {
		return x.PlayerRankings
	}
	return nil
}

func (x *ProtoRankingInfo) GetTopRankings() []*ProtoPlayerRankSetDetails {
	if x != nil {
		return x.TopRankings
	}
	return nil
}

func (x *ProtoRankingInfo) GetRankRangeScores() map[int64]float64 {
	if x != nil {
		return x.RankRangeScores
	}
	return nil
}

type ProtoPlayerRankSetDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId       string                                    `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerName     string                                    `protobuf:"bytes,2,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	PlayerLevel    int32                                     `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	Score          float64                                   `protobuf:"fixed64,4,opt,name=score,proto3" json:"score,omitempty"`
	Rank           int64                                     `protobuf:"varint,5,opt,name=rank,proto3" json:"rank,omitempty"`
	Identity       *protoidentity.ProtoPlayerIdentity        `protobuf:"bytes,6,opt,name=identity,proto3" json:"identity,omitempty"`
	PrimaryLoadout *protoidentity.ProtoPlayerLoadoutIdentity `protobuf:"bytes,7,opt,name=primary_loadout,json=primaryLoadout,proto3" json:"primary_loadout,omitempty"`
	IsPlayer       bool                                      `protobuf:"varint,8,opt,name=is_player,json=isPlayer,proto3" json:"is_player,omitempty"`
}

func (x *ProtoPlayerRankSetDetails) Reset() {
	*x = ProtoPlayerRankSetDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoPlayerRankSetDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPlayerRankSetDetails) ProtoMessage() {}

func (x *ProtoPlayerRankSetDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPlayerRankSetDetails.ProtoReflect.Descriptor instead.
func (*ProtoPlayerRankSetDetails) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoPlayerRankSetDetails) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *ProtoPlayerRankSetDetails) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *ProtoPlayerRankSetDetails) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *ProtoPlayerRankSetDetails) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ProtoPlayerRankSetDetails) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *ProtoPlayerRankSetDetails) GetIdentity() *protoidentity.ProtoPlayerIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *ProtoPlayerRankSetDetails) GetPrimaryLoadout() *protoidentity.ProtoPlayerLoadoutIdentity {
	if x != nil {
		return x.PrimaryLoadout
	}
	return nil
}

func (x *ProtoPlayerRankSetDetails) GetIsPlayer() bool {
	if x != nil {
		return x.IsPlayer
	}
	return false
}

type ProtoRankingEventClaim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShortCode     string                       `protobuf:"bytes,2,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	Claimed       bool                         `protobuf:"varint,3,opt,name=claimed,proto3" json:"claimed,omitempty"`
	CanClaim      bool                         `protobuf:"varint,4,opt,name=can_claim,json=canClaim,proto3" json:"can_claim,omitempty"`
	PlayerRanking []*ProtoPlayerRankSetDetails `protobuf:"bytes,5,rep,name=player_ranking,json=playerRanking,proto3" json:"player_ranking,omitempty"`
	TopRankings   []*ProtoPlayerRankSetDetails `protobuf:"bytes,6,rep,name=top_rankings,json=topRankings,proto3" json:"top_rankings,omitempty"`
	Statistics    *ProtoRankingEventStats      `protobuf:"bytes,7,opt,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *ProtoRankingEventClaim) Reset() {
	*x = ProtoRankingEventClaim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingEventClaim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingEventClaim) ProtoMessage() {}

func (x *ProtoRankingEventClaim) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingEventClaim.ProtoReflect.Descriptor instead.
func (*ProtoRankingEventClaim) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoRankingEventClaim) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProtoRankingEventClaim) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

func (x *ProtoRankingEventClaim) GetClaimed() bool {
	if x != nil {
		return x.Claimed
	}
	return false
}

func (x *ProtoRankingEventClaim) GetCanClaim() bool {
	if x != nil {
		return x.CanClaim
	}
	return false
}

func (x *ProtoRankingEventClaim) GetPlayerRanking() []*ProtoPlayerRankSetDetails {
	if x != nil {
		return x.PlayerRanking
	}
	return nil
}

func (x *ProtoRankingEventClaim) GetTopRankings() []*ProtoPlayerRankSetDetails {
	if x != nil {
		return x.TopRankings
	}
	return nil
}

func (x *ProtoRankingEventClaim) GetStatistics() *ProtoRankingEventStats {
	if x != nil {
		return x.Statistics
	}
	return nil
}

type ProtoRankingEventReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reward:
	//	*ProtoRankingEventReward_RewardItem
	//	*ProtoRankingEventReward_RewardCard
	Reward   isProtoRankingEventReward_Reward `protobuf_oneof:"reward"`
	Quantity int32                            `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ProtoRankingEventReward) Reset() {
	*x = ProtoRankingEventReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingEventReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingEventReward) ProtoMessage() {}

func (x *ProtoRankingEventReward) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingEventReward.ProtoReflect.Descriptor instead.
func (*ProtoRankingEventReward) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{3}
}

func (m *ProtoRankingEventReward) GetReward() isProtoRankingEventReward_Reward {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (x *ProtoRankingEventReward) GetRewardItem() string {
	if x, ok := x.GetReward().(*ProtoRankingEventReward_RewardItem); ok {
		return x.RewardItem
	}
	return ""
}

func (x *ProtoRankingEventReward) GetRewardCard() string {
	if x, ok := x.GetReward().(*ProtoRankingEventReward_RewardCard); ok {
		return x.RewardCard
	}
	return ""
}

func (x *ProtoRankingEventReward) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type isProtoRankingEventReward_Reward interface {
	isProtoRankingEventReward_Reward()
}

type ProtoRankingEventReward_RewardItem struct {
	RewardItem string `protobuf:"bytes,1,opt,name=reward_item,json=rewardItem,proto3,oneof"`
}

type ProtoRankingEventReward_RewardCard struct {
	RewardCard string `protobuf:"bytes,2,opt,name=reward_card,json=rewardCard,proto3,oneof"`
}

func (*ProtoRankingEventReward_RewardItem) isProtoRankingEventReward_Reward() {}

func (*ProtoRankingEventReward_RewardCard) isProtoRankingEventReward_Reward() {}

type ProtoRankingEventStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalParticipants  int64                                    `protobuf:"varint,1,opt,name=total_participants,json=totalParticipants,proto3" json:"total_participants,omitempty"`
	PopularRegion      *ProtoRankingEventStatPopularRegion      `protobuf:"bytes,2,opt,name=popular_region,json=popularRegion,proto3" json:"popular_region,omitempty"`
	PopularJob         *ProtoRankingEventStatPopularJob         `protobuf:"bytes,3,opt,name=popular_job,json=popularJob,proto3" json:"popular_job,omitempty"`
	PopularAbilityCard *ProtoRankingEventStatPopularAbilityCard `protobuf:"bytes,4,opt,name=popular_ability_card,json=popularAbilityCard,proto3" json:"popular_ability_card,omitempty"`
}

func (x *ProtoRankingEventStats) Reset() {
	*x = ProtoRankingEventStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingEventStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingEventStats) ProtoMessage() {}

func (x *ProtoRankingEventStats) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingEventStats.ProtoReflect.Descriptor instead.
func (*ProtoRankingEventStats) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{4}
}

func (x *ProtoRankingEventStats) GetTotalParticipants() int64 {
	if x != nil {
		return x.TotalParticipants
	}
	return 0
}

func (x *ProtoRankingEventStats) GetPopularRegion() *ProtoRankingEventStatPopularRegion {
	if x != nil {
		return x.PopularRegion
	}
	return nil
}

func (x *ProtoRankingEventStats) GetPopularJob() *ProtoRankingEventStatPopularJob {
	if x != nil {
		return x.PopularJob
	}
	return nil
}

func (x *ProtoRankingEventStats) GetPopularAbilityCard() *ProtoRankingEventStatPopularAbilityCard {
	if x != nil {
		return x.PopularAbilityCard
	}
	return nil
}

type ProtoRankingEventStatPopularJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobCardId   string `protobuf:"bytes,1,opt,name=job_card_id,json=jobCardId,proto3" json:"job_card_id,omitempty"`
	SubJobIndex int32  `protobuf:"varint,2,opt,name=sub_job_index,json=subJobIndex,proto3" json:"sub_job_index,omitempty"`
}

func (x *ProtoRankingEventStatPopularJob) Reset() {
	*x = ProtoRankingEventStatPopularJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingEventStatPopularJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingEventStatPopularJob) ProtoMessage() {}

func (x *ProtoRankingEventStatPopularJob) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingEventStatPopularJob.ProtoReflect.Descriptor instead.
func (*ProtoRankingEventStatPopularJob) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{5}
}

func (x *ProtoRankingEventStatPopularJob) GetJobCardId() string {
	if x != nil {
		return x.JobCardId
	}
	return ""
}

func (x *ProtoRankingEventStatPopularJob) GetSubJobIndex() int32 {
	if x != nil {
		return x.SubJobIndex
	}
	return 0
}

type ProtoRankingEventStatPopularRegion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionMapId string `protobuf:"bytes,1,opt,name=region_map_id,json=regionMapId,proto3" json:"region_map_id,omitempty"`
	NodeIndex   int32  `protobuf:"varint,2,opt,name=node_index,json=nodeIndex,proto3" json:"node_index,omitempty"`
}

func (x *ProtoRankingEventStatPopularRegion) Reset() {
	*x = ProtoRankingEventStatPopularRegion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingEventStatPopularRegion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingEventStatPopularRegion) ProtoMessage() {}

func (x *ProtoRankingEventStatPopularRegion) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingEventStatPopularRegion.ProtoReflect.Descriptor instead.
func (*ProtoRankingEventStatPopularRegion) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{6}
}

func (x *ProtoRankingEventStatPopularRegion) GetRegionMapId() string {
	if x != nil {
		return x.RegionMapId
	}
	return ""
}

func (x *ProtoRankingEventStatPopularRegion) GetNodeIndex() int32 {
	if x != nil {
		return x.NodeIndex
	}
	return 0
}

type ProtoRankingEventStatPopularAbilityCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityCardId string `protobuf:"bytes,1,opt,name=ability_card_id,json=abilityCardId,proto3" json:"ability_card_id,omitempty"`
}

func (x *ProtoRankingEventStatPopularAbilityCard) Reset() {
	*x = ProtoRankingEventStatPopularAbilityCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorank_rank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRankingEventStatPopularAbilityCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRankingEventStatPopularAbilityCard) ProtoMessage() {}

func (x *ProtoRankingEventStatPopularAbilityCard) ProtoReflect() protoreflect.Message {
	mi := &file_protorank_rank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRankingEventStatPopularAbilityCard.ProtoReflect.Descriptor instead.
func (*ProtoRankingEventStatPopularAbilityCard) Descriptor() ([]byte, []int) {
	return file_protorank_rank_proto_rawDescGZIP(), []int{7}
}

func (x *ProtoRankingEventStatPopularAbilityCard) GetAbilityCardId() string {
	if x != nil {
		return x.AbilityCardId
	}
	return ""
}

var File_protorank_rank_proto protoreflect.FileDescriptor

var file_protorank_rank_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x72, 0x61, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x1a, 0x1a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x0f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x74,
	0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x72, 0x61,
	0x6e, 0x6b, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x72, 0x61, 0x6e, 0x6b, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x1a, 0x42, 0x0a, 0x14, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcd, 0x02, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x12, 0x39, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x0f, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x6f,
	0x75, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0xc8, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x61, 0x6e, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x63, 0x61, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x46, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x42, 0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x43, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0xc1, 0x02, 0x0a, 0x16, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0b, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x4a, 0x6f,
	0x62, 0x52, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x5f, 0x0a,
	0x14, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x61,
	0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52, 0x12, 0x70, 0x6f, 0x70, 0x75,
	0x6c, 0x61, 0x72, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x22, 0x65,
	0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x4a, 0x6f,
	0x62, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x67, 0x0a, 0x22, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x51,
	0x0a, 0x27, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x42, 0x4b, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x75, 0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76,
	0x69, 0x75, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x61, 0x6e, 0x6b, 0xaa, 0x02, 0x11, 0x4d, 0x6f, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protorank_rank_proto_rawDescOnce sync.Once
	file_protorank_rank_proto_rawDescData = file_protorank_rank_proto_rawDesc
)

func file_protorank_rank_proto_rawDescGZIP() []byte {
	file_protorank_rank_proto_rawDescOnce.Do(func() {
		file_protorank_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_protorank_rank_proto_rawDescData)
	})
	return file_protorank_rank_proto_rawDescData
}

var file_protorank_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protorank_rank_proto_goTypes = []interface{}{
	(*ProtoRankingInfo)(nil),                        // 0: rank.ProtoRankingInfo
	(*ProtoPlayerRankSetDetails)(nil),               // 1: rank.ProtoPlayerRankSetDetails
	(*ProtoRankingEventClaim)(nil),                  // 2: rank.ProtoRankingEventClaim
	(*ProtoRankingEventReward)(nil),                 // 3: rank.ProtoRankingEventReward
	(*ProtoRankingEventStats)(nil),                  // 4: rank.ProtoRankingEventStats
	(*ProtoRankingEventStatPopularJob)(nil),         // 5: rank.ProtoRankingEventStatPopularJob
	(*ProtoRankingEventStatPopularRegion)(nil),      // 6: rank.ProtoRankingEventStatPopularRegion
	(*ProtoRankingEventStatPopularAbilityCard)(nil), // 7: rank.ProtoRankingEventStatPopularAbilityCard
	nil, // 8: rank.ProtoRankingInfo.RankRangeScoresEntry
	(*protoidentity.ProtoPlayerIdentity)(nil),        // 9: identity.ProtoPlayerIdentity
	(*protoidentity.ProtoPlayerLoadoutIdentity)(nil), // 10: identity.ProtoPlayerLoadoutIdentity
}
var file_protorank_rank_proto_depIdxs = []int32{
	1,  // 0: rank.ProtoRankingInfo.player_rankings:type_name -> rank.ProtoPlayerRankSetDetails
	1,  // 1: rank.ProtoRankingInfo.top_rankings:type_name -> rank.ProtoPlayerRankSetDetails
	8,  // 2: rank.ProtoRankingInfo.rank_range_scores:type_name -> rank.ProtoRankingInfo.RankRangeScoresEntry
	9,  // 3: rank.ProtoPlayerRankSetDetails.identity:type_name -> identity.ProtoPlayerIdentity
	10, // 4: rank.ProtoPlayerRankSetDetails.primary_loadout:type_name -> identity.ProtoPlayerLoadoutIdentity
	1,  // 5: rank.ProtoRankingEventClaim.player_ranking:type_name -> rank.ProtoPlayerRankSetDetails
	1,  // 6: rank.ProtoRankingEventClaim.top_rankings:type_name -> rank.ProtoPlayerRankSetDetails
	4,  // 7: rank.ProtoRankingEventClaim.statistics:type_name -> rank.ProtoRankingEventStats
	6,  // 8: rank.ProtoRankingEventStats.popular_region:type_name -> rank.ProtoRankingEventStatPopularRegion
	5,  // 9: rank.ProtoRankingEventStats.popular_job:type_name -> rank.ProtoRankingEventStatPopularJob
	7,  // 10: rank.ProtoRankingEventStats.popular_ability_card:type_name -> rank.ProtoRankingEventStatPopularAbilityCard
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_protorank_rank_proto_init() }
func file_protorank_rank_proto_init() {
	if File_protorank_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protorank_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingInfo); i {
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
		file_protorank_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoPlayerRankSetDetails); i {
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
		file_protorank_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingEventClaim); i {
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
		file_protorank_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingEventReward); i {
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
		file_protorank_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingEventStats); i {
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
		file_protorank_rank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingEventStatPopularJob); i {
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
		file_protorank_rank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingEventStatPopularRegion); i {
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
		file_protorank_rank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRankingEventStatPopularAbilityCard); i {
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
	file_protorank_rank_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ProtoRankingEventReward_RewardItem)(nil),
		(*ProtoRankingEventReward_RewardCard)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protorank_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protorank_rank_proto_goTypes,
		DependencyIndexes: file_protorank_rank_proto_depIdxs,
		MessageInfos:      file_protorank_rank_proto_msgTypes,
	}.Build()
	File_protorank_rank_proto = out.File
	file_protorank_rank_proto_rawDesc = nil
	file_protorank_rank_proto_goTypes = nil
	file_protorank_rank_proto_depIdxs = nil
}
