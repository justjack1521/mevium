package protosc

import "google.golang.org/protobuf/proto"

func NewJoinSocialChallengeRequest(bytes []byte) (*JoinSocialChallengeRequest, error) {
	req := &JoinSocialChallengeRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *JoinSocialChallengeResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func NewGetPlayerChallengeRequest(bytes []byte) (*GetPlayerChallengeRequest, error) {
	req := &GetPlayerChallengeRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GetPlayerChallengeResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
