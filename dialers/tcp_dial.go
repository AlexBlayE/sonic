package dialers

import "net"

type TcpDial struct {
}

func NewTcpDial() *TcpDial {
	return &TcpDial{}
}

func (td *TcpDial) Dial(address string) (net.Conn, error) {
	return net.Dial("tcp", address)
}
