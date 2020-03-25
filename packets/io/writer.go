package io

import (
	"encoding/binary"
	"io"
	"log"
)

type PacketWriter struct {
	Writer io.Writer
	data   []byte
}

func NewPacketWriter(writer io.Writer) *PacketWriter {
	return &PacketWriter{
		Writer: writer,
		data:   make([]byte, 0),
	}
}

func (writer *PacketWriter) Flush() {
	writer.data = append(writer.getVarInt(len(writer.data)), writer.data...)

	log.Print("Sending raw: ", writer.data)

	_, _ = writer.Writer.Write(writer.data)
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
