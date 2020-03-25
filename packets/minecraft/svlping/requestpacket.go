package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type RequestPacket struct {
	PacketId int
}

func (packet *RequestPacket) HandleRead(currentSession *session.Session) {
	// send response packet
	responsePacket := NewResponsePacket()
	responsePacket.Write(currentSession)
}

func (packet *RequestPacket) HandleWrite(currentSession *session.Session) {
	panic("implement me")
}

func (packet *RequestPacket) Write(currentSession *session.Session) {
	panic("implement me")
}

func (packet *RequestPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	return packet
}
