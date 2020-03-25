package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type RequestPacket struct {
	PacketId int
}

func (packet *RequestPacket) PreRead(_ *session.Session) {}

func (packet *RequestPacket) Read(packetId int, _ *io.PacketReader, _ int) packets.Packet {
	packet.PacketId = packetId
	return packet
}

func (packet *RequestPacket) PostRead(_ *session.Session) packets.Packet {
	// send response packet
	responsePacket := NewResponsePacket()

	return responsePacket
}

func (packet *RequestPacket) PreWrite(_ *session.Session) {}

func (packet *RequestPacket) Write(writer *io.PacketWriter) {}

func (packet *RequestPacket) PostWrite(_ *session.Session) {}
