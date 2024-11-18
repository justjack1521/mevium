package protomulti

import "google.golang.org/protobuf/proto"

func NewSessionCreateRequest(bytes []byte) (*SessionCreateRequest, error) {
	req := &SessionCreateRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *SessionCreateResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewSessionEndRequest(bytes []byte) (*SessionEndRequest, error) {
	req := &SessionEndRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *SessionEndResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbySearchRequest(bytes []byte) (*LobbySearchRequest, error) {
	req := &LobbySearchRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbySearchResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyStartRequest(bytes []byte) (*LobbyStartRequest, error) {
	req := &LobbyStartRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyStartResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantWatchRequest(bytes []byte) (*ParticipantWatchRequest, error) {
	req := &ParticipantWatchRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantWatchResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantUnwatchRequest(bytes []byte) (*ParticipantUnwatchRequest, error) {
	req := &ParticipantUnwatchRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantUnwatchResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyCreateRequest(bytes []byte) (*LobbyCreateRequest, error) {
	req := &LobbyCreateRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyCreateResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyCancelRequest(bytes []byte) (*LobbyCancelRequest, error) {
	req := &LobbyCancelRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyCancelResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantJoinRequest(bytes []byte) (*ParticipantJoinRequest, error) {
	req := &ParticipantJoinRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantJoinResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantLeaveRequest(bytes []byte) (*ParticipantLeaveRequest, error) {
	req := &ParticipantLeaveRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantLeaveResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantReadyRequest(bytes []byte) (*ParticipantReadyRequest, error) {
	req := &ParticipantReadyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantReadyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewParticipantUnreadyRequest(bytes []byte) (*ParticipantUnreadyRequest, error) {
	req := &ParticipantUnreadyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ParticipantUnreadyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyReadyRequest(bytes []byte) (*LobbyReadyRequest, error) {
	req := &LobbyReadyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyReadyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewLobbyStampRequest(bytes []byte) (*LobbyStampRequest, error) {
	req := &LobbyStampRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *LobbyStampResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGetGameRequest(bytes []byte) (*GetGameRequest, error) {
	req := &GetGameRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GetGameResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGameReadyPlayerRequest(bytes []byte) (*GameReadyPlayerRequest, error) {
	req := &GameReadyPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GameReadyPlayerResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGameEnqueueAbilityRequest(bytes []byte) (*GameEnqueueActionRequest, error) {
	req := &GameEnqueueActionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GameEnqueueActionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGameDequeueAbilityRequest(bytes []byte) (*GameDequeueActionRequest, error) {
	req := &GameDequeueActionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GameDequeueActionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGameLockActionRequest(bytes []byte) (*GameLockActionRequest, error) {
	req := &GameLockActionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GameLockActionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
