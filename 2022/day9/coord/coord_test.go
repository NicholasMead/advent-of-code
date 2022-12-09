package coord

import (
	"fmt"
	"testing"
)

func TestCoord(t *testing.T) {
	start := Coord{1, 1}

	t.Run("StepTorwards", func(t *testing.T) {
		t.Run("Moves", func(t *testing.T) {
			type moveCase struct {
				target, result Coord
			}

			cases := []moveCase{
				//Stand Still
				{Coord{1, 1}, Coord{1, 1}},
				{Coord{1, 2}, Coord{1, 1}},
				{Coord{2, 1}, Coord{1, 1}},
				{Coord{2, 2}, Coord{1, 1}},
				{Coord{0, 0}, Coord{1, 1}},
				{Coord{0, 1}, Coord{1, 1}},

				//Move up, down, left, right
				{Coord{1, 3}, Coord{1, 2}},
				{Coord{1, -1}, Coord{1, 0}},
				{Coord{3, 1}, Coord{2, 1}},
				{Coord{-1, 1}, Coord{0, 1}},

				//Move diagonal
				{Coord{2, 3}, Coord{2, 2}},
				{Coord{0, -1}, Coord{0, 0}},
				{Coord{3, 2}, Coord{2, 2}},
				{Coord{-1, 0}, Coord{0, 0}},
			}

			for _, c := range cases {
				t.Run(fmt.Sprint(c), func(t *testing.T) {
					r := start.StepTowards(c.target)

					if r.x != c.result.x || r.y != c.result.y {
						t.Error("Expected", c.result, "got", r)
					}
				})
			}
		})
	})
}
