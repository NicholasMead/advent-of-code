package coord

import (
	"fmt"
	"math"
)

type Coord struct {
	x, y int
}

type Direction rune

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

func (t Coord) StepTowards(h Coord) Coord {
	dx := h.x - t.x
	dy := h.y - t.y

	if abs(dx) >= 2 || abs(dy) >= 2 {
		t.x += sign(dx)
		t.y += sign(dy)
	}

	return t
}

func (t Coord) StepDirection(dir Direction) Coord {
	switch dir {
	case Up:
		t.y++
	case Down:
		t.y--
	case Left:
		t.x--
	case Right:
		t.x++
	default:
		panic(fmt.Sprintf("Unknown direction %c", dir))
	}

	return t
}

func abs(v int) int {
	//I'm a little upset with the math library for needing this statement.
	//Integers need absolute values too!
	return int(math.Abs(float64(v)))
}

func sign(v int) int {
	if v > 0 {
		return 1
	} else if v < 0 {
		return -1
	} else {
		return 0
	}
}
