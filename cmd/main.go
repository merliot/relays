// go run ./cmd/
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"github.com/merliot/relays"
	"github.com/merliot/device/runner"
)

func main() {
	relays := relays.New("relays01", "relays", "relays").(*relays.Relays)
	relays.SetDeployParams("target=nano-rp2040&relay1=kitchen+and+bath&relay2=closet&relay3=hallway&relay4=&gpio1=D2&gpio2=D4&gpio3=D5&gpio4=")
	runner.Run(relays)
}
