package mergesort

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	var (
		i, j int
	)
	a := []int{}
	for i < leftLen && j < rightLen {
		if left[i] < right[j] {
			a = append(a, left[i])
			i += 1
		} else {
			a = append(a, right[j])
			j += 1
		}
	}
	a = appendToEnd(a, left, i)
	a = appendToEnd(a, right, j)
	return a
}

func appendToEnd(a1, a2 []int, index int) []int {
	for index < len(a2) {
		a1 = append(a1, a2[index])
		index += 1
	}
	return a1
}
