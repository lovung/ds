package heaps

type ToInt interface {
	Int() int
}

type MinHeapInt[T ToInt] []T

func (h MinHeapInt[T]) Len() int           { return len(h) }
func (h MinHeapInt[T]) Less(i, j int) bool { return h[i].Int() < h[j].Int() }
func (h MinHeapInt[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push to heap
func (h *MinHeapInt[T]) Push(x any) {
	*h = append(*h, x.(T))
}

// Pop from heap
func (h *MinHeapInt[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
