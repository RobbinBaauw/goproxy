package main

import (
	"github.com/timanema/goproxy/server"
	"log"
)

//TODO: Catch signals for disabling etc
func main() {
	defer func() {
		log.Println("GoProxy is disabled")
	}()
	log.Println("Starting GoProxy")

	server := server.NewServer()
	server.StartServer()
}
