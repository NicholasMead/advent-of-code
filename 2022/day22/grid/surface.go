package grid

import "math"

type Surface map[Vector]Tile

func (surface Surface) ColumnBounds(c int) (int, int) {
	lower, upper := math.MaxInt, math.MinInt

	for v := range surface {
		if v[0] == c {
			lower = min(lower, v[1])
			upper = max(upper, v[1])
		}
	}

	return lower, upper
}

func (surface Surface) RowBounds(r int) (int, int) {
	lower, upper := math.MaxInt, math.MinInt

	for v := range surface {
		if v[1] == r {
			lower = min(lower, v[0])
			upper = max(upper, v[0])
		}
	}

	return lower, upper
}

func (surface Surface) String() (s string) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for c := range surface {
		minX = min(minX, c[0])
		minY = min(minY, c[1])

		maxX = max(maxX, c[0])
		maxY = max(maxY, c[1])
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			tile, found := surface[Vector{x, y}]
			if found {
				s += string(tile)
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return s
}

func (surface Surface) inBounds(v Vector) bool {
	_, found := surface[v]
	return found
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
