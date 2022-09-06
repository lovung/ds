package heaps

import "golang.org/x/exp/constraints"

// An MaxHeapWithValue is a max-heap of ints.
type MaxHeapWithValue[T constraints.Ordered] []*HeapItem[T]

func (h MaxHeapWithValue[T]) Len() int           { return len(h) }
func (h MaxHeapWithValue[T]) Less(i, j int) bool { return h[i].Ref > h[j].Ref }
func (h MaxHeapWithValue[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Push to heap
func (h *MaxHeapWithValue[T]) Push(x any) {
	n := len(*h)
	item := x.(*HeapItem[T])
	item.index = n
	*h = append(*h, item)
}

// Pop from heap
func (h *MaxHeapWithValue[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
