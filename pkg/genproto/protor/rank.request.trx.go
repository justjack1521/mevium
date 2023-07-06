package protor

import "google.golang.org/protobuf/proto"

func NewFetchPlayerRankingInfoRequest(bytes []byte) (*FetchPlayerRankingInfoRequest, error) {
	req := &FetchPlayerRankingInfoRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *FetchPlayerRankingInfoResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
