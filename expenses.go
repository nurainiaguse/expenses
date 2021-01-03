package main

import (
	"log"

	"github.com/nurainiaguse/expenses/internal/server"
)

func main() {
	server := server.NewServer()
	server.Port = 8080

	defer func() {
		_ = server.Shutdown()
	}()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
