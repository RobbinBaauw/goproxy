package auth

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type DisconnectPacket struct {
	PacketId int
	Reason   string
}

func NewDisconnectPacket() packets.Packet {
	packet := new(DisconnectPacket)
	packet.PacketId = 0
	packet.Reason = "{\"text\": \"yeet\", \"bold\": \"true\", \"color\": \"gold\"}"

	return packet
}

func (packet *DisconnectPacket) PreRead(_ *session.Session) {}

func (packet *DisconnectPacket) Read(packetId int, reader *io.PacketReader, _ int) packets.Packet {
	packet.PacketId = packetId
	packet.Reason = reader.ReadString()

	return packet
}

func (packet *DisconnectPacket) PostRead(_ *session.Session) packets.Packet {
	return nil
}

func (packet *DisconnectPacket) PreWrite(_ *session.Session) {}

func (packet *DisconnectPacket) Write(writer *io.PacketWriter) {
	writer.WriteVarInt(packet.PacketId)
	writer.WriteString(packet.Reason)
}

func (packet *DisconnectPacket) PostWrite(currentSession *session.Session) {
	currentSession.Close()
}
