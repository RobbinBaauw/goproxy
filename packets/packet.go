package packets

import (
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type Packet interface {
	Read(packetId int, reader *io.PacketReader) Packet
	Handle(currentSession *session.Session)
	Write(currentSession *session.Session)
}