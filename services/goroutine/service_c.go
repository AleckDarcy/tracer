package goroutine

import (
	"ucsc/tracer/protobuf"
	"ucsc/tracer/services"
	"ucsc/tracer/services/messages"

	"github.com/golang/protobuf/proto"
)

type ServiceC struct {
	Meta *services.Meta
}

func NewServiceC(meta *services.Meta) services.Receiver {
	return &ServiceC{Meta: meta}
}

func (s *ServiceC) ReceiveRequest(bytes []byte) error {
	reqB := &messages.RequestC{}
	err := protobuf.Unmarshal(bytes, reqB)
	if err != nil {
		if err == protobuf.MockedCrash {
			return nil
		}
	}

	req := reqB.Body

	repA := &messages.ResponseA{Body: &messages.Response{From: req.To, To: req.From, Status: 200}}

	s.SendResponse(req.From, repA)

	return nil
}

func (s *ServiceC) ReceiveResponse(bytes []byte) error {
	return nil
}

func (s *ServiceC) SendResponse(name string, rep proto.Message) error {
	bytes, err := protobuf.Marshal(rep)
	if err != nil {
		return err
	}

	return services.SendHelper.SendResponse(name, bytes)
}
