package protobuff

import (
	"errors"
	"log"
	"net"

	"google.golang.org/golang/protobuf/proto"
)

type ProtoHandler struct{}

func NewProtoHandler() *ProtoHandler {
	return new(ProtoHandler)
}

// func EncodeProto(obj interface{}) ([]byte, error) {
// 	// check whether obj contains a value of type "Ship"
// 	if v, ok := obj.(*Ship); ok {
// 		return proto.Marshal(v)
// 	}
// 	return nil, errors.New("Proto: unknown message type")
// }

// func DecodeProto(buffer []byte) (*Ship, error) {
// 	pb := new(Ship)
// 	return pb, proto.Unmarshal(buffer, pb)
// }


// encoder and sender
func (pSender *ProtoHandler) EncodeAndSend(obj interface{}, dest string) error {
	// check whether obj contains Ship
	v, ok := obj.(*Ship)
	if !ok {
		return errors.New("Proto: Unknown message type")
	}

	// encode message
	data, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	// send message
	return sendmessage(data, dest)
}

func (pSender *ProtoHandler) ListenAndDecode(listenaddress string) (chan interface{}, error) {

	l, err := net.Listen("tcp", listenaddress)
	if err != nil {
		return nil, err
	}

	go func() {
		defer l.Close()
		for {
			conn, err := l.Accept()
				
			go func(c net.Conn) {
				defer c.Close()
				for {
			
		

		// use select to control the wait status of channel --- non-block operations

				}
			}(c)
		}
	}()




}



func sendmessage(data []byte, dest string) error {
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		return err
	}

	defer conn.Close()

	log.Printf("Sending %d bytes to %s \n", len(data), dest)
	_, err = conn.Write(data)
	return err
}