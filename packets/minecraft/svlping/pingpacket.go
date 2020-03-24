package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type PingPacket struct {
	PacketId int
	Payload  int64
}

func (packet *PingPacket) Read(reader *io.PacketReader) packets.Packet {
	packet.Payload = reader.ReadLong()

	return packet
}

func (packet *PingPacket) Handle(currentSession *session.Session) {
	// send pong packet
	pongPacket := NewPongPacket(packet.Payload)
	pongPacket.Write(currentSession)

	// close connection
	currentSession.Close()
}
func (packet *PingPacket) Write(currentSession *session.Session) {

}
