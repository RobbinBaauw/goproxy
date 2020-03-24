package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
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
		go handleConnection(&ClientConnection{
			connection,
			bufio.NewReader(connection),
			INIT,
		})
	}
}

func handleConnection(conn *ClientConnection) {
	// Length
	length := readVarInt(&ByteStreamReader{conn.reader})
	
	// Message
	message := make([]byte, length)

	_, err := conn.reader.Read(message)
	if err != nil {
		log.Println("closed connection")
		_ = conn.conn.Close()
		return
	}

	arrayReader := &ByteArrayReader{bytes: message}

	// Packetid
	packetId := readVarInt(arrayReader)

	log.Println(fmt.Sprintf("\nLength: " + strconv.Itoa(length) + "\n" +
		"Packet: " + strconv.Itoa(packetId) + "\n"))

	data := arrayReader.getRest()

	if conn.nextState == INIT && packetId == 0 {
		handleHandshake(data, conn)
		handleConnection(conn)
	} else if conn.nextState == REQUEST && packetId == 0 {
		handleRequest(conn)
		handleConnection(conn)
	} else if conn.nextState == PING && packetId == 1 {
		handlePing(data, conn)
	}
}

