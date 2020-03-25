package server

import (
	"bufio"
	"encoding/json"
	"github.com/timanema/goproxy/packets"
	"github.com/timanema/goproxy/packets/io"
	"github.com/timanema/goproxy/server/handlers"
	"github.com/timanema/goproxy/server/session"
	"log"
	"net"
	"reflect"
)

type Server struct {
	sessions map[string]*session.Session
}

func NewServer() *Server {
	return &Server{
		sessions: make(map[string]*session.Session),
	}
}

func (server *Server) StartServer() {
	// start tcp server
	log.Print("Listening on 0.0.0.0:12345")

	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal("Unable to start GoProxy:", err)
	}

	for {
		// accept connections
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Could not accept conn:", err)
		}

		go server.acceptConnection(conn)
	}
}

func (server *Server) acceptConnection(conn net.Conn) {
	log.Println("Incoming connection from:", conn.RemoteAddr().String())
	newSession := session.NewSession(&conn)
	newSession.Writer = io.NewPacketWriter(conn)
	newSession.Reader = io.NewPacketReader(bufio.NewReader(conn))

	server.sessions[newSession.SessionId] = newSession
	server.acceptPacket(newSession)
}

func (server *Server) acceptPacket(currentSession *session.Session) {
	defer func() {
		if r := recover(); r != nil {
			addr := (*currentSession.Connection).RemoteAddr().String()
			log.Print("Error from client (", addr, "): ", r)
			log.Print("Client ", addr, " unexpectedly closed the connection")
		}

		delete(server.sessions, currentSession.SessionId)
	}()

	for {
		if currentSession.ConnectionClosed {
			break
		}

		// get packet id and packet
		packet, packetId := server.readPacket(currentSession, currentSession.Reader)

		// do any preread events
		packet.PreRead(currentSession)

		// read packet data
		packet.Read(packetId, currentSession.Reader, 0)

		// debug prints
		out, _ := json.Marshal(packet)
		log.Println("Got packet:", string(out), " of type ", reflect.TypeOf(packet))

		// do any postread events and get resulting reponse
		responsePacket := packet.PostRead(currentSession)

		// send response packet if needed
		if responsePacket != nil {
			responsePacket.PreWrite(currentSession)
			responsePacket.Write(currentSession.Writer)
			currentSession.Writer.Flush()
			responsePacket.PostWrite(currentSession)
		}
	}
}

func (server *Server) readPacket(currentSession *session.Session, packetReader *io.PacketReader) (packets.Packet, int) {
	packetLen := packetReader.ReadVarInt()
	packetId := packetReader.ReadVarInt()
	log.Print("Incoming packet of len ", packetLen, " with id ", packetId)

	switch currentSession.CurrentState {
	case session.Handshaking:
		return handlers.HandleHandshake(packetId), packetId
	case session.Status:
		return handlers.HandleStatus(packetId), packetId
	case session.Login:
		return handlers.HandleLogin(packetId), packetId
	case session.Play:
		return handlers.HandlePlay(packetId), packetId
	default:
		log.Panic("Unknown session state: ", currentSession.CurrentState)
		return nil, 0
	}
	return nil, 0
}
