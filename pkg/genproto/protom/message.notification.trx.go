package protom

import "google.golang.org/protobuf/proto"

func NewNotification(bytes []byte) (*Notification, error) {
	req := &Notification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewSocialDataNotification(bytes []byte) (*SocialDataNotification, error) {
	req := &SocialDataNotification{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *Notification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SocialDataNotification) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
