package protosocial

import "google.golang.org/protobuf/proto"

func NewFollowPlayerRequest(bytes []byte) (*FollowPlayerRequest, error) {
	req := &FollowPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *FollowPlayerResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewUnfollowPlayerRequest(bytes []byte) (*UnfollowPlayerRequest, error) {
	req := &UnfollowPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *UnfollowPlayerResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewFetchPlayerSocialInfo(bytes []byte) (*FetchPlayerSocialInfoRequest, error) {
	req := &FetchPlayerSocialInfoRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *FetchPlayerSocialInfoResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
