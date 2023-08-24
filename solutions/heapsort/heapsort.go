package heapsort

/**
TODO:
- Tests
- Implement heap data structure
- Implement max-heapify
- Implement build-max-heap
- Implement heap-sort
*/

func setLargest(a []int, i int) int {
	l := (i << 1) + 1
	r := (i << 1) + 2
	largest := i
	if l < len(a) && a[l] > a[i] {
		largest = l
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	return largest
}

func maxHeapify(a []int, i int) {
	largest := setLargest(a, i)
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest)
	}
}

func buildMaxHeap(a []int) {
	for i := (len(a) - 1) / 2; i >= 0; i-- {
		maxHeapify(a, i)
	}
}

func heapSort(a []int) {
	buildMaxHeap(a)
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		maxHeapify(a[:i], 0)
	}
}
