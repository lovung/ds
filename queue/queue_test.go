package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewSimpleQueue[int]()
	q.EnQueue(1)
	assert.Equal(t, 1, q.Len())
	v, _ := q.DeQueue()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, q.Len())
}

func Test_SimpleQueue(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		q := NewSimpleQueue[int]()
		_, ok := q.DeQueue()
		assert.False(t, ok)
		assert.True(t, q.EnQueue(1))
		assert.True(t, q.EnQueue(2))
		assert.True(t, q.EnQueue(3))
		assert.Equal(t, 3, q.Rear())
		assert.False(t, q.IsFull())
		_, ok = q.DeQueue()
		assert.True(t, ok)
		assert.True(t, q.EnQueue(4))
		assert.Equal(t, 4, q.Rear())
	})
	t.Run("2", func(t *testing.T) {
		q := NewSimpleQueue[int]()
		assert.True(t, q.EnQueue(6))
		assert.Equal(t, 6, q.Rear())
		v, ok := q.DeQueue()
		assert.Equal(t, 6, v)
		assert.True(t, ok)
		assert.True(t, q.EnQueue(5))
		assert.Equal(t, 5, q.Rear())
		v, ok = q.DeQueue()
		assert.Equal(t, 5, v)
		assert.True(t, ok)
		assert.Equal(t, 0, q.Front())
	})
	t.Run("3", func(t *testing.T) {
		q := NewSimpleQueue[int]()
		assert.Equal(t, 0, q.Rear())
		assert.True(t, q.EnQueue(1))
		q.DeQueue()
		assert.True(t, q.EnQueue(2))
		q.DeQueue()
		assert.True(t, q.EnQueue(3))
		q.DeQueue()
		assert.True(t, q.EnQueue(4))
		assert.Equal(t, 4, q.Rear())
		assert.Equal(t, 4, q.Front())
	})
}
