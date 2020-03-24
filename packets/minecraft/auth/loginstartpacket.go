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

func NewLoginStartPacket(name string) *LoginStartPacket {
	//TODO: Enforce max len name
	packet := new(LoginStartPacket)
	packet.PacketId = 0
	packet.Name = name

	return packet
}

func (packet *LoginStartPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId
	packet.Name = reader.ReadString(reader.ReadVarInt())

	return packet
}

func (packet *LoginStartPacket) Handle(currentSession *session.Session) {
	// send a disconnect packet for now
	disconnectPacket := NewDisconnectPacket()
	disconnectPacket.Write(currentSession)

	// close connection
	currentSession.Close()
}

func (packet *LoginStartPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteString(packet.Name)
	currentSession.Writer.Flush()
}
