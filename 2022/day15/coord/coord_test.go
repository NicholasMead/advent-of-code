package coord

import (
	"fmt"
	"testing"
)

func TestCoord(t *testing.T) {
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
			name := fmt.Sprintf("%c=>(%d,%d)", c.direction, c.result.X, c.result.Y)
			t.Run(name, func(t *testing.T) {
				if r := start.StepDirection(c.direction); r != c.result {
					t.Error("Expected", c.result, "got", r)
				}
			})
		}
	})

	t.Run("Trace", func(t *testing.T) {
		fmt.Println("trace")
		c := Coord{}

		trace := c.GetTrace(2)

		t.Log(trace)
		if len(trace) != 24 {
			t.Error("Lendth =", len(trace))
		}
	})
}
