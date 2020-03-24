package io

import (
	"bufio"
	"encoding/binary"
	"log"
)

type PacketReader struct {
	reader *bufio.Reader
}

func NewPacketReader(reader *bufio.Reader) *PacketReader {
	packetReader := new(PacketReader)
	packetReader.reader = reader

	return packetReader
}

func (reader *PacketReader) UpdateReader(newReader *bufio.Reader) {
	reader.reader = newReader
}

func (reader *PacketReader) safeRead() byte {
	res, err := reader.reader.ReadByte()

	if err != nil {
		log.Panic("Could not read byte: ", err)
	}

	return res
}

func (reader *PacketReader) ReadVarInt() int {
	var result = 0
	var read byte
	var i = 0

	// do
	read = reader.safeRead()
	val := int(read & 0b01111111)
	result |= val << (7 * i)
	i++

	// while
	for ; (read & 0b10000000) != 0; i++ {
		read = reader.safeRead()
		val := int(read & 0b01111111)
		result |= val << (7 * i)
	}

	return result
}

func (reader *PacketReader) ReadString(len int) string {
	var res string
	for i := 0; i < len; i++ {
		ch := reader.safeRead()
		res += string(ch)
	}

	return res
}

func (reader *PacketReader) ReadUnsignedShort() uint16 {
	buf1 := reader.safeRead()
	buf2 := reader.safeRead()
	in := []byte{buf1, buf2}
	return binary.BigEndian.Uint16(in)
}

func (reader *PacketReader) ReadLong() int64 {
	message := make([]byte, 8)

	for i := 0; i < 8; i++ {
		message[i] = reader.safeRead()
	}

	return int64(binary.BigEndian.Uint64(message))
}
