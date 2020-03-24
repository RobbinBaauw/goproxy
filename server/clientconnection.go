package server

import (
	"fmt"
	"github.com/finitum/goproxy/packets"
	"net"
	"strconv"
)

type ClientSession struct {
	Conn   net.Conn
	Reader *packets.ByteStreamReader
	State  int
	PlayerData PlayerData
}

type PlayerData struct {
	Username string
}

const (
	StateHandshaking = iota
	StateStatus
	StateLogin
	StatePlay
)

func HandleConnection(session *ClientSession) {
	session.Reader.ReadLength()

	packetId := packets.ReadVarInt(session.Reader)

	fmt.Println("PacketId: ", strconv.Itoa(packetId))

	if session.State == StateHandshaking {
		HandleHandshakeState(packetId, session)
	} else if session.State == StateStatus {
		HandleStatusState(packetId, session)
	} else if session.State == StateLogin {
		HandleLoginState(packetId, session)
	}
}

