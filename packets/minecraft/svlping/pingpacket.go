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

func (packet *PingPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	packet.Payload = reader.ReadLong()

	return packet
}

func (packet *PingPacket) HandleRead(currentSession *session.Session) packets.Packet {
	// send pong packet
	pongPacket := NewPongPacket(packet.Payload)

	return pongPacket
}

func (packet *PingPacket) Write(currentSession *session.Session) {
	log.Panic("Proxy should never send a ping packet")
}

func (packet *PingPacket) HandlePreWrite(currentSession *session.Session) {}

func (packet *PingPacket) HandleWrite(currentSession *session.Session) {}
