package auth

import (
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

func NewEncryptionResponsePacket(name string) *EncryptionResponsePacket {
	// TODO
	return nil
}

func (packet *EncryptionResponsePacket) Read(packetId int, reader *io.PacketReader) packets.Packet {
	packet.PacketId = packetId

	packet.SharedSecretLength = reader.ReadVarInt()
	packet.SharedSecret = reader.ReadBytes(packet.SharedSecretLength)

	packet.VerifyTokenLength = reader.ReadVarInt()
	packet.VerifyToken = reader.ReadBytes(packet.VerifyTokenLength)

	return packet
}

func (packet *EncryptionResponsePacket) Handle(currentSession *session.Session) {
	verifyToken := encryption.DecryptWithPrivateKey(packet.VerifyToken, encryption.EncryptionDataInstance.RSAKey)
	if !cmp.Equal(verifyToken, encryption.EncryptionDataInstance.VerifyToken) {
		log.Panic("Invalid verify token!")
	}

	decryptedSharedSecret := encryption.DecryptWithPrivateKey(packet.SharedSecret, encryption.EncryptionDataInstance.RSAKey)
	currentSession.SharedSecret = decryptedSharedSecret

	currentSession.PlayerData.UUID = "afa7545b-1b71-4a1b-a989-2c39b3f49600"

	currentSession.CurrentState = session.Play

	successPacket := NewLoginSuccessPacket(currentSession.PlayerData.Username, currentSession.PlayerData.UUID)
	successPacket.Write(currentSession)
}

func (packet *EncryptionResponsePacket) Write(currentSession *session.Session) {
	// TODO
	log.Panic("TODO!")
}
