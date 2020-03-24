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

func (packet *HandshakePacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	packet.ProtocolVersion = reader.ReadVarInt()
	packet.ServerAddrLen = reader.ReadVarInt()
	packet.ServerAddr = reader.ReadString(packet.ServerAddrLen)
	packet.ServerPort = reader.ReadUnsignedShort()
	packet.NextState = reader.ReadVarInt()

	return packet
}
