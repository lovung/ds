package heaps

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	var h PriorityQueue[string]
	heap.Init(&h)
	heap.Push(&h, &PrioritItem[string]{Value: "Hello", Priority: 1})
	w := PrioritItem[string]{Value: "World", Priority: 2}
	heap.Push(&h, &w)
	heap.Push(&h, &PrioritItem[string]{Value: "Golang", Priority: 3})
	assert.Equal(t, 3, h.Len())
	assert.Equal(t, "Golang", heap.Pop(&h).(*PrioritItem[string]).Value)
	h.Update(&w, "World2", 0)
	assert.Equal(t, "Hello", heap.Pop(&h).(*PrioritItem[string]).Value)
	assert.Equal(t, 1, h.Len())
	assert.Equal(t, "World2", heap.Pop(&h).(*PrioritItem[string]).Value)
}
