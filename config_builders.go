package sonic

import (
	"crypto/tls"
	"net"

	"github.com/AlexBlayE/sonic/dialers"
	"github.com/AlexBlayE/sonic/interfaces"
)

// port format is ":1234"
func MakeTcpListenerDialer(port string) (net.Listener, interfaces.Dialer) {
	listener, err := net.Listen("tcp", port)
	dialer := dialers.NewTcpDial()

	if err != nil {
		return nil, nil
	}

	return listener, dialer
}

func MakeTlsListenerDialer(port string, config *tls.Config) (net.Listener, interfaces.Dialer) {
	listener, err := tls.Listen("tcp", port, config)
	dialer := dialers.NewTlsDial(config)

	if err != nil {
		return nil, nil
	}

	return listener, dialer
}
