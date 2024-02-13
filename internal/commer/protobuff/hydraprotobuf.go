package protobuff

import (
	"errors"

	"google.golang.org/golang/protobuf/proto"
)

func EncodeProto(obj interface{}) ([]byte, error) {
	// check whether obj contains a value of type "Ship"
	if v, ok := obj.(*Ship); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("Proto: unknown message type")
}

func DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}