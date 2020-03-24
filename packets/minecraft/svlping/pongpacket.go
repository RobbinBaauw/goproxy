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

func NewPongPacket(payload int64) *PongPacket {
	packet := new(PongPacket)
	packet.PacketId = 1
	packet.Payload = payload

	return packet
}

func (packet *PongPacket) Read(reader *io.PacketReader) packets.Packet {
	packet.Payload = reader.ReadLong()

	return packet
}

func (packet *PongPacket) Handle(currentSession *session.Session) {

}
func (packet *PongPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteLong(packet.Payload)
	currentSession.Writer.Flush()
}
