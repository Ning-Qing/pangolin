package pangolin

import (
	"fmt"
	"io"
	"net"
)

func dial(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return nil, fmt.Errorf("addr parsing failed: %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("dial tcp failed: %v", err)
	}
	return conn, nil
}

func listen(addr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return nil, fmt.Errorf("addr parsing failed: %v", err)
	}
	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("dial tcp failed: %v", err)
	}
	return listen, nil
}

func joinConn(local,remote *net.TCPConn){
	defer local.Close()
	defer remote.Close()
	_,err := io.Copy(local,remote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func join2Conn(local,remoto *net.TCPConn){
	go joinConn(local,remoto)
	go joinConn(remoto,local)
}