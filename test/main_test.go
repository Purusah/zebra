// +build all integration

package test

import (
	"fmt"
	"github.com/purusah/zebra"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestMatchDefault(t *testing.T) {
	expected := http.StatusNotFound
	r := zebra.NewRouter()
	srv := httptest.NewServer(r)
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/", srv.URL))
	if err != nil {
		t.Errorf("unexpected error on request: %s", err)
	}
	if res.StatusCode != expected {
		t.Errorf("response code must be: %d", expected)
	}

	res, err = http.Get(fmt.Sprintf("%s/test", srv.URL))
	if err != nil {
		t.Errorf("unexpected error on request: %s", err)
	}
	if res.StatusCode != expected {
		t.Errorf("response code must be: %d", expected)
	}
}
