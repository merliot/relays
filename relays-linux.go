//go:build !tinygo

package relays

import (
	"net/http"
)

func (r *Relays) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.API(w, req, r)
}

func (r *Relays) Icon() []byte {
	icon, _ := fs.ReadFile("images/icon.png")
	return icon
}

func (r *Relays) DescHtml() []byte {
	desc, _ := fs.ReadFile("html/desc.html")
	return desc
}

func (r *Relays) SupportedTargets() string {
	return r.Targets.FullNames()
}
