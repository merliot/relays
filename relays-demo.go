//go:build !tinygo && !rpi

package relays

import (
	"github.com/merliot/dean"
)

type targetStruct struct {
	osStruct
}

func (r *Relays) targetNew() {
	r.osNew()
}

type Relay struct {
	Name   string
	Gpio   string
	State  bool
}

func (r *Relay) on() {
}

func (r *Relay) off() {
}

func (r *Relays) run(i *dean.Injector) {
	select {}
}
