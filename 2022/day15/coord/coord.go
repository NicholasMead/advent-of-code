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
		t.X += m.X
		t.Y += m.Y
	}

	return t
}

func Dist(a, b Coord) (dist int) {
	dx, dy := Diff(a, b)
	return abs(dx) + abs(dy)
}

func Diff(a, b Coord) (dx, dy int) {
	return b.X - a.X, b.Y - a.Y
}

// func (c Coord) GetTrace(dist int) []Coord {
// 	trace := map[Coord]interface{}{}

// 	start := Coord{c.X, c.Y + dist + 1}
// 	pos := start
// 	for {
// 		fmt.Println(pos)
// 		switch {
// 		case pos.X >= c.X && pos.Y > c.Y:
// 			pos = pos.StepDirection(Right)
// 			trace[pos] = new(interface{})
// 			pos = pos.StepDirection(Down)
// 			trace[pos] = new(interface{})

// 		case pos.X > c.X && pos.Y <= c.Y:
// 			pos = pos.StepDirection(Down)
// 			trace[pos] = new(interface{})
// 			pos = pos.StepDirection(Left)
// 			trace[pos] = new(interface{})

// 		case pos.X < c.X && pos.Y < c.Y:
// 			pos = pos.StepDirection(Left)
// 			trace[pos] = new(interface{})
// 			pos = pos.StepDirection(Up)
// 			trace[pos] = new(interface{})

// 		case pos.X < c.X && pos.Y >= c.Y:
// 			pos = pos.StepDirection(Up)
// 			trace[pos] = new(interface{})
// 			pos = pos.StepDirection(Right)
// 			trace[pos] = new(interface{})
// 		}

// 		if pos == start {
// 			break
// 		}
// 	}

// 	keys := make([]Coord, 0, len(trace))
// 	for c2 := range trace {
// 		keys = append(keys, c2)
// 	}
// 	return keys
// }

// func (c Coord) GetTrace(dist int) []Coord {
// 	x, y := c.X, c.Y
// 	trace := map[Coord]struct{}{}

// 	for i := 1; i <= dist; i++ {
// 		trace[Coord{x + i, y + (dist - i)}] = struct{}{}
// 		trace[Coord{x + i, y + (dist - i - 1)}] = struct{}{}

// 		trace[Coord{x - i, y + (dist - i)}] = struct{}{}
// 		trace[Coord{x - i, y + (dist - i - 1)}] = struct{}{}

// 		trace[Coord{x + i, y - (dist - i)}] = struct{}{}
// 		trace[Coord{x + i, y - (dist - i - 1)}] = struct{}{}

// 		trace[Coord{x - i, y - (dist - i)}] = struct{}{}
// 		trace[Coord{x - i, y - (dist - i - 1)}] = struct{}{}
// 	}
// 	trace[Coord{x + dist, y}] = struct{}{}
// 	trace[Coord{x - dist, y}] = struct{}{}

// 	unique := make([]Coord, 0, len(trace))
// 	for c2 := range trace {
// 		unique = append(unique, c2)
// 	}
// 	return unique
// }

func (c Coord) GetTrace(dist int) []Coord {
	trace := []Coord{}

	posT := Coord{c.X, c.Y + dist + 1}
	posB := Coord{c.X, c.Y - dist - 1}
	posR := Coord{c.X + dist + 1, c.Y}
	posL := Coord{c.X - dist - 1, c.Y}

	for x := 0; x <= dist; x++ {
		trace = append(trace, posT)
		trace = append(trace, posR)
		trace = append(trace, posB)
		trace = append(trace, posL)

		posT = posT.StepDirection(Right)
		posR = posR.StepDirection(Down)
		posB = posB.StepDirection(Left)
		posL = posL.StepDirection(Up)

		trace = append(trace, posT)
		trace = append(trace, posR)
		trace = append(trace, posB)
		trace = append(trace, posL)

		posT = posT.StepDirection(Down)
		posR = posR.StepDirection(Left)
		posB = posB.StepDirection(Up)
		posL = posL.StepDirection(Right)
	}

	return trace
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
