// Package provide tree structure to save mapping request paths to corresponding
// handlers.
package segment

import (
	"net/http"
	"regexp"

	"github.com/purusah/zebra/pkg/handlers"
)

// Segment Basic structure to represent one path segment. One level segments will be
// composed to one.
type Segment struct {
	descendants []*Segment
	name        string
	value       *regexp.Regexp
	handler     func(http.ResponseWriter, *http.Request)
	Middlewares []http.HandlerFunc // TODO Make private
}

// SetName Set segment name for full equal matching or matching with regular expressions
func (s *Segment) SetName(name string) {
	s.name = name
}

// SetValue Set matching segment to regular expression
func (s *Segment) SetValue(value string) {
	s.value = regexp.MustCompile(value)
}

// SetHandler Set handler func to path segment
func (s *Segment) SetHandler(handler http.HandlerFunc) {
	s.handler = handler
}

func (s *Segment) Match(name string) bool {
	if s.value == nil {
		if s.name == name {
			return true
		}
		return false
	}
	if s.value.MatchString(name) {
		return true
	}
	return false
}

func (s *Segment) Eq(other *Segment) bool {
	return s.name == other.name
	// TODO Think about segments comparing
	//if s.name != other.name {
	//	return false
	//} else if s.value == nil && other.value == nil {
	//	return false
	//} else if s.value == nil && other.value != nil {
	//	return false
	//} else if s.value != nil && other.value == nil {
	//	return false
	//} else if s.value.String() != other.value.String() {
	//	return true
	//}
	//return false
}

func (s *Segment) Handle(w http.ResponseWriter, r *http.Request) {
	s.handler(w, r)
}

func NewSegment(name string) *Segment {
	return &Segment{
		descendants: nil,
		name:        name,
		value:       nil,
		handler:     handlers.NotFoundHandler,
		Middlewares: nil,
	}
}

// InsertElement Insert new segment to handlers tree
func InsertElement(root *Segment, new *Segment) *Segment {
	for _, s := range root.descendants {
		if s.Eq(new) {
			return s
		}
	}
	root.descendants = append(root.descendants, new)
	return new
}

// FindElement Find relative endpoint for request url
func FindElement(root *Segment, name string) *Segment {
	for _, s := range root.descendants {
		if s.Match(name) {
			return s
		}
		continue
	}
	return nil
}
