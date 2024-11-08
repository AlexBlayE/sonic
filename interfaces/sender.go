package interfaces

import "net"

type Sender[T any] interface {
	Send(msg T, con net.Conn) error
}
