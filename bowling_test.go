package gobowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGutterGame(t *testing.T) {
	g := NewGame()
	rollMany(&g, 20, 0)
	assert.Equal(t, 0, g.CalculateScore(), "got 0 for gutter game")
}

func TestOnePinGame(t *testing.T) {
	g := NewGame()
	g.Roll(1)
	rollMany(&g, 19, 0)
	assert.Equal(t, 1, g.CalculateScore(), "got 1 for one pin game")
}

func TestSpare(t *testing.T) {
	g := NewGame()
	g.Roll(5)
	g.Roll(5)
	g.Roll(3)
	rollMany(&g, 17, 0)
	assert.Equal(t, 16, g.CalculateScore(), "got 16 for spare game")
}

func TestStrike(t *testing.T) {
	g := NewGame()
	g.Roll(10)
	g.Roll(6)
	g.Roll(2)
	rollMany(&g, 17, 0)
	assert.Equal(t, 26, g.CalculateScore(), "got 26 for strike game")
}

func TestPerfectGame(t *testing.T) {
	g := NewGame()
	rollMany(&g, 12, 10)
	assert.Equal(t, 300, g.CalculateScore(), "got 300 for perfect game")
}

func rollMany(g *Game, n int, pins int) {
	for i := 0; i < n; i++ {
		g.Roll(pins)
	}
}
