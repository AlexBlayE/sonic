package sonic

import (
	"net"
	"sync"

	"github.com/AlexBlayE/sonic/interfaces"
)

type ConnectionManager[T any] struct {
	reciever interfaces.Reciever[T]
	sender   interfaces.Sender[T]

	dialer interfaces.Dialer

	listener net.Listener

	connPool sync.Map
}

func NewManager[T any](
	reciever interfaces.Reciever[T],
	sender interfaces.Sender[T],
	listener net.Listener,
	dialer interfaces.Dialer,
) *ConnectionManager[T] {
	return &ConnectionManager[T]{
		reciever: reciever,
		sender:   sender,
		dialer:   dialer,
		listener: listener,
	}
}

func (m *ConnectionManager[T]) Recv() <-chan T {
	go func() {
		for {
			con, err := m.listener.Accept()

			if err != nil {
				continue
			}

			m.connPool.Store(con.RemoteAddr().String(), con)

			go m.reciever.ReadEvent(con)
		}
	}()

	return m.reciever.GetChannel()
}

// address is ip:port
func (m *ConnectionManager[T]) Send(msg T, address string) error {
	var con net.Conn

	val, ok := m.connPool.Load(address)

	if !ok {
		con, err := m.dialer.Dial(address)

		if err != nil {
			return err
		}

		m.sender.Send(msg, con)
		m.connPool.Store(address, con)

		return nil
	}

	con = val.(net.Conn)
	err := m.sender.Send(msg, con)

	return err
}

func (m *ConnectionManager[T]) ListAddresses() []string {
	list := []string{}

	m.connPool.Range(func(key, value any) bool {
		list = append(list, key.(string))
		return true
	})

	return list
}
