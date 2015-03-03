package stack

import "testing"

func TestLength(t *testing.T) {
	stack := NewStack()
	if stack.Length() != 0 {
		t.Errorf("Error, length should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := NewStack()
	if stack.IsEmpty() != true {
		t.Errorf("Error, should be true")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(2)
	if stack.Length() != 1 {
		t.Errorf("Error, length should be 1")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(2)
	stack.Push(1)
	if x := stack.Pop(); x != 1 {
		t.Errorf("Error, should be 1, got %d", x)
	}
	if stack.Length() != 1 {
		t.Errorf("Error, should be 1, got %d", stack.Length())
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(2)
	stack.Push(1)
	if x := stack.Peek(); x != 1 {
		t.Errorf("Error, should be 1")
	}
	if stack.Length() != 2 {
		t.Errorf("Error, should just get element")
	}
}
