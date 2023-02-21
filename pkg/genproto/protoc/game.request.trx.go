package protoc

import "google.golang.org/protobuf/proto"

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

func NewFetchPlayerDataRequest(bytes []byte) (*FetchPlayerDataRequest, error) {
	req := &FetchPlayerDataRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewFollowPlayerRequest(bytes []byte) (*FollowPlayerRequest, error) {
	req := &FollowPlayerRequest{}
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

func NewTeleportRequest(bytes []byte) (*TeleportRequest, error) {
	req := &TeleportRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUnfollowPlayerRequest(bytes []byte) (*UnfollowPlayerRequest, error) {
	req := &UnfollowPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}