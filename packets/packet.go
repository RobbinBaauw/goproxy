package packets

import (
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type Packet interface {
	Read(packetId int, reader *io.PacketReader) Packet
	HandleRead(currentSession *session.Session) Packet
	Write(currentSession *session.Session)
	HandlePreWrite(currentSession *session.Session)
	HandleWrite(currentSession *session.Session)
}
