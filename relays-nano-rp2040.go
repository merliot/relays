//go:build nano_rp2040

package relays

import "github.com/merliot/device/target"

func (r *Relays) pins() target.GpioPins {
	return r.Targets["nano-rp2040"].GpioPins
}
