package pangolin

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Server struct {
	addr   string
	server *net.TCPListener
}

func NewServer(addr string) (*Server, error) {
	svr, err := listen(addr)
	if err != nil {
		return nil, err
	}
	return &Server{
		server: svr,
	}, nil
}

func (s *Server) Run() {
	for {
		conn, err := s.server.Accept()
		if err != nil {
			continue
		}
		dec := gob.NewDecoder(conn)
		message := &Message{}
		err = dec.Decode(message)
		if err != nil {
			continue
		}
		fmt.Println(message.header.Get("type"))
		fmt.Println(message.body.Get("name").(string))
		encoder :=gob.NewEncoder(conn)
		h := NewDefaultHeader()
		h.Set("type", "test")
		b := NewDefaultBody()
		b.Set("name", "ningqing")
		m := NewMessage(h, b)
		encoder.Encode(m)
	}
}
