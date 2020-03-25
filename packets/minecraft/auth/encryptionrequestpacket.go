package auth

import (
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
	"github.com/timanema/goproxy/util/encryption"
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

func (packet *EncryptionRequestPacket) PreRead(_ *session.Session) {}

func (packet *EncryptionRequestPacket) Read(_ int, _ *io.PacketReader, _ int) packets.Packet {
	return nil
}

func (packet *EncryptionRequestPacket) PostRead(_ *session.Session) packets.Packet {
	return nil
}

func (packet *EncryptionRequestPacket) PreWrite(_ *session.Session) {}

func (packet *EncryptionRequestPacket) Write(writer *io.PacketWriter) {
	writer.WriteVarInt(packet.PacketId)
	writer.WriteString(packet.ServerId)

	writer.WriteVarInt(packet.PublicKeyLength)
	writer.WriteBytes(packet.PublicKey)

	writer.WriteVarInt(packet.VerifyTokenLength)
	writer.WriteBytes(packet.VerifyToken)
}

func (packet *EncryptionRequestPacket) PostWrite(_ *session.Session) {}
