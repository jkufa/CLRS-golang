package linearsearch

// Input: a sequence of n numbers <a1, a2, ..., an> stored in array A[1..n] and a value x
// Output: an index i such that x = A[i] or the special value NIL if x does not appear in A

func linearSearch(a []int, x int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return i
		}
	}
	return -1
}
