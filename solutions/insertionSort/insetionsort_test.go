package insertionsort

import (
	"testing"
)

func TestInsertionSort_Common(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	unsorted := []int{5, 4, 3, 2, 1}
	a := insertionSort(unsorted)
	for i := 0; i < len(a); i++ {
		if a[i] != sorted[i] {
			t.Errorf("insertionSort(%v) = %v; want %v", unsorted, a, sorted)
		}
	}
}

func TestInsertionSort_EmptyArray(t *testing.T) {
	a := []int{}
	sorted := []int{}
	a = insertionSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != sorted[i] {
			t.Errorf("insertionSort(%v) = %v; want %v", a, a, sorted)
		}
	}
}

func TestInsertionSort_ArraySizeOfOne(t *testing.T) {
	a := []int{1}
	sorted := []int{1}
	a = insertionSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != sorted[i] {
			t.Errorf("insertionSort(%v) = %v; want %v", a, a, sorted)
		}
	}
}
