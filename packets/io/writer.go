package io

import "bufio"

type PacketWriter struct {
	writer *bufio.Writer
	data   []byte
}

func NewPacketWriter(writer *bufio.Writer) *PacketWriter {
	packetWriter := new(PacketWriter)
	packetWriter.writer = writer
	packetWriter.data = make([]byte, 0)

	return packetWriter
}

func (writer *PacketWriter) UpdateWriter(newWriter *bufio.Writer) {
	writer.writer = newWriter
}

func (writer *PacketWriter) Flush() {
	writer.writer.Write(append(writer.getVarInt(len(writer.data)), writer.data...))
	writer.writer.Flush()
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
