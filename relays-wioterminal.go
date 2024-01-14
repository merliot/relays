//go:build wioterminal

package relays

import "github.com/merliot/target"

func (r *Relays) pins() target.GpioPins {
	return r.Targets["wioterminal"].GpioPins
}
