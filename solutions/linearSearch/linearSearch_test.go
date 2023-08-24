package linearsearch

import (
	"testing"
)

func TestLinearSearch_Common(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	x := 3
	i := linearSearch(a, x)
	if i != 2 {
		t.Errorf("linearSearch(%v, %v) = %v; want %v", a, x, i, 2)
	}
}

func TestLinearSearch_EmptyArray(t *testing.T) {
	a := []int{}
	x := 3
	i := linearSearch(a, x)
	if i != -1 {
		t.Errorf("linearSearch(%v, %v) = %v; want %v", a, x, i, -1)
	}
}

func TestLinearSearch_ArraySizeOfOne(t *testing.T) {
	a := []int{1}
	x := 1
	i := linearSearch(a, x)
	if i != 0 {
		t.Errorf("linearSearch(%v, %v) = %v; want %v", a, x, i, 0)
	}
}
