package frame

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalFrame(t *testing.T) {
	f := NewFrame([]int{3, 3})
	assert.Equal(t, 6, f.CalculateScore(), "got 6 for normal frame")
	assert.Equal(t, 2, f.CalculateOffset(), "got 2 for offset")
}

func TestSpareFrame(t *testing.T) {
	f := NewFrame([]int{5, 5, 7})
	assert.Equal(t, 17, f.CalculateScore(), "got 17 for spare frame")
	assert.Equal(t, 2, f.CalculateOffset(), "got 2 for offset")
}

func TestStrikeFrame(t *testing.T) {
	f := NewFrame([]int{10, 3, 5})
	assert.Equal(t, 18, f.CalculateScore(), "got 18 for strike frame")
	assert.Equal(t, 1, f.CalculateOffset(), "got 1 for offset")
}
