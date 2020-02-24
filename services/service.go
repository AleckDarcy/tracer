package services

import "time"

var Timeout = time.Second

type Meta struct {
	Name        string
	Port        int // mocked in goroutine model
	CalleeNames []string
	Sender      Sender
}

var Metas = map[string]*Meta{
	"A": &Meta{Name: "A", Port: 10000, CalleeNames: []string{"B", "C", "B"}},
	"B": &Meta{Name: "B", Port: 10001, CalleeNames: []string{}},
	"C": &Meta{Name: "C", Port: 10002, CalleeNames: []string{}},
}

type Sender interface {
	SendRequest(name string, bytes []byte) error
	SendResponse(name string, bytes []byte) error
}

var SendHelper Sender

type Receiver interface {
	ReceiveRequest(bytes []byte) error
	ReceiveResponse(bytes []byte) error
}
