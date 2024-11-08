package sonic

import (
	"net"

	"github.com/AlexBlayE/sonic/interfaces"
)

type Reciever[T any] struct {
	codec interfaces.Codec

	ch chan T
}

func NewReciever[T any](codec interfaces.Codec, channelSize int) *Reciever[T] {
	return &Reciever[T]{
		codec: codec,
		ch:    make(chan T, channelSize),
	}
}

func (r *Reciever[T]) ReadEvent(con net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := con.Read(buffer)

		if err != nil {
			continue // TODO
		}

		var it T
		err = r.codec.Decode(buffer[:n], &it)

		if err != nil {
			continue // TODO
		}

		r.ch <- it
	}
}

func (r *Reciever[T]) GetChannel() <-chan T {
	return r.ch
}
