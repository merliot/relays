//go:build !tinygo

package relays

import "embed"

//go:embed css go.mod *.go html images js template
var fs embed.FS
