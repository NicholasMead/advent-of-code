package shapes

import "testing"

func TestPatten(t *testing.T) {

	t.Run("AboveY", func(t *testing.T) {
		box := Box
		target := Patten{{0, 0}, {1, 0}}

		result := box.AboveY(1)

		if !target.Match(result) {
			t.Error("Expected match")
			t.Log("t:", target)
			t.Log("r:", result)
		}
	})

}
