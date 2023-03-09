// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: rank.response.proto

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

type SubmitScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeeklyRankUpdate    *protog.ProtoRankingInfo `protobuf:"bytes,1,opt,name=weekly_rank_update,json=weeklyRankUpdate,proto3" json:"weekly_rank_update,omitempty"`
	EventRankUpdate     *protog.ProtoRankingInfo `protobuf:"bytes,2,opt,name=event_rank_update,json=eventRankUpdate,proto3" json:"event_rank_update,omitempty"`
	WeeklyHighScore     bool                     `protobuf:"varint,3,opt,name=weekly_high_score,json=weeklyHighScore,proto3" json:"weekly_high_score,omitempty"`
	EventHighScore      bool                     `protobuf:"varint,4,opt,name=event_high_score,json=eventHighScore,proto3" json:"event_high_score,omitempty"`
	PreviousWeeklyScore float64                  `protobuf:"fixed64,5,opt,name=previous_weekly_score,json=previousWeeklyScore,proto3" json:"previous_weekly_score,omitempty"`
	PreviousEventScore  float64                  `protobuf:"fixed64,6,opt,name=previous_event_score,json=previousEventScore,proto3" json:"previous_event_score,omitempty"`
	PreviousWeeklyRank  int64                    `protobuf:"varint,7,opt,name=previous_weekly_rank,json=previousWeeklyRank,proto3" json:"previous_weekly_rank,omitempty"`
	PreviouslyEventRank int64                    `protobuf:"varint,8,opt,name=previously_event_rank,json=previouslyEventRank,proto3" json:"previously_event_rank,omitempty"`
}

func (x *SubmitScoreResponse) Reset() {
	*x = SubmitScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitScoreResponse) ProtoMessage() {}

func (x *SubmitScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rank_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitScoreResponse.ProtoReflect.Descriptor instead.
func (*SubmitScoreResponse) Descriptor() ([]byte, []int) {
	return file_rank_response_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitScoreResponse) GetWeeklyRankUpdate() *protog.ProtoRankingInfo {
	if x != nil {
		return x.WeeklyRankUpdate
	}
	return nil
}

func (x *SubmitScoreResponse) GetEventRankUpdate() *protog.ProtoRankingInfo {
	if x != nil {
		return x.EventRankUpdate
	}
	return nil
}

func (x *SubmitScoreResponse) GetWeeklyHighScore() bool {
	if x != nil {
		return x.WeeklyHighScore
	}
	return false
}

func (x *SubmitScoreResponse) GetEventHighScore() bool {
	if x != nil {
		return x.EventHighScore
	}
	return false
}

func (x *SubmitScoreResponse) GetPreviousWeeklyScore() float64 {
	if x != nil {
		return x.PreviousWeeklyScore
	}
	return 0
}

func (x *SubmitScoreResponse) GetPreviousEventScore() float64 {
	if x != nil {
		return x.PreviousEventScore
	}
	return 0
}

func (x *SubmitScoreResponse) GetPreviousWeeklyRank() int64 {
	if x != nil {
		return x.PreviousWeeklyRank
	}
	return 0
}

func (x *SubmitScoreResponse) GetPreviouslyEventRank() int64 {
	if x != nil {
		return x.PreviouslyEventRank
	}
	return 0
}

var File_rank_response_proto protoreflect.FileDescriptor

var file_rank_response_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x1a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x03, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x12, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x52, 0x61, 0x6e,
	0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x57, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x32, 0x0a, 0x15, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x6c, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75,
	0x73, 0x74, 0x6a, 0x61, 0x63, 0x6b, 0x31, 0x35, 0x32, 0x31, 0x2f, 0x6d, 0x65, 0x76, 0x69, 0x75,
	0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rank_response_proto_rawDescOnce sync.Once
	file_rank_response_proto_rawDescData = file_rank_response_proto_rawDesc
)

func file_rank_response_proto_rawDescGZIP() []byte {
	file_rank_response_proto_rawDescOnce.Do(func() {
		file_rank_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_rank_response_proto_rawDescData)
	})
	return file_rank_response_proto_rawDescData
}

var file_rank_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rank_response_proto_goTypes = []interface{}{
	(*SubmitScoreResponse)(nil),     // 0: rank.SubmitScoreResponse
	(*protog.ProtoRankingInfo)(nil), // 1: protog.ProtoRankingInfo
}
var file_rank_response_proto_depIdxs = []int32{
	1, // 0: rank.SubmitScoreResponse.weekly_rank_update:type_name -> protog.ProtoRankingInfo
	1, // 1: rank.SubmitScoreResponse.event_rank_update:type_name -> protog.ProtoRankingInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rank_response_proto_init() }
func file_rank_response_proto_init() {
	if File_rank_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rank_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitScoreResponse); i {
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
			RawDescriptor: file_rank_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rank_response_proto_goTypes,
		DependencyIndexes: file_rank_response_proto_depIdxs,
		MessageInfos:      file_rank_response_proto_msgTypes,
	}.Build()
	File_rank_response_proto = out.File
	file_rank_response_proto_rawDesc = nil
	file_rank_response_proto_goTypes = nil
	file_rank_response_proto_depIdxs = nil
}
