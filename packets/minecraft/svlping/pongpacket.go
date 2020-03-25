package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type PongPacket struct {
	PacketId int
	Payload  int64
}

func NewPongPacket(payload int64) packets.Packet {
	return &PongPacket{
		PacketId: 1,
		Payload:  payload,
	}
}

func (packet *PongPacket) HandleRead(currentSession *session.Session) {
	panic("implement me")
}

func (packet *PongPacket) HandleWrite(currentSession *session.Session) {
	panic("implement me")
}

func (packet *PongPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	panic("implement me")
}

func (packet *PongPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteLong(packet.Payload)
	currentSession.Writer.Flush(nil)
}
