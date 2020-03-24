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

func NewDisconnectPacket() *DisconnectPacket {
	packet := new(DisconnectPacket)
	packet.PacketId = 0
	packet.Reason = "{\"text\": \"yeet\", \"bold\": \"true\"}"

	return packet
}

func (packet *DisconnectPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	packet.Reason = reader.ReadString(reader.ReadVarInt())

	return packet
}

func (packet *DisconnectPacket) Handle(currentSession *session.Session) {
	// TODO
}

func (packet *DisconnectPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteString(packet.Reason)
	currentSession.Writer.Flush()
}
