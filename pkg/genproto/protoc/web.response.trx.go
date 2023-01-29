package protoc

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func (x *Response) Error() string {
	return fmt.Sprintf("status %d: err %v", x.Header.ErrorCode, x.Header.Error)
}

func (x *Response) MarshallBinary() ([]byte, error) {
	bytes, err := proto.Marshal(x)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (x *CreateProfileResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
