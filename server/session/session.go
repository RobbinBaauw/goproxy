package session

import (
	"github.com/timanema/goproxy/packets/io"
	"net"
)

const (
	Handshaking = iota
	Status
	Login
	Play
)

type Session struct {
	Connection   *net.Conn
	CurrentState int
	Reader       *io.PacketReader
	Writer       *io.PacketWriter
}

func NewSession(conn *net.Conn) *Session {
	session := new(Session)
	session.Connection = conn
	session.CurrentState = Handshaking

	return session
}
