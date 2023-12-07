// go run ./cmd/
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"os"

	"github.com/merliot/device/runner"
	"github.com/merliot/relays"
)

func get(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return value
}

var cfg = runner.Config{
	Port: get("PORT", "8000"),
	PortPrime: get("PORT_PRIME", "8001"),
	User: get("USER", ""),
	Passwd: get("PASSWD", ""),
	DialURLs: get("DIAL_URLS", "ws://127.0.0.1:8001/ws/?ping-period=4"),
}

var (
	id = get("ID", "relay01")
	model = get("MODEL", "relays")
	name = get("NAME", "relay01")
	deployParams = get("DEPLOY_PARAMS", "target=demo&relay1=kitchen&relay2=&relay3=&relay4=&gpio1=DEMO2&gpio2=&gpio3=&gpio4=")
)

func main() {
	relays := relays.New(id, model, name).(*relays.Relays)
	relays.SetDeployParams(deployParams)
	relays.ParseWifiAuth()
	println(cfg.User, cfg.Passwd)
	runner.Run(cfg, relays)
}
