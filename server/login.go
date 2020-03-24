package server

import (
	"fmt"
	"github.com/finitum/goproxy/packets"
	"github.com/finitum/goproxy/util"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"log"
)

func HandleLoginState(packetId int, session *ClientSession) {
	if packetId == 0 {
		HandleLoginStart(session)
		HandleConnection(session)
	} else if packetId == 1 {
		HandleEncryptionResponse(session)
		HandleConnection(session)
	} else {
		log.Panic("Unknown packet id ", packetId)
	}
}

func HandleLoginStart(session *ClientSession) {
	playerName := packets.ReadString(session.Reader)
	session.PlayerData.Username = playerName
	fmt.Println("Player: ", playerName)

	serverId := packets.WriteString("                    ")

	key := util.DERPublicKey
	keyLength := packets.WriteVarInt(len(key))

	verifyTokenLength := packets.WriteVarInt(len(util.VerifyToken))

	packets.Write(1, session.Conn, serverId, keyLength, key, verifyTokenLength, util.VerifyToken)
}

func HandleEncryptionResponse(session *ClientSession) {
	sharedSecretLength := packets.ReadVarInt(session.Reader)
	sharedSecret := packets.ReadBytes(session.Reader, sharedSecretLength)

	verifyTokenLength := packets.ReadVarInt(session.Reader)
	verifyToken := util.DecryptWithPrivateKey(packets.ReadBytes(session.Reader, verifyTokenLength))
	if !cmp.Equal(verifyToken, util.VerifyToken) {
		log.Panic("Invalid verify token!")
	}

	decryptedSharedSecret := util.DecryptWithPrivateKey(sharedSecret)
	session.PlayerData.SharedSecret = decryptedSharedSecret

	generatedUuid, _ := uuid.NewUUID()
	uuidBytes := packets.WriteString(generatedUuid.String())

	username := packets.WriteString(session.PlayerData.Username)

	packets.WriteEncrypted(decryptedSharedSecret, 2, session.Conn, uuidBytes, username)

	session.State = StatePlay
}




