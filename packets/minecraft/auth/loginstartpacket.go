package auth

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type LoginStartPacket struct {
	PacketId int
	Name     string
}

func NewLoginStartPacket(name string) packets.Packet {
	//TODO: Enforce max len name
	return &LoginStartPacket{
		PacketId: 0,
		Name:     name,
	}
}

func (packet *LoginStartPacket) PreRead(_ *session.Session) {}

func (packet *LoginStartPacket) Read(packetId int, reader *io.PacketReader, _ int) packets.Packet {
	packet.PacketId = packetId
	packet.Name = reader.ReadString()

	return packet
}

func (packet *LoginStartPacket) PostRead(currentSession *session.Session) packets.Packet {
	currentSession.PlayerData.Username = packet.Name

	shouldKick := false // TODO
	if shouldKick {
		// send a disconnect packet for now
		disconnectPacket := NewDisconnectPacket()

		return disconnectPacket
	} else {
		encryptionRequestPacket := NewEncryptionRequestPacket()

		return encryptionRequestPacket
	}
}

func (packet *LoginStartPacket) PreWrite(_ *session.Session) {}

func (packet *LoginStartPacket) Write(writer *io.PacketWriter) {
	writer.WriteVarInt(packet.PacketId)
	writer.WriteString(packet.Name)
}

func (packet *LoginStartPacket) PostWrite(_ *session.Session) {}
