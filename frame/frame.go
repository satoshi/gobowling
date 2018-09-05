package frame

type Frame struct {
	pins []int
}

func NewFrame(pins []int) Frame {
	f := Frame{}
	f.pins = pins
	return f
}

func (f *Frame) CalculateScore() int {
	first_roll := f.pins[0]
	second_roll := f.pins[1]

	if f.isStrike() {
		return 10 + second_roll + f.pins[2]
	} else if f.isSpare() {
		return 10 + f.pins[2]
	}
	return first_roll + second_roll
}

func (f *Frame) CalculateOffset() int {
	if f.isStrike() {
		return 1
	}
	return 2
}

func (f *Frame) isSpare() bool {
	if f.pins[0]+f.pins[1] == 10 {
		return true
	}
	return false
}

func (f *Frame) isStrike() bool {
	if f.pins[0] == 10 {
		return true
	}
	return false
}
