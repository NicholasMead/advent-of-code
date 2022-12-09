package coord

import (
	"math"
)

type Coord struct {
	x, y int
}

func (t Coord) StepTowards(h Coord) Coord {
	dx := h.x - t.x
	dy := h.y - t.y

	if abs(dx) >= 2 || abs(dy) >= 2 {
		switch {
		case dx > 0:
			t.x++
		case dx < 0:
			t.x--
		}

		switch {
		case dy > 0:
			t.y++
		case dy < 0:
			t.y--
		}
	}

	return t
}

func (t Coord) StepDirection(direction rune) Coord {
	switch direction {
	case 'U':
		t.y++
	case 'D':
		t.y--
	case 'L':
		t.x--
	case 'R':
		t.x++
	}

	return t
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}
