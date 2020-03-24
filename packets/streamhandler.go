package packets

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type ByteStreamReader struct {
	reader *bufio.Reader
	remainingBytes int
}

func ConstructByteStreamReader(reader *bufio.Reader) *ByteStreamReader {
	return &ByteStreamReader{
		reader:         reader,
		remainingBytes: 5,
	}
}

func (r *ByteStreamReader) ReadLength() {
	r.remainingBytes = 5
	r.remainingBytes = ReadVarInt(r)
	fmt.Println("Length: ", r.remainingBytes)
}

func (r *ByteStreamReader) NextByte() byte {
	if r.remainingBytes == 0 {
		log.Panic("Attempted to read beyond packet!")
	}

	readByte, err := r.reader.ReadByte()

	r.remainingBytes--

	if err != nil {
		log.Panic("Could not read byte: ", err)
	}

	return readByte
}

func (r *ByteStreamReader) AllBytes() []byte {
	if r.remainingBytes == 0 {
		log.Panic("Attempted to read beyond packet!")
	}

	result := make([]byte, r.remainingBytes)
	_, err := r.reader.Read(result)

	r.remainingBytes = 0

	if err != nil {
		log.Panic("Could not read byte: ", err)
	}

	return result
}

func Write(packetId int, data []byte, conn net.Conn) {
	packetIdBytes := WriteVarInt(packetId)

	length := len(packetIdBytes) + len(data)
	lengthBytes := WriteVarInt(length)

	message := append(lengthBytes, append(packetIdBytes, data...)...)

	fmt.Println("WRITING MESSAGE " + string(message))
	fmt.Println()

	_, _ = conn.Write(message)
}
