package protop

import "google.golang.org/protobuf/proto"

func (x *FollowPlayerResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UnfollowPlayerResponse) MarshallBinary() ([]byte, error) {
	return proto.Marshal(x)
}
