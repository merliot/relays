package main

import (
	"github.com/merliot/dean"
	"github.com/merliot/dean/tinynet"
	"github.com/merliot/device"
	"github.com/merliot/relays"
)

func main() {
	p := device.DeviceParams()
	tinynet.NetConnect(p.Ssid, p.Passphrase)
	thing := relays.New(p.Id, p.Model, p.Name).(*relays.Relays)
	thing.SetDeployParams(p.DeployParams)
	runner := dean.NewRunner(thing, p.User, p.Passwd)
	runner.Dial(p.DialURLs)
	runner.Run()
}
