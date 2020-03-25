package play

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type GenericPacket struct {
	PacketId int
	Payload  []byte
}

func NewGenericPacket() packets.Packet {
	return new(GenericPacket)
}

func (packet GenericPacket) PreRead(_ *session.Session) {}

func (packet GenericPacket) Read(packetId int, reader *io.PacketReader, length int) packets.Packet {
	packet.PacketId = packetId
	packet.Payload = reader.ReadBytes(length)

	return packet
}

func (packet GenericPacket) PostRead(currentSession *session.Session) packets.Packet {
	// TODO: Forward to real server
	return nil
}

func (packet GenericPacket) PreWrite(_ *session.Session) {}

func (packet GenericPacket) Write(writer *io.PacketWriter) {
	writer.WriteVarInt(packet.PacketId)
	writer.WriteBytes(packet.Payload)
}

func (packet GenericPacket) PostWrite(_ *session.Session) {}
