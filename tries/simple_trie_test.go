package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTrie(t *testing.T) {
	t.Run("apple", func(t *testing.T) {
		trie := NewSimpleTrie()
		trie.Insert("apple")
		assert.True(t, trie.Search("apple"))
		assert.False(t, trie.Search("app"))
		assert.True(t, trie.StartsWith("app"))
		trie.Insert("app")
		assert.True(t, trie.Search("app"))
		assert.False(t, trie.Search("applepay"))
	})
	t.Run("a", func(t *testing.T) {
		trie := NewSimpleTrie()
		assert.False(t, trie.StartsWith("a"))
	})
}
