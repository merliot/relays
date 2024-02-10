//go:build !tinygo

package relays

import (
	"net/http"
)

func (r *Relays) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.API(w, req, r)
}
