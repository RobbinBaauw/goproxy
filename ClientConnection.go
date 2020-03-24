package main

import (
	"bufio"
	"net"
)

type ClientConnection struct {
	conn net.Conn
	reader *bufio.Reader
	nextState int
}

const INIT = 0
const REQUEST = 1
const PING = 2
