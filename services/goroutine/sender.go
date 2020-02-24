package goroutine

import (
	"fmt"
	"ucsc/tracer/services"
)

var ThreadCount = 10

type Sender struct {
	Metas     map[string]*services.Meta
	Receivers map[string]services.Receiver
}

func (s *Sender) SendRequest(name string, bytes []byte) error {
	receiver, ok := s.Receivers[name]
	if !ok {
		panic("unreachable code")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic when calling", name, "to respond")
		}
	}()
	//fmt.Println("send request to", name)
	go receiver.ReceiveRequest(bytes)

	return nil
}

func (s *Sender) SendResponse(name string, bytes []byte) error {
	receiver, ok := s.Receivers[name]
	if !ok {
		panic("unreachable code")
	}

	//fmt.Println("send response to", name)

	go receiver.ReceiveResponse(bytes)

	return nil
}
