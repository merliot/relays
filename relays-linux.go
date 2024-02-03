//go:build !tinygo

package relays

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/merliot/device"
)

//go:embed css images js template
var fs embed.FS

type osStruct struct {
	templates *template.Template
}

func (r *Relays) osNew() {
	r.CompositeFs.AddFS(fs)
	r.templates = r.CompositeFs.ParseFS("template/*")
}

func (r *Relays) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch strings.TrimPrefix(req.URL.Path, "/") {
	case "state":
		device.ShowState(r.templates, w, r)
	default:
		r.API(r.templates, w, req)
	}
}

func (r *Relays) Icon() []byte {
	icon, _ := fs.ReadFile("images/icon.png")
	return icon
}

func (r *Relays) DescHtml() []byte {
	return nil
}
