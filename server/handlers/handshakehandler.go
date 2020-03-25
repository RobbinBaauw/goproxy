package handlers

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/minecraft/svlping"
	"log"
)

func HandleHandshake(packetId int) packets.Packet {
	switch packetId {
	case 0:
		return new(svlping.HandshakePacket)
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
