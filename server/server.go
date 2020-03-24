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
	sessions         []*session.Session
	handshakeHandler *handlers.HandshakeHandler
	statusHandler    *handlers.StatusHandler
}

func NewServer() *Server {
	server := new(Server)
	server.handshakeHandler = new(handlers.HandshakeHandler)
	server.statusHandler = new(handlers.StatusHandler)

	return server
}

func (server *Server) StartServer() {
	// start tcp server
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
	session := session.NewSession(&conn)
	session.Writer = io.NewPacketWriter(conn)
	session.Reader = io.NewPacketReader(bufio.NewReader(conn))

	server.sessions = append(server.sessions, session)
	server.acceptPacket(session)
}

func (server *Server) readPacket(currentSession *session.Session, packetReader *io.PacketReader) packets.Packet {
	len := packetReader.ReadVarInt()
	packetId := packetReader.ReadVarInt()
	log.Print("Incoming packet of len ", len, " with id ", packetId)

	switch currentSession.CurrentState {
	case session.Handshaking:
		return server.handshakeHandler.Handle(packetReader, packetId)
	case session.Status:
		return server.statusHandler.Handle(packetReader, packetId)
	default:
		log.Panic("Unknown session state: ", currentSession.CurrentState)
		return nil
	}
}

func (server *Server) acceptPacket(currentSession *session.Session) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("Client ", (*currentSession.Connection).RemoteAddr().String(), " unexpectedly closed the connection")
		}

		//TODO: Remove session from memory
	}()

	for {
		if currentSession.ConnectionClosed {
			break
		}

		// create reader and read packer
		packet := server.readPacket(currentSession, currentSession.Reader)

		// debug prints
		out, _ := json.Marshal(packet)
		log.Println("Got packet:", string(out), " of type ", reflect.TypeOf(packet))

		// handle packet
		packet.Handle(currentSession)
	}
}
