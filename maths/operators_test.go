package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum(1, 2, 3, 4))
	assert.Equal(t, "abcd", Sum("a", "b", "c", "d"))
}
