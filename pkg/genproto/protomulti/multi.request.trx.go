package protomulti

import "google.golang.org/protobuf/proto"

func NewCreateSessionRequest(bytes []byte) (*CreateSessionRequest, error) {
	req := &CreateSessionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *CreateSessionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewEndSessionRequest(bytes []byte) (*EndSessionRequest, error) {
	req := &EndSessionRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *EndSessionResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewSearchLobbyRequest(bytes []byte) (*SearchLobbyRequest, error) {
	req := &SearchLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *SearchLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewWatchLobbyRequest(bytes []byte) (*WatchLobbyRequest, error) {
	req := &WatchLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *WatchLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewDiscardLobbyRequest(bytes []byte) (*DiscardLobbyRequest, error) {
	req := &DiscardLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *DiscardLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewCreateLobbyRequest(bytes []byte) (*CreateLobbyRequest, error) {
	req := &CreateLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *CreateLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewCancelLobbyRequest(bytes []byte) (*CancelLobbyRequest, error) {
	req := &CancelLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *CancelLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewJoinLobbyRequest(bytes []byte) (*JoinLobbyRequest, error) {
	req := &JoinLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *JoinLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewReadyLobbyRequest(bytes []byte) (*ReadyLobbyRequest, error) {
	req := &ReadyLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *ReadyLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewUnreadyLobbyRequest(bytes []byte) (*UnreadyLobbyRequest, error) {
	req := &UnreadyLobbyRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *UnreadyLobbyResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGetLobbyPlayerRequest(bytes []byte) (*GetLobbyPlayerRequest, error) {
	req := &GetLobbyPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GetLobbyPlayerResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
