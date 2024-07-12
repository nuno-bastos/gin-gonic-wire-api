package main

import (
	"log"

	wire "github.com/nuno-bastos/gin-gonic-wire-api/wire_di"
)

func main() {
	server, diErr := wire.Inject()
	if diErr != nil {
		log.Fatal("Error - Cannot Start Server: ", diErr)
	} else {
		server.Start()
	}
}
