// +build !js

package test

import (
	"testing"

	_ "github.com/go-kivik/kivikd"
)

func init() {
	RegisterKivikdSuites()
}

func TestServer(t *testing.T) {
	ServerTest(t)
}
