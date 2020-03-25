package packets

import (
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
)

type Packet interface {
	PreRead(currentSession *session.Session)
	Read(packetId int, reader *io.PacketReader) Packet
	PostRead(currentSession *session.Session) Packet
	PreWrite(currentSession *session.Session)
	Write(currentSession *session.Session)
	PostWrite(currentSession *session.Session)
}
