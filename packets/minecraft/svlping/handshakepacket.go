package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type HandshakePacket struct {
	PacketId        int
	ProtocolVersion int
	ServerAddr      string
	ServerPort      uint16
	NextState       int
}

func (packet *HandshakePacket) PreRead(_ *session.Session) {}

func (packet *HandshakePacket) Read(packetId int, reader *io.PacketReader, _ int) packets.Packet {
	packet.PacketId = packetId
	packet.ProtocolVersion = reader.ReadVarInt()
	packet.ServerAddr = reader.ReadString()
	packet.ServerPort = reader.ReadUnsignedShort()
	packet.NextState = reader.ReadVarInt()

	return packet
}

func (packet *HandshakePacket) PostRead(currentSession *session.Session) packets.Packet {
	// update state
	currentSession.CurrentState = packet.NextState

	return nil
}

func (packet *HandshakePacket) PreWrite(_ *session.Session) {}

func (packet *HandshakePacket) Write(writer *io.PacketWriter) {}

func (packet *HandshakePacket) PostWrite(_ *session.Session) {}
