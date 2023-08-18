package insertionSort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	unsorted := []int{5, 4, 3, 2, 1}
	a := insertionSort(unsorted)
	for i := 0; i < len(a); i++ {
		if a[i] != sorted[i] {
			t.Errorf("insertionSort(%v) = %v; want %v", unsorted, a, sorted)
		}
	}
}
