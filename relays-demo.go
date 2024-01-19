//go:build !tinygo && !rpi

package relays

import (
	"github.com/merliot/dean"
)

type targetRelayStruct struct {
}

type targetStruct struct {
	osStruct
}

func (r *Relays) targetNew() {
	r.osNew()
}

func (r *Relay) on() {
}

func (r *Relay) off() {
}

func (r *Relays) run(i *dean.Injector) {
	select {}
}
