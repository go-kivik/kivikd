package couchserver

import (
	"net/http/httptest"
	"testing"

	"github.com/go-kivik/kivik/v3"
	_ "github.com/go-kivik/memorydb/v3"
	"gitlab.com/flimzy/testy"
)

func TestAllDBs(t *testing.T) {
	client, err := kivik.New("memory", "")
	if err != nil {
		panic(err)
	}
	h := &Handler{client: &clientWrapper{client}}
	handler := h.GetAllDBs()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/_all_dbs", nil)
	handler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	expected := []string{}
	if d := testy.DiffAsJSON(expected, resp.Body); d != nil {
		t.Error(d)
	}
}
