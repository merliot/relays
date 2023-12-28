// go run ./cmd
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"github.com/merliot/dean"
	"github.com/merliot/device/runner"
	"github.com/merliot/relays"
)

var (
	id           = dean.GetEnv("ID", "relay01")
	name         = dean.GetEnv("NAME", "Relays")
	deployParams = dean.GetEnv("DEPLOY_PARAMS", "target=demo&relay1=kitchen&relay2=&relay3=&relay4=&gpio1=DEMO2&gpio2=&gpio3=&gpio4=")
	port         = dean.GetEnv("PORT", "8000")
	portPrime    = dean.GetEnv("PORT_PRIME", "8001")
	user         = dean.GetEnv("USER", "")
	passwd       = dean.GetEnv("PASSWD", "")
	dialURLs     = dean.GetEnv("DIAL_URLS", "ws://127.0.0.1:8001/ws/?ping-period=4")
	ssids        = dean.GetEnv("WIFI_SSIDS", "")
	passphrases  = dean.GetEnv("WIFI_PASSPHRASES", "")
)

func main() {
	relays := relays.New(id, "relays", name).(*relays.Relays)
	relays.SetDeployParams(deployParams)
	relays.SetWifiAuth(ssids, passphrases)
	runner.Run(relays, port, portPrime, user, passwd, dialURLs)
}
