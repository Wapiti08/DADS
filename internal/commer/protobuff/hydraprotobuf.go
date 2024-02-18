package protobuff

import (
	"errors"
	"io"
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
	// make the channel for syntherizing the communication
	outChan := make(chan interface{})
	l, err := net.Listen("tcp", listenaddress)
	if err != nil {
		return nil, err
	}

	go func() {
		defer l.Close()
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			}
			go func(conn net.Conn) {
				defer conn.Close()
				// keep reading message
				for {
					buffer, err := io.ReadAll(conn)
					if err != nil {
						break
					}
					// the message is blank
					if len(buffer) == 0 {
						continue
					}

					// send message
					obj, err := pSender.DecodeProto(buffer)
					if err != nil {
						continue
					}

					// use select to control the wait status of channel --- non-block operations
					select {
					case outChan <- obj:
					// avoid block when none of the other cases are immediately ready
					default:
					}

				}
			}(conn)
		}
	}()
	return outChan, nil
}

func (pSender *ProtoHandler) DecodeProto(data []byte) (*Ship, error) {
	pb := new(Ship)
	err := proto.Unmarshal(data, pb)
	return pb, err
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