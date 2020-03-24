package packets

import (
	"encoding/binary"
	"log"
)

func ReadVarInt(reader *ByteStreamReader) int {
	numRead := 0
	result := 0
	var read byte

	for ok := true; ok; ok = (read & 0b10000000) != 0 {
		read = reader.NextByte()

		value := int(read & 0b01111111)
		result |= value << (7 * numRead)

		numRead++
		if numRead > 5 {
			log.Fatal("VARINT READ FAILED")
		}
	}

	return result
}

func ReadString(reader *ByteStreamReader) string {
	length := ReadVarInt(reader)
	return string(ReadBytes(reader, length))
}

func ReadBytes(reader *ByteStreamReader, length int) []byte {
	message := make([]byte, length)

	for i := 0; i < length; i++ {
		message[i] = reader.NextByte()
	}

	return message
}


func ReadUnsignedShort(reader *ByteStreamReader) uint16 {
	byte1 := reader.NextByte()
	byte2 := reader.NextByte()

	return binary.BigEndian.Uint16([]byte {byte1, byte2})
}

func ReadLong(reader *ByteStreamReader) int64 {
	message := make([]byte, 8)

	for i := 0; i < 8; i++ {
		message[i] = reader.NextByte()
	}

	return int64(binary.BigEndian.Uint64(message))
}
