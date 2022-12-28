package grid

import (
	"strconv"
	"testing"
)

var sampleInput = []string{
	"        ...#    ",
	"        .#..    ",
	"        #...    ",
	"        ....    ",
	"...#.......#    ",
	"........#...    ",
	"..#....#....    ",
	"..........#.    ",
	"        ...#....",
	"        .....#..",
	"        .#......",
	"        ......#.",
}

func TestFlat(t *testing.T) {
	flat := Flat(buildMock(sampleInput))

	tests := []struct {
		from, to Vector
		dir      Direction
	}{
		{Vector{11, 3}, Vector{8, 3}, Right},
		{Vector{2, 4}, Vector{2, 7}, Up},
		{Vector{8, 8}, Vector{15, 8}, Left},
		{Vector{10, 11}, Vector{10, 0}, Down},

		{Vector{10, 0}, Vector{10, 0}, Right},
		{Vector{0, 4}, Vector{0, 4}, Left},
		{Vector{7, 5}, Vector{7, 5}, Down},
		{Vector{7, 5}, Vector{7, 5}, Right},
	}

	for _, test := range tests {
		t.Run("Move", func(t *testing.T) {
			if res, _ := flat.Move(test.from, test.dir); res != test.to {
				t.Fail()
				t.Log(res)
			}
		})
	}

	t.Run("Part1", func(t *testing.T) {
		pos, dir := Vector{8, 0}, Right
		for _, i := range []string{"10", "R", "5", "L", "5", "R", "10", "L", "4", "R", "5", "L", "5"} {
			count, err := strconv.Atoi(i)

			if err == nil {
				for n := 0; n < count; n++ {
					pos, dir = flat.Move(pos, dir)
				}
			} else {
				r := Rotation(i)
				dir = dir.Rotate(r)
			}
		}
		if pos[0] != 7 {
			t.Errorf("Expected x=%v got x=%v", 7, pos[0])
		}
		if pos[1] != 5 {
			t.Errorf("Expected x=%v got x=%v", 5, pos[1])
		}
		if dir != Right {
			t.Errorf("Expected d=%v got d=%v", Right, dir)
		}
		if score := Score(pos, dir); score != 6032 {
			t.Errorf("Expected s=%v got s=%v", 6032, score)
		}
	})
}

func buildMock(input []string) Surface {
	s := Surface{}
	for y, line := range input {
		for x, r := range line {
			if r != ' ' {
				s[Vector{x, y}] = Tile(r)
			}
		}
	}
	return s
}
