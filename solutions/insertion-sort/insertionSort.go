package insertionSort

import (
	"github.com/jkufa/CLRS-solutions-golang/packages/util"
)

// Input: a sequence of n numbers <a1, a2, ..., an>
// Output: a permutation <a1', a2', ..., an'> of the input sequence such that a1' <= a2' <= ... <= an'
func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j -= 1
		}
		a[j+1] = key
	}
	return a
}

func main() {
	a := util.GenerateRandomArray(100000)
	defer util.Timer("insertionSort")()
	insertionSort(a)
}
