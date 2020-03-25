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

func NewLoginSuccessPacket(name string, uuid string) packets.Packet {
	return &LoginSuccessPacket{
		PacketId: 2,
		Name:     name,
		UUID:     uuid,
	}
}

func (packet *LoginSuccessPacket) PreRead(_ *session.Session) {}

func (packet *LoginSuccessPacket) Read(_ int, _ *io.PacketReader, _ int) packets.Packet {
	// TODO
	log.Panic("TODO!")
	return nil
}

func (packet *LoginSuccessPacket) PostRead(_ *session.Session) packets.Packet {
	return nil
}

func (packet *LoginSuccessPacket) PreWrite(currentSession *session.Session) {
	currentSession.CurrentState = session.Play
}

func (packet *LoginSuccessPacket) Write(writer *io.PacketWriter) {
	writer.WriteVarInt(packet.PacketId)
	writer.WriteString(packet.UUID)
	writer.WriteString(packet.Name)
}

func (packet *LoginSuccessPacket) PostWrite(_ *session.Session) {}
