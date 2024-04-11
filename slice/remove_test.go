package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, []int{}, RemoveHead(arr, 11))
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, RemoveHead(arr, 1))
	assert.Equal(t, []int{}, RemoveTail(arr, 11))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, RemoveTail(arr, 5))
	arr = RemoveAt(arr, 4)
	assert.Equal(t, []int{1, 2, 3, 4, 6, 7, 8, 9, 10}, arr)
}
