package main

import (
	"encoding/binary"
	"log"
)

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
			log.Fatal("VARINT READ FAILED")
		}
	}

	return result
}

func writeVarInt(value int) []byte {
	var result []byte

	for ok := true; ok; ok = value != 0 {
		temp := byte(value & 0b01111111)
		value = int(uint(value) >> 7)

		if value != 0 {
			temp |= 0b10000000;
		}

		result = append(result, temp);
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

func readLong(reader ByteReader) int64 {

	message := make([]byte, 8)

	for i := 0; i < 8; i++ {
		message[i] = reader.readNext()
	}

	return int64(binary.BigEndian.Uint64(message))
}
