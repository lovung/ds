package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSlices(t *testing.T) {
	t.Run("normal-slice", func(t *testing.T) {
		s := []int{2, 3, 4, 1, 6, 5, 8, 9}
		OrderedSlice(s)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9}, s)
	})

	t.Run("struct-slice", func(t *testing.T) {
		type Person struct {
			Name string
		}
		s := []*Person{
			{Name: "Bobby"},
			{Name: "Andy"},
		}
		SliceFn(s, func(p1, p2 *Person) bool { return p1.Name < p2.Name })
		assert.Equal(t, []*Person{
			{Name: "Andy"},
			{Name: "Bobby"},
		}, s)
	})
}
