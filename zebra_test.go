// +build all unit

package zebra

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	HandlerFuncDummy = func(_ http.ResponseWriter, _ *http.Request) {}
)

type RouterTest struct {
	name string
	handler http.HandlerFunc
	method string
	path []string
	expected map[string]int
}

func testRouter(t *testing.T, rt RouterTest) {
	client := &http.Client{}
	r := NewRouter()
	for _, p := range rt.path {
		r.Handle(p, rt.handler)
	}

	srv := httptest.NewServer(r)
	defer srv.Close()
	for url, expected := range rt.expected {
		req, err := http.NewRequest(rt.method, fmt.Sprintf("%s%s", srv.URL, url), nil)
		if err != nil {
			t.Errorf("unexpected error on request: %s for %s", err, rt.name)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("unexpected error from response: %s for %s", err, rt.name)
		}

		if res.StatusCode != expected {
			t.Errorf("unexpected status on request: %v for %s", res.StatusCode, rt.name)
		}
	}
}

func TestTailMatch(t *testing.T) {
	tests := []RouterTest{
		{
			name: "Test1",
			handler: HandlerFuncDummy,
			method: "GET",
			path: []string{"/"},
			expected: map[string]int{"/": http.StatusOK, "/a": http.StatusNotFound},
		},
		{
			name: "Test2",
			handler: HandlerFuncDummy,
			method: "GET",
			path: []string{"/api"},
			expected: map[string]int{"/": http.StatusNotFound, "/api": http.StatusOK, "/api/": http.StatusNotFound},
		},
		{
			name: "Test3",
			handler: HandlerFuncDummy,
			method: "GET",
			path: []string{"/api/"},
			expected: map[string]int{"/": http.StatusNotFound, "/api": http.StatusNotFound, "/api/": http.StatusOK},
		},
		{
			name: "Test4",
			handler: HandlerFuncDummy,
			method: "GET",
			path: []string{"/api/{version:v[0-9]+}"},
			expected: map[string]int{"/api/": http.StatusNotFound, "/api/v}": http.StatusNotFound, "/api/v1": http.StatusOK},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testRouter(t, test)
		})
	}
}
