package zebra

import (
	"net/http"
	"strings"

	sg "github.com/purusah/zebra/internal/segment"
	hd "github.com/purusah/zebra/pkg/handlers"
)

const PathSeparator = "/"
const PathNamedSegmentIndicator = ":"

type Router struct {
	root *sg.Segment
}

// ServeHTTP
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	next := rt.root
	parts := strings.Split(r.URL.Path, PathSeparator)[1:]
	for _, part := range parts {
		next = sg.FindElement(next, part)
		if next == nil {
			hd.NotFoundHandler(w, r)
			return
		}
	}
	for _, mw := range next.Middlewares {
		mw(w, r)
	}
	next.Handle(w, r)
}

// Handle Add new endpoint to router
func (rt *Router) Handle(pattern string, handlerFunc http.HandlerFunc, mw ...http.HandlerFunc) {
	var s *sg.Segment
	next := rt.root
	patternParts := strings.Split(pattern, PathSeparator)[1:]
	lastIx := len(patternParts) - 1
	for i, part := range patternParts {
		s = sg.NewSegment(part)
		partData := strings.Split(part, PathNamedSegmentIndicator)
		if len(partData) > 2 {
			panic("endpoint matching part must contain 2 parts separated with colon")
		}
		if len(partData) == 2 {
			s.SetName(partData[0][1:])
			s.SetValue(partData[1][:len(partData[1]) - 1])
		}
		if i == lastIx {
			s.SetHandler(handlerFunc)
		}
		sg.InsertElement(next, s)
	}
}

// NewRouter construct new router instance
func NewRouter() *Router {
	return &Router{
		root: sg.NewSegment(""),
	}
}
