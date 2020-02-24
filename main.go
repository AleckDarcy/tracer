package main

import (
	"fmt"
	"ucsc/tracer/services"
	"ucsc/tracer/services/goroutine"
	"ucsc/tracer/services/messages"
	"ucsc/tracer/trace"

	"github.com/golang/protobuf/proto"
)

//go:generate protoc --go_out=. services/messages/message.proto
//go:generate protoc --go_out=. trace/trace.proto

var sender services.Sender

func main() {
	// no parameters, start in goroutine model

	receivers := map[string]services.Receiver{}
	receivers["A"] = goroutine.NewServiceA(services.Metas["A"])
	receivers["B"] = goroutine.NewServiceB(services.Metas["B"])
	receivers["C"] = goroutine.NewServiceC(services.Metas["C"])

	services.SendHelper = &goroutine.Sender{Metas: services.Metas, Receivers: receivers}

	//threadCount := 10
	//// normal request
	//for i := 0; i < threadCount; i++ {
	//	go func() {
	//		for {
	//			reqA := &messages.RequestA{
	//				Body: &messages.Request{
	//					From: "client",
	//					To:   "A",
	//					Trace: &trace.Chain{
	//						Id: trace.NewID(),
	//					},
	//				},
	//			}
	//
	//			bytes, _ := proto.Marshal(reqA)
	//
	//			receivers["A"].ReceiveRequest(bytes)
	//		}
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//
	//return
	//fmt.Println("===================")

	// crash C after B
	{
		reqA := &messages.RequestA{
			Body: &messages.Request{
				From: "client",
				To:   "A",
				Trace: &trace.Chain{
					Id:  trace.NewID(),
					Tfi: &trace.TFI{Name: "RequestC", After: []string{"RequestB"}},
				},
			},
		}

		bytes, _ := proto.Marshal(reqA)
		receivers["A"].ReceiveRequest(bytes)
	}
	fmt.Println("===================")
	// crash B
	{
		reqA := &messages.RequestA{
			Body: &messages.Request{
				From: "client",
				To:   "A",
				Trace: &trace.Chain{
					Id:   trace.NewID(),
					Rlfi: &trace.RLFI{Name: "RequestB"},
				},
			},
		}

		bytes, _ := proto.Marshal(reqA)
		receivers["A"].ReceiveRequest(bytes)
	}
	fmt.Println("===================")
	// crash B after B, the demo will crash
	// the framework in practise knows how to restart the service, but we won't
	{
		reqA := &messages.RequestA{
			Body: &messages.Request{
				From: "client",
				To:   "A",
				Trace: &trace.Chain{
					Id:  trace.NewID(),
					Tfi: &trace.TFI{Name: "RequestB", After: []string{"RequestB"}},
				},
			},
		}

		bytes, _ := proto.Marshal(reqA)
		receivers["A"].ReceiveRequest(bytes)
	}
}
