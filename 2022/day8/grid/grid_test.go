package grid

import "testing"

type Case struct {
	x, y int
	v    bool
}

func toBool(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}

func TestSample(t *testing.T) {
	grid := NewGrid([][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	})

	cases := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
	}

	for x, row := range cases {
		for y := range row {
			v := grid.IsVisible(x, y)
			t.Log(grid.GetSize(x, y), v)
			if v != toBool(cases[x][y]) {
				t.Error(
					"Expected", toBool(cases[x][y]), "got", v,
					"for", x, y, "height:", grid.GetSize(x, y))
			}
		}
	}
}

func TestTrivial(t *testing.T) {
	grid := NewGrid([][]int{
		{3, 1, 3, 7, 3},
		{2, 0, 0, 0, 2},
		{6, 0, 0, 0, 2},
		{3, 0, 0, 0, 9},
		{3, 5, 3, 9, 0},
	})

	cases := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}

	for x, row := range cases {
		for y := range row {
			v := grid.IsVisible(x, y)
			t.Log(grid.GetSize(x, y), v)
			if v != toBool(cases[x][y]) {
				t.Error(
					"Expected", toBool(cases[x][y]), "got", v,
					"for", x, y, "height:", grid.GetSize(x, y))
			}
		}
	}
}
