package server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/finitum/goproxy/packets"
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

	key := getDERKey()
	keyLength := packets.WriteVarInt(len(key))

	verifyTokenLength := packets.WriteVarInt(4)
	verifyTokenBytes := make([]byte, 4)

	_, err := rand.Read(verifyTokenBytes)
	checkError(err)

	packets.Write(1, session.Conn, serverId, keyLength, key, verifyTokenLength, verifyTokenBytes)
}

func HandleEncryptionResponse(session *ClientSession) {
	playerName := packets.ReadString(session.Reader)
	session.PlayerData.Username = playerName
	fmt.Println("Player: ", playerName)

	serverId := packets.WriteString("                    ")

	key := getDERKey()
	keyLength := packets.WriteVarInt(len(key))

	verifyTokenLength := packets.WriteVarInt(4)
	verifyTokenBytes := make([]byte, 4)

	_, err := rand.Read(verifyTokenBytes)
	checkError(err)

	packets.Write(1, session.Conn, serverId, keyLength, key, verifyTokenLength, verifyTokenBytes)
}

func checkError(err error) {
	if err != nil {
		log.Panic("Fatal error ", err.Error())
	}
}

func getDERKey() []byte {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	checkError(err)

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	checkError(err)

	derKey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	return derKey.Bytes
}
