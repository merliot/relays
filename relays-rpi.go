//go:build rpi

package relays

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/merliot/dean"
	"github.com/merliot/target"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type targetStruct struct {
	osStruct
	adaptor *raspi.Adaptor
}

func (r *Relays) targetNew() {
	r.osNew()
	r.adaptor = raspi.NewAdaptor()
}

type Relay struct {
	Name   string
	Gpio   string
	State  bool
	driver *gpio.RelayDriver
}

func (r *Relay) start() {
	if r.driver != nil {
		r.driver.Start()
	}
}

func (r *Relay) on() {
	if r.driver != nil {
		r.driver.On()
	}
}

func (r *Relay) off() {
	if r.driver != nil {
		r.driver.Off()
	}
}

func (r *Relays) pins() target.GpioPins {
	return r.Targets["rpi"].GpioPins
}

// FailSafe by turning off all gpios
func (r *Relays) failSafe() {
	for _, pin := range r.pins() {
		rpin := strconv.Itoa(pin)
		driver := gpio.NewRelayDriver(r.adaptor, rpin)
		driver.Start()
		driver.Off()
	}
}

func (r *Relays) run(i *dean.Injector) {

	defer func() {
		if recover() != nil {
			r.failSafe()
		}
	}()

	r.adaptor.Connect()

	for i := range r.Relays {
		relay := &r.Relays[i]
		if relay.Gpio == "" {
			continue
		}
		if pin, ok := r.pins()[relay.Gpio]; ok {
			rpin := strconv.Itoa(pin)
			relay.driver = gpio.NewRelayDriver(r.adaptor, rpin)
			relay.start()
			relay.off()
		}
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-c:
		r.failSafe()
	}
}
