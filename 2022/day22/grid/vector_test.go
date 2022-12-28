package grid

import "testing"

func TestVector(t *testing.T) {

	t.Run("IsComparable", func(t *testing.T) {
		a := Vector{3, 4}
		b := Vector{3, 4}

		if a != b {
			t.Fail()
		}
	})

	t.Run("DoesMap", func(t *testing.T) {
		mapped := map[Vector]int{
			{1, 2}:   3,
			{-1, -2}: 3,
		}
		if mapped[Vector{1, 2}] != mapped[Vector{-1, -2}] {
			t.Fail()
			t.Log(Vector{1, 2})
			t.Log(Vector{-1, -2})
		}
	})
}
