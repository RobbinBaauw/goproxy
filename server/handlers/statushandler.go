package handlers

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/minecraft/svlping"
	"log"
)

func HandleStatus(packetId int) packets.Packet {
	switch packetId {
	case 0:
		return new(svlping.RequestPacket)
	case 1:
		return new(svlping.PingPacket)
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
