package heapsort

import (
	"container/heap"
	"testing"

	"github.com/jkufa/CLRS-solutions-golang/packages/solutions/heapsort/intheap"
)

func TestHeapSort_HeapPackage(t *testing.T) {
	a := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	h := &intheap.IntHeap{}
	heap.Init(h)
	// We push in order, but this ensures the heap is sorted how we hardcoded it,
	// making assertion easier
	for i := 0; i < len(a); i++ {
		heap.Push(h, a[i])
	}
	heap.Fix(h, 0)

	// Assert that heap is sorted as a min heap
	for i := 0; i < len(a); i++ {
		if a[i] != (*h)[i] {
			t.Errorf("Expected %d, got %d", a[i], (*h)[i])
		}
	}
}

func TestHeapSort_MaxHeapify(t *testing.T) {
	a := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	maxHeapify(a, 1)
	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_MaxHeapifyInvalidInput(t *testing.T) {
	a := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	maxHeapify(a, 11) // index out of range
	expected := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_BuildMaxHeap(t *testing.T) {
	a := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	buildMaxHeap(a)
	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_BuildMaxHeapOddLength(t *testing.T) {
	a := []int{4, 1, 3, 2, 16, 9, 10, 14, 8}
	buildMaxHeap(a)
	expected := []int{16, 14, 10, 8, 1, 9, 3, 2, 4}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_BuildMaxHeapAllDuplicates(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	buildMaxHeap(a)
	expected := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_Common(t *testing.T) {
	a := []int{4, 1, 3, 2, 16, 9, 10, 14, 8}
	heapSort(a)
	expected := []int{1, 2, 3, 4, 8, 9, 10, 14, 16}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_ArraySizeOfOne(t *testing.T) {
	a := []int{1}
	heapSort(a)
	expected := []int{1}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}

func TestHeapSort_EmptyArray(t *testing.T) {
	a := []int{}
	heapSort(a)
	expected := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], a[i])
		}
	}
}
