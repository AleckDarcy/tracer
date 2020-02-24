package protobuf

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
	"ucsc/tracer/trace"

	"github.com/golang/protobuf/proto"
)

func caller() string {
	_, file, _, _ := runtime.Caller(2)

	return file[strings.LastIndex(file, "/")+1:]
}

var MockedCrash = errors.New("mocked crash")

type tracingDataStore struct {
	Data map[int64]string
}

type Tracer interface {
	GetMsgID() int64
	GetTrace() *trace.Chain
	SetTrace(data *trace.Chain)

	proto.Message
}

func Marshal(msg proto.Message) ([]byte, error) {
	//fmt.Println("marshal", msg)
	if m, ok := msg.(Tracer); ok {
		goID := runtime.Goid()

		trace.Lock.Lock()
		defer trace.Lock.Unlock()

		if traceID, ok := trace.ThreadHelper.Get(goID); ok {
			name := reflect.TypeOf(msg).Elem().Name()

			if old, ok := trace.Helper.Get(traceID); ok {
				log.Printf("%s Marshal: append trace to trace %d", caller(), traceID)
				old.Data = &trace.Data{
					Type:        trace.Type_Send,
					Timestamp:   time.Now().UnixNano(),
					MessageName: name,
					Previous:    old.Data,
				}
				old.Depth++

				trace.Helper.Set(old)
				m.SetTrace(old)
			} else {
				log.Printf("%s Marshal: err: unrecognized trace %d", caller(), traceID)
				fmt.Printf("%v\n", trace.Helper)
				panic("unreachable code")
			}
		}
	}

	return proto.Marshal(msg)
}

func Unmarshal(bytes []byte, msg proto.Message) error {
	err := proto.Unmarshal(bytes, msg)
	if err != nil {
		return err
	}

	if m, ok := msg.(Tracer); ok {
		goID := runtime.Goid()

		trace.Lock.Lock()
		defer trace.Lock.Unlock()

		if data := m.GetTrace(); data != nil {
			name := reflect.TypeOf(msg).Elem().Name()

			old, ok := trace.Helper.Get(data.Id)
			if !ok { // new trace id
				log.Printf("%s Unmarshal: tracing enabled with trace %d", caller(), data.Id)
				old = data
			} else {
				if old.Id != data.Id {
					panic("trace id must be fixed")
				}
				log.Printf("%s Unmarshal: found existing trace %d", caller(), data.Id)
			}

			if old.Rlfi != nil && name == old.Rlfi.Name {
				fmt.Println("RLFI: mocked crash when receiving", name)
				return MockedCrash
			}

			// judge TFI
			if old.Tfi != nil && name == old.Tfi.Name {
				if trace.JudgeTFI(name, old) {
					fmt.Println("TFI: mocked crash when receiving", name, "after", old.Tfi.After)
					return MockedCrash
				}
			}

			old.Data = &trace.Data{
				Type:        trace.Type_Receive,
				Timestamp:   time.Now().UnixNano(),
				MessageName: name,
				Previous:    data.Data,
			}
			old.Depth++
			log.Printf("%s Unmarshal: set trace %d thread %d", caller(), old.Id, goID)

			trace.Helper.Set(old)
			trace.ThreadHelper.Set(goID, old.Id)
			m.SetTrace(old)
		} else {
			trace.ThreadHelper.Delete(goID)
		}
	}

	return nil
}
