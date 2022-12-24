package mountain

import "testing"

func TestPath(t *testing.T) {

	start := Coord{0, -1}
	end := Coord{5, 4}
	m := Mountain{
		Height:    4,
		Width:     6,
		Blizzards: make([]Blizzard, 0),
	}

	t.Run("Empty Mountain", func(t *testing.T) {
		result := FindShortestPath(m, start, end, 0)
		if result != 10 {
			t.Error("Expected 10 got ", result)
		}
	})

	t.Run("Sample Mountain", func(t *testing.T) {
		raw := [][]rune{
			{'>', '>', '.', '<', '^', '<'},
			{'.', '<', '.', '.', '<', '<'},
			{'>', 'v', '.', '>', '<', '>'},
			{'<', '^', 'v', '^', '^', '>'},
		}
		m = ParseMountain(raw)

		result := FindShortestPath(m, start, end, 0)
		if result != 18 {
			t.Error("Expected 18 got ", result)
		}
	})

	t.Run("Sample Mountain + snacks", func(t *testing.T) {
		raw := [][]rune{
			{'>', '>', '.', '<', '^', '<'},
			{'.', '<', '.', '.', '<', '<'},
			{'>', 'v', '.', '>', '<', '>'},
			{'<', '^', 'v', '^', '^', '>'},
		}
		m = ParseMountain(raw)

		p1 := FindShortestPath(m, start, end, 0)
		p2 := FindShortestPath(m, end, start, p1)
		p3 := FindShortestPath(m, start, end, p2)
		if p3 != 54 {
			t.Error("Expected 23 got ", 54)
		}
	})

}
