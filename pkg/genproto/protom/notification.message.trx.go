package protom

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
