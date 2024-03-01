package relays

import (
	"embed"
	"fmt"
	"net/http"
	"strconv"

	"github.com/merliot/dean"
	"github.com/merliot/device"
	"github.com/merliot/device/relay"
)

//go:embed css html images module.json js template
var fs embed.FS

const maxRelays int = 4

type Relays struct {
	*device.Device
	Relays []relay.Relay
}

type MsgClick struct {
	dean.ThingMsg
	Relay int
	State bool
}

var targets = []string{"demo", "rpi", "nano-rp2040", "wioterminal"}

func New(id, model, name string) dean.Thinger {
	println("NEW RELAYS")
	return &Relays{
		Device: device.New(id, model, name, fs, targets).(*device.Device),
		Relays: make([]relay.Relay, maxRelays),
	}
}

func (r *Relays) save(msg *dean.Msg) {
	msg.Unmarshal(r).Broadcast()
}

func (r *Relays) getState(msg *dean.Msg) {
	r.Path = "state"
	msg.Marshal(r).Reply()
}

func (r *Relays) click(msg *dean.Msg) {
	var msgClick MsgClick
	msg.Unmarshal(&msgClick)
	relay := &r.Relays[msgClick.Relay]
	relay.State = msgClick.State
	if r.IsMetal() {
		if msgClick.State {
			relay.On()
		} else {
			relay.Off()
		}
	}
	msg.Broadcast()
}

func (r *Relays) Subscribers() dean.Subscribers {
	return dean.Subscribers{
		"state":     r.save,
		"get/state": r.getState,
		"click":     r.click,
	}
}

func (r *Relays) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.API(w, req, r)
}

func (r *Relays) setRelay(num int, name, gpio string) {
	relay := &r.Relays[num]
	if name == "" {
		name = fmt.Sprintf("Relay %d", num+1)
	}
	relay.Name = name
	relay.Gpio = gpio
	relay.Configure()
}

func (r *Relays) parseParams() {
	for i := range r.Relays {
		num := strconv.Itoa(i + 1)
		name := r.ParamFirstValue("relay" + num)
		gpio := r.ParamFirstValue("gpio" + num)
		if gpio != "" {
			r.setRelay(i, name, gpio)
		}
	}
}

func (r *Relays) Setup() {
	r.Device.Setup()
	r.parseParams()
}
