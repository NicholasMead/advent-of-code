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
	if dx, dy := diff(t, h); abs(dx) >= 2 || abs(dy) >= 2 {
		t.x += sign(dx)
		t.y += sign(dy)
	}

	return t
}

func (t Coord) StepDirection(dir Direction) Coord {
	move := map[Direction]Coord{
		Up:    {0, 1},
		Down:  {0, -1},
		Left:  {-1, 0},
		Right: {1, 0},
	}

	if m, found := move[dir]; !found {
		panic(fmt.Sprintf("Unknown direction %c", dir))
	} else {
		t.x += m.x
		t.y += m.y
	}

	return t
}

func diff(a Coord, b Coord) (dx, dy int) {
	return b.x - a.x, b.y - a.y
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
