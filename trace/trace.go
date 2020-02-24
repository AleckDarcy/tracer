package trace

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var Lock sync.Mutex

func caller() string {
	_, file, _, _ := runtime.Caller(2)

	return file[strings.LastIndex(file, "/")+1:]
}

var idGenerator int64

func NewID() int64 {
	return atomic.AddInt64(&idGenerator, 1)
}

type helper struct {
	data map[int64]*Chain // key: trace id
}

var Helper = &helper{
	data: map[int64]*Chain{},
}

func (m *helper) Set(data *Chain) {
	//fmt.Printf("trace helper set trace id: %d data: %v\n", data.Id, data)
	m.data[data.Id] = data
	//log.Printf("%s set trace %s", caller(), data)
}

func (m *helper) Get(id int64) (*Chain, bool) {
	req, ok := m.data[id]
	if ok {
		tmp := *req
		delete(m.data, id)

		return &tmp, true
	}

	return nil, false
}

func (m *helper) Delete(id int64) {
	delete(m.data, id)
}

type threadHelper struct {
	data map[int64]int64 // key: thread id, value: trace id
}

var ThreadHelper = &threadHelper{
	data: map[int64]int64{},
}

func (h *threadHelper) Set(id int64, traceID int64) {
	//fmt.Printf("thread helper set thread id: %d data: %v\n", id, data)

	h.data[id] = traceID
}

func (h *threadHelper) Get(id int64) (int64, bool) {
	threadID, ok := h.data[id]

	return threadID, ok
}

func (h *threadHelper) Delete(id int64) {
	if _, ok := h.data[id]; ok {
		delete(h.data, id)
	}
}

func PrintTrace(chain *Chain) {
	depth := int(chain.Depth)
	data := chain.Data

	strs := make([]string, depth)

	for i := depth - 1; i >= 0; i-- {
		if data == nil {
			go func() {
				log.Fatal("oh no! ", chain)
			}()

			break
		}
		t := time.Unix(0, data.Timestamp)

		strs[i] = fmt.Sprintf(""+
			"%v: %v %v",
			t.Format("01/2/2006 15:04:05.999"), data.Type, data.MessageName,
		)

		data = data.Previous
	}

	for i := 0; i < depth; i++ {
		fmt.Println(strs[i])
	}
}

func JudgeTFI(name string, chain *Chain) bool {
	after := chain.Tfi.After
	if len(after) == 0 {
		return false
	}

	depth := int(chain.Depth)
	data := chain.Data

	count := 0
	counts := make([]int, len(after))
	for i := depth - 1; i >= 0; i-- {
		if data.Type == Type_Receive {
			for j, a := range after {
				if a == data.MessageName {
					if counts[j] == 0 {
						counts[j] = 1
						count++
					}
				}
			}
		}

		data = data.Previous
	}

	return count == len(after)
}
