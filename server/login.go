package server

import (
	"fmt"
	"github.com/finitum/goproxy/packets"
	"log"
)

func handleLoginState(packetId int, session *ClientSession) {
	if packetId == 0 {
		handleLoginStart(session)
		HandleConnection(session)
	} else {
		log.Panic("Unknown packet id ", packetId)
	}
}

func handleLoginStart(session *ClientSession) {
	playerName := packets.ReadString(session.Reader)
	session.PlayerData.Username = playerName
	fmt.Println("Player: ", playerName)

	//serverId := packets.WriteString("                    ")

}
