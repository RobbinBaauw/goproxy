package svlping

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
	"log"
)

type PingPacket struct {
	PacketId int
	Payload  int64
}

func (packet *PingPacket) HandleRead(currentSession *session.Session) {
	// send pong packet
	pongPacket := NewPongPacket(packet.Payload)
	pongPacket.Write(currentSession)

	// close connection
	currentSession.Close()
}

func (packet *PingPacket) HandleWrite(currentSession *session.Session) {
	panic("implement me")
}

func (packet *PingPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	packet.Payload = reader.ReadLong()

	return packet
}

func (packet *PingPacket) Write(currentSession *session.Session) {
	log.Panic("Proxy should never send a ping packet")
}
