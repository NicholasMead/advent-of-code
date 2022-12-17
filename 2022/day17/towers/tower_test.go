package towers

import (
	"aoc/day17/cycle"
	"aoc/day17/shapes"
	"testing"
)

func TestTower(t *testing.T) {
	t.Run("Falls Straight", func(t *testing.T) {
		jets := cycle.CreateCycle(0)
		tower := CreateTower(7)

		tower.Drop(shapes.Box, jets)

		if tower.Size() != 2 {
			t.Errorf("Expected size of 2 got %v", tower.Size())
		}
	})

	t.Run("Rocks Stack", func(t *testing.T) {
		jets := cycle.CreateCycle(0)
		tower := CreateTower(7)

		tower.Drop(shapes.Box, jets)
		tower.Drop(shapes.Box, jets)

		if tower.Size() != 4 {
			t.Errorf("Expected size of 2 got %v", tower.Size())
		}
	})
}
