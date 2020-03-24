package server

import (
	"fmt"
	"github.com/finitum/goproxy/packets"
	"log"
	"strconv"
)

func HandleHandshakeState(packetId int, session *ClientSession) {
	if packetId == 0 {
		HandleHandshake(session)
		HandleConnection(session)
	} else {
		log.Panic("Unknown packet id ", packetId)
	}
}

func HandleHandshake(session *ClientSession) {
	reader := session.Reader

	protocolVersion := packets.ReadVarInt(reader)
	serverAddress := packets.ReadString(reader)
	serverPort := packets.ReadUnsignedShort(reader)
	nextState := packets.ReadVarInt(reader)

	if nextState == 1 {
		session.State = StateStatus
	} else {
		session.State = StateLogin
	}

	fmt.Println("\nProtocol version: " + strconv.Itoa(protocolVersion) + "\n" +
		"Server address: " + serverAddress + "\n" +
		"Server port: " + strconv.FormatUint(uint64(serverPort), 10) + "\n" +
		"Next state: " + strconv.Itoa(nextState) + "\n")
}
