package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type RequestPacket struct {
	PacketId int
}

func (packet *RequestPacket) Write(currentSession *session.Session) {
	panic("implement me")
}

func (packet *RequestPacket) Handle(currentSession *session.Session) {
	// send response packet
	responsePacket := NewResponsePacket()
	responsePacket.Write(currentSession)
}

func (packet *RequestPacket) Read(reader *io.PacketReader) packets.Packet {
	return packet
}
