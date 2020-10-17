package segment

import (
	"net/http"
	"regexp"

	"github.com/purusah/zebra/pkg/handlers"
)

type Segment struct {
	descendants []*Segment
	name        string
	value       *regexp.Regexp
	handler     func(http.ResponseWriter, *http.Request)
	Middlewares []http.HandlerFunc // TODO Make private
}

func (s *Segment) SetName(name string) {
	s.name = name
}

func (s *Segment) SetValue(value string) {
	s.value = regexp.MustCompile(value)
}

func (s *Segment) SetHandler(handler http.HandlerFunc) {
	s.handler = handler
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

// Insert new segment to handlers tree
func InsertElement(root *Segment, new *Segment) *Segment {
	for _, s := range root.descendants {
		if s.Eq(new) {
			return s
		}
	}
	root.descendants = append(root.descendants, new)
	return new
}

// Find relative endpoint for request url
func FindElement(root *Segment, name string) *Segment {
	nameRaw := []byte(name)
	for _, s := range root.descendants {
		// No special matching
		if s.value == nil {
			if s.name == name {
				return s
			}
			continue
		}
		if s.value.Match(nameRaw) {
			return s
		}
	}
	return nil
}
