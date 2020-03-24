package io

import (
	"encoding/binary"
	"log"
	"net"
)

type PacketWriter struct {
	conn net.Conn
	data []byte
}

func NewPacketWriter(conn net.Conn) *PacketWriter {
	packetWriter := new(PacketWriter)
	packetWriter.conn = conn
	packetWriter.data = make([]byte, 0)

	return packetWriter
}

func (writer *PacketWriter) Flush() {
	writer.data = append(writer.getVarInt(len(writer.data)), writer.data...)

	log.Print("Sending raw: ", writer.data)

	writer.conn.Write(writer.data)
	writer.data = make([]byte, 0)
}

func (writer *PacketWriter) write(data []byte) {
	writer.data = append(writer.data, data...)
}

func (writer *PacketWriter) getVarInt(value int) []byte {
	var res []byte

	for ok := true; ok; ok = value != 0 {
		temp := byte(value & 0b01111111)
		value = int(uint(value) >> 7)

		if value != 0 {
			temp |= 0b10000000
		}

		res = append(res, temp)
	}

	return res
}

func (writer *PacketWriter) WriteVarInt(value int) {
	writer.write(writer.getVarInt(value))
}

func (writer *PacketWriter) WriteString(s string) {
	len := len(s)

	// write len of string
	writer.WriteVarInt(len)

	writer.write([]byte(s))
}

func (writer *PacketWriter) WriteLong(value int64) {
	res := make([]byte, 8)

	binary.BigEndian.PutUint64(res, uint64(value))
	writer.write(res)
}
