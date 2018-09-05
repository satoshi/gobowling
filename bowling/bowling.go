package bowling

import (
	"github.com/satoshi/go-bowling/frame"
)

type Game struct {
	rolls []int
}

func NewGame() Game {
	g := Game{}
	return g
}

func (g *Game) Roll(pins int) {
	g.rolls = append(g.rolls, pins)
}

func (g *Game) CalculateScore() int {
	score := 0
	current_roll := 0

	for i := 0; i < 10; i++ {
		var temp_array []int

		for j := current_roll; j < current_roll+3; j++ {
			if j < len(g.rolls) {
				temp_array = append(temp_array, g.rolls[j])
			}
		}
		f := frame.NewFrame(temp_array)
		score += f.CalculateScore()
		current_roll += f.CalculateOffset()
	}
	return score
}

func (g *Game) isSpare(first_roll, second_roll int) bool {
	if first_roll+second_roll == 10 {
		return true
	}
	return false
}

func (g *Game) isStrike(first_roll int) bool {
	if first_roll == 10 {
		return true
	}
	return false
}
