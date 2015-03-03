package list

import "testing"

func TestLength(t *testing.T) {
	list := NewList()
	if list.Length() != 0 {
		t.Error("Length should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewList()
	if list.IsEmpty() != true {
		t.Error("Should be empty")
	}
}

func TestPrepend(t *testing.T) {
	list := NewList()
	list.Prepend(10)
	if list.Head.Value != 10 {
		t.Error("Head should now be 10")
	}
	list.Prepend(9)
	if list.Head.Value != 9 {
		t.Error("Head should now be 9")
	}
}

func TestAppend(t *testing.T) {
	list := NewList()
	list.Append(1)
	if list.Head.Value != 1 {
		t.Error("Head should be 1")
	}
	list.Append(2)
	if list.Head.Value != 1 {
		t.Error("Head should continue to be 1")
	}
}
