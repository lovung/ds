package heaps

// This example demonstrates a priority queue built using the heap interface.

import (
	"container/heap"
)

// An PrioritItem is something we manage in a priority queue.
type PrioritItem[T any] struct {
	Value    T   // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*PrioritItem[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push a new item into the priority queue
func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PrioritItem[T])
	item.index = n
	*pq = append(*pq, item)
}

// Pop an item from the priority queue
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an PrioritItem in the queue.
func (pq *PriorityQueue[T]) Update(item *PrioritItem[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
