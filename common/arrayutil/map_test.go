package arrayutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(
		t,
		[]int{102, 103, 104},
		Map([]int{101, 102, 103}, func(v int) int { return v + 1 }),
	)
}
