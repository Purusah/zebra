// +build all unit

package segment

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestInsertDifferentElementsDummy(t *testing.T) {
	root := NewSegment("")
	subSegment1 := NewSegment("test1")
	subSegment2 := NewSegment("test2")
	_ = InsertElement(root, subSegment1)
	next := InsertElement(root, subSegment2)
	if next != subSegment2 {
		t.Errorf("must return passed segment")
	}
	if len(root.descendants) != 2 {
		t.Error("root must contain inserted segments")
	}
	if root.descendants[0] != subSegment1 && root.descendants[1] != subSegment2 {
		t.Error("inserted segments should be in root descendants")
	}
}

func TestInsertEqualElementsDummy(t *testing.T) {
	root := NewSegment("")
	subSegment1 := NewSegment("test1")
	subSegment2 := NewSegment("test1")
	_ = InsertElement(root, subSegment1)
	next := InsertElement(root, subSegment2)
	if next != subSegment1 {
		t.Errorf("must return first passed segment")
	}
	if len(root.descendants) != 1 {
		t.Error("root must contain one inserted segment")
	}
	if root.descendants[0] != subSegment1 {
		t.Error("inserted segments should be in root descendants")
	}
}

func TestInsertTwoLevelsElement(t *testing.T) {
	root := NewSegment("")
	subSegment1 := NewSegment("test1")
	subSegment2 := NewSegment("test12")
	level1 := InsertElement(root, subSegment1)
	level2 := InsertElement(level1, subSegment2)
	if level1 != subSegment1 && level2 != subSegment2 {
		t.Errorf("inserted segments shoud be returned from func")
	}
	if len(root.descendants) != 1 && len(level1.descendants) != 1 {
		t.Error(" must contain one inserted segment")
	}
	if level1.descendants[0] != level2 {
		t.Errorf("shoud be inserted at second level")
	}
}

func TestNewSegment(t *testing.T) {
	root := NewSegment("")
	if root.value != nil {
		t.Error("default segment value nil")
	}
	if root.descendants != nil {
		t.Error("default segment descendants nil")
	}
	if root.name != "" {
		t.Error("name should be set")
	}
}

func TestSegmentMethodSetName(t *testing.T) {
	name := "test"
	root := NewSegment("")
	root.SetName(name)
	if name != root.name {
		t.Error("names should be equal")
	}
}

func TestSegmentMethodSetValue(t *testing.T) {
	value := "[0-9]*"
	root := NewSegment("")
	root.SetValue(value)
	if regexp.MustCompile(value).String() != root.value.String() {
		t.Error("values should be equal")
	}
}

func TestSegmentMethodSetHandler(t *testing.T) {
	url := "http://example.com/test"
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != url {
			t.Error("url should be same")
		}
		_, _ = w.Write([]byte("test"))
	}
	root := NewSegment("")
	root.SetHandler(handler)
	root.Handle(w, req)
	if string(w.Body.Bytes()) != "test" {
		t.Error("func handler should be used")
	}
}

func TestSegmentEq(t *testing.T) {
	// TODO
	_ = NewSegment("")
}

func TestSegmentMatch(t *testing.T) {
	// TODO
	_ = NewSegment("")
}

func TestFindElement(t *testing.T) {
	// TODO
	_ = NewSegment("")
}
