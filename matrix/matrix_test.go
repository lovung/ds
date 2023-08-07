package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Matrix(t *testing.T) {
	t.Run("not-square", func(t *testing.T) {
		mat := New(3, 3, 1)
		assert.Equal(t, 1, mat.At(0, 2))
		mat.UpdateAt(0, 2, -1)
		assert.Equal(t, -1, mat.At(0, 2))
		assert.Equal(t, 1, mat.At(-1, 2))
		mat.UpdateAt(-1, 2, 0)
		mat.UpdateAt(1, -2, 0)
		sum := 0
		mat.ForEach(func(val int, _, _ int) {
			sum += val
		})
		assert.Equal(t, 7, sum)
		assert.Equal(t, -1, mat.Data()[0][2])
	})

	t.Run("square", func(t *testing.T) {
		mat := NewSquare(3, 1)
		assert.Equal(t, 1, mat.At(0, 2))
		mat.UpdateAt(0, 2, -1)
		assert.Equal(t, -1, mat.At(0, 2))
		assert.Equal(t, 1, mat.At(-1, 2))
		mat.UpdateAt(-1, 2, 0)
		mat.UpdateAt(1, -2, 0)
		sum := 0
		mat.ForEach(func(val int, _, _ int) {
			sum += val
		})
		assert.Equal(t, 7, sum)
		assert.Equal(t, -1, mat.Data()[0][2])
		mat.DoWithNearBy(0, 0, func(_ int, i, j int) { mat.UpdateAt(i, j, 0) })
		sum = 0
		mat.ForEach(func(val int, _, _ int) {
			sum += val
		})
		assert.Equal(t, 5, sum)
	})
}
