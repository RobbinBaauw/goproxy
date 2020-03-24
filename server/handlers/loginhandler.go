package handlers

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/packets/minecraft/auth"
	"log"
)

func HandleLogin(packetReader *io.PacketReader, packetId int) packets.Packet {
	switch packetId {
	case 0:
		return new(auth.LoginStartPacket).Read(packetId, packetReader)
	case 1:
		return new(auth.EncryptionResponsePacket).Read(packetId, packetReader)
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
