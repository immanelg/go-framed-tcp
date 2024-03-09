package framedtcp

import (
	"encoding/binary"
	"io"
	"net"
)

type HeaderLengthFramed struct {
	conn net.Conn
}

func (self *HeaderLengthFramed) Conn() net.Conn {
	return self.conn
}

func (self *HeaderLengthFramed) Receive() ([]byte, error) {
	header := make([]byte, 0, 4)
	if _, err := io.ReadFull(self.conn, header); err != nil {
		return nil, err
	}

	len := binary.BigEndian.Uint32(header)
	data := make([]byte, 0, len)
	if _, err := io.ReadFull(self.conn, data); err != nil {
		return nil, err
	}

	return data[4:], nil
}

func (self *HeaderLengthFramed) Send(msg []byte) error {
	msg = append([]byte{0, 0, 0, 0}, msg...)
	binary.BigEndian.PutUint32(msg[:4], uint32(len(msg)-4))
	_, err := self.conn.Write(msg)
	return err
}
