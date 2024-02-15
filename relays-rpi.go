//go:build rpi

package relays

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/merliot/dean"
)

func (r *Relays) Run(i *dean.Injector) {

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-c:
		r.FailSafe()
	}
}
