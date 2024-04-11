package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFront(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3}, PushFront([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{1}, PushFront([]int{}, 1))
	assert.Equal(t, []int{1}, PushFront(nil, 1))
}
