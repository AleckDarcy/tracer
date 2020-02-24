package goroutine

import (
	"errors"
	"fmt"
	"time"
	"ucsc/tracer/protobuf"
	"ucsc/tracer/services"
	"ucsc/tracer/services/messages"
	"ucsc/tracer/trace"
)

type ServiceA struct {
	Meta *services.Meta

	ResponseB chan *messages.Response
	ResponseC chan *messages.Response
}

func NewServiceA(meta *services.Meta) services.Receiver {
	return &ServiceA{
		Meta:      meta,
		ResponseB: make(chan *messages.Response), // mocked
		ResponseC: make(chan *messages.Response), // mocked
	}
}

func (s *ServiceA) SendRequest(name string, req protobuf.Tracer, signal chan *messages.Response) *messages.Response {
	bytes, err := protobuf.Marshal(req)
	if err != nil {
		return &messages.Response{Status: -1}
	}
	if err := services.SendHelper.SendRequest(name, bytes); err != nil {
		return &messages.Response{Status: -2}
	}

	var rep *messages.Response
	select {
	case <-time.After(services.Timeout):
		messages.Helper.Delete(req.GetMsgID())
		rep = &messages.Response{Status: -1}
	case rep = <-signal:
	}

	return rep
}

func (s *ServiceA) SendRequestBuggy(name string, req protobuf.Tracer, signal chan *messages.Response) *messages.Response {
	bytes, err := protobuf.Marshal(req)
	if err != nil {
		return &messages.Response{Status: -1}
	}
	if err := services.SendHelper.SendRequest(name, bytes); err != nil {
		return &messages.Response{Status: -2}
	}

	var rep *messages.Response
	select {
	case <-time.After(services.Timeout):
		messages.Helper.Delete(req.GetMsgID())
		//rep = &Response{Status: -1}
	case rep = <-signal:
	}

	return rep
}

func (s *ServiceA) ReceiveRequest(bytes []byte) error {
	reqA := &messages.RequestA{}
	protobuf.Unmarshal(bytes, reqA)
	req := reqA.Body

	reqB1 := &messages.RequestB{Body: &messages.Request{Id: messages.NewID(), From: "A", To: "B", Trace: req.Trace}}
	repB1 := s.SendRequest("B", reqB1, s.ResponseB)
	if repB1.Status != 200 {
		fmt.Println("calling B fail, status code:", repB1.Status)

		return errors.New("calling B fail")
	}

	reqC := &messages.RequestC{Body: &messages.Request{Id: messages.NewID(), From: "A", To: "C", Trace: req.Trace}}
	repC := s.SendRequest("C", reqC, s.ResponseC)
	if repC.Status != 200 {
		fmt.Println("calling C fail, status code:", repC.Status)

		return errors.New("calling C fail")
	}

	reqB2 := &messages.RequestB{Body: &messages.Request{Id: messages.NewID(), From: "A", To: "B", Trace: req.Trace}}
	repB2 := s.SendRequestBuggy("B", reqB2, s.ResponseB)
	if repB2.Status != 200 { // crashes when timeout happens
		fmt.Println("calling B fail, status code:", repB2.Status)

		return errors.New("calling B fail")
	}

	//// respond to client
	//rep := &Response{}
	//b, err := protobuf.Marshal(rep)
	//if err != nil {
	//	return err
	//}

	trace.PrintTrace(repB2.GetTrace())

	return nil
}

func (s *ServiceA) ReceiveResponse(bytes []byte) error {
	repA := &messages.ResponseA{}
	if err := protobuf.Unmarshal(bytes, repA); err != nil {
		return err
	}

	rep := repA.Body

	switch rep.From {
	case "B":
		fmt.Println("A received response from B")
		s.ResponseB <- rep
	case "C":
		fmt.Println("A received response from C")
		s.ResponseC <- rep
	default:
		panic("unreachable code")
	}

	return nil
}
