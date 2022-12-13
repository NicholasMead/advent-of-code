package grid

import (
	"fmt"
	"math"
)

type Node struct {
	x, y, h int
	grid    *grid
}

func (this Node) GetNeighbours() []Node {
	neighbours := make([]Node, 0, 4)
	for _, transform := range []struct{ x, y int }{
		{-1, 0},
		{+1, 0},
		{0, -1},
		{0, +1},
	} {
		x, y := this.x+transform.x, this.y+transform.y
		if neighbour, err := this.grid.GetNode(x, y); err == nil {
			if diff(this, neighbour) < 2 {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
}

func (this Node) GetNeighbourDist(n Node) int {
	return this.GetHuristic(n)
}

func (from Node) GetHuristic(to Node) int {
	return abs(from.x-to.x) + abs(from.y-to.y)
}

func diff(a, b Node) int {
	return b.h - a.h
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func (n Node) Hash() uint64 {
	return (uint64(n.x) << 32) | uint64(n.y)
}

func (n Node) String() string {
	return fmt.Sprintf("(%d, %d, %d)", n.x, n.y, n.h)
}
