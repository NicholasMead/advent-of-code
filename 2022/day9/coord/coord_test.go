package coord

import (
	"fmt"
	"testing"
)

func TestCoord(t *testing.T) {
	t.Run("StepTorwards", func(t *testing.T) {
		type testCase struct {
			input, result Coord
		}

		start := Coord{1, 1}
		cases := []testCase{
			//Stand Still
			{Coord{1, 1}, Coord{1, 1}},
			{Coord{1, 2}, Coord{1, 1}},
			{Coord{2, 1}, Coord{1, 1}},
			{Coord{2, 2}, Coord{1, 1}},
			{Coord{0, 0}, Coord{1, 1}},
			{Coord{0, 1}, Coord{1, 1}},

			//Move left, up, right, down
			{Coord{-1, +1}, Coord{0, 1}},
			{Coord{+3, +1}, Coord{2, 1}},
			{Coord{+1, -1}, Coord{1, 0}},
			{Coord{+1, +3}, Coord{1, 2}},

			//Move diagonal
			{Coord{+2, +3}, Coord{2, 2}},
			{Coord{+3, +2}, Coord{2, 2}},
			{Coord{+0, -1}, Coord{0, 0}},
			{Coord{-1, +0}, Coord{0, 0}},
		}

		for _, c := range cases {
			t.Run(fmt.Sprint(c), func(t *testing.T) {
				if r := start.StepTowards(c.input); r != c.result {
					t.Error("Expected", c.result, "got", r)
				}
			})
		}
	})
}
