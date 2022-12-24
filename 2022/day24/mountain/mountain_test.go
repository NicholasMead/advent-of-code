package mountain

import "testing"

func TestBlizzard(t *testing.T) {
	b := Blizzard{Coord{0, 0}, Coord{1, 0}}

	b.Next()
	b.Next()

	if b.Pos.X != 2 {
		t.Error("Extected x = 2, got x =", b.Pos.X)
	}
}
