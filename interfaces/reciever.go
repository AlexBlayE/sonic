package interfaces

import "net"

type Reciever[T any] interface {
	ReadEvent(net.Conn)
	GetChannel() <-chan T
}
