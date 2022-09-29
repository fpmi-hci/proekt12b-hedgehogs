package main

import (
	"os"

	"github.com/fpmi-hci/proekt12b-hedgehogs/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	server := new(internal.Server)
	if err := server.Run(os.Getenv("SERVPORT")); err != nil {
		log.Fatalf("error while running server %s", err.Error())
	}
}
