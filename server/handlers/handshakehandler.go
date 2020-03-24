package handlers

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/packets/minecraft/svlping"
	"log"
)

func HandleHandshake(packetReader *io.PacketReader, packetId int) packets.Packet {
	switch packetId {
	case 0:
		return new(svlping.HandshakePacket).Read(packetId, packetReader)
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
