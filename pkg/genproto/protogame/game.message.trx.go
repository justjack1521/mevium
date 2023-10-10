package protogame

import "google.golang.org/protobuf/proto"

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
