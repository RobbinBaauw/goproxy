package io

import (
	"encoding/binary"
	"github.com/timanema/goproxy/util/encryption"
	"log"
	"net"
)

type PacketWriter struct {
	conn net.Conn
	data []byte
}

func NewPacketWriter(conn net.Conn) *PacketWriter {
	return &PacketWriter{
		conn: conn,
		data: make([]byte, 0),
	}
}

func (writer *PacketWriter) Flush(encryptionKey *[]byte) {
	writer.data = append(writer.getVarInt(len(writer.data)), writer.data...)

	log.Print("Sending raw: ", writer.data)

	if encryptionKey != nil {
		writer.data = encryption.Encrypt(encryptionKey, &writer.data)
	}

	_, _ = writer.conn.Write(writer.data)
	writer.data = make([]byte, 0)
}

func (writer *PacketWriter) WriteBytes(data []byte) {
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
	writer.WriteBytes(writer.getVarInt(value))
}

func (writer *PacketWriter) WriteString(s string) {
	stringLen := len(s)

	writer.WriteVarInt(stringLen)
	writer.WriteBytes([]byte(s))
}

func (writer *PacketWriter) WriteLong(value int64) {
	res := make([]byte, 8)

	binary.BigEndian.PutUint64(res, uint64(value))
	writer.WriteBytes(res)
}
