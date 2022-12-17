package towers

import (
	"aoc/day17/cycle"
	"aoc/day17/shapes"
	"fmt"
)

type Tower interface {
	Drop(p shapes.Patten, j cycle.Cycle[int])
	AsPatten() shapes.Patten
	Copy() Tower

	Size() int
}

func CreateTower(width int) Tower {
	return &tower{
		width: width,
		fixed: map[shapes.Coord]interface{}{},
	}
}

type tower struct {
	width, height int
	base          int
	fixed         map[shapes.Coord]interface{}
}

// Copy implements Tower
func (t tower) Copy() Tower {
	fixed := map[shapes.Coord]interface{}{}
	for f, i := range t.fixed {
		fixed[f] = i
	}
	t.fixed = fixed
	return &t
}

// Reduce implements Tower
func (t *tower) reduce() {
	base := t.topCompletedRow()
	if base > t.base {
		fixed := map[shapes.Coord]interface{}{}
		for f, i := range t.fixed {
			if f.Y >= base {
				fixed[f] = i
			}
		}
		t.fixed = fixed
		t.base = base
	}
}

// Reduce implements Tower
func (t *tower) topCompletedRow() int {
	// top := math.MaxInt

	for y := t.Size(); y >= t.base; y-- {
		if t.rowCompleted(y) {
			return y
		}
	}
	return t.base
}

func (t *tower) rowCompleted(yIndex int) bool {
	scan := 3

	for x := 0; x < t.width; x++ {
		found := false
		for y := yIndex; y < yIndex+scan; y++ {

			_, found = t.fixed[shapes.Coord{X: x, Y: y}]
			if found {
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// AsPatten implements Tower
func (t *tower) AsPatten() shapes.Patten {
	p := shapes.Patten{}
	for r := range t.fixed {
		r.Y -= t.base
		p = append(p, r)
	}
	return p
}

// Add implements Tower
func (t *tower) Drop(p shapes.Patten, j cycle.Cycle[int]) {
	current := shapes.CreateRock(2, t.height+3, p)
	var candidate shapes.Rock

	for {
		//apply jet
		jet, _ := j.Next()
		candidate = current.Move(jet, 0)
		if !t.conflicts(candidate) {
			current = candidate
		}

		//apply fall
		candidate = current.Move(0, -1)
		if !t.conflicts(candidate) {
			current = candidate
		} else {
			t.fix(current)
			t.reduce()
			return
		}
	}
}

// Size implements Tower
func (t *tower) Size() int {
	return t.height
}

func (t *tower) String() string {
	lines := ""
	for y := t.height; y >= t.base; y-- {
		lines += "|"
		for x := 0; x < t.width; x++ {
			_, found := t.fixed[shapes.Coord{X: x, Y: y}]
			if found {
				lines += "#"
			} else {
				lines += " "
			}
		}
		lines += fmt.Sprintf("|%v\n", y)
	}
	lines += "|"
	for x := 0; x < t.width; x++ {
		lines += "~"
	}
	lines += fmt.Sprintf("|%v\n", t.base-1)
	return lines
}

func (t *tower) fix(shape shapes.Rock) {
	for _, stone := range shape.GetStones() {
		if stone.Y >= t.height {
			t.height = stone.Y + 1
		}
		t.fixed[stone] = *new(interface{})
	}
}

func (t *tower) conflicts(shape shapes.Rock) bool {
	for _, stone := range shape.GetStones() {
		_, overlap := t.fixed[stone]
		switch {
		case stone.X < 0:
			return true
		case stone.X >= t.width:
			return true
		case stone.Y < t.base:
			return true
		case overlap:
			return true
		}
	}
	return false
}
