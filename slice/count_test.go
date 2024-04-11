package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert.Equal(t, 3, Count([]int{1, 2, 3, 4, 5, 3, 3}, 3))
	assert.Equal(t, 0, Count([]int{1, 2, 3, 4, 5, 3, 3}, 0))
}

func TestCountIf(t *testing.T) {
	assert.Equal(t, 2, CountIf([]int{1, 2, 3, 4, 5, 3, 3}, func(a int) bool {
		return a%2 == 0
	}))
	assert.Equal(t, 5, CountIf([]int{1, 2, 3, 4, 5, 3, 3}, func(a int) bool {
		return a%2 != 0
	}))
}
