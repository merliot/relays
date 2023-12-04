// go run ./cmd/
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"github.com/merliot/device/runner"
	"github.com/merliot/relays"
)

var cfg = runner.Config{
	Port: "8000",
	PortPrime: "8001",
	User: "user",
	Passwd: "passwd",
	DialURLs: "ws://127.0.0.1:8001/ws/?ping-period=4",
}

func main() {
	relays := relays.New("relay01", "relays", "relays").(*relays.Relays)
	relays.SetDeployParams("target=nano-rp2040&relay1=kitchen+and+bath&relay2=closet&relay3=hallway&relay4=&gpio1=D2&gpio2=D4&gpio3=D5&gpio4=")
	runner.Run(cfg, relays)
}
