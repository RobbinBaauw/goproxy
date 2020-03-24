package main

import (
	"github.com/timanema/goproxy/server"
	"log"
)

func main() {
	log.Println("Starting GoProxy")
	server := server.NewServer()

	server.StartServer()
}
