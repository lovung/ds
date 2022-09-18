package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	nums := []int{0, -1, 2, -3, 4, 5}
	QuickSort(nums)
	assert.Equal(t, []int{-3, -1, 0, 2, 4, 5}, nums)
}

func TestQuickSortSlice(t *testing.T) {
	type A struct {
		age int
	}

	As := []*A{
		{3},
		{1},
		{2},
		{5},
		{4},
	}
	QuickSortSlice(As, func(i, j int) bool {
		return As[i].age < As[j].age
	})
	assert.Equal(t, []*A{
		{1},
		{2},
		{3},
		{4},
		{5},
	}, As)
}
