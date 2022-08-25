package sorts

import (
	"golang.org/x/exp/constraints"
)

// QuickSort a Slice of interger
func QuickSort[T constraints.Ordered](slice []T) {
	quickSort(slice, 0, len(slice)-1)
}

func quickSort[T constraints.Ordered](slice []T, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := partition(slice, lo, hi)
	quickSort(slice, lo, p-1)
	quickSort(slice, p+1, hi)
}

func partition[T constraints.Ordered](slice []T, lo, hi int) int {
	pivot := slice[hi]
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	i++
	slice[i], slice[hi] = slice[hi], slice[i]
	return i
}
