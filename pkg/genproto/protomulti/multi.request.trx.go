package protomulti

import "google.golang.org/protobuf/proto"

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
