package queue

import "testing"

func TestLength(t *testing.T) {
	queue := NewQueue()
	if queue.Length() != 0 {
		t.Error("Length should be 0")
	}
}

func TestIsEmpty(t *testing.T) {
	queue := NewQueue()
	if queue.IsEmpty() != true {
		t.Error("Should be empty")
	}
}

func TestPush(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	if queue.Length() != 2 {
		t.Error("Length should be 2")
	}
}

func TestPop(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	if x := queue.Pop(); x != 1 {
		t.Errorf("Error, should be 1, got %d", x)
	}
	if queue.Length() != 1 {
		t.Error("Length should be 1")
	}
}

func TestPeek(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	if x := queue.Peek(); x != 1 {
		t.Errorf("Error, should be 1, got %d", x)
	}
	if queue.Length() != 2 {
		t.Errorf("Length should be 2")
	}
}
