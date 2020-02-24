// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/messages/message.proto

package messages

import (
	fmt "fmt"
	math "math"
	trace "ucsc/tracer/trace"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestA struct {
	Body                 *Request `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestA) Reset()         { *m = RequestA{} }
func (m *RequestA) String() string { return proto.CompactTextString(m) }
func (*RequestA) ProtoMessage()    {}
func (*RequestA) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{0}
}

func (m *RequestA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestA.Unmarshal(m, b)
}
func (m *RequestA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestA.Marshal(b, m, deterministic)
}
func (m *RequestA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestA.Merge(m, src)
}
func (m *RequestA) XXX_Size() int {
	return xxx_messageInfo_RequestA.Size(m)
}
func (m *RequestA) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestA.DiscardUnknown(m)
}

var xxx_messageInfo_RequestA proto.InternalMessageInfo

func (m *RequestA) GetBody() *Request {
	if m != nil {
		return m.Body
	}
	return nil
}

type RequestB struct {
	Body                 *Request `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestB) Reset()         { *m = RequestB{} }
func (m *RequestB) String() string { return proto.CompactTextString(m) }
func (*RequestB) ProtoMessage()    {}
func (*RequestB) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{1}
}

func (m *RequestB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestB.Unmarshal(m, b)
}
func (m *RequestB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestB.Marshal(b, m, deterministic)
}
func (m *RequestB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestB.Merge(m, src)
}
func (m *RequestB) XXX_Size() int {
	return xxx_messageInfo_RequestB.Size(m)
}
func (m *RequestB) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestB.DiscardUnknown(m)
}

var xxx_messageInfo_RequestB proto.InternalMessageInfo

func (m *RequestB) GetBody() *Request {
	if m != nil {
		return m.Body
	}
	return nil
}

type RequestC struct {
	Body                 *Request `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestC) Reset()         { *m = RequestC{} }
func (m *RequestC) String() string { return proto.CompactTextString(m) }
func (*RequestC) ProtoMessage()    {}
func (*RequestC) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{2}
}

func (m *RequestC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestC.Unmarshal(m, b)
}
func (m *RequestC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestC.Marshal(b, m, deterministic)
}
func (m *RequestC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestC.Merge(m, src)
}
func (m *RequestC) XXX_Size() int {
	return xxx_messageInfo_RequestC.Size(m)
}
func (m *RequestC) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestC.DiscardUnknown(m)
}

var xxx_messageInfo_RequestC proto.InternalMessageInfo

func (m *RequestC) GetBody() *Request {
	if m != nil {
		return m.Body
	}
	return nil
}

type Request struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From                 string       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string       `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Trace                *trace.Chain `protobuf:"bytes,10,opt,name=trace,proto3" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Request) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Request) GetTrace() *trace.Chain {
	if m != nil {
		return m.Trace
	}
	return nil
}

type ResponseA struct {
	Body                 *Response `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResponseA) Reset()         { *m = ResponseA{} }
func (m *ResponseA) String() string { return proto.CompactTextString(m) }
func (*ResponseA) ProtoMessage()    {}
func (*ResponseA) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{4}
}

func (m *ResponseA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseA.Unmarshal(m, b)
}
func (m *ResponseA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseA.Marshal(b, m, deterministic)
}
func (m *ResponseA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseA.Merge(m, src)
}
func (m *ResponseA) XXX_Size() int {
	return xxx_messageInfo_ResponseA.Size(m)
}
func (m *ResponseA) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseA.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseA proto.InternalMessageInfo

func (m *ResponseA) GetBody() *Response {
	if m != nil {
		return m.Body
	}
	return nil
}

type ResponseB struct {
	Body                 *Response `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResponseB) Reset()         { *m = ResponseB{} }
func (m *ResponseB) String() string { return proto.CompactTextString(m) }
func (*ResponseB) ProtoMessage()    {}
func (*ResponseB) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{5}
}

func (m *ResponseB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseB.Unmarshal(m, b)
}
func (m *ResponseB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseB.Marshal(b, m, deterministic)
}
func (m *ResponseB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseB.Merge(m, src)
}
func (m *ResponseB) XXX_Size() int {
	return xxx_messageInfo_ResponseB.Size(m)
}
func (m *ResponseB) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseB.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseB proto.InternalMessageInfo

func (m *ResponseB) GetBody() *Response {
	if m != nil {
		return m.Body
	}
	return nil
}

type ResponseC struct {
	Body                 *Response `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResponseC) Reset()         { *m = ResponseC{} }
func (m *ResponseC) String() string { return proto.CompactTextString(m) }
func (*ResponseC) ProtoMessage()    {}
func (*ResponseC) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{6}
}

func (m *ResponseC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseC.Unmarshal(m, b)
}
func (m *ResponseC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseC.Marshal(b, m, deterministic)
}
func (m *ResponseC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseC.Merge(m, src)
}
func (m *ResponseC) XXX_Size() int {
	return xxx_messageInfo_ResponseC.Size(m)
}
func (m *ResponseC) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseC.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseC proto.InternalMessageInfo

func (m *ResponseC) GetBody() *Response {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From                 string       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string       `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Status               int64        `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Trace                *trace.Chain `protobuf:"bytes,10,opt,name=trace,proto3" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7bec9058cd15991, []int{7}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Response) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Response) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetTrace() *trace.Chain {
	if m != nil {
		return m.Trace
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestA)(nil), "messages.RequestA")
	proto.RegisterType((*RequestB)(nil), "messages.RequestB")
	proto.RegisterType((*RequestC)(nil), "messages.RequestC")
	proto.RegisterType((*Request)(nil), "messages.Request")
	proto.RegisterType((*ResponseA)(nil), "messages.ResponseA")
	proto.RegisterType((*ResponseB)(nil), "messages.ResponseB")
	proto.RegisterType((*ResponseC)(nil), "messages.ResponseC")
	proto.RegisterType((*Response)(nil), "messages.Response")
}

func init() { proto.RegisterFile("services/messages/message.proto", fileDescriptor_d7bec9058cd15991) }

var fileDescriptor_d7bec9058cd15991 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0xb3, 0xcb, 0x8a, 0x30, 0x1a, 0x13, 0xe6, 0x60, 0x1a, 0x2f, 0x92, 0x4d, 0x34, 0x9c,
	0x4a, 0x94, 0x27, 0x80, 0x7d, 0x83, 0xbe, 0x41, 0x61, 0x47, 0xed, 0x01, 0x8a, 0x9d, 0x62, 0xa2,
	0x4f, 0x6f, 0x98, 0x96, 0xb8, 0xde, 0x16, 0x2e, 0xcd, 0xf4, 0x9f, 0xff, 0x4b, 0x67, 0xfe, 0x14,
	0x1e, 0x99, 0xc2, 0x97, 0xdb, 0x10, 0xcf, 0xb7, 0xc4, 0x6c, 0xdf, 0xff, 0x0a, 0xbd, 0x0f, 0x3e,
	0x7a, 0x1c, 0x9d, 0xf4, 0x87, 0x49, 0x0c, 0x76, 0x43, 0x73, 0x39, 0x53, 0xb3, 0x7e, 0x81, 0x91,
	0xa1, 0xcf, 0x03, 0x71, 0x5c, 0xe2, 0x13, 0x54, 0x6b, 0xdf, 0x7e, 0xab, 0x62, 0x5a, 0xcc, 0x6e,
	0x5e, 0x27, 0xfa, 0xc4, 0xe9, 0xec, 0x30, 0xd2, 0xee, 0x20, 0xab, 0xf3, 0x91, 0xa6, 0x2f, 0x62,
	0xe1, 0x3a, 0x0b, 0x78, 0x07, 0xa5, 0x6b, 0xc5, 0x3f, 0x30, 0xa5, 0x6b, 0x11, 0xa1, 0x7a, 0x0b,
	0x7e, 0xab, 0xca, 0x69, 0x31, 0x1b, 0x1b, 0xa9, 0x8f, 0x9e, 0xe8, 0xd5, 0x40, 0x94, 0x32, 0x7a,
	0xac, 0xe1, 0x4a, 0xd6, 0x54, 0x20, 0xcf, 0xdc, 0xea, 0xb4, 0x74, 0xf3, 0x61, 0xdd, 0xce, 0xa4,
	0x56, 0xbd, 0x80, 0xb1, 0x21, 0xde, 0xfb, 0x1d, 0xd3, 0x12, 0x9f, 0xff, 0x8d, 0x85, 0xdd, 0xb1,
	0x92, 0x25, 0xcf, 0xd5, 0x81, 0x56, 0x97, 0x40, 0x4d, 0x6f, 0xe8, 0xe7, 0x18, 0x5a, 0x52, 0x2e,
	0x8a, 0xe0, 0x1e, 0x86, 0x1c, 0x6d, 0x3c, 0xb0, 0xaa, 0x84, 0xcb, 0xb7, 0x3e, 0xd1, 0xac, 0x87,
	0xf2, 0x3b, 0x16, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xd1, 0x6b, 0x3a, 0x5d, 0x02, 0x00,
	0x00,
}
