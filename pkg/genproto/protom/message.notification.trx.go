package protom

import "google.golang.org/protobuf/proto"

func NewNotification(bytes []byte) (*Notification, error) {
	req := &Notification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *Notification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewSocialDataNotification(bytes []byte) (*SocialDataNotification, error) {
	req := &SocialDataNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *SocialDataNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewBattleComplete(bytes []byte) (*BattleComplete, error) {
	req := &BattleComplete{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *BattleComplete) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewBattleRankingInfo(bytes []byte) (*BattleRankingInfo, error) {
	req := &BattleRankingInfo{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *BattleRankingInfo) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewBattleRankingClaim(bytes []byte) (*BattleRankingClaim, error) {
	req := &BattleRankingClaim{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *BattleRankingClaim) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
