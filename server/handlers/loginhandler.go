package handlers

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/minecraft/auth"
	"log"
)

func HandleLogin(packetId int) packets.Packet {
	switch packetId {
	case 0:
		return new(auth.LoginStartPacket)
	case 1:
		return new(auth.EncryptionResponsePacket)
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
