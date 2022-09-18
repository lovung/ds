package sorts

import (
	"golang.org/x/exp/constraints"
)

// QuickSort a Slice of ordered types
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
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	i++
	slice[i], slice[hi] = slice[hi], slice[i]
	return i
}

// QuickSort a Slice of ordered types
func QuickSortSlice[T any](slice []T, less func(i, j int) bool) {
	quickSortSlice(slice, 0, len(slice)-1, less)
}

func quickSortSlice[T any](slice []T, lo, hi int, less func(i, j int) bool) {
	if lo >= hi || lo < 0 {
		return
	}
	p := partitionSlice(slice, lo, hi, less)
	quickSortSlice(slice, lo, p-1, less)
	quickSortSlice(slice, p+1, hi, less)
}

func partitionSlice[T any](slice []T, lo, hi int, less func(i, j int) bool) int {
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		if less(j, hi) {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	i++
	slice[i], slice[hi] = slice[hi], slice[i]
	return i
}
