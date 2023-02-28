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
