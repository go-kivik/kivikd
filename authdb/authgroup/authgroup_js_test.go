//go:build js
// +build js

package authgroup

import (
	"net/http"

	"github.com/gopherjs/gopherjs/js"
)

func init() {
	js.Global.Set("XMLHttpRequest", js.Global.Call("require", "xhr2"))
	http.DefaultTransport = &http.XHRTransport{}
}
