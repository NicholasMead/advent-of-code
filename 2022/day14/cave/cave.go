package cave

import (
	"aoc/day14/coord"
	"strings"
)

type Material rune

const (
	Air, Rock, Sand Material = '.', '#', 'o'
)

type Cave interface {
	Copy() Cave
	AtCoord(x, y int) Material
	At(coord coord.Coord) Material
	Add(at coord.Coord, mat Material)
	AddRange(from, to coord.Coord, mat Material)
	Drop(from coord.Coord, mat Material) (landsAt coord.Coord, isFloor bool)
	Count(mat Material) int
}

func Create() Cave {
	return &cave{map[coord.Coord]Material{}, 0}
}

type cave struct {
	positions map[coord.Coord]Material
	floor     int
}

// Count implements Cave
func (c *cave) Count(material Material) int {
	count := 0
	for _, mat := range c.positions {
		if mat == material {
			count++
		}
	}
	return count
}

func (c *cave) Copy() Cave {
	copy := Create()

	for pos, mat := range c.positions {
		copy.Add(pos, mat)
	}

	return copy
}

func (c *cave) AtCoord(x, y int) Material {
	coord := coord.Coord{X: x, Y: y}
	return c.At(coord)
}

func (c *cave) At(coord coord.Coord) Material {
	m, f := c.positions[coord]

	if f {
		return m
	} else {
		return Air
	}
}

func (c *cave) Add(at coord.Coord, mat Material) {
	c.positions[at] = mat

	if mat == Rock && at.Y+2 > c.floor {
		c.floor = at.Y + 2
	}
}

func (c *cave) AddRange(from, to coord.Coord, mat Material) {
	pos := from

	for {
		c.Add(pos, mat)

		if pos == to {
			return
		} else {
			pos = pos.StepTowards(to)
		}
	}
}

func (c *cave) Drop(from coord.Coord, mat Material) (landsAt coord.Coord, isFloor bool) {
	if c.isOnFloor(from) {
		c.Add(from, mat)
		return from, true
	}

	for _, next := range []coord.Coord{
		from.StepDirection(coord.Down),
		from.StepDirection(coord.Down).StepDirection(coord.Left),
		from.StepDirection(coord.Down).StepDirection(coord.Right),
	} {
		if m := c.At(next); m == Air {
			return c.Drop(next, mat)
		}
	}

	c.Add(from, mat)
	return from, false
}

func (c *cave) isOnFloor(from coord.Coord) bool {
	return from.Y+1 == c.floor
}

func (c cave) String() string {
	var xMin, xMax int
	var yMin, yMax int = 0, c.floor
	first := true

	for pos := range c.positions {
		if first {
			xMin, xMax = pos.X, pos.X
			first = false
		} else {
			if pos.X < xMin {
				xMin = pos.X
			}
			if pos.X > xMax {
				xMax = pos.X
			}
			if pos.Y > yMax {
				yMax = pos.Y
			}
		}
	}

	xMin -= 2
	xMax += 2
	gridRows := []string{}
	for y := yMin; y <= yMax; y++ {
		row := ""
		for x := xMin; x <= xMax; x++ {
			if y == yMax {
				row += "#"
			} else {
				row += string(c.At(coord.Coord{x, y}))
			}
		}
		gridRows = append(gridRows, row)
	}

	return strings.Join(gridRows, "\n")
}
