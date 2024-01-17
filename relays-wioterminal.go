//go:build wioterminal

package relays

import "github.com/merliot/device/target"

func (r *Relays) pins() target.GpioPins {
	return r.Targets["wioterminal"].GpioPins
}
