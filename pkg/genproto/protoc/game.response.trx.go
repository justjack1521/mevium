package protoc

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func (x *Response) Error() string {
	return fmt.Sprintf("status %d: err %v", x.Header.ErrorCode, x.Header.Error)
}

func (x *Response) MarshallBinary() ([]byte, error) {
	bytes, err := proto.Marshal(x)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (x *BattleCompleteResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *BattleReviveResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *BattleStartResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CreateProfileResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CardAugmentResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CardFavouriteResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CardFusionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CardBoostFusionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CardSaleResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClaimMailBoxItemResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClaimEventRankingResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClaimRentalCardRewardResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DeleteAllMailboxItemResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DeckEditAllResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClaimLoginCampaignResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ConfirmDailyMissionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *FirstDailyLoginResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *FetchPlayerDataResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *StaminaRestoreResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *TeleportResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UpdateProfileResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
