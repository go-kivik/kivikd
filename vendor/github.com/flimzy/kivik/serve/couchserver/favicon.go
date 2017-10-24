package couchserver

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/errors"
)

//go:generate go-bindata -pkg couchserver -nocompress -prefix files -o files.go files

// GetFavicon serves GET /favicon.ico
func (h *Handler) GetFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ico io.Reader
		if h.Favicon == "" {
			asset, err := Asset("favicon.ico")
			if err != nil {
				panic(err)
			}
			ico = bytes.NewBuffer(asset)
		} else {
			file, err := os.Open(h.Favicon)
			if err != nil {
				if os.IsNotExist(err) {
					err = errors.Status(kivik.StatusNotFound, "not found")
				}
				h.HandleError(w, err)
				return
			}
			ico = file
			defer file.Close()
		}
		w.Header().Set("Content-Type", "image/x-icon")
		_, err := io.Copy(w, ico)
		h.HandleError(w, err)
	}
}
