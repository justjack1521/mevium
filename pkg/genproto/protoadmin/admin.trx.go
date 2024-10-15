package protoadmin

import "google.golang.org/protobuf/proto"

func NewGrantItemRequest(bytes []byte) (*GrantItemRequest, error) {
	req := &GrantItemRequest{}
	if err := proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (x *GrantItemResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
