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

func (packet *LoginStartPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	packet.Name = reader.ReadString()

	return packet
}

func (packet *LoginStartPacket) HandleRead(currentSession *session.Session) packets.Packet {
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

func (packet *LoginStartPacket) HandleWrite(currentSession *session.Session) {
	panic("implement me")
}

func (packet *LoginStartPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteString(packet.Name)
	currentSession.Writer.Flush(nil)
}
