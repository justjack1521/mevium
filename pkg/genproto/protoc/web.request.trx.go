package protoc

import "google.golang.org/protobuf/proto"

func NewCreateProfileRequest(bytes []byte) (*CreateProfileRequest, error) {
	req := &CreateProfileRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *CreateProfileResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
