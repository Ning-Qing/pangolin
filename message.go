package pangolin

import (
	"crypto/sha1"
	"fmt"
	"unsafe"
)

// type Value interface {}


type Header interface {
	Get(key string) string
	Set(key,value string)
	IsNewConnection() bool
	GetMethod() string
}

type Body interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

type Message struct {
	header Header
	body   Body
	hash   []byte
}

func NewMessage(header Header, body Body) *Message {
	return &Message{
		header: header,
		body:   body,
	}
}

func (m *Message) Summary() {
	h := sha1.New()
	data := m.convert()
	m.hash = h.Sum(data)
}

func (m *Message) Header() Header {
	return m.header
}

func (m *Message) Body() Body {
	return m.body
}

type toStructToBytes struct {
	header Header
	body   Body
}

type sliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func (m *Message) convert() []byte {
	s := &toStructToBytes{
		header: m.header,
		body:   m.body,
	}
	len := unsafe.Sizeof(*s)
	mock := &sliceMock{
		addr: uintptr(unsafe.Pointer(s)),
		cap:  int(len),
		len:  int(len),
	}
	return *(*[]byte)(unsafe.Pointer(mock))
}

func (m *Message) Check() error {
	h := sha1.New()

	data := m.convert()
	if string(h.Sum(data)) != string(m.hash) {
		return fmt.Errorf("checkdata may be tampered with")
	}

	return nil
}