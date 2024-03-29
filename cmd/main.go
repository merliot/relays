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
	deployParams = dean.GetEnv("DEPLOY_PARAMS", "")
	//deployParams = dean.GetEnv("DEPLOY_PARAMS", "target=demo&relay1=kitchen&relay2=bar&relay3=porch&relay4=foo&gpio1=DEMO2&gpio2=DEMO0&gpio3=DEMO3&gpio4=DEMO1")
	wsScheme    = dean.GetEnv("WS_SCHEME", "ws://")
	port        = dean.GetEnv("PORT", "8000")
	portPrime   = dean.GetEnv("PORT_PRIME", "8001")
	user        = dean.GetEnv("USER", "")
	passwd      = dean.GetEnv("PASSWD", "")
	dialURLs    = dean.GetEnv("DIAL_URLS", "")
	ssids       = dean.GetEnv("WIFI_SSIDS", "")
	passphrases = dean.GetEnv("WIFI_PASSPHRASES", "")
)

func main() {
	relays := relays.New(id, "relays", name).(*relays.Relays)
	relays.SetDeployParams(deployParams)
	relays.SetWifiAuth(ssids, passphrases)
	relays.SetDialURLs(dialURLs)
	relays.SetWsScheme(wsScheme)
	runner.Run(relays, port, portPrime, user, passwd, dialURLs, wsScheme)
}
