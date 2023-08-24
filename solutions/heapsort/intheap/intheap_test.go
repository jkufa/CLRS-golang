package intheap

import (
	"testing"
)

func TestIntHeap_Len(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	if h.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", h.Len())
	}
}

func TestIntHeap_LessMinHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	if h.Less(0, 1) {
		t.Errorf("Expected 2 to be less than 1 (true), got %t", h.Less(0, 1))
	}
}

func TestIntHeap_LessMaxHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	if !h.Less(1, 0) {
		t.Errorf("Expected 1 to be less than 2 (false), got %t", h.Less(1, 0))
	}
}

func TestIntHeap_Swap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	h.Swap(0, 1)
	if h.At(0) != 1 || h.At(1) != 2 {
		t.Errorf("Expected 1, 2, 5, got %d, %d, %d", h.At(0), h.At(1), h.At(2))
	}
}

func TestIntHeap_Push(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	h.Push(3)
	if h.At(3) != 3 {
		t.Errorf("Expected 3, got %d", h.At(3))
	}
}

func TestIntHeap_Pop(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	popped := h.Pop()
	if h.Len() != 2 {
		t.Errorf("Expected length of 2, got %d", h.Len())
	}
	if h.At(0) != 2 {
		t.Errorf("Expected 2, got %d", h.At(0))
	}
	if popped != 5 {
		t.Errorf("Expected 5, got %d", popped)
	}
}
