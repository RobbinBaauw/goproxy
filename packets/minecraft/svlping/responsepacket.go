package svlping

import (
	"encoding/json"
	"github.com/timanema/goproxy/server/session"
)

type ResponsePacket struct {
	PacketId     int
	jsonResponse string
}

func NewResponsePacket() *ResponsePacket {
	response := ListResponse{
		Version: ListVersion{
			Name:     "KLAPPE",
			Protocol: 578,
		},
		Players: ListPlayers{
			Max:    69,
			Online: 42,
			Sample: []ListPlayerSample{},
		},
		Description: ListDescription{
			Text: "\u00A7bKlappe",
		},
		Favicon: "",
	}

	data, _ := json.Marshal(response)

	packet := new(ResponsePacket)
	packet.PacketId = 0
	packet.jsonResponse = string(data)

	return packet
}

func (packet *ResponsePacket) Write(currentSession *session.Session) {
	currentSession.Writer.WriteVarInt(packet.PacketId)
	currentSession.Writer.WriteString(packet.jsonResponse)
	currentSession.Writer.Flush()
}

type ListResponse struct {
	Version     ListVersion     `json:"version"`
	Players     ListPlayers     `json:"players"`
	Description ListDescription `json:"description"`
	Favicon     string          `json:"favicon"`
}

type ListVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type ListPlayers struct {
	Max    int                `json:"max"`
	Online int                `json:"online"`
	Sample []ListPlayerSample `json:"sample"`
}

type ListPlayerSample struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ListDescription struct {
	Text string `json:"text"`
}
