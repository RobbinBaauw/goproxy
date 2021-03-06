package auth

import (
	"bufio"
	"github.com/google/go-cmp/cmp"
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/session"
	"github.com/timanema/goproxy/util/encryption"
	"log"
)

type EncryptionResponsePacket struct {
	PacketId           int
	SharedSecretLength int
	SharedSecret       []byte
	VerifyTokenLength  int
	VerifyToken        []byte
}

func NewEncryptionResponsePacket(name string) packets.Packet {
	// TODO
	return nil
}

func (packet *EncryptionResponsePacket) PreRead(_ *session.Session) {}

func (packet *EncryptionResponsePacket) Read(packetId int, reader *io.PacketReader, _ int) packets.Packet {
	packet.PacketId = packetId

	packet.SharedSecretLength = reader.ReadVarInt()
	packet.SharedSecret = reader.ReadBytes(packet.SharedSecretLength)

	packet.VerifyTokenLength = reader.ReadVarInt()
	packet.VerifyToken = reader.ReadBytes(packet.VerifyTokenLength)

	return packet
}

func (packet *EncryptionResponsePacket) PostRead(currentSession *session.Session) packets.Packet {
	verifyToken := encryption.DecryptWithPrivateKey(packet.VerifyToken, encryption.EncryptionDataInstance.RSAKey)
	if !cmp.Equal(verifyToken, encryption.EncryptionDataInstance.VerifyToken) {
		log.Panic("Invalid verify token!")
	}

	decryptedSharedSecret := encryption.DecryptWithPrivateKey(packet.SharedSecret, encryption.EncryptionDataInstance.RSAKey)
	codec := encryption.NewCFBCodec(&decryptedSharedSecret)

	currentSession.PlayerData.UUID = "159e238f-c6a5-499f-97bd-cdcdd8012135"

	successPacket := NewLoginSuccessPacket(currentSession.PlayerData.Username, currentSession.PlayerData.UUID)

	// set encryption streams
	codec.Encrypt(currentSession.Writer.Writer)
	codec.Decrypt(currentSession.Reader.Reader)

	// update readers
	currentSession.Writer = io.NewPacketWriter(codec.Writer)
	currentSession.Reader = io.NewPacketReader(bufio.NewReader(codec.Reader))

	return successPacket
}

func (packet *EncryptionResponsePacket) PreWrite(_ *session.Session) {}

func (packet *EncryptionResponsePacket) Write(writer *io.PacketWriter) {
	// TODO
}

func (packet *EncryptionResponsePacket) PostWrite(_ *session.Session) {}
