package cave

import (
	"aoc/day14/coord"
	"testing"
)

func expect(t *testing.T, cave Cave, x, y int, mat Material) {
	if m := cave.AtCoord(x, y); m != mat {
		t.Errorf("Expected %c found %c", mat, m)
	}
}

func TestCave(t *testing.T) {
	cave := Create()

	t.Run("At", func(t *testing.T) {
		cave := cave.Copy()
		cave.Add(coord.Coord{X: 1, Y: 1}, Rock)

		expect(t, cave, 1, 1, Rock)
		expect(t, cave, 2, 2, Air)
	})

	t.Run("Add", func(t *testing.T) {
		cave := cave.Copy()
		from, to := coord.Coord{X: 1, Y: 1}, coord.Coord{X: 1, Y: 3}

		cave.AddRange(from, to, Rock)

		for i := 1; i <= 3; i++ {
			expect(t, cave, 1, i, Rock)
		}
		expect(t, cave, 1, 0, Air)
		expect(t, cave, 1, 4, Air)
		expect(t, cave, 2, 2, Air)
	})

	t.Run("Drop", func(t *testing.T) {
		cave := cave.Copy()
		for _, line := range []struct {
			startx, starty, endx, endy int
			mat                        Material
		}{
			{498, 4, 498, 6, Rock},
			{498, 6, 496, 6, Rock},

			{530, 4, 502, 4, Rock},
			{502, 4, 502, 9, Rock},
			{502, 9, 494, 9, Rock},
		} {
			cave.AddRange(
				coord.Coord{X: line.startx, Y: line.starty},
				coord.Coord{X: line.endx, Y: line.endy},
				line.mat)
		}

		t.Run("Lands", func(t *testing.T) {
			from := []struct {
				fromX, fromY int
				landX, landY int
			}{
				{500, 0, 500, 8},
				{500, 0, 499, 8},
				{500, 0, 501, 8},
				{500, 0, 500, 7},
			}

			for _, from := range from {
				cave.Drop(coord.Coord{from.fromX, from.fromY}, Sand)
				expect(t, cave, from.landX, from.landY, Sand)
			}
		})

		t.Run("DetectsFloor", func(t *testing.T) {
			from := []struct {
				fromX, fromY int
				lands        bool
			}{
				{500, 0, false},
				{600, 0, true},
			}

			for _, from := range from {
				_, landed := cave.Drop(coord.Coord{from.fromX, from.fromY}, Sand)
				if landed != from.lands {
					t.Error("Expected lands=", from.lands, "got", landed)
				}
			}
		})
	})

	t.Run("Sample", func(t *testing.T) {
		cave := cave.Copy()

		for _, line := range []struct {
			startx, starty, endx, endy int
			mat                        Material
		}{
			{498, 4, 498, 6, Rock},
			{498, 6, 496, 6, Rock},

			{530, 4, 502, 4, Rock},
			{502, 4, 502, 9, Rock},
			{502, 9, 494, 9, Rock},
		} {
			cave.AddRange(
				coord.Coord{X: line.startx, Y: line.starty},
				coord.Coord{X: line.endx, Y: line.endy},
				line.mat)
		}

		for i := 0; i < 24; i++ {
			_, landed := cave.Drop(coord.Coord{X: 500, Y: 0}, Sand)
			if !landed {
				t.Errorf("Expected sand #%d to land", i)
			}
		}

		_, landed := cave.Drop(coord.Coord{X: 500, Y: 0}, Sand)
		if landed {
			t.Errorf("Expected sand #%d to fall forever", 25)
		}
	})
}
