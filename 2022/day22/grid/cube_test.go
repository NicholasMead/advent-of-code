package grid

import "testing"

var inputMock = []string{
	"    ........",
	"    ........",
	"    ........",
	"    ........",
	"    ....    ",
	"    ....    ",
	"    ....    ",
	"    ....    ",
	"........    ",
	"........    ",
	"........    ",
	"........    ",
	"....        ",
	"....        ",
	"....        ",
	"....        ",
}

func TestCube(t *testing.T) {
	cube := Cube(buildMock(inputMock), 4)

	tests := []struct {
		from, to       Vector
		fromDir, toDir Direction
		description    string
	}{
		{Vector{4, 1}, Vector{0, 10}, Left, Right, "A>D"},
		{Vector{5, 0}, Vector{0, 13}, Up, Right, "A>F"},

		{Vector{9, 0}, Vector{1, 15}, Up, Up, "B>F"},
		{Vector{11, 1}, Vector{7, 10}, Right, Left, "B>E"},
		{Vector{9, 3}, Vector{7, 5}, Down, Left, "B>C"},

		{Vector{4, 5}, Vector{1, 8}, Left, Down, "C>D"},
		{Vector{7, 5}, Vector{9, 3}, Right, Up, "C>B"},
	}

	for _, test := range tests {
		t.Run("Move:"+test.description, func(t *testing.T) {
			pos, dir := cube.Move(test.from, test.fromDir)

			if pos != test.to {
				t.Fail()
				t.Log(pos)
			}
			if dir != test.toDir {
				t.Fail()
				t.Log(dir)
			}
		})
	}
}
