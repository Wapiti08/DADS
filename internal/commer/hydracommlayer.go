package commer

import "DADS/internal/commer/protobuff"

const {
	// iota defines a sequence starting with 0
	Protobuf unit8 = iota
}

// define scalable connection interface
type HydraConnection interface {
	EncodeAndSend(obj interface{}, dest string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}


// define the connection bwetween buf and tcp stream
func NewConnection(connType string) HydraConnection {
	switch connType {
	case Protobuf:
		return protobuff.NewProtoHandler()
	}
	return nil
}
