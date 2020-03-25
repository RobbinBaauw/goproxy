package handlers

import (
	"github.com/timanema/goproxy/packets"
	"log"
)

func HandlePlay(packetId int) packets.Packet {
	switch packetId {
	default:
		log.Panic("Unknown packet id: ", packetId)
		return nil
	}
}
