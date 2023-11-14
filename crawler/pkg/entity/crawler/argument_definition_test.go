package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgumentDefinition(t *testing.T) {
	def := ArgumentDefinition{}
	def["a"] = 1
	def["b"] = "hoge"
	def["c"] = []map[string]string{
		{"k001": "v001"},
	}
	assert.Equal(t, 1, DefaultGetFromArgumentDefinition[int](def, "a", -1))
	assert.Equal(t, "hoge", DefaultGetFromArgumentDefinition[string](def, "b", "foo"))
	assert.Equal(t, "foo", DefaultGetFromArgumentDefinition[string](def, "no key", "foo")) // no key
	assert.Equal(t, 2, DefaultGetFromArgumentDefinition[int](def, "b", 2))                 // Cannot cast
	assert.Equal(t, []map[string]string{
		{"k001": "v001"},
	}, DefaultGetFromArgumentDefinition[[]map[string]string](def, "c", nil))
}

func TestNoKeyInAgumentDefinitionError(t *testing.T) {
	err := NoKeyInAgumentDefinitionError{
		Key: "foo",
	}
	assert.Equal(t, err.Error(), "Key 'foo' is not found in AgumentDefinition")
}

func TestCannotCastValueInAgumentDefinitionError(t *testing.T) {
	err := CannotCastValueInAgumentDefinitionError{
		Key: "foo",
		V:   1,
	}
	assert.Equal(t, err.Error(), "Value of 'foo' cannot be cast into type 'int' in AgumentDefinition")
}
