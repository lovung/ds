package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	assert.Equal(t, 1, uf.Find(1))
	assert.Equal(t, 9, uf.Find(9))

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 4)
	uf.Union(4, 1) // same set case
	assert.Equal(t, 0, uf.Find(4))

	uf.Union(9, 0) // lower rank case
	assert.Equal(t, 0, uf.Find(9))

	uf.Union(6, 5)
	uf.Union(5, 7)
	assert.Equal(t, 6, uf.Find(7))

	uf.Union(6, 2)
	assert.Equal(t, 6, uf.Find(1))
	assert.Equal(t, 6, uf.Find(4))
}
