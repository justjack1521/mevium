package protosocial

import "google.golang.org/protobuf/proto"

func NewPlayerPresenceUpdateMessage(bytes []byte) (*PlayerPresenceUpdateMessage, error) {
	req := &PlayerPresenceUpdateMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerPresenceUpdateMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerPositionUpdateMessage(bytes []byte) (*PlayerPositionUpdateMessage, error) {
	req := &PlayerPositionUpdateMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerPositionUpdateMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerCompanionUpdateMessage(bytes []byte) (*PlayerCompanionUpdateMessage, error) {
	req := &PlayerCompanionUpdateMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerCompanionUpdateMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewPlayerCommentUpdateMessage(bytes []byte) (*PlayerCommentUpdateMessage, error) {
	req := &PlayerCommentUpdateMessage{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *PlayerCommentUpdateMessage) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
