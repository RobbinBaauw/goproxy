package session

import (
	"github.com/timanema/goproxy/packets/io"
	"log"
	"net"
)

const (
	Handshaking = iota
	Status
	Login
	Play
)

type Session struct {
	Connection       *net.Conn
	CurrentState     int
	Reader           *io.PacketReader
	Writer           *io.PacketWriter
	ConnectionClosed bool
	SharedSecret     []byte
	PlayerData       PlayerData
}

type PlayerData struct {
	Username string
	UUID     string
}

func NewSession(conn *net.Conn) *Session {
	session := new(Session)
	session.Connection = conn
	session.CurrentState = Handshaking

	return session
}

func (session *Session) Close() {
	log.Print("Closed connection from: ", (*session.Connection).RemoteAddr().String())
	session.ConnectionClosed = true
	_ = (*session.Connection).Close()
}
