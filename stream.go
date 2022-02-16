package pangolin

import (
	"encoding/gob"
	"io"
	"net"
)

var _ io.ReadWriter = (*Stream)(nil)

type Stream struct {
	conn *net.TCPConn
}

func NewStream(conn *net.TCPConn) *Stream {
	return &Stream{
		conn: conn,
	}
}

func (s *Stream) Write(b []byte) (int, error) {
	return s.conn.Write(b)
}

func (s *Stream) Read(b []byte) (int, error) {
	return s.conn.Read(b)
}

func (s *Stream) ReadMessage()(*Message,error) {
	dec := gob.NewDecoder(s.conn)
	message := &Message{}
	err := dec.Decode(message)
	if err != nil {
		return nil,err
	}
	return message,nil
}

var _ io.Writer = (*Response)(nil)

type Response struct {
	conn net.Conn
}

func NewResponse(conn net.Conn) *Response {
	return &Response{
		conn: conn,
	}
}

// func (s ) ReadMessage()(*Message,error) {
// 	dec := gob.NewDecoder(s.conn)
// 	message := &Message{}
// 	err := dec.Decode(message)
// 	if err != nil {
// 		return nil,err
// 	}
// 	return message,nil
// }

func (r *Response) Write(b []byte) (int, error) {
	return r.conn.Write(b)
}

type Request struct {
}
