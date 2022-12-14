package coord

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func FromString(s string) Coord {
	segments := strings.Split(s, ",")
	x, _ := strconv.Atoi(segments[0])
	y, _ := strconv.Atoi(segments[1])
	return Coord{x, y}
}

type Direction rune

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

func (t Coord) StepTowards(h Coord) Coord {
	if dx, dy := diff(t, h); abs(dx) >= 1 || abs(dy) >= 1 {
		t.X += sign(dx)
		t.Y += sign(dy)
	}

	return t
}

func (t Coord) StepDirection(dir Direction) Coord {
	move := map[Direction]Coord{
		Up:    {0, -1},
		Down:  {0, 1},
		Left:  {-1, 0},
		Right: {1, 0},
	}

	if m, found := move[dir]; !found {
		panic(fmt.Sprintf("Unknown direction %c", dir))
	} else {
		t.X += m.X
		t.Y += m.Y
	}

	return t
}

func diff(a Coord, b Coord) (dx, dy int) {
	return b.X - a.X, b.Y - a.Y
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
