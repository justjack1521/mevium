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
