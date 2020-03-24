package auth

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
	"log"
)

type LoginSuccessPacket struct {
	PacketId int
	Name     string
	UUID     string
}

func NewLoginSuccessPacket(name string, uuid string) *LoginSuccessPacket {
	return &LoginSuccessPacket{
		PacketId: 0,
		Name:     name,
		UUID:     uuid,
	}
}

func (packet *LoginSuccessPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	// TODO
	log.Panic("TODO!")
	return nil
}

func (packet *LoginSuccessPacket) Handle(currentSession *session.Session) {
}

func (packet *LoginSuccessPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteString(packet.Name)
	currentSession.Writer.WriteString(packet.UUID)
	currentSession.Writer.Flush(&currentSession.SharedSecret)
}
