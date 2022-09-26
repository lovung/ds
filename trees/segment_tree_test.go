package trees

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmentTreeWithArray(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		st := NewSegmentTreeWithArray(10, func(a, b int) int {
			if a > b {
				return a
			}
			return b
		})
		assert.Zero(t, st.Query(0, 10))
		st.Update(5, 10)
		assert.Equal(t, 10, st.Query(0, 10))
	})
	t.Run("min", func(t *testing.T) {
		st := NewSegmentTreeWithArray(10, func(a, b int) int {
			if a < b {
				return a
			}
			return b
		})
		st.SetInitQueryValue(math.MaxInt64)
		assert.Zero(t, st.Query(0, 10))
		for i := 0; i < 10; i++ {
			st.Update(i, i+100)
		}
		assert.Equal(t, 100, st.Query(0, 10))
	})
	t.Run("sum", func(t *testing.T) {
		st := NewSegmentTreeWithArray(10, func(a, b int) int {
			return a + b
		})
		assert.Zero(t, st.Query(0, 10))
		st.Update(5, 10)
		assert.Equal(t, 10, st.Query(0, 10))
		st.Update(6, 10)
		assert.Equal(t, 10, st.Query(5, 6))
		assert.Equal(t, 20, st.Query(5, 7))
	})
}

func TestSegmentTreeWithMap(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		st := NewSegmentTreeWithMap(10, func(a, b int) int {
			if a > b {
				return a
			}
			return b
		})
		assert.Zero(t, st.Query(0, 10))
		st.Update(5, 10)
		assert.Equal(t, 10, st.Query(0, 10))
	})
	t.Run("min", func(t *testing.T) {
		st := NewSegmentTreeWithMap(10, func(a, b int) int {
			if a < b {
				return a
			}
			return b
		})
		st.SetInitQueryValue(math.MaxInt64)
		assert.Zero(t, st.Query(0, 10))
		for i := 0; i < 10; i++ {
			st.Update(i, i+100)
		}
		assert.Equal(t, 100, st.Query(0, 10))
	})
	t.Run("sum", func(t *testing.T) {
		st := NewSegmentTreeWithMap(10, func(a, b int) int {
			return a + b
		})
		assert.Zero(t, st.Query(0, 10))
		st.Update(5, 10)
		assert.Equal(t, 10, st.Query(0, 10))
		st.Update(6, 10)
		assert.Equal(t, 10, st.Query(5, 6))
		assert.Equal(t, 20, st.Query(5, 7))
	})
}
