package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum(1, 2, 3, 4))
	assert.Equal(t, "abcd", Sum("a", "b", "c", "d"))
}

func TestGCD(t *testing.T) {
	assert.Equal(t, 10, GCD(10, 20))
	assert.Equal(t, 5, GCD(15, 20))
	assert.Equal(t, 15, GCD(15, 0))
	assert.Equal(t, 4, GCD(12, 16))
}
