package cave

import (
	"aoc/day15/coord"
	"strings"
)

type Material rune

const (
	Unknown, Empty, Sensor, Beacon Material = '.', '#', 'S', 'B'
)

type Cave interface {
	Copy() Cave
	AtCoord(x, y int) Material
	At(coord coord.Coord) Material
	Add(at coord.Coord, mat Material)
	AddArea(from coord.Coord, dist int, mat Material)
	Count(mat Material) int
	Row(index int) []Material
}

func Create() Cave {
	return &cave{map[coord.Coord]Material{}}
}

type cave struct {
	positions map[coord.Coord]Material
}

// Row implements Cave
func (c *cave) Row(index int) []Material {
	materials := []Material{}
	for pos, mat := range c.positions {
		if pos.Y == index {
			materials = append(materials, mat)
		}
	}
	return materials
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
		return Unknown
	}
}

func (c *cave) Add(at coord.Coord, mat Material) {
	c.positions[at] = mat
}

func (c *cave) AddArea(from coord.Coord, dist int, mat Material) {
	for x := from.X - dist; x <= from.X+dist; x++ {
		for y := from.Y - dist; y <= from.Y+dist; y++ {
			pos := coord.Coord{x, y}

			if coord.Dist(from, pos) <= dist {
				c.Add(pos, mat)
			}
		}
	}
}

func (c cave) String() string {
	var xMin, xMax int
	var yMin, yMax int
	first := true

	for pos := range c.positions {
		if first {
			xMin, xMax = pos.X, pos.X
			yMin, yMax = pos.Y, pos.Y
			first = false
		} else {
			if pos.X < xMin {
				xMin = pos.X
			}
			if pos.X > xMax {
				xMax = pos.X
			}
			if pos.Y < yMin {
				yMin = pos.Y
			}
			if pos.Y > yMax {
				yMax = pos.Y
			}
		}
	}

	xMin -= 1
	yMin -= 1
	xMax += 1
	yMax += 1

	gridRows := []string{}
	for y := yMin; y <= yMax; y++ {
		row := ""
		for x := xMin; x <= xMax; x++ {
			mat := c.At(coord.Coord{x, y})
			row += string(mat)
		}
		gridRows = append(gridRows, row)
	}

	return strings.Join(gridRows, "\n")
}
