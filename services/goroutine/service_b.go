package goroutine

import (
	"log"
	"ucsc/tracer/protobuf"
	"ucsc/tracer/services"
	"ucsc/tracer/services/messages"

	"github.com/golang/protobuf/proto"
)

type ServiceB struct {
	Meta *services.Meta
}

func NewServiceB(meta *services.Meta) services.Receiver {
	return &ServiceB{Meta: meta}
}

func (s *ServiceB) ReceiveRequest(bytes []byte) error {
	reqB := &messages.RequestB{}
	if err := protobuf.Unmarshal(bytes, reqB); err != nil {
		if err == protobuf.MockedCrash {
			return nil
		}

		log.Printf("ServiceB.ReceiveRequest Unmarshal err: %s", err)
		return err
	}

	req := reqB.Body

	repA := &messages.ResponseA{Body: &messages.Response{From: req.To, To: req.From, Status: 200}}

	s.SendResponse(req.From, repA)

	return nil
}

func (s *ServiceB) ReceiveResponse(bytes []byte) error {
	return nil
}

func (s *ServiceB) SendResponse(name string, rep proto.Message) error {
	bytes, err := protobuf.Marshal(rep)
	if err != nil {
		return err
	}

	return services.SendHelper.SendResponse(name, bytes)
}
