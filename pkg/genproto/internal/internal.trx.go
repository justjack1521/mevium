package protointernal

import "google.golang.org/protobuf/proto"

func NewClientConnectedMessage(bytes []byte) (*ClientConnected, error) {
	req := &ClientConnected{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ClientConnected) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewClientDisconnectedMessage(bytes []byte) (*ClientDisconnected, error) {
	req := &ClientDisconnected{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ClientDisconnected) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewClientHeartbeatMessage(bytes []byte) (*ClientHeartbeat, error) {
	req := &ClientHeartbeat{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ClientHeartbeat) MarshallBinary() ([]byte, error) {
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
