package interfaces

import "net"

type Dialer interface {
	Dial(address string) (net.Conn, error)
}
