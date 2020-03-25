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

func (packet *PongPacket) PreRead(_ *session.Session) {}

func (packet *PongPacket) Read(_ int, _ *io.PacketReader, _ int) packets.Packet {
	return nil
}

func (packet *PongPacket) PostRead(_ *session.Session) packets.Packet {
	return nil
}

func (packet *PongPacket) PreWrite(_ *session.Session) {}

func (packet *PongPacket) Write(writer *io.PacketWriter) {
	writer.WriteVarInt(packet.PacketId)
	writer.WriteLong(packet.Payload)
}

func (packet *PongPacket) PostWrite(currentSession *session.Session) {
	currentSession.Close()
}
