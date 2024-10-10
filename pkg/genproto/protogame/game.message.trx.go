package protogame

import "google.golang.org/protobuf/proto"

func NewBattleCompleteMessage(bytes []byte) (*BattleCompleteMessage, error) {
	req := &BattleCompleteMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *BattleCompleteMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerLevelUpdateMessage(bytes []byte) (*PlayerLevelUpdateMessage, error) {
	req := &PlayerLevelUpdateMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerLevelUpdateMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerRentalCardUpdateMessage(bytes []byte) (*PlayerRentalCardUpdateMessage, error) {
	req := &PlayerRentalCardUpdateMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerRentalCardUpdateMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerProfileCreatedMessage(bytes []byte) (*PlayerProfileCreatedMessage, error) {
	req := &PlayerProfileCreatedMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerProfileCreatedMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerProfileDeletedMessage(bytes []byte) (*PlayerProfileDeletedMessage, error) {
	req := &PlayerProfileDeletedMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerProfileDeletedMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
