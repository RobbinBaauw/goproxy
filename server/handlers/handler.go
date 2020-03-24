package handlers

import (
	"github.com/timanema/goproxy/packets/io"
)

type PacketHandler interface {
	Handle(packetReader *io.PacketReader, packetId int)
}
