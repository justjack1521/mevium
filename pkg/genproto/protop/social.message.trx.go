package protop

import "google.golang.org/protobuf/proto"

func NewUpdatePlayerPresenceRequest(bytes []byte) (*UpdatePlayerPresenceRequest, error) {
	req := &UpdatePlayerPresenceRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUpdatePlayerPositionRequest(bytes []byte) (*UpdatePlayerPositionRequest, error) {
	req := &UpdatePlayerPositionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUpdatePlayerCompanionRequest(bytes []byte) (*UpdatePlayerCompanionRequest, error) {
	req := &UpdatePlayerCompanionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUpdatePlayerCommentRequest(bytes []byte) (*UpdatePlayerCommentRequest, error) {
	req := &UpdatePlayerCommentRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *UpdatePlayerPresenceRequest) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UpdatePlayerPositionRequest) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UpdatePlayerCompanionRequest) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UpdatePlayerCommentRequest) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
