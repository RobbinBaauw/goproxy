package auth

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
	"github.com/timanema/goproxy/util/encryption"
	"log"
)

type EncryptionRequestPacket struct {
	PacketId          int
	ServerId          string
	PublicKeyLength   int
	PublicKey         []byte
	VerifyTokenLength int
	VerifyToken       []byte
}

func NewEncryptionRequestPacket() packets.Packet {
	return &EncryptionRequestPacket{
		PacketId:          1,
		ServerId:          "",
		PublicKeyLength:   len(encryption.EncryptionDataInstance.DERPubkey),
		PublicKey:         encryption.EncryptionDataInstance.DERPubkey,
		VerifyTokenLength: len(encryption.EncryptionDataInstance.VerifyToken),
		VerifyToken:       encryption.EncryptionDataInstance.VerifyToken,
	}
}

func (packet *EncryptionRequestPacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	log.Panic("Proxy should never read an encryption request packet")
	return nil
}

func (packet *EncryptionRequestPacket) HandleRead(currentSession *session.Session) {
	panic("implement me")
}

func (packet *EncryptionRequestPacket) HandleWrite(currentSession *session.Session) {
	panic("implement me")
}

func (packet *EncryptionRequestPacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteString(packet.ServerId)

	currentSession.Writer.WriteVarInt(packet.PublicKeyLength)
	currentSession.Writer.WriteBytes(packet.PublicKey)

	currentSession.Writer.WriteVarInt(packet.VerifyTokenLength)
	currentSession.Writer.WriteBytes(packet.VerifyToken)

	currentSession.Writer.Flush(nil)
}
