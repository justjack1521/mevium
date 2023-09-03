package protosc

import "google.golang.org/protobuf/proto"

func NewFetchPlayerRankingInfoRequest(bytes []byte) (*JoinSocialChallengeRequest, error) {
	req := &JoinSocialChallengeRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *JoinSocialChallengeResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
