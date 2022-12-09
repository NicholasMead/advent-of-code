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
			name := fmt.Sprintf("(%d,%d)=>(%d,%d)", c.input.x, c.input.y, c.result.x, c.result.y)
			t.Run(name, func(t *testing.T) {
				if r := start.StepTowards(c.input); r != c.result {
					t.Error("Expected", c.result, "got", r)
				}
			})
		}
	})

	t.Run("StepDirection", func(t *testing.T) {
		type testCase struct {
			direction Direction
			result    Coord
		}

		start := Coord{1, 1}
		cases := []testCase{
			{'U', Coord{1, 2}},
			{'D', Coord{1, 0}},
			{'L', Coord{0, 1}},
			{'R', Coord{2, 1}},
		}

		for _, c := range cases {
			name := fmt.Sprintf("%c=>(%d,%d)", c.direction, c.result.x, c.result.y)
			t.Run(name, func(t *testing.T) {
				if r := start.StepDirection(c.direction); r != c.result {
					t.Error("Expected", c.result, "got", r)
				}
			})
		}
	})
}
