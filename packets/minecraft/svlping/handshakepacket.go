package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type HandshakePacket struct {
	PacketId        int
	ProtocolVersion int
	ServerAddrLen   int
	ServerAddr      string
	ServerPort      uint16
	NextState       int
}

func (packet *HandshakePacket) Write(currentSession *session.Session) {
	panic("implement me")
}

func (packet *HandshakePacket) Handle(currentSession *session.Session) {
	// update state
	currentSession.CurrentState = packet.NextState
}

func (packet *HandshakePacket) Read(reader *io.PacketReader) packets.Packet {
	protocolVersion := reader.ReadVarInt()
	serverAddrLen := reader.ReadVarInt()
	serverAddr := reader.ReadString(serverAddrLen)
	serverPort := reader.ReadUnsignedShort()
	nextState := reader.ReadVarInt()

	packet.ProtocolVersion = protocolVersion
	packet.ServerAddrLen = serverAddrLen
	packet.ServerAddr = serverAddr
	packet.ServerPort = serverPort
	packet.NextState = nextState

	return packet
}
