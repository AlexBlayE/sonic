package sonic

import (
	"net"

	"github.com/AlexBlayE/sonic/interfaces"
)

type Sender[T any] struct {
	codec interfaces.Codec
}

func NewSender[T any](codec interfaces.Codec) *Sender[T] {
	return &Sender[T]{
		codec: codec,
	}
}

func (s *Sender[T]) Send(msg T, con net.Conn) error {
	encoded, err := s.codec.Encode(msg)

	if err != nil {
		return err
	}

	_, err = con.Write(encoded)

	return err
}
