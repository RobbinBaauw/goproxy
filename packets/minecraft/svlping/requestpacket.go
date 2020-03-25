package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type RequestPacket struct {
	PacketId int
}

func (packet *RequestPacket) PreRead(currentSession *session.Session) {}

func (packet *RequestPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	return packet
}

func (packet *RequestPacket) PostRead(currentSession *session.Session) packets.Packet {
	// send response packet
	responsePacket := NewResponsePacket()

	return responsePacket
}

func (packet *RequestPacket) PreWrite(currentSession *session.Session) {}

func (packet *RequestPacket) Write(currentSession *session.Session) {}

func (packet *RequestPacket) PostWrite(currentSession *session.Session) {}
