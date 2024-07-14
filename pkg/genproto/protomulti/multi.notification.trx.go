package protomulti

import "google.golang.org/protobuf/proto"

func NewParticipantJoinNotification(bytes []byte) (*ParticipantJoinNotification, error) {
	req := &ParticipantJoinNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantJoinNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantLeaveNotification(bytes []byte) (*ParticipantLeaveNotification, error) {
	req := &ParticipantLeaveNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantLeaveNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantReadyNotification(bytes []byte) (*ParticipantReadyNotification, error) {
	req := &ParticipantReadyNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantReadyNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantUnreadyNotification(bytes []byte) (*ParticipantUnreadyNotification, error) {
	req := &ParticipantUnreadyNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantUnreadyNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewStampSendNotification(bytes []byte) (*StampSendNotification, error) {
	req := &StampSendNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *StampSendNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantDeckChangeNotification(bytes []byte) (*ParticipantDeckChangeNotification, error) {
	req := &ParticipantDeckChangeNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantDeckChangeNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyCancelNotification(bytes []byte) (*LobbyCancelNotification, error) {
	req := &LobbyCancelNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyCancelNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyStartNotification(bytes []byte) (*LobbyStartNotification, error) {
	req := &LobbyStartNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyStartNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
