package sonic

import (
	"crypto/tls"
	"testing"

	"github.com/AlexBlayE/sonic/codec"
)

func TestTcpConnection(t *testing.T) {
	go func() {
		reciever := NewReciever[string](codec.NewGobCodec(), 10)
		sender := NewSender[string](codec.NewGobCodec())

		listener, dialer := MakeTcpListenerDialer(":8080")

		manager := NewManager(reciever, sender, listener, dialer)

		for msg := range manager.Recv() {
			if msg != "Test 1" {
				t.Fail()
				return
			}

			manager.Send("Test 2", "localhost:3000")
		}
	}()

	go func() {
		reciever := NewReciever[string](codec.NewGobCodec(), 10)
		sender := NewSender[string](codec.NewGobCodec())

		listener, dialer := MakeTcpListenerDialer(":3000")

		manager := NewManager(reciever, sender, listener, dialer)

		manager.Send("Test 1", "localhost:8080")

		for msg := range manager.Recv() {
			if msg != "Test 2" {
				t.Fail()
				return
			}
		}
	}()
}

func TestTlsConnection(t *testing.T) {
	config := &tls.Config{}

	go func() {
		reciever := NewReciever[string](codec.NewGobCodec(), 10)
		sender := NewSender[string](codec.NewGobCodec())

		listener, dialer := MakeTlsListenerDialer(":8080", config.Clone())

		manager := NewManager(reciever, sender, listener, dialer)

		for msg := range manager.Recv() {
			if msg != "Test 1" {
				t.Fail()
				return
			}

			manager.Send("Test 2", "localhost:3000")
		}
	}()

	go func() {
		reciever := NewReciever[string](codec.NewGobCodec(), 10)
		sender := NewSender[string](codec.NewGobCodec())

		listener, dialer := MakeTlsListenerDialer(":3000", config.Clone())

		manager := NewManager(reciever, sender, listener, dialer)

		manager.Send("Test 1", "localhost:8080")

		for msg := range manager.Recv() {
			if msg != "Test 2" {
				t.Fail()
				return
			}
		}
	}()
}
