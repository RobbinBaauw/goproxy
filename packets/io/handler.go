package io

type PacketHandler interface {
	Handle(packetReader *PacketReader, packetId int)
}
