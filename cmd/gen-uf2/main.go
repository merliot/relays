package main

import (
	"log"
	"os"

	"github.com/merliot/relays"
)

func main() {
	relays := relays.New("proto", "relays", "proto").(*relays.Relays)
	if err := relays.GenerateUf2s(); err != nil {
		log.Println("Error generating UF2s:", err)
		os.Exit(1)
	}
}
