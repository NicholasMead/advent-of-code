package elfs

import "testing"

func TestRun(t *testing.T) {
	elfs := []Coord{
		{0, 0},
		{1, 0},
		{0, -1},
		{0, -3},
		{1, -3},
	}

	result, _ := Run(elfs, 10)

	minX, maxX, minY, maxY := Bounds(result)
	area := (1 + maxX - minX) * (1 + maxY - minY)
	score := area - len(elfs)
	if score != 25 {
		t.Error("Expected 25 got", result)
	}
}
