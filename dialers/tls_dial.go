package dialers

import (
	"crypto/tls"
	"net"
)

type TlsDial struct {
	config *tls.Config
}

func NewTlsDial(tlsConfig *tls.Config) *TlsDial {
	return &TlsDial{config: tlsConfig}
}

func (td *TlsDial) Dial(address string) (net.Conn, error) {
	return tls.Dial("tcp", address, nil)
}
