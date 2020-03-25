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

func (packet *PingPacket) PreRead(_ *session.Session) {}

func (packet *PingPacket) Read(packetId int, reader *io.PacketReader, _ int) packets.Packet {
	packet.PacketId = packetId
	packet.Payload = reader.ReadLong()

	return packet
}

func (packet *PingPacket) PostRead(_ *session.Session) packets.Packet {
	// send pong packet
	pongPacket := NewPongPacket(packet.Payload)

	return pongPacket
}

func (packet *PingPacket) PreWrite(_ *session.Session) {}

func (packet *PingPacket) Write(writer *io.PacketWriter) {
	log.Panic("Proxy should never send a ping packet")
}

func (packet *PingPacket) PostWrite(_ *session.Session) {}
