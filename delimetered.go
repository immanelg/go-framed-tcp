package framedtcp

import (
	"bufio"
	"bytes"
	"errors"
	"net"
)

type DelimeterFramed struct {
	delimeter []byte
	conn      net.Conn
	scanner   bufio.Scanner
}

func NewDelimeterFramed(conn net.Conn, delimiter []byte) DelimeterFramed {
	c := DelimeterFramed{
		conn:    conn,
		scanner: *bufio.NewScanner(conn),
	}
	c.scanner.Split(func(data []byte, eof bool) (advance int, token []byte, err error) {
		if i := bytes.Index(data, delimiter); i >= 0 {
			return i + len(delimiter), data[0:i], nil
		}
		if eof {
			return len(data), data, nil
		}
		return 0, nil, nil
	})
	return c
}

func (self *DelimeterFramed) Conn() net.Conn {
	return self.conn
}

func (self *DelimeterFramed) Receive() ([]byte, error) {
	if !self.scanner.Scan() {
		return nil, errors.New("eof")
	}
	if err := self.scanner.Err(); err != nil {
		return nil, err
	}
	data := self.scanner.Bytes()
	return data, nil
}

func (self *DelimeterFramed) Send(msg []byte) error {
	_, err := self.conn.Write(append(msg, self.delimeter...))
	return err
}
