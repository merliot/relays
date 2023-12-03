//go:build prime

package main

import (
	"github.com/merliot/dean"
	"github.com/merliot/prime"
	"github.com/merliot/relays"
)

func main() {
	relays := relays.New("relays01", "relays", "relays").(*relays.Relays)
	relays.SetDeployParams("target=nano-rp2040&relay1=kitchen+and+bath&relay2=closet&relay3=hallway&relay4=&gpio1=D2&gpio2=D4&gpio3=D5&gpio4=")
	device := prime.New("p1", "prime", "p1")
	server := dean.NewServer(device)
	server.AdoptThing(relays)
	server.Run()
}
