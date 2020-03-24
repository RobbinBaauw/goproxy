package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	log.Println("Starting GoProxy")

	// start tcp server
	listener, error := net.Listen("tcp", "0.0.0.0:12345")
	if error != nil {
		log.Fatal("Unable to start GoProxy:", error)
	}

	for {
		// accept connections
		connection, error := listener.Accept()
		if error != nil {
			log.Fatal("Could not accept connection:", error)
		}

		log.Println("Incoming connection from:", connection.RemoteAddr().String())
		go handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	// Length
	length := readVarInt(&ByteStreamReader{reader})

	// Message
	message := make([]byte, length)
	_, _ = reader.Read(message)
	arrayReader := &ByteArrayReader{bytes: message}

	// Packetid
	packetId := readVarInt(arrayReader)
	data := arrayReader.getRest()

	log.Println(fmt.Sprintf("\nLength: " + strconv.Itoa(length) + "\n" +
		"Packet: " + strconv.Itoa(packetId) + "\n"))

	if packetId == 0 {
		handleHandshake(data)
	}

	handleConnection(conn)
}


func handleHandshake(data []byte) {
	arrayReader := &ByteArrayReader{bytes: data}

	protocolVersion := readVarInt(arrayReader)
	serverAddress := readString(arrayReader)
	serverPort := readUnsignedShort(arrayReader)
	nextState := readVarInt(arrayReader)

	log.Println("\nProtocol version: " + strconv.Itoa(protocolVersion) + "\n" +
		"Server address: " + serverAddress + "\n" +
		"Server port: " + strconv.FormatUint(uint64(serverPort), 10) + "\n" +
		"Next state: " + strconv.Itoa(nextState) + "\n")
}


type ByteReader interface {
	readNext() byte
}

// Stream
type ByteStreamReader struct {
	reader *bufio.Reader
}

func (r *ByteStreamReader) readNext() byte {
	readByte, _ := r.reader.ReadByte()
	return readByte
}

// Array
type ByteArrayReader struct {
	bytes []byte
	currentIndex int
}

func (r *ByteArrayReader) readNext() byte {
	if r.currentIndex >= len(r.bytes) {
		log.Fatal("OOF 1")
	}

	readByte := r.bytes[r.currentIndex]
	r.currentIndex++
	return readByte
}

func (r *ByteArrayReader) getRest() []byte {
	if r.currentIndex >= len(r.bytes) {
		log.Fatal("OOF 2")
	}

	readByte := r.bytes[r.currentIndex:]
	return readByte
}

// Read
func readVarInt(reader ByteReader) int {
	numRead := 0
	result := 0
	var read byte

	for ok := true; ok; ok = (read & 0b10000000) != 0 {
		read = reader.readNext()

		value := int(read & 0b01111111)
		result |= value << (7 * numRead)

		numRead++
		if numRead > 5 {
			log.Fatal("OOF 3")
		}
	}

	return result
}

func readString(reader ByteReader) string {
	length := readVarInt(reader)

	message := make([]byte, length)

	for i := 0; i < length; i++ {
		message[i] = reader.readNext()
	}

	return string(message)
}

func readUnsignedShort(reader ByteReader) uint16 {
	byte1 := reader.readNext()
	byte2 := reader.readNext()

	return binary.BigEndian.Uint16([]byte {byte1, byte2})
}
