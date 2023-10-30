package arrayutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	assert.Equal(
		t,
		[]int{1, 2, 3, 4, 5},
		Uniq([]int{1, 2, 3, 3, 4, 5}),
	)
	assert.Equal(
		t,
		[]string{"1", "2", "3", "4", "5"},
		Uniq([]string{"1", "2", "3", "3", "4", "5"}),
	)
}
