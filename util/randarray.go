package util

import (
	"math/rand"
)

func GenerateRandomArray(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n)
	}
	return a
}
