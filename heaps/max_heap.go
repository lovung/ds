package heaps

import "golang.org/x/exp/constraints"

// MaxHeap is a min-heap of ints.
type MaxHeap[T constraints.Ordered] []T

func (h MaxHeap[T]) Len() int           { return len(h) }
func (h MaxHeap[T]) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push to heap
func (h *MaxHeap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(T))
}

// Pop from heap
func (h *MaxHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap[T]) Index(e T) int {
	for i, v := range h {
		if v == e {
			return i
		}
	}
	return -1
}
