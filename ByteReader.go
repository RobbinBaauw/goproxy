package main

import (
	"bufio"
	"log"
)

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
		log.Fatal("READING INVALID INDEX")
	}

	readByte := r.bytes[r.currentIndex]
	r.currentIndex++
	return readByte
}

func (r *ByteArrayReader) getRest() []byte {
	if r.currentIndex > len(r.bytes) {
		log.Fatal("PROGRESSED TOO FAR!")
	}

	readByte := r.bytes[r.currentIndex:]
	return readByte
}
