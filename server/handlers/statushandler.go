package handlers

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/packets/minecraft/svlping"
	"log"
)

func HandleStatus(packetReader *io.PacketReader, packetId int) packets.Packet {
	switch packetId {
	case 0:
		return new(svlping.RequestPacket).Read(packetId, packetReader)
	case 1:
		return new(svlping.PingPacket).Read(packetId, packetReader)
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
