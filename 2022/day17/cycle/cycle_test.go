package cycle

import "testing"

func TestCycle(t *testing.T) {

	c := CreateCycle(0, 1, 2, 3, 4, 5)

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			next, pos := c.Next()

			if j != next {
				t.Errorf("%v: Expected next %v got %v", i, j, next)
				t.FailNow()
			}
			if j != pos {
				t.Errorf("%v: Expected pos %v got %v", i, j, pos)
				t.FailNow()
			}
		}
	}
}
