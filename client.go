package pangolin

import (
	"context"
	"encoding/gob"
	"fmt"
	"net"
)

// Client
// 客户端
type Client struct {
	RemoteServerAddr  string
	RemoteControlAddr string
	LocalServerAddr   string
	router            *Router
}

func NewClient(remoteServerAddr, remoteControlAddr, localServerAddr string, router *Router) (*Client, error) {
	conn, err := dial(remoteControlAddr)
	if err != nil {
		return nil, fmt.Errorf("dial %s: %s", remoteControlAddr, err)
	}
	c := &Client{
		RemoteServerAddr:  remoteServerAddr,
		RemoteControlAddr: remoteControlAddr,
		LocalServerAddr:   localServerAddr,
		router:            router,
	}
	go c.runController(conn)
	listen, err := listen(localServerAddr)
	if err != nil {
		return nil, fmt.Errorf("listen %s: %s", localServerAddr, err)
	}
	go c.runServer(listen)
	return c, nil
}

func (c *Client) runServer(listen *net.TCPListener) {
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		dec := gob.NewDecoder(conn)
		message := &Message{}
		err = dec.Decode(message)
		if err != nil {
			continue
		}
		ctx := &PangolinCtx{
			ctx:    context.Background(),
			header: message.Header(),
			body:   message.Body(),
			w:      conn,
		}
		go c.router.run(message.header.GetMethod())(ctx)
	}
}

func (c *Client) runController(conn *net.TCPConn) {
	stream := NewStream(conn)
	for {
		message, err := stream.ReadMessage()
		if err != nil {
			continue
		}
		if message.Header().IsNewConnection() {
			go c.connetLocalAndRemote()
		}
	}
}

func (c *Client) connetLocalAndRemote() {
	local, err := dial(c.LocalServerAddr)
	if err != nil {
		local.Close()
		fmt.Println(err)
	}
	remote, err := dial(c.RemoteServerAddr)
	if err != nil {
		remote.Close()
		fmt.Println(err)
	}
	join2Conn(local, remote)
}
