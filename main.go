package main

import (
	"bufio"
	"github.com/finitum/goproxy/packets"
	"github.com/finitum/goproxy/server"
	"log"
	"net"
)

func main() {
	log.Println("Starting GoProxy")

	// start tcp server
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal("Unable to start GoProxy:", err)
	}

	for {
		// accept connections
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Could not accept connection:", err)
		}

		log.Println("Incoming connection from:", connection.RemoteAddr().String())

		go server.HandleConnection(&server.ClientSession{
			Conn:   connection,
			Reader: packets.ConstructByteStreamReader(bufio.NewReader(connection)),
			State:  server.StateHandshaking,
		})
	}
}
