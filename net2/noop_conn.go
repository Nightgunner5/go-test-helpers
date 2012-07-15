package net2

import (
	"net"
	"time"
)

type noop_conn int

// Compile-time interface check
var _ net.Conn = noop_conn(0)

func NoopConn() net.Conn {
	return noop_conn(0)
}

func (noop_conn) Close() error {
	return nil
}

func (noop_conn) LocalAddr() net.Addr {
	return pipeAddr(0)
}

func (noop_conn) RemoteAddr() net.Addr {
	return pipeAddr(0)
}

func (noop_conn) Read([]byte) (int, error) {
	return 0, nil
}

func (noop_conn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (noop_conn) SetDeadline(time.Time) error {
	return nil
}

func (noop_conn) SetReadDeadline(time.Time) error {
	return nil
}

func (noop_conn) SetWriteDeadline(time.Time) error {
	return nil
}
