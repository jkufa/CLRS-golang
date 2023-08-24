package mergesort

import (
	"testing"
)

func TestMergeSort_Common(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	unsorted := []int{5, 4, 3, 2, 1}
	a := mergeSort(unsorted)
	for i := 0; i < len(a); i++ {
		if a[i] != sorted[i] {
			t.Errorf("MergeSort(%v) = %v; want %v", unsorted, a, sorted)
		}
	}
}

func TestMergeSort_ArraySizeOfOne(t *testing.T) {
	a := []int{1}
	sorted := []int{1}
	a = mergeSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != sorted[i] {
			t.Errorf("MergeSort(%v) = %v; want %v", a, a, sorted)
		}
	}
}
