package framedtcp

import (
	"net"
)

type Framed interface {
	Conn() net.Conn
	Receive() ([]byte, error)
	Send([]byte) error
}
