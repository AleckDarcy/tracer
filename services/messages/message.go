package messages

import (
	"sync"
	"sync/atomic"
	"ucsc/tracer/trace"
)

var idGenerator int64

func NewID() int64 {
	return atomic.AddInt64(&idGenerator, 1)
}

type helper struct { // mocked, handle message id, timeout
	data map[int64]struct{} // key: message id

	lock sync.Mutex
}

var Helper = &helper{
	data: map[int64]struct{}{},
}

func (m *helper) Set(id int64) {
	m.lock.Lock()
	m.data[id] = struct{}{}
	m.lock.Unlock()
}

func (m *helper) Get(id int64) bool {
	m.lock.Lock()
	_, ok := m.data[id]
	m.lock.Unlock()

	return ok
}

func (m *helper) Delete(id int64) bool {
	m.lock.Lock()
	_, ok := m.data[id]
	if ok {
		delete(m.data, id)
	}
	m.lock.Unlock()

	return ok
}

func (r *RequestA) GetMsgID() int64 {
	return r.Body.Id
}

func (r *RequestA) GetTrace() *trace.Chain {
	return r.Body.Trace
}

func (r *RequestA) SetTrace(data *trace.Chain) {
	r.Body.Trace = data
}

func (r *RequestB) GetMsgID() int64 {
	return r.Body.Id
}

func (r *RequestB) GetTrace() *trace.Chain {
	return r.Body.Trace
}

func (r *RequestB) SetTrace(data *trace.Chain) {
	r.Body.Trace = data
}

func (r *RequestC) GetMsgID() int64 {
	return r.Body.Id
}

func (r *RequestC) GetTrace() *trace.Chain {
	return r.Body.Trace
}

func (r *RequestC) SetTrace(data *trace.Chain) {
	r.Body.Trace = data
}

func (r *ResponseA) GetMsgID() int64 {
	return r.Body.Id
}

func (r *ResponseA) GetTrace() *trace.Chain {
	return r.Body.Trace
}

func (r *ResponseA) SetTrace(data *trace.Chain) {
	r.Body.Trace = data
}

func (r *ResponseB) GetMsgID() int64 {
	return r.Body.Id
}

func (r *ResponseB) GetTrace() *trace.Chain {
	return r.Body.Trace
}

func (r *ResponseB) SetTrace(data *trace.Chain) {
	r.Body.Trace = data
}

func (r *ResponseC) GetMsgID() int64 {
	return r.Body.Id
}

func (r *ResponseC) GetTrace() *trace.Chain {
	return r.Body.Trace
}

func (r *ResponseC) SetTrace(data *trace.Chain) {
	r.Body.Trace = data
}
