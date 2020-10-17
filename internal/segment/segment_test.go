package segment

import (
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
