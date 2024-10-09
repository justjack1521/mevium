package protocommon

import (
	"google.golang.org/protobuf/proto"
)

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

func NewUserCreatedMessage(bytes []byte) (*UserCreated, error) {
	req := &UserCreated{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *UserCreated) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewUserDeletedMessage(bytes []byte) (*UserDeleted, error) {
	req := &UserDeleted{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *UserDeleted) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
