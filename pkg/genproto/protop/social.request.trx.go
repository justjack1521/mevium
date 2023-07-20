package protop

import "google.golang.org/protobuf/proto"

func NewFollowPlayerRequest(bytes []byte) (*FollowPlayerRequest, error) {
	req := &FollowPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewUnfollowPlayerRequest(bytes []byte) (*UnfollowPlayerRequest, error) {
	req := &UnfollowPlayerRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func NewFetchPlayerSocialInfo(bytes []byte) (*FetchPlayerSocialInfoRequest, error) {
	req := &FetchPlayerSocialInfoRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}
