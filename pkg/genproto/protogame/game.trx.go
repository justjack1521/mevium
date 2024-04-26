package protogame

import (
	"google.golang.org/protobuf/proto"
)

func NewCreateProfileRequest(bytes []byte) (*CreateProfileRequest, error) {
	req := &CreateProfileRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewBattleCompleteRequest(bytes []byte) (*BattleCompleteRequest, error) {
	req := &BattleCompleteRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewBattleReviveRequest(bytes []byte) (*BattleReviveRequest, error) {
	req := &BattleReviveRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewBattleStartRequest(bytes []byte) (*BattleStartRequest, error) {
	req := &BattleStartRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCardAugmentRequest(bytes []byte) (*CardAugmentRequest, error) {
	req := &CardAugmentRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCardFavouriteRequest(bytes []byte) (*CardFavouriteRequest, error) {
	req := &CardFavouriteRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCardFusionRequest(bytes []byte) (*CardFusionRequest, error) {
	req := &CardFusionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCardBoostFusionRequest(bytes []byte) (*CardBoostFusionRequest, error) {
	req := &CardBoostFusionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCardSaleRequest(bytes []byte) (*CardSaleRequest, error) {
	req := &CardSaleRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCardTransferRequest(bytes []byte) (*CardTransferRequest, error) {
	req := &CardTransferRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewClaimLoginBonusRequest(bytes []byte) (*ClaimLoginCampaignRequest, error) {
	req := &ClaimLoginCampaignRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewClaimDailyMissionRequest(bytes []byte) (*ClaimDailyMissionRequest, error) {
	req := &ClaimDailyMissionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewSkillPanelUnlockRequest(bytes []byte) (*SkillPanelUnlockRequest, error) {
	req := &SkillPanelUnlockRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewClaimMailBoxItemRequest(bytes []byte) (*ClaimMailBoxItemRequest, error) {
	req := &ClaimMailBoxItemRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewClaimEventRankingRequest(bytes []byte) (*ClaimEventRankingRequest, error) {
	req := &ClaimEventRankingRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewClaimRentalCardRewardRequest(bytes []byte) (*ClaimRentalCardRewardRequest, error) {
	req := &ClaimRentalCardRewardRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewConfirmDailyMissionRequest(bytes []byte) (*ConfirmDailyMissionRequest, error) {
	req := &ConfirmDailyMissionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewDeckEditRequest(bytes []byte) (*DeckEditRequest, error) {
	req := &DeckEditRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewDeckEditAllRequest(bytes []byte) (*DeckEditAllRequest, error) {
	req := &DeckEditAllRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewFirstDailyLoginRequest(bytes []byte) (*FirstDailyLoginRequest, error) {
	req := &FirstDailyLoginRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewFetchPlayerDataRequest(bytes []byte) (*FetchPlayerDataRequest, error) {
	req := &FetchPlayerDataRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewDeleteAllMailboxItemRequest(bytes []byte) (*DeleteAllMailboxItemRequest, error) {
	req := &DeleteAllMailboxItemRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewStaminaRestoreRequest(bytes []byte) (*StaminaRestoreRequest, error) {
	req := &StaminaRestoreRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewExpandAbilityCardSlotRequest(bytes []byte) (*ExpandAbilityCardSlotRequest, error) {
	req := &ExpandAbilityCardSlotRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewTeleportRequest(bytes []byte) (*TeleportRequest, error) {
	req := &TeleportRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewProcessRegionEventRequest(bytes []byte) (*ProcessRegionEventRequest, error) {
	req := &ProcessRegionEventRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewProcessRegionNodeEventRequest(bytes []byte) (*ProcessRegionNodeEventRequest, error) {
	req := &ProcessRegionNodeEventRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUpdateProfileRequest(bytes []byte) (*UpdateProfileRequest, error) {
	req := &UpdateProfileRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewPurchaseItemRequest(bytes []byte) (*ItemShopItemPurchaseRequest, error) {
	req := &ItemShopItemPurchaseRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewPurchaseCardRequest(bytes []byte) (*ItemShopCardPurchaseRequest, error) {
	req := &ItemShopCardPurchaseRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewAbilityShopPurchaseRequest(bytes []byte) (*AbilityShopPurchaseRequest, error) {
	req := &AbilityShopPurchaseRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewCompleteRegionMapRequest(bytes []byte) (*CompleteRegionMapRequest, error) {
	req := &CompleteRegionMapRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewAbilityCardSummonRequest(bytes []byte) (*AbilityCardSummonRequest, error) {
	req := &AbilityCardSummonRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *AbilityCardSummonResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *BattleCompleteResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CompleteRegionMapResponse) MarshallBinary() ([]byte, error) {
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

func (x *SkillPanelUnlockResponse) MarshallBinary() ([]byte, error) {
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

func (x *CardTransferResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClaimRentalCardRewardResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClaimDailyMissionResponse) MarshallBinary() ([]byte, error) {
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

func (x *ExpandAbilityCardSlotResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *TeleportResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ProcessRegionEventResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ProcessRegionNodeEventResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UpdateProfileResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ItemShopItemPurchaseResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *AbilityShopPurchaseResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ItemShopCardPurchaseResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
