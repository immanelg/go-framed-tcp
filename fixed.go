package framedtcp

import (
	"io"
	"net"
)

type FixedLengthFramed struct {
	len  int
	conn net.Conn
}

func NewFixedLengthFramed(conn net.Conn, len int) FixedLengthFramed {
	c := FixedLengthFramed{
		len:  len,
		conn: conn,
	}
	return c
}

func (self *FixedLengthFramed) Conn() net.Conn {
	return self.conn
}

func (self *FixedLengthFramed) Close() error {
	return self.conn.Close()
}

func (self *FixedLengthFramed) Receive() ([]byte, error) {
	data := make([]byte, 0, self.len)
	_, err := io.ReadFull(self.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *FixedLengthFramed) Send(msg []byte) error {
	if self.len != len(msg) {
		panic("wrong length of message")
	}
	_, err := self.conn.Write(msg)
	return err
}
