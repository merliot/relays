//go:build tinygo

package relays

import (
	"machine"

	"github.com/merliot/dean"
)

type targetRelayStruct struct {
	pin machine.Pin
}

type targetStruct struct {
}

func (r *Relays) targetNew() {
}

func (r *Relay) on() {
	if r.pin != machine.NoPin {
		r.pin.High()
	}
}

func (r *Relay) off() {
	if r.pin != machine.NoPin {
		r.pin.Low()
	}
}

func (r *Relays) run(i *dean.Injector) {

	for i := range r.Relays {
		relay := &r.Relays[i]
		relay.pin = machine.NoPin
		if pin, ok := r.pins()[relay.Gpio]; ok {
			relay.pin = machine.Pin(pin)
			relay.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		}
	}

	select {}
}
