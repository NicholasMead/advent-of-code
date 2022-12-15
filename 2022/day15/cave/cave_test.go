package cave

import (
	"aoc/day15/coord"
	"testing"
)

func expect(t *testing.T, cave Cave, x, y int, mat Material) {
	if m := cave.AtCoord(x, y); m != mat {
		t.Errorf("Expected %c found %c at (%d,%d)", mat, m, x, y)
	}
}

func TestCave(t *testing.T) {
	cave := Create()

	t.Run("At", func(t *testing.T) {
		cave := cave.Copy()
		cave.Add(coord.Coord{X: 1, Y: 1}, Empty)

		expect(t, cave, 1, 1, Empty)
		expect(t, cave, 2, 2, Unknown)
	})

	t.Run("AddArea", func(t *testing.T) {
		cave := cave.Copy()
		pos := coord.Coord{X: 1, Y: 1}

		cave.AddArea(pos, 3, Empty)

		for _, pos := range []struct {
			x, y int
			mat  Material
		}{
			{1, 1, Empty},
			{1, 4, Empty},
			{1, 5, Unknown},
			{4, 1, Empty},
			{5, 1, Unknown},
			{-2, 1, Empty},
			{-3, 1, Unknown},
			{1, -2, Empty},
			{1, -3, Unknown},
			{2, 3, Empty},
			{3, 3, Unknown},
		} {
			expect(t, cave, pos.x, pos.y, pos.mat)
		}
	})
}
