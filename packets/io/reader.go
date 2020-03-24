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
	numRead := 0
	result := 0
	var read byte

	for ok := true; ok; ok = (read & 0b10000000) != 0 {
		read = reader.safeRead()

		value := int(read & 0b01111111)
		result |= value << (7 * numRead)

		numRead++
		if numRead > 5 {
			log.Panic("VARINT READ FAILED")
		}
	}

	return result
}

func (reader *PacketReader) ReadString() string {
	length := reader.ReadVarInt()
	return string(reader.ReadBytes(length))
}

func (reader *PacketReader) ReadBytes(length int) []byte {
	message := make([]byte, length)

	for i := 0; i < length; i++ {
		message[i] = reader.safeRead()
	}

	return message
}

func (reader *PacketReader) ReadUnsignedShort() uint16 {
	return binary.BigEndian.Uint16([]byte{
		reader.safeRead(),
		reader.safeRead(),
	})
}

func (reader *PacketReader) ReadLong() int64 {
	message := make([]byte, 8)

	for i := 0; i < 8; i++ {
		message[i] = reader.safeRead()
	}

	return int64(binary.BigEndian.Uint64(message))
}
